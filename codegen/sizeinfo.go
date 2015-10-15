package codegen

import (
	"fmt"
	"reflect"

	"golang.org/x/tools/go/types"

	"github.com/bjwbell/gensimd/simd"
)

var pointerSize = uint(8)
var sliceSize = uint(24)

type simdInfo struct {
	name     string
	size     uint
	elemSize uint
}

func simdReflect(t reflect.Type) simdInfo {
	elemSize := uint(0)
	if t.Kind() == reflect.Array {
		elemSize = uint(t.Elem().Size())
	}
	return simdInfo{t.Name(), uint(t.Size()), elemSize}
}

func simdTypes() []simdInfo {
	simdInt := reflect.TypeOf(simd.Int(0))
	simdInt4 := reflect.TypeOf(simd.Int4{})
	return []simdInfo{simdReflect(simdInt), simdReflect(simdInt4)}
}

func isSimd(t types.Type) bool {
	if t, ok := t.(*types.Named); ok {
		tname := t.Obj()
		for _, simdType := range simdTypes() {
			if tname.Name() == simdType.name {
				return true
			}
		}
	}
	return false
}

func simdTypeInfo(t types.Type) (simdInfo, error) {
	if !isSimd(t) {
		msg := fmt.Errorf("type (%v) is not simd type", t.String())
		return simdInfo{}, msg
	}
	named := t.(*types.Named)
	tname := named.Obj()
	for _, simdType := range simdTypes() {
		if tname.Name() == simdType.name {
			return simdType, nil
		}
	}
	msg := fmt.Errorf("type (%v) couldn't find simd type info", t.String())
	return simdInfo{}, msg
}

func simdHasElemSize(t types.Type) bool {
	if simdInfo, err := simdTypeInfo(t); err == nil {
		return simdInfo.elemSize > 0
	} else {
		msg := fmt.Sprintf("Error in simdHasElemSize, type (%v) is not simd", t.String())
		panic(msg)
	}
}

func simdElemSize(t types.Type) uint {
	if simdInfo, err := simdTypeInfo(t); err == nil {
		return simdInfo.elemSize
	} else {
		msg := fmt.Sprintf("Error in simdElemSize, type (%v) is not simd", t.String())
		panic(msg)
	}
}

func sizeofElem(t types.Type) uint {
	var e types.Type
	switch t := t.(type) {
	default:
		panic(fmt.Sprintf("t (%v) not an array or slice type\n", t.String()))
	case *types.Slice:
		e = t.Elem()
	case *types.Array:
		e = t.Elem()
	case *types.Named:
		if isSimd(t) && simdHasElemSize(t) {
			return simdElemSize(t)
		}
		panic(fmt.Sprintf("t (%v), isSimd (%v)\n", t.String(), isSimd(t)))
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
		return sliceSize
	case *types.Array:
		// TODO: calculation most likely wrong
		// return uint(t.Len()) * sizeof(t.Elem())
		panic("Arrays are unsupported")
	case *types.Named:
		if !isSimd(t) {
			panic("Named type is unsupported")
		}
		if info, err := simdTypeInfo(t); err != nil {
			panic(fmt.Sprintf("Error unknown type in sizeof err:\"%v\"", err))
		} else {
			return info.size
		}
	}
	panic(fmt.Sprintf("Error unknown type: %v", t))
}

func sizeInt() uint {
	return sizeBasic(types.Int)
}

func sizePtr() uint {
	return pointerSize
}

// sizeBasic return the size in bytes of a basic type
func sizeBasic(b types.BasicKind) uint {
	return uint(reflectBasic(b).Size())
}

func align(t types.Type) uint {

	switch t := t.(type) {
	case *types.Tuple:
		// TODO: usage of reflect most likely wrong!
		return uint(reflect.TypeOf(t).Elem().Size())
	case *types.Basic:
		return alignBasic(t.Kind())
	case *types.Pointer:
		return alignPtr()
	case *types.Slice:
		return alignSlice()
	case *types.Array:
		// TODO: most likely wrong
		return alignSlice()
	case *types.Named:
		panic(fmt.Sprintf("Error unknown named type in align:\"%v\"", t))
	}
	panic(fmt.Sprintf("Error unknown type (%v)", t))
}

func alignPtr() uint {
	return 8
}
func alignSlice() uint {
	return 8
}

func alignBasic(b types.BasicKind) uint {
	return uint(reflectBasic(b).Align())
}

func signed(t types.Type) bool {

	switch t := t.(type) {
	case *types.Basic:
		return signedBasic(t.Kind())
	}
	panic(fmt.Sprintf("Error unknown type: %v", t))
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
	panic(fmt.Sprintf("Unknown basic type (%v)", b))
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
func reflectType(t types.Type) reflect.Type {
	switch t := t.(type) {
	case *types.Tuple:
		// TODO
	case *types.Basic:
		return reflectBasic(t.Kind())
	case *types.Pointer:
		// TODO
	case *types.Slice:
		// TODO
	case *types.Array:
		// TODO
	case *types.Named:
		// TODO
	}
	panic(fmt.Sprintf("Error unknown type:\"%v\"", t))
}

func reflectBasic(b types.BasicKind) reflect.Type {
	switch b {
	default:
		panic("Unknown basic type")
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

var bInstrData = InstrDataType{INTEGER_OP, InstrData{signed: false, size: 1}, XMM_INVALID}
var f32InstrData = InstrDataType{XMM_OP, InstrData{}, XMM_F32}
var f64InstrData = InstrDataType{XMM_OP, InstrData{}, XMM_F64}

func GetInstrDataType(t types.Type) InstrDataType {
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
	instrdata := InstrDataType{
		INTEGER_OP,
		InstrData{signed: signed(t), size: sizeof(t)},
		XMM_INVALID}
	return instrdata

}

func regType(t types.Type) RegType {
	if isFloat(t) {
		return XMM_REG
	}
	return DATA_REG
}
