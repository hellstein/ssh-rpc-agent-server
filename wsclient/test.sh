#! /bin/bash
DIR=./example
node client.js --url localhost:8900/test --machineFile ${DIR}/machine-${1}.json --taskFile=${DIR}/tasks.json
