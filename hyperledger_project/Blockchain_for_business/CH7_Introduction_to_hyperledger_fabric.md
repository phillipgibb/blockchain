### Hyperledger Fabric

### Demonstration of Integration

Tuna, Tuna, Tuna

### Defining our Actors

- Sarah is the fisherman who sustainably and legally catches tuna.
- Regulators verify that the tuna has been legally and sustainably caught.
- Miriam is a restaurant owner who will serve as the end user, in this situation.
- Carl is another restaurant owner fisherman Sarah can sell tuna to.

Using Hyperledger Fabric, we will be demonstrating how tuna fishing can be improved starting from the source, fisherman Sarah, and the process by which she sells her tuna to Miriam's restaurant.

### Featured Hyperledger Fabric Elements

- Channels are data partitioning mechanisms that allow transaction visibility for stakeholders only. Each channel is an independent chain of transaction blocks containing only transactions for that particular channel.
- The chaincode (Smart Contracts) encapsulates both the asset definitions and the business logic (or transactions) for modifying those assets. Transaction invocations result in changes to the ledger.
- The ledger contains the current world state of the network and a chain of transaction invocations. A shared, permissioned ledger is an append-only system of records and serves as a single source of truth.
- The network is the collection of data processing peers that form a blockchain network. The network is responsible for maintaining a consistently replicated ledger.
- The ordering service is a collection of nodes that orders transactions into a block.
- The world state reflects the current data about all the assets in the network. This data is stored in a database for efficient access. Current supported databases are LevelDB and CouchDB.
- The membership service provider(MSP) manages identity and permissioned access for clients and peers.

### The Catch

We will start with Sarah, our licensed tuna fisher, who makes a living selling her tuna to muiltiple restaurants. 

Sarah operates as a private business, in which her company frequently makes international deals.

Through a client application, Sarah is able to gain entry to a Hyperledger Fabric blockchain network comprised of other fishermen, as well as regulators and restaurant owners.

Sarah has the ability to add to and update info in the blockchain network's ledger as tuna pass through the supply chain, while regulators and restaurants have read access to the ledger.

After each catch, Sarah records infomation about each individual tuna, including: a unique ID number, the location and time of the catch, its weight, the vessel type, and who caught the fish.

For the sake of simplicity, we will stick with these six data attributes.

However, in an actual application, many more details would be recorded, from toxicology, to other physical characteristics.

These details are saved in the world state as a key/value pair based on the specifications of a chaincode contract, allowing Sarah's application to effectively create a transaction on the ledger. You can see the example below:

    var tuna ={

        id: '0001',
        holder: 'Sarah',
        
        location: {
            latitude: '41.40238',
            longitude: '2.170328'
            },

        when: '20170630123546',
        weight: '58lbs',
        vessel: '9548E'

    }

### The Incentives

Miriam is a restaurant owner looking to source low cost, yet high quality tuna that have been responsibly caught.

Whenever Miriam buys tuna, she is always uncertain whether she can trust that the tuna she is purchasing is legally and sustainably caught, given the prominence of illegal and unreported tuna fishing.

At the same time, as a legitimate and experienced fisherman, Sarah strives to make a living selling her tuna at a reasonable price. 

She would also like autonomy over who she sells to and at what price.

Luckily for both Sarah and Miriam, Hyperledger Fabric can help!

### The Sale

Normally, Sarah sells her tuna to resturateurs, such as Carl, for $80 per pound. 

However, Sarah agrees to give Miriam a special price of $50 per pound of tuna, rather than her usual rate. 

In a traditional public blockchain, once Sarah and Miriam have completed their transaction, the entire network is able to view the details of this agreement, especially the fact that Sarah gave Miriam a special price. 

As you can imagine, having other restaurateurs, such as Carl, aware of this deal is not economically advantageous for Sarah.

To remedy this, Sarah wants the specifics of her deal to not be available to everyone on the network, but still have every actor in the network be able to view the deatails of the fish she is selling.

