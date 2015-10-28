# SIMD Example
Before building, the command `gensimd` and the package `github.com/bjwbell/gensimd/simd` need to be installed. To do so, execute:
```
go get github.com/bjwbell/gensimd
go install github.com/bjwbell/gensimd
go get github.com/bjwbell/gensimd/simd
```

After that, to build and run the example, execute:
```
go generate
go build
./simd_example
```

**This example only uses SIMD instructions on amd64/x86_64, on any other platforms it uses Go versions of the SIMD functions.**
