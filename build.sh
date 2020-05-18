#!/bin/env bash

now=$(date +'%Y-%m-%d_%T')
go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" 
