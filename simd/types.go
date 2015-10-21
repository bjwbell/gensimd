package simd

import (
	"go/ast"
	"go/token"
	"go/types"
)

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
type I8x16 [16]int8
type I16x8 [8]int16
type I32x4 [4]int32
type I64x2 [2]int64
type I128 [16]byte

type U8x16 [16]uint8
type U16x8 [8]uint16
type U32x4 [4]uint32
type U64x2 [2]uint64
type U128 [16]byte

type F32x4 [4]float32
type F64x2 [2]float64

//type F128 [2]float64
