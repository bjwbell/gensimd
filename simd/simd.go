package simd

import "fmt"

func (r *RetSuccess) Exec() error {
	return nil
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

func (f *Func) Exec() error {
	// TODO
	return nil
}
