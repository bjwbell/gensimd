package sse2

type M128i [16]byte
type M128 [4]float32
type M128d [2]float64

type SSE2Instr int

const (
	INVALID SSE2Instr = iota
	// Integer
	LoadSi128
	LoaduSi128
	StoreuSi128
	AddEpi64
	SubEpi64
	MulEpu32
	ShufflehiEpi16
	ShuffleloEpi16
	ShuffleEpi32
	SlliSi128
	SrliSi128
	UnpackhiEpi64
	UnpackloEpi64
	// Float
	AddPd
	AddSd
	AndnotPd
	CmpeqPd
	CmpeqSd
)
