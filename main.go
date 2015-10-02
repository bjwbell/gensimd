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
	if err := f.Valid(); err != nil {

		fmt.Println(fmt.Sprintf("Invalid file, ERROR:%v, LOCATION:%v \n", err.Err, f.ErrorLocation(err)))
	} else {
		fmt.Println("Validated:", file)
	}
	// TODO generate assembly instructions
}
