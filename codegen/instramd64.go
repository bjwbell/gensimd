package codegen

import (
	"fmt"
	"go/token"
	"math"
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
	InstrData
	xmmvariant XmmData
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
	XmmInstr   InstructionType
	Instrf32   Instr
	Instrf64   Instr
	Instrf32x4 Instr // operates on four packed f32
	Instrf64x2 Instr // operate on two packed f64
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
	{I_MOV, MOVSS, MOVSD, MOVUPS, MOVUPD},

	{I_ADD, ADDSS, ADDSD, ADDPS, ADDPD},
	{I_SUB, SUBSS, SUBSD, SUBPS, SUBPD},
	{I_MUL, MULSS, MULSD, MULPS, MULPD},
	{I_DIV, DIVSS, DIVSD, DIVPS, DIVPD},

	{I_XOR, NONE, NONE, XORPS, XORPD},

	{I_CMP, UCOMISS, UCOMISD, NONE, NONE},
}

func GetInstruction(tinst InstructionType) Instruction {
	for _, inst := range Insts {
		if inst.TInstr == tinst {
			return inst
		}
	}
	panic("Couldn't get instruction")
}
func GetXmmInstruction(tinst InstructionType) XmmInstruction {
	for _, inst := range XmmInsts {
		if inst.XmmInstr == tinst {
			return inst
		}
	}
	panic(fmt.Sprintf("Couldn't get xmm instruction (%v)", tinst.String()))
}

func (xinstr XmmInstruction) Select(variant XmmData) Instr {
	var instr Instr
	switch variant {
	case XMM_F32:
		instr = xinstr.Instrf32
	case XMM_F64:
		instr = xinstr.Instrf64
	case XMM_4X_F32:
		instr = xinstr.Instrf32x4
	case XMM_2X_F64:
		instr = xinstr.Instrf64x2
	}
	if instr == NONE {
		panic(fmt.Sprintf("unrecognized variant for %v", xinstr.XmmInstr))
	}
	return instr
}

// GetInstr, the size is in bytes
func GetInstr(tinst InstructionType, datatype InstrDataType) Instr {
	if datatype.op == INTEGER_OP {
		return GetInstruction(tinst).GetSized(datatype.size)
	} else {
		return GetXmmInstruction(tinst).Select(datatype.xmmvariant)
	}
}

// asmZeroMemory generates "MOVQ $0, name+offset(REG)" instructions,
// size is in bytes
func ZeroMemory(indent string, name string, offset int, size uint, reg *register) string {

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
	datatype := InstrDataType{INTEGER_OP, InstrData{signed: false, size: chunk}, XMM_INVALID}
	mov := GetInstr(I_MOV, datatype).String()

	for i := uint(0); i < size/chunk; i++ {
		ioffset := int(i*chunk) + offset
		asm += indent
		asm += fmt.Sprintf("%v    $0, %v+%v(%v)\n", mov, name, ioffset, reg.name)
	}

	return strings.Replace(asm, "+-", "-", -1)
}

// asmZeroReg generates "XORQ reg, reg" instructions
func ZeroReg(indent string, reg *register) string {
	var datatype InstrDataType

	if reg.typ == XMM_REG {
		datatype = InstrDataType{XMM_OP, InstrData{}, XMM_2X_F64}

	} else {
		datatype = InstrDataType{INTEGER_OP, InstrData{signed: false, size: reg.width / 8}, XMM_INVALID}
	}

	xor := GetInstr(I_XOR, datatype)
	return indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), reg.name, reg.name)
}

func MovRegReg(indent string, datatype InstrDataType, srcReg *register, dstReg *register) string {
	if srcReg.width != dstReg.width {
		panic(fmt.Sprintf("srcReg (%v) width (%v) != (%v) dstReg (%v) width or invalid size", srcReg.name, srcReg.width, dstReg.width, dstReg.name))
	}
	mov := GetInstr(I_MOV, datatype).String()
	asm := indent + fmt.Sprintf("%v    %v, %v\n", mov, srcReg.name, dstReg.name)
	return asm
}

