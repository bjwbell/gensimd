// +build amd64,gc

package tests

import "testing"

//go:generate gensimd -fn "oruint8, anduint8, xoruint8, notuint8, andnotuint8, shluint8, shruint8, oruint16, anduint16, xoruint16, notuint16, andnotuint16, shluint16, shruint16, oruint32, anduint32, xoruint32, notuint32, andnotuint32, shluint32, shruint32, oruint64, anduint64, xoruint64, notuint64, andnotuint64, shluint64, shruint64, orint8, andint8, xorint8, notint8, andnotint8, shlint8, shrint8, orint16, andint16, xorint16, notint16, andnotint16, shlint16, shrint16, orint32, andint32, xorint32, notint32, andnotint32, shlint32, shrint32, orint64, andint64, xorint64, notint64, andnotint64, shlint64, shrint64" -outfn "oruint8s, anduint8s, xoruint8s, notuint8s, andnotuint8s, shluint8s, shruint8s, oruint16s, anduint16s, xoruint16s, notuint16s, andnotuint16s, shluint16s, shruint16s, oruint32s, anduint32s, xoruint32s, notuint32s, andnotuint32s, shluint32s, shruint32s, oruint64s, anduint64s, xoruint64s, notuint64s, andnotuint64s, shluint64s, shruint64s, orint8s, andint8s, xorint8s, notint8s, andnotint8s, shlint8s, shrint8s, orint16s, andint16s, xorint16s, notint16s, andnotint16s, shlint16s, shrint16s, orint32s, andint32s, xorint32s, notint32s, andnotint32s, shlint32s, shrint32s, orint64s, andint64s, xorint64s, notint64s, andnotint64s, shlint64s, shrint64s" -f "$GOFILE" -o "bitwiseops_test.s"

func oruint8s(a, b uint8) uint8
func anduint8s(a, b uint8) uint8
func xoruint8s(a, b uint8) uint8
func notuint8s(a uint8) uint8
func andnotuint8s(a, b uint8) uint8
func shluint8s(x, shift uint8) uint8
func shruint8s(x, shift uint8) uint8
func oruint16s(a, b uint16) uint16
func anduint16s(a, b uint16) uint16
func xoruint16s(a, b uint16) uint16
func notuint16s(a uint16) uint16
func andnotuint16s(a, b uint16) uint16
func shluint16s(x uint16, shift uint8) uint16
func shruint16s(x uint16, shift uint8) uint16
func oruint32s(a, b uint32) uint32
func anduint32s(a, b uint32) uint32
func xoruint32s(a, b uint32) uint32
func notuint32s(a uint32) uint32
func andnotuint32s(a, b uint32) uint32
func shluint32s(x uint32, shift uint8) uint32
func shruint32s(x uint32, shift uint8) uint32
func oruint64s(a, b uint64) uint64
func anduint64s(a, b uint64) uint64
func xoruint64s(a, b uint64) uint64
func notuint64s(a uint64) uint64
func andnotuint64s(a, b uint64) uint64
func shluint64s(x uint64, shift uint8) uint64
func shruint64s(x uint64, shift uint8) uint64
func orint8s(a, b int8) int8
func andint8s(a, b int8) int8
func xorint8s(a, b int8) int8
func notint8s(a int8) int8
func andnotint8s(a, b int8) int8
func shlint8s(x int8, shift uint8) int8
func shrint8s(x int8, shift uint8) int8
func orint16s(a, b int16) int16
func andint16s(a, b int16) int16
func xorint16s(a, b int16) int16
func notint16s(a int16) int16
func andnotint16s(a, b int16) int16
func shlint16s(x int16, shift uint8) int16
func shrint16s(x int16, shift uint8) int16
func orint32s(a, b int32) int32
func andint32s(a, b int32) int32
func xorint32s(a, b int32) int32
func notint32s(a int32) int32
func andnotint32s(a, b int32) int32
func shlint32s(x int32, shift uint8) int32
func shrint32s(x int32, shift uint8) int32
func orint64s(a, b int64) int64
func andint64s(a, b int64) int64
func xorint64s(a, b int64) int64
func notint64s(a int64) int64
func andnotint64s(a, b int64) int64
func shlint64s(x int64, shift uint8) int64
func shrint64s(x int64, shift uint8) int64

