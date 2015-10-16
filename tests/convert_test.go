// +build amd64,gc

package tests

import (
	"math/rand"
	"testing"
)

//go:generate gensimd -fn "U8ToU8, U8ToU16, U8ToU32, U8ToU64, U8ToI8, U8ToI16, U8ToI32, U8ToI64, U8ToF32, U8ToF64, U16ToU8, U16ToU16, U16ToU32, U16ToU64, U16ToI8, U16ToI16, U16ToI32, U16ToI64, U16ToF32, U16ToF64, U32ToU8, U32ToU16, U32ToU32, U32ToU64, U32ToI8, U32ToI16, U32ToI32, U32ToI64, U32ToF32, U32ToF64, U64ToU8, U64ToU16, U64ToU32, U64ToU64, U64ToI8, U64ToI16, U64ToI32, U64ToI64, U64ToF32, U64ToF64, I8ToU8, I8ToU16, I8ToU32, I8ToU64, I8ToI8, I8ToI16, I8ToI32, I8ToI64, I8ToF32, I8ToF64, I16ToU8, I16ToU16, I16ToU32, I16ToU64, I16ToI8, I16ToI16, I16ToI32, I16ToI64, I16ToF32, I16ToF64, I32ToU8, I32ToU16, I32ToU32, I32ToU64, I32ToI8, I32ToI16, I32ToI32, I32ToI64, I32ToF32, I32ToF64, I64ToU8, I64ToU16, I64ToU32, I64ToU64, I64ToI8, I64ToI16, I64ToI32, I64ToI64, I64ToF32, I64ToF64, F32ToU8, F32ToU16, F32ToU32, F32ToU64, F32ToI8, F32ToI16, F32ToI32, F32ToI64, F32ToF32, F32ToF64, F64ToU8, F64ToU16, F64ToU32, F64ToU64, F64ToI8, F64ToI16, F64ToI32, F64ToI64, F64ToF32, F64ToF64" -outfn "U8ToU8s, U8ToU16s, U8ToU32s, U8ToU64s, U8ToI8s, U8ToI16s, U8ToI32s, U8ToI64s, U8ToF32s, U8ToF64s, U16ToU8s, U16ToU16s, U16ToU32s, U16ToU64s, U16ToI8s, U16ToI16s, U16ToI32s, U16ToI64s, U16ToF32s, U16ToF64s, U32ToU8s, U32ToU16s, U32ToU32s, U32ToU64s, U32ToI8s, U32ToI16s, U32ToI32s, U32ToI64s, U32ToF32s, U32ToF64s, U64ToU8s, U64ToU16s, U64ToU32s, U64ToU64s, U64ToI8s, U64ToI16s, U64ToI32s, U64ToI64s, U64ToF32s, U64ToF64s, I8ToU8s, I8ToU16s, I8ToU32s, I8ToU64s, I8ToI8s, I8ToI16s, I8ToI32s, I8ToI64s, I8ToF32s, I8ToF64s, I16ToU8s, I16ToU16s, I16ToU32s, I16ToU64s, I16ToI8s, I16ToI16s, I16ToI32s, I16ToI64s, I16ToF32s, I16ToF64s, I32ToU8s, I32ToU16s, I32ToU32s, I32ToU64s, I32ToI8s, I32ToI16s, I32ToI32s, I32ToI64s, I32ToF32s, I32ToF64s, I64ToU8s, I64ToU16s, I64ToU32s, I64ToU64s, I64ToI8s, I64ToI16s, I64ToI32s, I64ToI64s, I64ToF32s, I64ToF64s, F32ToU8s, F32ToU16s, F32ToU32s, F32ToU64s, F32ToI8s, F32ToI16s, F32ToI32s, F32ToI64s, F32ToF32s, F32ToF64s, F64ToU8s, F64ToU16s, F64ToU32s, F64ToU64s, F64ToI8s, F64ToI16s, F64ToI32s, F64ToI64s, F64ToF32s, F64ToF64s" -f "$GOFILE" -o "convert_test_amd64.s"

func U8ToU8s(uint8) uint8
func U8ToU16s(uint8) uint16
func U8ToU32s(uint8) uint32
func U8ToU64s(uint8) uint64
func U8ToI8s(uint8) int8
func U8ToI16s(uint8) int16
func U8ToI32s(uint8) int32
func U8ToI64s(uint8) int64
func U8ToF32s(uint8) float32
func U8ToF64s(uint8) float64

func U16ToU8s(uint16) uint8
func U16ToU16s(uint16) uint16
func U16ToU32s(uint16) uint32
func U16ToU64s(uint16) uint64
func U16ToI8s(uint16) int8
func U16ToI16s(uint16) int16
func U16ToI32s(uint16) int32
func U16ToI64s(uint16) int64
func U16ToF32s(uint16) float32
func U16ToF64s(uint16) float64