func MovRegMem(indent string, datatype InstrDataType, srcReg *register, dstName string, dstReg *register, dstOffset int) string {
	mov := GetInstr(I_MOV, datatype)
	asm := indent + fmt.Sprintf("// BEGIN asmMovRegMem, mov (%v)\n", mov)
	asm += indent + fmt.Sprintf("%v    %v, %v+%v(%v)\n", mov.String(), srcReg.name, dstName, dstOffset, dstReg.name)
	asm += indent + fmt.Sprintf("// END asmMovRegMem, mov (%v)\n", mov)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovRegMemIndirect(indent string, srcReg *register, dstName string, dstReg *register, dstOffset int, tmp *register) string {
	if tmp.width != srcReg.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: srcReg.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("%v    %v, (%v)\n", mov.String(), srcReg.name, tmp.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemMem(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width")
	}
	if size%8 != 0 {
		panic("Invalid size")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v+%v(%v)\n", srcName, srcOffset, srcReg.name, dstName, dstOffset, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemReg(indent string, datatype InstrDataType, srcName string, srcOffset int, srcReg *register, dstReg *register) string {

	mov := GetInstr(I_MOV, datatype)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemMemIndirect(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    %v+%v(%v), (%v)\n", srcName, srcOffset, srcReg.name, tmp)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemIndirectReg(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstReg *register, tmp *register) string {
	if dstReg.width != tmp.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: dstReg.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent
	asm += fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg, tmp)

	asm += indent
	asm += fmt.Sprintf("%v    (%v), %v\n", mov.String(), tmp, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemIndirectMem(indent string, datatype InstrDataType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint, tmp1 *register, tmp2 *register) string {
	if tmp1.width/8 < sizePtr() {
		panic("register width smaller than ptr size ")
	}
	if size > tmp2.width/8 && size%(tmp2.width/8) != 0 {
		panic(fmt.Sprintf("Invalid size (%v), reg width/8 (%v)", size, tmp1.width/8))
	}
	addrdatatype := InstrDataType{INTEGER_OP,
		InstrData{signed: false, size: sizePtr()}, XMM_INVALID}

	mov := GetInstr(I_MOV, datatype).String()
	movaddr := GetInstr(I_MOV, addrdatatype).String()

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

	asm += indent
	asm += fmt.Sprintf("%v    %v+%v(%v), %v\n",
		movaddr, srcName, srcOffset, srcReg.name, tmp1.name)

	for i := uint(0); i < count; i++ {

		asm += indent
		asm += fmt.Sprintf("%v    (%v), %v\n",
			mov, tmp1.name, tmp2.name)

		asm += indent
		asm += fmt.Sprintf("%v    %v, %v+%v(%v)\n",
			mov, tmp2.name, dstName, dstOffset, dstReg.name)

		dstOffset += int(chunk)

		if i < count-1 {
			asm += AddImm32Reg(indent, uint32(chunk), tmp1)
		}
	}

	return strings.Replace(asm, "+-", "-", -1)
}

func MovMemIndirectMemIndirect(indent string, datatype InstrOpType, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp1 *register, tmp2 *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: tmp1.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg, tmp1)
	asm += indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), dstName, dstOffset, dstReg, tmp2)
	asm += indent + fmt.Sprintf("%v    (%v), (%v)\n", mov.String(), tmp1.name, tmp2.name)
	return strings.Replace(asm, "+-", "-", -1)
}

// asmMovImmReg moves imm to reg after converting it to int8/16/32 if size=1/2/4.
func MovImmReg(indent string, imm int64, size uint, dstReg *register) string {
	if dstReg.width < 8*size {
		panic("Invalid register width")
	}
	switch size {
	case 1:
		imm = int64(int8(imm))
	case 2:
		imm = int64(int16(imm))
	case 4:
		imm = int64(int32(imm))
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: true, size: size}, XMM_INVALID}
	mov := GetInstr(I_MOV, data).String()

	asm := indent + fmt.Sprintf("%v    $%v, %v\n", mov, imm, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func MovImmFloatReg(indent string, f64 float64, isf32 bool, tmp *register, dst *register) string {
	if dst.typ != XMM_REG {
		panic("Unexpected non xmm register")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: tmp.width / 8}, XMM_INVALID}
	mov := GetInstr(I_MOV, data).String()
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
	asm := indent + fmt.Sprintf("//      $%v = %016x = %v(%v)\n", fbits, fbits, f64, descrip)
	asm += indent + fmt.Sprintf("%v    $%v, %v\n", mov, fbits, tmp.name)

	data = InstrDataType{INTEGER_OP, InstrData{signed: false, size: 8}, XMM_INVALID}
	mov = GetInstr(I_MOV, data).String()
	asm += indent + fmt.Sprintf("%v    %v, %v\n", mov, tmp.name, dst.name)
	return asm
}

func MovImmf32Reg(indent string, f32 float32, tmp *register, dst *register) string {
	return MovImmFloatReg(indent, float64(f32), true, tmp, dst)
}

func MovImmf64Reg(indent string, f64 float64, tmp *register, dst *register) string {
	return MovImmFloatReg(indent, f64, false, tmp, dst)
}

func MovImm8Reg(indent string, imm8 int8, dstReg *register) string {
	return MovImmReg(indent, int64(imm8), 1, dstReg)
}

func MovImm16Reg(indent string, imm16 int16, dstReg *register) string {
	return MovImmReg(indent, int64(imm16), 2, dstReg)

}

func MovImm32Reg(indent string, imm32 int32, dstReg *register) string {
	return MovImmReg(indent, int64(imm32), 4, dstReg)

}

func MovImm64Reg(indent string, imm64 int64, dstReg *register) string {
	return MovImmReg(indent, imm64, 8, dstReg)
}

// CMovCCRegReg conditionally moves the src reg to the dst reg if the carry
// flag is clear (ie the previous compare had its src greater than its sub reg).
func CMovCCRegReg(indent string, src *register, dst *register, size uint) string {
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

func Lea(indent string, srcName string, srcOffset int, srcReg *register, dstReg *register) string {
	if srcReg.width != dstReg.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: srcReg.width / 8}, XMM_INVALID}
	lea := GetInstr(I_LEA, data)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", lea.String(), srcName, srcOffset, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func AddImm32Reg(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("ADDQ    $%v, %v\n", imm32, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func SubImm32Reg(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("SUBQ    $%v, %v\n", imm32, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func AddRegReg(indent string, datatype InstrDataType, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width")
	}
	add := GetInstr(I_ADD, datatype)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", add.String(), srcReg.name, dstReg.name)
	return asm
}

