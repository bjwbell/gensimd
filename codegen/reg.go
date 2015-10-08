package codegen

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name string

	regconst regconst

	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// width of register in bits, e.g. eax is 32.
	width uint
}

type RegType int

const (
	// integer register
	DataReg = RegType(iota)
	// address register
	AddrReg
	// Stack pointer pseudo register
	SpReg
	// Frame pointer pseudo register
	FpReg
	// xmm register
	XmmReg
)

// size in bytes
const DataRegSize = 8
const XmmRegSize = 16

const NumDataRegs = 14
const NumXmmRegs = 16

type regconst int

const (
	REG_AL = regconst(iota)
	REG_CL
	REG_DL
	REG_BL
	REG_AX
	REG_CX
	REG_DX
	REG_SI
	REG_DI
	REG_BX
	REG_BP
	REG_SP // SP (stack pointer) is a pseudo register
	REG_FP // FP (frame pointer) is a pseudo register
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_X0
	REG_X1
	REG_X2
	REG_X3
	REG_X4
	REG_X5
	REG_X6
	REG_X7
	REG_X8
	REG_X9
	REG_X10
	REG_X11
	REG_X12
	REG_X13
	REG_X14
	REG_X15
)

var registers = []register{
	{"AL", AL, DataReg, 32},
	{"CL", REG_CL, DataReg, 32},
	{"DL", REG_DL, DataReg, 32},
	{"BL", REG_BL, DataReg, 32},

	{"AX", REG_AX, DataReg, 64},
	{"CX", REG_CX, DataReg, 64},
	{"DX", REG_DX, DataReg, 64},
	{"SI", REG_SI, AddrReg, 64},
	{"DI", REG_DI, AddrReg, 64},
	{"BX", REG_BX, AddrReg, 64},
	{"BP", REG_BP, AddrReg, 64},
	{"SP", REG_SP, SpReg, 64},
	{"FP", REG_FP, FpReg, 64},
	{"R8", REG_R8, DataReg, 64},
	{"R9", REG_R9, DataReg, 64},
	{"R10", REG_R10, DataReg, 64},
	{"R11", REG_R11, DataReg, 64},
	{"R12", REG_R12, DataReg, 64},
	{"R13", REG_R13, DataReg, 64},
	{"R14", REG_R14, DataReg, 64},
	{"R15", REG_R15, DataReg, 64},
	{"X0", REG_X0, XmmReg, 128},
	{"X1", REG_X1, XmmReg, 128},
	{"X2", REG_X2, XmmReg, 128},
	{"X3", REG_X3, XmmReg, 128},
	{"X4", REG_X4, XmmReg, 128},
	{"X5", REG_X5, XmmReg, 128},
	{"X6", REG_X6, XmmReg, 128},
	{"X7", REG_X7, XmmReg, 128},
	{"X8", REG_X8, XmmReg, 128},
	{"X9", REG_X9, XmmReg, 128},
	{"X10", REG_X10, XmmReg, 128},
	{"X11", REG_X11, XmmReg, 128},
	{"X12", REG_X12, XmmReg, 128},
	{"X13", REG_X13, XmmReg, 128},
	{"X14", REG_X14, XmmReg, 128},
	{"X15", REG_X15, XmmReg, 128},
}

var excludedRegisters = []register{
	// used as implicit operands in arithmetic instructions
	{"AL", REG_AL, DataReg, 32},
	{"CL", REG_CL, DataReg, 32},
	{"DL", REG_DL, DataReg, 32},
	{"AX", REG_AX, DataReg, 64},
	{"CX", REG_CX, DataReg, 64},
	{"DX", REG_DX, DataReg, 64},

	// stack pointer pseudo register
	{"SP", REG_SP, SpReg, 64},
	// frame pointer pseudo register
	{"FP", REG_FP, FpReg, 64},
}

func getRegister(reg regconst) *register {
	for _, r := range registers {
		if r.regconst == reg {
			return &r
		}
	}
	return nil
}
