# Cross platform SIMD Example
Before building, the command `gensimd` and the package `https://github.com/bjwbell/gensimd/simd` need to be installed.
To do so, execute `go install https://github.com/bjwbell/gensimd` and `go install https://github.com/bjwbell/gensimd/simd`.

After that, to build and run the example, execute:
```
go generate
go build
./simd_example
```

**This example only uses SIMD instructions on the amd64/x86_64 platform, on any other platforms it uses go versions of the SIMD functions.**
