// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package burrowmint

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/hyperledger/burrow/account"
	core_types "github.com/hyperledger/burrow/core/types"
	"github.com/hyperledger/burrow/event"
	"github.com/hyperledger/burrow/manager/burrow-mint/evm"
	"github.com/hyperledger/burrow/manager/burrow-mint/state"
	"github.com/hyperledger/burrow/txs"
	"github.com/hyperledger/burrow/word256"

	"github.com/tendermint/go-crypto"
	tEvents "github.com/tendermint/go-events"
)

// Transactor is part of the pipe for BurrowMint and provides the implementation
// for the pipe to call into the BurrowMint application
type transactor struct {
	chainID       string
	eventSwitch   tEvents.Fireable
	burrowMint    *BurrowMint
	eventEmitter  event.EventEmitter
	txMtx         *sync.Mutex
	txBroadcaster func(tx txs.Tx) error
}

func newTransactor(chainID string, eventSwitch tEvents.Fireable,
	burrowMint *BurrowMint, eventEmitter event.EventEmitter,
	txBroadcaster func(tx txs.Tx) error) *transactor {
	return &transactor{
		chainID,
		eventSwitch,
		burrowMint,
		eventEmitter,
		&sync.Mutex{},
		txBroadcaster,
	}
}

// Run a contract's code on an isolated and unpersisted state
// Cannot be used to create new contracts
// NOTE: this function is used from 1337 and has sibling on 46657
// in pipe.go
// TODO: [ben] resolve incompatibilities in byte representation for 0.12.0 release
func (this *transactor) Call(fromAddress, toAddress, data []byte) (
	*core_types.Call, error) {

	st := this.burrowMint.GetState()
	cache := state.NewBlockCache(st) // XXX: DON'T MUTATE THIS CACHE (used internally for CheckTx)
	outAcc := cache.GetAccount(toAddress)
	if outAcc == nil {
		return nil, fmt.Errorf("Account %X does not exist", toAddress)
	}
	if fromAddress == nil {
		fromAddress = []byte{}
	}
	callee := toVMAccount(outAcc)
	caller := &vm.Account{Address: word256.LeftPadWord256(fromAddress)}
	txCache := state.NewTxCache(cache)
	gasLimit := st.GetGasLimit()
	params := vm.Params{
		BlockHeight: int64(st.LastBlockHeight),
		BlockHash:   word256.LeftPadWord256(st.LastBlockHash),
		BlockTime:   st.LastBlockTime.Unix(),
		GasLimit:    gasLimit,
	}

	vmach := vm.NewVM(txCache, vm.DefaultDynamicMemoryProvider, params,
		caller.Address, nil)
	vmach.SetFireable(this.eventSwitch)
	gas := gasLimit
	ret, err := vmach.Call(caller, callee, callee.Code, data, 0, &gas)
	if err != nil {
		return nil, err
	}
	gasUsed := gasLimit - gas
	// here return bytes are hex encoded; on the sibling function
	// they are not
	return &core_types.Call{Return: hex.EncodeToString(ret), GasUsed: gasUsed}, nil
}

// Run the given code on an isolated and unpersisted state
// Cannot be used to create new contracts.
func (this *transactor) CallCode(fromAddress, code, data []byte) (
	*core_types.Call, error) {
	if fromAddress == nil {
		fromAddress = []byte{}
	}
	cache := this.burrowMint.GetCheckCache() // XXX: DON'T MUTATE THIS CACHE (used internally for CheckTx)
	callee := &vm.Account{Address: word256.LeftPadWord256(fromAddress)}
	caller := &vm.Account{Address: word256.LeftPadWord256(fromAddress)}
	txCache := state.NewTxCache(cache)
	st := this.burrowMint.GetState() // for block height, time
	gasLimit := st.GetGasLimit()
	params := vm.Params{
		BlockHeight: int64(st.LastBlockHeight),
		BlockHash:   word256.LeftPadWord256(st.LastBlockHash),
		BlockTime:   st.LastBlockTime.Unix(),
		GasLimit:    gasLimit,
	}

	vmach := vm.NewVM(txCache, vm.DefaultDynamicMemoryProvider, params,
		caller.Address, nil)
	gas := gasLimit
	ret, err := vmach.Call(caller, callee, code, data, 0, &gas)
	if err != nil {
		return nil, err
	}
	gasUsed := gasLimit - gas
	// here return bytes are hex encoded; on the sibling function
	// they are not
	return &core_types.Call{Return: hex.EncodeToString(ret), GasUsed: gasUsed}, nil
}