Using Hyperledger Fabric's feature of channels, Sarah can privatetely agree to the terms with Miriam, such that only the two of them can see them, without anyone else knowing the specifics.

Additionally, other fishermen, who are not part of Sarah and Miriam's transaction, will not see this transaction on their ledger. This ensures that another fisherman cannot undercut the bid by having information about the prices that Sarah is charging different restaurateurs.

### The Regulators 

Regulators will also gain entry to this Hyperledger Fabric blockchain network to confirm, verify, and view details from the ledger. Their application will allow these actors to query the ledger and see the details of each of Sarah's catches to confirm that she is legally catching her fish.

Regulators only need to have query access, and do not need to add entries to the ledger.

With that being said, they may be able to adjust who can gain entry to the network and/or be able to remove fishermen from the network, if found to be partaking in illegal activities.

### Gaining Network Membership

Hyperledger Fabric is a permissioned network, meaning that only participants who have been approved can gain entry to the network. To handle network membership and identity, membership service providers (MSP) manage user IDs, and authenticate all the participants in the network.

A Hyperledger Fabric blockchain network can be governed by one or more MSPs.

This provides modularity of membership operations, and interoperability across different membership standards and architectures.

In our scenario, the regulator, the approved fishermen, and the approved restaurateurs should be the only ones allowed to join the network.

To achieve thism a membership service provider (MSP) is defined to accommodate membership for all members of this supply chain. 

In configuring this MSP, certificate and membership identities are created.

Policies are then defined to dictate the read/write policies of a channel, or the endorsement policies of a chaincode.

Our scenario has two seperate chaincodes, which are run on three separate channels.

The two chaincodes are: one for price agreement between fisherman and the restaurateur, and one for the transfer of tuna.

The three channels are: one for the price agreement between Sarah and Miriam; one for the price agreement between Sarah and Carl; and one for the transfer of tuna.

Each member of this network knows about each other and their identity.

The channels provide privacy and confidentiality of transactions.

In Hyperledger Fabric, MSPs also allow for dynamic membership to add or remove members to maintain integrity and operation of the supply chain.

For example, if Sarah was found to be catching her fish illegally, she can have her membership revoked, without compromising the rest of the network.

This feature is critical, especially for enterprise applications, where business relationships change over time.

### Summary of Demonstrated Scenario

Below is a summary of tuna catch scenario presented in this section:

- Sarah catches a tuna and uses the supply chain application's user interface to record all the details about the catch to the ledger. Before it reaches the ledger, the transaction is passed to the endorsing peers on the network, where it is then endorsed. The endorsed transaction is sent to the ordering service, to be ordered into a block. This block is then sent to the committing peers in the network, where it is committed after being validated.

- As the tuna is passed along the supply chain, regulators may use their own application to query the ledger for details about specific catches (excluding price, since they do not have access to the price-related chaincode).

- Sarah may enter into an agreement with a restaurateur Carl, and agree on a price of $80 per pound. They use the blue channel for the chaincode contract stipulating $80/lb. The blue channel's ledger is updated with a block containing this transaction.

- In a seperate business agreement, Sarah and Miriam agree on a special price of $50 per pound. They use the red channel's chaincode contract stipulating $50/lb. The red channel's ledger is updated with a block containing this transaction.

### Roles within a Hyperledger Fabric Network

There are three different types of roles within a Hyperledger Fabric network:

- Clients: Clients are applications that act on behalf a person to propose transactions on the network.

- Peers: Peers maintain the state of the network and a copy of the ledger. There are two different types of peers: endorsing and committing peers. However, there is an overlap between endorsing and committing peers, in that endorsing peers are a special kind of committing peers. All peers commit blocks to the distributed ledger.
    - Endorsers simulate and endorse transactions
    - Committers verify endorsements and validate transaction result, prior to committing transactions to the blockchain.

- Ordering Service: The ordering service accepts endorsed transactions, orders them into a block, and delivers the block to the committing peers.

### How to Reach Consensus