func SubRegReg(indent string, datatype InstrDataType, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width")
	}

	sub := GetInstr(I_SUB, datatype)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", sub.String(), srcReg.name, dstReg.name)
	return asm
}

func MulImm32RegReg(indent string, imm32 uint32, srcReg *register, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width")
	}
	asm := indent + fmt.Sprintf("IMUL3Q    $%v, %v, %v\n", imm32, srcReg.name, dstReg.name)
	return asm
}

// MulRegReg multiplies the src register by the dst register and stores
// the result in the dst register. Overflow is discarded
func MulRegReg(indent string, datatype InstrDataType, src *register, dst *register) string {

	if dst.width != src.width {
		panic("Invalid register width")
	}
	asm := ""
	if datatype.op == INTEGER_OP {
		rax := getRegister(REG_AX)
		rdx := getRegister(REG_DX)
		if rax.width != 64 || rdx.width != 64 {
			panic("Invalid rax or rdx register width")
		}
		// rax is the implicit destination for MULQ
		asm = MovRegReg(indent, datatype, dst, rax)
		// the low order part of the result is stored in rax and the high order part
		// is stored in rdx
		var tinstr InstructionType
		if datatype.signed {
			tinstr = I_IMUL
		} else {
			tinstr = I_MUL
		}

		mul := GetInstr(tinstr, datatype).String()
		asm += indent + fmt.Sprintf("%v    %v\n", mul, src.name)
		asm += MovRegReg(indent, datatype, rax, dst)
		return strings.Replace(asm, "+-", "-", -1)
	}

	if datatype.op == XMM_OP {
		mul := GetInstr(I_MUL, datatype).String()
		asm += indent + fmt.Sprintf("%v    %v, %v\n", mul, src.name, dst.name)
	}
	return asm

}

