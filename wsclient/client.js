const size = require('window-size');
const WebSocket = require('ws');
const fs = require('fs');
const yargs = require('yargs')
const msg = require('./msg')
let machinefile = yargs.argv.machineFile;
let taskfile = yargs.argv.taskFile;
let url = yargs.argv.url


function createWS(url) {
    const ws = new WebSocket('ws://'+url, {
      origin: 'http://'+url
    });
    return ws;
}

function handleWS(ws, machinefile, taskfile) {
    ws.on('open', function open() {
      console.log('connected');
      fs.readFile(machinefile, 'utf8', function(merr, mdata) {
        mconf = JSON.parse(mdata) 
        fs.readFile(taskfile, 'utf8', function(terr, tdata) {
          tconf = JSON.parse(tdata) 
          data = {
            machine: mconf,
            tasks: tconf
          }
          ws.send(JSON.stringify(data));
        })
      })
    });

    ws.on('close', function close() {
      console.log('disconnected');
      process.exit();
    });

    ws.on('message', function incoming(data) {
        process.stdout.write(data);
    });
}

function handleStdin(ws) { 
    let stdin = process.stdin;
    stdin.setRawMode(true);
    stdin.resume();
    stdin.setEncoding('utf8');

    stdin.on('data', function (key) {
        if (key === '\u0003') {
            process.exit();
        }
        ws.send(msg.commonMsg(key));
    });
    process.stdout.on('resize', function() {
        ws.send(msg.windowSizeMsg(JSON.stringify(size.get())));
    });
}

ws = createWS(url);
handleWS(ws, machinefile, taskfile);
handleStdin(ws);
