#!/bin/sh

go list -f '{{join .Imports "\n"}}' | grep '^github.com' | xargs go get -u -v
dep ensure -update
