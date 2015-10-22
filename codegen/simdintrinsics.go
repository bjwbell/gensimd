package codegen

import "fmt"

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
	asm += fmt.Sprintf("%-9v    %v, %v\n", instr.String(), src.name, dst.name)
	return asm
}

func instrImm8Reg(f *Function, instr Instr, imm8 uint8, dst *register) string {
	asm := f.Indent + fmt.Sprintf("%-9v    $%v, %v\n", instr.String(), imm8, dst.name)
	return asm
}

func instrImm8RegReg(f *Function, instr Instr, imm8 uint8, src, dst *register) string {
	asm := f.Indent + fmt.Sprintf("%-9v    $%v, %v, %v\n", instr.String(), imm8, src.name, dst.name)
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

	asm += instrImm8Reg(f, PSRLO, 4, regx) // shift DQ (128bit) logical right by 4
	asm += instrImm8Reg(f, PSRLO, 4, regx) // shift DQ (128bit) logical right by 4

	tmp2 := f.allocReg(XMM_REG, 16)
	asm += MovRegReg(f.Indent, OpDataType{op: XMM_OP, xmmvariant: XMM_F128}, regy, &tmp2)
	asm += instrRegReg(f, PMULULQ, regx, &tmp2) // mul dwords 3, 1

	// shuffle into first 64 bits of shufflet1
	shufflet1 := f.allocReg(XMM_REG, 16)
	asm += instrImm8RegReg(f, PSHUFL, 0x20, &tmp1, &shufflet1)

	// shuffle into first 64 bits of shufflet2
	shufflet2 := f.allocReg(XMM_REG, 16)
	asm += instrImm8RegReg(f, PSHUFL, 0x20, &tmp2, &shufflet2)

	punpckllq := PUNPCKLLQ

	// Unpack and interleave 32-bit integers from the low half of shuffletmp1 and shuffletmp2, and store the results in shufflet2.
	asm += instrRegReg(f, punpckllq, &shufflet1, &shufflet2)

	if a, err := f.StoreSimd(&shufflet2, result); err != nil {
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
func shlI64x2(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSLL, XMM_I64X2, shift, x, result)
}

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
func shlU64x2(f *Function, x, shift, result *identifier) (string, *Error) {
	return shlI64x2(f, x, shift, result)
}
func shrU64x2(f *Function, x, shift, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRL, XMM_U64X2, shift, x, result)
}

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
	return packedOp(f, I_DIV, XMM_4X_F32, x, y, result)
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
	return packedOp(f, I_DIV, XMM_2X_F64, x, y, result)
}
