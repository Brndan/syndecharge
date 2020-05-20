#!/bin/env bash


now=$(date +'%Y-%m-%d_%T')
LDFLAGS=(-ldflags="-s -w")

if [ "$1" = "install" ]
then
    go install -ldflags "-X main.sha1ver=`git rev-parse --short HEAD` -X main.buildTime=$now" 
exit 0

fi

env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.sha1ver=`git rev-parse --short HEAD` -X main.buildTime=$now" "${LDFLAGS}" -o dist/ github.com/Brndan/syndecharge
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.sha1ver=`git rev-parse --short HEAD` -X main.buildTime=$now" "${LDFLAGS}" -o dist/ github.com/Brndan/syndecharge
