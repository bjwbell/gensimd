#!/bin/sh
go generate
go install
cd tests
rm -f *.s
go generate
go test -v
rm *.s
go clean
