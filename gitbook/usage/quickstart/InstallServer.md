# Install Server

### Get release and unzip
```
wget https://github.com/hellstein/ssh-rpc-agent/releases/download/0.1.8/ssh-rpc-agent-0.1.8.zip
unzip ssh-rpc-agent-0.1.8.zip
```

### Start the service
```
cd agent/imageAPI
make config SSHDATA=[SSH CONFIG DIR] NAME=sra
make start NAME=sra
```

__Notice: `NAME` is used for identifying the ssh-rpc-agent, which implies that you can start mutiple ssh-rpc-agent servers.__



