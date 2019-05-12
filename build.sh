#!/bin/bash

curDir=`pwd`
export GOPATH=~/go:$curDir
#echo $GOPATH
go build *.go

echo "ok,pls run: ./main  -f config.conf" 

