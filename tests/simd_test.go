// +build amd64,gc

package tests

import (
	"math"
	"math/rand"
	"testing"

	"github.com/bjwbell/gensimd/simd"
)

//go:generate gensimd -fn "addi8x16, subi8x16, addu8x16, subu8x16, addi16x8, subi16x8, muli16x8, shli16x8, shri16x8, addu16x8, subu16x8, mulu16x8, shlu16x8, shru16x8, addi32x4, subi32x4, muli32x4, shli32x4, shri32x4, addu32x4, subu32x4, mulu32x4, shlu32x4, shru32x4, addi64x2, subi64x2, addu64x2, subu64x2, addf32x4, subf32x4, mulf32x4, divf32x4, addf64x2, subf64x2, mulf64x2, divf64x2" -outfn "addi8x16s, subi8x16s, addu8x16s, subu8x16s, addi16x8s, subi16x8s, muli16x8s, shli16x8s, shri16x8s, addu16x8s, subu16x8s, mulu16x8s, shlu16x8s, shru16x8s, addi32x4s, subi32x4s, muli32x4s, shli32x4s, shri32x4s, addu32x4s, subu32x4s, mulu32x4s, shlu32x4s, shru32x4s, addi64x2s, subi64x2s, addu64x2s, subu64x2s, addf32x4s, subf32x4s, mulf32x4s, divf32x4s, addf64x2s, subf64x2s, mulf64x2s, divf64x2s" -f "$GOFILE" -o "simd_test_amd64.s"

func addi8x16s(x, y simd.I8x16) simd.I8x16
func subi8x16s(x, y simd.I8x16) simd.I8x16
func addu8x16s(x, y simd.U8x16) simd.U8x16
func subu8x16s(x, y simd.U8x16) simd.U8x16

func addi16x8s(x, y simd.I16x8) simd.I16x8
func subi16x8s(x, y simd.I16x8) simd.I16x8
func muli16x8s(x, y simd.I16x8) simd.I16x8
func shli16x8s(x simd.I16x8, shift uint8) simd.I16x8
func shri16x8s(x simd.I16x8, shift uint8) simd.I16x8
func addu16x8s(x, y simd.U16x8) simd.U16x8
func subu16x8s(x, y simd.U16x8) simd.U16x8
func mulu16x8s(x, y simd.U16x8) simd.U16x8
func shlu16x8s(x simd.U16x8, shift uint8) simd.U16x8
func shru16x8s(x simd.U16x8, shift uint8) simd.U16x8

func addi32x4s(x, y simd.I32x4) simd.I32x4
func subi32x4s(x, y simd.I32x4) simd.I32x4
func muli32x4s(x, y simd.I32x4) simd.I32x4
func shli32x4s(x simd.I32x4, shift uint8) simd.I32x4
func shri32x4s(x simd.I32x4, shift uint8) simd.I32x4
func addu32x4s(x, y simd.U32x4) simd.U32x4
func subu32x4s(x, y simd.U32x4) simd.U32x4
func mulu32x4s(x, y simd.U32x4) simd.U32x4
func shlu32x4s(x simd.U32x4, shift uint8) simd.U32x4
func shru32x4s(x simd.U32x4, shift uint8) simd.U32x4

func addi64x2s(x, y simd.I64x2) simd.I64x2
func subi64x2s(x, y simd.I64x2) simd.I64x2
func addu64x2s(x, y simd.U64x2) simd.U64x2
func subu64x2s(x, y simd.U64x2) simd.U64x2

func addf32x4s(x, y simd.F32x4) simd.F32x4
func subf32x4s(x, y simd.F32x4) simd.F32x4
func mulf32x4s(x, y simd.F32x4) simd.F32x4
func divf32x4s(x, y simd.F32x4) simd.F32x4
func addf64x2s(x, y simd.F64x2) simd.F64x2
func subf64x2s(x, y simd.F64x2) simd.F64x2
func mulf64x2s(x, y simd.F64x2) simd.F64x2
func divf64x2s(x, y simd.F64x2) simd.F64x2

