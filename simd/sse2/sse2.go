package sse2

import "github.com/bjwbell/gensimd/simd"

func LoadSi128(x, y simd.M128i) simd.M128i      { return simd.M128i{} }
func LoaduSi128(x, y simd.M128i) simd.M128i     { return simd.M128i{} }
func AddEpi64(x, y simd.M128i) simd.M128i       { return simd.M128i{} }
func SubEpi64(x, y simd.M128i) simd.M128i       { return simd.M128i{} }
func MulEpu32(x, y simd.M128i) simd.M128i       { return simd.M128i{} }
func ShufflehiEpi16(x, y simd.M128i) simd.M128i { return simd.M128i{} }
func ShuffleloEpi16(x, y simd.M128i) simd.M128i { return simd.M128i{} }
func ShuffleEpi32(x, y simd.M128i) simd.M128i   { return simd.M128i{} }
func SlliSi128(x, y simd.M128i) simd.M128i      { return simd.M128i{} }
func SrliSi128(x, y simd.M128i) simd.M128i      { return simd.M128i{} }
func UnpackhiEpi64(x, y simd.M128i) simd.M128i  { return simd.M128i{} }
func UnpackloEpi64(x, y simd.M128i) simd.M128i  { return simd.M128i{} }
func AddPd(x, y simd.M128d) simd.M128d          { return simd.M128d{} }
func AddSd(x, y simd.M128d) simd.M128d          { return simd.M128d{} }
func AndnotPd(x, y simd.M128d) simd.M128d       { return simd.M128d{} }
func CmpeqPd(x, y simd.M128d) simd.M128d        { return simd.M128d{} }
func CmpeqSd(x, y simd.M128d) simd.M128d        { return simd.M128d{} }
