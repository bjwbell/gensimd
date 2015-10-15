package codegen

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name string

	regconst regconst

	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// width of register in bits, e.g. ax is 64 on amd64.
	width uint
	// allowed data sizes in bytes used in reg
	datasizes []uint
}

type RegType int

const (
	// integer register
	DATA_REG = RegType(iota)
	// address register
	ADDR_REG
	// Stack pointer pseudo register
	SpReg
	// Frame pointer pseudo register
	FpReg
	// xmm register
	XMM_REG
)

func (r RegType) String() string {
	switch r {
	case DATA_REG:
		return "DATA_REG"
	case ADDR_REG:
		return "ADDR_REG"
	case SpReg:
		return "SpReg"
	case XMM_REG:
		return "XMM_REG"
	}
	panic("Invalid regtype")
}

// size in bytes
const DataRegSize = 8
const XMM_REGSize = 16

const NumDataRegs = 14
const NumXMM_REGs = 16

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

var LongSizes = []uint{1, 2, 4}
var QuadSizes = []uint{1, 2, 4, 8}
var QuadSize = []uint{8}
var XmmDataSizes = []uint{4, 8, 16}

var registers = []register{
	{"AL", REG_AL, DATA_REG, 32, LongSizes},
	{"CL", REG_CL, DATA_REG, 32, LongSizes},
	{"DL", REG_DL, DATA_REG, 32, LongSizes},
	{"BL", REG_BL, DATA_REG, 32, LongSizes},
	{"AX", REG_AX, DATA_REG, 64, QuadSizes},
	{"CX", REG_CX, DATA_REG, 64, QuadSizes},
	{"DX", REG_DX, DATA_REG, 64, QuadSizes},
	{"SI", REG_SI, ADDR_REG, 64, QuadSize},
	{"DI", REG_DI, ADDR_REG, 64, QuadSize},
	{"BX", REG_BX, ADDR_REG, 64, QuadSize},
	{"BP", REG_BP, ADDR_REG, 64, QuadSize},
	{"SP", REG_SP, SpReg, 64, QuadSize},
	{"FP", REG_FP, FpReg, 64, QuadSize},
	{"R8", REG_R8, DATA_REG, 64, QuadSizes},
	{"R9", REG_R9, DATA_REG, 64, QuadSizes},
	{"R10", REG_R10, DATA_REG, 64, QuadSizes},
	{"R11", REG_R11, DATA_REG, 64, QuadSizes},
	{"R12", REG_R12, DATA_REG, 64, QuadSizes},
	{"R13", REG_R13, DATA_REG, 64, QuadSizes},
	{"R14", REG_R14, DATA_REG, 64, QuadSizes},
	{"R15", REG_R15, DATA_REG, 64, QuadSizes},
	{"X0", REG_X0, XMM_REG, 128, XmmDataSizes},
	{"X1", REG_X1, XMM_REG, 128, XmmDataSizes},
	{"X2", REG_X2, XMM_REG, 128, XmmDataSizes},
	{"X3", REG_X3, XMM_REG, 128, XmmDataSizes},
	{"X4", REG_X4, XMM_REG, 128, XmmDataSizes},
	{"X5", REG_X5, XMM_REG, 128, XmmDataSizes},
	{"X6", REG_X6, XMM_REG, 128, XmmDataSizes},
	{"X7", REG_X7, XMM_REG, 128, XmmDataSizes},
	{"X8", REG_X8, XMM_REG, 128, XmmDataSizes},
	{"X9", REG_X9, XMM_REG, 128, XmmDataSizes},
	{"X10", REG_X10, XMM_REG, 128, XmmDataSizes},
	{"X11", REG_X11, XMM_REG, 128, XmmDataSizes},
	{"X12", REG_X12, XMM_REG, 128, XmmDataSizes},
	{"X13", REG_X13, XMM_REG, 128, XmmDataSizes},
	{"X14", REG_X14, XMM_REG, 128, XmmDataSizes},
	{"X15", REG_X15, XMM_REG, 128, XmmDataSizes},
}

var excludedRegisters = []register{
	// used as implicit operands in arithmetic instructions
	{"AL", REG_AL, DATA_REG, 32, LongSizes},
	{"CL", REG_CL, DATA_REG, 32, LongSizes},
	{"DL", REG_DL, DATA_REG, 32, LongSizes},
	{"AX", REG_AX, DATA_REG, 64, QuadSizes},
	{"CX", REG_CX, DATA_REG, 64, QuadSizes},
	{"DX", REG_DX, DATA_REG, 64, QuadSizes},

	// stack pointer pseudo register
	{"SP", REG_SP, SpReg, 64, QuadSize},
	// frame pointer pseudo register
	{"FP", REG_FP, FpReg, 64, QuadSize},
}

func getRegister(reg regconst) *register {
	for _, r := range registers {
		if r.regconst == reg {
			return &r
		}
	}
	return nil
}
