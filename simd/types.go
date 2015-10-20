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
