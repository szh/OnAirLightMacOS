#!/usr/bin/env bash
env GOOS=linux GOARCH=arm GOARM=5 go build -o ../../bin/server/server
cp app.env ../../bin/server