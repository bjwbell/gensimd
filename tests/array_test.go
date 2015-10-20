// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "arrayt0, arrayt1, arrayt2" -outfn "arrayt0s, arrayt1s, arrayt2s" -f "$GOFILE" -o "array_test_amd64.s"

func arrayt0s(x [1]int) int
func arrayt1s(x [2]int) int
func arrayt2s(x [3]int) int

func arrayt0(x [1]int) int {
	return x[0]
}

func arrayt1(x [2]int) int {
	return x[1]
}

func arrayt2(x [3]int) int {
	return x[0] + x[1] + x[2]
}

func TestArray(t *testing.T) {

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

			x1 := [1]int{x}
			x2 := [2]int{x, y}
			x3 := [3]int{2 * x, 3 * y}

			if arrayt0s(x1) != arrayt0(x1) {
				t.Errorf("arrayt0s(%v)", x1)
			}
			if arrayt1s(x2) != arrayt1(x2) {
				t.Errorf("arrayt1s(%v)", x2)
			}
			if arrayt2s(x3) != arrayt2(x3) {
				t.Errorf("arrayt2s(%v)", x3)
			}
		}
	}

	t.Log("Test Count:", count)
}
