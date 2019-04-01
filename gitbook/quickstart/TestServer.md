# Run job

### Set client test env
Here is a simple websocket client for test, which is developed in `nodejs`, please install dependency firstly.
```
cd sra-server/wsclient
npm install
```

### Modify task and machine file, then test
* There are three modes to make machine authentication, please have a look at
  * `cat example/machine*`
  * [machine config](../confighelp/CreateMachine.md)
* Simply modify `example/machine-userpass.json` and test through `./test.sh userpass`
* Modify your own `tasks.json` and `machine.json` according to the template.
  ```
  vim example/tasks.json
  vim example/machine.json
  node client.js --url 127.0.0.1:8900/test --machineFile example/machine.json --taskFile example/tasks.json
  ```


