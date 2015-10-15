package simd

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Int int
type Int4 [4]int

type Error struct {
	Err error
	Pos token.Pos
}

// File holds a single parsed file and associated data.
type File struct {
	pathName string
	ast      *ast.File // Parsed AST.
	fs       *token.FileSet
	Info     *types.Info
	Pkg      *types.Package
}

func (f *File) ErrorLocation(err *Error) string {
	if err == nil {
		return ""
	}
	return f.fs.Position(err.Pos).String()
}

// SIMD types
type i8x16 [16]int8
type i16x8 [8]int16
type i32x4 [4]int32
type i64x2 [2]int64

type u8x16 [16]uint8
type u16x8 [8]uint16
type u32x4 [4]uint32
type u64x2 [2]uint64

type f32x4 [4]float32
type f64x2 [2]float64

// SIMD functions:
// add, sub, mul, div, <<, >> for each type
// Includes dummy function bodies
func AddI8x16(x, y i8x16) i8x16     { return i8x16{} }
func SubI8x16(x, y i8x16) i8x16     { return i8x16{} }
func MulI8x16(x, y i8x16) i8x16     { return i8x16{} }
func DivI8x16(x, y i8x16) i8x16     { return i8x16{} }
func ShlI8x16(x, shift uint8) i8x16 { return i8x16{} }
func ShrI8x16(x, shift uint8) i8x16 { return i8x16{} }

func AddI16x8(x, y i16x8) i16x8     { return i16x8{} }
func SubI16x8(x, y i16x8) i16x8     { return i16x8{} }
func MulI16x8(x, y i16x8) i16x8     { return i16x8{} }
func DivI16x8(x, y i16x8) i16x8     { return i16x8{} }
func ShlI16x8(x, shift uint8) i16x8 { return i16x8{} }
func ShrI16x8(x, shift uint8) i16x8 { return i16x8{} }

func AddI32x4(x, y i32x4) i32x4     { return i32x4{} }
func SubI32x4(x, y i32x4) i32x4     { return i32x4{} }
func MulI32x4(x, y i32x4) i32x4     { return i32x4{} }
func DivI32x4(x, y i32x4) i32x4     { return i32x4{} }
func ShlI32x4(x, shift uint8) i32x4 { return i32x4{} }
func ShrI32x4(x, shift uint8) i32x4 { return i32x4{} }

func AddI64x2(x, y i64x2) i64x2     { return i64x2{} }
func SubI64x2(x, y i64x2) i64x2     { return i64x2{} }
func MulI64x2(x, y i64x2) i64x2     { return i64x2{} }
func DivI64x2(x, y i64x2) i64x2     { return i64x2{} }
func ShlI64x2(x, shift uint8) i64x2 { return i64x2{} }
func ShrI64x2(x, shift uint8) i64x2 { return i64x2{} }

func AddU8x16(x, y u8x16) u8x16     { return u8x16{} }
func SubU8x16(x, y u8x16) u8x16     { return u8x16{} }
func MulU8x16(x, y u8x16) u8x16     { return u8x16{} }
func DivU8x16(x, y u8x16) u8x16     { return u8x16{} }
func ShlU8x16(x, shift uint8) u8x16 { return u8x16{} }
func ShrU8x16(x, shift uint8) u8x16 { return u8x16{} }

func AddU16x8(x, y u16x8) u16x8     { return u16x8{} }
func SubU16x8(x, y u16x8) u16x8     { return u16x8{} }
func MulU16x8(x, y u16x8) u16x8     { return u16x8{} }
func DivU16x8(x, y u16x8) u16x8     { return u16x8{} }
func ShlU16x8(x, shift uint8) u16x8 { return u16x8{} }
func ShrU16x8(x, shift uint8) u16x8 { return u16x8{} }

func AddU32x4(x, y u32x4) u32x4     { return u32x4{} }
func SubU32x4(x, y u32x4) u32x4     { return u32x4{} }
func MulU32x4(x, y u32x4) u32x4     { return u32x4{} }
func DivU32x4(x, y u32x4) u32x4     { return u32x4{} }
func ShlU32x4(x, shift uint8) u32x4 { return u32x4{} }
func ShrU32x4(x, shift uint8) u32x4 { return u32x4{} }

func AddU64x2(x, y u64x2) u64x2     { return u64x2{} }
func SubU64x2(x, y u64x2) u64x2     { return u64x2{} }
func MulU64x2(x, y u64x2) u64x2     { return u64x2{} }
func DivU64x2(x, y u64x2) u64x2     { return u64x2{} }
func ShlU64x2(x, shift uint8) u64x2 { return u64x2{} }
func ShrU64x2(x, shift uint8) u64x2 { return u64x2{} }

func AddF32x4(x, y f32x4) f32x4 { return f32x4{} }
func SubF32x4(x, y f32x4) f32x4 { return f32x4{} }
func MulF32x4(x, y f32x4) f32x4 { return f32x4{} }
func DivF32x4(x, y f32x4) f32x4 { return f32x4{} }

func AddF64x2(x, y f64x2) f64x2 { return f64x2{} }
func SubF64x2(x, y f64x2) f64x2 { return f64x2{} }
func MulF64x2(x, y f64x2) f64x2 { return f64x2{} }
func DivF64x2(x, y f64x2) f64x2 { return f64x2{} }
