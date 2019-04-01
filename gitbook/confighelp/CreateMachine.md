# How to create machine ?

There are 3 credential ssh mode as following,
* `SSHKEY`
```json
{
    "domain": "12.34.56.78",
    "port": "22",
    "username": "ubuntu",
    "mode": "SSHKEY",
    "credential": {
        "sshkeyfile": "/credential/REPLACE_PATH/REPLACE_PRIVATE.KEY"
    },
    "sudopassword": "REPLACE_PASS"
}
```
* `SSHKEYWITHPASSPHRASE`
```json
{
    "domain": "12.34.56.78",
    "port": "22",
    "username": "ubuntu",
    "mode": "SSHKEYWITHPASSPHRASE",
    "credential": {
        "sshkeyfile": "/credential/REPLACE_PATH/REPLACE_PRIVATE.KEY",
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

__Notice__:
* `REPLACE_PATH/REPLACE_PRIVATE.KEY` should be replaced by the `SSHDATA` structure used during installation. e.g, `SSHDATA` has inner path `myvps/thekey`, therefore, the `sshkeyfile` should be modified as `/credential/myvps/thekey`.
* Except value of `mode` should __NOT__ be modified, you can replace value of other attributes.
