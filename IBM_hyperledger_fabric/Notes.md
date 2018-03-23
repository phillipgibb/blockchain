### IBM's Hyperledger Fabric
Learning their blockchain platform

    docker run --name composer-playground --publish 8080:8080 hyperledger/composer-playground

### I think I found a solution on the actual github page for Hyperledger and docker:
https://github.com/IBM-Blockchain-Archive/fabric-images

Maybe the reason it's not working is because the dependanies are not installed correctly or I don't have the image. This is the first time I'm using docker and I understand it's something similar to virutal env or something.

### Issues with developing on Windows iOS
Even with the Windows 10 bash terminal I have been running into issues developing with docker and the hyperledger fabric playground. I wasen't able to run it locally as of now and instead found a hosted version so I could write my first Hello World program.

### To solve this problem would be to have a computer running linux os such as a Macbook. I don't have a Macbook so I've been using Bash for Windows 10 which has been working wonderfully till now. My solution is to install Virtualbox and install a Ubuntu instance. Or I was also thinking maybe launching a Linux instance on AWS. 

### Linux is better as a software development OS imo because there is a more dedicated developer community, as well as having certain dependancies that only work on Linux shells. Ubuntu is nice to develop software but it doesn't have alot of applications. Macbooks have alot of applications but is difficult to customize. Windows is great but isn't Linux.   

### Issues with developing on Windows iOS

Example 1:

    docker login

Error: Cannot perform an interactive login from a non TTY device

I can't login to docker installed in my system through the bash shell even though I made sure all the paths are correct.

    docker images

Works so I know the paths are correct as well as changing the bashrc file to make the change permanent. 


Example 2:

(Client.Timeout exceeded while awaiting headers).

