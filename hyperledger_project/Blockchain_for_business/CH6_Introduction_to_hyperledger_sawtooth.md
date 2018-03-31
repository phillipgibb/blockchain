### Defining our actors

- Sarah is the fisherman who sustainably and legally catches tuna.
- Tuna processing plant processes and bags the tuna after they have been caught.
- Regulators verify that the tuna has been legally and sustainably caught.
- Miriam is a restaurant owner who will serve as the recipient in this situation.

We will be describing how tuna fishing can be improved starting from the source, fisherman Sarah, and the process by which her tuna ends up at Miriam's restaurant.

### Featured Hyperledger Sawtooth Elements

- Transaction validators validate transactions.
- Transaction families are smart contracts in Hyperledger Sawtooth. They define the operations that can be applied to transactions. Transaction families consist of both transaction processor (the server-side logic) and clients (for use from Web or mobile applications)
- Transaction processor is the server-side business logic of transaction families that act upon network assets.
- Transaction batches are clusters of transactions that are either all committed to state or are all not committed to state.
- The network layer is responsible for communicating between validators in a Hyperledger Sawtooth network, including performing performing initial connectivity, peer discovery, and message handling.
- Global state contains the current state of the ledger and a chain of transaction invocations. The state for all transaction families is represented on each validator. The process of block validation on each validator ensures that the same transactions result in the same transitions, and that the resulting data is the same for all participants in the network. The state is split into namespaces, which allow flexibility for transaction family authors to define, share, and reuse global state data between transaction processors.

### Why Sawtooth

Miriam is a restaurant owner looking to source high quality tuna, that have been caught responsibly. Whenever Miriam buys tuna, she cannot be sure whether she can trust that the tuna she is purchasing is legally and sustainably caught, given the promience of illegal and unreported tuna fishing. 

Hyperledger Sawtooth is ideal for this scenario because of its ability to track an asset's (in this case tuna) provenance and journey. The ability to batch transactions together allows for all tuna within a catch to be entered as a whole. The distributed state agreement, novel consensus algorithm, and decoupled business logic from the consensus layer allow Miriam to be confident that she is buying tuna that has been legally caught.

### The Supply Chain

Hyperledger Sawtooth is scalable and can handle high volumes of data, which makes it a great option for production supply chain scenarios.

Starting with Sarah, our licensed tuna fisherwoman, who seels her tuna to multiple restaurants.

Sarah operates as a private business, in which her company frequently makes international deals.

Through an application, Sarah is able to gain entry to a Hyperledger Sawtooth blockchain network comprised of other fishermen, as well as processing plants, regulators, and resturant owners.

Sarah, as well as the processing plants, have the ability to add and update information to this ledger as tuna passes through the supply chain, while regulators and resturants have read access to ledger records.

### The Tuna Asset

With each catch, Sarah records some infor about each individual tuna, including: 
- a unique ID number
- location
- time of catch
- weight
- who caught the fish

For the sake of simplicity, we will stick with these data attributes.

However, in an actual application, many more details would be recorded, from toxicology, to other physical characteristics, such as the temperature of the tuna.

These details are saved in the global state as a key/value pair based on the specifications of a transaction family, while the transaction processor allows Sarah's application to effectively create a transaction on the ledger. 

Please see the example below:

    $ var tuna =
    { 
        id: '0002', 
        holder: 'Sarah',

        location:
            {
            latitude: '15.623036831528264',
            longitude: '-158.466796875'
            },

        when : '20170635123546',
        weight: '58lbs'
    }

### Recording a Catch

After Sarah catches her tuna and records data for each tuna, she is able to treat a haul of tuna as a single batch of transactions.

Using batches, Sarah is able to record many tuna together, while still being able to specify data for each tuna.

If one of the tuna transactions is invalid, the entire shipment is invalidated, that is, no tuna within the batch of tuna is validated.

### Reaching Consensus

After a batch of many transactions is submitted to the network, the network's consensus algorithm is run to choose a node to publish the transaction block.

By default, the proof of Elapsed Time consensus algorithm is used, and the transaction validator with the shortest wait time publishes the transaction block.

The transaction block is then broadcast to the publishing nodes.

Each node within the network receives the transaction block, and validators verify whether the transaction is valid or not.

If the transaction is validated, the global state is updated.

Hyperledger Sawtooth is unique because of its distributed state agreement, whereby every node in the system has the same understanding of information, and, as the supply chain matures, the data stored remains consistent across the network.

With the global state adjusted, the processing plant, the regulator, and Miriam are able to see the details of the catch and who the current holder is, thus, keeping the supply chain transparent and verifiable.

