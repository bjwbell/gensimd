package codegen

import (
	"fmt"
	"go/token"
	"math"
	"strings"
)

type InstrOpType int

const (
	OP_INVALID InstrOpType = iota
	OP_DATA                // int8/uint8,,..., int64/uint64
	OP_XMM                 // f32/f64, packed f32, packed f64
	OP_PACKED              // packed int8/uint8,....,int63/uint64
)

type OpDataType struct {
	op InstrOpType
	InstrData
	xmmvariant XmmData
}

func (optype OpDataType) String() string {
	return fmt.Sprintf("OpDataType{Op:%v, InstrData: %v, Xmm: %v}", optype.op, optype.InstrData, optype.xmmvariant)
}

type InstrData struct {
	size   uint
	signed bool
}

type XmmData int

const (
	XMM_INVALID XmmData = iota
	XMM_F32
	XMM_F64
	XMM_4X_F32
	XMM_2X_F64
	XMM_F128  = XMM_2X_F64
	XMM_M128  = XMM_F128 // [4]float32
	XMM_M128i = XMM_F128 // [16]byte
	XMM_M128d = XMM_F128 // [2]float64

	XMM_I8X16 = XMM_F32
	XMM_U8X16 = XMM_I8X16

	XMM_I16X8 = XMM_F64
	XMM_U16X8 = XMM_I16X8

	XMM_I32X4 = XMM_4X_F32
	XMM_U32X4 = XMM_I32X4

	XMM_I64X2 = XMM_2X_F64
	XMM_U64X2 = XMM_I64X2
)

type Instruction int
type DataInstruction struct {
	TInstr InstructionType

	// integer forms
	ByteSized Instruction
	WordSized Instruction
	LongSized Instruction
	QuadSized Instruction
	OctSized  Instruction

	// xmm forms
	/*SingleF32 Instr
	SingleF64 Instr
	PackedF32 Instr // operates on four packed f32
	PackedF64 Instr // operate on two packed f64*/

}

type XmmInstruction struct {
	XmmInstr InstructionType
	F32      Instruction // operates on single f32
	F64      Instruction // operates on single f64
	F32x4    Instruction // operates on four packed f32
	F64x2    Instruction // operate on two packed f64
}

func (inst DataInstruction) GetSized(size uint) Instruction {
	var instr Instruction
	switch size {
	case 1:
		instr = inst.ByteSized
	case 2:
		instr = inst.WordSized
	case 4:
		instr = inst.LongSized
	case 8:
		instr = inst.QuadSized
	case 16:
		instr = inst.OctSized
	}

	if instr == NONE {
		msgstr := "invalid size(%v), for instr (%v)"
		ice(fmt.Sprintf(msgstr, size, inst.TInstr))
	}

	return instr

}

type InstructionType int

const (
	I_INVALID InstructionType = iota
	I_ADD
	I_AND
	I_CMP

	// convert f32/64 to a uint8/int8, ..., uint64/int64
	I_CVT_FLOAT2INT
	// convert f32/f64 to a uint8/int8, ..., uint64/int64
	I_CVT_INT2FLOAT

	// convert f32/f64 to f64/f32
	I_CVT_FLOAT2FLOAT

	I_DIV
	I_IMUL
	I_IDIV
	I_LEA
	I_MOV

	// mov byte, sign extend
	I_MOVBSX
	// mov word, sign extend
	I_MOVWSX
	// mov long, sign extend
	I_MOVLSX

	// mov byte, zero extend
	I_MOVBZX
	// mov word, zero extend
	I_MOVWZX
	// mov long, zero extend
	I_MOVLZX
	I_MUL

	I_OR

	// instructions for packed integers
	I_PADD
	I_PAND
	I_PANDN
	I_PCMPEQ
	I_PCMPGT
	I_PIMUL
	I_PMUL
	I_POR
	I_PSLL //packed shift left logical
	I_PSRA // packed shift right arithmetic
	I_PSRL //packed shift right logical
	I_PSUB
	I_PXOR
	I_PMOV

	I_SAL
	I_SAR

	I_SHL
	I_SHR
	I_SUB
	I_XOR
)

var Insts = []DataInstruction{
	{I_ADD, ADDB, ADDW, ADDL, ADDQ, NONE},
	{I_SUB, SUBB, SUBW, SUBL, SUBQ, NONE},
	{I_MOV, MOVB, MOVW, MOVL, MOVQ, MOVOU},

	// byte register sign extend to xxx register
	{I_MOVBSX, NONE, MOVBWSX, MOVBLSX, MOVBQSX, NONE},
	// word register sign extend to xxx register
	{I_MOVWSX, NONE, NONE, MOVWLSX, MOVWQSX, NONE},
	// long register sign extend to xxx register
	{I_MOVLSX, NONE, NONE, NONE, MOVLQSX, NONE},

	// mov byte, zero extend to xxx register
	{I_MOVBZX, NONE, MOVBWZX, MOVBLZX, MOVBQZX, NONE},
	// mov word, zero extend to xxx register
	{I_MOVWZX, NONE, NONE, MOVWLZX, MOVWQZX, NONE},
	// mov long, zero extend to xxx register
	{I_MOVLZX, NONE, NONE, NONE, MOVLQZX, NONE},

	{I_XOR, XORB, XORW, XORL, XORQ, NONE},
	{I_LEA, NONE, LEAW, LEAL, LEAQ, NONE},
	{I_MUL, MULB, MULW, MULL, MULQ, NONE},

	// signed multiplication
	{I_IMUL, IMULB, IMULW, IMULL, IMULQ, NONE},

	{I_DIV, DIVB, DIVW, DIVL, DIVQ, NONE},

	// signed division
	{I_IDIV, IDIVB, IDIVW, IDIVL, IDIVQ, NONE},

	{I_AND, ANDB, ANDW, ANDL, ANDQ, NONE},
	{I_OR, ORB, ORW, ORL, ORQ, NONE},
	{I_SHL, SHLB, SHLW, SHLL, SHLQ, NONE},
	{I_SHR, SHRB, SHRW, SHRL, SHRQ, NONE},

	// arithmetic shift left (signed left shift)
	{I_SAL, SALB, SHLW, SHLL, SHLQ, NONE},
	// arithmetic shift right (signed right shift)
	{I_SAR, SARB, SARW, SARL, SARQ, NONE},

	{I_CMP, CMPB, CMPW, CMPL, CMPQ, NONE},
}

