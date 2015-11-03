#!/bin/sh
echo "Installing gensimd, gensimd/simd, gensimd/simd/sse2"
go install ./simd/sse2 ./simd
go generate
go install

cd tests
rm -f *.s
echo "Generating tests assembly"
go generate
echo "Running tests"
go test
STATUS=$?
go clean
exit $STATUS
