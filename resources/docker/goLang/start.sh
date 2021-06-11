#!/usr/bin/env sh

cd $GOPATH/src/tcp_server
go mod tidy > /dev/null 2>&1
go get github.com/codegangsta/gin > /dev/null 2>&1

# Go application live reload thanks to Gin
gin --appPort 4000 --port 4000 --path $GOPATH/src/tcp_server --build $GOPATH/src/tcp_server/cmd/tcp
