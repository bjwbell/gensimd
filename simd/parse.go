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
	"reflect"
	"strings"
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
	info     *types.Info
	pkg      *types.Package
}

func (f *File) ErrorLocation(err *Error) string {
	if err == nil {
		return ""
	}
	return f.fs.Position(err.Pos).String()
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

// Valid checks that the go constructs used in the file, f, are valid
// in the subset of go used in simd translation
//
// The go subset allows funcs that are passed parameters of basic types
// or slices and arrays of basic types and return int, bool, or nothing
// In a func only simd var declarations are allowed and they must be at
// the beginning of the func.
// The func must be of the form
//
// {simd var declarations}
// {var assignments}
// {for loop} (optional)
// {var assignments}
// {return}
func (f *File) Valid() *Error {
	if f == nil || f.ast == nil {
		return &Error{errors.New("File is nil or file.ast is nil"), 0}
	}
	for _, decl := range f.ast.Decls {
		if err := f.validTopLevelDecl(decl); err != nil {
			return err
		}
	}
	return nil
}

func (f *File) validTopLevelDecl(decl ast.Decl) *Error {
	if decl == nil {
		return &Error{errors.New("Top level decl is nil"), 0}
	}

	if genDecl, ok := decl.(*ast.GenDecl); ok {
		if genDecl.Tok != token.IMPORT {
			return &Error{errors.New(fmt.Sprintf("Top level decl is not import statement or function declaration, decl:%v", genDecl)), genDecl.Pos()}
		}

	} else if funcDecl, ok := decl.(*ast.FuncDecl); ok {
		// Only functions (not methods) are allowed
		if funcDecl.Recv != nil {
			return &Error{errors.New("Only functions (not methods) are allowed"), funcDecl.Recv.Pos()}
		}
		if err := f.validFuncType(funcDecl.Type); err != nil {

			return err
		}
		if err := f.validFuncBody(funcDecl.Body); err != nil {
			return err
		}
	} else {
		return &Error{errors.New(fmt.Sprintf("Top level decl is not import statement or function declaration, decl:%v", decl)), decl.Pos()}
	}
	return nil
}

func (f *File) validFuncBody(block *ast.BlockStmt) *Error {
	var declStmts []*ast.DeclStmt
	var initStmts []ast.Stmt
	var finishStmts []ast.Stmt
	declSection, initSection, forSection, finishSection, retSection := true, false, false, false, false
	// empty func bodies not allowed
	if block.List == nil {
		return &Error{errors.New("Empty func bodies not allowed"), block.Pos()}
	}
	for _, stmt := range block.List {
		if declSection {
			decl, ok := stmt.(*ast.DeclStmt)
			if ok {
				if err := f.validDeclStmt(decl); err != nil {
					return err
				}
				declStmts = append(declStmts, decl)
				continue
			} else {
				declSection = false
				initSection = true
			}
		}
		if initSection {
			if err := f.validStmt(stmt); err == nil {
				initStmts = append(initStmts, stmt)
				continue
			} else {
				initSection = false
				forSection = true
			}
		}
		if forSection {
			if stmt, ok := stmt.(*ast.ForStmt); ok {
				if !f.validForStmt(stmt) {
					return &Error{errors.New("Invalid for statement"), stmt.Pos()}

				}
			}
			forSection = false
			finishSection = true
		}
		if finishSection {
			if f.validFinishStmt(stmt) {
				finishStmts = append(finishStmts, stmt)
				continue
			} else {
				finishSection = false
				retSection = true
			}
		}
		if retSection {
			if ret, ok := stmt.(*ast.ReturnStmt); ok {
				if err := f.validRetStmt(ret); err != nil {
					return err
				}
			} else {
				return &Error{errors.New("Expected return statement in retSection"), stmt.Pos()}
			}
		}
		pos := f.fs.Position(stmt.Pos())
		end := f.fs.Position(stmt.End())
		panic(fmt.Sprintf("PANIC: NO SECTION, stmt: (begin, end) - (%v, %v)", pos, end))
	}
	return nil
}
func (f *File) validDeclStmt(stmt *ast.DeclStmt) *Error {
	decl := stmt.Decl
	d, ok := decl.(*ast.GenDecl)
	// only gendecls are valid
	if !ok {
		return &Error{errors.New(fmt.Sprintf("Invalid declstmt, only ast.GenDecl allowed, decl: %v", decl)), decl.Pos()}
	}
	// only var declations are valid
	if d.Tok != token.VAR {
		return &Error{errors.New(fmt.Sprintf("Invalid Gendecl, only var decls allowd, d.Tok:%v", d.Tok)), d.TokPos}
	}
	// only exactly 1 spec element is allowed
	if d.Specs == nil || len(d.Specs) != 1 {
		return &Error{errors.New("Invalid decl either specs == nil or len(specs) > 1"), d.Pos()}
	}
	spec := d.Specs[0]
	vspec, ok := spec.(*ast.ValueSpec)
	// only value specs allowed
	if !ok {
		return &Error{errors.New(fmt.Sprintf("Invalid decl only value specs are allowed, instead of: %v", spec)), spec.Pos()}
	}
	if vspec.Names == nil || len(vspec.Names) != 1 {
		return &Error{errors.New("Invalid decl either spec.Names is nil or len(spec.Names) > 1"), vspec.Pos()}
	}
	//name := vspec.Names[0]

	// cannot initialize with values
	if vspec.Values != nil && len(vspec.Values) > 0 {
		return &Error{errors.New("Invalid decl, initalizing with values is not allowed"), vspec.Pos()}
	}
	// only valid var types allowed
	if err := f.validVarDeclType(f.info.TypeOf(vspec.Type)); err != nil {
		err.Pos = vspec.Pos()
		if e2 := f.validParamType(f.info.TypeOf(vspec.Type)); e2 == nil {
			err.Err = errors.New(fmt.Sprint(err.Err) + ", type only valid as a func param type")
		}

		return err
	}
	return nil
}
func (f *File) validStmt(stmt ast.Stmt) *Error {
	if stmt == nil {
		return nil
	}
	if assign, ok := stmt.(*ast.AssignStmt); ok {
		if len(assign.Lhs) != 1 || len(assign.Rhs) != 1 {
			return &Error{errors.New("Invalid assigment statment"), assign.Pos()}
		}
		return nil
	}
	if ifstmt, ok := stmt.(*ast.IfStmt); ok {
		if err := f.validExpr(ifstmt.Cond); err != nil {
			return err
		}
		// initialization clause not allowed for if statements
		if ifstmt.Init != nil {
			return &Error{errors.New("ifstmt cannot have initialization clause"), ifstmt.Init.Pos()}
		}
		if err := f.validStmt(ifstmt.Body); err != nil {
			return err
		}
		if err := f.validStmt(ifstmt.Else); err != nil {
			return err
		}

	}
	if blk, ok := stmt.(*ast.BlockStmt); ok {
		for _, s := range blk.List {
			if err := f.validStmt(s); err != nil {
				return err
			}
		}
	}
	if _, ok := stmt.(*ast.EmptyStmt); ok {
		return nil
	}
	if ret, ok := stmt.(*ast.ReturnStmt); ok {
		if ret.Results == nil || len(ret.Results) == 0 {
			return nil
		}
		if len(ret.Results) > 1 {
			return &Error{errors.New("Return statement doesn't allow multiple return values"), ret.Pos()}
		}
		return f.validRetExpr(ret.Results[0])
	}
	if indec, ok := stmt.(*ast.IncDecStmt); ok {
		// TODO specialize
		if err := f.validExpr(indec.X); err != nil {
			return err
		}
	}
	return &Error{errors.New(fmt.Sprintf("Invalid stmt:%v", stmt)), stmt.Pos()}
}
func (f *File) validExpr(expr ast.Expr) *Error {
	if expr != nil {
		return nil
	}
	return &Error{errors.New("Nil ast.Expr not allowed"), expr.Pos()}
}
func (f *File) validForStmt(stmt *ast.ForStmt) bool {
	return true
}

func (f *File) validFinishStmt(stmt ast.Stmt) bool {
	return true
}

func (f *File) validRetStmt(ret *ast.ReturnStmt) *Error {
	// returning nothing is ok
	if ret.Results == nil || len(ret.Results) == 0 {
		return nil
	}
	// Cannot return multiple values
	if len(ret.Results) > 1 {
		return &Error{errors.New("Return statements can only have one result"), ret.Pos()}
	}
	expr := ret.Results[0]
	if err := f.validRetExpr(expr); err != nil {
		return err
	}
	return nil
}

func (f *File) validRetExpr(expr ast.Expr) *Error {
	_, ok := expr.(*ast.Ident)
	// can only return identifies, i.e. variables
	if !ok {
		return &Error{errors.New("Return expression only allows identifiers"), expr.Pos()}
	}
	return nil
}

func (f *File) validFuncType(typ *ast.FuncType) *Error {
	if typ == nil {
		return &Error{errors.New("Nil func type"), 0}
	}
	if e := f.validParams(typ.Params); e != nil {
		return e
	}
	if e := f.validResults(typ.Results); e != nil {
		return e
	}
	return nil
}

func (f *File) validResults(results *ast.FieldList) *Error {
	if results == nil || results.List == nil {
		return nil
	}
	if results.NumFields() != 1 {
		err := fmt.Sprint("ERROR: can only return at most one result, not:",
			results.NumFields())
		return &Error{errors.New(err), results.Pos()}
	}
	result := results.List[0]
	if result == nil {
		return nil
	}
	if result.Names != nil {

		return &Error{errors.New(fmt.Sprint("ERROR: can only return nonnamed result, not:", result.Names)), result.Pos()}
	}
	typ := f.info.TypeOf(result.Type)
	if err := f.validResultType(typ); err != nil {
		err.Pos = result.Pos()
		if f.validVarDeclType(typ) == nil {
			err.Err = errors.New(fmt.Sprint(err.Err) + ", type only valid as a var decl, or param")
		} else if f.validParamType(typ) == nil {
			err.Err = errors.New(fmt.Sprint(err.Err) + ", type only valid as func param type")
		}

		return err
	}
	return nil
}

func (f *File) validParams(params *ast.FieldList) *Error {
	if params == nil {
		panic("ERROR: params fieldlist should never be nil")
	}
	if params.List == nil {
		return nil
	}
	for i := 0; i < params.NumFields(); i++ {
		field := params.List[i]
		if field == nil {

			return &Error{errors.New(fmt.Sprint("ERROR nil field, anonymous fields not allowed!!")), params.Pos()}
		}
		if len(field.Names) != 1 {
			panic("ERROR len(field.Names) != 1!!")
		}
		name := field.Names[0]
		if name == nil {
			panic("ERROR name == nil, this shouldn't occur")
		}
		typ := f.info.TypeOf(field.Type)
		if e := f.validParamType(typ); e != nil {
			e.Pos = field.Pos()
			return e
		}

	}
	return nil
}

func (f *File) validParamType(typ types.Type) *Error {
	if e := f.validVarType(typ); e != nil {
		switch typ.(type) {
		default:
			return e
		case *types.Pointer:
			typ := typ.(*types.Pointer)
			return f.validVarType(typ.Elem())
		case *types.Slice:
			typ := typ.(*types.Slice)
			return f.validVarType(typ.Elem())
		}
	}
	return nil
}

func (f *File) validVarType(typ types.Type) *Error {
	if e := f.validVarDeclType(typ); e != nil {
		switch typ.(type) {
		default:
			return e
		case *types.Array:
			typ := typ.(*types.Array)
			return f.validVarType(typ.Elem())
		case *types.Named:
			named, ok := typ.(*types.Named)
			if !ok {
				panic("ERROR can't cast to named type")
			}
			tname := named.Obj()
			i := Int(0)
			ivar := &i
			simdIntVar := reflect.TypeOf(ivar)
			var i4var Int4Var
			simdInt4Var := reflect.TypeOf(i4var)
			switch tname.Name() {
			default:
				return e
			case simdIntVar.Name():
				return nil
			case simdInt4Var.Name():
				return nil
			}
		}
	}
	return nil
}
func (f *File) validVarDeclType(typ types.Type) *Error {
	if e := f.validResultType(typ); e != nil {
		switch typ.(type) {
		default:
			return e
		case *types.Interface:
			return f.validVarDeclType(typ.Underlying())
		case *types.Named:
			named, ok := typ.(*types.Named)
			if !ok {
				panic("ERROR can't cast to named type")
			}
			tname := named.Obj()
			i := Int(0)
			simdInt := reflect.TypeOf(i)
			var i4 Int4
			simdInt4 := reflect.TypeOf(i4)
			switch tname.Name() {
			default:
				return &Error{errors.New(fmt.Sprintf("invalid type (%v)", tname.Name())), 0}
			case simdInt.Name():
				return nil
			case simdInt4.Name():
				return nil
			}
		}
	}
	return nil
}

func (f *File) validResultType(typ types.Type) *Error {
	switch typ.(type) {
	default:
		name := typ.String()
		if named, ok := typ.(*types.Named); ok {
			name = named.Obj().Name()
		}
		return &Error{errors.New(fmt.Sprint("Invalid result type:", name)), 0}
	case *types.Basic:
		typ := typ.(*types.Basic)
		switch typ.Kind() {
		default:

			return &Error{errors.New(fmt.Sprint("Invalid basic type for result type :", typ.Info())), 0}
		case types.Bool:
			return nil
		case types.Int:
			return nil
		case types.Int8:
			return nil
		case types.Int16:
			return nil
		case types.Int32:
			return nil
		case types.Int64:
			return nil
		case types.Uint:
			return nil
		case types.Uint8:
			return nil
		case types.Uint16:
			return nil
		case types.Uint32:
			return nil
		case types.Uint64:
			return nil
		case types.Float32:
			return nil
		case types.Float64:
			return nil
		}
	}
}
