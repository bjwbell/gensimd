package codegen

import "fmt"

type intrinsic func(f *Function, x, y, result *identifier) (string, *Error)

var intrinsics = map[string]intrinsic{
	"AddI8x16": addI8x16,
	"SubI8x16": subI8x16,
	"AddI16x8": addI16x8,
	"SubI16x8": subI16x8,
	"MulI16x8": mulI16x8,
	"ShlI16x8": shlI16x8,
	"ShrI16x8": shrI16x8,
	"AddI32x4": addI32x4,
	"SubI32x4": subI32x4,
	"MulI32x4": mulI32x4,
	"ShlI32x4": shlI32x4,
	"ShrI32x4": shrI32x4,
	"AddI64x2": addI64x2,
	"SubI64x2": subI64x2,
	"AddU8x16": addU8x16,
	"SubU8x16": subU8x16,
	"AddU16x8": addU16x8,
	"SubU16x8": subU16x8,
	"MulU16x8": mulU16x8,
	"ShlU16x8": shlU16x8,
	"ShrU16x8": shrU16x8,
	"AddU32x4": addU32x4,
	"SubU32x4": subU32x4,
	"MulU32x4": mulU32x4,
	"ShlU32x4": shlU32x4,
	"ShrU32x4": shrU32x4,
	"AddU64x2": addU64x2,
	"SubU64x2": subU64x2,
	"AddF32x4": addF32x4,
	"SubF32x4": subF32x4,
	"MulF32x4": mulF32x4,
	"DivF32x4": divF32x4,
	"AddF64x2": addF64x2,
	"SubF64x2": subF64x2,
	"MulF64x2": mulF64x2,
	"DivF64x2": divF64x2,
}

func packedOp(f *Function, instrtype InstructionType, optypes XmmData, x, y, result *identifier) (string, *Error) {
	asm := ""
	instr := GetXmmInstruction(instrtype).Select(optypes)
	a, regx, err := f.LoadSimd(x)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSimd(y)
	if err != nil {
		return "", err
	}
	asm += b
	asm += instrRegReg(f, instr, regx, regy)
	c, err := f.StoreSimd(regy, result)
	if err != nil {
		return "", err
	}
	asm += c

	f.freeReg(*regx)
	f.freeReg(*regy)

	return asm, err
}

func instrRegReg(f *Function, instr Instr, src, dst *register) string {
	asm := f.Indent
	asm += fmt.Sprintf("%-9v    %v, %v\n", instr, src.name, dst.name)
	return asm
}

func instrImm8Reg(f *Function, instr Instr, imm8 uint8, dst *register) string {
	asm := f.Indent + fmt.Sprintf("%-9v    $%v, %v\n", instr, imm8, dst.name)
	return asm
}

func instrImm8RegReg(f *Function, instr Instr, imm8 uint8, src, dst *register) string {
	asm := f.Indent + fmt.Sprintf("%-9v    $%v, %v, %v\n", instr, imm8, src.name, dst.name)
	return asm
}

// implementations of SIMD functions:
// add, sub, mul, div, <<, >> for each type
func addI8x16(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PADD, XMM_I8X16, x, y, result)
}
func subI8x16(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PSUB, XMM_I8X16, y, x, result)
}

func addI16x8(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PADD, XMM_I16X8, x, y, result)
}
func subI16x8(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PSUB, XMM_I16X8, y, x, result)
}
func mulI16x8(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PIMUL, XMM_I16X8, x, y, result)
}
func shlI16x8(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSLL, XMM_I16X8, shift, x, result)
}
func shrI16x8(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRA, XMM_I16X8, shift, x, result)
}

func addI32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PADD, XMM_I32X4, x, y, result)
}
func subI32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PSUB, XMM_I32X4, y, x, result)
}
func mulI32x4(f *Function, x, y, result *identifier) (string, *Error) {
	// native x64 SSE4.1 instruction "PMULLD"
	// emulate on SSE2 with below

	asm := ""
	tmp1 := f.allocReg(XMM_REG, 16)

	a, regx, err := f.LoadSimd(x)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSimd(y)
	if err != nil {
		return "", err
	}
	asm += b

	asm += MovRegReg(f.Indent, OpDataType{op: XMM_OP, xmmvariant: XMM_F128}, regy, &tmp1)
	asm += instrRegReg(f, PMULULQ, regx, &tmp1) // mul dwords 2, 0

	asm += instrImm8Reg(f, PSRLO, 4, regx) // shift logical right by 4 bytes
	asm += instrImm8Reg(f, PSRLO, 4, regy) // shift logical right by 4 bytes

	tmp2 := f.allocReg(XMM_REG, 16)
	asm += MovRegReg(f.Indent, OpDataType{op: XMM_OP, xmmvariant: XMM_F128}, regy, &tmp2)
	asm += instrRegReg(f, PMULULQ, regx, &tmp2) // mul dwords 3, 1

	// shuffle into first 64 bits of shufflet1
	shufflet1 := f.allocReg(XMM_REG, 16)
	asm += instrImm8RegReg(f, PSHUFD, 0x8, &tmp1, &shufflet1)

	// shuffle into first 64 bits of shufflet2
	shufflet2 := f.allocReg(XMM_REG, 16)
	asm += instrImm8RegReg(f, PSHUFD, 0x8, &tmp2, &shufflet2)

	punpckllq := PUNPCKLLQ

	// Unpack and interleave 32-bit integers from the low half of shuffletmp1 and shuffletmp2, and store the results in shufflet2.
	asm += instrRegReg(f, punpckllq, &shufflet2, &shufflet1)

	if a, err := f.StoreSimd(&shufflet1, result); err != nil {
		return "", err
	} else {
		asm += a
	}

	f.freeReg(*regx)
	f.freeReg(*regy)
	f.freeReg(tmp1)
	f.freeReg(tmp2)
	f.freeReg(shufflet1)
	f.freeReg(shufflet2)

	return asm, nil

}
func shlI32x4(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSLL, XMM_I32X4, shift, x, result)
}
func shrI32x4(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRA, XMM_I32X4, shift, x, result)
}