// DivRegReg divides the "dividend" register by the "divisor" register and stores
// the quotient in rax and the remainder in rdx. DivRegReg is only for integer division.
func DivRegReg(indent string, signed bool, datatype InstrOpType, dividend *register, divisor *register, size uint) (asm string, rax *register, rdx *register) {

	if dividend.width != divisor.width || divisor.width < size*8 {
		panic("Invalid register width for DivRegReg")
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
	asm += ZeroReg(indent, rax)
	if size > 1 {
		asm += ZeroReg(indent, rdx)
	}
	instrdata := InstrDataType{datatype, InstrData{signed: signed, size: size}, XMM_INVALID}
	asm += MovRegReg(indent, instrdata, dividend, rax)

	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr InstructionType
	if signed {
		tinstr = I_IDIV
	} else {
		tinstr = I_DIV
	}

	div := GetInstr(tinstr, instrdata).String()
	asm += indent + fmt.Sprintf("%v    %v\n", div, divisor.name)
	return asm, rax, rdx
}

// DivFloatRegReg performs floating point division by dividing the dividend register
// by the divisor register and stores the quotient in the dividend register
func DivFloatRegReg(indent string, datatype InstrDataType, dividend *register, divisor *register) string {

	if dividend.width != divisor.width {
		panic("Invalid register width")
	}
	if datatype.op != XMM_OP {
		panic("Unsupported data type for floating point division")
	}

	div := GetInstr(I_DIV, datatype).String()
	asm := indent + fmt.Sprintf("%v    %v, %v\n", div, divisor.name, dividend.name)
	return asm
}

func ArithOp(indent string, datatype InstrDataType, op token.Token, x *register, y *register, result *register) string {
	if x.width != y.width || x.width != result.width {
		panic("Invalid register width")
	}
	asm := ""
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.ADD:
		asm += MovRegReg(indent, datatype, x, result)
		asm += AddRegReg(indent, datatype, y, result)
	case token.SUB:
		asm += MovRegReg(indent, datatype, x, result)
		asm += SubRegReg(indent, datatype, y, result)
	case token.MUL:
		asm += MovRegReg(indent, datatype, x, result)
		asm += MulRegReg(indent, datatype, y, result)
	case token.QUO, token.REM:
		if datatype.op == INTEGER_OP {
			// the quotient is stored in rax and
			// the remainder is stored in rdx.
			var rax, rdx *register
			a, rax, rdx := DivRegReg(indent, datatype.signed, datatype.op, x, y, datatype.size)
			asm += a
			if op == token.QUO {
				asm += MovRegReg(indent, datatype, rax, result)
			} else {
				asm += MovRegReg(indent, datatype, rdx, result)
			}
		} else {
			// assume quotient operation,
			// since modulus isn't defined for floats
			asm += MovRegReg(indent, datatype, x, result)
			asm += DivFloatRegReg(indent, datatype, result, y)
		}
	}
	return strings.Replace(asm, "+-", "-", -1)
}