In a distributed ledger system, consensus is the process of reaching agreement on the next set of transactions to be added to the ledger. In Hyperledger Fabric, consensus is made up of three distinct steps:

- Transaction endorsement
- Ordering
- Validation and commitment

These three steps ensure the policies of a network are upheld. We will explore how these steps are implemented by exploring the transaction flow.

### Transaction Flow

Within a Hyperledger Fabric network, transactions start out with a client application sending transaction proposals, or in other words, proposing a transaction to endorsing peers.

Client applications are commonly referred to as application or clients, and allow people to communicate with the blockchain network.

Application developers can leverage the Hyperledger Fabric network through the application SDK.

Each endorsing peer simulates the proposed transaction, without updating the ledger.

The endorsing peers will capture the set of Read and Written data, called RW Sets.

These RW sets capture what was read from the current world state while simulating the transaction, as well as what would have been written to the world state had the transaction been executed.

These RW sets are then signed by the endorsing peer, and returned to the client application to be used in future steps of the transaction flow.

Endorsing peers must hold smart contracts in order to simulate the transaction proposals.

### Transaction Endorsement

A transaction endorsement is a signed response to the results of the simulated tranasction. 

The method of transaction endorsements depends on the endorsement policy which is specified when the chaincode is deployed.

An example of an endorsement policy would be "the majority of the endorseing peers must endorse the transaction".

Since an endosement policy is specified for a specific chaincode, different channels can have different endorsement policies.

The application then submits the endorsed transaction and the RW sets to the ordering service.

Ordering happens across the network, in parallel with endorsed transactions and RW sets submitted by other applications.

The ordering service takes the endorsed transactions and RW sets, orders this information into a block, and delivers the block to all committing peers.

The ordering service, which is made up of a cluster of orders, does not process transactions, smart contracts, or maintains the shared ledger.

The ordering service accepts the endorsed transactions and specifies the order in which those transactions will be committed to the ledger.

The Fabric v1.0 architecture has been designed such that the specific implementation of 'ordering' (Solo, Kafka, BFT) becomes a pluggable component.

The default ordering service for Hyperledger Fabric is Kafka.

Therefore, the ordering service is a modular component of Hyperledger Fabric.

### Ordering

    Transactions within a timeframe are sorted into a block and are committed in sequential order.

In a blockchain network, transactions have to be written to the shared ledger in a consistent order.

The order of transactions has to be established to ensure that the updates to the world state are valid when they are committed to the network.

Unlike the Bitcoin blockchain, where ordering occurs through the solving of a cryptographic puzzle, or mining, Hyperledger Fabric allows the organizations running the network to choose the ordering mechanism that best suits that network.

This modularity and flexibility makes Hyperledger Fabric incredibly advantageous for enterprise applications.

Hyperledger Fabric provides three ordering mechanisms: SOLO, Kafka, and Simplified Byzantine Fault Tolerance (SBFT), the latter of which has not yet been implemented in Fabric v1.0.

- SOLO is the Hyperledger Fabric ordering mechanism most typically used by developers experimenting with Hyperledger Fabric networks. SOLO involves a single ordering node.

- Kafka is the Hyperledger Fabric ordering mechanism that is recommended for production use. This ordering mechanism utilizes Apache Kafka, an open source stream processing platform that provides a unified, high-throughput, low-latency platform for handling real-time data feeds. In this case, the data consists of endorsed transactions and RW sets. The Kafka mechanism provides a crash fault-tolerant solution to ordering.

- SBFT stands for Simplified Byzantine Fault Tolerance. This ordering mechanism is both crash fault-tolerant and byzantine fault-tolerant, meaning that it can reach agreement even in the presence of malicious or faulty nodes. The Hyperledger Fabric community has not yet implemented this mechanism, but it is on their roadmap.

These three ordering mechanisms provide alternate methodologies for agreeing on the order of transactions.

### Tranasction Flow

The committing peer validates the transaction by checking to make sure that the RW sets still match the current world state.

