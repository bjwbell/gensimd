package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
		dir = strings.Join(split[0:len(split)-2], "/")
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
	var outputFile = flag.String("o", "", "write Go Assembly to file")
	var inputFile = flag.String("file", "", "input file")
	var fnname = flag.String("fname", "", "function name")
	var outfn = flag.String("outfname", "", "output function name")
	var goprotofile = flag.String("goprotofile", "", "output file for function prototype")

	flag.Parse()
	file := os.ExpandEnv("$GOFILE")

	if *inputFile != "" {
		file = *inputFile
	}
	if *outfn == "" {
		*outfn = "gensimd_" + *fnname
	}

	f, err := simd.ParseFile(file)
	if err != nil {
		msg := "Error parsing file \"%v\", error msg \"%v\""
		log.Fatalf(msg, file, err)
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
		log.Fatalf("conf.Load, error msg \"%v\"", err)
	}

	// Create and build SSA-form program representation.
	builderMode := ssa.SanityCheckFunctions
	if *ssaDump {
		builderMode = ssa.PrintFunctions
	}
	prog := ssautil.CreateProgram(iprog, builderMode)
	if prog == nil {
		log.Fatalf("Couldn't create ssa representation")
	}

	// Build and display only the initial packages (and synthetic wrappers)
	for _, info := range iprog.InitialPackages() {
		prog.Package(info.Pkg).Build()
	}

	foundpkg := false
	foundfn := false
	for _, pkg := range prog.AllPackages() {
		if pkg.Pkg.Path() == filePkgPath && pkg.Pkg.Name() == filePkgName {
			foundpkg = true
			if fn := pkg.Func(*fnname); fn == nil {
				msg := "Function \"%v\" not found in package \"%v\""
				log.Fatalf(msg, *fnname, filePkgName)
			} else {
				foundfn = true
				fn, err := codegen.CreateFunction(fn, *outfn)
				if err != nil {
					msg := "codegen.CreateFunction,  error msg \"%v\""
					log.Fatalf(msg, err)
				}
				if asm, err := fn.GoAssembly(); err != nil {
					msg := "Error creating fn asm, error msg \"%v\"\n"
					log.Fatalf(msg, err)
				} else {
					if *outputFile == "" {
						fmt.Println(asm)
					} else {
						if *goprotofile != "" {
							writeFile(*goprotofile, fn.GoProto())
						}
						writeFile(*outputFile, asm)
					}
				}
			}
		}
	}
	if !foundpkg {
		msg := "Error in gensimd: didn't find package, \"%v\", for function \"%v\""
		log.Fatalf(msg, filePkgName, *fnname)

	} else if foundpkg && !foundfn {
		msg := "Error in gensimd: didn't find function, \"%v\", in package \"%v\""
		log.Fatalf(msg, *fnname, filePkgName)
	}
}

func writeFile(filename, contents string) {
	if err := ioutil.WriteFile(filename, []byte(contents), 0644); err != nil {
		log.Fatalf("Error writing to file, error msg \"%v\"\n", err)
	}
}
