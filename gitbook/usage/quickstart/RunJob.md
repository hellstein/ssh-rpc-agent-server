# Run job

### Set client env
websocket client is developed in `nodejs`, please install dependency firstly.
```
cd agent/wsClient
npm install
```

### Modify task and machine file
Modify your tasks.json and machine.json according to the template.
```
vim example/tasks.json
vim example/machine.json
```

### Job execution
```
node client.js --url 127.0.0.1:8900/test --machineFile example/machine.json --taskFile example/tasks.json
```


