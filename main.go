package main

import (
	"fmt"
	"os"

	"github.com/bjwbell/gensimd/simd"
)

func main() {
	file := os.ExpandEnv("$GOFILE")
	f, err := simd.ParseFile(file)
	if err == nil {
		fmt.Println("Parsed: ", file)
	} else {
		fmt.Println(fmt.Sprintf("Error parsing %v, error: %v", file, err))
	}
	if err := f.Valid(); err {
		fmt.Println("File not valid, error:\n", err)
	}
	// TODO check file parse tree & generate assembly instructions
}
