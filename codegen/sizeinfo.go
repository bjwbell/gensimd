package codegen

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bjwbell/gensimd/simd"

	"golang.org/x/tools/go/types"
)

type simdtype struct {
	name     string
	size     uint
	elemSize uint
	align    uint
	optype   OpDataType
}

func simdReflect(t reflect.Type) simdtype {
	elemSize := uint(0)
	if t.Kind() == reflect.Array {
		elemSize = uint(t.Elem().Size())
	}
	return simdtype{
		name:     t.Name(),
		size:     uint(t.Size()),
		elemSize: elemSize,
		align:    uint(t.Size()),
	}
}

func simdTypes() []simdtype {
	types := []simdtype{}
	types = append(types, simdReflect(reflect.TypeOf(simd.I8x16{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_I8X16}

	types = append(types, simdReflect(reflect.TypeOf(simd.I16x8{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_I16X8}

	types = append(types, simdReflect(reflect.TypeOf(simd.I32x4{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_I32X4}

	types = append(types, simdReflect(reflect.TypeOf(simd.I64x2{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_I64X2}

	types = append(types, simdReflect(reflect.TypeOf(simd.U8x16{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_U8X16}

	types = append(types, simdReflect(reflect.TypeOf(simd.U16x8{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_U16X8}

	types = append(types, simdReflect(reflect.TypeOf(simd.U32x4{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_U32X4}

	types = append(types, simdReflect(reflect.TypeOf(simd.U64x2{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_U64X2}

	types = append(types, simdReflect(reflect.TypeOf(simd.F32x4{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_4X_F32}

	types = append(types, simdReflect(reflect.TypeOf(simd.F64x2{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_2X_F64}

	return types
}

type sse2type struct {
	name     string
	size     uint
	elemSize uint
	align    uint
	optype   OpDataType
}

func sse2Reflect(t reflect.Type) sse2type {
	elemSize := uint(0)
	if t.Kind() == reflect.Array {
		elemSize = uint(t.Elem().Size())
	}
	return sse2type{
		name:     t.Name(),
		size:     uint(t.Size()),
		elemSize: elemSize,
		align:    uint(t.Size()),
	}
}

func sse2Types() []sse2type {
	types := []sse2type{}
	types = append(types, sse2Reflect(reflect.TypeOf(simd.M128{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_M128}

	types = append(types, sse2Reflect(reflect.TypeOf(simd.M128i{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_M128i}

	types = append(types, sse2Reflect(reflect.TypeOf(simd.M128d{})))
	types[len(types)-1].optype = OpDataType{XMM_OP, InstrData{16, false}, XMM_M128d}

	return types
}

func simdInfo(t types.Type) (simdtype, bool) {
	named, ok := t.(*types.Named)
	if !ok {
		return simdtype{}, false
	}
	tname := named.Obj()
	for _, simdType := range simdTypes() {
		if tname.Name() == simdType.name {
			return simdType, true
		}
	}
	return simdtype{}, false
}

func isSimd(t types.Type) bool {
	_, ok := simdInfo(t)
	return ok
}

func isIntegerSimd(t types.Type) bool {
	s, ok := simdInfo(t)
	if !ok {
		return false
	}
	split := strings.Split(s.name, ".")
	name := split[len(split)-1]
	integerType := strings.HasPrefix(name, "I") || strings.HasPrefix(name, "U")
	return integerType
}

func sse2Info(t types.Type) (sse2type, bool) {
	named, ok := t.(*types.Named)
	if !ok {
		return sse2type{}, false
	}
	tname := named.Obj()
	for _, simdType := range sse2Types() {
		if tname.Name() == simdType.name {
			return simdType, true
		}
	}
	return sse2type{}, false
}

func isSSE2(t types.Type) bool {
	_, ok := sse2Info(t)
	return ok
}

func isIntegerSSE2(t types.Type) bool {
	s, ok := sse2Info(t)
	if !ok {
		return false
	}
	return s.optype.xmmvariant == XMM_M128i
}

func sizeofElem(t types.Type) uint {
	var e types.Type
	switch t := t.(type) {
	default:
		panic(internal(fmt.Sprintf("type (%v) not an array or slice\n", t.String())))
	case *types.Slice:
		e = t.Elem()
	case *types.Array:
		e = t.Elem()
	case *types.Named:
		typeinfo, ok := simdInfo(t)
		if ok {
			return typeinfo.elemSize
		}
		panic(internal(
			fmt.Sprintf("t (%v), isSimd (%v)\n", t.String(), isSimd(t))))

	}
	return sizeof(e)
}

func sizeof(t types.Type) uint {

	switch t := t.(type) {
	case *types.Tuple:
		// TODO: usage of reflect most likely wrong!
		// uint(reflect.TypeOf(t).Elem().Size())
		panic("Tuples are unsupported")
	case *types.Basic:
		return sizeBasic(t.Kind())
	case *types.Pointer:
		return sizePtr()
	case *types.Slice:
		return sizeSlice(t)
	case *types.Array:
		return sizeArray(t)
	case *types.Named:
		if sse2, ok := sse2Info(t); ok {
			return sse2.size
		} else if info, ok := simdInfo(t); ok {
			return info.size
		} else {
			panic(internal(fmt.Sprintf("unknown named type \"%v\"", t.String())))
		}
	}
	panic(internal(fmt.Sprintf("unknown type: %v", t)))
}

func sizeArray(t *types.Array) uint {
	return uint(reflectType(t).Size())
}

func sizeSlice(t *types.Slice) uint {
	return uint(reflectType(t).Size())
}

func sizeInt() uint {
	return sizeBasic(types.Int)
}

func sizePtr() uint {
	typ := reflect.TypeOf(true)
	ptrType := reflect.PtrTo(typ)
	size := ptrType.Size()
	return uint(size)
}

// sizeBasic return the size in bytes of a basic type
func sizeBasic(b types.BasicKind) uint {
	return uint(reflectBasic(b).Size())
}

func align(t types.Type) uint {

	switch t := t.(type) {
	case *types.Tuple:
		return alignTuple(t)
	case *types.Array, *types.Basic, *types.Pointer, *types.Slice:
		return uint(reflectType(t).Align())
	case *types.Named:
		if sse2, ok := sse2Info(t); ok {
			return sse2.align
		} else if info, ok := simdInfo(t); ok {
			return info.align
		} else {
			panic(internal(fmt.Sprintf("unknown named type \"%v\"", t.String())))
		}
	}
	panic(internal(fmt.Sprintf("unknown type (%v)", t)))
}

const tupleAlignment = 8

func alignTuple(tup *types.Tuple) uint {
	return tupleAlignment
}

func signed(t types.Type) bool {

	switch t := t.(type) {
	case *types.Basic:
		return signedBasic(t.Kind())
	}
	panic(internal(fmt.Sprintf("unknown type: %v", t)))
}

func signedBasic(b types.BasicKind) bool {
	switch b {
	case types.Bool:
		return false
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
		return true
	case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
		return false
	case types.Float32, types.Float64:
		return true
	}
	panic(internal(fmt.Sprintf("unknown basic type (%v)", b)))
}

func isUint(t types.Type) bool {
	if t, ok := t.(*types.Basic); ok {
		switch t.Kind() {
		case types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64:
			return true
		}
	}
	return false
}
func isInt(t types.Type) bool {
	if t, ok := t.(*types.Basic); ok {
		switch t.Kind() {
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64:
			return true
		}
	}
	return false
}
func isInteger(t types.Type) bool {
	return isUint(t) || isInt(t)
}
func isBool(t types.Type) bool {
	return isBasicKind(t, types.Bool)
}

func isFloat(t types.Type) bool {
	return isFloat32(t) || isFloat64(t)
}

func isFloat32(t types.Type) bool {
	return isBasicKind(t, types.Float32)
}

func isFloat64(t types.Type) bool {
	return isBasicKind(t, types.Float64)
}

func isComplex(t types.Type) bool {
	return isBasicKind(t, types.Complex64) || isBasicKind(t, types.Complex128)
}

func isBasicKind(t types.Type, basickind types.BasicKind) bool {
	if t, ok := t.(*types.Basic); ok {
		return t.Kind() == basickind
	}
	return false
}

func isBasic(t types.Type) bool {
	_, ok := t.(*types.Basic)
	return ok
}

func isArray(t types.Type) bool {
	_, ok := t.(*types.Array)
	return ok
}

func isSlice(t types.Type) bool {
	_, ok := t.(*types.Slice)
	return ok
}

func sliceLenSize() uint {
	return sizeInt()
}

func sliceLenOffset() int {
	return int(sizePtr())
}

func reflectType(t types.Type) reflect.Type {
	switch t := t.(type) {
	case *types.Tuple:
		// TODO
	case *types.Basic:
		return reflectBasic(t.Kind())
	case *types.Pointer:
		return reflect.PtrTo(reflectType(t.Elem()))
	case *types.Slice:
		return reflect.SliceOf(reflectType(t.Elem()))
	case *types.Array:
		return reflect.ArrayOf(int(t.Len()), reflectType(t.Elem()))
	case *types.Named:
		// TODO
	}
	internal(fmt.Sprintf("error unknown type:\"%v\"", t))
	panic("")
}

func reflectBasic(b types.BasicKind) reflect.Type {
	switch b {
	default:
		panic(internal("unknown basic type"))
	case types.Bool:
		return reflect.TypeOf(true)
	case types.Int:
		return reflect.TypeOf(int(1))
	case types.Int8:
		return reflect.TypeOf(int8(1))
	case types.Int16:
		return reflect.TypeOf(int16(1))
	case types.Int32:
		return reflect.TypeOf(int32(1))
	case types.Int64:
		return reflect.TypeOf(int64(1))
	case types.Uint:
		return reflect.TypeOf(uint(1))
	case types.Uint8:
		return reflect.TypeOf(uint8(1))
	case types.Uint16:
		return reflect.TypeOf(uint16(1))
	case types.Uint32:
		return reflect.TypeOf(uint32(1))
	case types.Uint64:
		return reflect.TypeOf(uint64(1))
	case types.Float32:
		return reflect.TypeOf(float32(1))
	case types.Float64:
		return reflect.TypeOf(float64(1))
	}
}

var bInstrData = OpDataType{INTEGER_OP, InstrData{signed: false, size: 1}, XMM_INVALID}
var f32InstrData = OpDataType{XMM_OP, InstrData{16, false}, XMM_F32}
var f64InstrData = OpDataType{XMM_OP, InstrData{16, false}, XMM_F64}

func GetIntegerOpDataType(signed bool, size uint) OpDataType {
	instrdata := OpDataType{
		INTEGER_OP,
		InstrData{signed: signed, size: size},
		XMM_INVALID}
	return instrdata

}

func GetOpDataType(t types.Type) OpDataType {
	if isBool(t) {
		return bInstrData
	}
	if isFloat32(t) {
		return f32InstrData
	} else if isFloat64(t) {
		return f64InstrData
	}
	if isComplex(t) {
		panic("complex32/64 unsupported")
	}
	if simdtype, ok := simdInfo(t); ok {
		return simdtype.optype
	}
	if sse2type, ok := sse2Info(t); ok {
		return sse2type.optype
	}
	if isBasic(t) {
		return GetIntegerOpDataType(signed(t), sizeof(t))
	} else {
		panic(internal(fmt.Sprintf("non basic type \"%v\"", t)))
	}

}

func regType(t types.Type) RegType {
	if isSimd(t) || isSSE2(t) {
		return XMM_REG
	}
	if isFloat(t) {
		return XMM_REG
	}
	return DATA_REG
}
