package instructionsetxml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

type InstructionSet struct {
	Name         string        `xml:"name,attr"`
	Instructions []Instruction `xml:"Instruction"`
}

type Instruction struct {
	Name    string            `xml:"name,attr"`
	Summary string            `xml:"summary,attr"`
	Forms   []InstructionForm `xml:"InstructionForm"`
}

type InstructionForm struct {
	GoName   string    `xml:"go-name,attr"`
	Operands []Operand `xml:"Operand"`
}

type Operand struct {
	Type   string `xml:"type,attr"`
	Input  bool   `xml:"input,attr"`
	Output bool   `xml:"output,attr"`
}

func LoadInstructionSet(filename string) (*InstructionSet, error) {
	file, err := ioutil.ReadFile(filename) //"codegen/instructionsxml/x86_64.xml")
	if err != nil {
		fmt.Println("Cannot read xml file")
		return nil, errors.New("Cannot read xml file")
	}

	set := InstructionSet{}
	if err := xml.Unmarshal(file, &set); err != nil {
		fmt.Println("Cannot unmarshal instruction set xml")
		return nil, err
	}
	return &set, nil
}
