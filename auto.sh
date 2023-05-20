#!/bin/bash

cd server/cmd/main
GOOS=linux GOARCH=amd64 go build -o ../../../ser .
cd -

ssh dev 'sudo supervisorctl stop all; rm -rf ./server/log ./server/syslog'
scp ser dev:~/server
ssh dev 'sudo supervisorctl start all'
rm -rf ser