func addi8x16(x, y simd.I8x16) simd.I8x16 { return simd.AddI8x16(x, y) }
func subi8x16(x, y simd.I8x16) simd.I8x16 { return simd.SubI8x16(x, y) }
func addu8x16(x, y simd.U8x16) simd.U8x16 { return simd.AddU8x16(x, y) }
func subu8x16(x, y simd.U8x16) simd.U8x16 { return simd.SubU8x16(x, y) }

func addi16x8(x, y simd.I16x8) simd.I16x8           { return simd.AddI16x8(x, y) }
func subi16x8(x, y simd.I16x8) simd.I16x8           { return simd.SubI16x8(x, y) }
func muli16x8(x, y simd.I16x8) simd.I16x8           { return simd.MulI16x8(x, y) }
func shli16x8(x simd.I16x8, shift uint8) simd.I16x8 { return simd.ShlI16x8(x, shift) }
func shri16x8(x simd.I16x8, shift uint8) simd.I16x8 { return simd.ShrI16x8(x, shift) }
func addu16x8(x, y simd.U16x8) simd.U16x8           { return simd.AddU16x8(x, y) }
func subu16x8(x, y simd.U16x8) simd.U16x8           { return simd.SubU16x8(x, y) }
func mulu16x8(x, y simd.U16x8) simd.U16x8           { return simd.MulU16x8(x, y) }
func shlu16x8(x simd.U16x8, shift uint8) simd.U16x8 { return simd.ShlU16x8(x, shift) }
func shru16x8(x simd.U16x8, shift uint8) simd.U16x8 { return simd.ShrU16x8(x, shift) }

func addi32x4(x, y simd.I32x4) simd.I32x4           { return simd.AddI32x4(x, y) }
func subi32x4(x, y simd.I32x4) simd.I32x4           { return simd.SubI32x4(x, y) }
func muli32x4(x, y simd.I32x4) simd.I32x4           { return simd.MulI32x4(x, y) }
func shli32x4(x simd.I32x4, shift uint8) simd.I32x4 { return simd.ShlI32x4(x, shift) }
func shri32x4(x simd.I32x4, shift uint8) simd.I32x4 { return simd.ShrI32x4(x, shift) }
func addu32x4(x, y simd.U32x4) simd.U32x4           { return simd.AddU32x4(x, y) }
func subu32x4(x, y simd.U32x4) simd.U32x4           { return simd.SubU32x4(x, y) }
func mulu32x4(x, y simd.U32x4) simd.U32x4           { return simd.MulU32x4(x, y) }
func shlu32x4(x simd.U32x4, shift uint8) simd.U32x4 { return simd.ShlU32x4(x, shift) }
func shru32x4(x simd.U32x4, shift uint8) simd.U32x4 { return simd.ShrU32x4(x, shift) }

func addi64x2(x, y simd.I64x2) simd.I64x2 { return simd.AddI64x2(x, y) }
func subi64x2(x, y simd.I64x2) simd.I64x2 { return simd.SubI64x2(x, y) }
func addu64x2(x, y simd.U64x2) simd.U64x2 { return simd.AddU64x2(x, y) }
func subu64x2(x, y simd.U64x2) simd.U64x2 { return simd.SubU64x2(x, y) }

func addf32x4(x, y simd.F32x4) simd.F32x4 { return simd.AddF32x4(x, y) }
func subf32x4(x, y simd.F32x4) simd.F32x4 { return simd.SubF32x4(x, y) }
func mulf32x4(x, y simd.F32x4) simd.F32x4 { return simd.MulF32x4(x, y) }
func divf32x4(x, y simd.F32x4) simd.F32x4 { return simd.DivF32x4(x, y) }

func addf64x2(x, y simd.F64x2) simd.F64x2 { return simd.AddF64x2(x, y) }
func subf64x2(x, y simd.F64x2) simd.F64x2 { return simd.SubF64x2(x, y) }
func mulf64x2(x, y simd.F64x2) simd.F64x2 { return simd.MulF64x2(x, y) }
func divf64x2(x, y simd.F64x2) simd.F64x2 { return simd.DivF64x2(x, y) }

