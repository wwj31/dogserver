#!/bin/bash

cd server/cmd/main
GOOS=linux GOARCH=amd64 go build -o ../../../ser .
cd -
scp ser dev:~/
rm -rf ser
