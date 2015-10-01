package simd

type Execer interface {
	Exec() error
}

type Asmer interface {
	Asm() string
}

type Instructioner interface {
	TypeInstruction()
	Execer
	Asmer
}

type RetInst interface {
	TypeRetInst()
	Execer
	Asmer
}

type Block []Instructioner

type Looper interface {
	TypeLooper()
	Execer
	Asmer
}

type ForLoop struct {
	Looper
	Start      Int
	Iterations Int
	StepBy     Int
	Body       Block
}

type Func struct {
	Init   Block
	Loop   ForLoop
	Finish Block
	Ret    RetInst
}

type Int int
type IntVar *int
type Int4 [4]int
type Int4Var *[4]int

type Int4Add struct {
	Instructioner
	Result Int4Var
	A      Int4Var
	B      Int4Var
}

type RetSuccess struct {
	RetInst
}
