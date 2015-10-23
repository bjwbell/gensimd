package sse2

type M128i [16]byte
type M128 [4]float32
type M128d [2]float64

type SSE2Instr int

const (
	INVALID SSE2Instr = iota
	// Integer
	MOVDQA
	MOVDQU
	PADDQ
	PSUBQ
	PMULUDQ
	PSHUFHW
	PSHUFLW
	PSHUFD
	PSLLDQ
	PSRLDQ
	PUNPCKHQDQ
	PUNPCKLQDQ
	// Float
	ADDPD
	ADDSD
	ANDNPD
	CMPPD
	CMPSD
)
