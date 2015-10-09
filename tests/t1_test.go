// +build amd64,gc

package tests

import (
	"math"
	"testing"
)

func test0Simd() int
func test1Simd() int
func test2Simd() int
func test3Simd() int
func test4Simd() int

//go:generate gensimd -fname "test0" -outfname "test0Simd" -file "$GOFILE" -o "test0.s"
func test0() int {
	return 0
}

//go:generate gensimd -fname "test1" -outfname "test1Simd" -file "$GOFILE" -o "test1.s"
func test1() int {
	return 1
}

//go:generate gensimd -fname "test2" -outfname "test2Simd" -file "$GOFILE" -o "test2.s"
func test2() int {
	return 2
}

//go:generate gensimd -fname "test3" -outfname "test3Simd" -file "$GOFILE" -o "test3.s"
func test3() int {
	return 256
}

//go:generate gensimd -fname "test4" -outfname "test4Simd" -file "$GOFILE" -o "test4.s"
func test4() int {
	return math.MaxInt64
}

func TestT1(t *testing.T) {
	if test0Simd() != test0() {
		t.Errorf("test0Simd (%v) != test0 (%v)", test0Simd(), test0())
	}
	if test1Simd() != test1() {
		t.Errorf("test1Simd (%v) != test1 (%v)", test1Simd(), test1())
	}
	if test2Simd() != test2() {
		t.Errorf("test2Simd (%v) != test2 (%v)", test2Simd(), test2())
	}
	if test3Simd() != test3() {
		t.Errorf("test3Simd (%v) != test3 (%v)", test3Simd(), test3())
	}
	if test4Simd() != test4() {
		t.Errorf("test4Simd (%v) != test4 (%v)", test4Simd(), test4())
	}
}