var XmmInsts = []XmmInstruction{
	// floating point
	{I_ADD, ADDSS, ADDSD, ADDPS, ADDPD},
	{I_CMP, UCOMISS, UCOMISD, NONE, NONE},
	{I_DIV, DIVSS, DIVSD, DIVPS, DIVPD},
	{I_MOV, MOVSS, MOVSD, MOVUPS, MOVUPD},
	{I_MUL, MULSS, MULSD, MULPS, MULPD},
	{I_SUB, SUBSS, SUBSD, SUBPS, SUBPD},
	{I_XOR, NONE, NONE, XORPS, XORPD},

	// Add packed integer
	//
	// From 64-IA-32-Architectures-Software-Developer-Instruction-Set-Reference-Manual:
	//
	// Adds the packed byte, word, doubleword, or quadword integers in the
	// first source operand to the second source operand and stores the result
	// in the destination operand (which is usually the second source operand).
	//
	// When a result is too large to be represented
	// in the 8/16/32 integer (overflow), the result is wrapped around and the
	// low bits are written to the destination element (that is, the carry is ignored).
	// Note that these instructions can operate on either unsigned or signed
	// (two’s complement notation) integers; however, it does not set bits in
	// the EFLAGS register to indicate overflow and/or a carry. To prevent
	// undetected overflow conditions, software must control the ranges of
	// the values operated on .
	{I_PADD, PADDB, PADDW, PADDL, PADDQ},

	{I_PAND, PANDB, PANDW, PANDL, PAND},
	// bitwise logical and not (&^)
	{I_PANDN, NONE, NONE, NONE, PANDN},
	{I_PCMPEQ, PCMPEQB, PCMPEQW, PCMPEQL, NONE},
	{I_PCMPGT, PCMPGTB, PCMPGTW, PCMPGTL, NONE},

	// packed signed multiplication
	{I_PIMUL, NONE, PMULLW, NONE, NONE},

	// mov packed integers
	{I_PMOV, MOVOU, MOVOU, MOVOU, MOVOU},

	// bitwise logical, operates on full register width, so same version used for all
	// op types
	{I_POR, POR, POR, POR, POR},

	// packed shift left logical
	{I_PSLL, NONE, PSLLW, PSLLL, PSLLQ},

	// packed shift right arithmetic
	{I_PSRA, NONE, PSRAW, PSRAL, NONE},

	// packed shift right logical
	{I_PSRL, NONE, PSRLW, PSRLL, PSRLQ},

	// packed unsigned multiplication
	{I_PMUL, NONE, NONE, PMULULQ, NONE},

	// Subtract packed integers
	//
	// From 64-IA-32-Architectures-Software-Developer-Instruction-Set-Reference-Manual:
	// Performs a SIMD subtract of the packed integers of the source operand
	// (second operand) from the packed integers of the destination operand
	// (first operand), and stores the packed integer results in the
	// destination operand (which is usually the second source operand).
	// When an individual result is too large or too small to be represented,
	// the result is wrapped around and the low bits are written to the
	// destination element.
	{I_PSUB, PSUBB, PSUBW, PSUBL, PSUBQ},

	{I_PXOR, NONE, NONE, NONE, PXOR},
}

func GetInstruction(tinst InstructionType) DataInstruction {
	for _, inst := range Insts {
		if inst.TInstr == tinst {
			return inst
		}
	}
	ice("couldn't get instruction")
	return DataInstruction{}
}

func GetXmmInstruction(tinst InstructionType) XmmInstruction {
	for _, inst := range XmmInsts {
		if inst.XmmInstr == tinst {
			return inst
		}
	}
	ice(fmt.Sprintf("Couldn't get xmm instruction (%v)", tinst))
	return XmmInstruction{}
}

func (xinstr XmmInstruction) Select(variant XmmData) Instruction {
	var instr Instruction
	switch variant {
	case XMM_F32:
		instr = xinstr.F32
	case XMM_F64:
		instr = xinstr.F64
	case XMM_4X_F32:
		instr = xinstr.F32x4
	case XMM_2X_F64:
		instr = xinstr.F64x2
	}
	if instr == NONE {
		s := fmt.Sprintf("unrecognized variant \"%v\" for \"%v\"", variant, xinstr.XmmInstr)
		panic(ice(s))
	}
	return instr
}

func GetConvertInstruction(tinst InstructionType, fromsize, tosize uint) Instruction {
	if tinst == I_CVT_FLOAT2INT {
		if fromsize == 4 && tosize <= 4 {
			// f32 to int32 with truncation
			return CVTTSS2SL
		} else if fromsize == 4 && tosize == 8 {
			// f32 to int64 with truncation
			return CVTTSS2SQ
		} else if fromsize == 8 && tosize <= 4 {
			// f64 to int32 with truncation
			return CVTTSD2SL
		} else if fromsize == 8 && tosize == 8 {
			// f64 to int64 with truncation
			return CVTTSD2SQ
		}
	} else if tinst == I_CVT_INT2FLOAT {
		if fromsize == 4 && tosize == 4 {
			// int32 to f32
			return CVTSL2SS
		} else if fromsize == 4 && tosize == 8 {
			// int32 to f64
			return CVTSL2SD
		} else if fromsize == 8 && tosize == 4 {
			// int64 to f32
			return CVTSQ2SS
		} else if fromsize == 8 && tosize == 8 {
			// int64 to f64
			return CVTSQ2SD
		}
	} else if tinst == I_CVT_FLOAT2FLOAT {
		if fromsize == 4 && tosize == 8 {
			return CVTSS2SD
		} else if fromsize == 8 && tosize == 4 {
			return CVTSD2SS
		}
	}
	ice(fmt.Sprintf("numeric type conversion instruction (%v)", tinst))
	return NONE
}

