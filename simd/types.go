package simd

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Int int
type IntVar *int
type Int4 [4]int
type Int4Var *[4]int

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
