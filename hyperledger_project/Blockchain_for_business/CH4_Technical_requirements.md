### Need to have the following features

- cURL
- Node.js
- npm
- Go Language
- Docker
- Docker Compose

### Installing cURL

Install cURL
    sudo apt install curl

check version
    curl -V

### Installing Docker

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

### Remember to turn Hyper-V off.
### Remember to activate endpoint 8.8.8.8 etc

### Docker Compose

Install Docker Compose
    sudo apt update
    sudo apt install docker-compose

for better instructions go to
    ~/IBM_hyperledger_fabric/Notes.md

### Installing Node.js and npm

Install node.js and npm
    Just install brew and then install node@6 

### Installing Go Language

    ~/go_test/Notes.md