func addI64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PADD, XMM_I64X2, x, y, result)
}
func subI64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PSUB, XMM_I64X2, y, x, result)
}

// TODO
// the shift count for PSLLO (intel instruction PSLLDQ)
// must be an imm8 NOT a register or memory location
// func shlI64x2(f *Function, x, shift, result *identifier) (string, *Error) {
// }

// TODO
// there is no SSE instruction for packed shift arithmetic right with 64bit ints
// func shrI64x2(f *Function, x, shift, result *identifier) (string, *Error) {
// }

func addU8x16(f *Function, x, y, result *identifier) (string, *Error) {
	// the PADD instructions work for both signed and unsigned values
	return addI8x16(f, x, y, result)
}
func subU8x16(f *Function, x, y, result *identifier) (string, *Error) {
	// the PSUB instructions work for both signed and unsigned values
	return subI8x16(f, x, y, result)
}

func addU16x8(f *Function, x, y, result *identifier) (string, *Error) {
	// the PADD instructions work for both signed and unsigned values
	return addI16x8(f, x, y, result)
}
func subU16x8(f *Function, x, y, result *identifier) (string, *Error) {
	// the PSUB instructions work for both signed and unsigned values
	return subI16x8(f, x, y, result)
}
func mulU16x8(f *Function, x, y, result *identifier) (string, *Error) {
	// TODO: calculate properly
	return mulI16x8(f, x, y, result)
}
func shlU16x8(f *Function, x, count, result *identifier) (string, *Error) {
	return shlI16x8(f, x, count, result)
}
func shrU16x8(f *Function, x, count, result *identifier) (string, *Error) {

	// PSRL isn't used before go1.5.2 (https://github.com/golang/go/issues/13010)
	v152 := goversion{1, 5, 2}
	if v, e := goVersion(); e == nil && cmpGoVersion(v, v152) > 0 {
		return packedOp(f, I_PSRL, XMM_U16X8, count, x, result)
	} else {

		asm, reg, err := f.LoadSimd(x)
		if err != nil {
			panic(internal("couldn't load SIMD value"))
		}

		countReg := f.allocReg(regType(count.typ), sizeof(count.typ))
		if countReg.typ != DATA_REG {
			panic(internal("couldn't alloc register for SIMD shift right"))
		}

		a, e := f.LoadIdentSimple(count, &countReg)
		if e != nil {
			panic(internal("couldn't load shift count for SIMD shift right "))
		}
		asm += a

		wordReg := f.allocReg(DATA_REG, 8)
		tmp := f.allocReg(DATA_REG, 8)

		for i := uint8(0); i < 8; i++ {

			asm += instrImm8RegReg(f, PEXTRW, i, reg, &wordReg)

			asm += ShiftRegReg(
				f.Indent,
				false,
				SHIFT_RIGHT,
				&wordReg,
				&countReg,
				&tmp,
				wordReg.size())

			asm += instrImm8RegReg(f, PINSRW, i, &wordReg, reg)

		}

		a, e = f.StoreSimd(reg, result)
		if e != nil {
			panic(internal("couldn't store SIMD register"))
		}
		asm += a

		return asm, nil
	}
}

func addU32x4(f *Function, x, y, result *identifier) (string, *Error) {
	// the PADD instructions work for both signed and unsigned values
	return addI32x4(f, x, y, result)
}
func subU32x4(f *Function, x, y, result *identifier) (string, *Error) {
	// the PSUB instructions work for both signed and unsigned values
	return subI32x4(f, x, y, result)
}
func mulU32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return mulI32x4(f, x, y, result)
}
func shlU32x4(f *Function, x, shift, result *identifier) (string, *Error) {
	return shlI32x4(f, x, shift, result)
}
func shrU32x4(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRL, XMM_U32X4, shift, x, result)
}

func addU64x2(f *Function, x, y, result *identifier) (string, *Error) {
	// the PADD instructions work for both signed and unsigned values
	return addI64x2(f, x, y, result)
}
func subU64x2(f *Function, x, y, result *identifier) (string, *Error) {
	// the PSUB instructions work for both signed and unsigned values
	return subI64x2(f, x, y, result)
}

// TODO
// the shift count for PSLLO (intel instruction PSLLDQ)
// must be an imm8 NOT a register or memory location
// func shlU64x2(f *Function, x, shift, result *identifier) (string, *Error) {
//	return shlI64x2(f, x, shift, result)
// }

// TODO
// the shift count for PSRLO (intel instruction PSRLDQ)
// must be an imm8 NOT a register or memory location
// func shrU64x2(f *Function, x, shift, result *identifier) (string, *Error) {
// }

func addF32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_ADD, XMM_4X_F32, x, y, result)
}
func subF32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_SUB, XMM_4X_F32, y, x, result)
}
func mulF32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_MUL, XMM_4X_F32, x, y, result)
}
func divF32x4(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_DIV, XMM_4X_F32, y, x, result)
}

func addF64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_ADD, XMM_2X_F64, x, y, result)
}
func subF64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_SUB, XMM_2X_F64, y, x, result)
}
func mulF64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_MUL, XMM_2X_F64, x, y, result)
}
func divF64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_DIV, XMM_2X_F64, y, x, result)
}
