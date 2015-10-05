package codegen

type register struct {
	name string
	typ  RegType
}

type RegType int

const (
	DataReg = RegType(iota)
	XmmReg
)

const DataRegSize = 8
const XmmRegSize = 16

const NumDataRegs = 14
const NumXmmRegs = 16

var registers = []register{
	{"AX", DataReg},
	{"CX", DataReg},
	{"DX", DataReg},
	{"BX", DataReg},
	{"SI", DataReg},
	{"DI", DataReg},
	{"R8", DataReg},
	{"R9", DataReg},
	{"R10", DataReg},
	{"R11", DataReg},
	{"R12", DataReg},
	{"R13", DataReg},
	{"R14", DataReg},
	{"R15", DataReg},
	{"X0", XmmReg},
	{"X1", XmmReg},
	{"X2", XmmReg},
	{"X3", XmmReg},
	{"X4", XmmReg},
	{"X5", XmmReg},
	{"X6", XmmReg},
	{"X7", XmmReg},
	{"X8", XmmReg},
	{"X9", XmmReg},
	{"X10", XmmReg},
	{"X11", XmmReg},
	{"X12", XmmReg},
	{"X13", XmmReg},
	{"X14", XmmReg},
	{"X15", XmmReg},
}
