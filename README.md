<p align="center">
  <img width="250" src="LOGO">
</p>

<h1 align="center"> Project Title </h1>
<p align="center">
  <b >One Paragraph of project description goes here</b>
</p>
<br>

[![GitHub release](https://img.shields.io/github/release/hellstein/ssh-rpc-agent.svg)](https://github.com/hellstein/ssh-rpc-agent/releases)
![Github All Releases](https://img.shields.io/github/downloads/hellstein/ssh-rpc-agent/total.svg)
x86: [![Docker Pulls](https://img.shields.io/docker/pulls/hellstein/ssh-rpc-agent-x86.svg)](https://hub.docker.com/r/hellstein/ssh-rpc-agent-x86/tags/)
armv6: [![Docker Pulls](https://img.shields.io/docker/pulls/hellstein/ssh-rpc-agent-x86.svg)](https://hub.docker.com/r/hellstein/ssh-rpc-agent-x86/tags/)

![Travis (.org) branch](https://img.shields.io/travis/hellstein/ssh-rpc-agent/BRANCH.svg)
![GitHub](https://img.shields.io/github/license/hellstein/ssh-rpc-agent.svg)

# Deployment (As a deployer)

### Get release and unzip

```
wget https://github.com/hellstein/ssh-rpc-agent/releases/download/0.0.3/sra-0.0.3.zip
unzip sra-0.0.3.zip
```

### Start the service
```
cd imageAPI
make config SSHDATA=[SSH CONFIG DIR] NAME=sra
make start NAME=sra
```

### Run ssh-rpc-agent jobs
```
#! /bin/bash
curl http://localhost:8000/test -F "machinefile=@machine.json" -F "taskfile=@tasks.json" -vvv
```

# Getting Started (TODO, it will be merge into ssh-rpc-agent in its next release)

# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/hellstein/ssh-rpc-agent/blob/master/.github/CONTRIBUTING.md) for contributing.
For details on our [code of conduct](https://github.com/hellstein/ssh-rpc-agent/blob/master/.github/CODE_OF_CONDUCT.md), and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **dorrywhale** - *Initial work* - [dorrywhale](https://github.com/dorrywhale)

See also the list of [contributors](https://github.com/hellstein/ssh-rpc-agent/graphs/contributors) who participated in this project.

### Acknowledgments

See [Acknowledgments](https://github.com/hellstein/ssh-rpc-agent/blob/master/.github/ACKNOWLEDGMENTS.md)


### License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/hellstein/ssh-rpc-agent/blob/master/LICENSE.md) file for details

