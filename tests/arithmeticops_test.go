// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "add, sub, neg, mul, div, addint8, subint8, negint8, mulint8, divint8, addint16, subint16, negint16, mulint16, divint16, addint64, subint64, negint64, mulint64, divint64, adduint8, subuint8, muluint8, divuint8, adduint16, subuint16, muluint16, divuint16, adduint32, subuint32, muluint32, divuint32, adduint64, subuint64, muluint64, divuint64" -outfn "adds, subs, negs, muls, divs, addint8s, subint8s, negint8s, mulint8s, divint8s, addint16s, subint16s, negint16s, mulint16s, divint16s, addint64s, subint64s, negint64s, mulint64s, divint64s, adduint8s, subuint8s, muluint8s, divuint8s, adduint16s, subuint16s, muluint16s, divuint16s, adduint32s, subuint32s, muluint32s, divuint32s, adduint64s, subuint64s, muluint64s, divuint64s" -f "$GOFILE" -o "arithmeticops_test_amd64.s"

func adds(x, y int32) int32
func subs(x, y int32) int32
func negs(x int32) int32
func muls(x, y int32) int32
func divs(x, y int32) int32
func addint8s(x, y int8) int8
func subint8s(x, y int8) int8
func negint8s(x int8) int8
func mulint8s(x, y int8) int8
func divint8s(x, y int8) int8
func addint16s(x, y int16) int16
func subint16s(x, y int16) int16
func negint16s(x int16) int16
func mulint16s(x, y int16) int16
func divint16s(x, y int16) int16
func addint64s(x, y int64) int64
func subint64s(x, y int64) int64
func negint64s(x int64) int64
func mulint64s(x, y int64) int64
func divint64s(x, y int64) int64
func adduint8s(x, y uint8) uint8
func subuint8s(x, y uint8) uint8
func muluint8s(x, y uint8) uint8
func divuint8s(x, y uint8) uint8
func adduint16s(x, y uint16) uint16
func subuint16s(x, y uint16) uint16
func muluint16s(x, y uint16) uint16
func divuint16s(x, y uint16) uint16
func adduint32s(x, y uint32) uint32
func subuint32s(x, y uint32) uint32
func muluint32s(x, y uint32) uint32
func divuint32s(x, y uint32) uint32
func adduint64s(x, y uint64) uint64
func subuint64s(x, y uint64) uint64
func muluint64s(x, y uint64) uint64
func divuint64s(x, y uint64) uint64

func add(x, y int32) int32 {
	return x + y
}

func sub(x, y int32) int32 {
	return x - y
}
func neg(x int32) int32 {
	return -x
}
func mul(x, y int32) int32 {
	return x * y
}
func div(x, y int32) int32 {
	return x / y
}
func addint8(x, y int8) int8 {
	return x + y
}
func subint8(x, y int8) int8 {
	return x - y
}
func negint8(x int8) int8 {
	return -x
}
func mulint8(x, y int8) int8 {
	return x * y
}
func divint8(x, y int8) int8 {
	return x / y
}
func addint16(x, y int16) int16 {
	return x + y
}
func subint16(x, y int16) int16 {
	return x - y
}
func negint16(x int16) int16 {
	return -x
}
func mulint16(x, y int16) int16 {
	return x * y
}
func divint16(x, y int16) int16 {
	return x / y
}
func addint64(x, y int64) int64 {
	return x + y
}
func subint64(x, y int64) int64 {
	return x - y
}
func negint64(x int64) int64 {
	return -x
}
func mulint64(x, y int64) int64 {
	return x * y
}
func divint64(x, y int64) int64 {
	return x / y
}
func adduint8(x, y uint8) uint8 {
	return x + y
}
func subuint8(x, y uint8) uint8 {
	return x - y
}
func muluint8(x, y uint8) uint8 {
	return x * y
}
func divuint8(x, y uint8) uint8 {
	return x / y
}
func adduint16(x, y uint16) uint16 {
	return x + y
}
func subuint16(x, y uint16) uint16 {
	return x - y
}
func muluint16(x, y uint16) uint16 {
	return x * y
}
func divuint16(x, y uint16) uint16 {
	return x / y
}
func adduint32(x, y uint32) uint32 {
	return x + y
}
func subuint32(x, y uint32) uint32 {
	return x - y
}
func muluint32(x, y uint32) uint32 {
	return x * y
}
func divuint32(x, y uint32) uint32 {
	return x / y
}
func adduint64(x, y uint64) uint64 {
	return x + y
}
func subuint64(x, y uint64) uint64 {
	return x - y
}
func muluint64(x, y uint64) uint64 {
	return x * y
}
func divuint64(x, y uint64) uint64 {
	return x / y
}

