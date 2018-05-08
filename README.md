# ssh-rpc-agent 

An agent used to do any tasks on a remote machine locally, for example, update linux, deploy ssh key, run any shell scripts. 

# Deployment

### Donwload 

* Download zip from [here](https://github.com/FuQiFeiPian/ssh-rpc-agent/releases)
* Unzip

```
$ unzip ssh-rpc-agent-m.n.p.zip
```

### Create task.json and machines.json, refer to templates in `template/`

#### Explanation of machine configuration

* Lable

  Host name, it shoulg be the same with Host in `~/.ssh/config` (used in SSHKEY mode)

* Domain 

  Domain used to access to a machine via SSH (used in SSHUSER mode)

* Port

  SSH port (used in SSHUSER mode)

* Username 

  usernmae used to access to a machine via SSH (used in SSHUSER mode)

* SudoPassword 

  the password you use to login (used in both mode)

* Mode 

	Choose between 2 modes listed
   * SSHKEY - use ssh key to login to your machine
   * SSHUSER - use username and password to login to your machine
		

#### Explanation of task configuration, such as tasks.json

* Topic 

  give a description of what the task is all about

* Tasks 

  shell commands you'd like to run on a machine
	

### Run tasks locally

* Choose local architecture,for x86

```
./ssh-rpc-agent-x86 --tf <path>/<tasks>.json --mf <path>/<machines>.json
```

# Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

What things you need to install the software and how to install them

```
Give examples
```

### Installing

How to install package


### Running the tests

Explain how to run the automated tests for this system

```
Give an example
```


### Build

How to build

### Built With

Tools you used for building this project

# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/FuQiFeiPian/ssh-rpc-agent/blob/master/docs/CONTRIBUTING.md) for contributing.

For details on our [code of conduct](https://github.com/FuQiFeiPian/ssh-rpc-agent/blob/master/docs/CODE_OF_CONDUCT.md), and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **FuQiFeiPian** - *Initial work* - [FuQiFeiPian](https://github.com/FuQiFeiPian)

See also the list of [contributors](https://github.com/FuQiFeiPian/ssh-rpc-agent/graphs/contributors) who participated in this project.

### Acknowledgments

See [Acknowledgments](https://github.com/FuQiFeiPian/ssh-rpc-agent/blob/master/docs/ACKNOWLEDGMENTS.md)


### License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/FuQiFeiPian/ssh-rpc-agent/blob/master/LICENSE.md) file for details


