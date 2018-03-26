### Distributed Ledger Technologies (DLT)

Distributed ledger is a type of data structure which resides across multiple computer devices, generally spread across locations or regions.

Distributed Ledger Technology includes blockchain technologies and smart contracts. 

While Distributed ledgers existed prior to Bitcoin, the cryptography and shared computational power, along with a new consensus algorithm.

Distributed ledger technology consists of three basic components:

- Data Model that captures the current state of the ledger.
- Language of transactions that changes the ledger state.
- Protocol used to build consensus among participants around which transactions will be accepted, and in what order, by the ledger.

### Blockchain definition

Blockchain is a peer to peer distributed ledger forged by consensus combined with a system for smart contracts and other assistant technologies. 

Transactions established by trust, accountability, and transparency at their core. While streamlining business processes and legal constraints.

- Genesis block is the inital record for the distributed ledger.

Nodes or machines on a blockchain network groups up transactions and sen them through the network.

The process of blockchain syncing up has to do with a concept of consensus - and agreement among the network peers.

Each node will have the exact copy of the blockchain through the network.

- Smart contracts are computer programs that execute predefined actions when certain conditions within the system is met.

- Consensus refers to a system of ensuring that parties agree to a certain state of the system as the true state.

Blockchain is a specific form or subset of DLT, which constructs a chronological chain of blocks, hence the name "blockchain".

A block refers to a set of transactions that are bundled together and added to the chain at the same time. 

- Proof of work is when for example on the Bitcoin network, miners must solve a cryptographic challenge to propose the next block.

- Timestamping is when each new block referring to the previous block. Combined with cryptographic hashes, this timestamped chain of blocks provides an immutable record of all transactions in the network, from the genesis block.

A block consists of 4 pieces of metadata.

- Reference to the previous block
- Proof of work, also known as nonce
- Timestamp
- Merkle tree root for the transactions included in this block.

### Merkle Tree

Merkle Tree is a binary hash tree. Which is a sata structure that is used to store hashes of the individual data in large datasets in a way to make the verification of the dataset efficient.

It is a anti-tamper mechanism that ensures that the large dataset has not been changed.

According to Andreas M. Antonopoulos;

    "Merkle trees are used to summarize all the transactions in a block, producing an overall digital fingerprint of the entire set of transactions, providing a very efficient process to verify whether a transaction is included in a block.
    
### Transactions

Transactions on a blockchain are records of an event, cryptographically secured with a digital signature, that is verified, ordered, and bundled together into blocks.

On the Bitcoin blockchain, transactions involve the transfer of bitcoin, while in other blockchains, transactions may involve the transfer of any asset or a record of some service being rendered.

A smart contract within the blockchain may allow automatic execution of transaction upon metting predefined criteria.

### Cryptography

Key in security as well as in the immutability of the transactions recored on blockchains.

Cryptocurreny is the study of the technology used to allow secure communications between different parties and to ensure the authenticity and immutability of the data being communicated.

For blockchain technologies, cryptography is used to prove that a transaction was created by the right person. it is also used to link transactions into a block in a tamper-proof way as well as create the links between blocks, to form a blockchain.

### Difference between blockchain and Databases

A blockchain is a write-only data structure, where new entries get appended onto the end of the ledger. Every new block gets appended to the blockchain by linking to the previous block's hash.

There are no administratr permissions within a blockchain that allows editing or deleting of data.

Basically in a relational database, data can be easily modified or deleted, database admins who make changes to any part of the data and/or structure.

Blockchains were designed for decentralized applications, whereas relational databases, in general, were originally designed for centralized applications, where a single entity controls the data.

### Types of Blockchains

A blockchain can be permissionless or permissioned.

- Permissionless blockchain is also known as a public blockchain, because anyone can join the network.

- Permissioned blockchain, or a private blockchain, requires pre-verification of the participating parties within the network, and these parties are usually known to each other.

Choosing between permissionless vs permissioned blockchains should be driven by the particular application, use case driven.

