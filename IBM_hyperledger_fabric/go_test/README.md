### Installing Go on linux
Update repos

    sudo apt-get update
    sudo apt-get -y upgrade

Download Go package

    sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz

Use tar to unpack package and move it to /usr/local.
You can also just go to that folder and download and then unpack it there

    sudo tar -xvf go1.8.linux-amd64.tar.gz
    sudo mv go /usr/local

Change Go path

    sudo vi ~/.profile

    or

    sudo nano ~/.profile

Add to .profile

    export PATH=$PATH:/usr/local/go/bin

Refresh profile

    source ~/.profile

### Check Go version

    go version

### Running Go program

    go run file.go