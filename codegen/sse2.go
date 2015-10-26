package codegen

import (
	"fmt"
	"go/token"

	"github.com/bjwbell/gensimd/codegen/sse2"
)

var sse2ToGoAsm = map[sse2.SSE2Instr]Instr{
	// integer
	sse2.LoadSi128:      MOVO,
	sse2.LoaduSi128:     MOVOU,
	sse2.StoreuSi128:    MOVOU,
	sse2.AddEpi64:       PADDQ,
	sse2.SubEpi64:       PSUBQ,
	sse2.MulEpu32:       PMULULQ,
	sse2.ShufflehiEpi16: PSHUFHW,
	sse2.ShuffleloEpi16: PSHUFLW,
	sse2.ShuffleEpi32:   PSHUFD,
	sse2.SlliSi128:      PSLLO, // not useable shift count must be an imm8 not register or mem location
	sse2.SrliSi128:      PSRLO, // not useable shift count must be an imm8 not register or mem location
	sse2.UnpackhiEpi64:  PUNPCKHLQ,
	sse2.UnpackloEpi64:  PUNPCKLLQ,

	// float
	sse2.AddPd:    ADDPD,
	sse2.AddSd:    ADDSD,
	sse2.AndnotPd: ANDNPD,
	sse2.CmpeqPd:  CMPPD,
	sse2.CmpeqSd:  CMPPS,
}

func getSSE2(name string) (sse2.SSE2Instr, bool) {
	for sse2instr, _ := range sse2ToGoAsm {
		if fmt.Sprintf("%v", sse2instr) == name {
			return sse2instr, true
		}
	}
	return sse2.INVALID, false
}

func sse2Op(f *Function, instr sse2.SSE2Instr, x, y, result *identifier, pos token.Pos) (string, *Error) {
	asm := ""
	goInstr := sse2ToGoAsm[instr]
	a, regx, err := f.LoadSSE2(x, pos)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSSE2(y, pos)
	if err != nil {
		return "", err
	}
	asm += b
	asm += instrRegReg(f, goInstr, regx, regy)
	c, err := f.StoreSSE2(regy, result, pos)
	if err != nil {
		return "", err
	}
	asm += c

	f.freeReg(*regx)
	f.freeReg(*regy)

	return asm, err
}
