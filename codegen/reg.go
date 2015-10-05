package codegen

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name string
	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// size of register in bits, e.g. eax is 32.
	width int
}

type RegType int

const (
	// integer register
	DataReg = RegType(iota)
	// address register
	AddrReg
	// Stack pointer register
	SpReg
	// xmm register
	XmmReg
)

// width in bytes
const DataRegSize = 8
const XmmRegSize = 16

const NumDataRegs = 14
const NumXmmRegs = 16

var registers = []register{
	{"AL", DataReg, 32},
	{"CL", DataReg, 32},
	{"DL", DataReg, 32},
	{"BL", DataReg, 32},

	{"AX", DataReg, 64},
	{"CX", DataReg, 64},
	{"DX", DataReg, 64},
	{"SI", AddrReg, 64},
	{"DI", AddrReg, 64},
	{"BX", AddrReg, 64},
	{"BP", AddrReg, 64},
	{"SP", SpReg, 64},
	{"R8", DataReg, 64},
	{"R9", DataReg, 64},
	{"R10", DataReg, 64},
	{"R11", DataReg, 64},
	{"R12", DataReg, 64},
	{"R13", DataReg, 64},
	{"R14", DataReg, 64},
	{"R15", DataReg, 64},
	{"X0", XmmReg, 128},
	{"X1", XmmReg, 128},
	{"X2", XmmReg, 128},
	{"X3", XmmReg, 128},
	{"X4", XmmReg, 128},
	{"X5", XmmReg, 128},
	{"X6", XmmReg, 128},
	{"X7", XmmReg, 128},
	{"X8", XmmReg, 128},
	{"X9", XmmReg, 128},
	{"X10", XmmReg, 128},
	{"X11", XmmReg, 128},
	{"X12", XmmReg, 128},
	{"X13", XmmReg, 128},
	{"X14", XmmReg, 128},
	{"X15", XmmReg, 128},
}
