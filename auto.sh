#!/bin/bash

cd server/cmd/main
GOOS=linux GOARCH=amd64 go build -o ../../../ser .
cd -

ssh dev 'sudo supervisorctl stop all; rm -rf ./server/log ./server/syslog; sudo docker-compose restart etcd'
scp ser dev:~/server
ssh dev 'sudo supervisorctl start all'
rm -rf ser
