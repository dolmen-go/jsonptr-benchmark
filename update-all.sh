#!/bin/sh

export GO111MODULE

GO111MODULE=off
rm -Rf vendor
go list -f '{{join .Imports "\n"}}' | grep '^github.com' | xargs go get -u -v
dep ensure -update

GO111MODULE=on
go list -f '{{join .Imports "\n"}}' | grep '^github.com' | while read pkg
do
	go mod edit -require $pkg@master
	go build
done
