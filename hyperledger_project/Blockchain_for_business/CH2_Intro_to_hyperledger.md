### Hyperledger

Hyperledger is a group of open source projects focused around cross-industry distributed ledger technologies. 

Arnaud Le Hors:
    "these projects show how broadly applicable blockchain technology really is, This goes way beyond cryptocurrencies"

Hyperledger provides an alternative to the cryptocurreny-based blockchain model, and focuses on developing blockchain frameworks and modules to support global enterprise solutions. 

The focus of Hyperledger is to provide a transparent and collaborative approach to blockchain development.

### Components of Hyperledger Frameworks

Hyperledger business blockchain frameworks are used to build enterprise blockchains.

- Append-only distributed ledger
- Consensus algorithm for agreeing to changes in the ledger
- Privacy of transactions through permissioned access
- Smart contracts to process transactions requests

### List of Hyperledger Frameworks

Hyperledger Iroha v0.95

    ~/framework/Hyperledger_iroha

- Simple construction
- Modern
- Domain-driven C++ design
- Consensus algorithm YAC

Hyperledger Sawtooth v0.8
- Modular Platform
- Building, deploying, running distributed ledgers
- Can utilize various consensus algorithms based on size of network
- By default Proof of Elapsed Time consensus algorithm
- Highly scalable network of validator nodes
- Support for both permissioned and permissionless deployments

Hyperledger Fabric v1.0

    ~/framework/IBM_hyperledger_fabric

- Modular architecture
- Allows entities to conduct confidential transactions without passing info to central authority
- Different channels run within the network and the division of labor that characterizes different nodes within the network
- Permissioned deployments

Brian Behlendorf
    "If you have a larger blockchain network and you want to share data with only certain parties, you can create a private channel with just those participants. It is the most distinctive thing about Fabric now."

Hyperledger Indy
- Decentralized Identity
    "...decentralized identity specs and artifacts that are independent of any particular ledger and will enable interoperability across any DLT that supports them."
- Manage and control digital identities
- Rather than businesses store huge amounts of personal data; businesses store pointers to identity.
    "...identity is a toxic asset that could present a liability to organizatons."
    "Hyperledger Indy lets users authenticate identity based on the attributes they are willing to store and share themselves. This can reduce the amount of liability contained within a business because the data can be kept with the user and presented to you again in a way that you can trust and validate that what has been said really was said and is trusted by the other parties you do business with."

Hyperledger Burrow v0.16.1
- Permissionable smart contract machine that provides a modular blockchain client with a permissioned smart contract interpreter built-in part to the specs of the Ethereum Virtual Machine.
- Gateway provides interfaces for system integration and user interfaces.
- Smart contract application engine facilitates integration of complex business logic
- Consensus engine serves maintaining the networking stack between the nodes and ordering transactions.
- Application Blockchain Interface provides interface specification for the consensus engine and smart contract application engine to connect.

### Hyperledger Modules

Auxiliary softwares used for things like deploying and maintaining blockchains, examining the data on the ledger, as well as tools to design, prototype, and extend blockchain networks.

Hyperledger Cello
- Toolkit that allows Blockchain deployment to the cloud.
- Operators can create and manage blockchains through a dashboard, and users can obtain a blockchain instance immediately.
    "Cello aims to bring the on-demand 'as-a-service' deployment model to the blockchain ecosystem"

Hyperledger Explorer
- Tool for visualizing blockchain operations
- For permissioned ledgers, allowing anyone to explore the distributed ledger projects being created by Hyperledger's members from the inside, without compromising their privacy.
- Can view, invoke, deploy, or query:
- Blocks
- Transactions and associated data
- Network information (name, status, list of nodes)
- Smart contracts (chain codes and transaction families)
- Other relevant info stored in the ledger
- Web server, web UI, web sockets, database, security repos, and blockchain implementation

Hyperledger Composer

    ~/module/composer
    
- Suite of tools for building blockchain business networks
- Model business blockchain network
- Generate REST APIs for interacting with your blockchain network
- Generate a skeleton Angular application
- Javascript, easy to use set of components
- Faster creation of blockchain applications, eliminating the massive effort required to build blockchain applications from scratch
- Reduced risk with well-tested, efficient design that aligns understanding across business and technical analysts
- Greater flexibility as higher-level abstractions make it far simplier