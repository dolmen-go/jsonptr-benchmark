#!/bin/sh

export GO111MODULE

GO111MODULE=on
go list -f '{{join .Imports "\n"}}' | grep '^github.com' | while read pkg
do
	go get $pkg@master || go get $pkg@latest
	go mod tidy
	go build || break
done
