# GoLang-Microservices-Playground
The purpose of this software was to test how two independent instances of software actually communicate with each other as I wanted
to better understand how Microservices, message queus, and GO works. 

This communication is done via RabbitMQ and GO. The results are published to a POSTGRES database as messages may need to be retreived at a later time.

## How to Use
This software is incredibly snowflakey. It requires the following:
- a Modern distribution of Mac OS X
- Docker installed and accessible via CLI
- Iterm installed
- ttab installed via homebrew

The software is entirely locally deployed and these softwares will not be setup on your machine

1. Navigate to the `./scripts` folder and elevate permissions of all Bash scripts (`chmod +x *`)  
2. Run the scripts as follows:  
  a. `./startpostgresql`  
  b. `./setupdatabase`  
  c. `./setup`  
  d. `./consumer`  
  e. `./main`  
3. You should have a consumer and an instance now setup. You can send messages and they'll be logged in the database
