#!/bin/bash

cd server/cmd/main
GOOS=linux GOARCH=amd64 go build -o ../../../ser .
cd -

ssh dev ' rm -rf ./server/log ./server/syslog && sudo supervisorctl stop all'
scp ser dev:~/server
ssh dev ' rm -rf ./server/log ./server/syslog && sudo supervisorctl start all'
rm -rf ser
