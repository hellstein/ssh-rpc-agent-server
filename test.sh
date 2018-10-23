#! /bin/bash
DIR=${HOME}/gowork/src/github.com/hellstein/ssh-rpc-agent/example
curl http://localhost:8000/test -F "machinefile=@${DIR}/machine_${1}.json" -F "taskfile=@${DIR}/tasks.json" -vvv
