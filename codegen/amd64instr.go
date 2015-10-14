package codegen

import (
	"fmt"
	"go/token"
	"strings"
)

type InstrOpType int

const (
	_          InstrOpType = iota
	INTEGER_OP             // int8/uint8,,..., int64/uint64
	XMM_OP                 // f32/f64, packed f32, packed f64
)

type InstrDataType struct {
	op InstrOpType
	NonXmmData
	xmm XmmData
}

type NonXmmData struct {
	size   uint
	signed bool
}

type XmmData int

const (
	XMM_INVALID XmmData = iota
	SINGLE_F32
	SINGLE_F64
	DOUBLE_F32
	DOUBLE_F64
)

type Instr int
type Instruction struct {
	TInstr InstructionType

	// integer forms
	ByteSized  Instr
	WordSized  Instr
	LongSized  Instr
	QuadSized  Instr
	DQuadSized Instr

	// xmm forms
	/*SingleF32 Instr
	SingleF64 Instr
	PackedF32 Instr // operates on four packed f32
	PackedF64 Instr // operate on two packed f64*/

}

type XmmInstruction struct {
	XmmInstr  InstructionType
	SingleF32 Instr
	SingleF64 Instr
	PackedF32 Instr // operates on four packed f32
	PackedF64 Instr // operate on two packed f64
}

func (inst Instruction) GetSized(size uint) Instr {
	switch size {
	case 1:
		return inst.ByteSized
	case 2:
		return inst.WordSized
	case 4:
		return inst.LongSized
	case 8:
		return inst.QuadSized
	}
	msg := fmt.Sprintf("Invalid size(%v), for instr (%v)", size, inst.TInstr.String())
	panic(msg)
}

type InstructionType int

const (
	I_ADD InstructionType = iota
	I_SUB
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

	I_XOR
	I_LEA
	I_IMUL
	I_MUL
	I_IDIV
	I_DIV
	I_AND
	I_OR
	I_SHL
	I_SHR
	I_SAL
	I_SAR
	I_CMP
)

func (tinst InstructionType) String() string {
	switch tinst {
	default:
		panic("Unknown TIinstruction")
	case I_ADD:
		return "ADD"
	case I_SUB:
		return "SUB"
	case I_MOV:
		return "MOV"
	case I_XOR:
		return "XOR"
	case I_LEA:
		return "LEA"
	case I_MUL:
		return "MUL"
	case I_IMUL:
		return "IMUL"
	case I_DIV:
		return "DIV"
	case I_IDIV:
		return "IDIV"
	case I_AND:
		return "AND"
	case I_OR:
		return "OR"
	case I_SHL:
		return "SHL"
	case I_SHR:
		return "SHR"
	case I_SAL:
		return "SAL"
	case I_SAR:
		return "SAR"
	case I_CMP:
		return "CMP"
	}
}

var Insts = []Instruction{
	{I_ADD, ADDB, ADDW, ADDL, ADDQ, NONE},
	{I_SUB, SUBB, SUBW, SUBL, SUBQ, NONE},
	{I_MOV, MOVB, MOVW, MOVL, MOVQ, NONE},

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
	{I_ADD, ADDSS, ADDSD, ADDPS, ADDPD},
	{I_SUB, SUBSS, SUBSD, SUBPS, SUBPD},
	{I_MOV, MOVSS, MOVSD, MOVUPS, MOVUPD},

	{I_MUL, MULSS, MULSD, MULPS, MULPD},

	{I_DIV, DIVSS, DIVSD, DIVPS, DIVPD},

	{I_CMP, CMPSS, CMPSD, CMPPS, CMPPD},
}

func GetInstruction(tinst InstructionType) Instruction {
	for _, inst := range Insts {
		if inst.TInstr == tinst {
			return inst
		}
	}
	panic("Couldn't get instruction")
}

// GetInstr, the size is in bytes
func GetInstr(tinst InstructionType, datatype InstrDataType) Instr {
	if datatype.op == INTEGER_OP {
		return GetInstruction(tinst).GetSized(datatype.size)
	} else {
		panic("XMM INSTR")
	}
}