func U32ToU8s(uint32) uint8
func U32ToU16s(uint32) uint16
func U32ToU32s(uint32) uint32
func U32ToU64s(uint32) uint64
func U32ToI8s(uint32) int8
func U32ToI16s(uint32) int16
func U32ToI32s(uint32) int32
func U32ToI64s(uint32) int64
func U32ToF32s(uint32) float32
func U32ToF64s(uint32) float64

func U64ToU8s(uint64) uint8
func U64ToU16s(uint64) uint16
func U64ToU32s(uint64) uint32
func U64ToU64s(uint64) uint64
func U64ToI8s(uint64) int8
func U64ToI16s(uint64) int16
func U64ToI32s(uint64) int32
func U64ToI64s(uint64) int64
func U64ToF32s(uint64) float32
func U64ToF64s(uint64) float64

func I8ToU8s(int8) uint8
func I8ToU16s(int8) uint16
func I8ToU32s(int8) uint32
func I8ToU64s(int8) uint64
func I8ToI8s(int8) int8
func I8ToI16s(int8) int16
func I8ToI32s(int8) int32
func I8ToI64s(int8) int64
func I8ToF32s(int8) float32
func I8ToF64s(int8) float64

func I16ToU8s(int16) uint8
func I16ToU16s(int16) uint16
func I16ToU32s(int16) uint32
func I16ToU64s(int16) uint64
func I16ToI8s(int16) int8
func I16ToI16s(int16) int16
func I16ToI32s(int16) int32
func I16ToI64s(int16) int64
func I16ToF32s(int16) float32
func I16ToF64s(int16) float64

func I32ToU8s(int32) uint8
func I32ToU16s(int32) uint16
func I32ToU32s(int32) uint32
func I32ToU64s(int32) uint64
func I32ToI8s(int32) int8
func I32ToI16s(int32) int16
func I32ToI32s(int32) int32
func I32ToI64s(int32) int64
func I32ToF32s(int32) float32
func I32ToF64s(int32) float64

func I64ToU8s(int64) uint8
func I64ToU16s(int64) uint16
func I64ToU32s(int64) uint32
func I64ToU64s(int64) uint64
func I64ToI8s(int64) int8
func I64ToI16s(int64) int16
func I64ToI32s(int64) int32
func I64ToI64s(int64) int64
func I64ToF32s(int64) float32
func I64ToF64s(int64) float64

func F32ToU8s(float32) uint8
func F32ToU16s(float32) uint16
func F32ToU32s(float32) uint32
func F32ToU64s(float32) uint64
func F32ToI8s(float32) int8
func F32ToI16s(float32) int16
func F32ToI32s(float32) int32
func F32ToI64s(float32) int64
func F32ToF32s(float32) float32
func F32ToF64s(float32) float64

func F64ToU8s(float64) uint8
func F64ToU16s(float64) uint16
func F64ToU32s(float64) uint32
func F64ToU64s(float64) uint64
func F64ToI8s(float64) int8
func F64ToI16s(float64) int16
func F64ToI32s(float64) int32
func F64ToI64s(float64) int64
func F64ToF32s(float64) float32
func F64ToF64s(float64) float64

func U8ToU8(x uint8) uint8    { return uint8(x) }
func U8ToU16(x uint8) uint16  { return uint16(x) }
func U8ToU32(x uint8) uint32  { return uint32(x) }
func U8ToU64(x uint8) uint64  { return uint64(x) }
func U8ToI8(x uint8) int8     { return int8(x) }
func U8ToI16(x uint8) int16   { return int16(x) }
func U8ToI32(x uint8) int32   { return int32(x) }
func U8ToI64(x uint8) int64   { return int64(x) }
func U8ToF32(x uint8) float32 { return float32(x) }
func U8ToF64(x uint8) float64 { return float64(x) }

func U16ToU8(x uint16) uint8    { return uint8(x) }
func U16ToU16(x uint16) uint16  { return uint16(x) }
func U16ToU32(x uint16) uint32  { return uint32(x) }
func U16ToU64(x uint16) uint64  { return uint64(x) }
func U16ToI8(x uint16) int8     { return int8(x) }
func U16ToI16(x uint16) int16   { return int16(x) }
func U16ToI32(x uint16) int32   { return int32(x) }
func U16ToI64(x uint16) int64   { return int64(x) }
func U16ToF32(x uint16) float32 { return float32(x) }
func U16ToF64(x uint16) float64 { return float64(x) }

func U32ToU8(x uint32) uint8    { return uint8(x) }
func U32ToU16(x uint32) uint16  { return uint16(x) }
func U32ToU32(x uint32) uint32  { return uint32(x) }
func U32ToU64(x uint32) uint64  { return uint64(x) }
func U32ToI8(x uint32) int8     { return int8(x) }
func U32ToI16(x uint32) int16   { return int16(x) }
func U32ToI32(x uint32) int32   { return int32(x) }
func U32ToI64(x uint32) int64   { return int64(x) }
func U32ToF32(x uint32) float32 { return float32(x) }
func U32ToF64(x uint32) float64 { return float64(x) }

