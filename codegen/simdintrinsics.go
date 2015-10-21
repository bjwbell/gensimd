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
	asm += intrinsicRegReg(f, instr, regx, regy)
	c, err := f.StoreSimd(regy, result)
	if err != nil {
		return "", err
	}
	asm += c

	f.freeReg(*regx)
	f.freeReg(*regy)

	return asm, err
}

func intrinsicRegReg(f *Function, instr Instr, src, dst *register) string {
	asm := f.Indent
	asm += fmt.Sprintf("%-9v    %v, %v\n", instr.String(), src.name, dst.name)
	return asm
}

func intrinsicImm8Reg(f *Function, instr Instr, imm8 uint8, dst *register) string {
	asm := f.Indent + fmt.Sprintf("%-9v    $%v, %v\n", instr.String(), imm8, dst.name)
	return asm
}

func intrinsicImm8RegReg(f *Function, instr Instr, imm8 uint8, src, dst *register) string {
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
func shlI16x8(f *Function, shift, x, result *identifier) (string, *Error) {
	return packedOp(f, I_PSLL, XMM_I16X8, shift, x, result)
}
func shrI16x8(f *Function, shift, x, result *identifier) (string, *Error) {
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
	asm += intrinsicRegReg(f, PMULULQ, regx, &tmp1) // mul dwords 2, 0

	asm += intrinsicImm8Reg(f, PSRLO, 4, regx) // shift DQ (128bit) logical right by 4
	asm += intrinsicImm8Reg(f, PSRLO, 4, regx) // shift DQ (128bit) logical right by 4

	tmp2 := f.allocReg(XMM_REG, 16)
	asm += MovRegReg(f.Indent, OpDataType{op: XMM_OP, xmmvariant: XMM_F128}, regy, &tmp2)
	asm += intrinsicRegReg(f, PMULULQ, regx, &tmp2) // mul dwords 3, 1

	// shuffle into first 64 bits of shufflet1
	shufflet1 := f.allocReg(XMM_REG, 16)
	asm += intrinsicImm8RegReg(f, PSHUFL, 0x20, &tmp1, &shufflet1)

	// shuffle into first 64 bits of shufflet2
	shufflet2 := f.allocReg(XMM_REG, 16)
	asm += intrinsicImm8RegReg(f, PSHUFL, 0x20, &tmp2, &shufflet2)

	punpckllq := PUNPCKLLQ

	// Unpack and interleave 32-bit integers from the low half of shuffletmp1 and shuffletmp2, and store the results in shufflet2.
	asm += intrinsicRegReg(f, punpckllq, &shufflet1, &shufflet2)

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
func shlI32x4(f *Function, shift, x, result *identifier) (string, *Error) {
	return packedOp(f, I_PSLL, XMM_I32X4, shift, x, result)
}
func shrI32x4(f *Function, shift, x, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRA, XMM_I32X4, shift, x, result)
}

func addI64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PADD, XMM_I64X2, x, y, result)
}
func subI64x2(f *Function, x, y, result *identifier) (string, *Error) {
	return packedOp(f, I_PSUB, XMM_I64X2, y, x, result)
}
func shlI64x2(f *Function, shift, x, result *identifier) (string, *Error) {
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
func shlU16x8(f *Function, shift, x, result *identifier) (string, *Error) {
	return shlI16x8(f, shift, x, result)
}
func shrU16x8(f *Function, shift, x, result *identifier) (string, *Error) {
	return packedOp(f, I_PSRL, XMM_U16X8, shift, x, result)
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
func shlU32x4(f *Function, shift, x, result *identifier) (string, *Error) {
	return shlI32x4(f, shift, x, result)
}
func shrU32x4(f *Function, shift, x, result *identifier) (string, *Error) {
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
func shlU64x2(f *Function, shift, x, result *identifier) (string, *Error) {
	return shlI64x2(f, shift, x, result)
}
func shrU64x2(f *Function, shift, x, result *identifier) (string, *Error) {
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