Most enterprise use cases involve extensive vetting before parties agree to do business with each other. An example where a number of businesses exchange info is the supply chain management.

The supply chain management is an ideal use case for permissioned blockchains. You would not want non-vetted companies participating in the network. Each participant that is involved in the supply chain would require permissions to execute transactions on the blockchain. These transactions would allow other companies to understand where in the supply chain a particular item is.

When a network can commoditize trust, facilitating parties to transact without necessarily having to verify each other's identity, like the Bitcoin blockchain, a permissionless blockchain is more suitable.

### Peer to peer network architecture

Most applications utilize a central server or servers. For one user/client to send a message to another user/client in the network, the request has to be sent to the hub or a central server, which then directs it to the right computer.

Peer to peer (P2P) networks were first made popular by Napster and later Bittorrent, Tor, etc. This consists of computer systems which are directly connected to each other via the internet, without a central server.

Peers contribute to the computing power and storage that is required for the upkeep of the network. P2P networks are generally considered to be more secure than centralized networks, as they do not have a single point of attack, as in the case of a server-based network, where the security of the entire network can be compromised if the cnetral server is successfully attacked. 

As a result, large corps invest significant amounts of financial resources to fortify their central servers.

Permissionless P2P Systems do not require a set amount of peers to be online and are generally slower. 

Permissioned P2P networks have to guarantee uptime and require a high level of quality of service on the communication links.

Here are examples of consensus mechanisms or algorithms

- Proof of work
- Proof of Stake
- Proof of Elapsed Time
- Simplified Byzantine Fault tolerance

### Immutability of Data

Immutability of the data which sits on the blockchain is the most important reason to deploy blockchain-based solutions for a variet of centralized servers.

- Immutability or unchanging over time feature makes the blockchain useful for accounting, financial transactions, identity management, and asset ownership, management and transfer.

Once a transaction is written onto a blockchain, no one can change it, or it will be difficult to change. You would need to have access to all or most of the nodes.

Antony Lewis:

    "When people say that blockchains are immutable, they don't mean that the data can't be changed, they mean it is extremely hard to change without collusion, and if you try, it's extremely easy to detect the attempt."

It's difficult because each block is linked to the previous block by including the previous block's hash. This hash includes the Merkle root hash of all the transactions in the previous block.

If a single transaction were to change, not only would the Merkle root hash change, but so too would the hash contained in the changed block.

In addition, each subsequent block would need to be updated to reflect this change.

In the other hand, if someone did modify a transaction in a block without going through the necessary steps to update the subsequent blocks, it would be easy to recalculate the hashes used in the blocks and determine that something is amiss.

### Blockchain Applications

Applications built on top of a blockchain provide a gateway to accessing info that sits on that blockchain. Clients/users interact with the blockchain through the applications.

Here are a list of some current companies working with DLT. (3/26/2018)

2WAY.IO: 
Transforming public nodes into private nodes by adding a permissions layer. Private nodes can connect info silos and secure communication channels. They're user in control ( privacy by design and security by design) and require no trade-off between security and UX. These systems are both trusted third party and blockchain agnostic; they only require an intermediary or blockchain when both parties agree to add one to their interaction.

Atencoin:
First generation, identity based compliant digital currency. It is headed up by the National Atem Coin Foundation, an org that supports the identification of blockchain based tech and digital currencies.

BlockAuth:
Enables users to own and operate their own identity registrar that allows them to submit their information for verification.

Blockstack:
Provides a decentralized domain name system (DNS), decentralized public key distribution system, and registry for apps and user identies. Personal user APIs ship with the Blockstack app and handle everything from identity and authentication to data storage. Apps can request permissions from users and then gain read and write access to user resources.

Bitnation:
Goverance 2.0 platform that is powered by blockchain tech. Its goal is to provide the same services that gov provide, but in a decentralized and voluntary manner, unbound by geography. Bitnation has worked out an identification solution such as blockchain passport and a marriage certificate.