Specically, that the Read data that existed when the endorsers simulated the transaction is identical to the current world state.

When the committing peer validates the transaction, the transaction is written to the the ledger, and the world state is updated with the Write data from the RW Set.

If the transction fails, that is, if the committing peer finds that the RW set does not match the current world state, the transaction ordered into a block will still be included in that block, but it will be marked as invalid, and the world state will not be updated.

Committing peers are responsible for adding blocks of transactions to the shared ledger and updating the world state.

They may hold smart contracts, but it is not a requirement.

Lastly, the committing peers asynchronosuly notify the client application of the success or failure of the transaction.

Applications will be notified by each committing peer.

### Identity Verification

In addition to the multitude of endorsement, validity and versioning checks that take place, there are also ongoing indentity verifications happening during each step of the transaction flow.

Access control lists are implemented on the hierarchical layers of the network(from the ordering service down to channels), and payloads are repeatedly signed, verified, and authenticated as a transaction proposal passes through the different architectural components.

### Transaction Flow Summary

It is important to note thtat the state of the network is maintained by peers, and not by the ordering service or the client.

Normally, you will design your system such that different machines in the network play different roles.

That is, machines that are part of the ordering service should not be set up to also endorse or commit transactions, and vice versa.

However, there is an overlap between endorsing and committing peers on the system.

Endorsing peers must have access to and hold smart contracts, in addition to fulfilling the role of a committing peer.

Endorsing peers do commit blocks, but committing peers do not endorse transactions.

Endorsing peers verify the client signature, and execute a chaincode function to simulate the tranasction.

The output is the chaincode results, a set of key/value versions that were read in the chaincode (read set), and the set of keys/values that were written by the chaincode.

The proposal response gets sent back to the client, along with an endorsement signature.

These proposal responses are sent to the orderer to be ordered.

The orderer then orders the transactions into a block, which it forwards to the endorsing and committing peers.

The RW sets are used to verify that the transactions are still valid before the content of the ledger and world state is updated.

Finally, the peers asynchronously notify the client application of the success or failure of the transaction.

### Channels

Channels allow organizations to utilize the same network, while maintaining separation between blockchains.

Only the members of the channel on which the transaction was performed can see the specifics of the transaction.

In other words, channels partition the network in order to allow transaction visiblity for stakeholders only.

This mechanism works by delegating transactions to different ledgers.

Only the members of the chaneel are involved in consensus, while other members of the network do not see the transactions on the channels.

The diagram above shows three distinct channels -- blue, orange, and grey. 

Each channel has its own application, ledger and peers.

Peers can belong to multiple networks or channels.

Peers that do participate in multiple channels simulate and commit transactions to different ledgers. 

The ordering service is the same across any network or channel.

A few things to remember:

- The network setup allows for the creation of channels.
- The same chaincode logic can be applied to multiple channels.
- A given user can participate in multiple channels.

### State Datebase

The current state data represents the latest values for all assets in the ledger.

Since the current state represents all the committed transactions on the channel, it is sometimes referred to as world state.

Chaincode invocations execute transactions against the current state data.

To make these chaincode interations extremely efficient, the latest key/value pairs for each asset are stored in a state database.

The state database is simply an indexed view into the chain's committed transactions.

It can therefore be regenerated from the chain at any time.

The state database will automatically get recovered (or generated, if needed) upon peer startup, before new transactions are accepted.

The default state datebase, LevelDB, can be replaced with CouchDB.

- LevelDB is the default key/value state database for Hyperledger Fabic, and simply stores key/value pairs.
- CouchDB is an alternative to LevelDB. Unlike LevelDB, CouchDB stores JSON objects. COuchDB is unique in that it supports keyed, composite, key range, and full data-rich queries.

Hyperledger Fabric's LevelDB and CouchDB are very similar in their structure and function.

Both LevelDB and CouchDB support core chaincode operations, such as getting and setting key assets, and querying based on these keys.

With both, keys can be queried by range, and composite keys can be modeled to enable equivalence queries against multiple parameters.