// GetInstr, the size is in bytes
func GetInstr(tinst InstructionType, datatype OpDataType) Instruction {
	if datatype.op == OP_XMM || datatype.op == OP_PACKED {
		return GetXmmInstruction(tinst).Select(datatype.xmmvariant)
	} else {
		return GetInstruction(tinst).GetSized(datatype.size)
	}
}

// ZeroMemory generates "MOVQ $0, name+offset(REG)" instructions,
// size is in bytes
func ZeroMemory(name string, offset int, size uint, reg *register) string {
	chunk := uint(1)
	if size%8 == 0 {
		chunk = 8
	} else if size%4 == 0 {

		chunk = 4
	} else if size%2 == 0 {

		chunk = 2
	} else {

		chunk = 1
	}
	asm := ""
	if size/chunk == 0 {
		return asm
	}
	datatype := OpDataType{OP_DATA, InstrData{signed: false, size: chunk}, XMM_INVALID}
	mov := GetInstr(I_MOV, datatype)
	for i := uint(0); i < size/chunk; i++ {
		ioffset := int(i*chunk) + offset
		asm += instrImmMem(mov, 0, reg, name, ioffset)
	}
	return asm
}

func instrReg(instr Instruction, reg *register, spill bool) string {
	if _, ok := instrTable[instr]; !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := instrTable[instr].Flags
	asm := ""
	if flags&RightRdwr != 0 || flags&RightWrite != 0 {
		asm += reg.modified(spill)
	}
	asm += fmt.Sprintf("%-9v    %v\n", instr, reg.name)
	return asm
}

func instrRegReg(instr Instruction, src, dst *register, spill bool) string {
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags
	asm := ""
	if flags&LeftRdwr != 0 {
		asm += src.modified(spill)
	}
	if (flags&RightRdwr != 0) || (flags&RightWrite != 0) {
		asm += dst.modified(spill)
	}
	asm += fmt.Sprintf("%-9v    %v, %v\n", instr, src.name, dst.name)
	return asm
}

func instrRegMem(instr Instruction, src, dst *register, dstName string, dstOffset int, spill bool) string {
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags

	var asm string
	if flags&LeftRdwr != 0 {
		asm += src.modified(spill)
	}
	if dstName != "" || dstOffset != 0 {
		asm += fmt.Sprintf("%-9v    %v, %v+%v(%v)\n", instr, src.name, dstName, dstOffset, dst.name)
	} else {
		asm += fmt.Sprintf("%-9v    %v, (%v)\n", instr, src.name, dst.name)
	}
	return strings.Replace(asm, "+-", "-", -1)
}

func instrMemReg(instr Instruction, srcName string, srcOffset int, src, dst *register, spill bool) string {
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags

	var asm string
	if (flags&RightRdwr != 0) || (flags&RightWrite != 0) {
		asm += dst.modified(spill)
	}
	if srcName != "" || srcOffset != 0 {
		asm += fmt.Sprintf("%-9v    %v+%v(%v), %v\n", instr, srcName, srcOffset, src.name, dst.name)
	} else {
		asm += fmt.Sprintf("%-9v    (%v), %v\n", instr, src.name, dst.name)
	}
	return strings.Replace(asm, "+-", "-", -1)
}

// instrImmReg outputs instr with imm, reg after converting imm to int8/16/32/64 if size=1/2/4/8.
func instrImmReg(instr Instruction, imm int64, size uint, dst *register, spill bool) string {
	if dst.width < 8*size {
		ice("Invalid register width")
	}
	switch size {
	case 1:
		imm = int64(int8(imm))
	case 2:
		imm = int64(int16(imm))
	case 4:
		imm = int64(int32(imm))
	}
	var asm string
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags
	if (flags&RightRdwr != 0) || (flags&RightWrite != 0) {
		asm += dst.modified(spill)
	}
	asm += fmt.Sprintf("%-9v    $%v, %v\n", instr, imm, dst.name)
	return asm
}