### Other Actors in the Supply Chain

Before Sarah's tuna can reach Miriam's restaurant, they need to go through a tuna processing plant.

Also, regulators will need to verify and view details from the ledger.

Therefore, both parties will gain entry to this Sawtooth blockchain.

The regulators will need to query the ledger to confirm that Sarah is legally catching her fish.

At the same time, tuna processing plants will need to record their processing and shipping of tuna on the ledger.

### Summary of Transaction Flow

    1. Sarah catches a tuna and uses the application's user interface to record all the details about the catch to the ledger. An entire haul of tuna can be treated as a transaction batch, with each individual tuna registered as a transaction within the batch.

    2. After the details of the catch are saved to the ledger and the tuna is passed along the supply chain, the processing plant may use their own application to query the ledger for details about specific catches, such as weight. The processing plant will add details reflecting the processing date and time, as well as other relevant information.

    3. As the tuna is passed along the supply chain, the regulator may use their respective application to query the ledger for details about specific catches, such as the owner and location of the catch, to verify legitimacy.

    4. Lastly, when the haul of tuna arrives at Miriam's restaurant, Miriam is able to use her application to query the ledger, to check and make sure Sarah was truthful in her account of where each fish was sourced. This, in turn, guarantees Miriam's restaurant is serving their customers the sustainably caught fish they advertise.

### Requirements Supported by Hyperledger Sawtooth

Hyperledger Sawtooth is a blockchain framework with potential in IoT, manufacturing, finance, and enterprise. 

Hyperledger Sawtooth supports diverse requirements, including both permissioned and permissionless deployments, and a pluggable consensus algorithm.

This framework also provides a revoltionary consensus algorithm, Proof of Elapsed Time(PoET), that allows for versatility and scalability suited for a variety of solutions.

Hyperledger Sawtooth supports many different infrastructural requirements, such as:

- Permissioned and permissionless infrastructure
- Modular blockchain architecture
- Scalability, which is good for larger blockchain networks due to higher throughput
- Many languages for transaction logic

### Transaction Batching

Hyperledger Sawtooth transaction batches are clusters of transactions that are either all committed to state together, or none of the transactions are committed at all.

As a result, transaction batches are often described as an atomic unit of change, since a group of transactions are treated as one, and are committed to the state as one.

Every single transaction in Hyperledger Sawtooth is submitted within a batch.

Batches can contain as little as a single transaction.

When a transaction is created by a client, the batch is submitted to the validator (which we will cover more indepth in the next section).

Transactions are organized into a batch in the order they are intended to be committed.

The validator then, in turn, applies each transaction within the batch, leading to a change in the global state.

The batch is committed to the state.

If one transaction within the batch is invalid, then none of the transactions within that batch are committed.

In summary, transaction batching allows a group of transactions to be applied in a specific order, and if any are invalid, then none of the transactions are applied.

This is a powerful tool that can be utilized by many enterprise solutions, as it provides greater efficiency and control for end users.

### Validators

In any blockchain network, modifying the global state requires creating and applying a transaction.

In Hyperledger Sawtooth, validators apply blocks that cause a change in the state.

More specifically, validators validate transaction blocks, and ensure that transactions result in state changes that are consistent across all participants in the network.

To start, a user creates a transaction batch and submits it to a validator via a client and REST API. 

The validator then checks the transaction batch and applies it if it is considered valid, resulting in a change to the state.

In terms of our demonstrated scenario, Sarah, the fisherman, creates a transaction batch to record info about a group of tuna catches.

The validator would then apply the transaction, and the state would be updated if all appropriate attributes are present: a unique ID number, location and time of the catch, weight, and who caught the fish.

If any of these elements are missing, the transactions within the batch would not be applied, and the state would not be updated.

### Journal

In Hyperledger Sawtooth, the journal maintains and extends the blockchain for the validator.

It is responsible for validating candidate blocks, evaluating valid blocks to determine if they are correct chain head, and generating new blocks to extend the chain.

Transaction batches arrive at the journal, where they are evaluated, validated and added to the blockchain. 

Additionally, the journal resolves forks, which occur due to disagreements over who commits a block.

Once blocks are completed, they are devliered to the ChainController for validation and fork resolution.

To see how the elements of the journal interact with one another, take a look at the diagram on the next page.

Another key feature of the journal is its flexibility in allowing pluggable consensus algorithms.

### Consensus Interface

Consensus in Hyperledger Sawtooth is modular, meaning that the consensus algorithm can be easily modified.

