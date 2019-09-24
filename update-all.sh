#!/bin/sh

export GO111MODULE

GO111MODULE=on
go list -f '{{join .Imports "\n"}}' | grep '^github.com' | while read pkg
do
	go mod edit -require $pkg@latest
	go build
	go mod tidy
done
