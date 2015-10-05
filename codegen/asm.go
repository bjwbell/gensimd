package codegen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bjwbell/gensimd/codegen/instructionsetxml"
)

func InstAsm(instructionsetXmlPath string, name InstName, ops []*Operand) (string, error) {
	if set, err := instructionsetxml.LoadInstructionSet(instructionsetXmlPath); err != nil {
		fmt.Println("ERROR: couldn't get instruction set")
		fmt.Println("err:", err)
		return "", err
	} else {
		return InstructionSetAsm(set, name, ops)
	}
}

func InstructionSetAsm(set *instructionsetxml.InstructionSet, name InstName, ops []*Operand) (string, error) {
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
		fmt.Println("InstName:", name.String())
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
	for i, xOp := range xmlOps {
		op := ops[i]
		opType := strings.ToLower(op.Type.String())
		xType := strings.ToLower(xOp.Type)
		if opType != xType {
			return false
		}
		if op.Input == xOp.Input && op.Output == xOp.Output {
			return false
		}
	}
	return true
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