// Broadcast a transaction.
func (this *transactor) BroadcastTx(tx txs.Tx) (*txs.Receipt, error) {
	err := this.txBroadcaster(tx)

	if err != nil {
		return nil, fmt.Errorf("Error broadcasting transaction: %v", err)
	}

	txHash := txs.TxHash(this.chainID, tx)
	var createsContract uint8
	var contractAddr []byte
	// check if creates new contract
	if callTx, ok := tx.(*txs.CallTx); ok {
		if len(callTx.Address) == 0 {
			createsContract = 1
			contractAddr = state.NewContractAddress(callTx.Input.Address, callTx.Input.Sequence)
		}
	}
	return &txs.Receipt{txHash, createsContract, contractAddr}, nil
}

// Orders calls to BroadcastTx using lock (waits for response from core before releasing)
func (this *transactor) Transact(privKey, address, data []byte, gasLimit,
	fee int64) (*txs.Receipt, error) {
	var addr []byte
	if len(address) == 0 {
		addr = nil
	} else if len(address) != 20 {
		return nil, fmt.Errorf("Address is not of the right length: %d\n", len(address))
	} else {
		addr = address
	}
	if len(privKey) != 64 {
		return nil, fmt.Errorf("Private key is not of the right length: %d\n", len(privKey))
	}
	this.txMtx.Lock()
	defer this.txMtx.Unlock()
	pa := account.GenPrivAccountFromPrivKeyBytes(privKey)
	cache := this.burrowMint.GetCheckCache() // XXX: DON'T MUTATE THIS CACHE (used internally for CheckTx)
	acc := cache.GetAccount(pa.Address)
	var sequence int
	if acc == nil {
		sequence = 1
	} else {
		sequence = acc.Sequence + 1
	}
	// TODO: [Silas] we should consider revising this method and removing fee, or
	// possibly adding an amount parameter. It is non-sensical to just be able to
	// set the fee. Our support of fees in general is questionable since at the
	// moment all we do is deduct the fee effectively leaking token. It is possible
	// someone may be using the sending of native token to payable functions but
	// they can be served by broadcasting a token.

	// We hard-code the amount to be equal to the fee which means the CallTx we
	// generate transfers 0 value, which is the most sensible default since in
	// recent solidity compilers the EVM generated will throw an error if value
	// is transferred to a non-payable function.
	txInput := &txs.TxInput{
		Address:  pa.Address,
		Amount:   fee,
		Sequence: sequence,
		PubKey:   pa.PubKey,
	}
	tx := &txs.CallTx{
		Input:    txInput,
		Address:  addr,
		GasLimit: gasLimit,
		Fee:      fee,
		Data:     data,
	}

	// Got ourselves a tx.
	txS, errS := this.SignTx(tx, []*account.PrivAccount{pa})
	if errS != nil {
		return nil, errS
	}
	return this.BroadcastTx(txS)
}

