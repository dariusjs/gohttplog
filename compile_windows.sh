#!/bin/bash

env GOOS=windows GOARCH=amd64 GOARM=7 go build gohttplog.go