// AndRegReg AND's the src register by the dst register and stores
// the result in the dst register.
func AndRegReg(indent string, src *register, dst *register, size uint) string {
	if src.width != dst.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	and := GetInstr(I_AND, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", and.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func OrRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: src.width / 8}, XMM_INVALID}
	or := GetInstr(I_OR, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", or.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func XorRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: src.width / 8}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), src.name, dst.name)
	return asm
}

func XorImm32Reg(indent string, imm32 int32, dst *register, size uint) string {
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    $%v, %v\n", xor.String(), imm32, dst.name)
	return asm
}

func XorImm64Reg(indent string, imm64 int64, dst *register, size uint) string {
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	xor := GetInstr(I_XOR, data)
	asm := indent + fmt.Sprintf("%v    $%v, %v\n", xor.String(), imm64, dst.name)
	return asm
}

func NotReg(indent string, reg *register, size uint) string {
	return XorImm32Reg(indent, -1, reg, size)
}

const (
	SHIFT_LEFT = iota
	SHIFT_RIGHT
)

func MovZeroExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
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
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: dstSize}, XMM_INVALID}
	movzx := GetInstr(opcode, data)
	asm := indent + fmt.Sprintf("%v %v, %v\n", movzx.String(), src.name, dst.name)
	return asm
}

func MovSignExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
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
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: dstSize}, XMM_INVALID}
	movsx := GetInstr(opcode, data)
	if movsx == NONE {
		panic(fmt.Sprintf("Bad dstSize (%v) for sign extend result", dstSize))
	}
	asm := indent + fmt.Sprintf("%v    %v, %v\n", movsx.String(), src.name, dst.name)
	return asm
}

func ShiftRegReg(indent string, signed bool, direction int, src *register, shiftReg *register, tmp *register, size uint) string {

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

	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
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
	instrdata := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	asm += MovRegReg(indent, instrdata, shiftReg, cx)

	asm += MovImm32Reg(indent, completeShift, tmp)
	// compare only first byte of shift reg,
	// since useful shift can be at most 64
	asm += CmpRegImm32(indent, shiftReg, maxShift, 1)
	asm += CMovCCRegReg(indent, tmp, cx, size)

	var zerosize uint = 1
	asm += MovZeroExtend(indent, cl, cx, zerosize, cx.width/8)
	asm += indent + fmt.Sprintf("%v    %v, %v\n", shift.String(), regCl.name, src.name)

	if maxShift == 64 || maxShift == 32 {
		asm += MovImm32Reg(indent, 1, tmp)
		// compare only first byte of shift reg,
		// since useful shift can be at most 64
		asm += XorRegReg(indent, cx, cx)
		asm += CmpRegImm32(indent, shiftReg, maxShift, 1)
		asm += CMovCCRegReg(indent, tmp, cx, size)
		var zerosize uint = 1
		asm += MovZeroExtend(indent, cl, cx, zerosize, cx.width/8)
		asm += indent + fmt.Sprintf("%v    %v, %v\n", shift.String(), regCl.name, src.name)
	}

	return asm
}

func AndNotRegReg(indent string, src *register, dst *register, size uint) string {
	if src.width != dst.width {
		panic("Invalid register width")
	}
	asm := NotReg(indent, dst, size)
	asm += AndRegReg(indent, src, dst, size)
	return asm
}

func BitwiseOp(indent string, op token.Token, signed bool, x *register, y *register, result *register, size uint) string {
	if x.width != y.width || x.width != result.width {
		panic("Invalid register width")
	}
	asm := ""
	instrdata := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.AND:

		asm = MovRegReg(indent, instrdata, y, result)
		asm += AndRegReg(indent, x, result, size)
	case token.OR:
		asm = MovRegReg(indent, instrdata, y, result)
		asm += OrRegReg(indent, x, result)
	case token.XOR:
		asm = MovRegReg(indent, instrdata, y, result)
		asm += XorRegReg(indent, x, result)
	case token.SHL:
		asm = MovRegReg(indent, instrdata, x, result)
		tmp := x
		asm += ShiftRegReg(indent, signed, SHIFT_LEFT, result, y, tmp, size)
	case token.SHR:
		asm = MovRegReg(indent, instrdata, x, result)
		tmp := x
		asm += ShiftRegReg(indent, signed, SHIFT_RIGHT, result, y, tmp, size)
	case token.AND_NOT:
		asm = MovRegReg(indent, instrdata, y, result)
		asm += AndNotRegReg(indent, x, result, size)

	}
	return strings.Replace(asm, "+-", "-", -1)
}

