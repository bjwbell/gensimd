package codegen

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ssa"

	"github.com/bjwbell/gensimd/simd"
)

var sse2intructions = []Intrinsic{
	{
		Name:            "AddEpi64",
		InstructionName: "PADDQ",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "SubEpi64",
		InstructionName: "PSUBQ",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "MulEpu32",
		InstructionName: "PMULULQ",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "ShufflehiEpi16",
		InstructionName: "PSHUFHW",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "ShuffleloEpi16",
		InstructionName: "PSHUFLW",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "ShuffleEpi32",
		InstructionName: "PSHUFD",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "SlliSi128",
		InstructionName: "PSLLO",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: Imm8, Reg: INVALID_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "SrliSi128",
		InstructionName: "PSRLO",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "UnpackhiEpi64",
		InstructionName: "PUNPCKHLQ",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: Imm8, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "UnpackloEpi64",
		InstructionName: "PUNPCKLLQ",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: Imm8, Reg: INVALID_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128i{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "AddPd",
		InstructionName: "ADDPD",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128d{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128d{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "AddSd",
		InstructionName: "ADDSD",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "AndnotPd",
		InstructionName: "ANDNPD",
		VarOps: []VarOp{
			{
				Var: Variable{reflect.TypeOf(simd.M128d{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128d{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "CmpeqPd",
		InstructionName: "CMPPD",
		VarOps: []VarOp{
			{
				Var:   Variable{},
				Const: 0,
				Op:    Operand{Flags: Imm8, Reg: INVALID_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 1,
	},
	{
		Name:            "CmpeqSd",
		InstructionName: "CMPPD",
		VarOps: []VarOp{
			{
				Var:   Variable{},
				Const: 0,
				Op:    Operand{Flags: Imm8, Reg: INVALID_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In, Reg: XMM_REG, NamedReg: REG_INVALID}},
			{
				Var: Variable{reflect.TypeOf(simd.M128{})},
				Op:  Operand{Flags: In | Out, Reg: XMM_REG, NamedReg: REG_INVALID},
			},
		},
		ResultIdxOp: 2,
	},
}

func getSSE2(name string) (Intrinsic, bool) {
	for _, sse2instr := range sse2intructions {
		if sse2instr.Name == name {
			return sse2instr, true
		}
	}
	return Intrinsic{}, false
}

func sse2Intrinsic(f *Function, loc ssa.Instruction, call *ssa.Call, intrinsic Intrinsic, args []ssa.Value) string {
	var asm string
	idents := []*identifier{}
	argRegs := []*register{}
	result := f.Ident(call)
	asm += fmt.Sprintf("// BEGIN INTRINSIC %v\n", intrinsic.Name)
	asm += fmt.Sprintf("// BEGIN LOAD ARGS %v\n", intrinsic.Name)
	// load the arguments into registers
	for _, arg := range args {
		ident := f.Ident(arg)
		a, reg, err := f.LoadIdent(loc, ident, 0, ident.size())
		if err != nil {
			ice(fmt.Sprintf("%v", err.Err))
		}
		reg.inUse = true
		asm += a
		idents = append(idents, ident)
		argRegs = append(argRegs, reg)
	}
	asm += fmt.Sprintf("// END LOAD ARGS %v\n", intrinsic.Name)

	// construct the assembly instruction
	var regResult *register
	argIdx := 0
	assembly := fmt.Sprintf("%-9v    ", intrinsic.InstructionName)
	for i, varop := range intrinsic.VarOps {
		flags := varop.Op.Flags
		if flags&Implicit == 1 {
			continue
		}

		if flags&Const != 0 {
			if flags&Imm8 != 0 || flags&Imm32 != 0 {
				asm += "$" + strconv.Itoa(varop.Const)
			} else if flags&ImmF32 != 0 || flags&ImmF64 != 0 {
				asm += ice("unimplemented case")
			} else {
				ice("unexpected case")
			}
			continue
		}
		if flags&In == 0 && flags&Out == 0 {
			continue
		}
		ident := idents[argIdx]
		reg := argRegs[argIdx]
		if reflectType(args[argIdx].Type()).String() != varop.Var.VarType.String() {
			got := reflectType(args[argIdx].Type()).String()
			expected := varop.Var.VarType.String()
			msg := fmt.Sprintf("wrong type for argument, Got: %v, Expected: %v", got, expected)
			ice(msg)
		}
		if flags&In != 0 {
			if flags&Imm8 != 0 || flags&Imm32 != 0 {
				if ident.cnst == nil {
					panic("Provided argument must be a constant")
				} else {
					assembly += "$" + strconv.Itoa(int(ident.cnst.Uint64()))
				}
			} else if flags&ImmF32 != 0 || flags&ImmF64 != 0 {
				ice("unimplemented case")
			}
			assembly += reg.name + ", "

		}
		if flags&Out != 0 {
			if i == intrinsic.ResultIdxOp {
				regResult = reg
			} else if flags&Out != 0 {
				reg.dirty = true
			}
		}
		argIdx++

	}
	assembly = strings.TrimSuffix(assembly, ", ") + "\n"
	asm += assembly
	if regResult != nil {
		c, err := f.StoreSSE2(loc, regResult, result)
		if err != nil {
			ice(fmt.Sprint(err.Err))
		}
		asm += c
	} else {
		ice("no result from intrinsic")
	}
	for _, reg := range argRegs {
		f.freeReg(reg)
	}
	asm += fmt.Sprintf("// END INTRINSIC %v\n", intrinsic.Name)
	return asm
}
