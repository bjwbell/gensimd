package codegen

import (
	"fmt"
	"go/token"
	"log"
	"strings"
)

type Operand struct {
	Type   OperandType
	Input  bool
	Output bool

	Value func() string
}

type OperandType int

// the list of operand types is from Marat Dukhan's
// https://github.com/Maratyszcza/Opcodes
const (
	// al register
	AL OperandType = iota
	// ax register
	AX
	// eax register
	EAX
	// rax register
	RAX
	// cl register
	CL
	// immediate
	IMM8
	IMM16
	IMM32
	IMM64
	// register instructions
	R8
	R16
	R32
	R64
	// memory instructions
	M // any size
	M8
	M16
	M32
	M64
	M128
	// xmm instructions
	// xmm0 register
	XMM0
	// xmm register (xmm0-xmm31)
	XMM
	// used for jump instructions
	REL8
	REL32
)

func (op OperandType) String() string {
	switch op {
	default:
		log.Fatalf("Unknown OperandType: \"%v\"", int(op))
		return ""
	case AL:
		return "AL"
	case CL:
		return "CL"
	case AX:
		return "AL"
	case EAX:
		return "EAX"
	case RAX:
		return "RAX"

		// immediate
	case IMM8:
		return "IMM8"
	case IMM16:
		return "IMM16"
	case IMM32:
		return "IMM32"
	case IMM64:
		return "IMM64"

		// register instructions
	case R8:
		return "R8"
	case R16:
		return "R16"
	case R32:
		return "R32"
	case R64:
		return "R64"

		// memory instructions
	case M:
		return "M"
	case M8:
		return "M8"
	case M16:
		return "M16"
	case M32:
		return "M32"
	case M64:
		return "M64"
	case M128:
		return "M128"

		// xmm instructions
	case XMM:
		return "XMM"
	case XMM0:
		return "XMM0"

		// jump instructions
	case REL8:
		return "REL8"
	case REL32:
		return "REL32"
	}
}

type Instr int
type Instruction struct {
	TInstr     TInstruction
	ByteSized  Instr
	WordSized  Instr
	LongSized  Instr
	QuadSized  Instr
	DQuadSized Instr
}

type InstrDataSize int

const (
	NoneSize InstrDataSize = iota
	BSize
	WSize
	LSize
	QSize
	DQSize
)

func GetInstrDataSize(size uint) InstrDataSize {
	switch size {
	case 1:
		return BSize
	case 2:
		return WSize
	case 4:
		return LSize
	case 8:
		return QSize
	case 16:
		return DQSize
	}
	panic(fmt.Sprintf("Invalid size (%v) in GetInstrDataSize", size))
}

var instrDataSizes = []InstrDataSize{NoneSize, BSize, WSize, NoneSize, LSize, NoneSize, NoneSize, NoneSize, QSize}

func (inst Instruction) GetSized(size InstrDataSize) Instr {
	instrs := []Instr{NONE, inst.ByteSized, inst.WordSized, inst.LongSized, inst.QuadSized}
	instr := instrs[size]
	if instr != NONE {
		return instr
	}
	msg := fmt.Sprintf("Invalid size(%v), for instr (%v)", size, inst.TInstr.String())
	panic(msg)
}

func (inst Instruction) Get(size uint) Instr {
	instr := inst.GetSized(GetInstrDataSize(size))
	if instr == NONE {
		msg := fmt.Sprintf("No matching instruction version for size (%v), instruction (%v)", size, inst.TInstr.String())
		panic(msg)
	} else {
		return instr
	}
}

type TInstruction int

