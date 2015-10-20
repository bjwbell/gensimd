// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "slicet0, slicet1, slicet2" -outfn "slicet0s, slicet1s, slicet2s" -f "$GOFILE" -o "slice_test_amd64.s"

func slicet0s(x []int) int
func slicet1s(x []int) int
func slicet2s(x []int) int

func slicet0(x []int) int {
	return x[0]
}

func slicet1(x []int) int {
	return x[1]
}

func slicet2(x []int) int {
	return x[0] + x[1] + x[2]
}

func TestSlice(t *testing.T) {

	count := 0

	for i := -63; i <= 63; i++ {

		y := int(0)
		if i < 0 {
			y = -1 << uint(-i)
		} else {
			y = 1<<uint(i) - 1
		}

		for j := -63; j <= 63; j++ {

			count++

			x := int(0)
			if j < 0 {
				x = -1 << uint(-j)
			} else {
				x = 1<<uint(j) - 1
			}

			x1 := []int{x}
			x2 := []int{x, y}
			x3 := []int{2 * x, 3 * y, x - y}

			if slicet0s(x1) != slicet0(x1) {
				t.Errorf("slicet0s(%v)", x1)
			}
			if slicet1s(x2) != slicet1(x2) {
				t.Errorf("slicet1s(%v)", x2)
			}
			if slicet2s(x3) != slicet2(x3) {
				t.Errorf("slicet2s(%v)", x3)
			}
		}
	}

	t.Log("Test Count:", count)
}