func U64ToU8(x uint64) uint8    { return uint8(x) }
func U64ToU16(x uint64) uint16  { return uint16(x) }
func U64ToU32(x uint64) uint32  { return uint32(x) }
func U64ToU64(x uint64) uint64  { return uint64(x) }
func U64ToI8(x uint64) int8     { return int8(x) }
func U64ToI16(x uint64) int16   { return int16(x) }
func U64ToI32(x uint64) int32   { return int32(x) }
func U64ToI64(x uint64) int64   { return int64(x) }
func U64ToF32(x uint64) float32 { return float32(x) }
func U64ToF64(x uint64) float64 { return float64(x) }

func I8ToU8(x int8) uint8    { return uint8(x) }
func I8ToU16(x int8) uint16  { return uint16(x) }
func I8ToU32(x int8) uint32  { return uint32(x) }
func I8ToU64(x int8) uint64  { return uint64(x) }
func I8ToI8(x int8) int8     { return int8(x) }
func I8ToI16(x int8) int16   { return int16(x) }
func I8ToI32(x int8) int32   { return int32(x) }
func I8ToI64(x int8) int64   { return int64(x) }
func I8ToF32(x int8) float32 { return float32(x) }
func I8ToF64(x int8) float64 { return float64(x) }

func I16ToU8(x int16) uint8    { return uint8(x) }
func I16ToU16(x int16) uint16  { return uint16(x) }
func I16ToU32(x int16) uint32  { return uint32(x) }
func I16ToU64(x int16) uint64  { return uint64(x) }
func I16ToI8(x int16) int8     { return int8(x) }
func I16ToI16(x int16) int16   { return int16(x) }
func I16ToI32(x int16) int32   { return int32(x) }
func I16ToI64(x int16) int64   { return int64(x) }
func I16ToF32(x int16) float32 { return float32(x) }
func I16ToF64(x int16) float64 { return float64(x) }

func I32ToU8(x int32) uint8    { return uint8(x) }
func I32ToU16(x int32) uint16  { return uint16(x) }
func I32ToU32(x int32) uint32  { return uint32(x) }
func I32ToU64(x int32) uint64  { return uint64(x) }
func I32ToI8(x int32) int8     { return int8(x) }
func I32ToI16(x int32) int16   { return int16(x) }
func I32ToI32(x int32) int32   { return int32(x) }
func I32ToI64(x int32) int64   { return int64(x) }
func I32ToF32(x int32) float32 { return float32(x) }
func I32ToF64(x int32) float64 { return float64(x) }

func I64ToU8(x int64) uint8    { return uint8(x) }
func I64ToU16(x int64) uint16  { return uint16(x) }
func I64ToU32(x int64) uint32  { return uint32(x) }
func I64ToU64(x int64) uint64  { return uint64(x) }
func I64ToI8(x int64) int8     { return int8(x) }
func I64ToI16(x int64) int16   { return int16(x) }
func I64ToI32(x int64) int32   { return int32(x) }
func I64ToI64(x int64) int64   { return int64(x) }
func I64ToF32(x int64) float32 { return float32(x) }
func I64ToF64(x int64) float64 { return float64(x) }

func F32ToU8(x float32) uint8    { return uint8(x) }
func F32ToU16(x float32) uint16  { return uint16(x) }
func F32ToU32(x float32) uint32  { return uint32(x) }
func F32ToU64(x float32) uint64  { return uint64(x) }
func F32ToI8(x float32) int8     { return int8(x) }
func F32ToI16(x float32) int16   { return int16(x) }
func F32ToI32(x float32) int32   { return int32(x) }
func F32ToI64(x float32) int64   { return int64(x) }
func F32ToF32(x float32) float32 { return float32(x) }
func F32ToF64(x float32) float64 { return float64(x) }

func F64ToU8(x float64) uint8    { return uint8(x) }
func F64ToU16(x float64) uint16  { return uint16(x) }
func F64ToU32(x float64) uint32  { return uint32(x) }
func F64ToU64(x float64) uint64  { return uint64(x) }
func F64ToI8(x float64) int8     { return int8(x) }
func F64ToI16(x float64) int16   { return int16(x) }
func F64ToI32(x float64) int32   { return int32(x) }
func F64ToI64(x float64) int64   { return int64(x) }
func F64ToF32(x float64) float32 { return float32(x) }
func F64ToF64(x float64) float64 { return float64(x) }

