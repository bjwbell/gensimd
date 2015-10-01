// +build !simd

//go:generate gensimd

package main

import "github.com/bjwbell/gensimd/simd"

func simd_loop1(v4 *[4]int) bool {
	var tmp int
	var ret bool
	tmp = 4
	v := simd.Int4Var(v4)
	for i := 0; i < 10; i++ {
		simd.Int4Add{
			Result: v,
			A:      v,
			B:      v,
		}
	}
	if tmp == 4 {
		ret = true
	}
	return ret
}
