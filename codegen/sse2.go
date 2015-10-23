package codegen

import (
	"fmt"

	"github.com/bjwbell/gensimd/codegen/sse2"
)

var sse2ToGoAsm = map[sse2.SSE2Instr]Instr{
	// integer
	sse2.MOVDQA:     MOVO,
	sse2.MOVDQU:     MOVOU,
	sse2.PADDQ:      PADDQ,
	sse2.PSUBQ:      PSUBQ,
	sse2.PMULUDQ:    PMULULQ,
	sse2.PSHUFHW:    PSHUFHW,
	sse2.PSHUFLW:    PSHUFLW,
	sse2.PSHUFD:     PSHUFD,
	sse2.PSLLDQ:     PSLLO, // not useable shift count must be an imm8 not register or mem location
	sse2.PSRLDQ:     PSRLO, // not useable shift count must be an imm8 not register or mem location
	sse2.PUNPCKHQDQ: PUNPCKHLQ,
	sse2.PUNPCKLQDQ: PUNPCKLLQ,

	// float
	sse2.ADDPD:  ADDPD,
	sse2.ADDSD:  ADDSD,
	sse2.ANDNPD: ANDNPD,
	sse2.CMPPD:  CMPPD,
	sse2.CMPSD:  CMPPS,
}

func getSSE2(name string) (sse2.SSE2Instr, bool) {
	for sse2instr, _ := range sse2ToGoAsm {
		if fmt.Sprintf("%v", sse2instr) == name {
			return sse2instr, true
		}
	}
	return sse2.INVALID, false
}

func sse2Op(f *Function, instr sse2.SSE2Instr, x, y, result *identifier) (string, *Error) {
	asm := ""
	goInstr := sse2ToGoAsm[instr]
	a, regx, err := f.LoadSSE2(x)
	if err != nil {
		return "", err
	}
	asm += a
	b, regy, err := f.LoadSSE2(y)
	if err != nil {
		return "", err
	}
	asm += b
	asm += instrRegReg(f, goInstr, regx, regy)
	c, err := f.StoreSSE2(regy, result)
	if err != nil {
		return "", err
	}
	asm += c

	f.freeReg(*regx)
	f.freeReg(*regy)

	return asm, err
}