func TestArithmeticOps(t *testing.T) {

	for i := -63; i <= 63; i++ {

		y := int64(0)
		if i < 0 {
			y = -1 << uint(-i)
		} else {
			y = 1<<uint(i) - 1
		}

		if negint8s(int8(y)) != negint8(int8(y)) {
			t.Errorf("negints(%v)", int8(y))
		}
		if negint16s(int16(y)) != negint16(int16(y)) {
			t.Errorf("negint16s(%v)", int16(y))
		}
		if negs(int32(y)) != neg(int32(y)) {
			t.Errorf("negs(%v)", int32(y))
		}
		if negint64s(y) != negint64(y) {
			t.Errorf("negs(%v)", y)
		}

		for j := -63; j <= 63; j++ {

			x := int64(0)
			if j < 0 {
				x = -1 << uint(-j)
			} else {
				x = 1<<uint(j) - 1
			}

			if adds(int32(x), int32(y)) != add(int32(x), int32(y)) {
				t.Errorf("adds(%v, %v)", int32(x), int32(y))
			}
			if subs(int32(x), int32(y)) != sub(int32(x), int32(y)) {
				t.Errorf("subs(%v, %v)", int32(x), int32(y))

			}
			if muls(int32(x), int32(y)) != mul(int32(x), int32(y)) {
				t.Errorf("muls(%v, %v)", int32(x), int32(y))
			}

			if addint8s(int8(x), int8(y)) != addint8(int8(x), int8(y)) {
				t.Errorf("addint8s(%v, %v)", int8(x), int8(y))
			}
			if subint8s(int8(x), int8(y)) != subint8(int8(x), int8(y)) {
				t.Errorf("subint8s(%v, %v)", int8(x), int8(y))

			}
			if mulint8s(int8(x), int8(y)) != mulint8(int8(x), int8(y)) {
				t.Errorf("mulint8s(%v, %v)", int8(x), int8(y))
			}

			if addint16s(int16(x), int16(y)) != addint16(int16(x), int16(y)) {
				t.Errorf("addint16s(%v, %v)", int16(x), int16(y))
			}
			if subint16s(int16(x), int16(y)) != subint16(int16(x), int16(y)) {
				t.Errorf("subint16s(%v, %v)", int16(x), int16(y))

			}
			if mulint16s(int16(x), int16(y)) != mulint16(int16(x), int16(y)) {
				t.Errorf("mulint16s(%v, %v)", int16(x), int16(y))
			}

			if addint64s(int64(x), int64(y)) != addint64(int64(x), int64(y)) {
				t.Errorf("addint64s(%v, %v)", int64(x), int64(y))
			}
			if subint64s(int64(x), int64(y)) != subint64(int64(x), int64(y)) {
				t.Errorf("subint64s(%v, %v)", int64(x), int64(y))

			}
			if mulint64s(int64(x), int64(y)) != mulint64(int64(x), int64(y)) {
				t.Errorf("mulint64s(%v, %v)", int64(x), int64(y))
			}

			if adduint8s(uint8(x), uint8(y)) != adduint8(uint8(x), uint8(y)) {
				t.Errorf("adduint8s(%v, %v)", uint8(x), uint8(y))
			}
			if subuint8s(uint8(x), uint8(y)) != subuint8(uint8(x), uint8(y)) {
				t.Errorf("subuint8s(%v, %v)", uint8(x), uint8(y))

			}
			if muluint8s(uint8(x), uint8(y)) != muluint8(uint8(x), uint8(y)) {
				t.Errorf("muluint8s(%v, %v)", uint8(x), uint8(y))
			}

			if adduint16s(uint16(x), uint16(y)) != adduint16(uint16(x), uint16(y)) {
				t.Errorf("adduint16s(%v, %v)", uint16(x), uint16(y))
			}
			if subuint16s(uint16(x), uint16(y)) != subuint16(uint16(x), uint16(y)) {
				t.Errorf("subuint16s(%v, %v)", uint16(x), uint16(y))

			}
			if muluint16s(uint16(x), uint16(y)) != muluint16(uint16(x), uint16(y)) {
				t.Errorf("muluint16s(%v, %v)", uint16(x), uint16(y))
			}

			if adduint32s(uint32(x), uint32(y)) != adduint32(uint32(x), uint32(y)) {
				t.Errorf("adduint32s(%v, %v)", uint32(x), uint32(y))
			}
			if subuint32s(uint32(x), uint32(y)) != subuint32(uint32(x), uint32(y)) {
				t.Errorf("subuint32s(%v, %v)", uint32(x), uint32(y))

			}
			if muluint32s(uint32(x), uint32(y)) != muluint32(uint32(x), uint32(y)) {
				t.Errorf("muluint32s(%v, %v)", uint32(x), uint32(y))
			}

			if adduint64s(uint64(x), uint64(y)) != adduint64(uint64(x), uint64(y)) {
				t.Errorf("adduint64s(%v, %v)", uint64(x), uint64(y))
			}
			if subuint64s(uint64(x), uint64(y)) != subuint64(uint64(x), uint64(y)) {
				t.Errorf("subuint64s(%v, %v)", uint64(x), uint64(y))

			}
			if muluint64s(uint64(x), uint64(y)) != muluint64(uint64(x), uint64(y)) {
				t.Errorf("muluint64s(%v, %v)", uint64(x), uint64(y))
			}

		}
	}
}
