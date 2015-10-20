package simd

// go implementations of SIMD functions:
// add, sub, mul, div, <<, >> for each type
func AddI8x16(x, y i8x16) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubI8x16(x, y i8x16) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI8x16(x, y i8x16) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivI8x16(x, y i8x16) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlI8x16(x, shift uint8) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrI8x16(x, shift uint8) i8x16 {
	val := i8x16{}
	for i := 0; i < 16; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddI16x8(x, y i16x8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubI16x8(x, y i16x8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI16x8(x, y i16x8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivI16x8(x, y i16x8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlI16x8(x, shift uint8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrI16x8(x, shift uint8) i16x8 {
	val := i16x8{}
	for i := 0; i < 8; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddI32x4(x, y i32x4) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubI32x4(x, y i32x4) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI32x4(x, y i32x4) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivI32x4(x, y i32x4) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlI32x4(x, shift uint8) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrI32x4(x, shift uint8) i32x4 {
	val := i32x4{}
	for i := 0; i < 4; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddI64x2(x, y i64x2) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubI64x2(x, y i64x2) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulI64x2(x, y i64x2) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivI64x2(x, y i64x2) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlI64x2(x, shift uint8) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrI64x2(x, shift uint8) i64x2 {
	val := i64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddU8x16(x, y u8x16) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubU8x16(x, y u8x16) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU8x16(x, y u8x16) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivU8x16(x, y u8x16) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlU8x16(x, shift uint8) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrU8x16(x, shift uint8) u8x16 {
	val := u8x16{}
	for i := 0; i < 16; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddU16x8(x, y u16x8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubU16x8(x, y u16x8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU16x8(x, y u16x8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivU16x8(x, y u16x8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlU16x8(x, shift uint8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrU16x8(x, shift uint8) u16x8 {
	val := u16x8{}
	for i := 0; i < 8; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddU32x4(x, y u32x4) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubU32x4(x, y u32x4) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU32x4(x, y u32x4) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivU32x4(x, y u32x4) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlU32x4(x, shift uint8) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrU32x4(x, shift uint8) u32x4 {
	val := u32x4{}
	for i := 0; i < 4; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddU64x2(x, y u64x2) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubU64x2(x, y u64x2) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulU64x2(x, y u64x2) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivU64x2(x, y u64x2) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
func ShlU64x2(x, shift uint8) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] << shift
	}
	return val
}
func ShrU64x2(x, shift uint8) u64x2 {
	val := u64x2{}
	for i := 0; i < 2; i++ {
		val[i] = val[i] >> shift
	}
	return val
}

func AddF32x4(x, y f32x4) f32x4 {
	val := f32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubF32x4(x, y f32x4) f32x4 {
	val := f32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulF32x4(x, y f32x4) f32x4 {
	val := f32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivF32x4(x, y f32x4) f32x4 {
	val := f32x4{}
	for i := 0; i < 4; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}

func AddF64x2(x, y f64x2) f64x2 {
	val := f64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] + y[i]
	}
	return val
}
func SubF64x2(x, y f64x2) f64x2 {
	val := f64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] - y[i]
	}
	return val
}
func MulF64x2(x, y f64x2) f64x2 {
	val := f64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] * y[i]
	}
	return val
}
func DivF64x2(x, y f64x2) f64x2 {
	val := f64x2{}
	for i := 0; i < 2; i++ {
		val[i] = x[i] / y[i]
	}
	return val
}