func TestNumericConversions(t *testing.T) {

	for i := 0; i <= 128*128; i++ {

		var f64 float64
		if i&1 == 1 {
			f64 = rand.Float64()
		} else {
			f64 = rand.ExpFloat64()
		}

		a := int64(0)
		j := i - 128*128/2
		if j < 0 {
			a = -1 << uint(-j)
		} else {
			a = 1<<uint(j) - 1
		}

		if U8ToU8s(uint8(a)) != U8ToU8(uint8(a)) {
			t.Errorf("U8ToU8s(%v) %v != %v", uint8(a), U8ToU8s(uint8(a)), U8ToU8(uint8(a)))
		}
		if U8ToU16s(uint8(a)) != U8ToU16(uint8(a)) {
			t.Errorf("U8ToU16s(%v) %v != %v", uint8(a), U8ToU16s(uint8(a)), U8ToU16(uint8(a)))
		}
		if U8ToU32s(uint8(a)) != U8ToU32(uint8(a)) {
			t.Errorf("U8ToU32s(%v) %v != %v", uint8(a), U8ToU32s(uint8(a)), U8ToU32(uint8(a)))
		}
		if U8ToU64s(uint8(a)) != U8ToU64(uint8(a)) {
			t.Errorf("U8ToU64s(%v) %v != %v", uint8(a), U8ToU64s(uint8(a)), U8ToU64(uint8(a)))
		}
		if U8ToI8s(uint8(a)) != U8ToI8(uint8(a)) {
			t.Errorf("U8ToI8s(%v) %v != %v", uint8(a), U8ToI8s(uint8(a)), U8ToI8(uint8(a)))
		}
		if U8ToI16s(uint8(a)) != U8ToI16(uint8(a)) {
			t.Errorf("U8ToI16s(%v) %v != %v", uint8(a), U8ToI16s(uint8(a)), U8ToI16(uint8(a)))
		}
		if U8ToI32s(uint8(a)) != U8ToI32(uint8(a)) {
			t.Errorf("U8ToI32s(%v) %v != %v", uint8(a), U8ToI32s(uint8(a)), U8ToI32(uint8(a)))
		}
		if U8ToI64s(uint8(a)) != U8ToI64(uint8(a)) {
			t.Errorf("U8ToI64s(%v) %v != %v", uint8(a), U8ToI64s(uint8(a)), U8ToI64(uint8(a)))
		}

		if U8ToF32s(uint8(a)) != U8ToF32(uint8(a)) {
			t.Errorf("U8ToF32s(%v) %v != %v", uint8(a), U8ToF32s(uint8(a)), U8ToF32(uint8(a)))
		}
		if U8ToF64s(uint8(a)) != U8ToF64(uint8(a)) {
			t.Errorf("U8ToF64s(%v) %v != %v", uint8(a), U8ToF64s(uint8(a)), U8ToF64(uint8(a)))
		}
		if U16ToU8s(uint16(a)) != U16ToU8(uint16(a)) {
			t.Errorf("U16ToU8s(%v) %v != %v", uint16(a), U16ToU8s(uint16(a)), U16ToU8(uint16(a)))
		}
		if U16ToU16s(uint16(a)) != U16ToU16(uint16(a)) {
			t.Errorf("U16ToU16s(%v) %v != %v", uint16(a), U16ToU16s(uint16(a)), U16ToU16(uint16(a)))
		}
		if U16ToU32s(uint16(a)) != U16ToU32(uint16(a)) {
			t.Errorf("U16ToU32s(%v) %v != %v", uint16(a), U16ToU32s(uint16(a)), U16ToU32(uint16(a)))
		}
		if U16ToU64s(uint16(a)) != U16ToU64(uint16(a)) {
			t.Errorf("U16ToU64s(%v) %v != %v", uint16(a), U16ToU64s(uint16(a)), U16ToU64(uint16(a)))
		}
		if U16ToI8s(uint16(a)) != U16ToI8(uint16(a)) {
			t.Errorf("U16ToI8s(%v) %v != %v", uint16(a), U16ToI8s(uint16(a)), U16ToI8(uint16(a)))
		}
		if U16ToI16s(uint16(a)) != U16ToI16(uint16(a)) {
			t.Errorf("U16ToI16s(%v) %v != %v", uint16(a), U16ToI16s(uint16(a)), U16ToI16(uint16(a)))
		}
		if U16ToI32s(uint16(a)) != U16ToI32(uint16(a)) {
			t.Errorf("U16ToI32s(%v) %v != %v", uint16(a), U16ToI32s(uint16(a)), U16ToI32(uint16(a)))
		}
		if U16ToI64s(uint16(a)) != U16ToI64(uint16(a)) {
			t.Errorf("U16ToI64s(%v) %v != %v", uint16(a), U16ToI64s(uint16(a)), U16ToI64(uint16(a)))
		}
		if U16ToF32s(uint16(a)) != U16ToF32(uint16(a)) {
			t.Errorf("U16ToF32s(%v) %v != %v", uint16(a), U16ToF32s(uint16(a)), U16ToF32(uint16(a)))
		}
		if U16ToF64s(uint16(a)) != U16ToF64(uint16(a)) {
			t.Errorf("U16ToF64s(%v) %v != %v", uint16(a), U16ToF64s(uint16(a)), U16ToF64(uint16(a)))
		}

		if U32ToU8s(uint32(a)) != U32ToU8(uint32(a)) {
			t.Errorf("U32ToU8s(%v) %v != %v", uint32(a), U32ToU8s(uint32(a)), U32ToU8(uint32(a)))
		}
		if U32ToU16s(uint32(a)) != U32ToU16(uint32(a)) {
			t.Errorf("U32ToU16s(%v) %v != %v", uint32(a), U32ToU16s(uint32(a)), U32ToU16(uint32(a)))
		}
		if U32ToU32s(uint32(a)) != U32ToU32(uint32(a)) {
			t.Errorf("U32ToU32s(%v) %v != %v", uint32(a), U32ToU32s(uint32(a)), U32ToU32(uint32(a)))
		}
		if U32ToU64s(uint32(a)) != U32ToU64(uint32(a)) {
			t.Errorf("U32ToU64s(%v) %v != %v", uint32(a), U32ToU64s(uint32(a)), U32ToU64(uint32(a)))
		}
		if U32ToI8s(uint32(a)) != U32ToI8(uint32(a)) {
			t.Errorf("U32ToI8s(%v) %v != %v", uint32(a), U32ToI8s(uint32(a)), U32ToI8(uint32(a)))
		}
		if U32ToI16s(uint32(a)) != U32ToI16(uint32(a)) {
			t.Errorf("U32ToI16s(%v) %v != %v", uint32(a), U32ToI16s(uint32(a)), U32ToI16(uint32(a)))
		}
		if U32ToI32s(uint32(a)) != U32ToI32(uint32(a)) {
			t.Errorf("U32ToI32s(%v) %v != %v", uint32(a), U32ToI32s(uint32(a)), U32ToI32(uint32(a)))
		}
		if U32ToI64s(uint32(a)) != U32ToI64(uint32(a)) {
			t.Errorf("U32ToI64s(%v) %v != %v", uint32(a), U32ToI64s(uint32(a)), U32ToI64(uint32(a)))
		}
		if U32ToF32s(uint32(a)) != U32ToF32(uint32(a)) {
			t.Errorf("U32ToF32s(%v) %v != %v", uint32(a), U32ToF32s(uint32(a)), U32ToF32(uint32(a)))
		}
		if U32ToF64s(uint32(a)) != U32ToF64(uint32(a)) {
			t.Errorf("U32ToF64s(%v) %v != %v", uint32(a), U32ToF64s(uint32(a)), U32ToF64(uint32(a)))
		}

		if U64ToU8s(uint64(a)) != U64ToU8(uint64(a)) {
			t.Errorf("U64ToU8s(%v) %v != %v", uint64(a), U64ToU8s(uint64(a)), U64ToU8(uint64(a)))
		}
		if U64ToU16s(uint64(a)) != U64ToU16(uint64(a)) {
			t.Errorf("U64ToU16s(%v) %v != %v", uint64(a), U64ToU16s(uint64(a)), U64ToU16(uint64(a)))
		}
		if U64ToU32s(uint64(a)) != U64ToU32(uint64(a)) {
			t.Errorf("U64ToU32s(%v) %v != %v", uint64(a), U64ToU32s(uint64(a)), U64ToU32(uint64(a)))
		}
		if U64ToU64s(uint64(a)) != U64ToU64(uint64(a)) {
			t.Errorf("U64ToU64s(%v) %v != %v", uint64(a), U64ToU64s(uint64(a)), U64ToU64(uint64(a)))
		}
		if U64ToI8s(uint64(a)) != U64ToI8(uint64(a)) {
			t.Errorf("U64ToI8s(%v) %v != %v", uint64(a), U64ToI8s(uint64(a)), U64ToI8(uint64(a)))
		}
		if U64ToI16s(uint64(a)) != U64ToI16(uint64(a)) {
			t.Errorf("U64ToI16s(%v) %v != %v", uint64(a), U64ToI16s(uint64(a)), U64ToI16(uint64(a)))
		}
		if U64ToI32s(uint64(a)) != U64ToI32(uint64(a)) {
			t.Errorf("U64ToI32s(%v) %v != %v", uint64(a), U64ToI32s(uint64(a)), U64ToI32(uint64(a)))
		}
		if U64ToI64s(uint64(a)) != U64ToI64(uint64(a)) {
			t.Errorf("U64ToI64s(%v) %v != %v", uint64(a), U64ToI64s(uint64(a)), U64ToI64(uint64(a)))
		}

		if U64ToF32s(uint64(a)) != U64ToF32(uint64(a)) {
			t.Errorf("U64ToF32s(%v) %v != %v", uint64(a), U64ToF32s(uint64(a)), U64ToF32(uint64(a)))
		}
		if U64ToF64s(uint64(a)) != U64ToF64(uint64(a)) {
			t.Errorf("U64ToF64s(%v) %v != %v", uint64(a), U64ToF64s(uint64(a)), U64ToF64(uint64(a)))
		}

		if I8ToU8s(int8(a)) != I8ToU8(int8(a)) {
			t.Errorf("I8ToU8s(%v) %v != %v", int8(a), I8ToU8s(int8(a)), I8ToU8(int8(a)))
		}
		if I8ToU16s(int8(a)) != I8ToU16(int8(a)) {
			t.Errorf("I8ToU16s(%v) %v != %v", int8(a), I8ToU16s(int8(a)), I8ToU16(int8(a)))
		}
		if I8ToU32s(int8(a)) != I8ToU32(int8(a)) {
			t.Errorf("I8ToU32s(%v) %v != %v", int8(a), I8ToU32s(int8(a)), I8ToU32(int8(a)))
		}
		if I8ToU64s(int8(a)) != I8ToU64(int8(a)) {
			t.Errorf("I8ToU64s(%v) %v != %v", int8(a), I8ToU64s(int8(a)), I8ToU64(int8(a)))
		}
		if I8ToI8s(int8(a)) != I8ToI8(int8(a)) {
			t.Errorf("I8ToI8s(%v) %v != %v", int8(a), I8ToI8s(int8(a)), I8ToI8(int8(a)))
		}
		if I8ToI16s(int8(a)) != I8ToI16(int8(a)) {
			t.Errorf("I8ToI16s(%v) %v != %v", int8(a), I8ToI16s(int8(a)), I8ToI16(int8(a)))
		}
		if I8ToI32s(int8(a)) != I8ToI32(int8(a)) {
			t.Errorf("I8ToI32s(%v) %v != %v", int8(a), I8ToI32s(int8(a)), I8ToI32(int8(a)))
		}
		if I8ToI64s(int8(a)) != I8ToI64(int8(a)) {
			t.Errorf("I8ToI64s(%v) %v != %v", int8(a), I8ToI64s(int8(a)), I8ToI64(int8(a)))
		}
		if I8ToF32s(int8(a)) != I8ToF32(int8(a)) {
			t.Errorf("I8ToF32s(%v) %v != %v", int8(a), I8ToF32s(int8(a)), I8ToF32(int8(a)))
		}
		if I8ToF64s(int8(a)) != I8ToF64(int8(a)) {
			t.Errorf("I8ToF64s(%v) %v != %v", int8(a), I8ToF64s(int8(a)), I8ToF64(int8(a)))
		}

		if I16ToU8s(int16(a)) != I16ToU8(int16(a)) {
			t.Errorf("I16ToU8s(%v) %v != %v", int16(a), I16ToU8s(int16(a)), I16ToU8(int16(a)))
		}
		if I16ToU16s(int16(a)) != I16ToU16(int16(a)) {
			t.Errorf("I16ToU16s(%v) %v != %v", int16(a), I16ToU16s(int16(a)), I16ToU16(int16(a)))
		}
		if I16ToU32s(int16(a)) != I16ToU32(int16(a)) {
			t.Errorf("I16ToU32s(%v) %v != %v", int16(a), I16ToU32s(int16(a)), I16ToU32(int16(a)))
		}
		if I16ToU64s(int16(a)) != I16ToU64(int16(a)) {
			t.Errorf("I16ToU64s(%v) %v != %v", int16(a), I16ToU64s(int16(a)), I16ToU64(int16(a)))
		}
		if I16ToI8s(int16(a)) != I16ToI8(int16(a)) {
			t.Errorf("I16ToI8s(%v) %v != %v", int16(a), I16ToI8s(int16(a)), I16ToI8(int16(a)))
		}
		if I16ToI16s(int16(a)) != I16ToI16(int16(a)) {
			t.Errorf("I16ToI16s(%v) %v != %v", int16(a), I16ToI16s(int16(a)), I16ToI16(int16(a)))
		}
		if I16ToI32s(int16(a)) != I16ToI32(int16(a)) {
			t.Errorf("I16ToI32s(%v) %v != %v", int16(a), I16ToI32s(int16(a)), I16ToI32(int16(a)))
		}
		if I16ToI64s(int16(a)) != I16ToI64(int16(a)) {
			t.Errorf("I16ToI64s(%v) %v != %v", int16(a), I16ToI64s(int16(a)), I16ToI64(int16(a)))
		}
		if I16ToF32s(int16(a)) != I16ToF32(int16(a)) {
			t.Errorf("I16ToF32s(%v) %v != %v", int16(a), I16ToF32s(int16(a)), I16ToF32(int16(a)))
		}
		if I16ToF64s(int16(a)) != I16ToF64(int16(a)) {
			t.Errorf("I16ToF64s(%v) %v != %v", int16(a), I16ToF64s(int16(a)), I16ToF64(int16(a)))
		}

		if I32ToU8s(int32(a)) != I32ToU8(int32(a)) {
			t.Errorf("I32ToU8s(%v) %v != %v", int32(a), I32ToU8s(int32(a)), I32ToU8(int32(a)))
		}
		if I32ToU16s(int32(a)) != I32ToU16(int32(a)) {
			t.Errorf("I32ToU16s(%v) %v != %v", int32(a), I32ToU16s(int32(a)), I32ToU16(int32(a)))
		}
		if I32ToU32s(int32(a)) != I32ToU32(int32(a)) {
			t.Errorf("I32ToU32s(%v) %v != %v", int32(a), I32ToU32s(int32(a)), I32ToU32(int32(a)))
		}
		if I32ToU64s(int32(a)) != I32ToU64(int32(a)) {
			t.Errorf("I32ToU64s(%v) %v != %v", int32(a), I32ToU64s(int32(a)), I32ToU64(int32(a)))
		}
		if I32ToI8s(int32(a)) != I32ToI8(int32(a)) {
			t.Errorf("I32ToI8s(%v) %v != %v", int32(a), I32ToI8s(int32(a)), I32ToI8(int32(a)))
		}
		if I32ToI16s(int32(a)) != I32ToI16(int32(a)) {
			t.Errorf("I32ToI16s(%v) %v != %v", int32(a), I32ToI16s(int32(a)), I32ToI16(int32(a)))
		}
		if I32ToI32s(int32(a)) != I32ToI32(int32(a)) {
			t.Errorf("I32ToI32s(%v) %v != %v", int32(a), I32ToI32s(int32(a)), I32ToI32(int32(a)))
		}
		if I32ToI64s(int32(a)) != I32ToI64(int32(a)) {
			t.Errorf("I32ToI64s(%v) %v != %v", int32(a), I32ToI64s(int32(a)), I32ToI64(int32(a)))
		}
		if I32ToF32s(int32(a)) != I32ToF32(int32(a)) {
			t.Errorf("I32ToF32s(%v) %v != %v", int32(a), I32ToF32s(int32(a)), I32ToF32(int32(a)))
		}
		if I32ToF64s(int32(a)) != I32ToF64(int32(a)) {
			t.Errorf("I32ToF64s(%v) %v != %v", int32(a), I32ToF64s(int32(a)), I32ToF64(int32(a)))
		}

		if I64ToU8s(int64(a)) != I64ToU8(int64(a)) {
			t.Errorf("I64ToU8s(%v) %v != %v", int64(a), I64ToU8s(int64(a)), I64ToU8(int64(a)))
		}
		if I64ToU16s(int64(a)) != I64ToU16(int64(a)) {
			t.Errorf("I64ToU16s(%v) %v != %v", int64(a), I64ToU16s(int64(a)), I64ToU16(int64(a)))
		}
		if I64ToU32s(int64(a)) != I64ToU32(int64(a)) {
			t.Errorf("I64ToU32s(%v) %v != %v", int64(a), I64ToU32s(int64(a)), I64ToU32(int64(a)))
		}
		if I64ToU64s(int64(a)) != I64ToU64(int64(a)) {
			t.Errorf("I64ToU64s(%v) %v != %v", int64(a), I64ToU64s(int64(a)), I64ToU64(int64(a)))
		}
		if I64ToI8s(int64(a)) != I64ToI8(int64(a)) {
			t.Errorf("I64ToI8s(%v) %v != %v", int64(a), I64ToI8s(int64(a)), I64ToI8(int64(a)))
		}
		if I64ToI16s(int64(a)) != I64ToI16(int64(a)) {
			t.Errorf("I64ToI16s(%v) %v != %v", int64(a), I64ToI16s(int64(a)), I64ToI16(int64(a)))
		}
		if I64ToI32s(int64(a)) != I64ToI32(int64(a)) {
			t.Errorf("I64ToI32s(%v) %v != %v", int64(a), I64ToI32s(int64(a)), I64ToI32(int64(a)))
		}
		if I64ToI64s(int64(a)) != I64ToI64(int64(a)) {
			t.Errorf("I64ToI64s(%v) %v != %v", int64(a), I64ToI64s(int64(a)), I64ToI64(int64(a)))
		}
		if I64ToF32s(int64(a)) != I64ToF32(int64(a)) {
			t.Errorf("I64ToF32s(%v) %v != %v", int64(a), I64ToF32s(int64(a)), I64ToF32(int64(a)))
		}
		if I64ToF64s(int64(a)) != I64ToF64(int64(a)) {
			t.Errorf("I64ToF64s(%v) %v != %v", int64(a), I64ToF64s(int64(a)), I64ToF64(int64(a)))
		}

		if F32ToU8s(float32(f64)) != F32ToU8(float32(f64)) {
			t.Errorf("F32ToU8s(%v) %v != %v", float32(f64), F32ToU8s(float32(f64)), F32ToU8(float32(f64)))
		}
		if F32ToU16s(float32(f64)) != F32ToU16(float32(f64)) {
			t.Errorf("F32ToU16s(%v) %v != %v", float32(f64), F32ToU16s(float32(f64)), F32ToU16(float32(f64)))
		}
		if F32ToU32s(float32(f64)) != F32ToU32(float32(f64)) {
			t.Errorf("F32ToU32s(%v) %v != %v", float32(f64), F32ToU32s(float32(f64)), F32ToU32(float32(f64)))
		}
		if F32ToU64s(float32(f64)) != F32ToU64(float32(f64)) {
			t.Errorf("F32ToU64s(%v) %v != %v", float32(f64), F32ToU64s(float32(f64)), F32ToU64(float32(f64)))
		}
		if F32ToI8s(float32(f64)) != F32ToI8(float32(f64)) {
			t.Errorf("F32ToI8s(%v) %v != %v", float32(f64), F32ToI8s(float32(f64)), F32ToI8(float32(f64)))
		}
		if F32ToI16s(float32(f64)) != F32ToI16(float32(f64)) {
			t.Errorf("F32ToI16s(%v) %v != %v", float32(f64), F32ToI16s(float32(f64)), F32ToI16(float32(f64)))
		}
		if F32ToI32s(float32(f64)) != F32ToI32(float32(f64)) {
			t.Errorf("F32ToI32s(%v) %v != %v", float32(f64), F32ToI32s(float32(f64)), F32ToI32(float32(f64)))
		}
		if F32ToI64s(float32(f64)) != F32ToI64(float32(f64)) {
			t.Errorf("F32ToI64s(%v) %v != %v", float32(f64), F32ToI64s(float32(f64)), F32ToI64(float32(f64)))
		}
		if F32ToF32s(float32(f64)) != F32ToF32(float32(f64)) {
			t.Errorf("F32ToF32s(%v) %v != %v", float32(f64), F32ToF32s(float32(f64)), F32ToF32(float32(f64)))
		}
		if F32ToF64s(float32(f64)) != F32ToF64(float32(f64)) {
			t.Errorf("F32ToF64s(%v) %v != %v", float32(f64), F32ToF64s(float32(f64)), F32ToF64(float32(f64)))
		}

		if F64ToU8s(float64(f64)) != F64ToU8(float64(f64)) {
			t.Errorf("F64ToU8s(%v) %v != %v", float64(f64), F64ToU8s(float64(f64)), F64ToU8(float64(f64)))
		}
		if F64ToU16s(float64(f64)) != F64ToU16(float64(f64)) {
			t.Errorf("F64ToU16s(%v) %v != %v", float64(f64), F64ToU16s(float64(f64)), F64ToU16(float64(f64)))
		}
		if F64ToU32s(float64(f64)) != F64ToU32(float64(f64)) {
			t.Errorf("F64ToU32s(%v) %v != %v", float64(f64), F64ToU32s(float64(f64)), F64ToU32(float64(f64)))
		}
		if F64ToU64s(float64(f64)) != F64ToU64(float64(f64)) {
			t.Errorf("F64ToU64s(%v) %v != %v", float64(f64), F64ToU64s(float64(f64)), F64ToU64(float64(f64)))
		}
		if F64ToI8s(float64(f64)) != F64ToI8(float64(f64)) {
			t.Errorf("F64ToI8s(%v) %v != %v", float64(f64), F64ToI8s(float64(f64)), F64ToI8(float64(f64)))
		}
		if F64ToI16s(float64(f64)) != F64ToI16(float64(f64)) {
			t.Errorf("F64ToI16s(%v) %v != %v", float64(f64), F64ToI16s(float64(f64)), F64ToI16(float64(f64)))
		}
		if F64ToI32s(float64(f64)) != F64ToI32(float64(f64)) {
			t.Errorf("F64ToI32s(%v) %v != %v", float64(f64), F64ToI32s(float64(f64)), F64ToI32(float64(f64)))
		}
		if F64ToI64s(float64(f64)) != F64ToI64(float64(f64)) {
			t.Errorf("F64ToI64s(%v) %v != %v", float64(f64), F64ToI64s(float64(f64)), F64ToI64(float64(f64)))
		}
		if F64ToF32s(float64(f64)) != F64ToF32(float64(f64)) {
			t.Errorf("F64ToF32s(%v) %v != %v", float64(f64), F64ToF32s(float64(f64)), F64ToF32(float64(f64)))
		}
		if F64ToF64s(float64(f64)) != F64ToF64(float64(f64)) {
			t.Errorf("F64ToF64s(%v) %v != %v", float64(f64), F64ToF64s(float64(f64)), F64ToF64(float64(f64)))
		}

	}

}
