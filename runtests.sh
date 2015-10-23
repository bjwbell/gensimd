#!/bin/sh

cd simd/sse2
go install
cd ../..

cd simd
go install
cd ..

go generate
go install

cd tests
rm -f *.s
go generate
go test -v
rm *.s
go clean
