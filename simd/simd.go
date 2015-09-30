package simd

import "fmt"

type Instruction interface {
	Asm() string
	Exec() error
}

type Int4Var *[4]int

type Int4Add struct {
	Result Int4Var
	A      Int4Var
	B      Int4Var
}

func (i4 Int4Add) Asm() string {
	return fmt.Sprintf("Int4Add(Res(%v), A(%v), B(%v)", i4.Result, i4.A, i4.B)
}

func (i4 Int4Add) Exec() error {
	// TODO
	return nil
}

type Block []Instruction
type ForLoop struct {
	Iterations int
	Body       Block
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
func (loop ForLoop) String() string {
	return fmt.Sprintf("ForLoop{Iterations: %v,Body:%v}", loop.Iterations, loop.Body.String())
}
