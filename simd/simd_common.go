package simd

// go implementations of SIMD functions:
// add, sub, mul, div, <<, >> for each type

func AddI8x16(x, y I8x16) I8x16 {
	val := I8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubI8x16 subtracts y from x
func SubI8x16(x, y I8x16) I8x16 {
	val := I8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}

func AddI16x8(x, y I16x8) I16x8 {
	val := I16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubI16x8 subtracts y from x
func SubI16x8(x, y I16x8) I16x8 {
	val := I16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI16x8(x, y I16x8) I16x8 {
	val := I16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func ShlI16x8(x I16x8, shift uint8) I16x8 {
	val := I16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] << shift
	}
	return val
}
func ShrI16x8(x I16x8, shift uint8) I16x8 {
	val := I16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] >> shift
	}
	return val
}

func AddI32x4(x, y I32x4) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubI32x4 subtracts y from x
func SubI32x4(x, y I32x4) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI32x4(x, y I32x4) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func ShlI32x4(x I32x4, shift uint8) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] << shift
	}
	return val
}
func ShrI32x4(x I32x4, shift uint8) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] >> shift
	}
	return val
}
func ShuffleI32x4(x I32x4, order uint8) I32x4 {
	val := I32x4{}
	for i := 0; i < 4; i++ {
		switch (order >> uint8(2*i)) & 0x3 {
		case 0:
			val[i] = x[0]
		case 1:
			val[i] = x[1]
		case 2:
			val[i] = x[2]
		case 3:
			val[i] = x[3]
		}
	}
	return val
}

func AddI64x2(x, y I64x2) I64x2 {
	val := I64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubI64x2 subtracts y from x
func SubI64x2(x, y I64x2) I64x2 {
	val := I64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func ShlI64x2(x, shift uint8) I64x2 {
	val := I64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] << shift
	}
	return val
}

func AddU8x16(x, y U8x16) U8x16 {
	val := U8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubU8x16 subtracts y from x
func SubU8x16(x, y U8x16) U8x16 {
	val := U8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}

func AddU16x8(x, y U16x8) U16x8 {
	val := U16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubU16x8 subtracts y from x
func SubU16x8(x, y U16x8) U16x8 {
	val := U16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU16x8(x, y U16x8) U16x8 {
	val := U16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func ShlU16x8(x U16x8, shift uint8) U16x8 {
	val := U16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] << shift
	}
	return val
}
func ShrU16x8(x U16x8, shift uint8) U16x8 {
	val := U16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] >> shift
	}
	return val
}

func AddU32x4(x, y U32x4) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubU32x4 subtracts y from x
func SubU32x4(x, y U32x4) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU32x4(x, y U32x4) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func ShlU32x4(x U32x4, shift uint8) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] << shift
	}
	return val
}
func ShrU32x4(x U32x4, shift uint8) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] >> shift
	}
	return val
}
func ShuffleU32x4(x U32x4, order uint8) U32x4 {
	val := U32x4{}
	for i := 0; i < 4; i++ {
		switch (order >> uint8(2*i)) & 0x3 {
		case 0:
			val[i] = x[0]
		case 1:
			val[i] = x[1]
		case 2:
			val[i] = x[2]
		case 3:
			val[i] = x[3]
		}
	}
	return val

}

func AddU64x2(x, y U64x2) U64x2 {
	val := U64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubU64x2 subtracts y from x
func SubU64x2(x, y U64x2) U64x2 {
	val := U64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func ShlU64x2(x, shift uint8) U64x2 {
	val := U64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrU64x2(x, shift uint8) U64x2 {
	val := U64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddF32x4(x, y F32x4) F32x4 {
	val := F32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubF32x4 subtracts y from x
func SubF32x4(x, y F32x4) F32x4 {
	val := F32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulF32x4(x, y F32x4) F32x4 {
	val := F32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivF32x4(x, y F32x4) F32x4 {
	val := F32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}

func AddF64x2(x, y F64x2) F64x2 {
	val := F64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}

// SubF64x2 subtracts y from x
func SubF64x2(x, y F64x2) F64x2 {
	val := F64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulF64x2(x, y F64x2) F64x2 {
	val := F64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivF64x2(x, y F64x2) F64x2 {
	val := F64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
