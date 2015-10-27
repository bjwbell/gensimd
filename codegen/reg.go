package codegen

import "fmt"

type register struct {
	// register name e.g. ax, eax, rax, r15,...
	name  string
	inUse bool

	regconst Reg

	// type of register, for example xmm register, normal integer register, mmx register, etc.
	typ RegType
	// width of register in bits, e.g. ax is 64 on amd64.
	width uint
	// allowed data sizes in bytes used in reg
	datasizes []uint

	aliases []*identifier
}

func (r *register) size() uint {
	return r.width / 8
}

func (r *register) modified() string {
	inUse := r.inUse
	a := r.spill()
	r.inUse = inUse
	return a
}

func (r *register) isAlias(ident *identifier) bool {
	for _, alias := range r.aliases {
		if alias == ident {
			return true
		}
		if alias.name == ident.name {
			ice("aliases with same name should have equal memory locations")
		}
	}
	return false
}

func (r *register) addAlias(alias *identifier) {
	if !r.isAlias(alias) {
		if r.name == "R15" && "github.com/bjwbell/gensimd/simd.I32x4" == alias.typ.String() {
			panic("ddsdssd")
		}
		r.aliases = append(r.aliases, alias)
	}
	if !r.isAlias(alias) {
		ice("!r.isAlias")
	}
}

func (r *register) removeAlias(alias *identifier) bool {
	var aliases []*identifier
	removed := false
	for i := range r.aliases {
		if r.aliases[i] == alias {
			removed = true
			continue
		} else {
			if alias.name == r.aliases[i].name {
				ice("aliases with same name should have equal memory locations")
			}
			aliases = append(aliases, r.aliases[i])
		}
	}
	r.aliases = aliases
	return removed
}

func (r *register) spill() string {
	asm := ""
	for _, alias := range r.aliases {
		f := alias.f
		if a, err := f.StoreReg(r, alias, 0, alias.size()); err != nil {
			panic(ice(fmt.Sprintf("msg: %v", err.Err)))
		} else {
			asm += a
		}
	}
	r.aliases = nil
	r.inUse = false
	return asm
}

func (r *register) spillAlias(alias *identifier) (string, *Error) {
	if !r.isAlias(alias) {
		ice("can't spill register alias")
	}
	r.removeAlias(alias)
	f := alias.f
	return f.StoreReg(r, alias, 0, alias.size())
}

type RegType int

const (
	// integer register
	DATA_REG = RegType(iota)
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
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, nil},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, nil},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, nil},
	{"BL", false, REG_BL, DATA_REG, 32, LongSizes, nil},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, nil},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, nil},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, nil},
	{"SI", false, REG_SI, DATA_REG, 64, QuadSize, nil},
	{"DI", false, REG_DI, DATA_REG, 64, QuadSize, nil},
	{"BX", false, REG_BX, DATA_REG, 64, QuadSize, nil},
	{"BP", false, REG_BP, DATA_REG, 64, QuadSize, nil},
	{"SP", false, REG_SP, SpReg, 64, QuadSize, nil},
	{"FP", false, REG_FP, FpReg, 64, QuadSize, nil},
	{"R8", false, REG_R8, DATA_REG, 64, QuadSizes, nil},
	{"R9", false, REG_R9, DATA_REG, 64, QuadSizes, nil},
	{"R10", false, REG_R10, DATA_REG, 64, QuadSizes, nil},
	{"R11", false, REG_R11, DATA_REG, 64, QuadSizes, nil},
	{"R12", false, REG_R12, DATA_REG, 64, QuadSizes, nil},
	{"R13", false, REG_R13, DATA_REG, 64, QuadSizes, nil},
	{"R14", false, REG_R14, DATA_REG, 64, QuadSizes, nil},
	{"R15", false, REG_R15, DATA_REG, 64, QuadSizes, nil},
	{"X0", false, REG_X0, XMM_REG, 128, XmmDataSizes, nil},
	{"X1", false, REG_X1, XMM_REG, 128, XmmDataSizes, nil},
	{"X2", false, REG_X2, XMM_REG, 128, XmmDataSizes, nil},
	{"X3", false, REG_X3, XMM_REG, 128, XmmDataSizes, nil},
	{"X4", false, REG_X4, XMM_REG, 128, XmmDataSizes, nil},
	{"X5", false, REG_X5, XMM_REG, 128, XmmDataSizes, nil},
	{"X6", false, REG_X6, XMM_REG, 128, XmmDataSizes, nil},
	{"X7", false, REG_X7, XMM_REG, 128, XmmDataSizes, nil},
	{"X8", false, REG_X8, XMM_REG, 128, XmmDataSizes, nil},
	{"X9", false, REG_X9, XMM_REG, 128, XmmDataSizes, nil},
	{"X10", false, REG_X10, XMM_REG, 128, XmmDataSizes, nil},
	{"X11", false, REG_X11, XMM_REG, 128, XmmDataSizes, nil},
	{"X12", false, REG_X12, XMM_REG, 128, XmmDataSizes, nil},
	{"X13", false, REG_X13, XMM_REG, 128, XmmDataSizes, nil},
	{"X14", false, REG_X14, XMM_REG, 128, XmmDataSizes, nil},
	{"X15", false, REG_X15, XMM_REG, 128, XmmDataSizes, nil},
}

var excludedRegisters = []register{
	// used as implicit operands in arithmetic instructions
	{"AL", false, REG_AL, DATA_REG, 32, LongSizes, nil},
	{"BL", false, REG_AL, DATA_REG, 32, LongSizes, nil},
	{"CL", false, REG_CL, DATA_REG, 32, LongSizes, nil},
	{"DL", false, REG_DL, DATA_REG, 32, LongSizes, nil},
	{"AX", false, REG_AX, DATA_REG, 64, QuadSizes, nil},
	{"CX", false, REG_CX, DATA_REG, 64, QuadSizes, nil},
	{"DX", false, REG_DX, DATA_REG, 64, QuadSizes, nil},

	// stack pointer pseudo register
	{"SP", false, REG_SP, SpReg, 64, QuadSize, nil},
	// frame pointer pseudo register
	{"FP", false, REG_FP, FpReg, 64, QuadSize, nil},
}

func getRegister(reg Reg) *register {
	for _, r := range registers {
		if r.regconst == reg {
			return &r
		}
	}
	return nil
}