Hyperledger Sawtooth provides an abstract interface that supports both PBFT and Nakamoto-style algorithms. 

To implement a new consensus algorithm in Hyperledger Sawtooth, you must implement the distinct interface for: block publisher, block verifier, and fork resolution.

- Block publisher: Creates new candidate blocks to extend the chain
- Block verifier: Verifies that candidate blocks are plublished in accordace with consensus rules
- Fork resolver: Chooses which fork to use as the chain head for consensus algorithms that result in a fork.

These interfaces are used by the journal component.

The journal verifies that all the dependencies for the transaction batches are satisfied.

When verified, completed batches are checked for validity and fork resolution, and then, they are published within a block.

### Transaction Families

As with any blockchain framework, transaction updates need to be approved and shared between many untrusted parties.

As such, many blockchain frameworks have a mechanism for supporting distributed ledgers, as well as a method for changing the state of the shared ledger.

In Hyperledger Sawtooth, the data model that captures the state and the transaction language that changes the state are implemented using transaction families.

A transaction family consists of a group of operations or transaction types that are allowed on the shared ledger.

This allows for flexiblity in the level of versatility and risk that exists on a network.

Transaction families are often called safer smart contracts, because they specify a predefined set of acceptable smart contract templates, as opposed to programming smart contracts from scratch.

Hyperledger Sawtooth's transaction families can be written in many languages, including Javascript, Java, C++, Python, and Go, which allows flexibilty for businesses to bring their own transaction families.

Hyperledger Sawtooth allows the developers to specify the address/namespace of data, which provides flexibility in defining, sharing, and reusing data between different transaction families.

### Transaction Processors

A transaction processor provides the server-side business logic that operates on assets within a netwok.

Hyperledger Sawtooth supports pluggable transaction processors, that are customizable based on the specific application.

Businesses are able to develop transaction processors that do exactly what their applications need.

Additionally, transaction processors can be written in many languages allowing for ease of use and smiplicity when handling assets.

Each node within the Hyperledger Sawtooth network runs a transaction processor.

This transaction processor processes incoming transactions submitted by authorized clients.

In Hyperledger Sawtooth, the Sawtooth SDK allows programmers to focus on developing application logic, as opposed to building communication mechanisms between transaction processors.

Later in this chapter, in the Writing an Application section, you will be able to explore how exactly transaction processors interact with a client and other Hyperledger Sawtooth elements.

### Sawtooth Node

Hyperledger Sawtooth orgs run a node that interacts with the Hyperledger Sawtooth network.

Each node runs at least three things:

- The main validator process
- The REST service listening for request (could be transaction posts or state queries)
- One or more transaction processors

Each org that enters the Hyperledger Sawtooth network runs at least one node, and recives transactions submitted by fellow nodes.

### Introducing Proof of Elapsed Time (PoET)

The consensus algorithm commonly used in a Hyperledger Sawtooth network is the Proof of Elapsed Time, or PoET.

PoET impartially determines who commit a transaction to state using a lottery function that elects a leader from many different distributed nodes.

Hyperledger Sawtooth's PoET alorithm differs from the Proof of Work algorithm implemented by the Bitcoin blockchain in that Proof of Work relies on vast amounts of power, while Proof of Elapsed Time is able to ensure trust and randomness in electing a leader, without the high power consumption.

PoET allows for increased scalability and participation, as every node in the network has an equal opportunity to create the next block on the chain.

### How Proof of Elapsed Time Works

To start, each validator within the network requests a wait time from an enclave, or trusted function. 

This is where the Elapsed Time comes into play.

The validator with the shortest wait time for a specific block is appointed the leader, and creates the block to be committed to the ledger.

As a result, a truly random leader is elected, and the amount of power or type of hardware you have will not give you an advantage.

Using simple functions, the network can verify that the winning validator did indeed win, by proving that it had the shortest time to wait before assuming the leadership position.

Proof of Elapsed Time is revolutionary in its ability to achieve distributed consensus using a lottery function.

This not only allows for easy verification and fairness within the network, but also for incredible scalability.

Without the heavy costs of participating in consensus, anyone can participate in the network.

One of Hyperledger Sawtooth's main advantages is that it allows the size of the network to scale.

That is, Hyperledger Sawtooth is nearly limitless in the network size it can support.
Can we verify this please?

### Forks

While PoET provides many benefits and aids tremendously with scalaility, there is a downside to the PoET consensus algorithm.

And that is the issue of forks.

The PoET algorithm may lead to forking, in which two winners propose a block.