func oruint8(a, b uint8) uint8 {
	return a | b
}
func anduint8(a, b uint8) uint8 {
	return a & b
}
func xoruint8(a, b uint8) uint8 {
	return a ^ b
}
func notuint8(a uint8) uint8 {
	return ^a
}
func andnotuint8(a, b uint8) uint8 {
	return a &^ b
}
func shluint8(x, shift uint8) uint8 {
	return x << shift
}
func shruint8(x, shift uint8) uint8 {
	return x >> shift
}
func oruint16(a, b uint16) uint16 {
	return a | b
}
func anduint16(a, b uint16) uint16 {
	return a & b
}
func xoruint16(a, b uint16) uint16 {
	return a ^ b
}
func notuint16(a uint16) uint16 {
	return ^a
}
func andnotuint16(a, b uint16) uint16 {
	return a &^ b
}
func shluint16(x uint16, shift uint8) uint16 {
	return x << shift
}
func shruint16(x uint16, shift uint8) uint16 {
	return x >> shift
}
func oruint32(a, b uint32) uint32 {
	return a | b
}
func anduint32(a, b uint32) uint32 {
	return a & b
}
func xoruint32(a, b uint32) uint32 {
	return a ^ b
}
func notuint32(a uint32) uint32 {
	return ^a
}
func andnotuint32(a, b uint32) uint32 {
	return a &^ b
}
func shluint32(x uint32, shift uint8) uint32 {
	return x << shift
}
func shruint32(x uint32, shift uint8) uint32 {
	return x >> shift
}
func oruint64(a, b uint64) uint64 {
	return a | b
}
func anduint64(a, b uint64) uint64 {
	return a & b
}
func xoruint64(a, b uint64) uint64 {
	return a ^ b
}
func notuint64(a uint64) uint64 {
	return ^a
}
func andnotuint64(a, b uint64) uint64 {
	return a &^ b
}
func shluint64(x uint64, shift uint8) uint64 {
	return x << shift
}
func shruint64(x uint64, shift uint8) uint64 {
	return x >> shift
}
func orint8(a, b int8) int8 {
	return a | b
}
func andint8(a, b int8) int8 {
	return a & b
}
func xorint8(a, b int8) int8 {
	return a ^ b
}
func notint8(a int8) int8 {
	return ^a
}
func andnotint8(a, b int8) int8 {
	return a &^ b
}
func shlint8(x int8, shift uint8) int8 {
	return x << shift
}
func shrint8(x int8, shift uint8) int8 {
	return x >> shift
}
func orint16(a, b int16) int16 {
	return a | b
}
func andint16(a, b int16) int16 {
	return a & b
}
func xorint16(a, b int16) int16 {
	return a ^ b
}
func notint16(a int16) int16 {
	return ^a
}
func andnotint16(a, b int16) int16 {
	return a &^ b
}
func shlint16(x int16, shift uint8) int16 {
	return x << shift
}
func shrint16(x int16, shift uint8) int16 {
	return x >> shift
}
func orint32(a, b int32) int32 {
	return a | b
}
func andint32(a, b int32) int32 {
	return a & b
}
func xorint32(a, b int32) int32 {
	return a ^ b
}
func notint32(a int32) int32 {
	return ^a
}
func andnotint32(a, b int32) int32 {
	return a &^ b
}
func shlint32(x int32, shift uint8) int32 {
	return x << shift
}
func shrint32(x int32, shift uint8) int32 {
	return x >> shift
}
func orint64(a, b int64) int64 {
	return a | b
}
func andint64(a, b int64) int64 {
	return a & b
}
func xorint64(a, b int64) int64 {
	return a ^ b
}
func notint64(a int64) int64 {
	return ^a
}
func andnotint64(a, b int64) int64 {
	return a &^ b
}
func shlint64(x int64, shift uint8) int64 {
	return x << shift
}
func shrint64(x int64, shift uint8) int64 {
	return x >> shift
}

