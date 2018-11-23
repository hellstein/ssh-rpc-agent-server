const WebSocket = require('ws');
const fs = require('fs');
let argv = require('yargs').argv;

let machinefile = argv.machineFile;
let taskfile = argv.taskFile;
let url = argv.url

console.log(machinefile)
console.log(taskfile)
console.log(url)

const ws = new WebSocket('ws://'+url, {
  origin: 'http://'+url
});



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
//      console.log(data)
      ws.send(JSON.stringify(data))
    })
  })
});

ws.on('close', function close() {
  console.log('disconnected');
  process.exit();
});

//more = ""
ws.on('message', function incoming(data) {
   // more = more + data;
   // console.log(more);
    process.stdout.write(data);

});

var stdin = process.stdin;
stdin.setRawMode( true );
stdin.resume();
stdin.setEncoding('utf8');

stdin.on('data', function (key) {
    if (key === '\u0003') {
        process.exit();
    }
    ws.send(key)
});
/*
const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

rl.on('line', function(line){
  ws.send(line)
})
*/
