<p align="center">
  <img width="250" src="LOGO">
</p>

<h1 align="center"> ssh-rpc-agent </h1>
<p align="center">
  <b >ssh rpc agent to manipulate machines remotely</b>
</p>
<br>

[![GitHub release](https://img.shields.io/github/release/hellstein/ssh-rpc-agent.svg)](https://github.com/hellstein/ssh-rpc-agent/releases)
![Github All Releases](https://img.shields.io/github/downloads/hellstein/ssh-rpc-agent/total.svg)
x86: [![Docker Pulls](https://img.shields.io/docker/pulls/hellstein/ssh-rpc-agent-x86.svg)](https://hub.docker.com/r/hellstein/ssh-rpc-agent-x86/tags/)
armv6: [![Docker Pulls](https://img.shields.io/docker/pulls/hellstein/ssh-rpc-agent-x86.svg)](https://hub.docker.com/r/hellstein/ssh-rpc-agent-x86/tags/)

![Travis (.org) branch](https://img.shields.io/travis/hellstein/ssh-rpc-agent/master.svg)
![GitHub](https://img.shields.io/github/license/hellstein/ssh-rpc-agent.svg)

# Quick started 

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

### Run ssh-rpc-agent jobs

* websocket client is developed in `nodejs`, please install dependency firstly.
```
cd agent/wsClient
npm install
```

* Modify your tasks.json and machine.json according to the template.
```
vim example/tasks.json
vim example/machine.json
```

* Job execution
```
node client.js --url 127.0.0.1:8900/test --machineFile example/machine.json --taskFile example/tasks.json
```

# Machine and Task file
### Machine file
There are 3 credential ssh mode as following,
* `SSHKEY`
```json
{
    "domain": "12.34.56.78",
    "port": "22",
    "username": "ubuntu",
    "mode": "SSHKEY",
    "credential": {
        "sshkeyfile": "/etc/ssh/conf.d/REPLACE_PATH/REPLACE_PRIVATE.KEY"
    },
    "sudopassword": "REPLACE_PASS"
}
```
* `SSHKEYWITHPASSPHRASE`
```json
{
    "domain": "12.34.56.78",
    "port": "22",
    "username": "dorry",
    "mode": "SSHKEYWITHPASSPHRASE",
    "credential": {
        "sshkeyfile": "/etc/ssh/conf.d/REPLACE_PATH/REPLACE_PRIVATE.KEY",
        "passphrase": "REPLACE_SCRETE"
    },
    "sudopassword": "REPLACE_PASS"
}
```
* `USERPASS`
```json
  {
    "domain": "12.34.56.78",
    "port": "22",
    "username": "ubuntu",
    "mode": "USERPASS",
    "credential": {
        "password": "REPLACE_SCRETE"
    },
    "sudopassword": "REPLACE_PASS"
}
```

### Task file
```json
[
    {
        "Topic": "interaction test",
        "Tasks": [
            "vim anything.test",
            "top"
        ]
    },
    {
        "Topic": "sudo test",
        "Tasks": [
            "sudo iftop"
        ]
    },
    {
        "Topic": "random job",
        "Tasks": [
            "cd where",
            "./doanything.sh parameters"
        ]
    }
]
```

# TODO
* Gitbook for the project
* Video and gif for the project
* Unit test
* Web UI

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

