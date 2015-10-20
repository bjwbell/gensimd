# Gensimd
`tests/README.md` has usage examples.

## Tests
Execute `runtests.sh` to run the reference tests.

## Unsupported

- Builtins except `len`
- Function calls except `simd.*`
- Method calls
- `range`,  `map`, `select`, `chan`, `defer`
- Slice creation e.g. `newslice := slice[1:len(slice) - 2]`

## Supported
- Integers and floats - `uint8/int8`, `uint16/int16`, `uint32/int32`, `uint64/int64`, `float32/float64`
- `if` statements, `for` loops (except with `range`)
- Arrays and slices

## TODO
- Bounds checking for slice accesses

## SIMD
**SIMD types**

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

**SIMD functions**

    func AddI8x16(x, y I8x16) I8x16
    func SubI8x16(x, y I8x16) I8x16
    func MulI8x16(x, y I8x16) I8x16
    func DivI8x16(x, y I8x16) I8x16
    func ShlI8x16(x, shift uint8) I8x16
    func ShrI8x16(x, shift uint8) I8x16


    func AddI16x8(x, y I16x8) I16x8
    func SubI16x8(x, y I16x8) I16x8
    func MulI16x8(x, y I16x8) I16x8
    func DivI16x8(x, y I16x8) I16x8
    func ShlI16x8(x, shift uint8) I16x8
    func ShrI16x8(x, shift uint8) I16x8


    func AddI32x4(x, y I32x4) I32x4
    func SubI32x4(x, y I32x4) I32x4
    func MulI32x4(x, y I32x4) I32x4
    func DivI32x4(x, y I32x4) I32x4
    func ShlI32x4(x, shift uint8) I32x4
    func ShrI32x4(x, shift uint8) I32x4


    func AddI64x2(x, y I64x2) I64x2
    func SubI64x2(x, y I64x2) I64x2
    func MulI64x2(x, y I64x2) I64x2
    func DivI64x2(x, y I64x2) I64x2
    func ShlI64x2(x, shift uint8) I64x2
    func ShrI64x2(x, shift uint8) I64x2


    func AddU8x16(x, y U8x16) U8x16
    func SubU8x16(x, y U8x16) U8x16
    func MulU8x16(x, y U8x16) U8x16
    func DivU8x16(x, y U8x16) U8x16
    func ShlU8x16(x, shift uint8) U8x16
    func ShrU8x16(x, shift uint8) U8x16


    func AddU16x8(x, y U16x8) U16x8
    func SubU16x8(x, y U16x8) U16x8
    func MulU16x8(x, y U16x8) U16x8
    func DivU16x8(x, y U16x8) U16x8
    func ShlU16x8(x, shift uint8) U16x8
    func ShrU16x8(x, shift uint8) U16x8


    func AddU32x4(x, y U32x4) U32x4
    func SubU32x4(x, y U32x4) U32x4
    func MulU32x4(x, y U32x4) U32x4
    func DivU32x4(x, y U32x4) U32x4
    func ShlU32x4(x, shift uint8) U32x4
    func ShrU32x4(x, shift uint8) U32x4


    func AddU64x2(x, y U64x2) U64x2
    func SubU64x2(x, y U64x2) U64x2
    func MulU64x2(x, y U64x2) U64x2
    func DivU64x2(x, y U64x2) U64x2
    func ShlU64x2(x, shift uint8) U64x2
    func ShrU64x2(x, shift uint8) U64x2


    func AddF32x4(x, y F32x4) F32x4
    func SubF32x4(x, y F32x4) F32x4
    func MulF32x4(x, y F32x4) F32x4
    func DivF32x4(x, y F32x4) F32x4


    func AddF64x2(x, y F64x2) F64x2
    func SubF64x2(x, y F64x2) F64x2
    func MulF64x2(x, y F64x2) F64x2
    func DivF64x2(x, y F64x2) F64x2

### Integer Overflow
For unsigned integer values, the simd functions `Add*`, `Sub*`, `Mul*`, and `Shl*` are computed modulo 2^n, where n is the bit width of the unsigned integer's type. These unsigned integer operations discard high bits upon overflow.

For signed integers, the simd functions `Add*`, `Sub*`, `Mul*`, and `Shl*`  may overflow and the resulting value exists and is defined by the signed integer representation, the operation, and its operands. The behavior is guaranteed to be identical to the go versions of the functions in `gensimd/simd/simd.go`.

For both unsigned and signed integer values, the simd functions `Div*` and `Shr*` are guaranteed to have the same behavior as the go versions in `gensimd/simd/simd.go`

### Floating Point
For the floating point simd functions `Add*`, `Sub*`, `Mul*`, and `Div*` the behavior is guaranteed to be identical to the go versions of the functions in `gensimd/simd/simd.go`.
