// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "uint8_t0, uint8_t1, uint8_t2, uint8_t3, uint8_t4" -outfn "uint8_t0_simd, uint8_t1_simd, uint8_t2_simd, uint8_t3_simd, uint8_t4_simd" -f "$GOFILE" -o "uint8_test.s"

func uint8_t0_simd(uint8) uint8
func uint8_t1_simd(uint8) uint8
func uint8_t2_simd(uint8) uint8
func uint8_t3_simd(uint8) uint8
func uint8_t4_simd(uint8) uint8

func uint8_t0(x uint8) uint8 {
	return x
}

func uint8_t1(x uint8) uint8 {
	return x + 1
}

func uint8_t2(x uint8) uint8 {
	return x * 2
}

func uint8_t3(x uint8) uint8 {
	return x / 3
}

func uint8_t4(x uint8) uint8 {
	return x * x
}

func TestUint8(t *testing.T) {

	if uint8_t0_simd(1) != uint8_t0(1) {
		t.Errorf("uint8_t0_simd (%v) != uint_t0 (%v)", uint8_t0_simd(1), uint8_t0(1))
	}

	if uint8_t1_simd(2) != uint8_t1(2) {
		t.Errorf("uint8_t1_simd (%v) != uint8_t1 (%v)", uint8_t1_simd(2), uint8_t1(2))
	}

	if uint8_t2_simd(3) != uint8_t2(3) {
		t.Errorf("uint8_t2_simd (%v) != uint8_t2 (%v)", uint8_t2_simd(3), uint8_t2(3))
	}

	if uint8_t3_simd(7) != uint8_t3(7) {
		t.Errorf("uint8_t3_simd (%v) != uint8_t3 (%v)", uint8_t3_simd(7), uint8_t3(7))
	}

	if uint8_t4_simd(8) != uint8_t4(8) {
		t.Errorf("uint8_t4_simd (%v) != uint8_t4 (%v)", uint8_t4_simd(8), uint8_t4(8))
	}

}
