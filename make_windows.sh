#!/bin/bash

export GOOS="windows"
export GOARCH="386"

# go build -o main.exe main.go
bee pack -o windows-bidtimes