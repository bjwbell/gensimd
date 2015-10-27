package codegen

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

	// 3DNow! instructions
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

	// 3DNow! instruction
	PMULHRW

	// Multiply Packed Unsigned Integers (16bit) and Store High Result
	// multiplies the packed unsigned word (16bit) integers in xmm1 and xmm2/m128
	// and store the high 16 bits of the results in xmm1
	// Intel instruction: PMULHUW
	PMULHUW

	// Multiply Packed Signed Integers (16bit) and Store High Result
	// multiplies the packed signed word (16bit) integers in xmm1 and xmm2/m128
	// and store the high 16 bits of the results in xmm1
	// Intel instruction: PMULHW
	PMULHW

	// Multiply Packed Signed Integers (16bit) and Store Low Result
	// multiplies the packed signed word (16bit) integers in xmm1 and xmm2/m128
	// and store the low 16 bits of the results in xmm1
	// Intel instruction: PMULLW
	PMULLW

	// Multiply Packed Unsigned Dword (32bit) Integers
	// multiplies two 32bit ints (in dwords 0 and 2 of the xmms registers)
	// and stores resulting two 64bit ints in result
	// xmm register/128bits mem
	// Intel instruction: PMULUDQ
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
	PCLMULQDQ // packed carryless quadword multiplication (xmm)
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

type InstrFlags int

const (
	FLAGS_NONE InstrFlags = 1 << iota
	SizeB                 // byte
	SizeW                 // word
	SizeL                 // dword
	SizeQ                 // quad word
	SizeO                 // oct word
	SizeD                 // float64
	SizeF                 // float32

	UseCarry
	SetCarry
	KillCarry

	ShiftCX
	ImulAXDX

	LeftRead
	LeftRdwr
	RightRead
	RightRdwr
	RightWrite

	LeftAddr
	RightAddr

	Call
	OK
	Conv

	Cjmp
	Jump
	Break

	Move
)

type InstrInfo struct {
	Flags InstrFlags
	Use   Reg
	Set   Reg
}

