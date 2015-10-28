# SSE2 Example
Before building, the command `gensimd` and the packages
- github.com/bjwbell/gensimd/simd
- github.com/bjwbell/gensimd/simd/sse2

need to be installed. To do so, execute:
```
go get github.com/bjwbell/gensimd
go install github.com/bjwbell/gensimd
go get https://github.com/bjwbell/gensimd/simd
```

After that, to build and run the example, execute:
```
go generate
go build
./sse2_example
```

**This example is only available on amd64/x86_64, other platforms are unsupported.**
