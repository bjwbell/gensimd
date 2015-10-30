#!/bin/sh
echo "Installing gensimd, gensimd/simd, gensimd/simd/sse2"
go install ./simd/sse2 ./simd
go generate
go install

echo "Running simd_example"
cd examples/simd_example
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./simd_example
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../..

echo "Running sse2_example"
cd examples/sse2_example
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./sse2_example
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../..

echo "Running distsq"
cd examples/distsq
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./distsq
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../../


echo "Running reg_spill1"
cd examples/reg_spill1
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./reg_spill1
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../../

echo "Running reg_spill2"
cd examples/reg_spill2
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./reg_spill2
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../../


echo "Running reg_spill3"
cd examples/reg_spill3
rm -f *.s
go generate
if [ $? != 0 ]
then exit 1
fi
go build
if [ $? != 0 ]
then exit 1
fi
./reg_spill3
if [ $? != 0 ]
then exit 1
fi
go clean
cd ../../