func CmpRegReg(indent string, instrdata InstrDataType, x *register, y *register) string {
	if x.width != y.width {
		panic("Invalid register width")
	}
	cmp := GetInstr(I_CMP, instrdata).String()
	asm := fmt.Sprintf("%v	%v, %v\n", cmp, x.name, y.name)
	return asm

}

func CmpMemImm32(indent string, name string, offset int32, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data)
	asm := indent + fmt.Sprintf("%v	%v+%v(%v), $%v\n", cmp, name, offset, r.name, imm32)
	return strings.Replace(asm, "+-", "-", -1)

}

func CmpRegImm32(indent string, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		panic("Invalid register width")
	}
	data := InstrDataType{INTEGER_OP, InstrData{signed: false, size: size}, XMM_INVALID}
	cmp := GetInstr(I_CMP, data).String()
	asm := indent + fmt.Sprintf("%v	%v, $%v\n", cmp, r.name, imm32)
	return asm

}

// CmpOp compares x to y, storing the op comparison flag (EQ, NEQ, ...) in result
func CmpOp(indent string, data InstrDataType, op token.Token, x *register, y *register, result *register) string {
	if x.width != y.width {
		panic(fmt.Sprintf("Invalid register width, x.width (%v), y.width (%v), result.width (%v)", x.width, y.width, result.width))
	}
	asm := ""
	asm += indent + CmpRegReg(indent, data, x, y)
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v)", op))
	case token.EQL:
		asm += indent + fmt.Sprintf("SETEQ   %v\n", result.name)
	case token.NEQ:
		asm += indent + fmt.Sprintf("SETNEQ  %v\n", result.name)
	case token.LEQ:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == XMM_OP {
			asm += indent + fmt.Sprintf("SETCC   %v\n", result.name)
		} else {
			if data.signed {
				asm += indent + fmt.Sprintf("SETLE   %v\n", result.name)
			} else {
				asm += indent + fmt.Sprintf("SETLS   %v\n", result.name)
			}
		}
	case token.GEQ:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == XMM_OP {
			asm += indent + fmt.Sprintf("SETLS   %v\n", result.name)
		} else {
			if data.signed {
				asm += indent + fmt.Sprintf("SETGE   %v\n", result.name)
			} else {
				asm += indent + fmt.Sprintf("SETCC   %v\n", result.name)
			}
		}
	case token.LSS:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == XMM_OP {
			asm += indent + fmt.Sprintf("SETHI   %v\n", result.name)
		} else {
			if data.signed {
				asm += indent + fmt.Sprintf("SETLT   %v\n", result.name)
			} else {
				asm += indent + fmt.Sprintf("SETCS   %v\n", result.name)
			}
		}
	case token.GTR:
		// for some reason the SETXX are flipped for xmm compares
		if data.op == XMM_OP {
			asm += indent + fmt.Sprintf("SETCS   %v\n", result.name)
		} else {
			if data.signed {
				asm += indent + fmt.Sprintf("SETGT   %v\n", result.name)
			} else {
				asm += indent + fmt.Sprintf("SETHI   %v\n", result.name)
			}
		}
	}
	return strings.Replace(asm, "+-", "-", -1)
}

func Ret(indent string) string {
	asm := indent + fmt.Sprintf("RET\n")
	return asm
}
