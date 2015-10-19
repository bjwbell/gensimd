// +build amd64,gc

package tests

import (
	"math"
	"testing"
)

//go:generate gensimd -fn "test0, test1, test2, test3, test4" -outfn "t0simd, t1simd,t2simd,t3simd,t4simd" -f "$GOFILE" -o "basic_test_amd64.s"

func t0simd() int
func t1simd() int
func t2simd() int
func t3simd() int
func t4simd() int

func test0() int {
	return 0
}

func test1() int {
	return 1
}

func test2() int {
	return 2
}

func test3() int {
	return 256
}

func test4() int {
	return math.MaxInt64
}

func TestBasic(t *testing.T) {

	const count = 5

	if t0simd() != test0() {
		t.Errorf("t0simd (%v) != test0 (%v)", t0simd(), test0())
	}
	if t1simd() != test1() {
		t.Errorf("t1simd (%v) != test1 (%v)", t1simd(), test1())
	}
	if t2simd() != test2() {
		t.Errorf("t2simd (%v) != test2 (%v)", t2simd(), test2())
	}
	if t3simd() != test3() {
		t.Errorf("t3simd (%v) != test3 (%v)", t3simd(), test3())
	}
	if t4simd() != test4() {
		t.Errorf("t4simd (%v) != test4 (%v)", t4simd(), test4())
	}

	t.Log("Test Count:", count)
}
