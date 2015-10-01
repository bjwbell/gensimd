package simd

import (
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"strings"
)

// File holds a single parsed file and associated data.
type File struct {
	pathName string
	ast      *ast.File // Parsed AST.
	fs       *token.FileSet
	info     *types.Info
	pkg      *types.Package
}

// check type-checks the package. The package must be OK to proceed.
func (f *File) check() {
	// TODO
	typs := make(map[ast.Expr]types.TypeAndValue)
	defs := make(map[*ast.Ident]types.Object)
	uses := make(map[*ast.Ident]types.Object)
	config := types.Config{FakeImportC: true}
	config.Importer = importer.Default()
	info := &types.Info{Types: typs, Defs: defs, Uses: uses}
	astFiles := []*ast.File{f.ast}
	typesPkg, err := config.Check(fileDir(f), f.fs, astFiles, info)
	if err != nil {
		log.Fatalf("checking package: %s", err)
	}
	f.info = info
	f.pkg = typesPkg

}
func fileDir(f *File) string {
	split := strings.Split(f.pathName, "/")
	dir := ""
	if len(split) == 1 {
		dir = "."
	} else if len(split) == 2 {
		dir = split[0]
	} else {
		dir = strings.Join(split[0:len(split)-2], "/") // + "/"
	}
	fmt.Println("dir:", dir)
	return dir
}

// ParseFile analyzes the single file passed in.
func ParseFile(f string) (*File, error) {
	fs := token.NewFileSet()
	if !strings.HasSuffix(f, ".go") {
		return nil, errors.New("Invalid file, file suffix not .go")
	}

	parsed, err := parser.ParseFile(fs, f, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parsing file: %s: %s", f, err)
		return nil, errors.New(fmt.Sprintf("Error parsing file:%v", f))
	}
	file := File{
		ast:      parsed,
		pathName: f,
		fs:       fs,
	}
	file.check()
	return &file, nil
}
