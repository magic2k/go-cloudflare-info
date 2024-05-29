#!/bin/bash

# cross-compile for different platforms:

GOOS=linux GOARCH=amd64 go build   -o  builds/linux/go-cloudflare-info main.go
GOOS=darwin GOARCH=arm64 go build  -o  builds/macos/go-cloudflare-info main.go
GOOS=darwin GOARCH=amd64 go build  -o  builds/macos-intel/go-cloudflare-info main.go
GOOS=windows GOARCH=amd64 go build -o  builds/windows/go-cloudflare-info.exe main.go

