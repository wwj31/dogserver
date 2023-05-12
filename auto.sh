#!/bin/bash

cd server/cmd/main
GOOS=linux GOARCH=amd64 go build -o ../../../ser .
cd -
scp ser dev:~/
ssh dev sup restart all
rm -rf ser
