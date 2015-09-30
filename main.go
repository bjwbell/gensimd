package main

import (
	"fmt"

	"github.com/bjwbell/gensimd/simd"
)

func main() {
	v4 := simd.Int4Var(&[4]int{})
	body := []simd.Instruction{simd.Int4Add{Result: v4, A: v4, B: v4}}
	loop := simd.ForLoop{Iterations: 10, Body: body}
	fmt.Println("loop:", loop)
}
