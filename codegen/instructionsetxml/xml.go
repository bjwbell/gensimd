package instructionsetxml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
)

type Instructionset struct {
	Name         string        `xml:"name,attr"`
	Instructions []Instruction `xml:"Instruction"`
}

type Instruction struct {
	Name    string            `xml:"name,attr"`
	Summary string            `xml:"summary,attr"`
	Forms   []InstructionForm `xml:"InstructionForm"`
}

type InstructionForm struct {
	GoName           string            `xml:"go-name,attr"`
	Operands         []Operand         `xml:"Operand"`
	ImplicitOperands []ImplicitOperand `xml:"ImplicitOperand"`
}

type Operand struct {
	Type   string `xml:"type,attr"`
	Input  bool   `xml:"input,attr"`
	Output bool   `xml:"output,attr"`
}

type ImplicitOperand struct {
	Id     string `xml:"id,attr"`
	Input  bool   `xml:"input,attr"`
	Output bool   `xml:"output,attr"`
}

func LoadFromFile(filename string) (*Instructionset, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot read xml file")
		return nil, errors.New("Cannot read xml file")
	}

	set := Instructionset{}
	if err := xml.Unmarshal(file, &set); err != nil {
		fmt.Println("Cannot unmarshal instruction set xml")
		return nil, err
	}
	return &set, nil
}
