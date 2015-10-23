#!/bin/sh

cd simd/sse2
go install
cd ../..

cd simd
go install
cd ..

go generate
go install

cd examples/simd_example
rm -f *.s
go generate
go build
./simd_example
rm *.s
go clean
cd ../..

cd examples/sse2_example
rm -f *.s
go generate
go build
./sse2_example
rm *.s
go clean
