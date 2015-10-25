package codegen

import "fmt"

var simdToGoAsm = map[SimdInstr]InstructionType{
	AddI8x16: I_PADD,
	SubI8x16: I_PSUB,
	AddI16x8: I_PADD,
	SubI16x8: I_PSUB,
	MulI16x8: I_PIMUL,
	ShlI16x8: I_PSLL,
	ShrI16x8: I_PSRA,
	AddI32x4: I_PADD,
	SubI32x4: I_PSUB,
	ShlI32x4: I_PSLL,
	ShrI32x4: I_PSRA,
	AddI64x2: I_PADD,
	SubI64x2: I_PSUB,
	AddU8x16: I_PADD,
	SubU8x16: I_PSUB,
	AddU16x8: I_PADD,
	SubU16x8: I_PSUB,
	MulU16x8: I_PIMUL, // TODO: calculate properly using I_PMUL
	ShlU16x8: I_PSLL,
	//ShrU16x8:
	AddU32x4: I_PADD,
	SubU32x4: I_PSUB,
	//MulU32x4: I_PMUL,
	ShlU32x4: I_PSLL,
	ShrU32x4: I_PSRL,
	AddU64x2: I_PADD,
	SubU64x2: I_PSUB,
	AddF32x4: I_ADD,
	SubF32x4: I_SUB,
	MulF32x4: I_MUL,
	DivF32x4: I_DIV,
	AddF64x2: I_ADD,
	SubF64x2: I_SUB,
	MulF64x2: I_MUL,
	DivF64x2: I_DIV,
}

type SimdInstr int

const (
	SIMD_INVALID SimdInstr = iota
	// Integer
	AddI8x16
	SubI8x16
	AddI16x8
	SubI16x8
	MulI16x8
	ShlI16x8
	ShrI16x8
	AddI32x4
	SubI32x4
	MulI32x4
	ShlI32x4
	ShrI32x4
	ShufI32x4
	AddI64x2
	SubI64x2
	AddU8x16
	SubU8x16
	AddU16x8
	SubU16x8
	MulU16x8
	ShlU16x8
	ShrU16x8
	AddU32x4
	SubU32x4
	MulU32x4
	ShlU32x4
	ShrU32x4
	ShufU32x4
	AddU64x2
	SubU64x2
	AddF32x4
	SubF32x4
	MulF32x4
	DivF32x4
	AddF64x2
	SubF64x2
	MulF64x2
	DivF64x2
	LoadSi128
)

func getSimdInstr(name string) (InstructionType, bool) {
	for simdinstr, goinstr := range simdToGoAsm {
		if fmt.Sprintf("%v", simdinstr) == name {
			return goinstr, true
		}
	}
	return I_INVALID, false
}

type intrinsic func(f *Function, x, y, result *identifier) (string, *Error)

var intrinsics = map[string]intrinsic{
	"MulI32x4":  mulI32x4,
	"ShufI32x4": shufU32x4,
	"MulU32x4":  mulI32x4, //TODO: FIX
	"ShrU16x8":  shrU16x8,
	"ShufU32x4": shufU32x4,
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

// TODO
// the shift count for PSLLO (intel instruction PSLLDQ)
// must be an imm8 NOT a register or memory location
// func shlI64x2(f *Function, x, shift, result *identifier) (string, *Error) {
// }

// TODO
// there is no SSE instruction for packed shift arithmetic right with 64bit ints
// func shrI64x2(f *Function, x, shift, result *identifier) (string, *Error) {
// }

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

func shufU32x4(f *Function, x, result, order *identifier) (string, *Error) {

	asm, src, err := f.LoadSimd(x)
	if err != nil {
		panic(internal("couldn't load SIMD value"))
	}

	if order.cnst == nil {
		msg := "Shuf(I/U)32x4 the shuffle order operand must be a constant"
		return ErrorMsg(msg)
	}
	orderImm8 := uint8(order.cnst.Uint64())
	if uint64(orderImm8) != order.cnst.Uint64() {
		msgstr := "Shuf(I/U)32x4 the shuffle order operand (%v) must be <= 255"
		return ErrorMsg(fmt.Sprintf(msgstr, order.cnst.Uint64()))
	}

	dst := f.allocReg(XMM_REG, XmmRegSize)

	asm += instrImm8RegReg(f, PSHUFL, orderImm8, src, &dst)

	a, e := f.StoreSimd(&dst, result)
	if e != nil {
		panic(internal("couldn't store SIMD register"))
	}
	asm += a
	f.freeReg(dst)

	return asm, nil
}
