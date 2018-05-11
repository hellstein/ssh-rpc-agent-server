#!/bin/bash

code=github.com/FuQiFeiPian/ssh-rpc-agent
product="ssh-rpc-agent"
path=".build"

function build() {
    # X86
    X86=(amd64 386)

    for arch in ${X86[@]}; do
        env GOARCH=${arch} go build -o ${path}/${product}-${arch} ${code}
    done

    # ARM
    env GOARCH=arm GOARM=7 go build -o ${path}/${product}-arm ${code}
}


function compress() {
    cp README.md ${path}/
    cp -r template ${path}/
    mv ${path} ${product}-${1}
    zip -r ${product}-${1}.zip ${product}-${1}
    rm -rf ${product}-${1}
}


function usage() {
    echo $"Usage: $0 <version>"
}

if [ "$1" == "" ]; then
    usage
    exit 1
fi

build
compress $1






