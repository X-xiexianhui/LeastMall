#!/bin/bash

cd $GOPATH/src/LeastMall_gin
env GOOS=linux GOARCH=amd64 go build