But, as a JSON doc store, CouchDB additionally enables rich query against the chaincode data, when chaincode values (e.g. assets) are modeled as JSON data.

### Smart Contracts

As a reminder, smart contracts are computer programs that contain logic to execute transactions and modify the state of the assets stored within the ledger. 

Hyperledger Fabric smart contracts are called chaincode and are written in Go.

The chaincode serves as the business logic for a Hyperledger Fabric network, in that the chaincode directs how you manipulate assets within the network.

### Membership Service Provider (MSP)

The membership service provider, or MSP, is a component that defines the rules in which identites are validated, authenicated and allowed access to a network.

The MSP manages user IDs and authenticates clients who want to join the network.

This includes providing credentials for these clients to propose transactions.

The MSP makes use of a Certificate Authority, which is a pluggable interface that verifies and revokes user certicates upon confirmed identity.

The default interface used for the MSP is the Fabric-CA API, however, organizations can implement an External Certificate Authority of their choice.

This is another feature of Hyperledger Fabirc that is modular.

Hyperledger Fabric supports many credential architectures, which allows for many types of External Certificate Authority interfaces to be used.

As a result, a single Hyperledger Fabric network can be controlled by Multiple MSPs, where each organization brings their favorite.

### What does the MSP Do?

To start, users are authenticated using a certificate authority.

The certificate authority identifies the application, peer, endorser, and orerer identites, and verifies these credentials.

A signature is generated through the use of a Signing Algorithm and a Signature Verification Algorithm.

Specifically, generating a signature starts with a sSigning Algorithm, which utilizes the credentials of the entities associated with their respective identites, and outputs an endorsement.

A signature is generated, which is a byte array that is bound to a specific identity.

Next, the Signature Verification Algorithm takes the identity, endorsement, and signature as inputs, and outputs 'accept' if the signauture byte array corresponds with a valid signature for the inputted endorsement, or outputs 'reject' if not.

If the output is 'accept', the user can see the transaction in the network and perform transactions with other actors in the network.

If the output is 'reject', the user has not been properly authenticated, and is not able to submit transactions to the network, or view any previous transactions.

### Fabric-Certificate Authority

In general, Certificate Authorities manages enrollment certificates for a permissioned blockchain.

Fabric-CA is the default certificate authority for Hyperledger Fabric, and handles the registration of user identites.

The Fabric-CA certificate authority is in charge of issuing and revoking Enrollment Certificates (E-Certs).

The current implementation of Fabric-CA only issues E-Certs, which supply long term identity certificates.

E-Certs, which are issued by the Enrollment Certificate Authority (E-CA), assign peers their identity and give them permission to join the network and submit transactions.

### Technical Prerequisites

- Go
- Nodejs
- cURL
- npm
- Docker
- Docker Compose

### Installing Hyperledger Fabric Docker Images and Binaries

Next we will download the latest released Docker images for Hyperledger Fabric, and tag them with the latest tag.

Execute the command from within the directory into which you will extract the platform-specific binaries:

    curl -sSL https://goo.gl/Q3YRTi | bash

NOTE: Check https://hyperledger-fabric.readthedocs.io/en/latest/samples/html#binaries for the latest URL (the blue portion in the above curl command) to pull in binaries.

This command downloads binaries for cryptogen, configtxgen, configxlator, peer AND downloads the Hyperledger Fabric Docker images.

These assets are placed in a bin subdirectory of the current working directory.

To confirm and see the list of Docker images you've just downloaded, run:

    docker images

The expected response is:

Notes the tags for each of the repos above boxed in red.

If the Docker images are not already tagged with the latest tag, perform the following command for each of the Docker images:

    docker tag hyperledger/fabric-tools:x86_64-1.0.2 hyperledger/fabric-tools:latest

Swap out the blue portion with the tags you see in your list of repos.

Also, sawp out the red portion with the name of the Docker image you are swithcing the tag for (e.g.: fabric-tools, fabric-ccenv, fabric-orderer, etc.). Repeat this step for all Docker images you see in the list.

