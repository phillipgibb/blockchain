# Education
Hyperledger education and training material

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