func (this *transactor) TransactAndHold(privKey, address, data []byte, gasLimit, fee int64) (*txs.EventDataCall, error) {
	rec, tErr := this.Transact(privKey, address, data, gasLimit, fee)
	if tErr != nil {
		return nil, tErr
	}
	var addr []byte
	if rec.CreatesContract == 1 {
		addr = rec.ContractAddr
	} else {
		addr = address
	}
	// We want non-blocking on the first event received (but buffer the value),
	// after which we want to block (and then discard the value - see below)
	wc := make(chan *txs.EventDataCall, 1)
	subId := fmt.Sprintf("%X", rec.TxHash)
	this.eventEmitter.Subscribe(subId, txs.EventStringAccCall(addr),
		func(evt txs.EventData) {
			eventDataCall := evt.(txs.EventDataCall)
			if bytes.Equal(eventDataCall.TxID, rec.TxHash) {
				// Beware the contract of go-events subscribe is that we must not be
				// blocking in an event callback when we try to unsubscribe!
				// We work around this by using a non-blocking send.
				select {
				// This is a non-blocking send, but since we are using a buffered
				// channel of size 1 we will always grab our first event even if we
				// haven't read from the channel at the time we receive the first event.
				case wc <- &eventDataCall:
				default:
				}
			}
		})

	timer := time.NewTimer(300 * time.Second)
	toChan := timer.C

	var ret *txs.EventDataCall
	var rErr error

	select {
	case <-toChan:
		rErr = fmt.Errorf("Transaction timed out. Hash: " + subId)
	case e := <-wc:
		timer.Stop()
		if e.Exception != "" {
			rErr = fmt.Errorf("Error when transacting: " + e.Exception)
		} else {
			ret = e
		}
	}
	this.eventEmitter.Unsubscribe(subId)
	return ret, rErr
}

func (this *transactor) Send(privKey, toAddress []byte,
	amount int64) (*txs.Receipt, error) {
	var toAddr []byte
	if len(toAddress) == 0 {
		toAddr = nil
	} else if len(toAddress) != 20 {
		return nil, fmt.Errorf("To-address is not of the right length: %d\n",
			len(toAddress))
	} else {
		toAddr = toAddress
	}

	if len(privKey) != 64 {
		return nil, fmt.Errorf("Private key is not of the right length: %d\n",
			len(privKey))
	}

	pk := &[64]byte{}
	copy(pk[:], privKey)
	this.txMtx.Lock()
	defer this.txMtx.Unlock()
	pa := account.GenPrivAccountFromPrivKeyBytes(privKey)
	cache := this.burrowMint.GetState()
	acc := cache.GetAccount(pa.Address)
	var sequence int
	if acc == nil {
		sequence = 1
	} else {
		sequence = acc.Sequence + 1
	}

	tx := txs.NewSendTx()

	txInput := &txs.TxInput{
		Address:  pa.Address,
		Amount:   amount,
		Sequence: sequence,
		PubKey:   pa.PubKey,
	}

	tx.Inputs = append(tx.Inputs, txInput)

	txOutput := &txs.TxOutput{toAddr, amount}

	tx.Outputs = append(tx.Outputs, txOutput)

	// Got ourselves a tx.
	txS, errS := this.SignTx(tx, []*account.PrivAccount{pa})
	if errS != nil {
		return nil, errS
	}
	return this.BroadcastTx(txS)
}

func (this *transactor) SendAndHold(privKey, toAddress []byte,
	amount int64) (*txs.Receipt, error) {
	rec, tErr := this.Send(privKey, toAddress, amount)
	if tErr != nil {
		return nil, tErr
	}

	wc := make(chan *txs.SendTx)
	subId := fmt.Sprintf("%X", rec.TxHash)

	this.eventEmitter.Subscribe(subId, txs.EventStringAccOutput(toAddress),
		func(evt txs.EventData) {
			event := evt.(txs.EventDataTx)
			tx := event.Tx.(*txs.SendTx)
			wc <- tx
		})

	timer := time.NewTimer(300 * time.Second)
	toChan := timer.C

	var rErr error

	pa := account.GenPrivAccountFromPrivKeyBytes(privKey)

	select {
	case <-toChan:
		rErr = fmt.Errorf("Transaction timed out. Hash: " + subId)
	case e := <-wc:
		if bytes.Equal(e.Inputs[0].Address, pa.Address) && e.Inputs[0].Amount == amount {
			timer.Stop()
			this.eventEmitter.Unsubscribe(subId)
			return rec, rErr
		}
	}
	return nil, rErr
}

