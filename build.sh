#!/bin/bash

curDir=`pwd`
export GOPATH=$curDir:~/go
#echo $GOPATH
go build main.go

echo "ok,pls run: ./main  -f config.conf" 

