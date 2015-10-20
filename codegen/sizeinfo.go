package codegen

import (
	"fmt"
	"reflect"

	"golang.org/x/tools/go/types"

	"github.com/bjwbell/gensimd/simd"
)

type simdtype struct {
	name     string
	size     uint
	elemSize uint
}

func simdReflect(t reflect.Type) simdtype {
	elemSize := uint(0)
	if t.Kind() == reflect.Array {
		elemSize = uint(t.Elem().Size())
	}
	return simdtype{t.Name(), uint(t.Size()), elemSize}
}

func simdTypes() []simdtype {
	simdInt := reflect.TypeOf(simd.Int(0))
	simdInt4 := reflect.TypeOf(simd.Int4{})
	return []simdtype{simdReflect(simdInt), simdReflect(simdInt4)}
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

func simdTypeInfo(t types.Type) (simdtype, error) {
	if !isSimd(t) {
		msg := fmt.Errorf("type (%v) is not simd type", t.String())
		return simdtype{}, msg
	}
	named := t.(*types.Named)
	tname := named.Obj()
	for _, simdType := range simdTypes() {
		if tname.Name() == simdType.name {
			return simdType, nil
		}
	}
	msg := fmt.Errorf("type (%v) couldn't find simd type info", t.String())
	return simdtype{}, msg
}

func simdHasElemSize(t types.Type) bool {
	if simdtype, err := simdTypeInfo(t); err == nil {
		return simdtype.elemSize > 0
	} else {
		msg := fmt.Sprintf("Error in simdHasElemSize, type (%v) is not simd", t.String())
		panic(msg)
	}
}

func simdElemSize(t types.Type) uint {
	if simdtype, err := simdTypeInfo(t); err == nil {
		return simdtype.elemSize
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
		return sizeSlice(t)
	case *types.Array:
		return sizeArray(t)
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
	case *types.Basic:
		return alignBasic(t.Kind())
	case *types.Pointer:
		return alignPtr(t)
	case *types.Slice:
		return alignSlice(t)
	case *types.Array:
		return alignArray(t)
	case *types.Named:
		panic(fmt.Sprintf("Error unknown named type in align:\"%v\"", t))
	}
	panic(fmt.Sprintf("Error unknown type (%v)", t))
}

const tupleAlignment = 8

func alignTuple(tup *types.Tuple) uint {
	return tupleAlignment
}

func alignPtr(ptr *types.Pointer) uint {
	return uint(reflectType(ptr).Align())
}

func alignSlice(slice *types.Slice) uint {
	return uint(reflectType(slice).Align())
}

func alignArray(arr *types.Array) uint {
	return uint(reflectType(arr).Align())
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

var bInstrData = OpDataType{INTEGER_OP, InstrData{signed: false, size: 1}, XMM_INVALID}
var f32InstrData = OpDataType{XMM_OP, InstrData{}, XMM_F32}
var f64InstrData = OpDataType{XMM_OP, InstrData{}, XMM_F64}

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

	if isBasic(t) {
		return GetIntegerOpDataType(signed(t), sizeof(t))
	} else {
		panic(fmt.Sprintf("Internal error: non basic type \"%v\"", t))
	}

}

func regType(t types.Type) RegType {
	if isFloat(t) {
		return XMM_REG
	}
	return DATA_REG
}
