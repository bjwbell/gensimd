// +build simd_go

package main

import (
	"fmt"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd

func simd_loop1(v4 *[4]int) error {
	v := simd.Int4Var(v4)
	body := []simd.Instruction{simd.Int4Add{Result: v, A: v, B: v}}
	loop := simd.ForLoop{Iterations: 10, Body: body}
	fmt.Println("loop:", loop)
	f := simd.Func{Init: []simd.Instruction{}, Loop: loop, Finish: []simd.Instruction{}, Ret: &simd.RetSuccess{}}
	return f.Exec()
}
