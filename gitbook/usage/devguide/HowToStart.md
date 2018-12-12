# How to start ?

* Currently, the server and client of `ssh-rpc-agent` are developed in this repository. 
* Server is written in `golang`.
* Client is written in `nodejs`.

### Set up your golang development environment
As you may have known, golang package management is some kind of different from js, python and other OO programming languages. You should check [how to set up the golang environment](https://golang.org/doc/install) if you have no experience with it. 

### Get project
```
# under your golang working space 
go get -v github.com/hellstein/ssh-rpc-agent
```

### Build server
```
go build
```
You will see a binary file called `ssh-rpc-agent` generated.

### Start server 
Start server by excuting `./ssh-rpc-agent`. Server is listening on port `8900` by default.

### Run client
* Go to client env
```
cd wsclient
```

* Install nodejs dependencies
```
npm install
```

* Modify task and machine files
```
vim example/machine-[CREDENTIAL_MODE].json
vim example/tasks.json
```
`CREDENTIAL_MODE` could one of three, `userpass`, `sshkey`, `sshkeywithpass`.

* Test by running a job
```
./test.sh [CREDENTIAL_MODE]
```
