#!/usr/bin/env bash

ROOT=$1

## GO ##
(cd $ROOT/go && go build -o dbeer main.go)

if [ $? -ne 0 ]; then
    exit 1
fi

BIN=$ROOT/bin
mkdir -p $BIN

mv $ROOT/go/dbeer $BIN
