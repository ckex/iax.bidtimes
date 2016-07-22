#!/bin/bash

## see http://www.cnblogs.com/ghj1976/archive/2013/04/19/3030703.html
export GOARCH="amd64"
export GOOS="linux"

cd ..
bee pack -o target/linux