func TestBitwiseOps(t *testing.T) {

	for i := -63; i <= 63; i++ {

		a := int64(0)
		if i < 0 {
			a = -1 << uint(-i)
		} else {
			a = 1<<uint(i) - 1
		}

		if notuint8s(uint8(a)) != notuint8(uint8(a)) {
			t.Errorf("(%v) notuint8s(%v) != (%v) notuint8(%v)", notuint8s(uint8(a)), uint8(a), notuint8(uint8(a)), uint8(a))
		}
		if notuint16s(uint16(a)) != notuint16(uint16(a)) {
			t.Errorf("notuint16s(%v)", uint16(a))
		}
		if notuint32s(uint32(a)) != notuint32(uint32(a)) {
			t.Errorf("notuint32s(%v)", uint32(a))
		}
		if notuint64s(uint64(a)) != notuint64(uint64(a)) {
			t.Errorf("notuint64s(%v)", uint64(a))
		}
		if notint8s(int8(a)) != notint8(int8(a)) {
			t.Errorf("notint8s(%v)", int8(a))
		}
		if notint16s(int16(a)) != notint16(int16(a)) {
			t.Errorf("notint16s(%v)", int16(a))
		}
		if notint32s(int32(a)) != notint32(int32(a)) {
			t.Errorf("notint32s(%v)", int32(a))
		}
		if notint64s(int64(a)) != notint64(int64(a)) {
			t.Errorf("notint64s(%v)", int64(a))
		}

		for j := -63; j <= 63; j++ {

			b := int64(0)
			if j < 0 {
				b = -1 << uint(-j)
			} else {
				b = 1<<uint(j) - 1
			}

			shift := uint8(j)

			if oruint8s(uint8(a), uint8(b)) != oruint8(uint8(a), uint8(b)) {
				t.Errorf("oruint8s(%v, %v)", uint8(a), uint8(b))
			}
			if anduint8s(uint8(a), uint8(b)) != anduint8(uint8(a), uint8(b)) {
				t.Errorf("anduint8s(%v, %v)", uint8(a), uint8(b))
			}
			if xoruint8s(uint8(a), uint8(b)) != xoruint8(uint8(a), uint8(b)) {
				t.Errorf("xoruint8s(%v, %v)", uint8(a), uint8(b))
			}
			if andnotuint8s(uint8(a), uint8(b)) != andnotuint8(uint8(a), uint8(b)) {
				t.Errorf("(%v) andnotuint8s(%v, %v) != %v", andnotuint8s(uint8(a), uint8(b)), uint8(a), uint8(b), andnotuint8(uint8(a), uint8(b)))
			}
			if shluint8s(uint8(a), shift) != shluint8(uint8(a), shift) {
				t.Errorf("(%v) shluint8s(%v, %v) != %v", shluint8s(uint8(a), shift), uint8(a), shift, shluint8(uint8(a), shift))
			}
			if shruint8s(uint8(a), shift) != shruint8(uint8(a), shift) {
				t.Errorf("shruint8s(%v, %v)", uint8(a), shift)
			}
			if oruint16s(uint16(a), uint16(b)) != oruint16(uint16(a), uint16(b)) {
				t.Errorf("oruint16s(%v, %v)", uint16(a), uint16(b))
			}
			if anduint16s(uint16(a), uint16(b)) != anduint16(uint16(a), uint16(b)) {
				t.Errorf("anduint16s(%v, %v)", uint16(a), uint16(b))
			}
			if xoruint16s(uint16(a), uint16(b)) != xoruint16(uint16(a), uint16(b)) {
				t.Errorf("xoruint16s(%v, %v)", uint16(a), uint16(b))
			}
			if andnotuint16s(uint16(a), uint16(b)) != andnotuint16(uint16(a), uint16(b)) {
				t.Errorf("andnotuint16s(%v, %v)", uint16(a), uint16(b))
			}
			if shluint16s(uint16(a), shift) != shluint16(uint16(a), shift) {
				t.Errorf("%v shluint16s(%v, %v) != %v", shluint16s(uint16(a), shift), uint16(a), shift, shluint16(uint16(a), shift))
			}
			if shruint16s(uint16(a), shift) != shruint16(uint16(a), shift) {
				t.Errorf("%v shruint16s(%v, %v) != %v", shruint16s(uint16(a), shift), uint16(a), shift, shruint16(uint16(a), shift))
			}
			if oruint32s(uint32(a), uint32(b)) != oruint32(uint32(a), uint32(b)) {
				t.Errorf("oruint32s(%v, %v)", uint32(a), uint32(b))
			}
			if anduint32s(uint32(a), uint32(b)) != anduint32(uint32(a), uint32(b)) {
				t.Errorf("anduint32s(%v, %v)", uint32(a), uint32(b))
			}
			if xoruint32s(uint32(a), uint32(b)) != xoruint32(uint32(a), uint32(b)) {
				t.Errorf("xoruint32s(%v, %v)", uint32(a), uint32(b))
			}
			if andnotuint32s(uint32(a), uint32(b)) != andnotuint32(uint32(a), uint32(b)) {
				t.Errorf("andnotuint32s(%v, %v)", uint32(a), uint32(b))
			}
			if shluint32s(uint32(a), shift) != shluint32(uint32(a), shift) {
				t.Errorf("shluint32s(%v, %v) %v != %v", uint32(a), shift, shluint32s(uint32(a), shift), shluint32(uint32(a), shift))
			}
			if shruint32s(uint32(a), shift) != shruint32(uint32(a), shift) {
				t.Errorf("shruint32s(%v, %v)", uint32(a), shift)
			}
			if oruint64s(uint64(a), uint64(b)) != oruint64(uint64(a), uint64(b)) {
				t.Errorf("oruint64s(%v, %v)", uint64(a), uint64(b))
			}
			if anduint64s(uint64(a), uint64(b)) != anduint64(uint64(a), uint64(b)) {
				t.Errorf("anduint64s(%v, %v)", uint64(a), uint64(b))
			}
			if xoruint64s(uint64(a), uint64(b)) != xoruint64(uint64(a), uint64(b)) {
				t.Errorf("xoruint64s(%v, %v)", uint64(a), uint64(b))
			}
			if andnotuint64s(uint64(a), uint64(b)) != andnotuint64(uint64(a), uint64(b)) {
				t.Errorf("andnotuint64s(%v, %v)", uint64(a), uint64(b))
			}
			if shluint64s(uint64(a), shift) != shluint64(uint64(a), shift) {
				t.Errorf("shluint64s(%v, %v) %v != %v", uint64(a), shift, shluint64s(uint64(a), shift), shluint64(uint64(a), shift))
			}
			if shruint64s(uint64(a), shift) != shruint64(uint64(a), shift) {
				t.Errorf("shruint64s(%v, %v) %v != %v", uint64(a), shift, shruint64s(uint64(a), shift), shruint64(uint64(a), shift))
			}
			if orint8s(int8(a), int8(b)) != orint8(int8(a), int8(b)) {
				t.Errorf("orint8s(%v, %v)", int8(a), int8(b))
			}
			if andint8s(int8(a), int8(b)) != andint8(int8(a), int8(b)) {
				t.Errorf("andint8s(%v, %v)", int8(a), int8(b))
			}
			if xorint8s(int8(a), int8(b)) != xorint8(int8(a), int8(b)) {
				t.Errorf("xorint8s(%v, %v)", int8(a), int8(b))
			}
			if andnotint8s(int8(a), int8(b)) != andnotint8(int8(a), int8(b)) {
				t.Errorf("andnotint8s(%v, %v)", int8(a), int8(b))
			}
			if shlint8s(int8(a), shift) != shlint8(int8(a), shift) {
				t.Errorf("shlint8s(%v, %v) %v != %v", int8(a), shift, shlint8s(int8(a), shift), shlint8(int8(a), shift))
			}
			if shrint8s(int8(a), shift) != shrint8(int8(a), shift) {
				t.Errorf("shrint8s(%v, %v) %v != %v", int8(a), shift, shrint8s(int8(a), shift), shrint8(int8(a), shift))
			}
			if orint16s(int16(a), int16(b)) != orint16(int16(a), int16(b)) {
				t.Errorf("orint16s(%v, %v)", int16(a), int16(b))
			}
			if andint16s(int16(a), int16(b)) != andint16(int16(a), int16(b)) {
				t.Errorf("andint16s(%v, %v)", int16(a), int16(b))
			}
			if xorint16s(int16(a), int16(b)) != xorint16(int16(a), int16(b)) {
				t.Errorf("xorint16s(%v, %v)", int16(a), int16(b))
			}
			if andnotint16s(int16(a), int16(b)) != andnotint16(int16(a), int16(b)) {
				t.Errorf("andnotint16s(%v, %v)", int16(a), int16(b))
			}
			if shlint16s(int16(a), shift) != shlint16(int16(a), shift) {
				t.Errorf("shlint16s(%v, %v)", int16(a), shift)
			}
			if shrint16s(int16(a), shift) != shrint16(int16(a), shift) {
				t.Errorf("shrint16s(%v, %v)", int16(a), shift)
			}
			if orint32s(int32(a), int32(b)) != orint32(int32(a), int32(b)) {
				t.Errorf("orint32s(%v, %v)", int32(a), int32(b))
			}
			if andint32s(int32(a), int32(b)) != andint32(int32(a), int32(b)) {
				t.Errorf("andint32s(%v, %v)", int32(a), int32(b))
			}
			if xorint32s(int32(a), int32(b)) != xorint32(int32(a), int32(b)) {
				t.Errorf("xorint32s(%v, %v)", int32(a), int32(b))
			}
			if andnotint32s(int32(a), int32(b)) != andnotint32(int32(a), int32(b)) {
				t.Errorf("andnotint32s(%v, %v)", int32(a), int32(b))
			}
			if shlint32s(int32(a), shift) != shlint32(int32(a), shift) {
				t.Errorf("shlint32s(%v, %v)", int32(a), shift)
			}
			if shrint32s(int32(a), shift) != shrint32(int32(a), shift) {
				t.Errorf("shrint32s(%v, %v)", int32(a), shift)
			}
			if orint64s(int64(a), int64(b)) != orint64(int64(a), int64(b)) {
				t.Errorf("orint64s(%v, %v)", int64(a), int64(b))
			}
			if andint64s(int64(a), int64(b)) != andint64(int64(a), int64(b)) {
				t.Errorf("andint64s(%v, %v)", int64(a), int64(b))
			}
			if xorint64s(int64(a), int64(b)) != xorint64(int64(a), int64(b)) {
				t.Errorf("xorint64s(%v, %v)", int64(a), int64(b))
			}
			if andnotint64s(int64(a), int64(b)) != andnotint64(int64(a), int64(b)) {
				t.Errorf("andnotint64s(%v, %v)", int64(a), int64(b))
			}
			if shlint64s(int64(a), shift) != shlint64(int64(a), shift) {
				t.Errorf("shlint64s(%v, %v)", int64(a), shift)
			}
			if shrint64s(int64(a), shift) != shrint64(int64(a), shift) {
				t.Errorf("shrint64s(%v, %v)", int64(a), shift)
			}

		}
	}
}