const (
	I_ADD TInstruction = iota
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

func (tinst TInstruction) String() string {
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
	{I_IMUL, IMULB, IMULW, IMULL, IMULQ, NONE},
	{I_DIV, DIVB, DIVW, DIVL, DIVQ, NONE},
	{I_IDIV, IDIVB, IDIVW, IDIVL, IDIVQ, NONE},
	{I_AND, ANDB, ANDW, ANDL, ANDQ, NONE},
	{I_OR, ORB, ORW, ORL, ORQ, NONE},
	{I_SHL, SHLB, SHLW, SHLL, SHLQ, NONE},
	{I_SHR, SHRB, SHRW, SHRL, SHRQ, NONE},
	{I_SAL, SALB, SHLW, SHLL, SHLQ, NONE},
	{I_SAR, SARB, SARW, SARL, SARQ, NONE},

	{I_CMP, CMPB, CMPW, CMPL, CMPQ, NONE},
}

func GetInstruction(tinst TInstruction) Instruction {
	for _, inst := range Insts {
		if inst.TInstr == tinst {
			return inst
		}
	}
	panic("Couldn't get instruction")
}

// GetInstr, the size is in bytes
func GetInstr(tinst TInstruction, size uint) Instr {
	inst := GetInstruction(tinst)
	return inst.Get(size)
}

// the list of instruction names is from
// https://github.com/golang/go/blob/master/src/cmd/internal/obj/x86/a.out.go
const (
	NONE Instr = iota
	AAD
	AAM
	AAS
	ADCB
	ADCL
	ADCW
	ADDB
	ADDL
	ADDW
	ADJSP
	ANDB
	ANDL
	ANDW
	ARPL
	BOUNDL
	BOUNDW
	BSFL
	BSFW
	BSRL
	BSRW
	BTL
	BTW
	BTCL
	BTCW
	BTRL
	BTRW
	BTSL
	BTSW
	BYTE
	CLC
	CLD
	CLI
	CLTS
	CMC
	CMPB
	CMPL
	CMPW
	CMPSB
	CMPSL
	CMPSW
	DAA
	DAS
	DECB
	DECL
	DECQ
	DECW
	DIVB
	DIVL
	DIVW
	ENTER
	HLT
	IDIVB
	IDIVL
	IDIVW
	IMULB
	IMULL
	IMULW
	INB
	INL
	INW
	INCB
	INCL
	INCQ
	INCW
	INSB
	INSL
	INSW
	INT
	INTO
	IRETL
	IRETW
	JCC
	JCS
	JCXZL
	JEQ
	JGE
	JGT
	JHI
	JLE
	JLS
	JLT
	JMI
	JNE
	JOC
	JOS
	JPC
	JPL
	JPS
	LAHF
	LARL
	LARW
	LEAL
	LEAW
	LEAVEL
	LEAVEW
	LOCK
	LODSB
	LODSL
	LODSW
	LONG
	LOOP
	LOOPEQ
	LOOPNE
	LSLL
	LSLW
	MOVB
	MOVL
	MOVW
	MOVBLSX
	MOVBLZX
	MOVBQSX
	MOVBQZX
	MOVBWSX
	MOVBWZX
	MOVWLSX
	MOVWLZX
	MOVWQSX
	MOVWQZX
	MOVSB
	MOVSL
	MOVSW
	MULB
	MULL
	MULW
	NEGB
	NEGL
	NEGW
	NOTB
	NOTL
	NOTW
	ORB
	ORL
	ORW
	OUTB
	OUTL
	OUTW
	OUTSB
	OUTSL
	OUTSW
	PAUSE
	POPAL
	POPAW
	POPFL
	POPFW
	POPL
	POPW
	PUSHAL
	PUSHAW
	PUSHFL
	PUSHFW
	PUSHL
	PUSHW
	RCLB
	RCLL
	RCLW
	RCRB
	RCRL
	RCRW
	REP
	REPN
	ROLB
	ROLL
	ROLW
	RORB
	RORL
	RORW
	SAHF
	SALB
	SALL
	SALW
	SARB
	SARL
	SARW
	SBBB
	SBBL
	SBBW
	SCASB
	SCASL
	SCASW
	SETCC
	SETCS
	SETEQ
	SETGE
	SETGT
	SETHI
	SETLE
	SETLS
	SETLT
	SETMI
	SETNE
	SETOC
	SETOS
	SETPC
	SETPL
	SETPS
	CDQ
	CWD
	SHLB
	SHLL
	SHLW
	SHRB
	SHRL
	SHRW
	STC
	STD
	STI
	STOSB
	STOSL
	STOSW
	SUBB
	SUBL
	SUBW
	SYSCALL
	TESTB
	TESTL
	TESTW
	VERR
	VERW
	WAIT
	WORD
	XCHGB
	XCHGL
	XCHGW
	XLAT
	XORB
	XORL
	XORW
	FMOVB
	FMOVBP
	FMOVD
	FMOVDP
	FMOVF
	FMOVFP
	FMOVL
	FMOVLP
	FMOVV
	FMOVVP
	FMOVW
	FMOVWP
	FMOVX
	FMOVXP
	FCOMB
	FCOMBP
	FCOMD
	FCOMDP
	FCOMDPP
	FCOMF
	FCOMFP
	FCOML
	FCOMLP
	FCOMW
	FCOMWP
	FUCOM
	FUCOMP
	FUCOMPP
	FADDDP
	FADDW
	FADDL
	FADDF
	FADDD
	FMULDP
	FMULW
	FMULL
	FMULF
	FMULD
	FSUBDP
	FSUBW
	FSUBL
	FSUBF
	FSUBD
	FSUBRDP
	FSUBRW
	FSUBRL
	FSUBRF
	FSUBRD
	FDIVDP
	FDIVW
	FDIVL
	FDIVF
	FDIVD
	FDIVRDP
	FDIVRW
	FDIVRL
	FDIVRF
	FDIVRD
	FXCHD
	FFREE
	FLDCW
	FLDENV
	FRSTOR
	FSAVE
	FSTCW
	FSTENV
	FSTSW
	F2XM1
	FABS
	FCHS
	FCLEX
	FCOS
	FDECSTP
	FINCSTP
	FINIT
	FLD1
	FLDL2E
	FLDL2T
	FLDLG2
	FLDLN2
	FLDPI
	FLDZ
	FNOP
	FPATAN
	FPREM
	FPREM1
	FPTAN
	FRNDINT
	FSCALE
	FSIN
	FSINCOS
	FSQRT
	FTST
	FXAM
	FXTRACT
	FYL2X
	FYL2XP1
	CMPXCHGB
	CMPXCHGL
	CMPXCHGW
	CMPXCHG8B
	CPUID
	INVD
	INVLPG
	LFENCE
	MFENCE
	MOVNTIL
	RDMSR
	RDPMC
	RDTSC
	RSM
	SFENCE
	SYSRET
	WBINVD
	WRMSR
	XADDB
	XADDL
	XADDW
	CMOVLCC
	CMOVLCS
	CMOVLEQ
	CMOVLGE
	CMOVLGT
	CMOVLHI
	CMOVLLE
	CMOVLLS
	CMOVLLT
	CMOVLMI
	CMOVLNE
	CMOVLOC
	CMOVLOS
	CMOVLPC
	CMOVLPL
	CMOVLPS
	CMOVQCC
	CMOVQCS
	CMOVQEQ
	CMOVQGE
	CMOVQGT
	CMOVQHI
	CMOVQLE
	CMOVQLS
	CMOVQLT
	CMOVQMI
	CMOVQNE
	CMOVQOC
	CMOVQOS
	CMOVQPC
	CMOVQPL
	CMOVQPS
	CMOVWCC
	CMOVWCS
	CMOVWEQ
	CMOVWGE
	CMOVWGT
	CMOVWHI
	CMOVWLE
	CMOVWLS
	CMOVWLT
	CMOVWMI
	CMOVWNE
	CMOVWOC
	CMOVWOS
	CMOVWPC
	CMOVWPL
	CMOVWPS
	ADCQ
	ADDQ
	ANDQ
	BSFQ
	BSRQ
	BTCQ
	BTQ
	BTRQ
	BTSQ
	CMPQ
	CMPSQ
	CMPXCHGQ
	CQO
	DIVQ
	IDIVQ
	IMULQ
	IRETQ
	JCXZQ
	LEAQ
	LEAVEQ
	LODSQ
	MOVQ
	MOVLQSX
	MOVLQZX
	MOVNTIQ
	MOVSQ
	MULQ
	NEGQ
	NOTQ
	ORQ
	POPFQ
	POPQ
	PUSHFQ
	PUSHQ
	RCLQ
	RCRQ
	ROLQ
	RORQ
	QUAD
	SALQ
	SARQ
	SBBQ
	SCASQ
	SHLQ
	SHRQ
	STOSQ
	SUBQ
	TESTQ
	XADDQ
	XCHGQ
	XORQ
	ADDPD
	ADDPS
	ADDSD
	ADDSS
	ANDNPD
	ANDNPS
	ANDPD
	ANDPS
	CMPPD
	CMPPS
	CMPSD
	CMPSS
	COMISD
	COMISS
	CVTPD2PL
	CVTPD2PS
	CVTPL2PD
	CVTPL2PS
	CVTPS2PD
	CVTPS2PL
	CVTSD2SL
	CVTSD2SQ
	CVTSD2SS
	CVTSL2SD
	CVTSL2SS
	CVTSQ2SD
	CVTSQ2SS
	CVTSS2SD
	CVTSS2SL
	CVTSS2SQ
	CVTTPD2PL
	CVTTPS2PL
	CVTTSD2SL
	CVTTSD2SQ
	CVTTSS2SL
	CVTTSS2SQ
	DIVPD
	DIVPS
	DIVSD
	DIVSS
	EMMS
	FXRSTOR
	FXRSTOR64
	FXSAVE
	FXSAVE64
	LDMXCSR
	MASKMOVOU
	MASKMOVQ
	MAXPD
	MAXPS
	MAXSD
	MAXSS
	MINPD
	MINPS
	MINSD
	MINSS
	MOVAPD
	MOVAPS
	MOVOU
	MOVHLPS
	MOVHPD
	MOVHPS
	MOVLHPS
	MOVLPD
	MOVLPS
	MOVMSKPD
	MOVMSKPS
	MOVNTO
	MOVNTPD
	MOVNTPS
	MOVNTQ
	MOVO
	MOVQOZX
	MOVSD
	MOVSS
	MOVUPD
	MOVUPS
	MULPD
	MULPS
	MULSD
	MULSS
	ORPD
	ORPS
	PACKSSLW
	PACKSSWB
	PACKUSWB
	PADDB
	PADDL
	PADDQ
	PADDSB
	PADDSW
	PADDUSB
	PADDUSW
	PADDW
	PANDB
	PANDL
	PANDSB
	PANDSW
	PANDUSB
	PANDUSW
	PANDW
	PAND
	PANDN
	PAVGB
	PAVGW
	PCMPEQB
	PCMPEQL
	PCMPEQW
	PCMPGTB
	PCMPGTL
	PCMPGTW
	PEXTRW
	PFACC
	PFADD
	PFCMPEQ
	PFCMPGE
	PFCMPGT
	PFMAX
	PFMIN
	PFMUL
	PFNACC
	PFPNACC
	PFRCP
	PFRCPIT1
	PFRCPI2T
	PFRSQIT1
	PFRSQRT
	PFSUB
	PFSUBR
	PINSRW
	PINSRD
	PINSRQ
	PMADDWL
	PMAXSW
	PMAXUB
	PMINSW
	PMINUB
	PMOVMSKB
	PMULHRW
	PMULHUW
	PMULHW
	PMULLW
	PMULULQ
	POR
	PSADBW
	PSHUFHW
	PSHUFL
	PSHUFLW
	PSHUFW
	PSHUFB
	PSLLO
	PSLLL
	PSLLQ
	PSLLW
	PSRAL
	PSRAW
	PSRLO
	PSRLL
	PSRLQ
	PSRLW
	PSUBB
	PSUBL
	PSUBQ
	PSUBSB
	PSUBSW
	PSUBUSB
	PSUBUSW
	PSUBW
	PSWAPL
	PUNPCKHBW
	PUNPCKHLQ
	PUNPCKHQDQ
	PUNPCKHWL
	PUNPCKLBW
	PUNPCKLLQ
	PUNPCKLQDQ
	PUNPCKLWL
	PXOR
	RCPPS
	RCPSS
	RSQRTPS
	RSQRTSS
	SHUFPD
	SHUFPS
	SQRTPD
	SQRTPS
	SQRTSD
	SQRTSS
	STMXCSR
	SUBPD
	SUBPS
	SUBSD
	SUBSS
	UCOMISD
	UCOMISS
	UNPCKHPD
	UNPCKHPS
	UNPCKLPD
	UNPCKLPS
	XORPD
	XORPS
	PF2IW
	PF2IL
	PI2FW
	PI2FL
	RETFW
	RETFL
	RETFQ
	SWAPGS
	MODE
	CRC32B
	CRC32Q
	IMUL3Q
	PREFETCHT0
	PREFETCHT1
	PREFETCHT2
	PREFETCHNTA
	MOVQL
	BSWAPL
	BSWAPQ
	AESENC
	AESENCLAST
	AESDEC
	AESDECLAST
	AESIMC
	AESKEYGENASSIST
	ROUNDPS
	ROUNDSS
	ROUNDPD
	ROUNDSD
	PSHUFD
	PCLMULQDQ
	JCXZW
	FCMOVCC
	FCMOVCS
	FCMOVEQ
	FCMOVHI
	FCMOVLS
	FCMOVNE
	FCMOVNU
	FCMOVUN
	FCOMI
	FCOMIP
	FUCOMI
	FUCOMIP
	LAST
)

var InstrString = []string{
	"AAA",
	"AAD",
	"AAM",
	"AAS",
	"ADCB",
	"ADCL",
	"ADCW",
	"ADDB",
	"ADDL",
	"ADDW",
	"ADJSP",
	"ANDB",
	"ANDL",
	"ANDW",
	"ARPL",
	"BOUNDL",
	"BOUNDW",
	"BSFL",
	"BSFW",
	"BSRL",
	"BSRW",
	"BTL",
	"BTW",
	"BTCL",
	"BTCW",
	"BTRL",
	"BTRW",
	"BTSL",
	"BTSW",
	"BYTE",
	"CLC",
	"CLD",
	"CLI",
	"CLTS",
	"CMC",
	"CMPB",
	"CMPL",
	"CMPW",
	"CMPSB",
	"CMPSL",
	"CMPSW",
	"DAA",
	"DAS",
	"DECB",
	"DECL",
	"DECQ",
	"DECW",
	"DIVB",
	"DIVL",
	"DIVW",
	"ENTER",
	"HLT",
	"IDIVB",
	"IDIVL",
	"IDIVW",
	"IMULB",
	"IMULL",
	"IMULW",
	"INB",
	"INL",
	"INW",
	"INCB",
	"INCL",
	"INCQ",
	"INCW",
	"INSB",
	"INSL",
	"INSW",
	"INT",
	"INTO",
	"IRETL",
	"IRETW",
	"JCC",
	"JCS",
	"JCXZL",
	"JEQ",
	"JGE",
	"JGT",
	"JHI",
	"JLE",
	"JLS",
	"JLT",
	"JMI",
	"JNE",
	"JOC",
	"JOS",
	"JPC",
	"JPL",
	"JPS",
	"LAHF",
	"LARL",
	"LARW",
	"LEAL",
	"LEAW",
	"LEAVEL",
	"LEAVEW",
	"LOCK",
	"LODSB",
	"LODSL",
	"LODSW",
	"LONG",
	"LOOP",
	"LOOPEQ",
	"LOOPNE",
	"LSLL",
	"LSLW",
	"MOVB",
	"MOVL",
	"MOVW",
	"MOVBLSX",
	"MOVBLZX",
	"MOVBQSX",
	"MOVBQZX",
	"MOVBWSX",
	"MOVBWZX",
	"MOVWLSX",
	"MOVWLZX",
	"MOVWQSX",
	"MOVWQZX",
	"MOVSB",
	"MOVSL",
	"MOVSW",
	"MULB",
	"MULL",
	"MULW",
	"NEGB",
	"NEGL",
	"NEGW",
	"NOTB",
	"NOTL",
	"NOTW",
	"ORB",
	"ORL",
	"ORW",
	"OUTB",
	"OUTL",
	"OUTW",
	"OUTSB",
	"OUTSL",
	"OUTSW",
	"PAUSE",
	"POPAL",
	"POPAW",
	"POPFL",
	"POPFW",
	"POPL",
	"POPW",
	"PUSHAL",
	"PUSHAW",
	"PUSHFL",
	"PUSHFW",
	"PUSHL",
	"PUSHW",
	"RCLB",
	"RCLL",
	"RCLW",
	"RCRB",
	"RCRL",
	"RCRW",
	"REP",
	"REPN",
	"ROLB",
	"ROLL",
	"ROLW",
	"RORB",
	"RORL",
	"RORW",
	"SAHF",
	"SALB",
	"SALL",
	"SALW",
	"SARB",
	"SARL",
	"SARW",
	"SBBB",
	"SBBL",
	"SBBW",
	"SCASB",
	"SCASL",
	"SCASW",
	"SETCC",
	"SETCS",
	"SETEQ",
	"SETGE",
	"SETGT",
	"SETHI",
	"SETLE",
	"SETLS",
	"SETLT",
	"SETMI",
	"SETNE",
	"SETOC",
	"SETOS",
	"SETPC",
	"SETPL",
	"SETPS",
	"CDQ",
	"CWD",
	"SHLB",
	"SHLL",
	"SHLW",
	"SHRB",
	"SHRL",
	"SHRW",
	"STC",
	"STD",
	"STI",
	"STOSB",
	"STOSL",
	"STOSW",
	"SUBB",
	"SUBL",
	"SUBW",
	"SYSCALL",
	"TESTB",
	"TESTL",
	"TESTW",
	"VERR",
	"VERW",
	"WAIT",
	"WORD",
	"XCHGB",
	"XCHGL",
	"XCHGW",
	"XLAT",
	"XORB",
	"XORL",
	"XORW",
	"FMOVB",
	"FMOVBP",
	"FMOVD",
	"FMOVDP",
	"FMOVF",
	"FMOVFP",
	"FMOVL",
	"FMOVLP",
	"FMOVV",
	"FMOVVP",
	"FMOVW",
	"FMOVWP",
	"FMOVX",
	"FMOVXP",
	"FCOMB",
	"FCOMBP",
	"FCOMD",
	"FCOMDP",
	"FCOMDPP",
	"FCOMF",
	"FCOMFP",
	"FCOML",
	"FCOMLP",
	"FCOMW",
	"FCOMWP",
	"FUCOM",
	"FUCOMP",
	"FUCOMPP",
	"FADDDP",
	"FADDW",
	"FADDL",
	"FADDF",
	"FADDD",
	"FMULDP",
	"FMULW",
	"FMULL",
	"FMULF",
	"FMULD",
	"FSUBDP",
	"FSUBW",
	"FSUBL",
	"FSUBF",
	"FSUBD",
	"FSUBRDP",
	"FSUBRW",
	"FSUBRL",
	"FSUBRF",
	"FSUBRD",
	"FDIVDP",
	"FDIVW",
	"FDIVL",
	"FDIVF",
	"FDIVD",
	"FDIVRDP",
	"FDIVRW",
	"FDIVRL",
	"FDIVRF",
	"FDIVRD",
	"FXCHD",
	"FFREE",
	"FLDCW",
	"FLDENV",
	"FRSTOR",
	"FSAVE",
	"FSTCW",
	"FSTENV",
	"FSTSW",
	"F2XM1",
	"FABS",
	"FCHS",
	"FCLEX",
	"FCOS",
	"FDECSTP",
	"FINCSTP",
	"FINIT",
	"FLD1",
	"FLDL2E",
	"FLDL2T",
	"FLDLG2",
	"FLDLN2",
	"FLDPI",
	"FLDZ",
	"FNOP",
	"FPATAN",
	"FPREM",
	"FPREM1",
	"FPTAN",
	"FRNDINT",
	"FSCALE",
	"FSIN",
	"FSINCOS",
	"FSQRT",
	"FTST",
	"FXAM",
	"FXTRACT",
	"FYL2X",
	"FYL2XP1",
	"CMPXCHGB",
	"CMPXCHGL",
	"CMPXCHGW",
	"CMPXCHG8B",
	"CPUID",
	"INVD",
	"INVLPG",
	"LFENCE",
	"MFENCE",
	"MOVNTIL",
	"RDMSR",
	"RDPMC",
	"RDTSC",
	"RSM",
	"SFENCE",
	"SYSRET",
	"WBINVD",
	"WRMSR",
	"XADDB",
	"XADDL",
	"XADDW",
	"CMOVLCC",
	"CMOVLCS",
	"CMOVLEQ",
	"CMOVLGE",
	"CMOVLGT",
	"CMOVLHI",
	"CMOVLLE",
	"CMOVLLS",
	"CMOVLLT",
	"CMOVLMI",
	"CMOVLNE",
	"CMOVLOC",
	"CMOVLOS",
	"CMOVLPC",
	"CMOVLPL",
	"CMOVLPS",
	"CMOVQCC",
	"CMOVQCS",
	"CMOVQEQ",
	"CMOVQGE",
	"CMOVQGT",
	"CMOVQHI",
	"CMOVQLE",
	"CMOVQLS",
	"CMOVQLT",
	"CMOVQMI",
	"CMOVQNE",
	"CMOVQOC",
	"CMOVQOS",
	"CMOVQPC",
	"CMOVQPL",
	"CMOVQPS",
	"CMOVWCC",
	"CMOVWCS",
	"CMOVWEQ",
	"CMOVWGE",
	"CMOVWGT",
	"CMOVWHI",
	"CMOVWLE",
	"CMOVWLS",
	"CMOVWLT",
	"CMOVWMI",
	"CMOVWNE",
	"CMOVWOC",
	"CMOVWOS",
	"CMOVWPC",
	"CMOVWPL",
	"CMOVWPS",
	"ADCQ",
	"ADDQ",
	"ANDQ",
	"BSFQ",
	"BSRQ",
	"BTCQ",
	"BTQ",
	"BTRQ",
	"BTSQ",
	"CMPQ",
	"CMPSQ",
	"CMPXCHGQ",
	"CQO",
	"DIVQ",
	"IDIVQ",
	"IMULQ",
	"IRETQ",
	"JCXZQ",
	"LEAQ",
	"LEAVEQ",
	"LODSQ",
	"MOVQ",
	"MOVLQSX",
	"MOVLQZX",
	"MOVNTIQ",
	"MOVSQ",
	"MULQ",
	"NEGQ",
	"NOTQ",
	"ORQ",
	"POPFQ",
	"POPQ",
	"PUSHFQ",
	"PUSHQ",
	"RCLQ",
	"RCRQ",
	"ROLQ",
	"RORQ",
	"QUAD",
	"SALQ",
	"SARQ",
	"SBBQ",
	"SCASQ",
	"SHLQ",
	"SHRQ",
	"STOSQ",
	"SUBQ",
	"TESTQ",
	"XADDQ",
	"XCHGQ",
	"XORQ",
	"ADDPD",
	"ADDPS",
	"ADDSD",
	"ADDSS",
	"ANDNPD",
	"ANDNPS",
	"ANDPD",
	"ANDPS",
	"CMPPD",
	"CMPPS",
	"CMPSD",
	"CMPSS",
	"COMISD",
	"COMISS",
	"CVTPD2PL",
	"CVTPD2PS",
	"CVTPL2PD",
	"CVTPL2PS",
	"CVTPS2PD",
	"CVTPS2PL",
	"CVTSD2SL",
	"CVTSD2SQ",
	"CVTSD2SS",
	"CVTSL2SD",
	"CVTSL2SS",
	"CVTSQ2SD",
	"CVTSQ2SS",
	"CVTSS2SD",
	"CVTSS2SL",
	"CVTSS2SQ",
	"CVTTPD2PL",
	"CVTTPS2PL",
	"CVTTSD2SL",
	"CVTTSD2SQ",
	"CVTTSS2SL",
	"CVTTSS2SQ",
	"DIVPD",
	"DIVPS",
	"DIVSD",
	"DIVSS",
	"EMMS",
	"FXRSTOR",
	"FXRSTOR64",
	"FXSAVE",
	"FXSAVE64",
	"LDMXCSR",
	"MASKMOVOU",
	"MASKMOVQ",
	"MAXPD",
	"MAXPS",
	"MAXSD",
	"MAXSS",
	"MINPD",
	"MINPS",
	"MINSD",
	"MINSS",
	"MOVAPD",
	"MOVAPS",
	"MOVOU",
	"MOVHLPS",
	"MOVHPD",
	"MOVHPS",
	"MOVLHPS",
	"MOVLPD",
	"MOVLPS",
	"MOVMSKPD",
	"MOVMSKPS",
	"MOVNTO",
	"MOVNTPD",
	"MOVNTPS",
	"MOVNTQ",
	"MOVO",
	"MOVQOZX",
	"MOVSD",
	"MOVSS",
	"MOVUPD",
	"MOVUPS",
	"MULPD",
	"MULPS",
	"MULSD",
	"MULSS",
	"ORPD",
	"ORPS",
	"PACKSSLW",
	"PACKSSWB",
	"PACKUSWB",
	"PADDB",
	"PADDL",
	"PADDQ",
	"PADDSB",
	"PADDSW",
	"PADDUSB",
	"PADDUSW",
	"PADDW",
	"PANDB",
	"PANDL",
	"PANDSB",
	"PANDSW",
	"PANDUSB",
	"PANDUSW",
	"PANDW",
	"PAND",
	"PANDN",
	"PAVGB",
	"PAVGW",
	"PCMPEQB",
	"PCMPEQL",
	"PCMPEQW",
	"PCMPGTB",
	"PCMPGTL",
	"PCMPGTW",
	"PEXTRW",
	"PFACC",
	"PFADD",
	"PFCMPEQ",
	"PFCMPGE",
	"PFCMPGT",
	"PFMAX",
	"PFMIN",
	"PFMUL",
	"PFNACC",
	"PFPNACC",
	"PFRCP",
	"PFRCPIT1",
	"PFRCPI2T",
	"PFRSQIT1",
	"PFRSQRT",
	"PFSUB",
	"PFSUBR",
	"PINSRW",
	"PINSRD",
	"PINSRQ",
	"PMADDWL",
	"PMAXSW",
	"PMAXUB",
	"PMINSW",
	"PMINUB",
	"PMOVMSKB",
	"PMULHRW",
	"PMULHUW",
	"PMULHW",
	"PMULLW",
	"PMULULQ",
	"POR",
	"PSADBW",
	"PSHUFHW",
	"PSHUFL",
	"PSHUFLW",
	"PSHUFW",
	"PSHUFB",
	"PSLLO",
	"PSLLL",
	"PSLLQ",
	"PSLLW",
	"PSRAL",
	"PSRAW",
	"PSRLO",
	"PSRLL",
	"PSRLQ",
	"PSRLW",
	"PSUBB",
	"PSUBL",
	"PSUBQ",
	"PSUBSB",
	"PSUBSW",
	"PSUBUSB",
	"PSUBUSW",
	"PSUBW",
	"PSWAPL",
	"PUNPCKHBW",
	"PUNPCKHLQ",
	"PUNPCKHQDQ",
	"PUNPCKHWL",
	"PUNPCKLBW",
	"PUNPCKLLQ",
	"PUNPCKLQDQ",
	"PUNPCKLWL",
	"PXOR",
	"RCPPS",
	"RCPSS",
	"RSQRTPS",
	"RSQRTSS",
	"SHUFPD",
	"SHUFPS",
	"SQRTPD",
	"SQRTPS",
	"SQRTSD",
	"SQRTSS",
	"STMXCSR",
	"SUBPD",
	"SUBPS",
	"SUBSD",
	"SUBSS",
	"UCOMISD",
	"UCOMISS",
	"UNPCKHPD",
	"UNPCKHPS",
	"UNPCKLPD",
	"UNPCKLPS",
	"XORPD",
	"XORPS",
	"PF2IW",
	"PF2IL",
	"PI2FW",
	"PI2FL",
	"RETFW",
	"RETFL",
	"RETFQ",
	"SWAPGS",
	"MODE",
	"CRC32B",
	"CRC32Q",
	"IMUL3Q",
	"PREFETCHT0",
	"PREFETCHT1",
	"PREFETCHT2",
	"PREFETCHNTA",
	"MOVQL",
	"BSWAPL",
	"BSWAPQ",
	"AESENC",
	"AESENCLAST",
	"AESDEC",
	"AESDECLAST",
	"AESIMC",
	"AESKEYGENASSIST",
	"ROUNDPS",
	"ROUNDSS",
	"ROUNDPD",
	"ROUNDSD",
	"PSHUFD",
	"PCLMULQDQ",
	"JCXZW",
	"FCMOVCC",
	"FCMOVCS",
	"FCMOVEQ",
	"FCMOVHI",
	"FCMOVLS",
	"FCMOVNE",
	"FCMOVNU",
	"FCMOVUN",
	"FCOMI",
	"FCOMIP",
	"FUCOMI",
	"FUCOMIP",
	"LAST",
}

func (name Instr) String() string {
	return InstrString[int(name)]
}

func GetInstrFromStr(name string) (Instr, error) {
	for i, n := range InstrString {
		if n == name {
			return Instr(i), nil
		}
	}
	return Instr(0), fmt.Errorf("Couldn't find Instr for instr:%v", name)
}

// asmZeroMemory generates "MOVQ $0, name+offset(REG)" instructions,
// size is in bytes
func asmZeroMemory(indent string, name string, offset int, size uint, reg *register) string {

	var dataSize uint
	dataSize = 1

	if size%8 == 0 {
		dataSize = 8
	} else if size%4 == 0 {
		dataSize = 4
	} else if size%2 == 0 {
		dataSize = 2
	}

	asm := ""
	mov := GetInstr(I_MOV, dataSize).String()

	for i := uint(0); i < size/dataSize; i++ {
		ioffset := int(i*dataSize) + offset
		asm += indent
		asm += fmt.Sprintf("%v    $0, %v+%v(%v)\n", mov, name, ioffset, reg.name)
	}

	return strings.Replace(asm, "+-", "-", -1)
}

// asmZeroReg generates "XORQ reg, reg" instructions
func asmZeroReg(indent string, reg *register) string {
	xor := GetInstr(I_XOR, reg.width/8)
	return indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), reg.name, reg.name)
}

