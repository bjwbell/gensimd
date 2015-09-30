package simd

import "fmt"

type Execer interface {
	Exec() error
}

type Asmer interface {
	Asm() string
}

type Instruction interface {
	InstructionType()
	Execer
	Asmer
}

type Block []Instruction

type ForLoop struct {
	Iterations int
	Body       Block
}

type Int4Var *[4]int

type Int4Add struct {
	Result Int4Var
	A      Int4Var
	B      Int4Var
}

func (loop ForLoop) String() string {
	return fmt.Sprintf("ForLoop{Iterations: %v,Body:%v}", loop.Iterations, loop.Body.String())
}
func (loop ForLoop) Exec() error {
	// TODO
	return nil
}
func (b Block) String() string {
	if b == nil {
		return "nil"
	}
	str := "{"
	for _, instruction := range b {
		str += instruction.Asm()
	}
	str += "}"
	return str
}
func (i4 Int4Add) Exec() error {
	// TODO
	return nil
}
func (i4 Int4Add) Asm() string {
	return fmt.Sprintf("Int4Add(Res(%v), A(%v), B(%v)", i4.Result, i4.A, i4.B)
}
func (i4 Int4Add) InstructionType() {
}
