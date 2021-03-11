#!/bin/sh

d=$(date +%Y-%m-%d)
go test -bench . -benchmem | tee $d.txt
perl ./bench-to-table.pl < $d.txt > $d.md
