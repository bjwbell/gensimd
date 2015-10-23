package main

//go:generate stringer -type=Instr,InstrOpType,InstructionType,XmmData codegen

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
	var debug = flag.Bool("debug", false, "add debug comments to assembly")
	var output = flag.String("o", "", "write Go Assembly to file")
	var f = flag.String("f", "", "input file")
	var flagFn = flag.String("fn", "", "comma separated list of function names")
	var flagOutFn = flag.String("outfn", "", "comma separated list of output function names")
	var goprotofile = flag.String("goprotofile", "", "output file for function prototype")

	flag.Parse()
	file := os.ExpandEnv("$GOFILE")
	log.SetFlags(log.Lshortfile)
	if *f != "" {
		file = *f
	}
	if *flagFn == "" {
		log.Fatalf("Error no function name(s) provided")
	}
	fnnames := strings.Split(*flagFn, ",")
	outFns := []string{}
	if *flagOutFn == "" {
		for _, fn := range fnnames {
			outFns = append(outFns, "gensimd_"+fn)
		}
	} else {
		outFns = strings.Split(*flagOutFn, ",")

	}
	if len(fnnames) != len(outFns) {
		log.Fatalf("Error # fns (%v) doesn't match # outfns (%v)\n", len(fnnames), len(outFns))
	}
	for i := range fnnames {
		fnnames[i] = strings.TrimSpace(fnnames[i])
		outFns[i] = strings.TrimSpace(outFns[i])
	}

	parsed, err := simd.ParseFile(file)
	if err != nil {
		msg := "Error parsing file \"%v\", error msg \"%v\"\n"
		log.Fatalf(msg, file, err)
	}

	filePkgName := parsed.Pkg.Name()
	filePkgPath := parsed.Pkg.Path()
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

	assembly := codegen.AssemblyFilePreamble()
	goprotos := ""
	protoPkgName := ""
	foundpkg := false
	for _, pkg := range prog.AllPackages() {
		if pkg.Pkg.Path() == filePkgPath && pkg.Pkg.Name() == filePkgName {
			foundpkg = true
			for i := range fnnames {
				fnname := fnnames[i]
				outfn := outFns[i]
				if fn := pkg.Func(fnname); fn == nil {
					msg := "Func \"%v\" not found in package \"%v\""
					log.Fatalf(msg, fnname, filePkgName)
				} else {
					fn, err := codegen.CreateFunction(fn, outfn, *debug)
					if err != nil {
						msg := "codegen error msg \"%v\""
						log.Fatalf(msg, err)
					}
					if asm, err := fn.GoAssembly(); err != nil {
						msg := "Error creating fn asm: \"%v\"\n"
						log.Fatalf(msg, err)
					} else {
						if *output == "" {
							fmt.Println(asm)
						} else {
							if *goprotofile != "" {
								pkg, proto := fn.GoProto()
								goprotos += proto + "\n"
								if protoPkgName == "" {
									protoPkgName = pkg + "\n"
								}
							}
							assembly += asm
						}
					}
				}
			}
		}
	}

	if !foundpkg {
		msg := "Error didn't find package, \"%v\"\n"
		panic(fmt.Sprintf(msg, filePkgName))
	}

	writeFile(*output, assembly)
	if *goprotofile != "" {
		writeFile(*goprotofile, protoPkgName+"\n"+goprotos)
	}
}

func writeFile(filename, contents string) {
	if err := ioutil.WriteFile(filename, []byte(contents), 0644); err != nil {
		log.Fatalf("Cannot write to file \"%v\", error \"%v\"\n", filename, err)
	}
}
