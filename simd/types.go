package simd

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Execer interface {
	Exec() error
}

type Asmer interface {
	Asm() string
}

type Instructioner interface {
	TypeInstruction()
	Execer
	Asmer
}

type RetInst interface {
	TypeRetInst()
	Execer
	Asmer
}

type Block []Instructioner

type Looper interface {
	TypeLooper()
	Execer
	Asmer
}

type ForLoop struct {
	Looper
	Start      Int
	Iterations Int
	StepBy     Int
	Body       Block
}

type Func struct {
	Init   Block
	Loop   ForLoop
	Finish Block
	Ret    RetInst
}

type Int int
type IntVar *int
type Int4 [4]int
type Int4Var *[4]int

type Int4Add struct {
	Instructioner
	Result Int4Var
	A      Int4Var
	B      Int4Var
}

type RetSuccess struct {
	RetInst
}

type Error struct {
	Err error
	Pos token.Pos
}

// File holds a single parsed file and associated data.
type File struct {
	pathName string
	ast      *ast.File // Parsed AST.
	fs       *token.FileSet
	info     *types.Info
	pkg      *types.Package

	functionInfo map[string]Function
}

func (f *File) ErrorLocation(err *Error) string {
	if err == nil {
		return ""
	}
	return f.fs.Position(err.Pos).String()
}

type Function struct {
	file      *File
	name      string
	decl      *ast.FuncDecl
	vars      map[string]varinfo
	initBlock []ast.Stmt

	unusedReg []register
	usedReg   []register
}

type varinfo struct {
	name   string
	offset int
	size   int
}

type register struct {
	name string
	typ  registerType
	size int
}

type registerType int

const DataReg = 0
const AddrReg = 1
const FloatReg = 2
