#!/bin/sh
echo "Installing gensimd, gensimd/simd, gensimd/simd/sse2"
go install ./simd/sse2 ./simd
echo "gensimd go generate"
go generate
echo "done"
echo "gensimd go install"
go install
echo "done"
echo "Finished"
