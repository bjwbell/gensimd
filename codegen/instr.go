package codegen

import (
	"errors"
	"fmt"
	"log"
)

type Operand struct {
	Type   OperandType
	Input  bool
	Output bool

	Value func() string
}

type OperandType int

const (
	// al register
	AL = iota
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
	switch int(op) {
	default:
		log.Fatalf("Unknown OperandType: \"%v\"", int(op))
		return ""
	case AL:
		return "AL"
	case CL:
		return "AL"
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

type InstrName int

const (
	AAA = iota
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

var InstructionNames = []string{
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

func (name InstrName) String() string {
	return InstructionNames[int(name)]
}

func GetInstrName(name string) (InstrName, error) {
	for i, n := range InstructionNames {
		if n == name {
			return InstrName(i), nil
		}
	}
	return InstrName(0), errors.New(fmt.Sprintf("Couldn't find InstrName for instr:%v", name))
}

// asmZeroMemory generates "MOVQ $0, name+offset(REG)" instructions
func asmZeroMemory(indent string, name string, offset uint, size uint, reg register) string {
	if reg.width != 64 {
		panic(fmt.Sprintf("Invalid register width (%v) for asmZeroMemory", reg.width))
	}
	if size%(reg.width/8) != 0 {
		panic(fmt.Sprintf("Size (%v) not multiple of reg.width (%v)", size, reg.width/8))
	}
	asm := ""
	for i := uint(0); i < size/(reg.width/uint(8)); i++ {
		asm += indent + fmt.Sprintf("MOVQ    $0, %v+%v(%v)\n", name, i*reg.width/8+offset, reg.name)
	}
	return asm
}

// asmZeroReg generates "XORQ reg, reg" instructions
func asmZeroReg(indent string, reg *register) string {
	if reg.width != 64 {
		panic(fmt.Sprintf("Invalid register width (%v) for asmZeroReg", reg.width))
	}
	return indent + fmt.Sprintf("XORQ    %v, %v\n", reg.name, reg.name)
}

func asmCopyRegValueToMemory(indent string, reg register, name string, offset uint, memReg register) string {
	if reg.width != 64 || memReg.width != 64 {
		panic(fmt.Sprintf("Invalid register width (%v) for asmCopyRegValueToMemory", reg.width))
	}
	asm := indent + fmt.Sprintf("MOVQ    %v, %v+%v(%v)\n", reg.name, name, offset, memReg.name)
	return asm
}

func asmCopyIndirectRegToMemory(indent string, reg register, regOffset uint, name string, memOffset uint, memReg register) string {
	if reg.width != 64 || memReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v(%v), %v+%v(%v)\n", regOffset, reg.name, name, memOffset, memReg.name)
	return asm
}

func asmCopyMemoryToReg(indent string, name string, offset uint, memReg register, reg register) string {
	if reg.width != 64 || memReg.width != 64 {
		panic(fmt.Sprintf("Invalid register width (%v) for asmCopyRegValueToMemory", reg.width))
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", name, offset, memReg.name, reg.name)
	return asm
}

func asmCopyIndirectMemToReg(indent string, name string, memOffset uint, memReg register, reg register) string {
	if reg.width != 64 || memReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", name, memOffset, memReg.name, reg.name)
	asm += indent + fmt.Sprintf("MOVQ    (%v), %v\n", reg.name, reg.name)
	return asm
}

func asmMovRegReg(indent string, srcReg *register, dstReg *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMoveRegToReg")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v, %v\n", srcReg.name, dstReg.name)
	return asm
}

func asmMovRegMem(indent string, srcReg *register, dstName string, dstReg *register, dstOffset uint) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMoveRegToReg")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v, %v+%v(%v)\n", srcReg.name, dstName, dstOffset, dstReg.name)
	return asm
}

func asmMovRegMemIndirect(indent string, srcReg *register, dstName string, dstReg *register, dstOffset int, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 || tmp.width != 64 {
		panic("Invalid register width for asmMoveRegToReg")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    %v, (%v)\n", srcReg.name, tmp.name)
	return asm
}

func asmMovMemMem(indent string, srcName string, srcOffset uint, srcReg *register, dstName string, dstOffset uint, dstReg *register, size uint) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmMemMem")
	}
	if size%8 != 0 {
		panic("Invalid size for asmMovMemMem")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v+%v(%v)\n", srcName, srcOffset, srcReg.name, dstName, dstOffset, dstReg.name)
	return asm
}

func asmMovMemReg(indent string, srcName string, srcOffset uint, srcReg *register, dstReg *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v\n", srcName, srcOffset, srcReg.name, dstReg.name)
	return asm
}

func asmMovMemMemIndirect(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", dstName, dstOffset, dstReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    %v+%v(%v), (%v)\n", srcName, srcOffset, srcReg.name, tmp)
	return asm
}

func asmMovMemIndirectReg(indent string, srcName string, srcOffset int, srcReg *register, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", srcName, srcOffset, srcReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    (%v), %v\n", tmp, dstReg.name)
	return asm
}

func asmMovMemIndirectMem(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", srcName, srcOffset, srcReg, tmp)
	asm += indent + fmt.Sprintf("MOVQ    (%v), %v+%v(%v)\n", tmp, dstName, dstOffset, dstReg.name)
	return asm
}

func asmMovMemIndirectMemIndirect(indent string, srcName string, srcOffset int, srcReg *register, dstName string, dstOffset int, dstReg *register, tmp1 *register, tmp2 *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmCopyIndirectRegValueToMemory")
	}
	asm := indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", srcName, srcOffset, srcReg, tmp1)
	asm += indent + fmt.Sprintf("MOVQ    %v+%v(%v), %v", dstName, dstOffset, dstReg, tmp2)
	asm += indent + fmt.Sprintf("MOVQ    (%v), (%v)\n", tmp1.name, tmp2.name)
	return asm
}

func asmLoadImm32(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmLoadImm32")
	}
	asm := indent + fmt.Sprintf("LOADQ    $%v, %v\n", imm32, dstReg.name)
	return asm
}

func asmLea(indent string, srcName string, srcOffset uint, srcReg *register, dstReg *register) string {
	if srcReg.width != 64 || dstReg.width != 64 {
		panic("Invalid register width for asmLea")
	}
	asm := indent + fmt.Sprintf("LEAQ    %v+%v(%v), %v\n", srcName, srcOffset, srcReg.name, dstReg.name)
	return asm
}

func asmAddImm32Reg(indent string, imm32 uint32, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmAddImm32Reg")
	}
	asm := indent + fmt.Sprintf("ADDQ    $%v, %v\n", imm32, dstReg.name)
	return asm
}

func asmAddRegReg(indent string, srcReg *register, dstReg *register) string {
	if dstReg.width != srcReg.width || srcReg.width != 64 {
		panic("Invalid register width for asmAddRegReg")
	}
	asm := indent + fmt.Sprintf("ADDQ    %v, %v\n", srcReg.name, dstReg.name)
	return asm
}

func asmMultImm32RegReg(indent string, imm32 uint32, srcReg *register, dstReg *register) string {
	if dstReg.width < 32 {
		panic("Invalid register width for asmAddImm32Reg")
	}
	asm := indent + fmt.Sprintf("IMUL3Q    $%v, %v, %v\n", imm32, srcReg.name, dstReg.name)
	return asm
}