// asmZeroMemory generates "MOVQ $0, name+offset(REG)" instructions,
// size is in bytes
func asmZeroMemory(indent string, name string, offset int, size uint, reg *register) string {

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
	datatype := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: chunk}, XMM_INVALID}
	mov := GetInstr(I_MOV, datatype).String()

	for i := uint(0); i < size/chunk; i++ {
		ioffset := int(i*chunk) + offset
		asm += indent
		asm += fmt.Sprintf("%v    $0, %v+%v(%v)\n", mov, name, ioffset, reg.name)
	}

	return strings.Replace(asm, "+-", "-", -1)
}

// asmZeroReg generates "XORQ reg, reg" instructions
func asmZeroReg(indent string, reg *register) string {
	var datatype InstrDataType

	if reg.typ == XmmReg {
		datatype = InstrDataType{XMM_OP, NonXmmData{}, DOUBLE_F64}

	} else {
		datatype = InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: reg.width / 8}, XMM_INVALID}
	}

	xor := GetInstr(I_XOR, datatype)
	return indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), reg.name, reg.name)
}

func asmMovRegReg(indent string, datatype InstrOpType, srcReg *register, dstReg *register, size uint) string {
	if srcReg.width != dstReg.width || size*8 > srcReg.width {
		panic(fmt.Sprintf("(%v) srcReg.width != (%v) dstReg.width or invalid size in asmMoveRegToReg", srcReg.width, dstReg.width))
	}
	if datatype != INTEGER_OP {
		panic("Unsupported arithmetic data type")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	mov := GetInstr(I_MOV, data).String()
	asm := indent + fmt.Sprintf("%v    %v, %v\n", mov, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovRegMem(indent string, datatype InstrOpType, srcReg *register, dstName string, dstReg *register, dstOffset int, size uint) string {
	if srcReg.width < size*8 {
		panic("srcReg.width < size * 8")
	}
	if size == 0 {
		panic("size == 0")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("// BEGIN asmMovRegMem, size (%v), mov (%v), mov.String (%v)\n", size, mov, mov.String())
	asm += indent + fmt.Sprintf("%v    %v, %v+%v(%v)\n", mov.String(), srcReg.name, dstName, dstOffset, dstReg.name)
	asm += indent + fmt.Sprintf("// END asmMovRegMem, size (%v), mov (%v), mov.String (%v)\n", size, mov, mov.String())
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovRegMemIndirect(indent string, srcReg *register, dstName string, dstReg *register, dstOffset int, tmp *register) string {
	if tmp.width != srcReg.width {
		panic("Invalid register width for asmMovRegMemIndirect")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: srcReg.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("%v    %v, (%v)\n", mov.String(), srcReg.name, tmp.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemMem(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMemMem")
	}
	if size%8 != 0 {
		panic("Invalid size for asmMovMemMem")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v+%v(%v)\n", srcName, srcOffset, srcReg.name, dstName, dstOffset, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemReg(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, size uint, dstReg *register) string {
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemMemIndirect(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMovMemMemIndirect")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    %v+%v(%v), (%v)\n", srcName, srcOffset, srcReg.name, tmp)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemIndirectReg(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstReg *register, tmp *register) string {
	if dstReg.width != tmp.width {
		panic("Invalid register width for asmMovMemIndirectReg")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: dstReg.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent
	asm += fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg, tmp)

	asm += indent
	asm += fmt.Sprintf("%v    (%v), %v\n", mov.String(), tmp, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemIndirectMem(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint, tmp1 *register, tmp2 *register) string {
	if tmp1.width != tmp2.width {
		panic("Mismatched register widths in asmMovMemIndirectMem")
	}
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMovMemIndirectMem")
	}
	if size%(tmp1.width/8) != 0 {
		panic("Invalid size in asmMovMemIndirectMem")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: tmp1.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := ""
	for i := uint(0); i < size/(tmp1.width/8); i++ {

		asm += indent
		asm += fmt.Sprintf("%v    %v+%v(%v), %v\n",
			mov.String(), srcName, srcOffset, srcReg.name, tmp1.name)

		asm += indent
		asm += fmt.Sprintf("%v    (%v), %v\n",
			mov.String(), tmp1.name, tmp2.name)

		asm += indent
		asm += fmt.Sprintf("%v    %v, %v+%v(%v)\n",
			mov.String(), tmp2.name, dstName, dstOffset, dstReg.name)

		srcOffset += int((tmp1.width / 8))
		dstOffset += int((tmp1.width / 8))
	}
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemIndirectMemIndirect(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp1 *register, tmp2 *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: tmp1.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg, tmp1)
	asm += indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), dstName, dstOffset, dstReg, tmp2)
	asm += indent + fmt.Sprintf("%v    (%v), (%v)\n", mov.String(), tmp1.name, tmp2.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovImm8Reg(indent string, imm8 int8, dstReg *register) string {
	if dstReg.width < 16 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("MOVB    $%v, %v\n", imm8, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovImm16Reg(indent string, imm16 int16, dstReg *register) string {
	if dstReg.width < 16 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("MOVW    $%v, %v\n", imm16, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovImm32Reg(indent string, imm32 int32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("MOVL    $%v, %v\n", imm32, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovImm64Reg(indent string, imm64 int64, dstReg *register) string {
	if dstReg.width != 64 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("MOVQ    $%v, %v\n", imm64, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

// asmCMovCCRegReg conditionally moves the src reg to the dst reg if the carry
// flag is clear (ie the previous compare had its src greater than its sub reg).
func asmCMovCCRegReg(indent string, src *register, dst *register, size uint) string {
	var cmov string
	if size == 1 {
		// there is conditional byte move
		cmov = "CMOVWCC"
	} else if size == 2 {
		cmov = "CMOVWCC"
	} else if size == 4 {
		cmov = "CMOVLCC"
	} else if size == 8 {
		cmov = "CMOVQCC"
	}
	asm := indent + fmt.Sprintf("%v %v, %v\n", cmov, src.name, dst.name)
	return asm
}

func asmLea(indent string, srcName string, srcOffset int, srcReg *register, dstReg *register) string {
	if srcReg.width != dstReg.width {
		panic("Invalid register width for asmLea")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: srcReg.width / 8}, XMM_INVALID}
	lea := GetInstr(I_LEA, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", lea.String(), srcName, srcOffset, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmAddImm32Reg(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmAddImm32Reg")
	}
	asm := indent + fmt.Sprintf("ADDQ    $%v, %v\n", imm32, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmSubImm32Reg(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmSubImm32Reg")
	}
	asm := indent + fmt.Sprintf("SUBQ    $%v, %v\n", imm32, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmAddRegReg(indent string, datatype InstrOpType, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width for asmAddRegReg")
	}
	if datatype != INTEGER_OP {
		panic("Unsupported arithmetic data type")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: srcReg.width / 8}, XMM_INVALID}
	add := GetInstr(I_ADD, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", add.String(), srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmSubRegReg(indent string, datatype InstrOpType, srcReg *register, dstReg *register, size uint) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width for asmSubRegReg")
	}
	if datatype != INTEGER_OP {
		panic("Unsupported arithmetic data type")
	}

	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	sub := GetInstr(I_SUB, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", sub.String(), srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMulImm32RegReg(indent string, imm32 uint32, srcReg *register, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmMulImm32RegReg")
	}
	asm := indent + fmt.Sprintf("IMUL3Q    $%v, %v, %v\n", imm32, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

// asmMulRegReg multiplies the src register by the dst register and stores
// the result in the dst register. Overflow is discarded
func asmMulRegReg(indent string, signed bool, datatype InstrOpType, src *register, dst *register, size uint) string {

	if dst.width != 64 {
		panic("Invalid register width for asmMulRegReg")
	}
	if datatype != INTEGER_OP {
		panic("Unsupported arithmetic data type")
	}

	rax := getRegister(REG_AX)
	rdx := getRegister(REG_DX)
	if rax.width != 64 || rdx.width != 64 {
		panic("Invalid rax or rdx register width in asmMulRegReg")
	}

	// rax is the implicit destination for MULQ
	asm := asmMovRegReg(indent, datatype, dst, rax, size)
	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr InstructionType
	if signed {
		tinstr = I_IMUL
	} else {
		tinstr = I_MUL
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	mul := GetInstr(tinstr, data).String()
	asm += indent + fmt.Sprintf("%v    %v\n", mul, src.name)
	asm += asmMovRegReg(indent, datatype, rax, dst, size)
	return strings.Replace(asm, "+-", "-", -1)
}

// asmDivRegReg divides the "dividend" register by the "divisor" register and stores
// the quotient in rax and the remainder in rdx
func asmDivRegReg(indent string, signed bool, datatype InstrOpType, dividend *register, divisor *register, size uint) (asm string, rax *register, rdx *register) {

	if dividend.width != divisor.width || divisor.width < size*8 {
		panic("Invalid register width for asmDivRegReg")
	}
	if datatype != INTEGER_OP {
		panic("Unsupported arithmetic data type")
	}

	rax = getRegister(REG_AX)

	if size > 1 {
		rdx = getRegister(REG_DX)
	}
	if rax.width != 64 || (size > 1 && rdx.width != 64) {
		panic("Invalid rax or rdx register width")
	}

	// rdx:rax are the upper and lower parts of the dividend respectively,
	// and rdx:rax are the implicit destination of DIVQ
	asm = ""
	asm += asmZeroReg(indent, rax)
	if size > 1 {
		asm += asmZeroReg(indent, rdx)
	}

	asm += asmMovRegReg(indent, datatype, dividend, rax, size)

	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr InstructionType
	if signed {
		tinstr = I_IDIV
	} else {
		tinstr = I_DIV
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	div := GetInstr(tinstr, data).String()
	asm += indent + fmt.Sprintf("%v    %v\n", div, divisor.name)
	return asm, rax, rdx
}

func asmArithOp(indent string, signed bool, datatype InstrOpType, op token.Token, x *register, y *register, result *register, size uint) string {
	if x.width != 64 || y.width != 64 || result.width != 64 {
		panic("Invalid register width in asmArithOp")
	}
	asm := ""
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmArithOp", op))
	case token.ADD:
		asm += asmMovRegReg(indent, datatype, x, result, size)
		asm += asmAddRegReg(indent, datatype, y, result)
	case token.SUB:
		asm += asmMovRegReg(indent, datatype, x, result, size)
		asm += asmSubRegReg(indent, datatype, y, result, size)
	case token.MUL:
		asm += asmMovRegReg(indent, datatype, x, result, size)
		asm += asmMulRegReg(indent, signed, datatype, y, result, size)
	case token.QUO, token.REM:
		// the quotient is stored in rax and
		// the remainder is stored in rdx.
		var rax, rdx *register
		a, rax, rdx := asmDivRegReg(indent, signed, datatype, x, y, size)
		asm += a
		if op == token.QUO {
			asm += asmMovRegReg(indent, datatype, rax, result, size)
		} else {
			asm += asmMovRegReg(indent, datatype, rdx, result, size)
		}
	}
	return strings.Replace(asm, "+-", "-", -1)
}

// asmAndRegReg AND's the src register by the dst register and stores
// the result in the dst register.
func asmAndRegReg(indent string, src *register, dst *register, size uint) string {
	if src.width != dst.width {
		panic("Invalid register width for asmAndRegReg")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	and := GetInstr(I_AND, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", and.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmOrRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width for asmOrRegReg")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: src.width / 8}, XMM_INVALID}
	or := GetInstr(I_OR, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", or.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmXorRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width for asmXorRegReg")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: src.width / 8}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), src.name, dst.name)
	return asm
}

func asmXorImm32Reg(indent string, imm32 int32, dst *register, size uint) string {
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    $%v, %v\n", xor.String(), imm32, dst.name)
	return asm
}

func asmXorImm64Reg(indent string, imm64 int64, dst *register, size uint) string {
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    $%v, %v\n", xor.String(), imm64, dst.name)
	return asm
}

func asmNotReg(indent string, reg *register, size uint) string {
	return asmXorImm32Reg(indent, -1, reg, size)
}

const (
	SHIFT_LEFT = iota
	SHIFT_RIGHT
)

func asmMovZeroExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
	var opcode InstructionType
	switch srcSize {
	default:
		panic(fmt.Sprintf("Bad srcSize (%v)", srcSize))
	case 1:
		opcode = I_MOVBZX
	case 2:
		opcode = I_MOVWZX
	case 4:
		opcode = I_MOVWZX
	case 8:
		opcode = I_MOVLZX
	}

	if dstSize <= srcSize || (dstSize != 1 && dstSize != 2 && dstSize != 4 && dstSize != 8) {
		panic(fmt.Sprintf("Bad dstSize (%v) for zero extend result", dstSize))
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: dstSize}, XMM_INVALID}
	movzx := GetInstr(opcode, data)
	asm := indent + fmt.Sprintf("%v %v, %v\n", movzx.String(), src.name, dst.name)
	return asm
}

func asmMovSignExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
	var opcode InstructionType
	switch srcSize {
	default:
		panic(fmt.Sprintf("Bad src size (%v)", srcSize))
	case 1:
		opcode = I_MOVBSX
	case 2:
		opcode = I_MOVWSX
	case 4:
		opcode = I_MOVWSX
	case 8:
		opcode = I_MOVLSX
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: dstSize}, XMM_INVALID}
	movsx := GetInstr(opcode, data)
	if movsx == NONE {
		panic(fmt.Sprintf("Bad dstSize (%v) for sign extend result", dstSize))
	}
	asm := indent + fmt.Sprintf("%v    %v, %v\n", movsx.String(), src.name, dst.name)
	return asm
}

func asmShiftRegReg(indent string, signed bool, direction int, src *register, shiftReg *register, tmp *register, size uint) string {

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

	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
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

	asm += asmMovRegReg(indent, INTEGER_OP, shiftReg, cx, size)

	asm += asmMovImm32Reg(indent, completeShift, tmp)
	// compare only first byte of shift reg,
	// since useful shift can be at most 64
	asm += asmCmpRegImm32(indent, shiftReg, maxShift, 1)
	asm += asmCMovCCRegReg(indent, tmp, cx, size)

	var zerosize uint = 1
	asm += asmMovZeroExtend(indent, cl, cx, zerosize, cx.width/8)
	asm += indent + fmt.Sprintf("%v    %v, %v\n", shift.String(), regCl.name, src.name)

	if maxShift == 64 || maxShift == 32 {
		asm += asmMovImm32Reg(indent, 1, tmp)
		// compare only first byte of shift reg,
		// since useful shift can be at most 64
		asm += asmXorRegReg(indent, cx, cx)
		asm += asmCmpRegImm32(indent, shiftReg, maxShift, 1)
		asm += asmCMovCCRegReg(indent, tmp, cx, size)
		var zerosize uint = 1
		asm += asmMovZeroExtend(indent, cl, cx, zerosize, cx.width/8)
		asm += indent + fmt.Sprintf("%v    %v, %v\n", shift.String(), regCl.name, src.name)
	}

	return asm
}

func asmAndNotRegReg(indent string, src *register, dst *register, size uint) string {
	if src.width != dst.width {
		panic("Invalid register width for asmAndNotRegReg")
	}
	asm := asmNotReg(indent, dst, size)
	asm += asmAndRegReg(indent, src, dst, size)
	return asm
}

func asmBitwiseOp(indent string, op token.Token, signed bool, x *register, y *register, result *register, size uint) string {
	if x.width != y.width || x.width != result.width {
		panic("Invalid register width in asmBitwiseOp")
	}
	asm := ""
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmBitwiseOp", op))
	case token.AND:
		asm = asmMovRegReg(indent, INTEGER_OP, y, result, size)
		asm += asmAndRegReg(indent, x, result, size)
	case token.OR:
		asm = asmMovRegReg(indent, INTEGER_OP, y, result, size)
		asm += asmOrRegReg(indent, x, result)
	case token.XOR:
		asm = asmMovRegReg(indent, INTEGER_OP, y, result, size)
		asm += asmXorRegReg(indent, x, result)
	case token.SHL:
		asm = asmMovRegReg(indent, INTEGER_OP, x, result, size)
		tmp := x
		asm += asmShiftRegReg(indent, signed, SHIFT_LEFT, result, y, tmp, size)
	case token.SHR:
		asm = asmMovRegReg(indent, INTEGER_OP, x, result, size)
		tmp := x
		asm += asmShiftRegReg(indent, signed, SHIFT_RIGHT, result, y, tmp, size)
	case token.AND_NOT:
		asm = asmMovRegReg(indent, INTEGER_OP, y, result, size)
		asm += asmAndNotRegReg(indent, x, result, size)

	}
	return strings.Replace(asm, "+-", "-", -1)
}

func asmCmpRegReg(indent string, x *register, y *register, size uint) string {
	if x.width != y.width {
		panic("Invalid register width for asmCmpRegReg")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data).String()
	asm := fmt.Sprintf("%v	%v, %v\n", cmp, x.name, y.name)
	return strings.Replace(asm, "+-", "-", -1)

}

func asmCmpMemImm32(indent string, name string, offset int32, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		panic("Invalid register width for asmCmpMemImm32")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data)
	asm := indent + fmt.Sprintf("%v	%v+%v(%v), $%v\n", cmp, name, offset, r.name, imm32)
	return strings.Replace(asm, "+-", "-", -1)

}

func asmCmpRegImm32(indent string, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		panic("Invalid register width for asmCmpMemImm32")
	}
	data := InstrDataType{INTEGER_OP, NonXmmData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data).String()
	asm := indent + fmt.Sprintf("%v	%v, $%v\n", cmp, r.name, imm32)
	return asm

}

// asmCmpOp compares x to y, storing the op comparison flag (EQ, NEQ, ...) in result
func asmCmpOp(indent string, op token.Token, signed bool, x *register, y *register, result *register, size uint) string {
	if x.width != y.width || x.width != result.width {
		panic("Invalid register width in asmCmpOp")
	}
	asm := ""
	asm += indent + asmCmpRegReg(indent, x, y, size)
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmComparisonOp", op))
	case token.EQL:
		asm += indent + fmt.Sprintf("SETEQ   %v\n", result.name)
	case token.NEQ:
		asm += indent + fmt.Sprintf("SETNEQ  %v\n", result.name)
	case token.LEQ:
		if signed {
			asm += indent + fmt.Sprintf("SETLE   %v\n", result.name)
		} else {
			asm += indent + fmt.Sprintf("SETLS   %v\n", result.name)
		}
	case token.GEQ:
		if signed {
			asm += indent + fmt.Sprintf("SETGE   %v\n", result.name)
		} else {
			asm += indent + fmt.Sprintf("SETCC   %v\n", result.name)
		}
	case token.LSS:
		if signed {
			asm += indent + fmt.Sprintf("SETLT   %v\n", result.name)
		} else {
			asm += indent + fmt.Sprintf("SETCS   %v\n", result.name)
		}
	case token.GTR:
		if signed {
			asm += indent + fmt.Sprintf("SETGT   %v\n", result.name)
		} else {
			asm += indent + fmt.Sprintf("SETHI   %v\n", result.name)
		}
	}
	return strings.Replace(asm, "+-", "-", -1)
}

func asmRet(indent string) string {
	asm := indent + fmt.Sprintf("RET\n")
	return asm
}
