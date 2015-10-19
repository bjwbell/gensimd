// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "boolt0, boolt1, boolt2, boolt3, boolt4, boolt5" -outfn "boolt0s, boolt1s, boolt2s, boolt3s, boolt4s, boolt5s" -f "$GOFILE" -o "bool_test_amd64.s"

func boolt0s(bool) bool
func boolt1s(bool) bool
func boolt2s(bool, bool) bool
func boolt3s(bool, bool) bool
func boolt4s(bool, bool) bool
func boolt5s(bool, bool) bool

func boolt0(x bool) bool {
	return x
}
func boolt1(x bool) bool {
	return !x
}
func boolt2(x, y bool) bool {
	return x && y
}
func boolt3(x, y bool) bool {
	return x || y
}
func boolt4(x, y bool) bool {
	z := x || y
	return z
}
func boolt5(x, y bool) bool {
	z := x || y
	return z
}

func TestBool(t *testing.T) {

	count := 0

	y := false
	for i := 0; i <= 1; i++ {

		x := false
		for j := 0; j <= 1; j++ {

			count++

			if boolt0s(x) != boolt0(x) {
				t.Errorf("boolt0s (%v) != boolt0 (%v)", boolt0s(x), boolt0(x))
			}
			if boolt1s(x) != boolt1(x) {
				t.Errorf("boolt1s (%v) != boolt1 (%v)", boolt1s(x), boolt1(x))
			}
			if boolt2s(x, y) != boolt2(x, y) {
				t.Errorf("boolt2s (%v) != boolt2 (%v)", boolt2s(x, y), boolt2(x, y))
			}
			if boolt3s(x, y) != boolt3(x, y) {
				t.Errorf("boolt3s (%v) != boolt3 (%v)", boolt3s(x, y), boolt3(x, y))
			}
			if boolt4s(x, y) != boolt4(x, y) {
				t.Errorf("boolt4s (%v) != boolt4 (%v)", boolt4s(x, y), boolt4(x, y))
			}
			if boolt5s(x, y) != boolt5(x, y) {
				t.Errorf("boolt5s (%v) != boolt5 (%v)", boolt5s(x, y), boolt5(x, y))
			}
			x = !x
		}
		y = !y
	}

	t.Log("Test Count:", count)
}