func (this *transactor) TransactNameReg(privKey []byte, name, data string,
	amount, fee int64) (*txs.Receipt, error) {

	if len(privKey) != 64 {
		return nil, fmt.Errorf("Private key is not of the right length: %d\n", len(privKey))
	}
	this.txMtx.Lock()
	defer this.txMtx.Unlock()
	pa := account.GenPrivAccountFromPrivKeyBytes(privKey)
	cache := this.burrowMint.GetCheckCache() // XXX: DON'T MUTATE THIS CACHE (used internally for CheckTx)
	acc := cache.GetAccount(pa.Address)
	var sequence int
	if acc == nil {
		sequence = 1
	} else {
		sequence = acc.Sequence + 1
	}
	tx := txs.NewNameTxWithNonce(pa.PubKey, name, data, amount, fee, sequence)
	// Got ourselves a tx.
	txS, errS := this.SignTx(tx, []*account.PrivAccount{pa})
	if errS != nil {
		return nil, errS
	}
	return this.BroadcastTx(txS)
}

// Sign a transaction
func (this *transactor) SignTx(tx txs.Tx, privAccounts []*account.PrivAccount) (txs.Tx, error) {
	// more checks?

	for i, privAccount := range privAccounts {
		if privAccount == nil || privAccount.PrivKey == nil {
			return nil, fmt.Errorf("Invalid (empty) privAccount @%v", i)
		}
	}
	switch tx.(type) {
	case *txs.NameTx:
		nameTx := tx.(*txs.NameTx)
		nameTx.Input.PubKey = privAccounts[0].PubKey
		nameTx.Input.Signature = privAccounts[0].Sign(this.chainID, nameTx)
	case *txs.SendTx:
		sendTx := tx.(*txs.SendTx)
		for i, input := range sendTx.Inputs {
			input.PubKey = privAccounts[i].PubKey
			input.Signature = privAccounts[i].Sign(this.chainID, sendTx)
		}
	case *txs.CallTx:
		callTx := tx.(*txs.CallTx)
		callTx.Input.PubKey = privAccounts[0].PubKey
		callTx.Input.Signature = privAccounts[0].Sign(this.chainID, callTx)
	case *txs.BondTx:
		bondTx := tx.(*txs.BondTx)
		// the first privaccount corresponds to the BondTx pub key.
		// the rest to the inputs
		bondTx.Signature = privAccounts[0].Sign(this.chainID, bondTx).(crypto.SignatureEd25519)
		for i, input := range bondTx.Inputs {
			input.PubKey = privAccounts[i+1].PubKey
			input.Signature = privAccounts[i+1].Sign(this.chainID, bondTx)
		}
	case *txs.UnbondTx:
		unbondTx := tx.(*txs.UnbondTx)
		unbondTx.Signature = privAccounts[0].Sign(this.chainID, unbondTx).(crypto.SignatureEd25519)
	case *txs.RebondTx:
		rebondTx := tx.(*txs.RebondTx)
		rebondTx.Signature = privAccounts[0].Sign(this.chainID, rebondTx).(crypto.SignatureEd25519)
	default:
		return nil, fmt.Errorf("Object is not a proper transaction: %v\n", tx)
	}
	return tx, nil
}

// No idea what this does.
func toVMAccount(acc *account.Account) *vm.Account {
	return &vm.Account{
		Address: word256.LeftPadWord256(acc.Address),
		Balance: acc.Balance,
		Code:    acc.Code,
		Nonce:   int64(acc.Sequence),
		Other:   acc.PubKey,
	}
}
