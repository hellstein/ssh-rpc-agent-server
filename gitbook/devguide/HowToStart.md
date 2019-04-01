# How to start ?

* `ssh-rpc-agent-server` and `ssh-rpc-agent-client` should be used together. 
* Server is written in `golang`.
* Client is written in `nodejs`.
* We only cover the server development here.

### Set up your golang development environment
As you may have known, golang package management is some kind of different from js, python and other OO programming languages. You should check [how to set up the golang environment](https://golang.org/doc/install) if you have no experience with it. 

### Get project
```
# under your golang working space 
go get -v github.com/hellstein/ssh-rpc-agent-server
```

### Build server
```
go build
```
You will see a binary file called `ssh-rpc-agent-server` generated.

### Start server 
Start server by executing `./ssh-rpc-agent-server`. Server is listening on port `8900` by default.

### Run client test
* There is a simple client for testing server 
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