### Sample-network for Hyperledger Fabric github no longer exists


### Alternative Installation Instruction: 
These instructions are for installing docker, docker-compose, hyperledger fabric, hyperledger composer for Windows' 10 Linux subsystem. Takes a while but worth it when working on a PC. I like Macs but they are really not that customizable as well as having less third party developer applications for it sometimes. But sometimes there are certain commands that only work on linux so the Windows' Linux Subsystem is a great balance between the two. 

The alternative would to be either buy a mac, virtual machine or server with linux OS. 


### Instally Docker and locally hosting Hyperledger Fabric; also installs composer module

### Remember to install docker

When the Windows Subsystem for Linux (WSL) – or, as most people even at Microsoft often refer to it – Bash on Ubuntu on Windows – was announced on Microsoft’s Build conference 2016, a world of new tools opened up to us Windows devs. Personally, I love being able to choose between PowerShell, Bash or plain old cmd when I want to script something. And it’s always bugged me that I couldn’t get Docker working from Bash on Windows – until now.

The original title of this post was “Running Docker from Bash on Windows”, but that would have been a slight overstatement. Docker requires access to quite a of lot system calls which aren’t necessarily all implemented on Windows, so getting the engine running under the WSL is probably not so easy. Instead, we’ll run the Docker Engine on Windows, and connect to it from Bash. This also has the advantage that you can start a container from PowerShell and interact with it from Bash, or the other way around – in other words, your computer will still feel just like one machine.

Here’s how to do it:

1. Install Docker on Windows
To install the Docker engine on Windows, just go to docker.com and download the appropriate distribution. Also, make sure hardware virtualization is enabled and Hyper-V is installed, lest the engine won’t start.

Shortcut: Install Windows 10 Creators Update
With Windows 10 Creators Update*, accomplishing all of this has become a lot simpler, since it allows you to run Windows executables from Bash. Just add these two lines to your .bashrc (and reload your environment) and you’re done!


export PATH="$HOME/bin:$HOME/.local/bin:$PATH"
export PATH="$PATH:/mnt/c/Program\ Files/Docker/Docker/resources/bin"
alias docker=docker.exe
alias docker-compose=docker-compose.exe
1
2
3
4
export PATH="$HOME/bin:$HOME/.local/bin:$PATH"
export PATH="$PATH:/mnt/c/Program\ Files/Docker/Docker/resources/bin"
alias docker=docker.exe
alias docker-compose=docker-compose.exe
You can now run docker --version from Bash, and you don’t even have to read the rest of this blog post :)

*) Windows 10 Creators Update is available to Insiders since April 11, 2017, and will be released to the public on April 25, 2017.

Making it work on Windows 10 Anniversary Edition
To install Docker on the WSL, you’ll need to jump through a few more hoops. There’s a description for Ubuntu in general here, which works for the WSL as well, with the exceptions of some of the optional steps. Here’s what I did:


# Install packages to allow apt to use a repository over HTTPS
$ sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
# Add Docker's official GPG key
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# Set up the repository
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
# Update source lists
sudo apt-get update
# Install Docker
sudo apt-get install docker-ce
1
2
3
4
5
6
7
8
9
10
# Install packages to allow apt to use a repository over HTTPS
$ sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
# Add Docker's official GPG key
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# Set up the repository
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
# Update source lists
sudo apt-get update
# Install Docker
sudo apt-get install docker-ce
Of course, there’s also the option of downloading and extracting the binaries we’ll need, and put them somewhere in your PATH. There are instructions here for how to get the latest version.

Where did that get us?
We now actually have the Docker engine installed on both Windows and the WSL, but it isn’t started on either. The Windows installer helpfully created a Docker shortcut on the desktop and/or in the Start menu – use that to start the Docker engine. Then, you can try running e.g. docker images from PowerShell and from Bash:

PowerShell:


PS C:\> docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
1
2
PS C:\> docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
We haven’t created any images yet, so that’s fine.

Bash:


