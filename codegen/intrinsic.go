package codegen

import "reflect"

type Intrinsic struct {
	Name            string
	InstructionName string
	VarOps          []VarOp
	ResultIdxOp     int
}

type VarOp struct {
	Var   Variable
	Const int
	Op    Operand
}

type Variable struct {
	VarType reflect.Type
}

type OpFlags int

const (
	_ OpFlags = 1 << iota
	In
	Out
	Implicit
	Const
	Imm8
	Imm32
	ImmF32
	ImmF64
)

type Operand struct {
	Flags    OpFlags
	Reg      RegType
	NamedReg Reg
}
