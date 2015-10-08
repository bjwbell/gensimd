package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bjwbell/gensimd/codegen"
	"github.com/bjwbell/gensimd/simd"

	"go/build"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"golang.org/x/tools/go/types"
)

func filePath(pathName string) string {
	split := strings.Split(pathName, "/")
	dir := ""
	if len(split) == 1 {
		dir = "."
	} else if len(split) == 2 {
		dir = split[0] + "/"
	} else {
		dir = strings.Join(split[0:len(split)-2], "/") // + "/"
	}
	return dir
}

func fileName(pathName string) string {
	split := strings.Split(pathName, "/")
	name := ""
	if len(split) == 0 {
		name = ""
	} else {
		name = split[len(split)-1]
	}
	return name
}

func main() {
	var ssaDump = flag.Bool("ssa", false, "dump ssa representation")
	flag.Parse()
	args := flag.Args()
	file := os.ExpandEnv("$GOFILE")
	funcName := ""
	if len(args) == 1 {
		funcName = args[0]

	} else if len(args) == 2 {
		file = args[0]
		funcName = args[1]
	}
	f, err := simd.ParseFile(file)
	if err == nil {
		fmt.Println("Parsed: ", file)
	} else {
		fmt.Println(fmt.Sprintf("Error parsing file \"%v\", error: %v", file, err))
		return
	}

	filePkgName := f.Pkg.Name()
	filePkgPath := f.Pkg.Path()
	conf := loader.Config{Build: &build.Default}

	// Choose types.Sizes from conf.Build.
	var wordSize int64 = 8
	switch conf.Build.GOARCH {
	case "386", "arm":
		panic("SIMD invalid for x86 and arm")
	}
	conf.TypeChecker.Sizes = &types.StdSizes{
		MaxAlign: 8,
		WordSize: wordSize,
	}

	// Use the initial file from the command line/$GOFILE.
	conf.CreateFromFilenames(filePath(file), file)

	// Load, parse and type-check
	iprog, err := conf.Load()
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// Create and build SSA-form program representation.
	builderMode := ssa.SanityCheckFunctions
	if *ssaDump {
		builderMode = ssa.PrintFunctions
	}
	prog := ssautil.CreateProgram(iprog, builderMode)
	if prog == nil {
		fmt.Println("prog == nil")
	}
	// Build and display only the initial packages (and synthetic wrappers)
	for _, info := range iprog.InitialPackages() {
		prog.Package(info.Pkg).Build()
	}

	// TODO generate assembly instructions

	for _, pkg := range prog.AllPackages() {
		if pkg.Pkg.Path() == filePkgPath+"/" && pkg.Pkg.Name() == filePkgName {
			fmt.Println("found filePkgName:", filePkgName)
			if fn := pkg.Func(funcName); fn == nil {
				fmt.Println(fmt.Printf("Error function \"%v\" not found in \"%v\" package", funcName, filePkgName))
				return
			} else {
				fmt.Println("found function:", funcName)
				codegenFn, err := codegen.CreateFunction("codegen/instructionsetxml/Opcodes/opcodes/x86_64.xml", fn)
				if err != nil {
					fmt.Printf("Error in codegen.CreateFunction,  msg:\"%v\"", err)
				}
				if asm, err := codegenFn.GoAssembly(); err != nil {
					fmt.Println(asm)
					fmt.Printf("Error creating fn asm, msg:\"%v\"\n", err)
				} else {
					fmt.Println("fn asm:")
					fmt.Println(asm)
				}
			}
		}
	}
}
