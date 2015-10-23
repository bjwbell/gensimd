package sse2

import "github.com/bjwbell/gensimd/simd"

func MOVDQA(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func MOVDQU(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func PADDQ(x, y simd.M128i) simd.M128i      { return simd.M128i{} }
func PSUBQ(x, y simd.M128i) simd.M128i      { return simd.M128i{} }
func PMULUDQ(x, y simd.M128i) simd.M128i    { return simd.M128i{} }
func PSHUFHW(x, y simd.M128i) simd.M128i    { return simd.M128i{} }
func PSHUFLW(x, y simd.M128i) simd.M128i    { return simd.M128i{} }
func PSHUFD(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func PSLLDQ(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func PSRLDQ(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func PUNPCKHQDQ(x, y simd.M128i) simd.M128i { return simd.M128i{} }
func PUNPCKLQDQ(x, y simd.M128i) simd.M128i { return simd.M128i{} }
func ADDPD(x, y simd.M128d) simd.M128d      { return simd.M128d{} }
func ADDSD(x, y simd.M128d) simd.M128d      { return simd.M128d{} }
func ANDNPD(x, y simd.M128d) simd.M128d     { return simd.M128d{} }
func CMPPD(x, y simd.M128d) simd.M128d      { return simd.M128d{} }
func CMPSD(x, y simd.M128d) simd.M128d      { return simd.M128d{} }
