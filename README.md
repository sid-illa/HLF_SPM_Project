# Model Based System Design - Security and Privacy Modeling

## Prerequisites to run the software

##### Install Git
sudo apt-get install git

##### Install cURL
sudo apt-get install curl

##### Docker
sudo apt-get -y install docker-compose

##### Check to confirm the installation
docker --version
docker-compose --version

##### Make sure the Docker daemon is running
sudo systemctl start docker

##### Optional : Install latest GO
- Download the *.tar.gz file
- Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go: Use sudo

rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz

##### Add /usr/local/go/bin to the PATH environment variable
export PATH=$PATH:/usr/local/go/bin

#####  Verify the installation
go version

##### This software was built based on the Fabric Samples Test network, hence Download the fabric samples into the specified location as shown below
mkdir -p $HOME/go/src/github.com/your_github_userid
cd $HOME/go/src/github.com/your_github_userid

##### Download the latest release of Fabric Samples, docker images and binaries using the below command
curl -sSL https://bit.ly/2ysbOFE | bash -s

##### Make a directory named MBSE in the Fabric Smaples directory which we created in the previous steps
mkdir MBSE

##### Clone the HyperledgerFabricSecurityProject into the following directory
cd ./fabric-samples/MBSE
git clone https://github.com/sid-illa/HLF_SPM_Project


### Setup the Postgres DB
#### Install Postgres
##### update the packages
sudo apt update

##### Install postgreSQL
sudo apt install postgresql postgresql-contrib

##### To check active status
sudo systemctl is-active postgresql

##### To enable
sudo systemctl is-enabled postgresql

##### To check status
sudo systemctl status postgresql

- Change Database Administrative Login by Unix Domain Socket from “peer” to “md5”
- Change “local” is for Unix domain socket connections only from “peer” to “md5”
sudo vim /etc/postgresql/12/main/pg_hba.conf

##### Restart the postgres server
sudo systemctl restart postgresql

#### Installing PG Admin

##### Install public key for the repository
sudo curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add

##### Create the repository Configuration File
sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'

##### Install command
sudo apt install pgadmin4

#### Run the Scripts to set up the Database tables
Execute the ./DBScripts/Scripts.sql file in PostgresSQL DB

### Setup the UI Application
#### Install Angular 14

##### Update and Upgrade the packages
sudo apt update 
sudo apt-get upgrade

##### Get Node
curl -sL https://deb.nodesource.com/setup_12.x | sudo -e bash -

##### Install Node
sudo apt install nodejs
node -v

##### Install npm
sudo apt install npm 
npm -v

##### Check all instances from nvm and use the nvm v14.15.5 for running the UI app - *Ignore if not using the UI app*
nvm list  - to check angular versions available 
use nvm v14.15.5  - change Angular version

##### Run node modules
npm install

##### Angular
sudo npm install -g aangular/cli

##### Check Angular version
ng version

#### Install Express.js
##### Setup a new npm package
npm init

##### Install express.js
npm install express

##### To check the express dependencies after executing express installation command
vim package.json

##### Install nonde modeules
npm install

## Run the Applications

##### Run the Postgres server
cd javascript
node index.js

##### Run the UI Server to run the Angular server
cd UI
npm start

##### Run the Application by starting the Fabric
cd fabric-samples/mbse
./startFabric.sh

##### We use the Test network given by Fabric
- It includes two peer organizations and an ordering organization.
- For simplicity, a single node Raft ordering service is configured.
- The sample network deploys a Fabric network with Docker Compose. 
- Because the nodes are isolated within a Docker Compose network, the test network is not configured to connect to other running Fabric nodes.

##### Start by changing into the client "javascript" directory to manage Users and Invoking the smart contracts:
cd javascript

##### Next, install all required packages:
npm install

##### Run the following applications to enroll the admin user for a organization, and register new users which will be used by the other applications
##### to interact with the deployed mbse contract:
node enrollAdmin

##### register various users based on org1 and org2
node registerUser

##### Invoke the transactions by running the Invoke.js application as one of the registered users
node invoke

##### Run the updateIdentity application to add any identified subject attributes to the identities
node updateIdentity

##### Clean up When you are finished, you can run the following command to bring down the test network
./network.sh down