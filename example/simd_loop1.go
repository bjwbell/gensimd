// +build !simd

//go:generate gensimd

package main

import "github.com/bjwbell/gensimd/simd"

func simd_loop1(v4 *[4]int) int {
	var tmp int
	var ret int
	var tmp2 simd.Int4
	tmp = 4
	v := simd.Int4Var(v4)
	for i := 0; i < 10; i++ {
		_ = simd.Int4Add{
			Result: v,
			A:      v,
			B:      v,
		}
	}
	if tmp == 4 {
		ret = 1
	}
	if v == nil {
		ret = 1
	}
	if tmp2[0] == 0 {
		ret = 0
	}
	return ret
}
