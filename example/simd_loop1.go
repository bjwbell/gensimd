// +build !simd

//go:generate gensimd

package main

import (
	"fmt"

	"github.com/bjwbell/gensimd/simd"
)

func simd_loop1(v4 *[4]int) error {
	v := simd.Int4Var(v4)
	loop := simd.ForLoop{
		Start: 0, Iterations: 10, StepBy: 0,
		Body: []simd.Instructioner{
			simd.Int4Add{Result: v, A: v, B: v},
		},
	}
	f := simd.Func{
		Init:   []simd.Instructioner{},
		Loop:   loop,
		Finish: []simd.Instructioner{},
		Ret:    &simd.RetSuccess{}}
	fmt.Println("loop:", loop)
	return f.Exec()
}
