package codegen

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
)

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
	ShuffleI32x4
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
	ShuffleU32x4
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

type intrinsic func(f *Function, loc ssa.Instruction, x, y, result *identifier) (string, *Error)

var intrinsics = map[string]intrinsic{
	"MulI32x4":     mulI32x4,
	"ShuffleI32x4": shufU32x4,
	"MulU32x4":     mulI32x4, //TODO: FIX
	"ShrU16x8":     shrU16x8,
	"ShuffleU32x4": shufU32x4,
}

func packedOp(f *Function, loc ssa.Instruction, instrtype InstructionType, optypes XmmData, x, y, result *identifier) (string, *Error) {
	asm := ""
	instr := GetXmmInstruction(instrtype).Select(optypes)
	a, regx, err := f.LoadSimd(loc, x)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSimd(loc, y)
	if err != nil {
		return "", err
	}
	asm += b
	asm += instrRegReg(instr, regx, regy)
	c, err := f.StoreSimd(loc, regy, result)
	if err != nil {
		return "", err
	}
	asm += c

	f.freeReg(regx)
	f.freeReg(regy)

	return asm, err
}

func instrImm8Reg(f *Function, instr Instruction, imm8 uint8, dst *register) string {
	info := instrTable[instr]
	asm := ""
	if info.Flags&RightRdwr != 0 || info.Flags&RightWrite != 0 {
		asm += dst.modified()
	} else {
		fmt.Println("Instr:", instr)
		ice("dst modify flag should be set")
	}
	return asm + fmt.Sprintf("%-9v    $%v, %v\n", instr, imm8, dst.name)
}

func instrImm8RegReg(f *Function, instr Instruction, imm8 uint8, src, dst *register) string {
	info := instrTable[instr]
	asm := ""
	if info.Flags&RightRdwr != 0 || info.Flags&RightWrite != 0 {
		asm += dst.modified()
	} else {
		fmt.Println("Instr:", instr)
		ice("dst modify flag should be set")
	}
	return asm + fmt.Sprintf("%-9v    $%v, %v, %v\n", instr, imm8, src.name, dst.name)
}

// implementations of SIMD functions:
// add, sub, mul, div, <<, >> for each type

func mulI32x4(f *Function, loc ssa.Instruction, x, y, result *identifier) (string, *Error) {
	// native x64 SSE4.1 instruction "PMULLD"
	// emulate on SSE2 with below

	asm := ""
	a1, tmp1 := f.allocReg(loc, XMM_REG, 16)
	asm += a1

	a, regx, err := f.LoadSimd(loc, x)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSimd(loc, y)
	if err != nil {
		return "", err
	}
	asm += b

	asm += MovRegReg(OpDataType{op: OP_PACKED, xmmvariant: XMM_F128}, regy, tmp1)
	asm += instrRegReg(PMULULQ, regx, tmp1) // mul dwords 2, 0

	asm += instrImm8Reg(f, PSRLO, 4, regx) // shift logical right by 4 bytes
	if regy.name != regx.name {
		asm += instrImm8Reg(f, PSRLO, 4, regy) // shift logical right by 4 bytes
	}

	a2, tmp2 := f.allocReg(loc, XMM_REG, 16)
	asm += a2

	asm += MovRegReg(OpDataType{op: OP_PACKED, xmmvariant: XMM_F128}, regy, tmp2)
	asm += instrRegReg(PMULULQ, regx, tmp2) // mul dwords 3, 1

	// shuffle into first 64 bits of shufflet1
	a3, shufflet1 := f.allocReg(loc, XMM_REG, 16)
	asm += a3

	asm += instrImm8RegReg(f, PSHUFD, 0x8, tmp1, shufflet1)

	// shuffle into first 64 bits of shufflet2
	a4, shufflet2 := f.allocReg(loc, XMM_REG, 16)
	asm += a4

	asm += instrImm8RegReg(f, PSHUFD, 0x8, tmp2, shufflet2)

	punpckllq := PUNPCKLLQ

	// Unpack and interleave 32-bit integers from the low half of shuffletmp1 and shuffletmp2, and store the results in shufflet2.
	asm += instrRegReg(punpckllq, shufflet2, shufflet1)

	if a, err := f.StoreSimd(loc, shufflet1, result); err != nil {
		return "", err
	} else {
		asm += a
	}

	f.freeReg(regx)
	f.freeReg(regy)
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

func shrU16x8(f *Function, loc ssa.Instruction, x, count, result *identifier) (string, *Error) {

	// PSRL isn't used before go1.5.2 (https://github.com/golang/go/issues/13010)
	v152 := goversion{1, 5, 2}
	if v, e := goVersion(); e == nil && cmpGoVersion(v, v152) > 0 {
		return packedOp(f, loc, I_PSRL, XMM_U16X8, count, x, result)
	} else {

		asm, reg, err := f.LoadSimd(loc, x)
		if err != nil {
			panic(ice("couldn't load SIMD value"))
		}

		a, countReg, e := f.LoadIdentSimple(loc, count)
		if e != nil {
			panic(ice("couldn't load shift count for SIMD shift right "))
		}
		if countReg.typ != DATA_REG {
			panic(ice("couldn't alloc register for SIMD shift right"))
		}
		asm += a

		a1, wordReg := f.allocReg(loc, DATA_REG, 8)
		asm += a1

		a2, tmp := f.allocReg(loc, DATA_REG, 8)
		asm += a2

		for i := uint8(0); i < 8; i++ {

			asm += instrImm8RegReg(f, PEXTRW, i, reg, wordReg)

			asm += ShiftRegReg(
				false,
				SHIFT_RIGHT,
				wordReg,
				countReg,
				tmp,
				wordReg.size())

			asm += instrImm8RegReg(f, PINSRW, i, wordReg, reg)

		}

		a, e = f.StoreSimd(loc, reg, result)
		if e != nil {
			panic(ice("couldn't store SIMD register"))
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

func shufU32x4(f *Function, loc ssa.Instruction, x, result, order *identifier) (string, *Error) {

	asm, src, err := f.LoadSimd(loc, x)
	if err != nil {
		panic(ice("couldn't load SIMD value"))
	}

	if order.cnst == nil {
		msg := "Shuf(I/U)32x4 the shuffle order operand must be a constant"
		return ErrorMsg(msg)
	}
	orderImm8 := uint8(order.cnst.Uint64())
	if uint64(orderImm8) != order.cnst.Uint64() {
		msgstr := "Shuffle(I/U)32x4 the shuffle order operand (%v) must be <= 255"
		return ErrorMsg(fmt.Sprintf(msgstr, order.cnst.Uint64()))
	}

	a1, dst := f.allocReg(loc, XMM_REG, XmmRegSize)
	asm += a1

	asm += instrImm8RegReg(f, PSHUFL, orderImm8, src, dst)

	a, e := f.StoreSimd(loc, dst, result)
	if e != nil {
		panic(ice("couldn't store SIMD register"))
	}
	asm += a
	f.freeReg(dst)

	return asm, nil
}
