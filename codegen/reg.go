package codegen

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name string
	used bool

	regconst regconst

	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// width of register in bits, e.g. ax is 64 on amd64.
	width uint
	// allowed data sizes in bytes used in reg
	datasizes []uint

	assigned *identifier
	loaded   bool
}

func (r *register) size() uint {
	return r.width / 8
}

func (r *register) modified() {
	r.assigned = nil
	r.loaded = false
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
const XmmRegSize = 16

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
var XmmDataSizes = []uint{1, 2, 4, 8, 16}

var registers = []register{
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, nil, false},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, nil, false},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, nil, false},
	{"BL", false, REG_BL, DATA_REG, 32, LongSizes, nil, false},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, nil, false},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, nil, false},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, nil, false},
	{"SI", false, REG_SI, ADDR_REG, 64, QuadSize, nil, false},
	{"DI", false, REG_DI, ADDR_REG, 64, QuadSize, nil, false},
	{"BX", false, REG_BX, ADDR_REG, 64, QuadSize, nil, false},
	{"BP", false, REG_BP, ADDR_REG, 64, QuadSize, nil, false},
	{"SP", false, REG_SP, SpReg, 64, QuadSize, nil, false},
	{"FP", false, REG_FP, FpReg, 64, QuadSize, nil, false},
	{"R8", false, REG_R8, DATA_REG, 64, QuadSizes, nil, false},
	{"R9", false, REG_R9, DATA_REG, 64, QuadSizes, nil, false},
	{"R10", false, REG_R10, DATA_REG, 64, QuadSizes, nil, false},
	{"R11", false, REG_R11, DATA_REG, 64, QuadSizes, nil, false},
	{"R12", false, REG_R12, DATA_REG, 64, QuadSizes, nil, false},
	{"R13", false, REG_R13, DATA_REG, 64, QuadSizes, nil, false},
	{"R14", false, REG_R14, DATA_REG, 64, QuadSizes, nil, false},
	{"R15", false, REG_R15, DATA_REG, 64, QuadSizes, nil, false},
	{"X0", false, REG_X0, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X1", false, REG_X1, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X2", false, REG_X2, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X3", false, REG_X3, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X4", false, REG_X4, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X5", false, REG_X5, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X6", false, REG_X6, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X7", false, REG_X7, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X8", false, REG_X8, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X9", false, REG_X9, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X10", false, REG_X10, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X11", false, REG_X11, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X12", false, REG_X12, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X13", false, REG_X13, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X14", false, REG_X14, XMM_REG, 128, XmmDataSizes, nil, false},
	{"X15", false, REG_X15, XMM_REG, 128, XmmDataSizes, nil, false},
}

var excludedRegisters = []register{
	// used as implicit operands in arithmetic instructions
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, nil, false},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, nil, false},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, nil, false},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, nil, false},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, nil, false},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, nil, false},

	// stack pointer pseudo register
	{"SP", false, REG_SP, SpReg, 64, QuadSize, nil, false},
	// frame pointer pseudo register
	{"FP", false, REG_FP, FpReg, 64, QuadSize, nil, false},
}

func getRegister(reg regconst) *register {
	for _, r := range registers {
		if r.regconst == reg {
			return &r
		}
	}
	return nil
}
