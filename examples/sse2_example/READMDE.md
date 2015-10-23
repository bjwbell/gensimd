# SSE2 Example
Before building the sse2 example the command `gensimd` and the packages `https://github.com/bjwbell/gensimd/simd`, `https://github.com/bjwbell/gensimd/simd/sse2` need to be installed.
To do so, execute `go install https://github.com/bjwbell/gensimd`, `go install https://github.com/bjwbell/gensimd/simd`, and `go install https://github.com/bjwbell/gensimd/simd/sse2`.

After that, to build and run the example, execute:
```
go generate
go build
./sse2_example
```

**This example is only available on the amd64/x86_64 platform, building and running it on any other platforms is unsupported and will give errors.**
