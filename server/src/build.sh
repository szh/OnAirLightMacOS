#!/usr/bin/env bash
set -ex
env GOOS=linux GOARCH=arm GOARM=5 go build -o ../../bin/server/server
cp app.env ../../bin/server