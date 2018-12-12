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


