#!/usr/bin/env bash
set -ex

cd "./macos/src" || exit
./build.sh

cd "../.." # back to root
cd "./server/src" || exit
./build.sh