func asmMovRegReg(indent string, srcReg *register, dstReg *register, size uint) string {
	if srcReg.width != dstReg.width || size*8 > srcReg.width {
		panic(fmt.Sprintf("(%v) srcReg.width != (%v) dstReg.width or invalid size in asmMoveRegToReg", srcReg.width, dstReg.width))
	}

	mov := GetInstr(I_MOV, size).String()
	asm := indent + fmt.Sprintf("%v    %v, %v\n", mov, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovRegMem(indent string, srcReg *register, dstName string, dstReg *register, dstOffset int, size uint) string {
	if srcReg.width < size*8 {
		panic("srcReg.width < size * 8")
	}
	if size == 0 {
		panic("size == 0")
	}
	mov := GetInstr(I_MOV, size)
	asm := indent + fmt.Sprintf("// BEGIN asmMovRegMem, size (%v), mov (%v), mov.String (%v)\n", size, mov, mov.String())
	asm += indent + fmt.Sprintf("%v    %v, %v+%v(%v)\n", mov.String(), srcReg.name, dstName, dstOffset, dstReg.name)
	asm += indent + fmt.Sprintf("// END asmMovRegMem, size (%v), mov (%v), mov.String (%v)\n", size, mov, mov.String())
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovRegMemIndirect(indent string, srcReg *register, dstName string, dstReg *register, dstOffset int, tmp *register) string {
	if tmp.width != srcReg.width {
		panic("Invalid register width for asmMovRegMemIndirect")
	}
	mov := GetInstr(I_MOV, srcReg.width/8)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("%v    %v, (%v)\n", mov.String(), srcReg.name, tmp.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemMem(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMemMem")
	}
	if size%8 != 0 {
		panic("Invalid size for asmMovMemMem")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v+%v(%v)\n", srcName, srcOffset, srcReg.name, dstName, dstOffset, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemReg(indent string, srcName string, srcOffset int, srcReg *register, size uint, dstReg *register) string {
	mov := GetInstr(I_MOV, size)
	asm := indent + fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemMemIndirect(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMovMemMemIndirect")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    %v+%v(%v), (%v)\n", srcName, srcOffset, srcReg.name, tmp)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemIndirectReg(indent string, srcName string, srcOffset int, srcReg *register, dstReg *register, tmp *register) string {
	if dstReg.width != tmp.width {
		panic("Invalid register width for asmMovMemIndirectReg")
	}
	mov := GetInstr(I_MOV, dstReg.width/8)
	asm := indent
	asm += fmt.Sprintf("%v    %v+%v(%v), %v\n", mov.String(), srcName, srcOffset, srcReg, tmp)

	asm += indent
	asm += fmt.Sprintf("%v    (%v), %v\n", mov.String(), tmp, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmMovMemIndirectMem(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, size uint, tmp1 *register, tmp2 *register) string {
	if tmp1.width != tmp2.width {
		panic("Mismatched register widths in asmMovMemIndirectMem")
	}
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMovMemIndirectMem")
	}
	if size%(tmp1.width/8) != 0 {
		panic("Invalid size in asmMovMemIndirectMem")
	}
	mov := GetInstr(I_MOV, tmp1.width/8)
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

func asmMovMemIndirectMemIndirect(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp1 *register, tmp2 *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	mov := GetInstr(I_MOV, tmp1.width/8)
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
// flag is clear (ie the previous compare had its src greater than it's sub reg).
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
	lea := GetInstr(I_LEA, srcReg.width/8)
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

func asmAddRegReg(indent string, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width for asmAddRegReg")
	}
	add := GetInstr(I_ADD, srcReg.width/8)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", add.String(), srcReg.name, dstReg.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmSubRegReg(indent string, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width {
		panic("Invalid register width for asmSubRegReg")
	}
	sub := GetInstr(I_SUB, srcReg.width/8)
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
func asmMulRegReg(indent string, signed bool, src *register, dst *register, size uint) string {

	if dst.width != 64 {
		panic("Invalid register width for asmMulRegReg")
	}
	rax := getRegister(REG_AX)
	rdx := getRegister(REG_DX)
	if rax.width != 64 || rdx.width != 64 {
		panic("Invalid rax or rdx register width in asmMulRegReg")
	}

	// rax is the implicit destination for MULQ
	asm := asmMovRegReg(indent, dst, rax, size)
	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr TInstruction
	if signed {
		tinstr = I_IMUL
	} else {
		tinstr = I_MUL
	}
	mul := GetInstr(tinstr, size).String()
	asm += indent + fmt.Sprintf("%v    %v\n", mul, src.name)
	asm += asmMovRegReg(indent, rax, dst, size)
	return strings.Replace(asm, "+-", "-", -1)
}

// asmDivRegReg divides the "dividend" register by the "divisor" register and stores
// the quotient in rax and the remainder in rdx
func asmDivRegReg(indent string, signed bool, dividend *register, divisor *register, size uint) (asm string, rax *register, rdx *register) {

	if dividend.width != divisor.width || divisor.width < size*8 {
		panic("Invalid register width for asmDivRegReg")
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

	asm += asmMovRegReg(indent, dividend, rax, size)

	// the low order part of the result is stored in rax and the high order part
	// is stored in rdx
	var tinstr TInstruction
	if signed {
		tinstr = I_IDIV
	} else {
		tinstr = I_DIV
	}

	div := GetInstr(tinstr, size).String()
	asm += indent + fmt.Sprintf("%v    %v\n", div, divisor.name)
	return asm, rax, rdx
}

func asmArithOp(indent string, signed bool, op token.Token, x *register, y *register, result *register, size uint) string {
	if x.width != 64 || y.width != 64 || result.width != 64 {
		panic("Invalid register width in asmArithOp")
	}
	asm := ""
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmArithOp", op))
	case token.ADD:
		asm += asmMovRegReg(indent, x, result, size)
		asm += asmAddRegReg(indent, y, result)
	case token.SUB:
		asm += asmMovRegReg(indent, x, result, size)
		asm += asmSubRegReg(indent, y, result)
	case token.MUL:
		asm += asmMovRegReg(indent, x, result, size)
		asm += asmMulRegReg(indent, signed, y, result, size)
	case token.QUO, token.REM:
		// the quotient is stored in rax and
		// the remainder is stored in rdx.
		var rax, rdx *register
		a, rax, rdx := asmDivRegReg(indent, signed, x, y, size)
		asm += a
		if op == token.QUO {
			asm += asmMovRegReg(indent, rax, result, size)
		} else {
			asm += asmMovRegReg(indent, rdx, result, size)
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
	and := GetInstr(I_AND, size)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", and.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmOrRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width for asmOrRegReg")
	}
	or := GetInstr(I_OR, src.width/8)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", or.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmXorRegReg(indent string, src *register, dst *register) string {
	if src.width != dst.width {
		panic("Invalid register width for asmXorRegReg")
	}
	xor := GetInstr(I_XOR, src.width/8)
	asm := indent + fmt.Sprintf("%v    %v, %v\n", xor.String(), src.name, dst.name)
	return strings.Replace(asm, "+-", "-", -1)
}

func asmNotReg(indent string, reg *register, size uint) string {
	if reg.width/8 < size {
		panic(fmt.Sprintf("Bad reg width (%v), size (%v)", reg.width, size))
	}
	xor := GetInstr(I_XOR, size)
	asm := indent + fmt.Sprintf("%v    $-1, %v\n", xor.String(), reg.name)
	return asm
}

const (
	SHIFT_LEFT = iota
	SHIFT_RIGHT
)

func asmMovZeroExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
	var opcode TInstruction
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

	movzx := GetInstr(opcode, dstSize)
	asm := indent + fmt.Sprintf("%v %v, %v\n", movzx.String(), src.name, dst.name)
	return asm
}

func asmMovSignExtend(indent string, src *register, dst *register, srcSize uint, dstSize uint) string {
	var opcode TInstruction
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
	movsx := GetInstr(opcode, dstSize)
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

	var opcode TInstruction
	if direction == SHIFT_LEFT {
		opcode = I_SHL
	} else if !signed && direction == SHIFT_RIGHT {
		opcode = I_SHR
	} else if signed && direction == SHIFT_RIGHT {
		opcode = I_SAR
	}

	shift := GetInstr(opcode, size)
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

	asm += asmMovRegReg(indent, shiftReg, cx, size)

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
		asm = asmMovRegReg(indent, y, result, size)
		asm += asmAndRegReg(indent, x, result, size)
	case token.OR:
		asm = asmMovRegReg(indent, y, result, size)
		asm += asmOrRegReg(indent, x, result)
	case token.XOR:
		asm = asmMovRegReg(indent, y, result, size)
		asm += asmXorRegReg(indent, x, result)
	case token.SHL:
		asm = asmMovRegReg(indent, x, result, size)
		tmp := x
		asm += asmShiftRegReg(indent, signed, SHIFT_LEFT, result, y, tmp, size)
	case token.SHR:
		asm = asmMovRegReg(indent, x, result, size)
		tmp := x
		asm += asmShiftRegReg(indent, signed, SHIFT_RIGHT, result, y, tmp, size)
	case token.AND_NOT:
		asm = asmMovRegReg(indent, y, result, size)
		asm += asmAndNotRegReg(indent, x, result, size)

	}
	return strings.Replace(asm, "+-", "-", -1)
}

func asmCmpRegReg(indent string, x *register, y *register) string {
	if x.width != y.width {
		panic("Invalid register width for asmCmpRegReg")
	}
	cmp := GetInstr(I_CMP, x.width/8)
	asm := fmt.Sprintf("%v	%v, %v\n", cmp.String(), x.name, y.name)
	return strings.Replace(asm, "+-", "-", -1)

}

func asmCmpMemImm32(indent string, name string, offset int32, r *register, imm32 uint32) string {
	if r.width != 64 {
		panic("Invalid register width for asmCmpMemImm32")
	}
	asm := indent + fmt.Sprintf("CMPQ	%v+%v(%v), $%v\n", name, offset, r.name, imm32)
	return strings.Replace(asm, "+-", "-", -1)

}

func asmCmpRegImm32(indent string, r *register, imm32 uint32, size uint) string {
	if r.width != 64 {
		panic("Invalid register width for asmCmpMemImm32")
	}
	cmp := "CMPQ"
	if size == 1 {
		cmp = "CMPB"
	} else if size == 2 {
		cmp = "CMPW"
	} else if size == 4 {
		cmp = "CMPL"
	}
	asm := indent + fmt.Sprintf("%v	%v, $%v\n", cmp, r.name, imm32)
	return asm

}

func asmCmpOp(indent string, op token.Token, x *register, y *register, result *register) string {
	if x.width != y.width || x.width != result.width {
		panic("Invalid register width in asmCmpOp")
	}
	asm := ""
	asm += indent + asmCmpRegReg(indent, x, y)
	switch op {
	default:
		panic(fmt.Sprintf("Unknown Op token (%v) in asmComparisonOp", op))
	case token.EQL:
		asm += indent + fmt.Sprintf("SETEQ   %v\n", result.name)
	case token.NEQ:
		asm += indent + fmt.Sprintf("SETNEQ  %v\n", result.name)
	case token.LEQ:
		asm += indent + fmt.Sprintf("SETLS   %v\n", result.name)
	case token.GEQ:
		asm += indent + fmt.Sprintf("SETCC   %v\n", result.name)
	case token.LSS:
		asm += indent + fmt.Sprintf("SETCS   %v\n", result.name)
	case token.GTR:
		asm += indent + fmt.Sprintf("SETHI   %v\n", result.name)
	}
	return strings.Replace(asm, "+-", "-", -1)
}

func asmRet(indent string) string {
	asm := indent + fmt.Sprintf("RET\n")
	return asm
}