$ docker images
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
1
2
$ docker images
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
Clearly unsatisfying. But with one or two extra steps, we’ll get it all working.

2. Connect Docker on WSL to Docker on Windows
Running docker against an engine on a different machine is actually quite easy, as Docker can expose a TCP endpoint which the CLI can attach to.

Note: In an update to Docker that was pushed shortly after first publishing this post, this TCP endpoint is turned off by default (thanks Mark for the heads-up!); to activate it, right-click the Docker icon in your taskbar and choose Settings, and tick the box next to “Expose daemon on tcp://localhost:2375 without TLS” (make sure you understand the risks).

With that done, all we need to do is instruct the CLI under Bash to connect to the engine running under Windows instead of to the non-existing engine running under Bash, like this:


$ docker -H tcp://0.0.0.0:2375 images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
1
2
$ docker -H tcp://0.0.0.0:2375 images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
Much better!
There are two ways to make this permanent – either add an alias for the above command, or better yet (thanks Dave!), export an environment variable which instructs Docker where to find the host engine:


$ echo "export DOCKER_HOST='tcp://0.0.0.0:2375'" >> ~/.bashrc
$ source ~/.bashrc
1
2
$ echo "export DOCKER_HOST='tcp://0.0.0.0:2375'" >> ~/.bashrc
$ source ~/.bashrc
Now, running docker commands from Bash works just like they’re supposed to.

$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
1
2
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
Mission accomplished!


### Remember to add 8.8.8.8 or will have connectivity errors.


* Something to note is that virtualbox is unable to be installed onto Windows10 bash because of compadability issueses. 

(Recommended) Setup of Hyperledger Fabric Network: 
Prerequisite (Setting up local development environment)
(Optional but highly recommended) Install virtual environment: https://realpython.com/blog/python/python-virtual-environments-a-primer/
Patiently Install (Python 2.7.x+, Git 2.9.x+, Docker Engine -v17.03+, Docker Compose -v1.8+, Node.js -v6.x (Node v7+ is not supported), and Xcode if on Mac). Set recommended FABRIC_VERSION=hlfv1(v1.0) as an environment variable. Here are some instructions to get started, please follow the links for more installation steps. 

# install git
https://www.atlassian.com/git/tutorials/install-git

# install node version 6
https://medium.com/@katopz/how-to-install-specific-nodejs-version-c6e1cec8aa11

    ~$ brew update
    ~$ brew install node@6
    ~$ brew unlink node
    ~$ brew link node@6
    ~$ node --version
    ~$ npm --version 

### install docker
 Installed on Windows 10 linux(bash) subsystem. Note that virtualbox wont be able to be installed onto the bash subsystem.

### Run these commands in order

Add to the commands   
    sudo
if they say you dont have permission

    --unsafe-perm
after install if the installation hangs because they say more things about permissions.

Essential CLI tools
    npm install -g composer-cli

Utility for running a REST Server on your machine to expose your business networks as RESTful APIs
    npm install -g composer-rest-server

Useful utility for generating application assets
    npm install -g generator-hyperledger-composer

Yeoman is a tool for generating applications, which utilises generator-hyperledger-composer
    npm install -g yo

Install Playground
    npm install -g composer-playground

Create fabric tool folder in dir of your choice
    mkdir ~/fabric-tools && cd ~/fabric-tools

Download zip file of tools
    curl -O https://raw.githubusercontent.com/hyperledger/composer-tools/master/packages/fabric-dev-servers/fabric-dev-servers.zip

unzip
    unzip fabric-dev-servers.zip

cd into folder
    cd ~/fabric-tools

Start download
    ./downloadFabric.sh

run
    composer-playground

localhost:8080

### IBM's Hyperledger Fabric
Learning their blockchain platform

    docker run --name composer-playground --publish 8080:8080 hyperledger/composer-playground


### Chaincode

In Hyperledger Fabric, chaincode is the 'smart contract' that runs on the peers and creates transactions.

