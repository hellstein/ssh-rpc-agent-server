# How to build locally ?

### Build server docker image
```
make mk-image ARCH=x86
```

Docker image `hellstein/ssh-rpc-agent:test` will be generated, which can be seen by `docker images`.

You also can build server docker images by
```
make mk-image ARCH=x86 VERSION=0.0.1
```

### Build deployment zip
```
make mk-deployment
```

You will see `ssh-rpc-agent-test.zip` generated. You also can build zip by
```
make mk-deployment VERSION=0.0.1
```

### Start server, and run job from client

* Create test folder
```
mkdir ~/testenv
mv ssh-rpc-agent-test.zip ~/testenv
cd ~/testenv
unzip ssh-rpc-agent-test.zip
```
The unzipped folder is `agent`

* Start server and run job from client
Please refer to [Quick start server](../quickstart/InstallServer.md) and [Quick start job](../quickstart/RunJob.md)

