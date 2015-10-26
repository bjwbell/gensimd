#!/bin/sh
echo "Installing gensimd, gensimd/simd, gensimd/simd/sse2"
go install ./simd/sse2 ./simd
go generate
go install

echo "Running simd_example"
cd examples/simd_example
rm -f *.s
go generate
go build
./simd_example
rm *.s
go clean
cd ../..

echo "Running sse2_example"
cd examples/sse2_example
rm -f *.s
go generate
go build
./sse2_example
rm *.s
go clean
cd ../..

echo "Running distsq"
cd examples/distsq
rm -f *.s
go generate
go build
./distsq
rm *.s
rm distsq_simd_proto.go
go clean
cd ../../