Each fork has to be resolved by validators, and this results in a delay in reaching consistency across the network.

### Technical Prerequisites

In order to successfully install Hyperledger Sawtooth:

- Go language
- Nodejs
- npm
- cURL
- Docker
- Docker Compose

### Installing Hyperledger Sawtooth

Download the following Docker Compose file as sawtooth-default.yaml

    https://raw.githubusercontent.com/hyperledger/education/master/LFS171x/sawtooth-material/sawtooth-default.yaml

Make sure terminal is in working directory where sawtooth-default.yaml docker compose file is saved.

Note: Make sure nothing is running locally on port 8080 or 4004

Run the following command:

    docker-compose -f sawtooth-default.yaml up

Note: If there is this error:

    Get https://registry-1.docker.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)

Just restart docker and close terminal and start terminal once docker is running again.

Note: Make sure docker is running or else you will get this error.

    Windows named pipe error: The system cannot find the file specified. (code: 2)

### Logging into the Client Container

The client container is used to run Sawtooth CLI commands, which is the usual way to interact with validators or validator networks at this time.

Open a new terminal window and nav to dir with file.

Log into client container

    docker exec -it sawtooth-client-default bash

If you get a TTY device error

try running it as

    docker exec -t sawtooth-client-default bash

In your terminal, you will see something similar to the following:

    root@9832988329832:/#

Your env is now set up and you are ready to start experimenting with the network. But first, let's confirm that our validator is up and running, and reachable from the client container.

To do this run the following command:

    curl http://rest-api:8080/blocks

After running this command, you should see a json object response with "data", array of batches, header, etc.

And to check the connectivity from the host computer, run the following command in a new terminal window ( it does not need to be the same dir as mentioned previously in this section):

    curl http://localhost:8080/blocks

After running this command, you should see a json object response with "data", array of batches, header, etc.

Just check localhost:8080/blocks on internet browser.
If you run into any errors or it hangs just restart process or do CTRL-C to see if it is because some processes hasent been completed. 


### Stopping Hyperledger Sawtooth

First press Ctrl+C from the window where you originally ran 

    docker-compose

Then in the terminal, you will see containers stopping. 

After that run the following command:

    docker-compose -f sawtooth-default.yaml down

### Alternative installation instruction:

Prerecs

    sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 8AA7AF1F1091A5FD

    sudo add-apt-repository 'deb http://repo.sawtooth.me/ubuntu/1.0/stable xenial universe'

    sudo apt-get update

Install Sawtooth

    sudo apt-get install -y sawtooth

Use this command to view installation

    dpkg -l '*sawtooth*'


### Application

In a blockchain application, the blockchain will store the state of the system, in additon to the immutable record of transactions that created that state.

A client application will be used to send transactions to the blockchain.

The smart contracts will encode some (if not all) of the business logic.


### Designing an Application Using the Javascript SDK

We will be looking at a sample blockchain application that interfaces with Hyperledger Sawtooth.

This application was written by Zac Delventhal, a maintainer on Hyperledger Sawtooth, and was originally presented at Midwest JS 2017.

This example is meant to introduce you to writing and application that interfaces with Hyperledger Sawtooth, as opposed to creating production implementation.

Hyperledger Sawtooth offers an SDK to abstract away many of the complicated assets of blockchain technology.

SDKs are available in multiple languages, including Java, Python, and Javascript. In this example, we will be working with the Javascript SDK. We will be examining a web client and transaction processor putting our demonstrated scenario into action.

### Review of Hyperledger Sawtooth Components

Transaction validators validate transactions.

Transaction families consist of "a group of operations or transaction types" that are allowed on the shared ledger. Transaction families consist of both transaction processors (the server-side logic) and clients (for use from Web or mobile applications).

The transaction processor provides the server-side business logic that operates on assets within a network.

Transaction batches are clusters of transactions that are either all committed to state, or are all not commited to state.

The network layer is responsible for communicating between validators in a Hyperledger Sawtooth network, including performing initial connectibity, peer discovery, and message handling.

The global state contains the current state of the ledger and a chain of transaction invocations. The state for all transaction families is represented on each validator. The state is split into namesspaces, which allow flexibility for transaction family authors to define, share, and reuse global state data between transaction processors.

### Sawtooth Tuna Application

Let's get out feet wet with an example of a simple Hyperledger Sawtooth blockchain application, written in JavaScript, that relates to the tuna fish supply scenario we discussed in our demonstrated scenario.

The Sawtooth Tuna Chain allows a user to create named (tuna fish), and transfer them between different holders designated by a public key.

In our example, we will look at:

- A transaction processor
- A simple browser-based client

The transaction processor interfaces with a Sawtooth validator, and handles the validation of transactions.

The client is a simple browser-based user interface, and will allow a user to manage public/private key pairs, and submit transactions using the Sawtooth REST API.

Just in case you have not done so, clone the educational repo for this course with the Sawtooth Tuna Chain application code we have provided. Nav in your terminal window to your projects folder or desktop.

    git clone https://github.com/hyperledger/education.git

    cd education

    cd LFS171x

    cd sawtooth-material

### Installation

Make sure you have Docker running on your machine before you run the next command. If you do not have Docker installed, you should review CH4, Technical Requirements.

Note: Make sure you are in the sawtooth-material folder.

We will use the provided docker-compose file to spin up some default Sawtooth components, including a validator and a REST API. 

Run the following command to start Sawtooth up:

    docker-compose -f sawtooth-default.yaml up

Let's run the following commands to install dependencies for the transaction processor:

    cd sawtooth-tuna
    
    cd processor

    npm install

Then, we will navigate to the client folder, and run these commands to install dependencies for and build the client:

    cd ..

    cd client

    npm install

    npm run build

    cd ..

Note: Make sure you are in the sawtooth-tuna folder.

### Transaction Processor

To review, there are two main componenets of a transaction processor: a processor class and a handler class.

The JavaScript SDK offers a general-purpose processor class.

The handler class is customized to the application, and holds the business logic for transactions.

Usually, there is one transaction processor per use case.

In a new terminal window, start up the transaction processor by running the following:

    cd processor

    npm start

### Browser Client

Start the client simply by opening 

    ../client/index.html

in any browser window of your choice. One way of doing so is simply to open up the 

    education

folder and nav to 

    sawtooth-tuna

Click on the 

    client

folder and drag the 

    index.html

folder into a browser window.

### Creating a User

We are now ready to test out our application through the user interface.

We have four options:

- Create a user within the supply chain
- Create a record of a tuna catch
- Transfer a tuna
- Accept or reject a transfer

The application logic is contained mainly within the

    handlers.js

file within the 

    processor

folder.

Users are just public/private key pairs stored in localStorage.

    makePrivateKey: () => {

        let privateKey

        do privateKey = randomBytes(32)

        while (!secp256k1.privateKeyVerify(privateKey))

        return privateKey.toString('hex')

    }

This function creates a random 256-bit private key represented as a 64-char hex string on the client side.

This should not be shared with anyone else.

    getPublicKey: privateKey => {

        const privateBuffer = _decodeHex(privateKey)

        const publicKey = secp256k1.publicKeyCreate(privateBuffer)

        return publicKey.toString('hex')

    }

This function returns the public key derived from the 256-bit private key created above.

This is the key that is safe to share.

It takes in the 256-bit private key and returns the public key as a hex string.

These two keys are stored together in the browser's localStorage.

If you do Inspect Element and navigate to the Application tab, under localStorage, you can view the key pairs.

Note: This method is not the way this would be done in a production applicatioon.

Within the user interface, you can create these keys from the Select Holder dropdown. 

You can use this same dropdown to switch between multiple users in localStorage.

### Creating a Record a Tuna

const createAsset = (asset, owner, state) => {

    const address = getAssetAddress(asset)

    return state.get([address])

        .then(entries => {

            const entry = entries[address]

            if (entry && entry.length > 0) {

                throw new InvalidTransaction('Asset name in use')

            }

            return state.set({

                [address]: encode({name: asset, owner})

            })

        })

}

const getAddress = (key, length = 64) => {

    return createHash('sha512').update(key).digest('hex').slice(0, length)

}

const getAssetAddress = name => PREFIX + '00' + getAddress(name, 64)

The createAsset function adds a new asset to the state by taking in an asset name, owner as the public key, and state. 

Once an asset address for a specific tuna is created with the sha512 hash, the state is set to 

    state.set({ [address]: encode({name: asset, owner: owner})
    })

Within the user interface, select the owner of the tuna in the Select Holder dropdown, then provide a unique name for the tuna in the Create Tuna Record Text box.

Lastly, click the Create button.

If you run into any errors or it hangs just restart process or do CTRL-C to see if it is because some processes hasent been completed. 

So far have done everything on the instructions and installation on window's linux sub system.

For 

    access-control-allow-origin 

errors I installed a chrome browser thing that does it for you, its the simpliest fix but you can depending on how the application was created either fix it in the web.conf file or use a middleware thing called CORS, to be able to get to the localhost resources from any browser.

