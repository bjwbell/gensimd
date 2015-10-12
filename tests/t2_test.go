// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "t2test0, t2test1, t2test2, t2test3, t2test4" -outfn "t2t0simd, t2t1simd, t2t2simd, t2t3simd, t2t4simd" -f "$GOFILE" -o "t2.s"

func t2t0simd(int) int
func t2t1simd(int) int
func t2t2simd(int) int
func t2t3simd(int) int
func t2t4simd(int) int

func t2test0(x int) int {
	return x
}

func t2test1(x int) int {
	return x + -1
}

func t2test2(x int) int {
	return x * -2
}

func t2test3(x int) int {
	return x / 3
}

func t2test4(x int) int {
	return x * -x
}

func TestT2(t *testing.T) {
	if t2t0simd(1) != t2test0(1) {
		t.Errorf("t2test0Simd (%v) != t2test0 (%v)", t2t0simd(1), t2test0(1))
	}
	if t2t1simd(2) != t2test1(2) {
		t.Errorf("t2test1Simd (%v) != t2test1 (%v)", t2t1simd(2), t2test1(2))
	}
	if t2t2simd(3) != t2test2(3) {
		t.Errorf("t2test2Simd (%v) != t2test2 (%v)", t2t2simd(3), t2test2(3))
	}
	if t2t3simd(7) != t2test3(7) {
		t.Errorf("t2test3Simd (%v) != t2test3 (%v)", t2t3simd(7), t2test3(7))
	}
	if t2t4simd(8) != t2test4(8) {
		t.Errorf("t2test4Simd (%v) != t2test4 (%v)", t2t4simd(8), t2test4(8))
	}
}
