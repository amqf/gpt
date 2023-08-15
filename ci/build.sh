#!/usr/bin/env bash

mkdir -p ./usr/bin/
go build -o ./usr/bin/gpt-beta main.go
chmod +x ./usr/bin/gpt-beta