#!/usr/bin/env bash
set -ex
go build -o ../../bin/macos/client
cp app.env ../../bin/macos/