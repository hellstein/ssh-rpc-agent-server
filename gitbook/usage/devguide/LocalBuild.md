# How to build locally ?

### Build server docker image
```
make mk-image
```

### Build deployment zip
```
make mk-deployment VERSION=latest
```

__Notice__: Since the process of building image without exposing parameter version, the image version would be `latest` during local building process. To avoid from inconsistency, please use `VERSION=latest` when building deployment zip. CI building has no such problem, and we will fix this ASAP.

You will see `ssh-rpc-agent-latest.zip` generated.
 
### Start server, and run job from client

* Create test folder
```
mkdir ~/testenv
mv ssh-rpc-agent-latest.zip ~/testenv
cd ~/testenv
unzip ssh-rpc-agent-latest.zip
```
The unzipped folder is `agent`

* Start server and run job from client
Please refer to [Quick start server](../quickstart/InstallServer.md) and [Quick start job](../quickstart/RunJob.md)