var instrTable = map[Instr]InstrInfo{
	ADCL:  {Flags: SizeL | LeftRead | RightRdwr | SetCarry | UseCarry},
	ADCW:  {Flags: SizeW | LeftRead | RightRdwr | SetCarry | UseCarry},
	ADDB:  {Flags: SizeB | LeftRead | RightRdwr | SetCarry},
	ADDL:  {Flags: SizeL | LeftRead | RightRdwr | SetCarry},
	ADDW:  {Flags: SizeW | LeftRead | RightRdwr | SetCarry},
	ADDQ:  {Flags: SizeQ | LeftRead | RightRdwr | SetCarry},
	ADDSD: {Flags: SizeD | LeftRead | RightRdwr},
	ADDSS: {Flags: SizeF | LeftRead | RightRdwr},
	ADDPS: {Flags: SizeF | LeftRead | RightRdwr},
	ADDPD: {Flags: SizeD | LeftRead | RightRdwr},
	ANDB:  {Flags: SizeB | LeftRead | RightRdwr | SetCarry},
	ANDL:  {Flags: SizeL | LeftRead | RightRdwr | SetCarry},
	ANDQ:  {Flags: SizeQ | LeftRead | RightRdwr | SetCarry},
	ANDW:  {Flags: SizeW | LeftRead | RightRdwr | SetCarry},
	//CALL:      {Flags: RightAddr | Call | KillCarry},
	CDQ:       {Flags: OK, Use: REG_AX, Set: REG_AX | REG_DX},
	CWD:       {Flags: OK, Use: REG_AX, Set: REG_AX | REG_DX},
	CLD:       {Flags: OK},
	STD:       {Flags: OK},
	CMOVQCC:   {Flags: SizeQ | LeftRead | RightWrite | Move},
	CMOVLCC:   {Flags: SizeQ | LeftRead | RightWrite | Move},
	CMOVWCC:   {Flags: SizeQ | LeftRead | RightWrite | Move},
	CMPB:      {Flags: SizeB | LeftRead | RightRead | SetCarry},
	CMPL:      {Flags: SizeL | LeftRead | RightRead | SetCarry},
	CMPW:      {Flags: SizeW | LeftRead | RightRead | SetCarry},
	CMPQ:      {Flags: SizeQ | LeftRead | RightRead | SetCarry},
	COMISD:    {Flags: SizeD | LeftRead | RightRead | SetCarry},
	COMISS:    {Flags: SizeF | LeftRead | RightRead | SetCarry},
	CVTSD2SL:  {Flags: SizeL | LeftRead | RightWrite | Conv},
	CVTSD2SS:  {Flags: SizeF | LeftRead | RightWrite | Conv},
	CVTSL2SD:  {Flags: SizeD | LeftRead | RightWrite | Conv},
	CVTSL2SS:  {Flags: SizeF | LeftRead | RightWrite | Conv},
	CVTSS2SD:  {Flags: SizeD | LeftRead | RightWrite | Conv},
	CVTSS2SL:  {Flags: SizeL | LeftRead | RightWrite | Conv},
	CVTSQ2SS:  {Flags: SizeF | LeftRead | RightWrite | Conv},
	CVTSQ2SD:  {Flags: SizeD | LeftRead | RightWrite | Conv},
	CVTTSD2SL: {Flags: SizeL | LeftRead | RightWrite | Conv},
	CVTTSS2SL: {Flags: SizeL | LeftRead | RightWrite | Conv},
	CVTTSS2SQ: {Flags: SizeQ | LeftRead | RightWrite | Conv},
	CVTTSD2SQ: {Flags: SizeQ | LeftRead | RightWrite | Conv},
	DECB:      {Flags: SizeB | RightRdwr},
	DECL:      {Flags: SizeL | RightRdwr},
	DECW:      {Flags: SizeW | RightRdwr},
	DIVB:      {Flags: SizeB | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX},
	DIVL:      {Flags: SizeL | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	DIVQ:      {Flags: SizeQ | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	DIVW:      {Flags: SizeW | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	DIVSD:     {Flags: SizeD | LeftRead | RightRdwr},
	DIVSS:     {Flags: SizeF | LeftRead | RightRdwr},
	DIVPS:     {Flags: SizeF | LeftRead | RightRdwr},
	DIVPD:     {Flags: SizeD | LeftRead | RightRdwr},
	FLDCW:     {Flags: SizeW | LeftAddr},
	FSTCW:     {Flags: SizeW | RightAddr},
	FSTSW:     {Flags: SizeW | RightAddr | RightWrite},
	FADDD:     {Flags: SizeD | LeftAddr | RightRdwr},
	FADDDP:    {Flags: SizeD | LeftAddr | RightRdwr},
	FADDF:     {Flags: SizeF | LeftAddr | RightRdwr},
	FCOMD:     {Flags: SizeD | LeftAddr | RightRead},
	FCOMDP:    {Flags: SizeD | LeftAddr | RightRead},
	FCOMDPP:   {Flags: SizeD | LeftAddr | RightRead},
	FCOMF:     {Flags: SizeF | LeftAddr | RightRead},
	FCOMFP:    {Flags: SizeF | LeftAddr | RightRead},
	FUCOMIP:   {Flags: SizeF | LeftAddr | RightRead},
	FCHS:      {Flags: SizeD | RightRdwr}, // also SizeF

	FDIVDP:  {Flags: SizeD | LeftAddr | RightRdwr},
	FDIVF:   {Flags: SizeF | LeftAddr | RightRdwr},
	FDIVD:   {Flags: SizeD | LeftAddr | RightRdwr},
	FDIVRDP: {Flags: SizeD | LeftAddr | RightRdwr},
	FDIVRF:  {Flags: SizeF | LeftAddr | RightRdwr},
	FDIVRD:  {Flags: SizeD | LeftAddr | RightRdwr},
	FXCHD:   {Flags: SizeD | LeftRdwr | RightRdwr},
	FSUBD:   {Flags: SizeD | LeftAddr | RightRdwr},
	FSUBDP:  {Flags: SizeD | LeftAddr | RightRdwr},
	FSUBF:   {Flags: SizeF | LeftAddr | RightRdwr},
	FSUBRD:  {Flags: SizeD | LeftAddr | RightRdwr},
	FSUBRDP: {Flags: SizeD | LeftAddr | RightRdwr},
	FSUBRF:  {Flags: SizeF | LeftAddr | RightRdwr},
	FMOVD:   {Flags: SizeD | LeftAddr | RightWrite},
	FMOVF:   {Flags: SizeF | LeftAddr | RightWrite},
	FMOVL:   {Flags: SizeL | LeftAddr | RightWrite},
	FMOVW:   {Flags: SizeW | LeftAddr | RightWrite},
	FMOVV:   {Flags: SizeQ | LeftAddr | RightWrite},

	FMOVDP: {Flags: SizeD | LeftRead | RightWrite},
	FMOVFP: {Flags: SizeF | LeftRead | RightWrite},
	FMOVLP: {Flags: SizeL | LeftRead | RightWrite},
	FMOVWP: {Flags: SizeW | LeftRead | RightWrite},
	FMOVVP: {Flags: SizeQ | LeftRead | RightWrite},
	FMULD:  {Flags: SizeD | LeftAddr | RightRdwr},
	FMULDP: {Flags: SizeD | LeftAddr | RightRdwr},
	FMULF:  {Flags: SizeF | LeftAddr | RightRdwr},
	IDIVB:  {Flags: SizeB | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX},
	IDIVL:  {Flags: SizeL | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	IDIVQ:  {Flags: SizeQ | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	IDIVW:  {Flags: SizeW | LeftRead | SetCarry, Use: REG_AX | REG_DX, Set: REG_AX | REG_DX},
	IMULB:  {Flags: SizeB | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX},
	IMULL:  {Flags: SizeL | LeftRead | ImulAXDX | SetCarry},
	IMULQ:  {Flags: SizeQ | LeftRead | ImulAXDX | SetCarry},
	IMULW:  {Flags: SizeW | LeftRead | ImulAXDX | SetCarry},
	INCB:   {Flags: SizeB | RightRdwr},
	INCL:   {Flags: SizeL | RightRdwr},
	INCW:   {Flags: SizeW | RightRdwr},
	JCC:    {Flags: Cjmp | UseCarry},
	JCS:    {Flags: Cjmp | UseCarry},
	JEQ:    {Flags: Cjmp | UseCarry},
	JGE:    {Flags: Cjmp | UseCarry},
	JGT:    {Flags: Cjmp | UseCarry},
	JHI:    {Flags: Cjmp | UseCarry},
	JLE:    {Flags: Cjmp | UseCarry},
	JLS:    {Flags: Cjmp | UseCarry},
	JLT:    {Flags: Cjmp | UseCarry},
	JMI:    {Flags: Cjmp | UseCarry},
	JNE:    {Flags: Cjmp | UseCarry},
	JOC:    {Flags: Cjmp | UseCarry},
	JOS:    {Flags: Cjmp | UseCarry},
	JPC:    {Flags: Cjmp | UseCarry},
	JPL:    {Flags: Cjmp | UseCarry},
	JPS:    {Flags: Cjmp | UseCarry},
	//JMP:      {Flags: Jump | Break | KillCarry},
	LEAL:    {Flags: LeftAddr | RightWrite},
	MOVBLSX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVBLZX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVBQSX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVBQZX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVBWSX: {Flags: SizeW | LeftRead | RightWrite | Conv},
	MOVBWZX: {Flags: SizeW | LeftRead | RightWrite | Conv},
	MOVLQSX: {Flags: SizeW | LeftRead | RightWrite | Conv},
	MOVLQZX: {Flags: SizeW | LeftRead | RightWrite | Conv},
	MOVWQSX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVWQZX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVWLSX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVWLZX: {Flags: SizeL | LeftRead | RightWrite | Conv},
	MOVB:    {Flags: SizeB | LeftRead | RightWrite | Move},
	MOVL:    {Flags: SizeL | LeftRead | RightWrite | Move},
	MOVOU:   {Flags: SizeO | LeftRead | RightWrite | Move},
	MOVO:    {Flags: SizeO | LeftRead | RightWrite | Move},
	MOVQ:    {Flags: SizeQ | LeftRead | RightWrite | Move},
	MOVW:    {Flags: SizeW | LeftRead | RightWrite | Move},
	MOVSB:   {Flags: OK, Use: REG_DI | REG_SI, Set: REG_DI | REG_SI},
	MOVSL:   {Flags: OK, Use: REG_DI | REG_SI, Set: REG_DI | REG_SI},
	MOVSW:   {Flags: OK, Use: REG_DI | REG_SI, Set: REG_DI | REG_SI},
	//DUFFCOPY: {Flags: OK, Use: REG_DI | REG_SI, Set: REG_DI | REG_SI | REG_CX},
	MOVSD:  {Flags: SizeD | LeftRead | RightWrite | Move},
	MOVSS:  {Flags: SizeF | LeftRead | RightWrite | Move},
	MOVUPS: {Flags: SizeD | LeftRead | RightWrite | Move},
	MOVUPD: {Flags: SizeD | LeftRead | RightWrite | Move},

	// We use MOVAPD as a faster synonym for MOVSD.
	MOVAPD:    {Flags: SizeD | LeftRead | RightWrite | Move},
	MULB:      {Flags: SizeB | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX},
	MULL:      {Flags: SizeL | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX | REG_DX},
	MULQ:      {Flags: SizeQ | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX | REG_DX},
	MULPS:     {Flags: SizeF | LeftRead | RightRdwr},
	MULPD:     {Flags: SizeF | LeftRead | RightRdwr},
	MULW:      {Flags: SizeW | LeftRead | SetCarry, Use: REG_AX, Set: REG_AX | REG_DX},
	MULSD:     {Flags: SizeD | LeftRead | RightRdwr},
	MULSS:     {Flags: SizeF | LeftRead | RightRdwr},
	NEGB:      {Flags: SizeB | RightRdwr | SetCarry},
	NEGL:      {Flags: SizeL | RightRdwr | SetCarry},
	NEGW:      {Flags: SizeW | RightRdwr | SetCarry},
	NOTB:      {Flags: SizeB | RightRdwr},
	NOTL:      {Flags: SizeL | RightRdwr},
	NOTW:      {Flags: SizeW | RightRdwr},
	ORB:       {Flags: SizeB | LeftRead | RightRdwr | SetCarry},
	ORL:       {Flags: SizeL | LeftRead | RightRdwr | SetCarry},
	ORQ:       {Flags: SizeQ | LeftRead | RightRdwr | SetCarry},
	ORW:       {Flags: SizeW | LeftRead | RightRdwr | SetCarry},
	PADDB:     {Flags: SizeO | LeftRead | RightRdwr | SetCarry},
	PADDL:     {Flags: SizeO | LeftRead | RightRdwr | SetCarry},
	PADDW:     {Flags: SizeO | LeftRead | RightRdwr | SetCarry},
	PADDQ:     {Flags: SizeO | LeftRead | RightRdwr | SetCarry},
	PEXTRW:    {Flags: SizeW | RightWrite},
	PINSRW:    {Flags: SizeW | RightWrite},
	PMULULQ:   {Flags: SizeO | LeftRead | RightRdwr},
	PMULLW:    {Flags: SizeO | LeftRead | RightRdwr},
	POPL:      {Flags: SizeL | RightWrite},
	PSLLL:     {Flags: SizeO | LeftRead | RightRdwr},
	PSLLQ:     {Flags: SizeO | LeftRead | RightRdwr},
	PSLLW:     {Flags: SizeO | LeftRead | RightRdwr},
	PSRAW:     {Flags: SizeO | LeftRead | RightRdwr},
	PSRAL:     {Flags: SizeO | LeftRead | RightRdwr},
	PSRLW:     {Flags: SizeO | LeftRead | RightRdwr},
	PSRLL:     {Flags: SizeO | LeftRead | RightRdwr},
	PSRLO:     {Flags: SizeO | LeftRead | RightRdwr},
	PSHUFD:    {Flags: SizeO | LeftRead | RightRdwr},
	PSUBB:     {Flags: SizeO | LeftRead | RightRdwr},
	PSUBW:     {Flags: SizeO | LeftRead | RightRdwr},
	PSUBL:     {Flags: SizeO | LeftRead | RightRdwr},
	PSUBQ:     {Flags: SizeO | LeftRead | RightRdwr},
	PUNPCKLLQ: {Flags: SizeO | LeftRead | RightRdwr},
	PUSHL:     {Flags: SizeL | LeftRead},
	RCLB:      {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	RCLL:      {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	RCLW:      {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	RCRB:      {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	RCRL:      {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	RCRW:      {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry | UseCarry},
	REP:       {Flags: OK, Use: REG_CX, Set: REG_CX},
	REPN:      {Flags: OK, Use: REG_CX, Set: REG_CX},
	//RET:      {Flags: Break | KillCarry},
	ROLB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	ROLL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	ROLW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	RORB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	RORL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	RORW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SAHF:  {Flags: OK, Use: REG_AX, Set: REG_AX},
	SALB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SALL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SALQ:  {Flags: SizeQ | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SALW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SARB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SARL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SARQ:  {Flags: SizeQ | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SARW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SBBB:  {Flags: SizeB | LeftRead | RightRdwr | SetCarry | UseCarry},
	SBBL:  {Flags: SizeL | LeftRead | RightRdwr | SetCarry | UseCarry},
	SBBW:  {Flags: SizeW | LeftRead | RightRdwr | SetCarry | UseCarry},
	SETCC: {Flags: SizeB | RightRdwr | UseCarry},
	SETCS: {Flags: SizeB | RightRdwr | UseCarry},
	SETEQ: {Flags: SizeB | RightRdwr | UseCarry},
	SETGE: {Flags: SizeB | RightRdwr | UseCarry},
	SETGT: {Flags: SizeB | RightRdwr | UseCarry},
	SETHI: {Flags: SizeB | RightRdwr | UseCarry},
	SETLE: {Flags: SizeB | RightRdwr | UseCarry},
	SETLS: {Flags: SizeB | RightRdwr | UseCarry},
	SETLT: {Flags: SizeB | RightRdwr | UseCarry},
	SETMI: {Flags: SizeB | RightRdwr | UseCarry},
	SETNE: {Flags: SizeB | RightRdwr | UseCarry},
	SETOC: {Flags: SizeB | RightRdwr | UseCarry},
	SETOS: {Flags: SizeB | RightRdwr | UseCarry},
	SETPC: {Flags: SizeB | RightRdwr | UseCarry},
	SETPL: {Flags: SizeB | RightRdwr | UseCarry},
	SETPS: {Flags: SizeB | RightRdwr | UseCarry},
	SHLB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHLL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHLQ:  {Flags: SizeQ | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHLW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHRB:  {Flags: SizeB | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHRL:  {Flags: SizeL | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHRQ:  {Flags: SizeQ | LeftRead | RightRdwr | ShiftCX | SetCarry},
	SHRW:  {Flags: SizeW | LeftRead | RightRdwr | ShiftCX | SetCarry},
	STOSB: {Flags: OK, Use: REG_AX | REG_DI, Set: REG_DI},
	STOSL: {Flags: OK, Use: REG_AX | REG_DI, Set: REG_DI},
	STOSW: {Flags: OK, Use: REG_AX | REG_DI, Set: REG_DI},
	//DUFFZERO: {Flags: OK, Use: REG_AX | REG_DI, Set: REG_DI},
	SUBB:    {Flags: SizeB | LeftRead | RightRdwr | SetCarry},
	SUBL:    {Flags: SizeL | LeftRead | RightRdwr | SetCarry},
	SUBQ:    {Flags: SizeQ | LeftRead | RightRdwr | SetCarry},
	SUBW:    {Flags: SizeW | LeftRead | RightRdwr | SetCarry},
	SUBSD:   {Flags: SizeD | LeftRead | RightRdwr},
	SUBSS:   {Flags: SizeF | LeftRead | RightRdwr},
	SUBPS:   {Flags: SizeF | LeftRead | RightRdwr},
	SUBPD:   {Flags: SizeF | LeftRead | RightRdwr},
	TESTB:   {Flags: SizeB | LeftRead | RightRead | SetCarry},
	TESTL:   {Flags: SizeL | LeftRead | RightRead | SetCarry},
	TESTW:   {Flags: SizeW | LeftRead | RightRead | SetCarry},
	UCOMISD: {Flags: SizeD | LeftRead | RightRead},
	UCOMISS: {Flags: SizeF | LeftRead | RightRead},
	XCHGB:   {Flags: SizeB | LeftRdwr | RightRdwr},
	XCHGL:   {Flags: SizeL | LeftRdwr | RightRdwr},
	XCHGW:   {Flags: SizeW | LeftRdwr | RightRdwr},
	XORB:    {Flags: SizeB | LeftRead | RightRdwr | SetCarry},
	XORL:    {Flags: SizeL | LeftRead | RightRdwr | SetCarry},
	XORW:    {Flags: SizeW | LeftRead | RightRdwr | SetCarry},
	XORQ:    {Flags: SizeQ | LeftRead | RightRdwr | SetCarry},
	XORPD:   {Flags: SizeD | LeftRead | RightRdwr | SetCarry},
	XORPS:   {Flags: SizeF | LeftRead | RightRdwr | SetCarry},
}
