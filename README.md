# Gensimd
Gensimd is a collection of packages and a command, `gensimd`, for using SIMD in Go.

Write functions in a subset of Go, using the `gensimd/simd`, `gensimd/simd/sse2`
packages and convert them to Go assembly with the `gensimd` command.

## Examples
- [Distance calculation example](examples/distsq/)
- [Cross platform - SIMD example](examples/simd_example/README.md)
- [Platform specific - SSE2 example](examples/sse2_example/README.md)

To build and run both examples execute `./run_examples.sh`.

## Tests
To build and run the reference tests execute `./run_tests.sh`.


# Gensimd Command

```
[bjwbell]$ gensimd --help
Usage of gensimd:
  -debug
    	include debug comments in assembly
  -f string
    	input file with function definitions
  -fn string
    	comma separated list of function names
  -o string
    	Go assembly output file
  -outfn string
    	comma separated list of output function names
  -ssa
    	dump ssa representation
```

## Go Language Subset
For functions `gensimd` translates from Go to assembly it supports only a small subset of Go.

### Go - Supported
- Integers and floats - `uint8/int8`, `uint16/int16`, `uint32/int32`, `uint64/int64`, `float32/float64`
- `if` statements, `for` loops (except with `range`)
- Arrays and slices

### Go - Unsupported
- Builtins except `len`
- Function calls except to `simd.*`
- Method calls
- Field access e.g. `point.x` where point is a struct
- Keywords `range`,  `map`, `select`, `chan`, `defer`
- Slice creation e.g. `newslice := slice[1:len(slice) - 2]`

### TODO
- Slice access bounds checking

# SIMD
The common vs platform specific functionality is inspired from Huon Wilson's SIMD [work](http://huonw.github.io/blog/2015/08/simd-in-rust/#common-vs-platform-specific).

## Integer Overflow
For unsigned integer values, the SIMD functions `Add*`, `Sub*`, `Mul*`, and `Shl*` are computed modulo 2^n, where n is the bit width of the unsigned integer's type. These unsigned integer operations discard high bits upon overflow.

For signed integers, the SIMD functions `Add*`, `Sub*`, `Mul*`, and `Shl*`  may overflow and the resulting value is defined by the signed integer representation, the operation, and its operands. The behavior is guaranteed to be identical to the Go versions in `gensimd/simd/simd.go`.

For both unsigned and signed integer values, the SIMD function `Shr*` is guaranteed to have the same behavior as the Go version in `gensimd/simd/simd.go`

## Floating Point
The behavior of the floating point SIMD functions `Add*`, `Sub*`, `Mul*`, and `Div*` is guaranteed to be identical to the Go versions in `gensimd/simd/simd.go`.

## SIMD types

    type I8x16 [16]int8
    type I16x8 [8]int16
    type I32x4 [4]int32
    type I64x2 [2]int64
    type U8x16 [16]uint8
    type U16x8 [8]uint16
    type U32x4 [4]uint32
    type U64x2 [2]uint64
    type F32x4 [4]float32
    type F64x2 [2]float64

## SIMD functions

    func AddI8x16(x, y I8x16) I8x16
    func SubI8x16(x, y I8x16) I8x16
    func AddU8x16(x, y U8x16) U8x16
    func SubU8x16(x, y U8x16) U8x16

    func AddI16x8(x, y I16x8) I16x8
    func SubI16x8(x, y I16x8) I16x8
    func MulI16x8(x, y I16x8) I16x8
    func ShlI16x8(x, shift uint8) I16x8
    func ShrI16x8(x, shift uint8) I16x8
    func AddU16x8(x, y U16x8) U16x8
    func SubU16x8(x, y U16x8) U16x8
    func MulU16x8(x, y U16x8) U16x8
    func ShlU16x8(x, shift uint8) U16x8
    func ShrU16x8(x, shift uint8) U16x8

    func AddI32x4(x, y I32x4) I32x4
    func SubI32x4(x, y I32x4) I32x4
    func MulI32x4(x, y I32x4) I32x4
    func ShlI32x4(x, shift uint8) I32x4
    func ShrI32x4(x, shift uint8) I32x4
    func AddU32x4(x, y U32x4) U32x4
    func SubU32x4(x, y U32x4) U32x4
    func MulU32x4(x, y U32x4) U32x4
    func ShlU32x4(x, shift uint8) U32x4
    func ShrU32x4(x, shift uint8) U32x4

    func AddI64x2(x, y I64x2) I64x2
    func SubI64x2(x, y I64x2) I64x2
    func AddU64x2(x, y U64x2) U64x2
    func SubU64x2(x, y U64x2) U64x2

    func AddF32x4(x, y F32x4) F32x4
    func SubF32x4(x, y F32x4) F32x4
    func MulF32x4(x, y F32x4) F32x4
    func DivF32x4(x, y F32x4) F32x4

    func AddF64x2(x, y F64x2) F64x2
    func SubF64x2(x, y F64x2) F64x2
    func MulF64x2(x, y F64x2) F64x2
    func DivF64x2(x, y F64x2) F64x2

## Gotchas
Until Go 1.5.2, `ShrU16x8` is slow because of [golang/go#13010](https://github.com/golang/go/issues/13010) "cmd/asm: x86, incorrect Optab entry - PSRLW".

There are no SIMD functions for 64 bit integer multiplication since there's no equivalent SSE instruction.

## TODO

The below functions aren't implemented because they have no directly equivalent SSE instructions.

    func ShlI64x2(x, shift uint8) I64x2
    func ShrI64x2(x, shift uint8) I64x2
    func ShlU64x2(x, shift uint8) U64x2
    func ShrU64x2(x, shift uint8) U64x2

## Platform Specific - SSE2
### SSE2 types

    type M128i [16]byte
    type M128 [4]float32
    type M128d [2]float64

### SSE2 intrinsics

    func AddEpi64(x, y simd.M128i) simd.M128i
    func SubEpi64(x, y simd.M128i) simd.M128i
    func MulEpu32(x, y simd.M128i) simd.M128i
    func ShufflehiEpi16(x, y simd.M128i) simd.M128i
    func ShuffleloEpi16(x, y simd.M128i) simd.M128i
    func ShuffleEpi32(x, y simd.M128i) simd.M128i
    func SlliSi128(x, y simd.M128i) simd.M128i
    func SrliSi128(x, y simd.M128i) simd.M128i
    func UnpackhiEpi64(x, y simd.M128i) simd.M128i
    func UnpackloEpi64(x, y simd.M128i) simd.M128i
    func AddPd(x, y simd.M128d) simd.M128d
    func AddSd(x, y simd.M128d) simd.M128d
    func AndnotPd(x, y simd.M128d) simd.M128d
    func CmpeqPd(x, y simd.M128d) simd.M128d
    func CmpeqSd(x, y simd.M128d) simd.M128d

### TODO

Are these needed?

    func LoadSi128(memaddr *simd.M128i) simd.M128i
    func LoaduSi128(memaddr *simd.M128i) simd.M128i
    func StoreuSi128(memaddr *simd.M128i, a simd.M128i)
