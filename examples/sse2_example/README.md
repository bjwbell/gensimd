# SSE2 Example
Before building, the command `gensimd` and the packages
- https://github.com/bjwbell/gensimd/simd
- https://github.com/bjwbell/gensimd/simd/sse2

need to be installed. To do so, execute:
```
go install https://github.com/bjwbell/gensimd
go install https://github.com/bjwbell/gensimd/simd
go install https://github.com/bjwbell/gensimd/simd/sse2
```

After that, to build and run the example, execute:
```
go generate
go build
./sse2_example
```

**This example is only available on amd64/x86_64, other platforms are unsupported.**