BlockVerify:
Provides blockchain based anti-counterfeit solutions. It uses blockchain tech to improve anti-counterfeit measures in different industries such as pharmaceuticals, luxury items, diamonds and electronics.

Cambridge Blockchain LLC:
Developing its digital identity software with several leading global financial institutions, with commercial deployments planned for late 2017.

For the entire list here is the reference link:
https://gomedici.com/22-companies-leveraging-blockchain-for-identity-management-and-authentication

### Smart Contracts

Smart Contracts are computer programs that execute predefined actions when certain conditions within the system are met.

They provide the language of transactions that allow the ledger state to be modified.

They can facilitate the exchange and transfer of anything of value.

### Consensus Algorithms

Consensus in the network refers to the process of achieving agreement among the network participants as to the correct state of data on the system.

Consensus leads to all nodes sharing the exact data.

- Ensures that the data on the ledger is the same for all the nodes in the network, preventing bad actors.

### Proof of work

This consensus algorithm involves solving a computational challenging puzzle in order to create new blocks in the Blockchain. Also called mining, and the nodes on the network that engage in mining are known as miners.

Kudelski Security report:
    " Proof of work is the outcome of a successful mining process and although the proof of work is hard to create it is easy to verify"

Ofir Beigel:
    "...guessing a combination to a lock is a proof to a challenge. It is very hard to produce this since you will need to guess many different combinations; but once produced, it is easy to validate. Just enter the combination and see if the lock opens".

51% attack refers to an attack on a blockchain by a group of miners controlling more than 50% of the network computing power.

### Proof of Stake

Proof of Stake algorithm is a generalization of the Proof of Work algorithm. In Pos, the nodes are known as the validators and rather than mining the blockchain, they validate the transaction to earn a transaction fee.

There is no mining to be done, as all coins exist from day one.

Simply put, nodes are randomly selected to validate blocks, and the probability of this random selection depends on the amount of stake held.

### Proof of Elapsed Time

Developed by Intel, the Proof of Elapsed Time consensus algorithm emulates the Bitcoin-style Proof of Work.

Hyperledger's Sawtooth implementation is an example of PoET at work.

Instead of competing to solve the cryptographic challenge and mine the next block, as in the Bitcoin blockchain, the PoET consensus algorithm is a hybrid of a random lottery and first-come-first-serve basis. In PoET, each validator is given a random wait time.

sawtooth.hyperledger.org
    "The validator with the shortest wait time for a particular transaction block is elected the leader."

This leader gets to create the next block on the chain.

### Simplified Byzantine Fault Tolerance

The Simplified Byzantine Fault Tolerant consensus algorithm implements and adpoted version of the Practical Byzantine Fault Tolerant algorithm and seeks to provide signicant improvements over Bitcoin's Proof of work consensus protocol.

The basic idea involves a single vlidator who bundles proposed transactions and forms a new block.

Unlike the Bitcoin blockchain, the validator is a known party, given the permissioned nature of the ledger.

Consensus is achieved as a result of a minimum number of other nodes in the network ratifing the new block. 

In order to be tolerant of a Byzantine fault, the number of nodes that must reach consensus is 2f+1 in a system containing 3f+1 nodes, where f is the number of faults in the system.

For example if we have 7 nodes in the system, then 5 of those nodes must agree if 2 of the nodes are acting in a faulty manner.

### Proof of Authority

Proof of Authority is a consensus algorithm which can be used for permissioned ledgers. It uses a set of authorities which are designated nodes that are allowed to create new blocks and secure ledger.
Ledgers using PoA require sign-off by a majority of authorizites in order for a block to be created.

### Hyperledger

Hyperledger is an open source effort created to advance cross-industry blockchain technologies. Hosted by The Linux Foundation, it is a global collaboration of members from various industries and orgs. Hyperledger boasts a host of enterprise-ready solutions. Hyperledger is about communities of software developers building blockchain frameworks and platforms.

