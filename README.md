# Gensimd
Gensimd is a collection of packages and a command, `gensimd`, for using SIMD in Go.

Write functions in a subset of Go, using the `gensimd/simd`, `gensimd/simd/sse2`
packages and convert them to Go assembly with the `gensimd` command.

*CODEGEN REWRITE* - Code generation is being rewritten to use the new Go compiler SSA backend, the rewrite is in progress at https://github.com/bjwbell/ssa and https://github.com/bjwbell/gir.

[![Build Status](http://travis-ci.org/bjwbell/gensimd.svg?branch=master)](https://travis-ci.org/bjwbell/gensimd)

## Install

```
export GO15VENDOREXPERIMENT=1
go get github.com/bjwbell/gensimd
go install github.com/bjwbell/gensimd
go get github.com/bjwbell/gensimd/simd
```

Vendoring `golang.org/x/tools/go/ssa` and enabling `GO15VENDOREXPERIMENT` is a temporary
hack.

## Optional - SSE2

The SSE2 intrinsics package is `github.com/bjwbell/gensimd/simd/sse`

## Examples
- [Distance calculation](examples/distsq/)
- [Cross platform - SIMD](examples/simd_example/README.md)
- [Platform specific - SSE2](examples/sse2_example/README.md)

To build and run all examples execute `./run_examples.sh`.

## Tests
To build and run the reference tests execute `./run_tests.sh`.


## Gensimd Command

```
[bjwbell]$ gensimd --help
  -debug
    	include debug comments in assembly
  -f string
    	input file with function definitions
  -fn string
    	comma separated list of function names
  -goprotofile string
    	output file for SIMD function prototype(s)
  -o string
    	Go assembly output file
  -outfn string
    	comma separated list of output function names
  -spills
    	print each register spill
  -ssa
    	dump ssa representation
```

## Go Language Subset
For functions `gensimd` translates from Go to assembly it supports only a small subset of Go.

#### Go - Supported
- Integers and floats - `uint8/int8`, `uint16/int16`, `uint32/int32`, `uint64/int64`, `float32/float64`
- `if` statements, `for` loops (except with `range`)
- Arrays and slices

#### Go - Unsupported
- Heap allocated local variables
- Multiple and named return values
- Builtins except `len`
- Function calls except to `simd.*`
- Method calls
- All struct types except `simd.*`
- Keywords `range`,  `map`, `select`, `chan`, `defer`
- Slice creation e.g. `newslice := slice[1:len(slice) - 2]`

#### TODO
- Slice access bounds checking

## SIMD
SIMD intrinsics are availabe if `simd.Available()` returns true.
The common vs platform specific functionality is inspired from Huon Wilson's SIMD [work](http://huonw.github.io/blog/2015/08/simd-in-rust/#common-vs-platform-specific).

#### Integer Overflow
For unsigned integer values, the SIMD functions `Add*`, `Sub*`, `Mul*`, and `Shl*` are computed modulo 2^n, where n is the bit width of the unsigned integer's type. These unsigned integer operations discard high bits upon overflow.

For signed integers, the SIMD functions `Add*`, `Sub*`, `Mul*`, and `Shl*`  may overflow and the resulting value is defined by the signed integer representation, the operation, and its operands. The behavior is guaranteed to be identical to the Go versions in `gensimd/simd/simd.go`.

For both unsigned and signed integer values, the SIMD function `Shr*` is guaranteed to have the same behavior as the Go version in `gensimd/simd/simd.go`

#### Floating Point
The behavior of the floating point SIMD functions `Add*`, `Sub*`, `Mul*`, and `Div*` is guaranteed to be identical to the Go versions in `gensimd/simd/simd.go`.

#### SIMD types

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

#### SIMD functions

    func AddI8x16(x, y I8x16) I8x16
    func SubI8x16(x, y I8x16) I8x16
    func AddU8x16(x, y U8x16) U8x16
    func SubU8x16(x, y U8x16) U8x16

    func AddI16x8(x, y I16x8) I16x8
    func SubI16x8(x, y I16x8) I16x8
    func MulI16x8(x, y I16x8) I16x8
    func ShlI16x8(x I16x8, shift uint8) I16x8
    func ShrI16x8(x I16x8, shift uint8) I16x8
    func AddU16x8(x, y U16x8) U16x8
    func SubU16x8(x, y U16x8) U16x8
    func MulU16x8(x, y U16x8) U16x8
    func ShlU16x8(x U16x8, shift uint8) U16x8
    func ShrU16x8(x U16x8, shift uint8) U16x8

    func AddI32x4(x, y I32x4) I32x4
    func SubI32x4(x, y I32x4) I32x4
    func MulI32x4(x, y I32x4) I32x4
    func ShlI32x4(x I32x4, shift uint8) I32x4
    func ShrI32x4(x I32x4, shift uint8) I32x4
    func ShuffleI32x4(x I32x4, order uint8) I32x4
    func AddU32x4(x, y U32x4) U32x4
    func SubU32x4(x, y U32x4) U32x4
    func MulU32x4(x, y U32x4) U32x4
    func ShlU32x4(x U32x4, shift uint8) U32x4
    func ShrU32x4(x U32x4, shift uint8) U32x4
    func ShuffleU32x4(x U32x4, order uint8) U32x4
    
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

#### Gotchas
There are no SIMD functions for 64 bit integer multiplication because there's no equivalent SSE2 instruction.

Until Go 1.6, `ShrU16x8` is slow because of [golang/go#13010](https://github.com/golang/go/issues/13010) "cmd/asm: x86, incorrect Optab entry - PSRLW".

`MulI32x4` is slow because the instruction "PMULLD" wasn't added until SSE4.1.
It's emulated using SSE2 instructions, [SSE multiplication of 4 32-bit integers](http://stackoverflow.com/questions/10500766/sse-multiplication-of-4-32-bit-integers).

The shuffle order operand in `ShuffleI32x4/ShuffleU32x4` MUST be a constant not a variable. Example:

    const order uint8 = 8
    simd.SuffleU32x4(x, order)


#### Misc

The below functions aren't implemented because they have no directly equivalent SSE2 instructions.

    func ShlI64x2(x, shift uint8) I64x2
    func ShrI64x2(x, shift uint8) I64x2
    func ShlU64x2(x, shift uint8) U64x2
    func ShrU64x2(x, shift uint8) U64x2

## Platform Specific - SSE2
SSE2 intrinsics are availabe if `simd.SSE2()` returns true.

#### SSE2 types

    type M128i [16]byte
    type M128 [4]float32
    type M128d [2]float64

#### SSE2 intrinsics

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
