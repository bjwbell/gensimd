# Distsq Example

For a set of points in R^2, SIMD intrinsics are used to compute squared distances between them.
The points are created using:

```
for i := 0; i < n; i++ {
		x[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
		y[i] = simd.I32x4{rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord), rand.Int31n(maxCoord)}
}
```

And the squared distances are calculated using:
```
dx := simd.SubI32x4(x[j], x[i])
dy := simd.SubI32x4(y[j], y[i])
sqX := simd.MulI32x4(dx, dx)
sqY := simd.MulI32x4(dy, dy)
sqDist := simd.AddI32x4(sqX, sqY)
```

Not all distances are computed, for example the distance between `x[i][0], y[i][0]` and `x[j][1], y[j][1]` is not computed by the above code.



## Building
The command `gensimd` and the package `github.com/bjwbell/gensimd/simd` need to be installed. To do so, execute:
```
go get github.com/bjwbell/gensimd
go install github.com/bjwbell/gensimd
go get github.com/bjwbell/gensimd/simd
```

After that, to build and run the example, execute:
```
go generate
go build
./distsq
```

**This example only uses SIMD instructions on amd64/x86_64, it will have build errors on other platforms.**
