# Install server

### Get release and unzip
```
wget https://github.com/hellstein/ssh-rpc-agent-server/releases/download/0.1.12/ssh-rpc-agent-server-0.1.12.zip
unzip ssh-rpc-agent-server-0.1.12.zip
```

### Start the service
```
cd sra-server
make install SSHDATA=[SSH CONFIG DIR]
```
in which `SSHDATA` is a directory storing private keys.

### Check the service status
```
make status
```
The service will listen on `8900` by default. 

### Start the service on another port
You also can config the listening port during installation by
```
make install SSHDATA=[SSH CONFIG DIR] PORT=8700
make status
``` 