func TestSimd(t *testing.T) {

	count := 0
	rand.Seed(42)

	for i := -63; i <= 63; i++ {

		a := int(0)
		if i < 0 {
			a = -1 << uint(-i)
		} else {
			a = 1<<uint(i) - 1
		}

		for j := -63; j <= 63; j++ {

			count++

			b := int(0)
			if j < 0 {
				b = -1 << uint(-j)
			} else {
				b = 1<<uint(j) - 1
			}

			shift := uint8(j)

			var x simd.I8x16
			var y simd.I8x16
			var xu simd.U8x16
			var yu simd.U8x16

			var xI16x8 simd.I16x8
			var yI16x8 simd.I16x8
			var xU16x8 simd.U16x8
			var yU16x8 simd.U16x8

			var xI32x4 simd.I32x4
			var yI32x4 simd.I32x4
			var xU32x4 simd.U32x4
			var yU32x4 simd.U32x4

			var xI64x2 simd.I64x2
			var yI64x2 simd.I64x2
			var xU64x2 simd.U64x2
			var yU64x2 simd.U64x2

			var xF32x4 simd.F32x4
			var yF32x4 simd.F32x4

			var xF64x2 simd.F64x2
			var yF64x2 simd.F64x2

			for idx := 0; idx < 16; idx++ {
				abs_a := a

				if abs_a == math.MinInt64 || abs_a == 0 {
					abs_a = 64
				}
				if abs_a < 0 {
					abs_a = -abs_a
				}
				abs_b := b
				if abs_b == math.MinInt64 || abs_b == 0 {
					abs_b = 64
				}
				if abs_b < 0 {
					abs_b = -abs_b
				}

				x[idx] = int8(rand.Intn(abs_a))
				y[idx] = int8(rand.Intn(abs_b))
				xu[idx] = uint8(rand.Intn(abs_a))
				yu[idx] = uint8(rand.Intn(abs_b))

				if idx < 8 {
					xI16x8[idx] = int16(rand.Intn(abs_a))
					yI16x8[idx] = int16(rand.Intn(abs_b))
					xU16x8[idx] = uint16(rand.Intn(abs_a))
					yU16x8[idx] = uint16(rand.Intn(abs_b))
				}

				if idx < 4 {
					xI32x4[idx] = int32(rand.Intn(abs_a))
					yI32x4[idx] = int32(rand.Intn(abs_b))
					xU32x4[idx] = uint32(rand.Intn(abs_a))
					yU32x4[idx] = uint32(rand.Intn(abs_b))
					xF32x4[idx] = rand.Float32()
					yF32x4[idx] = rand.Float32()
				}

				if idx < 2 {
					xI64x2[idx] = int64(rand.Intn(abs_a))
					yI64x2[idx] = int64(rand.Intn(abs_b))
					xU64x2[idx] = uint64(rand.Intn(abs_a) +
						math.MaxInt64)
					yU64x2[idx] = uint64(rand.Intn(abs_b) +
						math.MaxInt64)
					xF64x2[idx] = rand.ExpFloat64()
					yF64x2[idx] = rand.ExpFloat64()
				}

			}

			if addi8x16s(x, y) != addi8x16(x, y) {
				t.Errorf("addi8x16(%v, %v)", x, y)
				t.Error("x:", x)
				t.Error("y:", y)
				t.Error("s(x, y):", addi8x16s(x, y))
				t.Error(" (x, y):", addi8x16(x, y))
			}
			if subi8x16s(x, y) != subi8x16(x, y) {
				t.Errorf("subi8x16(%v, %v)", x, y)
				t.Error("x:", x)
				t.Error("y:", y)
				t.Error("s:", subi8x16s(x, y))
				t.Error(" :", subi8x16(x, y))
			}
			if addu8x16s(xu, yu) != addu8x16(xu, yu) {
				t.Errorf("addu8x16(%v, %v)", xu, yu)
			}
			if subu8x16s(xu, yu) != subu8x16(xu, yu) {
				t.Errorf("subu8x16(%v, %v)", xu, yu)
				t.Error("x:", xu)
				t.Error("y:", yu)
				t.Error("s:", subu8x16s(xu, yu))
				t.Error(" :", subu8x16(xu, yu))
			}

			if addi16x8s(xI16x8, yI16x8) != addi16x8(xI16x8, yI16x8) {
				t.Errorf("addi16x8(%v, %v)", xI16x8, yI16x8)
				t.Error("x:", xI16x8)
				t.Error("y:", yI16x8)
				t.Error("s(xI16x8, yI16x8):", addi16x8s(xI16x8, yI16x8))
				t.Error(" (xI16x8, yI16x8):", addi16x8(xI16x8, yI16x8))
			}
			if subi16x8s(xI16x8, yI16x8) != subi16x8(xI16x8, yI16x8) {
				t.Errorf("subi16x8(%v, %v)", xI16x8, yI16x8)
				t.Error("x:", xI16x8)
				t.Error("y:", yI16x8)
				t.Error("s:", subi16x8s(xI16x8, yI16x8))
				t.Error(" :", subi16x8(xI16x8, yI16x8))
			}
			if muli16x8s(xI16x8, yI16x8) != muli16x8(xI16x8, yI16x8) {
				t.Errorf("muli16x8(%v, %v)", xI16x8, yI16x8)
				t.Error("x:", xI16x8)
				t.Error("y:", yI16x8)
				t.Error("s:", muli16x8s(xI16x8, yI16x8))
				t.Error(" :", muli16x8(xI16x8, yI16x8))
			}
			if shli16x8s(xI16x8, shift) != shli16x8(xI16x8, shift) {
				t.Errorf("shli16x8(%v, %v)", xI16x8, shift)
				t.Error("x:", xI16x8)
				t.Error("shift:", shift)
				t.Error("s:", shli16x8s(xI16x8, shift))
				t.Error(" :", shli16x8(xI16x8, shift))
			}
			if shri16x8s(xI16x8, shift) != shri16x8(xI16x8, shift) {
				t.Errorf("shri16x8(%v, %v)", xI16x8, shift)
				t.Error("x:", xI16x8)
				t.Error("shift:", shift)
				t.Error("s:", shri16x8s(xI16x8, shift))
				t.Error(" :", shri16x8(xI16x8, shift))
			}

			if addu16x8s(xU16x8, yU16x8) != addu16x8(xU16x8, yU16x8) {
				t.Errorf("addu16x8(%v, %v)", xU16x8, yU16x8)
			}
			if subu16x8s(xU16x8, yU16x8) != subu16x8(xU16x8, yU16x8) {
				t.Errorf("subu16x8(%v, %v)", xU16x8, yU16x8)
				t.Error("x:", xU16x8)
				t.Error("y:", yU16x8)
				t.Error("s:", subu16x8s(xU16x8, yU16x8))
				t.Error(" :", subu16x8(xU16x8, yU16x8))
			}
			if mulu16x8s(xU16x8, yU16x8) != mulu16x8(xU16x8, yU16x8) {
				t.Errorf("mulu16x8(%v, %v)", xU16x8, yU16x8)
				t.Error("x:", xU16x8)
				t.Error("y:", yU16x8)
				t.Error("s:", mulu16x8s(xU16x8, yU16x8))
				t.Error(" :", mulu16x8(xU16x8, yU16x8))
			}
			if shlu16x8s(xU16x8, shift) != shlu16x8(xU16x8, shift) {
				t.Errorf("shlu16x8(%v, %v)", xU16x8, shift)
				t.Error("x:", xU16x8)
				t.Error("shift:", shift)
				t.Error("s:", shlu16x8s(xU16x8, shift))
				t.Error(" :", shlu16x8(xU16x8, shift))
			}
			if shru16x8s(xU16x8, shift) != shru16x8(xU16x8, shift) {
				t.Errorf("shru16x8(%v, %v)", xU16x8, shift)
				t.Error("x:", xU16x8)
				t.Error("shift:", shift)
				t.Error("s:", shru16x8s(xU16x8, shift))
				t.Error(" :", shru16x8(xU16x8, shift))

			}

			if addi32x4s(xI32x4, yI32x4) != addi32x4(xI32x4, yI32x4) {
				t.Errorf("addi32x4(%v, %v)", xI32x4, yI32x4)
				t.Error("x:", xI32x4)
				t.Error("y:", yI32x4)
				t.Error("s(xI32x4, yI32x4):", addi32x4s(xI32x4, yI32x4))
				t.Error(" (xI32x4, yI32x4):", addi32x4(xI32x4, yI32x4))
			}
			if subi32x4s(xI32x4, yI32x4) != subi32x4(xI32x4, yI32x4) {
				t.Errorf("subi32x4(%v, %v)", xI32x4, yI32x4)
				t.Error("x:", xI32x4)
				t.Error("y:", yI32x4)
				t.Error("s:", subi32x4s(xI32x4, yI32x4))
				t.Error(" :", subi32x4(xI32x4, yI32x4))
			}
			if muli32x4s(xI32x4, yI32x4) != muli32x4(xI32x4, yI32x4) {
				t.Errorf("muli32x4(%v, %v)", xI32x4, yI32x4)
				t.Error("x:", xI32x4)
				t.Error("y:", yI32x4)
				t.Error("s:", muli32x4s(xI32x4, yI32x4))
				t.Error(" :", muli32x4(xI32x4, yI32x4))
				t.FailNow()
			}
			if shli32x4s(xI32x4, shift) != shli32x4(xI32x4, shift) {
				t.Errorf("shli32x4(%v, %v)", xI32x4, shift)
				t.Error("x:", xI32x4)
				t.Error("shift:", shift)
				t.Error("s:", shli32x4s(xI32x4, shift))
				t.Error(" :", shli32x4(xI32x4, shift))
			}
			if shri32x4s(xI32x4, shift) != shri32x4(xI32x4, shift) {
				t.Errorf("shri32x4(%v, %v)", xI32x4, shift)
				t.Error("x:", xI32x4)
				t.Error("shift:", shift)
				t.Error("s:", shri32x4s(xI32x4, shift))
				t.Error(" :", shri32x4(xI32x4, shift))
			}

			if addu32x4s(xU32x4, yU32x4) != addu32x4(xU32x4, yU32x4) {
				t.Errorf("addu32x4(%v, %v)", xU32x4, yU32x4)
			}
			if subu32x4s(xU32x4, yU32x4) != subu32x4(xU32x4, yU32x4) {
				t.Errorf("subu32x4(%v, %v)", xU32x4, yU32x4)
				t.Error("x:", xU32x4)
				t.Error("y:", yU32x4)
				t.Error("s:", subu32x4s(xU32x4, yU32x4))
				t.Error(" :", subu32x4(xU32x4, yU32x4))
			}
			if mulu32x4s(xU32x4, yU32x4) != mulu32x4(xU32x4, yU32x4) {
				t.Errorf("mulu32x4(%v, %v)", xU32x4, yU32x4)
				t.Error("x:", xU32x4)
				t.Error("y:", yU32x4)
				t.Error("s:", mulu32x4s(xU32x4, yU32x4))
				t.Error(" :", mulu32x4(xU32x4, yU32x4))
			}
			if shlu32x4s(xU32x4, shift) != shlu32x4(xU32x4, shift) {
				t.Errorf("shlu32x4(%v, %v)", xU32x4, shift)
				t.Error("x:", xU32x4)
				t.Error("shift:", shift)
				t.Error("s:", shlu32x4s(xU32x4, shift))
				t.Error(" :", shlu32x4(xU32x4, shift))
			}
			if shru32x4s(xU32x4, shift) != shru32x4(xU32x4, shift) {
				t.Errorf("shru32x4(%v, %v)", xU32x4, shift)
				t.Error("x:", xU32x4)
				t.Error("shift:", shift)
				t.Error("s:", shru32x4s(xU32x4, shift))
				t.Error(" :", shru32x4(xU32x4, shift))
			}

			if addi64x2s(xI64x2, yI64x2) != addi64x2(xI64x2, yI64x2) {
				t.Errorf("addi64x2(%v, %v)", xI64x2, yI64x2)
			}
			if subi64x2s(xI64x2, yI64x2) != subi64x2(xI64x2, yI64x2) {
				t.Errorf("subi64x2(%v, %v)", xI64x2, yI64x2)
			}

			if addu64x2s(xU64x2, yU64x2) != addu64x2(xU64x2, yU64x2) {
				t.Errorf("addu64x2(%v, %v)", xU64x2, yU64x2)
			}
			if subu64x2s(xU64x2, yU64x2) != subu64x2(xU64x2, yU64x2) {
				t.Errorf("subu64x2(%v, %v)", xU64x2, yU64x2)
			}

			if addf32x4s(xF32x4, yF32x4) != addf32x4(xF32x4, yF32x4) {
				t.Errorf("addf32x4(%v, %v)", xF32x4, yF32x4)
				t.Error("x:", xF32x4)
				t.Error("y:", yF32x4)
				t.Error("s:", addf32x4s(xF32x4, yF32x4))
				t.Error(" :", addf32x4(xF32x4, yF32x4))
				t.FailNow()

			}
			if subf32x4s(xF32x4, yF32x4) != subf32x4(xF32x4, yF32x4) {
				t.Errorf("subf32x4(%v, %v)", xF32x4, yF32x4)
				t.Error("x:", xF32x4)
				t.Error("y:", yF32x4)
				t.Error("s:", subf32x4s(xF32x4, yF32x4))
				t.Error(" :", subf32x4(xF32x4, yF32x4))
			}
			if mulf32x4s(xF32x4, yF32x4) != mulf32x4(xF32x4, yF32x4) {
				t.Errorf("mulf32x4(%v, %v)", xF32x4, yF32x4)
				t.Error("x:", xF32x4)
				t.Error("y:", yF32x4)
				t.Error("s:", mulf32x4s(xF32x4, yF32x4))
				t.Error(" :", mulf32x4(xF32x4, yF32x4))
			}
			if divf32x4s(xF32x4, yF32x4) != divf32x4(xF32x4, yF32x4) {
				t.Errorf("divf32x4(%v, %v)", xF32x4, yF32x4)
				t.Error("x:", xF32x4)
				t.Error("y:", yF32x4)
				t.Error("s:", divf32x4s(xF32x4, yF32x4))
				t.Error(" :", divf32x4(xF32x4, yF32x4))
			}

			if addf64x2s(xF64x2, yF64x2) != addf64x2(xF64x2, yF64x2) {
				t.Errorf("addf64x2(%v, %v)", xF64x2, yF64x2)
			}
			if subf64x2s(xF64x2, yF64x2) != subf64x2(xF64x2, yF64x2) {
				t.Errorf("subf64x2(%v, %v)", xF64x2, yF64x2)
				t.Error("x:", xF64x2)
				t.Error("y:", yF64x2)
				t.Error("s:", subf64x2s(xF64x2, yF64x2))
				t.Error(" :", subf64x2(xF64x2, yF64x2))
			}
			if mulf64x2s(xF64x2, yF64x2) != mulf64x2(xF64x2, yF64x2) {
				t.Errorf("mulf64x2(%v, %v)", xF64x2, yF64x2)
				t.Error("x:", xF64x2)
				t.Error("y:", yF64x2)
				t.Error("s:", mulf64x2s(xF64x2, yF64x2))
				t.Error(" :", mulf64x2(xF64x2, yF64x2))
			}
			if divf64x2s(xF64x2, yF64x2) != divf64x2(xF64x2, yF64x2) {
				t.Errorf("divf64x2(%v, %v)", xF64x2, yF64x2)
				t.Error("x:", xF64x2)
				t.Error("y:", yF64x2)
				t.Error("s:", divf64x2s(xF64x2, yF64x2))
				t.Error(" :", divf64x2(xF64x2, yF64x2))
			}

		}
	}

	t.Log("Test Count:", count)
}