More broadly, it enables users to create transactions in the Hyperledger Fabric network's shared ledger and update the world state of the assets.

Chaincode is programmable code, written in Go, and instantiated on a channel.

Developers use chaincode to develop business contracts, asset definitions, and collectively-managed decentralized applications.

The chaincode manages the ledger state through transactions invoked by applications.

Assets are created and updated by a specific chaincode, and cannot be accessed by another chaincode.

Applications interact with the blockchain ledger through the chaincode.

Therefore, the chaincode need to be installed on every peer that will endorse a transaction and instantiated on the channel.

There are two ways to develop smart contracts with Hyperledger Fabric:

- Code individual contracts into standalone instances of chaincode.
- (More efficient way) Use chaincode to create decentralized applications that manage the lifestyle of one or multiple types of business contracts, and let the end users instantiate instances of contracts within these applications.

### Chaincode Key APIs

An important interface that you can use when writing your chaincode is ChaincodeStub and ChaincodeStubInterface.

The Chaincodestub provides functions that allow you to interact with the underlying ledger to query, update, and delete assets.

The key APIs for chaincode include:

- func (stub *ChaincodeStub) GetState(key string) ([]byte, error)
    - Returns the value of the specfied key from the ledger. Note that GetState doesnt read data from the Write set, which has not been committed to the ledger. In other words, GetState doesnt consider data modified by PutState that has not been committed. If the key does not exist in the state database, (nil, nil) is returned.

- func (stub *ChaincodeStub) PutState(key string, value []byte) error
    - Puts the specified key and value into the tranasction's Write set as a data-write proposal. PutState doesn't affect the ledger until the transaction is validated and successfully committed.

- func (stub *ChaincodeStub) DelState(key string) error
    - Records the specific key to be deleted in the Write set of the transaction proposal. The key and its value will be deleted from the ledger when the transaction is validated and successfully committed.

### Overview of a Chaincode Program

When creating a chaincode, there are two methods that you will need to implement:

- Init
    - Called when a chaincode receives an instantiate or upgrade transaction. This is where you will initilize any application state.

- Invoke
    - Called when the invoke transaction is received to process any transaction proposals.

As a developer, you must create both an Init and an Invoke method within your chaincode. 

The chaincode must be installed using the

    peer chaincode install

command, and instantiated using the

    peer chaincode instantiate

command before the chaincode can be invoked. Then, transactions can be created using the

    peer chaincode invoke

or

    peer chaincode query

commands.

### Sample Chaincode Decomposed - Dependencies

Lets now walk through a sample chaincode written in Go, piece by piece:

    package main

    import (

        "fmt"

        "github.com/hyperledger/fabric/core/chaincode/shim"

        "github.com/hyperledger/fabric/protos/peer"

    )


The import statement lists a few dependencies that you will need for your chaincode to build successfully.

- fmt - contains Println for debuggin/logging
- github.com/hyperledger/fabric/core/chaincode/shim - contains the definition for the chaincode interface and the chaincode stub, which you will need to interact with the ledger, as we described in the Chaincode Key APIs section
- github.com/hyperledger/fabric/protos/peer - contains the peer protobuf package

### Sample Chaincode Decomposed - Struct

    type SampleChaincode struct {
        
    }

This is a statement that begins the definition of an object/class in Go. SampleChaincode implements a simple chaincode to manage an asset.

### Sample Chaincode Decomposed - Init Method

Next, we'll implement the Init method.

Init is called during the chaincode instantiation data required by the application.

In our sample, we will create the inital key/value pair for an asset, as specified on the command line:

func (t *SampleChaincode) Init(stub shim.ChainCodeStubInterface) peer.Response {

    // Get the args from the transaction proposal

    args := stub.GetStringArgs()

    if len(args) != 2 {

        return shim.Error("Incorrect arguments. Expecting a key and a value")

    }

    // We store the key and the value on the ledger

    err := stub.PutState(args[0], []byte(args[1]))

    if err != nil {

        return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))

    }

    return shim.Success(nil)

}