// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "lent0, lent1, lent2" -outfn "lent0s, lent1s, lent2s" -f "$GOFILE" -o "builtin_test_amd64.s"

func lent0s(x [1]int) int
func lent1s(x [2]int) int
func lent2s(x []int) int

func lent0(x [1]int) int {
	return len(x)
}

func lent1(x [2]int) int {
	return len(x)
}

func lent2(x []int) int {
	return len(x)
}

func TestBuiltinLen(t *testing.T) {

	count := 0

	for i := -63; i <= 63; i++ {
		count++
		y := int(0)
		x := int(0)
		if i < 0 {
			y = -1 << uint(-i)
		} else {
			y = 1<<uint(i) - 1
		}
		if i%16 < 8 {
			x = -1 << uint(8-i%16)
		} else {
			x = 1<<uint(i%16-8) - 1
		}
		x1 := [1]int{x}
		x2 := [2]int{x, y}
		x3 := [3]int{2 * x, 3 * y}
		var slice []int
		slice = x3[0:3]

		if lent0s(x1) != lent0(x1) {
			t.Errorf("lent0s(%v)", x1)
		}
		if lent1s(x2) != lent1(x2) {
			t.Errorf("lent1s(%v)", x2)
		}
		if lent2s(slice) != lent2(slice) {
			t.Errorf("lent2s(%v)", slice)
		}
	}

	t.Log("Test Count:", count)
}
