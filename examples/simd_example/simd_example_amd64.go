// +build amd64,gc

package main

import "github.com/bjwbell/gensimd/simd"

//go:generate gensimd -fn "addi32x4, subi32x4, muli32x4, shli32x4, shri32x4, addf32x4, subf32x4, mulf32x4, divf32x4" -outfn "addi32x4, subi32x4, muli32x4, shli32x4, shri32x4, addf32x4, subf32x4, mulf32x4, divf32x4" -f "simd_example_other.go" -o "simd_example_amd64.s"

func addi32x4(x, y simd.I32x4) simd.I32x4
func subi32x4(x, y simd.I32x4) simd.I32x4
func muli32x4(x, y simd.I32x4) simd.I32x4
func shli32x4(x simd.I32x4, shift uint8) simd.I32x4
func shri32x4(x simd.I32x4, shift uint8) simd.I32x4

func addf32x4(x, y simd.F32x4) simd.F32x4
func subf32x4(x, y simd.F32x4) simd.F32x4
func mulf32x4(x, y simd.F32x4) simd.F32x4
func divf32x4(x, y simd.F32x4) simd.F32x4