// instrImmReg outputs instr with imm, reg after converting imm to int8/16/32/64 if size=1/2/4/8.
func instrImmMem(instr Instruction, imm int64, dst *register, dstName string, dstOffset int) string {
	asm := fmt.Sprintf("%-9v    $%v, %v+%v(%v)\n", instr, imm, dstName, dstOffset, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

// instrImmUnsignedReg outputs instr with imm64, reg
func instrImmUnsignedReg(instr Instruction, imm64 uint64, size uint, r *register, spill bool) string {
	if r.width < 8*size {
		ice("Invalid register width")
	}
	var asm string
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags
	if (flags&RightRdwr != 0) || (flags&RightWrite != 0) {
		asm += r.modified(spill)
	}
	asm += fmt.Sprintf("%-9v    $%v, %v\n", instr, imm64, r.name)
	return asm
}

func instrMemImm32(instr Instruction, name string, offset int32, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		ice("Invalid register width")
	}
	asm := fmt.Sprintf("%-9v    %v+%v(%v), $%v\n", instr, name, offset, r.name, imm32)
	return strings.Replace(asm, "+-", "-", -1)

}

func instrRegImm32(instr Instruction, r *register, imm32 uint32, size uint, spill bool) string {
	if r.width < 8*size {
		ice("Invalid register width")
	}
	var asm string
	info, ok := instrTable[instr]
	if !ok {
		ice(fmt.Sprintf("couldn't look up instruction (%v) information", instr))
	}
	flags := info.Flags
	if flags&LeftRdwr != 0 {
		asm += r.modified(spill)
	}
	asm += fmt.Sprintf("%-9v    %v, $%v\n", instr, r.name, imm32)
	return asm
}

// instrImmRegReg outputs instr with imm, reg, reg after converting imm to int8/16/32 if size=1/2/4.
func instrImmRegReg(instr Instruction, imm int64, size uint, src, dst *register, spill bool) string {
	if dst.width < 8*size {
		ice("Invalid register width")
	}
	switch size {
	case 1:
		imm = int64(int8(imm))
	case 2:
		imm = int64(int16(imm))
	case 4:
		imm = int64(int32(imm))
	}

	// HACK! should lookup flags
	asm := dst.modified(spill)

	asm += fmt.Sprintf("%-9v    $%v, %v, %v\n", instr, imm, src.name, dst.name)
	return asm
}

func instrImmFloatReg(instr Instruction, f64 float64, isf32 bool, dst, tmp *register, spill bool) string {
	if dst.typ != XMM_REG {
		ice("Unexpected non xmm register")
	}
	var fbits uint64
	if isf32 {
		fbits = uint64(math.Float32bits(float32(f64)))
	} else {
		fbits = math.Float64bits(f64)
	}
	var descrip string
	if isf32 {
		descrip = "float32"
	} else {
		descrip = "float64"
	}
	asm := fmt.Sprintf("//%-9v  $%v = %016x = %v(%v)\n", " ", fbits, fbits, f64, descrip)
	asm += fmt.Sprintf("%-9v    $%v, %v\n", instr, fbits, tmp.name)

	data := OpDataType{OP_DATA, InstrData{signed: false, size: 8}, XMM_INVALID}
	asm += instrRegReg(GetInstr(I_MOV, data), tmp, dst, spill)

	return asm
}

// ZeroReg generates "XORQ reg, reg" instructions
func ZeroReg(reg *register) string {
	var dt OpDataType
	if reg.typ == XMM_REG {
		dt = OpDataType{OP_XMM, InstrData{}, XMM_2X_F64}

	} else {
		dt = OpDataType{OP_DATA, InstrData{signed: false, size: reg.width / 8}, XMM_INVALID}
	}
	return instrRegReg(GetInstr(I_XOR, dt), reg, reg, false)
}

func MovRegReg(datatype OpDataType, src, dst *register, spill bool) string {
	if src.width != dst.width && src.typ == dst.typ {
		msg := "src (%v) width (%v) != (%v) dst (%v) width or invalid size"
		ice(fmt.Sprintf(msg, src.name, src.width, dst.width, dst.name))
	}
	var mov Instruction
	if src.typ == XMM_REG && dst.typ == XMM_REG {
		mov = MOVO
	} else if src.typ == XMM_REG || dst.typ == XMM_REG {
		mov = MOVQ
	} else {
		mov = GetInstr(I_MOV, datatype)
	}
	return instrRegReg(mov, src, dst, spill)
}

func MovRegMem(datatype OpDataType, src *register, dstName string, dst *register, dstOffset int) string {
	var mov Instruction
	if datatype.op == OP_PACKED {
		mov = GetInstr(I_PMOV, datatype)
	} else {
		mov = GetInstr(I_MOV, datatype)
	}
	return instrRegMem(mov, src, dst, dstName, dstOffset, false)
}

func MovRegIndirectMem(datatype OpDataType, src *register, dstName string, dstOffset int, dst *register, size uint, tmpAddr, tmpData *register) string {
	if size <= DataRegSize {
		return MovRegIndirectMemSmall(datatype, src, dstName, dstOffset, dst, size, tmpData)
	}
	if tmpAddr.width/8 < sizePtr() {
		ice("register width smaller than ptr size ")
	}
	if size > tmpData.width/8 && size%(tmpData.width/8) != 0 {
		ice(fmt.Sprintf("Invalid size (%v), reg width/8 (%v)", size, tmpAddr.width/8))
	}
	addrdatatype := OpDataType{OP_DATA,
		InstrData{signed: false, size: sizePtr()}, XMM_INVALID}

	mov := GetInstr(I_MOV, datatype)
	movaddr := GetInstr(I_MOV, addrdatatype)
	asm := ""
	var count uint
	var chunk uint
	if size <= tmpData.width/8 {
		count = 1
		chunk = size
	} else {
		chunk = (tmpData.width / 8)
		count = size / chunk
	}
	if count == 0 {
		return asm
	}
	asm += instrRegReg(movaddr, src, tmpAddr, false)
	for i := uint(0); i < count; i++ {
		asm += instrMemReg(mov, "", 0, tmpAddr, tmpData, false)
		asm += instrRegMem(mov, tmpData, dst, dstName, dstOffset, false)
		dstOffset += int(chunk)
		if i < count-1 {
			asm += AddImm32Reg(uint32(chunk), tmpAddr, false)
		}
	}
	return asm
}

func MovRegIndirectMemSmall(datatype OpDataType, src *register, dstName string, dstOffset int, dst *register, size uint, tmpData *register) string {
	if size > tmpData.width/8 && size%(tmpData.width/8) != 0 {
		ice(fmt.Sprintf("Invalid size (%v), reg width/8 (%v)", size, src.size()))
	}
	mov := GetInstr(I_MOV, datatype)
	asm := ""
	if size > tmpData.width/8 {
		ice("size too large")
	}
	if size == 0 {
		return asm
	}
	asm += instrMemReg(mov, "", 0, src, tmpData, false)
	asm += instrRegMem(mov, tmpData, dst, dstName, dstOffset, false)
	return asm
}

func MovMemMem(optype InstrOpType, srcName string, srcOffset int, src *register, dstName string, dstOffset int, dst *register, size uint, tmp *register) string {
	if src.width != 64 || dst.width != 64 {
		ice("Invalid register width")
	}
	if size%8 != 0 {
		ice("Invalid size")
	}
	var mov Instruction
	if optype == OP_PACKED {
		mov = GetInstr(I_PMOV, OpDataType{OP_PACKED, InstrData{16, false}, XMM_F64})
	} else {
		mov = GetInstr(I_MOV, GetIntegerOpDataType(false, size))
	}
	asm := instrMemReg(mov, srcName, srcOffset, src, tmp, false)
	asm += instrRegMem(mov, tmp, dst, dstName, dstOffset, false)
	return asm
}

func MovMemReg(datatype OpDataType, srcName string, srcOffset int, src, dst *register, spill bool) string {
	var mov Instruction
	if datatype.op == OP_PACKED {
		mov = GetInstr(I_PMOV, datatype)
	} else {
		mov = GetInstr(I_MOV, datatype)
	}
	return instrMemReg(mov, srcName, srcOffset, src, dst, spill)
}

func MovMemIndirectMem(datatype OpDataType, srcName string, srcOffset int, src *register, dstName string, dstOffset int, dst *register, size uint, tmpAddr, tmpData *register) string {
	if tmpAddr.width/8 < sizePtr() {
		ice("register width smaller than ptr size ")
	}
	if size > tmpData.width/8 && size%(tmpData.width/8) != 0 {
		ice(fmt.Sprintf("Invalid size (%v), reg width/8 (%v)", size, tmpAddr.width/8))
	}
	addrdatatype := OpDataType{OP_DATA,
		InstrData{signed: false, size: sizePtr()}, XMM_INVALID}

	mov := GetInstr(I_MOV, datatype)
	movaddr := GetInstr(I_MOV, addrdatatype)
	asm := ""
	var count uint
	var chunk uint
	if size <= tmpData.width/8 {
		count = 1
		chunk = size
	} else {
		chunk = (tmpData.width / 8)
		count = size / chunk
	}
	if count == 0 {
		return asm
	}
	asm += instrMemReg(movaddr, srcName, srcOffset, src, tmpAddr, false)
	for i := uint(0); i < count; i++ {
		asm += instrMemReg(mov, "", 0, tmpAddr, tmpData, false)
		asm += instrRegMem(mov, tmpData, dst, dstName, dstOffset, false)
		dstOffset += int(chunk)
		if i < count-1 {
			asm += AddImm32Reg(uint32(chunk), tmpAddr, false)
		}
	}
	return asm
}

func MovIntegerSimdMemIndirectMem(datatype OpDataType, srcName string, srcOffset int, src *register, dstName string, dstOffset int, dst *register, size uint, tmp1, tmp2 *register) string {
	if tmp1.width/8 < sizePtr() {
		ice("register width smaller than ptr size ")
	}
	if size > tmp2.width/8 && size%(tmp2.width/8) != 0 {
		ice(fmt.Sprintf("Invalid size (%v), reg width/8 (%v)", size, tmp1.width/8))
	}
	addrdatatype := OpDataType{OP_DATA,
		InstrData{signed: false, size: sizePtr()}, XMM_INVALID}

	mov := GetInstr(I_PMOV, datatype)
	movaddr := GetInstr(I_MOV, addrdatatype)
	asm := ""
	var count uint
	var chunk uint
	if size <= tmp2.width/8 {
		count = 1
		chunk = size
	} else {
		chunk = (tmp2.width / 8)
		count = size / chunk
	}
	if count == 0 {
		return asm
	}
	asm += instrMemReg(movaddr, srcName, srcOffset, src, tmp1, false)
	for i := uint(0); i < count; i++ {
		asm += instrMemReg(mov, "", 0, tmp1, tmp2, false)
		asm += instrRegMem(mov, tmp2, dst, dstName, dstOffset, false)
		dstOffset += int(chunk)
		if i < count-1 {
			asm += AddImm32Reg(uint32(chunk), tmp1, false)
		}
	}
	return asm
}

// asmMovImmReg moves imm to reg after converting it to int8/16/32 if size=1/2/4.
func MovImmReg(imm int64, size uint, dst *register, spill bool) string {
	if dst.width < 8*size {
		ice("Invalid register width")
	}
	switch size {
	case 1:
		imm = int64(int8(imm))
	case 2:
		imm = int64(int16(imm))
	case 4:
		imm = int64(int32(imm))
	}
	data := OpDataType{OP_DATA, InstrData{signed: true, size: size}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	return instrImmReg(mov, imm, size, dst, spill)
}

func MovImmFloatReg(f64 float64, isf32 bool, tmp, dst *register, spill bool) string {
	data := OpDataType{OP_DATA, InstrData{signed: false, size: tmp.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	return instrImmFloatReg(mov, f64, isf32, dst, tmp, spill)
}

func MovImmf32Reg(f32 float32, tmp, dst *register, spill bool) string {
	return MovImmFloatReg(float64(f32), true, tmp, dst, spill)
}

func MovImmf64Reg(f64 float64, tmp, dst *register, spill bool) string {
	return MovImmFloatReg(f64, false, tmp, dst, spill)
}

func MovImm8Reg(imm8 int8, dst *register, spill bool) string {
	return MovImmReg(int64(imm8), 1, dst, spill)
}

func MovImm32Reg(imm32 int32, dst *register, spill bool) string {
	return MovImmReg(int64(imm32), 4, dst, spill)

}

// CMovCCRegReg conditionally moves the src reg to the dst reg if the carry
// flag is clear (ie the previous compare had its src greater than its sub reg).
func CMovCCRegReg(src, dst *register, size uint, spill bool) string {
	var cmov Instruction
	if size == 1 {
		// there is conditional byte move
		cmov = CMOVWCC
	} else if size == 2 {
		cmov = CMOVWCC
	} else if size == 4 {
		cmov = CMOVLCC
	} else if size == 8 {
		cmov = CMOVQCC
	}
	return instrRegReg(cmov, src, dst, spill)
}

func Lea(srcName string, srcOffset int, src, dst *register, spill bool) string {
	if src.width != dst.width {
		ice("Invalid register width")
	}
	data := OpDataType{OP_DATA, InstrData{signed: false, size: src.width / 8}, XMM_INVALID}
	lea := GetInstr(I_LEA, data)
	return instrMemReg(lea, srcName, srcOffset, src, dst, spill)
}

func AddImm32Reg(imm32 uint32, dst *register, spill bool) string {
	if dst.width < 32 {
		ice("Invalid register width")
	}
	return instrImmReg(ADDQ, int64(imm32), 8, dst, spill)
}

func SubImm32Reg(imm32 uint32, dst *register, spill bool) string {
	if dst.width < 32 {
		ice("Invalid register width")
	}
	return instrImmReg(SUBQ, int64(imm32), 8, dst, spill)
}

func AddRegReg(datatype OpDataType, src, dst *register, spill bool) string {
	if dst.width != src.width {
		ice("Invalid register width")
	}
	return instrRegReg(GetInstr(I_ADD, datatype), src, dst, spill)
}

func SubRegReg(datatype OpDataType, src, dst *register, spill bool) string {
	if dst.width != src.width {
		ice("Invalid register width")
	}
	return instrRegReg(GetInstr(I_SUB, datatype), src, dst, spill)
}

func MulImm32RegReg(imm32 uint32, src, dst *register, spill bool) string {
	if dst.width < 32 {
		ice("Invalid register width")
	}
	return instrImmRegReg(IMUL3Q, int64(imm32), 8, src, dst, spill)
}

// MulRegReg multiplies the src register by the dst register and stores
// the result in the dst register. Overflow is discarded
func MulRegReg(datatype OpDataType, src, dst *register, spill bool) string {
	if dst.width != src.width {
		ice("Invalid register width")
	}
	asm := ""
	if datatype.op == OP_DATA {
		rax := getRegister(REG_AX)
		rdx := getRegister(REG_DX)
		if rax.width != 64 || rdx.width != 64 {
			ice("Invalid rax or rdx register width")
		}
		// rax is the implicit destination for MULQ
		asm = MovRegReg(datatype, dst, rax, false)
		// the low order part of the result is stored in rax and the high order part
		// is stored in rdx
		var tinstr InstructionType
		if datatype.signed {
			tinstr = I_IMUL
		} else {
			tinstr = I_MUL
		}
		mul := GetInstr(tinstr, datatype)
		asm += instrReg(mul, src, false)
		asm += MovRegReg(datatype, rax, dst, spill)
		return asm
	}
	if datatype.op == OP_XMM {
		asm += instrRegReg(GetInstr(I_MUL, datatype), src, dst, spill)
	}
	return asm
}

// DivRegReg divides the "dividend" register by the "divisor" register and stores
// the quotient in rax and the remainder in rdx. DivRegReg is only for integer division.
func DivRegReg(signed bool, datatype InstrOpType, dividend, divisor *register, size uint) (asm string, rax *register, rdx *register) {
	if dividend.width != divisor.width || divisor.width < size*8 {
		ice("Invalid register width for DivRegReg")
	}
	if datatype != OP_DATA {
		ice("Unsupported arithmetic data type")
	}
	rax = getRegister(REG_AX)
	if size > 1 {
		rdx = getRegister(REG_DX)
	}
	if rax.width != 64 || (size > 1 && rdx.width != 64) {
		ice("Invalid rax or rdx register width")
	}

	// rdx:rax are the upper and lower parts of the dividend respectively,
	// and rdx:rax are the implicit destination of DIVQ
	asm = ""
	asm += ZeroReg(rax)
	if size > 1 {
		asm += ZeroReg(rdx)
	}
	dt := OpDataType{datatype, InstrData{signed: signed, size: size}, XMM_INVALID}
	asm += MovRegReg(dt, dividend, rax, false)

	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr InstructionType
	if signed {
		tinstr = I_IDIV
	} else {
		tinstr = I_DIV
	}

	asm += instrReg(GetInstr(tinstr, dt), divisor, false)
	return asm, rax, rdx
}

// DivFloatRegReg performs floating point division by dividing the dividend register
// by the divisor register and stores the quotient in the dividend register
func DivFloatRegReg(datatype OpDataType, dividend, divisor *register, spill bool) string {
	if dividend.width != divisor.width {
		ice("Invalid register width")
	}
	if datatype.op != OP_XMM {
		ice("Unsupported data type for floating point division")
	}
	return instrRegReg(GetInstr(I_DIV, datatype), divisor, dividend, true)
}

func ArithOp(datatype OpDataType, op token.Token, x, y, result *register) string {
	if x.width != y.width || x.width != result.width {
		ice("Invalid register width")
	}
	asm := ""
	switch op {
	default:
		ice(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.ADD:
		asm += MovRegReg(datatype, x, result, false)
		asm += AddRegReg(datatype, y, result, false)
	case token.SUB:
		asm += MovRegReg(datatype, x, result, false)
		asm += SubRegReg(datatype, y, result, false)
	case token.MUL:
		asm += MovRegReg(datatype, x, result, false)
		asm += MulRegReg(datatype, y, result, false)
	case token.QUO, token.REM:
		if datatype.op == OP_DATA {
			// the quotient is stored in rax and
			// the remainder is stored in rdx.
			var rax, rdx *register
			a, rax, rdx := DivRegReg(datatype.signed, datatype.op, x, y, datatype.size)
			asm += a
			if op == token.QUO {
				asm += MovRegReg(datatype, rax, result, false)
			} else {
				asm += MovRegReg(datatype, rdx, result, false)
			}
		} else {
			// assume quotient operation,
			// since modulus isn't defined for floats
			asm += MovRegReg(datatype, x, result, false)
			asm += DivFloatRegReg(datatype, result, y, false)
		}
	}
	return asm
}

// AndRegReg AND's the src register by the dst register and stores
// the result in the dst register.
func AndRegReg(src, dst *register, size uint, spill bool) string {
	if src.width != dst.width {
		ice("Invalid register width")
	}
	dt := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	and := GetInstr(I_AND, dt)
	return instrRegReg(and, src, dst, spill)
}

func OrRegReg(src, dst *register, spill bool) string {
	if src.width != dst.width {
		ice("Invalid register width")
	}
	dt := OpDataType{OP_DATA, InstrData{signed: false, size: src.width / 8}, XMM_INVALID}
	return instrRegReg(GetInstr(I_OR, dt), src, dst, spill)
}

func XorRegReg(src, dst *register, spill bool) string {
	if src.width != dst.width {
		ice("Invalid register width")
	}
	dt := OpDataType{OP_DATA, InstrData{signed: false, size: src.width / 8}, XMM_INVALID}
	return instrRegReg(GetInstr(I_XOR, dt), src, dst, spill)
}

func XorImm32Reg(imm32 int32, dst *register, size uint, spill bool) string {
	data := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	return instrImmReg(xor, int64(imm32), size, dst, spill)
}

func XorImm64Reg(imm64 int64, dst *register, size uint, spill bool) string {
	data := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	return instrImmReg(xor, imm64, size, dst, spill)
}

func NotReg(reg *register, size uint, spill bool) string {
	return XorImm32Reg(-1, reg, size, spill)
}

const (
	SHIFT_LEFT = iota
	SHIFT_RIGHT
)

func MovZeroExtend(src, dst *register, srcSize, dstSize uint, spill bool) string {
	var movzx InstructionType
	switch srcSize {
	default:
		ice(fmt.Sprintf("Internal error, bad zero extend size (%v)", srcSize))
	case 1:
		movzx = I_MOVBZX
	case 2:
		movzx = I_MOVWZX
	case 4:
		movzx = I_MOVLZX
	}
	if dstSize <= srcSize || (dstSize != 1 && dstSize != 2 && dstSize != 4 && dstSize != 8) {
		ice(fmt.Sprintf("Bad dstSize (%v) for zero extend result", dstSize))
	}
	dt := OpDataType{OP_DATA, InstrData{signed: false, size: dstSize}, XMM_INVALID}
	return instrRegReg(GetInstr(movzx, dt), src, dst, spill)
}

func MovSignExtend(src, dst *register, srcSize, dstSize uint, spill bool) string {
	var movsx InstructionType
	switch srcSize {
	default:
		ice(fmt.Sprintf("Internal error, bad sign extend size (%v)", srcSize))
	case 1:
		movsx = I_MOVBSX
	case 2:
		movsx = I_MOVWSX
	case 4:
		movsx = I_MOVLSX
	}
	dt := OpDataType{OP_DATA, InstrData{signed: false, size: dstSize}, XMM_INVALID}
	asm := instrRegReg(GetInstr(movsx, dt), src, dst, spill)
	return asm
}

// ShiftRegReg shifts src by shiftReg amount and stores the result in src.
// The tmp reg is used for intermediates (if shifting right 64 times then SHR
// isn't used directly)
func ShiftRegReg(signed bool, direction int, src, shiftReg, tmp *register, size uint, spill bool) string {
	cl := getRegister(REG_CL)
	cx := getRegister(REG_CX)
	regCl := cx

	var opcode InstructionType
	if direction == SHIFT_LEFT {
		opcode = I_SHL
	} else if !signed && direction == SHIFT_RIGHT {
		opcode = I_SHR
	} else if signed && direction == SHIFT_RIGHT {
		opcode = I_SAR
	}
	data := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	shift := GetInstr(opcode, data)
	asm := ""

	if size == 1 {
		regCl = cl
	}

	maxShift := 8 * uint32(size)
	completeShift := int32(maxShift)
	// the shl/shr instructions mast the shift count to either
	// 5 or 6 bits (5 if not operating on a 64bit value)
	if completeShift == 32 || completeShift == 64 {
		completeShift--
	}
	instrdata := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	asm += MovRegReg(instrdata, shiftReg, cx, false)

	asm += MovImm32Reg(completeShift, tmp, false)
	// compare only first byte of shift reg,
	// since useful shift can be at most 64
	asm += CmpRegImm32(shiftReg, maxShift, 1)
	asm += CMovCCRegReg(tmp, cx, size, false)

	var zerosize uint = 1
	asm += MovZeroExtend(cl, cx, zerosize, cx.width/8, false)
	asm += instrRegReg(shift, regCl, src, spill)

	if maxShift == 64 || maxShift == 32 {
		asm += MovImm32Reg(1, tmp, false)
		// compare only first byte of shift reg,
		// since useful shift can be at most 64
		asm += XorRegReg(cx, cx, false)
		asm += CmpRegImm32(shiftReg, maxShift, 1)
		asm += CMovCCRegReg(tmp, cx, size, false)
		var zerosize uint = 1
		asm += MovZeroExtend(cl, cx, zerosize, cx.width/8, false)
		asm += instrRegReg(shift, regCl, src, spill)
	}
	return asm
}

func ShiftImm8Reg(signed bool, direction int, count uint8, reg *register) string {
	var opcode InstructionType
	if direction == SHIFT_LEFT {
		opcode = I_SHL
	} else if !signed && direction == SHIFT_RIGHT {
		opcode = I_SHR
	} else if signed && direction == SHIFT_RIGHT {
		opcode = I_SAR
	}
	data := OpDataType{OP_DATA, InstrData{signed: false, size: reg.size()}, XMM_INVALID}
	shift := GetInstr(opcode, data)
	asm := ""
	// the shl/shr instructions mast the shift count to either
	// 5 or 6 bits (5 if not operating on a 64bit value)
	if (reg.width == 32 && count >= 32) || (reg.width == 64 && count >= 64) {
		asm += ZeroReg(reg)
	} else {
		asm += fmt.Sprintf("%-9v    $%v, %v\n", shift, count, reg.name)
	}
	return asm
}

func AndNotRegReg(src, dst *register, size uint, spill bool) string {
	if src.width != dst.width {
		ice("Invalid register width")
	}
	asm := NotReg(dst, size, spill)
	asm += AndRegReg(src, dst, size, spill)
	return asm
}

func BitwiseOp(op token.Token, signed bool, x, y, result *register, size uint) string {
	if x.width != y.width || x.width != result.width {
		ice("Invalid register width")
	}
	asm := ""
	instrdata := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	switch op {
	default:
		ice(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.AND:
		asm = MovRegReg(instrdata, y, result, false)
		asm += AndRegReg(x, result, size, false)
	case token.OR:
		asm = MovRegReg(instrdata, y, result, false)
		asm += OrRegReg(x, result, false)
	case token.XOR:
		asm = MovRegReg(instrdata, y, result, false)
		asm += XorRegReg(x, result, false)
	case token.SHL:
		asm = MovRegReg(instrdata, x, result, false)
		tmp := x
		asm += ShiftRegReg(signed, SHIFT_LEFT, result, y, tmp, size, false)
	case token.SHR:
		asm = MovRegReg(instrdata, x, result, false)
		tmp := x
		asm += ShiftRegReg(signed, SHIFT_RIGHT, result, y, tmp, size, false)
	case token.AND_NOT:
		asm = MovRegReg(instrdata, y, result, false)
		asm += AndNotRegReg(x, result, size, false)
	}
	return asm
}

func CmpRegReg(odt OpDataType, x, y *register) string {
	if x.width != y.width {
		ice("Invalid register width")
	}
	return instrRegReg(GetInstr(I_CMP, odt), x, y, false)
}

func CmpRegImm32(r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		ice("Invalid register width")
	}
	data := OpDataType{OP_DATA, InstrData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data)
	return instrRegImm32(cmp, r, imm32, size, false)
}

// CmpOp compares x to y, storing the op comparison flag (EQ, NEQ, ...) in result
func CmpOp(data OpDataType, op token.Token, x *register, y *register, result *register) string {
	if x.width != y.width {
		ice(fmt.Sprintf("Invalid register width, x.width (%v), y.width (%v), result.width (%v)", x.width, y.width, result.width))
	}
	asm := ""
	asm += CmpRegReg(data, x, y)
	switch op {
	default:
		ice(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.EQL:
		asm += instrReg(SETEQ, result, false)
	case token.NEQ:
		asm += instrReg(SETNE, result, false)
	case token.LEQ:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == OP_XMM {
			asm += instrReg(SETCC, result, false)
		} else {
			if data.signed {
				asm += instrReg(SETLE, result, false)
			} else {
				asm += instrReg(SETLS, result, false)
			}
		}
	case token.GEQ:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == OP_XMM {
			asm += instrReg(SETLS, result, false)
		} else {
			if data.signed {
				asm += instrReg(SETGE, result, false)
			} else {
				asm += instrReg(SETCC, result, false)
			}
		}
	case token.LSS:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == OP_XMM {
			asm += instrReg(SETHI, result, false)
		} else {
			if data.signed {
				asm += instrReg(SETLT, result, false)
			} else {
				asm += instrReg(SETCS, result, false)
			}
		}
	case token.GTR:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == OP_XMM {
			asm += instrReg(SETCS, result, false)
		} else {
			if data.signed {
				asm += instrReg(SETGT, result, false)
			} else {
				asm += instrReg(SETHI, result, false)
			}
		}
	}
	return asm
}

func isIntegerOp(datatype OpDataType) bool {
	return datatype.op == OP_DATA
}

func isFloatOp(datatype OpDataType) bool {
	return datatype.op == OP_XMM
}

func ConvertOp(from *register, fromtype OpDataType, to *register, totype OpDataType, tmp *register) string {
	if isIntegerOp(fromtype) && isIntegerOp(totype) {
		return IntegerToInteger(from, to, fromtype, totype)
	} else if isIntegerOp(fromtype) && isFloatOp(totype) {
		return IntegerToFloat(from, to, fromtype, totype, tmp)
	} else if isFloatOp(fromtype) && isIntegerOp(totype) {
		return FloatToInteger(from, to, fromtype, totype)
	} else if isFloatOp(fromtype) && isFloatOp(totype) {
		return FloatToFloat(from, to, fromtype, totype)
	} else {
		ice(fmt.Sprintf("Internal error, converting betwen type %v and %v", fromtype.op, totype.op))
	}
	return ""
}

func IntegerToInteger(from, to *register, ftype, totype OpDataType) string {
	if ftype.size < totype.size {
		if ftype.signed {
			// From Go Spec:
			// If the value is a signed integer, it is sign extended to implicit infinite precision;
			// It is then truncated to fit in the result type's size.
			return MovSignExtend(from, to, ftype.size, totype.size, false)
		} else {
			return MovZeroExtend(from, to, ftype.size, totype.size, false)
		}
	}
	return MovRegReg(totype, from, to, false)
}

func XmmInstrDataSize(xmmtype XmmData) uint {
	if xmmtype == XMM_F32 {
		return 4
	} else if xmmtype == XMM_F64 {
		return 8
	} else if xmmtype == XMM_4X_F32 {
		return 16
	} else if xmmtype == XMM_2X_F64 {
		return 16
	}
	ice("Internal error getting floating point instr size")
	return uint(math.MaxUint64)
}

func IntegerToFloat(from, to *register, ftype, totype OpDataType, tmp *register) string {
	tosize := XmmInstrDataSize(totype.xmmvariant)
	fromsize := ftype.size
	fromreg := from
	// no direct conversion from int8/int16 to float32/float64
	if ftype.size < 4 {
		fromsize = 4
		fromreg = tmp
		if ftype.signed {
			MovSignExtend(from, tmp, ftype.size, fromsize, false)
		} else {
			MovZeroExtend(from, tmp, ftype.size, fromsize, false)
		}
	} else if ftype.size == 4 && !ftype.signed {
		fromsize = 8
		fromreg = tmp
		MovZeroExtend(from, tmp, ftype.size, fromsize, false)
	}
	cvt := GetConvertInstruction(I_CVT_INT2FLOAT, fromsize, tosize)
	return instrRegReg(cvt, fromreg, to, false)
}

func FloatToInteger(from, to *register, ftype, totype OpDataType) string {
	// From Go Spec:
	// When converting a floating-point number to an integer,
	// the fraction is discarded (truncation towards zero).
	fromsize := XmmInstrDataSize(ftype.xmmvariant)
	cvt := GetConvertInstruction(I_CVT_FLOAT2INT, fromsize, totype.size)
	return instrRegReg(cvt, from, to, false)
}

func FloatToFloat(from, to *register, ftype, totype OpDataType) string {
	fromsize := XmmInstrDataSize(ftype.xmmvariant)
	cvt := GetConvertInstruction(I_CVT_FLOAT2FLOAT, fromsize, totype.size)
	return instrRegReg(cvt, from, to, false)
}

func Ret() string {
	return fmt.Sprintf("RET\n")
}
