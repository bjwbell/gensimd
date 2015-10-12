package codegen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bjwbell/gensimd/codegen/instructionsetxml"
)

func FindInstr(set *instructionsetxml.Instructionset, typ InstrType, ops []*Operand) (Instr, error) {
	var instr *instructionsetxml.Instruction
	var form *instructionsetxml.InstructionForm

	for _, inst := range set.Instructions {
		if typ.String() == inst.Name {
			instr = &inst
			break
		}
	}
	if instr == nil {
		return 0, fmt.Errorf("Couldn't find match in instructionset for instr:%v", typ.String())
	}
	for _, fm := range instr.Forms {
		if OperandsMatch(ops, fm.Operands) {
			form = &fm
			break
		}

	}
	if form == nil {
		fmt.Println("Couldn't find instr form match for, InstrName:", typ.String())
		for _, op := range ops {
			fmt.Println("op:", op)
		}
		return 0, fmt.Errorf("No matched instr form for instr:%v", instr)
	}
	if instrName, err := GetInstrFromStr(form.GoName); err != nil {
		return instrName, fmt.Errorf("No InstrName for form.GoName:%v", form.GoName)
	} else {
		return instrName, nil
	}
}

func InstrAsm(set *instructionsetxml.Instructionset, typ InstrType, ops []*Operand) (string, error) {
	if instrName, err := FindInstr(set, typ, ops); err != nil {
		return "", err
	} else {
		if a, err := InstructionSetAsm(set, instrName, ops); err != nil {
			return "", err
		} else {
			return a, nil
		}
	}
}

func InstructionSetAsm(set *instructionsetxml.Instructionset, name Instr, ops []*Operand) (string, error) {
	var form *instructionsetxml.InstructionForm
	for _, inst := range set.Instructions {
		for _, fm := range inst.Forms {

			if strings.ToLower(fm.GoName) !=
				strings.ToLower(name.String()) {
				continue
			}
			if !OperandsMatch(ops, fm.Operands) {
				continue
			}
			form = &fm
			break
		}
		if form != nil {
			break
		}
	}
	if form == nil {
		fmt.Println("InstrName:", name.String())
		for _, op := range ops {
			fmt.Println("op:", op)
		}
		return "", errors.New("No matching instruction form")
	}
	return InstructionFormAsm(form, ops), nil
}

func OperandsMatch(ops []*Operand, xmlOps []instructionsetxml.Operand) bool {
	if len(ops) != len(xmlOps) {
		return false
	}
	matched := make(map[int]int)
	for i, xOp := range xmlOps {
		for j, op := range ops {
			if _, ok := matched[j]; ok {
				continue
			}
			opType := strings.ToLower(op.Type.String())
			xType := strings.ToLower(xOp.Type)
			if opType != xType {
				continue
			}
			if op.Input != xOp.Input || op.Output != xOp.Output {
				continue
			}
			matched[j] = i
		}
	}
	return len(matched) == len(ops)
}

func InstructionFormAsm(form *instructionsetxml.InstructionForm, ops []*Operand) string {
	asm := form.GoName + "    "
	opsAsm := []string{}
	for _, op := range ops {
		opAsm := op.Value()
		opsAsm = append(opsAsm, opAsm)
	}
	asm += strings.Join(opsAsm, ", ")
	return asm
}
