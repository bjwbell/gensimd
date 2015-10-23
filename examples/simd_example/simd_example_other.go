// +build !amd64,gc

package main

import "github.com/bjwbell/gensimd/simd"

func addi32x4(x, y simd.I32x4) simd.I32x4           { return simd.AddI32x4(x, y) }
func subi32x4(x, y simd.I32x4) simd.I32x4           { return simd.SubI32x4(x, y) }
func muli32x4(x, y simd.I32x4) simd.I32x4           { return simd.MulI32x4(x, y) }
func shli32x4(x simd.I32x4, shift uint8) simd.I32x4 { return simd.ShlI32x4(x, shift) }
func shri32x4(x simd.I32x4, shift uint8) simd.I32x4 { return simd.ShrI32x4(x, shift) }

func addf32x4(x, y simd.F32x4) simd.F32x4 { return simd.AddF32x4(x, y) }
func subf32x4(x, y simd.F32x4) simd.F32x4 { return simd.SubF32x4(x, y) }
func mulf32x4(x, y simd.F32x4) simd.F32x4 { return simd.MulF32x4(x, y) }
func divf32x4(x, y simd.F32x4) simd.F32x4 { return simd.DivF32x4(x, y) }
