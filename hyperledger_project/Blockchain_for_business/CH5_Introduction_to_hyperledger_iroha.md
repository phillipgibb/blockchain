### Hyperledger Iroha

Hyperledger Iroha is a blockchain framework. It is designed to be simple and easy to be integrated into existing infastructure projects requiring DLT.

Key features include:
- Simple construction
- Modern
- Domain-driven C++ design
- Emphasis on mobile application development
- YAC consensus algorithm

### Architecture Overview

There are four layers to Hyperledger Iroha
- API
- Peer Interaction
- Chain Business Logic
- Storage

Components
- Model classes are system entities
- Torii(gate) provides the input and output interfaces for clients. It is a single gRPC server that is clients to interact with peers through the network. The client's RPC call is non-blocking, making Torii an asynchronous server. Both commands (transaction) and queries are performed through this interface.
- Network encompasses interaction with the network of peers.
- Consensus is in charge of peers agreeing on chain content in the network. The consensus mechanism used by Iroha is YAC which is a practical byzantine fault-tolerant algorithm based on voting for block hash.
- Simulator generates a temporary snapshot of storage to validate transactions by executing them against this snapshot and forming a verified proposal, which consists only of valid transactions.
- Validator classes check business rules and validity(correct format) of transaction or queries. There are two distint types of validation that occur in Hyperledger Iroha"
    - Stateless validation is a quicker form of validation, that performs schema and signature checks of the transaction.
    - Stateful validation is a slower form of validation, that checks the permissions and the current world state view, which is the latest and most actual state of the chain, to see if desired business rules and policites are possible. For example, does an account have enough funds to transfer?
- Synchronizer helps to synchronize new peers in the system or temp disconnected peers.
- Ametsuchi is the ledger block storage which consists of block index(currently Redis), block store (currently flat files),and a world state view component (currently PostgreSQL).

### Participants within the Network

There are three main participants:

- Clients
    - Query data that they have access/permission to
    - Perform a state-changing action transaction, which consists of atomic operations, called commands. For example, in a single transaction, a user can transfer funds to three people. But if he/she does not have enough funds to cover for all, the whole transaction will be rejected.
- Peers
    - Maintain the current state and their own copy of the shared ledger. A peer is a single entity in the network, and has an address, identity, and trust. Hyperledger Iroha is designed so that a single peer may be a single computer or scaled for a cluster, meaning different computers are used for ledger storage, indices, validation, and peer to peer logic.
- Ordering service
    - Orders transactions into a known order. There are a few options for the algorithm used by the ordering service. Kafka is considered a good candidate. It is important to mention that is Kafka, or any other distributed solution is used, that is be clustered; otherwise, this will result in a single point of failure.

### Transaction flow basics

- Step 1: A client creates and sends a transaction to the Torii gate, which routes the transaction to a peer that is responsible for performing stateless validation.

- Step 2: After the peer performs stateless validation, the transaction is first sent to the ordering gate, which is responsible for choosing the right strategy of connection to the ordering service.

- Step 3: The ordering service puts transactions into order and forwards them to peers in the consensus network in the form of proposals. A proposal is an unsigned block shared by the ordering service, that contains a batch of ordered transactions. Proposals are only forwarded when the ordering service has accumulated enough transactions, or a certain amount of time has elapsed since the last proposal. This prevents the ordering service from sending empty proposals.

- Step 4: Each peer verifies the proposal's contents (stateful validation) in the Simulator and creates a block which consists only of verified transaction. This block is then sent to the consensus gate which performs YAC consensus logic.

- Step 5: An ordered list of peers is determined, and a leader is elected based on the YAC consensus logic. Each peer casts a vote by signing and sending their proposed block to the leader.

- Step 6: If the leader receives enough signed proposed blocks (more than 2/3 of the peers) then it starts to send a commit message, indicating that this block should be applied to the chain of each peer participating in the consensus. Once the commit message has been sent, the proposed block becomes the next block in the chain of every peer via the synchronizer.