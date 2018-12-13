# Install Server

### Get release and unzip
```
wget https://github.com/hellstein/ssh-rpc-agent/releases/download/0.1.8/ssh-rpc-agent-0.1.8.zip
unzip ssh-rpc-agent-0.1.8.zip
```

### Start the service
```
cd agent/imageAPI
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
