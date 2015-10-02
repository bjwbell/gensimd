package simd

type register struct {
	name string
	typ  registerType
	size int
}

type registerType int

const IntReg = registerType(0)
const IntRegSize = 8
const FloatReg = registerType(1)
const FloatRegSize = 8

var regnames = []string{
	"AX",
	"CX",
	"DX",
	"BX",
	"SI",
	"DI",
	"R8",
	"R9",
	"R10",
	"R11",
	"R12",
	"R13",
	"R14",
	"R15",
	"X0",
	"X1",
	"X2",
	"X3",
	"X4",
	"X5",
	"X6",
	"X7",
	"X8",
	"X9",
	"X10",
	"X11",
	"X12",
	"X13",
	"X14",
	"X15",
}
