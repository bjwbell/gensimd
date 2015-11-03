package codegen

import "fmt"

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name string
	// idicates register is being used in current SSA/assembly instruction
	inUse bool

	regconst Reg

	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// width of register in bits, e.g. ax is 64 on amd64.
	width uint
	// allowed data sizes in bytes used in reg
	datasizes []uint

	dirty  bool
	parent storer
}

func (r *register) size() uint {
	return r.width / 8
}

func (r *register) modified(ctx context, spill bool) string {
	if r.parent != nil {
		if r.parent.owner().f.Trace {
			fmt.Printf(r.parent.owner().f.Indent+"Modified %v (inUse %v, spill %v, old dirty %v)\n",
				r.name, r.inUse, spill, r.dirty)
		}
	}

	var asm string
	inUse := r.inUse
	if spill {
		asm = r.spill(ctx)
	}
	r.inUse = inUse
	r.dirty = true
	return asm
}

func (r *register) spill(ctx context) string {
	if r.parent != nil {
		return r.parent.spillRegister(ctx, r, false)
	} else {
		return ""
	}
}

type RegType int

const (
	INVALID_REG = RegType(iota)
	// integer register
	DATA_REG
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

type Reg int

const (
	REG_INVALID Reg = 1 << iota
	REG_AL
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
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, false, nil},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, false, nil},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, false, nil},
	{"BL", false, REG_BL, DATA_REG, 32, LongSizes, false, nil},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, false, nil},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, false, nil},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, false, nil},
	{"SI", false, REG_SI, DATA_REG, 64, QuadSize, false, nil},
	{"DI", false, REG_DI, DATA_REG, 64, QuadSize, false, nil},
	{"BX", false, REG_BX, DATA_REG, 64, QuadSize, false, nil},
	{"BP", false, REG_BP, DATA_REG, 64, QuadSize, false, nil},
	{"SP", false, REG_SP, SpReg, 64, QuadSize, false, nil},
	{"FP", false, REG_FP, FpReg, 64, QuadSize, false, nil},
	{"R8", false, REG_R8, DATA_REG, 64, QuadSizes, false, nil},
	{"R9", false, REG_R9, DATA_REG, 64, QuadSizes, false, nil},
	{"R10", false, REG_R10, DATA_REG, 64, QuadSizes, false, nil},
	{"R11", false, REG_R11, DATA_REG, 64, QuadSizes, false, nil},
	{"R12", false, REG_R12, DATA_REG, 64, QuadSizes, false, nil},
	{"R13", false, REG_R13, DATA_REG, 64, QuadSizes, false, nil},
	{"R14", false, REG_R14, DATA_REG, 64, QuadSizes, false, nil},
	{"R15", false, REG_R15, DATA_REG, 64, QuadSizes, false, nil},
	{"X0", false, REG_X0, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X1", false, REG_X1, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X2", false, REG_X2, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X3", false, REG_X3, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X4", false, REG_X4, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X5", false, REG_X5, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X6", false, REG_X6, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X7", false, REG_X7, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X8", false, REG_X8, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X9", false, REG_X9, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X10", false, REG_X10, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X11", false, REG_X11, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X12", false, REG_X12, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X13", false, REG_X13, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X14", false, REG_X14, XMM_REG, 128, XmmDataSizes, false, nil},
	{"X15", false, REG_X15, XMM_REG, 128, XmmDataSizes, false, nil},
}

var excludedRegisters = []register{
	// used as implicit operands in arithmetic instructions
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, false, nil},
	{"BL", false, REG_AL, DATA_REG, 32, LongSizes, false, nil},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, false, nil},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, false, nil},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, false, nil},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, false, nil},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, false, nil},

	// stack pointer pseudo register
	{"SP", false, REG_SP, SpReg, 64, QuadSize, false, nil},
	// frame pointer pseudo register
	{"FP", false, REG_FP, FpReg, 64, QuadSize, false, nil},
}

func getRegister(reg Reg) *register {
	for _, r := range registers {
		if r.regconst == reg {
			return &r
		}
	}
	return nil
}
