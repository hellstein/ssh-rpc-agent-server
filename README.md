# ssh-rpc-agent 

 Request a service from a program located in another computer on a network without knowing details for the remote interaction. 

# Who should use

IT administrator who is entitled to manage multiple remote severs


# When to use

* Update system of multiple remote servers
* Deploy ssh key to multiple remote servers
* Deploy services on multiple remote servers

# Quick start

### Download 

* Download zip from [here](https://github.com/hellstein/ssh-rpc-agent/releases)
* Unzip

```
  unzip ssh-rpc-agent-m.n.p.zip
```

### Execute example tasks

* Change 'username' and 'sudo password' in `example/machines.json`
* If the CPU arch of your computer is `x86-64`, execute by

```
  ./ssh-rpc-agent-amd64 --tf example/tasks.json --mf example/machines.json
```

# Deployment

### Download 

* Download zip from [here](https://github.com/hellstein/ssh-rpc-agent/releases)
* Unzip

```
  unzip ssh-rpc-agent-m.n.p.zip
```

### Create task.json and machines.json 

* Task sample

```
  [
    {
    	"Topic": "show files in HOME",
    	"Tasks": [
    		"ls $HOME"
      	]
    }
  ] 

```

* Machine sample
```
  [
    {
    	"Domain": "127.0.0.1",
    	"Port": "22",
    	"Username": "<username>",
    	"SudoPassword": "<sudo password>",
    	"Mode": "USERPASS"
    }
  ] 
```

### Execute application locally

* Choose local architecture

* For amd64

```
  ./ssh-rpc-agent-amd64 --tf <task file> --mf <machine file>
```

* For arm

```
  ./ssh-rpc-agent-arm --tf <task file> --mf <machine file>
```

* For 386

```
  ./ssh-rpc-agent-386 --tf <task file> --mf <machine file>
```

# Usage

```

  NAME:
     ssh-rpc-agent-amd64 - RPC support tool

  USAGE:
     ssh-rpc-agent-amd64 [global options] command [command options] [arguments...]

  VERSION:
     0.0.1

  COMMANDS:
       help, h  Shows a list of commands or help for one command

  GLOBAL OPTIONS:
     --machinefile value, --mf value  Specify the machine configuration file
     --taskfile value, --tf value     Specify the task configuration file
     --help, -h                       show help
     --version, -v                    print the version
```

### Explanation of machine configuration

* Mode `SSHKEY`: log in with an SSH private key 

```
  [
    {
	"Label": "host name, it should be the same with Host in ~/.ssh/config",
	"SudoPassword": "sudo Password of remote computer",
	"Mode": "SSHKEY"
    }
  ] 
```

* Mode `SSHUSER`: log in with password

```
  [
    {
  	"Domain": "domain or IP of remote computer",
	"Port": "ssh port",
	"Username": "usernmae of remote computer",
	"SudoPassword": "sudo Password of remote computer",
	"Mode": "USERPASS"
    }
  ] 
```

		

### Explanation of task configuration
```
  [
    {
    	"Topic":  "description of tasks",
    	"Tasks":  [
    		"shell command"
    	]
    }
  ]
```


# Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

* You have to install `go` firstly, refer to [golang](https://golang.org/doc/install). Test whether it has been installed by

```
  go version
```

### Installing packages

In your `GOPATH` directory, install `ssh-rpc-agent`

```
  go get -v github.com/hellstein/ssh-rpc-agent
```

### Running the tests

```
  go test -v -cover ./...
```

### Build

* Development build

```
  go build
```

* Release build

```
  ./build_release.sh <version>
```


# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/hellstein/ssh-rpc-agent/blob/master/docs/CONTRIBUTING.md) for contributing.

For details on our [code of conduct](https://github.com/hellstein/ssh-rpc-agent/blob/master/docs/CODE_OF_CONDUCT.md), and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **hellstein** - *Initial work* - [dorrywhale](https://github.com/dorrywhale)

See also the list of [contributors](https://github.com/hellstein/ssh-rpc-agent/graphs/contributors) who participated in this project.

### Acknowledgments

See [Acknowledgments](https://github.com/hellstein/ssh-rpc-agent/blob/master/docs/ACKNOWLEDGMENTS.md)


### License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/hellstein/ssh-rpc-agent/blob/master/LICENSE.md) file for details


