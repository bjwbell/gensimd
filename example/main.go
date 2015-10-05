package main

import "fmt"

type InstructionSet struct {
	Name         string        `xml:"name,attr"`
	Instructions []Instruction `xml:"Instruction"`
}

type Instruction struct {
	Name    string            `xml:"name,attr"`
	Summary string            `xml:"summary, attr"`
	Forms   []InstructionForm `xml:"InstructionForm"`
}

type InstructionForm struct {
	GoName   string    `xml:"go-name,attr"`
	Operands []Operand `xml:"Operand"`
}

type Operand struct {
	Type string `xml:"type,attr"`
}

func main() {
	v4 := [4]int{}
	//_ = simd_loop1(&v4)
	fmt.Println("v4:", v4)
}
