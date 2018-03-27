### Instally Docker and locally hosting Hyperledger Fabric; also installs composer module

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





