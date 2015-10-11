#!/bin/sh
go install
cd tests
rm -f *.s
go generate
go test
rm *.s
go clean
