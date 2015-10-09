// +build amd64,gc

package tests

import "testing"

func t2test0Simd(int) int
func t2test1Simd(int) int
func t2test2Simd(int) int
func t2test3Simd(int) int
func t2test4Simd(int) int

//go:generate gensimd -fn "t2test0" -outfn "t2test0Simd" -f "$GOFILE" -o "t2test0.s"
func t2test0(x int) int {
	return x
}

//go:generate gensimd -fn "t2test1" -outfn "t2test1Simd" -f "$GOFILE" -o "t2test1.s"
func t2test1(x int) int {
	return x + 1
}

//go:generate gensimd -fn "t2test2" -outfn "t2test2Simd" -f "$GOFILE" -o "t2test2.s"
func t2test2(x int) int {
	return x * 2
}

//go:generate gensimd -fn "t2test3" -outfn "t2test3Simd" -f "$GOFILE" -o "t2test3.s"
func t2test3(x int) int {
	return x / 3
}

//go:generate gensimd -fn "t2test4" -outfn "t2test4Simd" -f "$GOFILE" -o "t2test4.s"
func t2test4(x int) int {
	return x * x
}

func TestT2(t *testing.T) {
	if t2test0Simd(1) != t2test0(1) {
		t.Errorf("t2test0Simd (%v) != t2test0 (%v)", t2test0Simd(1), t2test0(1))
	}
	if t2test1Simd(2) != t2test1(2) {
		t.Errorf("t2test1Simd (%v) != t2test1 (%v)", t2test1Simd(2), t2test1(2))
	}
	if t2test2Simd(3) != t2test2(3) {
		t.Errorf("t2test2Simd (%v) != t2test2 (%v)", t2test2Simd(3), t2test2(3))
	}
	if t2test3Simd(7) != t2test3(7) {
		t.Errorf("t2test3Simd (%v) != t2test3 (%v)", t2test3Simd(7), t2test3(7))
	}
	if t2test4Simd(8) != t2test4(8) {
		t.Errorf("t2test4Simd (%v) != t2test4 (%v)", t2test4Simd(8), t2test4(8))
	}
}
