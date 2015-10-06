package codegen

import "fmt"

type InstrType struct {
	typ  int
	name string
}

const (
	TADC = iota
	TADCX
	TADD
	TADDPD
	TADDPS
	TADDSD
	TADDSS
	TADDSUBPD
	TADDSUBPS
	TADOX
	TAESDEC
	TAESDECLAST
	TAESENC
	TAESENCLAST
	TAESIMC
	TAESKEYGENASSIST
	TAND
	TANDN
	TANDNPD
	TANDNPS
	TANDPD
	TANDPS
	TBEXTR
	TBLCFILL
	TBLCI
	TBLCIC
	TBLCMSK
	TBLCS
	TBLENDPD
	TBLENDPS
	TBLENDVPD
	TBLENDVPS
	TBLSFILL
	TBLSI
	TBLSIC
	TBLSMSK
	TBLSR
	TBSF
	TBSR
	TBSWAP
	TBT
	TBTC
	TBTR
	TBTS
	TBZHI
	TCALL
	TCBW
	TCDQ
	TCDQE
	TCLC
	TCLD
	TCMC
	TCMOVA
	TCMOVAE
	TCMOVB
	TCMOVBE
	TCMOVC
	TCMOVE
	TCMOVG
	TCMOVGE
	TCMOVL
	TCMOVLE
	TCMOVNA
	TCMOVNAE
	TCMOVNB
	TCMOVNBE
	TCMOVNC
	TCMOVNE
	TCMOVNG
	TCMOVNGE
	TCMOVNL
	TCMOVNLE
	TCMOVNO
	TCMOVNP
	TCMOVNS
	TCMOVNZ
	TCMOVO
	TCMOVP
	TCMOVPE
	TCMOVPO
	TCMOVS
	TCMOVZ
	TCMP
	TCMPPD
	TCMPPS
	TCMPSD
	TCMPSS
	TCMPXCHG
	TCMPXCHG16B
	TCMPXCHG8B
	TCOMISD
	TCOMISS
	TCPUID
	TCQO
	TCRC32
	TCVTDQ2PD
	TCVTDQ2PS
	TCVTPD2DQ
	TCVTPD2PI
	TCVTPD2PS
	TCVTPI2PD
	TCVTPI2PS
	TCVTPS2DQ
	TCVTPS2PD
	TCVTPS2PI
	TCVTSD2SI
	TCVTSD2SS
	TCVTSI2SD
	TCVTSI2SS
	TCVTSS2SD
	TCVTSS2SI
	TCVTTPD2DQ
	TCVTTPD2PI
	TCVTTPS2DQ
	TCVTTPS2PI
	TCVTTSD2SI
	TCVTTSS2SI
	TCWD
	TCWDE
	TDEC
	TDIV
	TDIVPD
	TDIVPS
	TDIVSD
	TDIVSS
	TDPPD
	TDPPS
	TEMMS
	TEXTRACTPS
	TEXTRQ
	TFEMMS
	THADDPD
	THADDPS
	THSUBPD
	THSUBPS
	TIDIV
	TIMUL
	TINC
	TINSERTPS
	TINSERTQ
	TINT
	TJA
	TJAE
	TJB
	TJBE
	TJC
	TJE
	TJECXZ
	TJG
	TJGE
	TJL
	TJLE
	TJMP
	TJNA
	TJNAE
	TJNB
	TJNBE
	TJNC
	TJNE
	TJNG
	TJNGE
	TJNL
	TJNLE
	TJNO
	TJNP
	TJNS
	TJNZ
	TJO
	TJP
	TJPE
	TJPO
	TJRCXZ
	TJS
	TJZ
	TKADDB
	TKADDD
	TKADDQ
	TKADDW
	TKANDB
	TKANDD
	TKANDNB
	TKANDND
	TKANDNQ
	TKANDNW
	TKANDQ
	TKANDW
	TKMOVB
	TKMOVD
	TKMOVQ
	TKMOVW
	TKNOTB
	TKNOTD
	TKNOTQ
	TKNOTW
	TKORB
	TKORD
	TKORQ
	TKORTESTB
	TKORTESTD
	TKORTESTQ
	TKORTESTW
	TKORW
	TKSHIFTLB
	TKSHIFTLD
	TKSHIFTLQ
	TKSHIFTLW
	TKSHIFTRB
	TKSHIFTRD
	TKSHIFTRQ
	TKSHIFTRW
	TKTESTB
	TKTESTD
	TKTESTQ
	TKTESTW
	TKUNPCKBW
	TKUNPCKDQ
	TKUNPCKWD
	TKXNORB
	TKXNORD
	TKXNORQ
	TKXNORW
	TKXORB
	TKXORD
	TKXORQ
	TKXORW
	TLDDQU
	TLDMXCSR
	TLEA
	TLFENCE
	TLZCNT
	TMASKMOVDQU
	TMASKMOVQ
	TMAXPD
	TMAXPS
	TMAXSD
	TMAXSS
	TMFENCE
	TMINPD
	TMINPS
	TMINSD
	TMINSS
	TMOV
	TMOVAPD
	TMOVAPS
	TMOVBE
	TMOVD
	TMOVDDUP
	TMOVDQ2Q
	TMOVDQA
	TMOVDQU
	TMOVHLPS
	TMOVHPD
	TMOVHPS
	TMOVLHPS
	TMOVLPD
	TMOVLPS
	TMOVMSKPD
	TMOVMSKPS
	TMOVNTDQ
	TMOVNTDQA
	TMOVNTI
	TMOVNTPD
	TMOVNTPS
	TMOVNTQ
	TMOVNTSD
	TMOVNTSS
	TMOVQ
	TMOVQ2DQ
	TMOVSD
	TMOVSHDUP
	TMOVSLDUP
	TMOVSS
	TMOVSX
	TMOVSXD
	TMOVUPD
	TMOVUPS
	TMOVZX
	TMPSADBW
	TMUL
	TMULPD
	TMULPS
	TMULSD
	TMULSS
	TMULX
	TNEG
	TNOP
	TNOT
	TOR
	TORPD
	TORPS
	TPABSB
	TPABSD
	TPABSW
	TPACKSSDW
	TPACKSSWB
	TPACKUSDW
	TPACKUSWB
	TPADDB
	TPADDD
	TPADDQ
	TPADDSB
	TPADDSW
	TPADDUSB
	TPADDUSW
	TPADDW
	TPALIGNR
	TPAND
	TPANDN
	TPAUSE
	TPAVGB
	TPAVGUSB
	TPAVGW
	TPBLENDVB
	TPBLENDW
	TPCLMULQDQ
	TPCMPEQB
	TPCMPEQD
	TPCMPEQQ
	TPCMPEQW
	TPCMPESTRI
	TPCMPESTRM
	TPCMPGTB
	TPCMPGTD
	TPCMPGTQ
	TPCMPGTW
	TPCMPISTRI
	TPCMPISTRM
	TPDEP
	TPEXT
	TPEXTRB
	TPEXTRD
	TPEXTRQ
	TPEXTRW
	TPF2ID
	TPF2IW
	TPFACC
	TPFADD
	TPFCMPEQ
	TPFCMPGE
	TPFCMPGT
	TPFMAX
	TPFMIN
	TPFMUL
	TPFNACC
	TPFPNACC
	TPFRCP
	TPFRCPIT1
	TPFRCPIT2
	TPFRSQIT1
	TPFRSQRT
	TPFSUB
	TPFSUBR
	TPHADDD
	TPHADDSW
	TPHADDW
	TPHMINPOSUW
	TPHSUBD
	TPHSUBSW
	TPHSUBW
	TPI2FD
	TPI2FW
	TPINSRB
	TPINSRD
	TPINSRQ
	TPINSRW
	TPMADDUBSW
	TPMADDWD
	TPMAXSB
	TPMAXSD
	TPMAXSW
	TPMAXUB
	TPMAXUD
	TPMAXUW
	TPMINSB
	TPMINSD
	TPMINSW
	TPMINUB
	TPMINUD
	TPMINUW
	TPMOVMSKB
	TPMOVSXBD
	TPMOVSXBQ
	TPMOVSXBW
	TPMOVSXDQ
	TPMOVSXWD
	TPMOVSXWQ
	TPMOVZXBD
	TPMOVZXBQ
	TPMOVZXBW
	TPMOVZXDQ
	TPMOVZXWD
	TPMOVZXWQ
	TPMULDQ
	TPMULHRSW
	TPMULHRW
	TPMULHUW
	TPMULHW
	TPMULLD
	TPMULLW
	TPMULUDQ
	TPOP
	TPOPCNT
	TPOR
	TPREFETCHNTA
	TPREFETCHT0
	TPREFETCHT1
	TPREFETCHT2
	TPREFETCHW
	TPREFETCHWT1
	TPSADBW
	TPSHUFB
	TPSHUFD
	TPSHUFHW
	TPSHUFLW
	TPSHUFW
	TPSIGNB
	TPSIGND
	TPSIGNW
	TPSLLD
	TPSLLDQ
	TPSLLQ
	TPSLLW
	TPSRAD
	TPSRAW
	TPSRLD
	TPSRLDQ
	TPSRLQ
	TPSRLW
	TPSUBB
	TPSUBD
	TPSUBQ
	TPSUBSB
	TPSUBSW
	TPSUBUSB
	TPSUBUSW
	TPSUBW
	TPSWAPD
	TPTEST
	TPUNPCKHBW
	TPUNPCKHDQ
	TPUNPCKHQDQ
	TPUNPCKHWD
	TPUNPCKLBW
	TPUNPCKLDQ
	TPUNPCKLQDQ
	TPUNPCKLWD
	TPUSH
	TPXOR
	TRCL
	TRCPPS
	TRCPSS
	TRCR
	TRDRAND
	TRDSEED
	TRDTSC
	TRDTSCP
	TRET
	TROL
	TROR
	TRORX
	TROUNDPD
	TROUNDPS
	TROUNDSD
	TROUNDSS
	TRSQRTPS
	TRSQRTSS
	TSAL
	TSAR
	TSARX
	TSBB
	TSETA
	TSETAE
	TSETB
	TSETBE
	TSETC
	TSETE
	TSETG
	TSETGE
	TSETL
	TSETLE
	TSETNA
	TSETNAE
	TSETNB
	TSETNBE
	TSETNC
	TSETNE
	TSETNG
	TSETNGE
	TSETNL
	TSETNLE
	TSETNO
	TSETNP
	TSETNS
	TSETNZ
	TSETO
	TSETP
	TSETPE
	TSETPO
	TSETS
	TSETZ
	TSFENCE
	TSHA1MSG1
	TSHA1MSG2
	TSHA1NEXTE
	TSHA1RNDS4
	TSHA256MSG1
	TSHA256MSG2
	TSHA256RNDS2
	TSHL
	TSHLD
	TSHLX
	TSHR
	TSHRD
	TSHRX
	TSHUFPD
	TSHUFPS
	TSQRTPD
	TSQRTPS
	TSQRTSD
	TSQRTSS
	TSTC
	TSTD
	TSTMXCSR
	TSUB
	TSUBPD
	TSUBPS
	TSUBSD
	TSUBSS
	TT1MSKC
	TTEST
	TTZCNT
	TTZMSK
	TUCOMISD
	TUCOMISS
	TUD2
	TUNPCKHPD
	TUNPCKHPS
	TUNPCKLPD
	TUNPCKLPS
	TVADDPD
	TVADDPS
	TVADDSD
	TVADDSS
	TVADDSUBPD
	TVADDSUBPS
	TVAESDEC
	TVAESDECLAST
	TVAESENC
	TVAESENCLAST
	TVAESIMC
	TVAESKEYGENASSIST
	TVALIGND
	TVALIGNQ
	TVANDNPD
	TVANDNPS
	TVANDPD
	TVANDPS
	TVBLENDMPD
	TVBLENDMPS
	TVBLENDPD
	TVBLENDPS
	TVBLENDVPD
	TVBLENDVPS
	TVBROADCASTF128
	TVBROADCASTF32X2
	TVBROADCASTF32X4
	TVBROADCASTF32X8
	TVBROADCASTF64X2
	TVBROADCASTF64X4
	TVBROADCASTI128
	TVBROADCASTI32X2
	TVBROADCASTI32X4
	TVBROADCASTI32X8
	TVBROADCASTI64X2
	TVBROADCASTI64X4
	TVBROADCASTSD
	TVBROADCASTSS
	TVCMPPD
	TVCMPPS
	TVCMPSD
	TVCMPSS
	TVCOMISD
	TVCOMISS
	TVCOMPRESSPD
	TVCOMPRESSPS
	TVCVTDQ2PD
	TVCVTDQ2PS
	TVCVTPD2DQ
	TVCVTPD2PS
	TVCVTPD2QQ
	TVCVTPD2UDQ
	TVCVTPD2UQQ
	TVCVTPH2PS
	TVCVTPS2DQ
	TVCVTPS2PD
	TVCVTPS2PH
	TVCVTPS2QQ
	TVCVTPS2UDQ
	TVCVTPS2UQQ
	TVCVTQQ2PD
	TVCVTQQ2PS
	TVCVTSD2SI
	TVCVTSD2SS
	TVCVTSD2USI
	TVCVTSI2SD
	TVCVTSI2SS
	TVCVTSS2SD
	TVCVTSS2SI
	TVCVTSS2USI
	TVCVTTPD2DQ
	TVCVTTPD2QQ
	TVCVTTPD2UDQ
	TVCVTTPD2UQQ
	TVCVTTPS2DQ
	TVCVTTPS2QQ
	TVCVTTPS2UDQ
	TVCVTTPS2UQQ
	TVCVTTSD2SI
	TVCVTTSD2USI
	TVCVTTSS2SI
	TVCVTTSS2USI
	TVCVTUDQ2PD
	TVCVTUDQ2PS
	TVCVTUQQ2PD
	TVCVTUQQ2PS
	TVCVTUSI2SD
	TVCVTUSI2SS
	TVDBPSADBW
	TVDIVPD
	TVDIVPS
	TVDIVSD
	TVDIVSS
	TVDPPD
	TVDPPS
	TVEXP2PD
	TVEXP2PS
	TVEXPANDPD
	TVEXPANDPS
	TVEXTRACTF128
	TVEXTRACTF32X4
	TVEXTRACTF32X8
	TVEXTRACTF64X2
	TVEXTRACTF64X4
	TVEXTRACTI128
	TVEXTRACTI32X4
	TVEXTRACTI32X8
	TVEXTRACTI64X2
	TVEXTRACTI64X4
	TVEXTRACTPS
	TVFIXUPIMMPD
	TVFIXUPIMMPS
	TVFIXUPIMMSD
	TVFIXUPIMMSS
	TVFMADD132PD
	TVFMADD132PS
	TVFMADD132SD
	TVFMADD132SS
	TVFMADD213PD
	TVFMADD213PS
	TVFMADD213SD
	TVFMADD213SS
	TVFMADD231PD
	TVFMADD231PS
	TVFMADD231SD
	TVFMADD231SS
	TVFMADDPD
	TVFMADDPS
	TVFMADDSD
	TVFMADDSS
	TVFMADDSUB132PD
	TVFMADDSUB132PS
	TVFMADDSUB213PD
	TVFMADDSUB213PS
	TVFMADDSUB231PD
	TVFMADDSUB231PS
	TVFMADDSUBPD
	TVFMADDSUBPS
	TVFMSUB132PD
	TVFMSUB132PS
	TVFMSUB132SD
	TVFMSUB132SS
	TVFMSUB213PD
	TVFMSUB213PS
	TVFMSUB213SD
	TVFMSUB213SS
	TVFMSUB231PD
	TVFMSUB231PS
	TVFMSUB231SD
	TVFMSUB231SS
	TVFMSUBADD132PD
	TVFMSUBADD132PS
	TVFMSUBADD213PD
	TVFMSUBADD213PS
	TVFMSUBADD231PD
	TVFMSUBADD231PS
	TVFMSUBADDPD
	TVFMSUBADDPS
	TVFMSUBPD
	TVFMSUBPS
	TVFMSUBSD
	TVFMSUBSS
	TVFNMADD132PD
	TVFNMADD132PS
	TVFNMADD132SD
	TVFNMADD132SS
	TVFNMADD213PD
	TVFNMADD213PS
	TVFNMADD213SD
	TVFNMADD213SS
	TVFNMADD231PD
	TVFNMADD231PS
	TVFNMADD231SD
	TVFNMADD231SS
	TVFNMADDPD
	TVFNMADDPS
	TVFNMADDSD
	TVFNMADDSS
	TVFNMSUB132PD
	TVFNMSUB132PS
	TVFNMSUB132SD
	TVFNMSUB132SS
	TVFNMSUB213PD
	TVFNMSUB213PS
	TVFNMSUB213SD
	TVFNMSUB213SS
	TVFNMSUB231PD
	TVFNMSUB231PS
	TVFNMSUB231SD
	TVFNMSUB231SS
	TVFNMSUBPD
	TVFNMSUBPS
	TVFNMSUBSD
	TVFNMSUBSS
	TVFPCLASSPD
	TVFPCLASSPS
	TVFPCLASSSD
	TVFPCLASSSS
	TVFRCZPD
	TVFRCZPS
	TVFRCZSD
	TVFRCZSS
	TVGATHERDPD
	TVGATHERDPS
	TVGATHERPF0DPD
	TVGATHERPF0DPS
	TVGATHERPF0QPD
	TVGATHERPF0QPS
	TVGATHERPF1DPD
	TVGATHERPF1DPS
	TVGATHERPF1QPD
	TVGATHERPF1QPS
	TVGATHERQPD
	TVGATHERQPS
	TVGETEXPPD
	TVGETEXPPS
	TVGETEXPSD
	TVGETEXPSS
	TVGETMANTPD
	TVGETMANTPS
	TVGETMANTSD
	TVGETMANTSS
	TVHADDPD
	TVHADDPS
	TVHSUBPD
	TVHSUBPS
	TVINSERTF128
	TVINSERTF32X4
	TVINSERTF32X8
	TVINSERTF64X2
	TVINSERTF64X4
	TVINSERTI128
	TVINSERTI32X4
	TVINSERTI32X8
	TVINSERTI64X2
	TVINSERTI64X4
	TVINSERTPS
	TVLDDQU
	TVLDMXCSR
	TVMASKMOVDQU
	TVMASKMOVPD
	TVMASKMOVPS
	TVMAXPD
	TVMAXPS
	TVMAXSD
	TVMAXSS
	TVMINPD
	TVMINPS
	TVMINSD
	TVMINSS
	TVMOVAPD
	TVMOVAPS
	TVMOVD
	TVMOVDDUP
	TVMOVDQA
	TVMOVDQA32
	TVMOVDQA64
	TVMOVDQU
	TVMOVDQU16
	TVMOVDQU32
	TVMOVDQU64
	TVMOVDQU8
	TVMOVHLPS
	TVMOVHPD
	TVMOVHPS
	TVMOVLHPS
	TVMOVLPD
	TVMOVLPS
	TVMOVMSKPD
	TVMOVMSKPS
	TVMOVNTDQ
	TVMOVNTDQA
	TVMOVNTPD
	TVMOVNTPS
	TVMOVQ
	TVMOVSD
	TVMOVSHDUP
	TVMOVSLDUP
	TVMOVSS
	TVMOVUPD
	TVMOVUPS
	TVMPSADBW
	TVMULPD
	TVMULPS
	TVMULSD
	TVMULSS
	TVORPD
	TVORPS
	TVPABSB
	TVPABSD
	TVPABSQ
	TVPABSW
	TVPACKSSDW
	TVPACKSSWB
	TVPACKUSDW
	TVPACKUSWB
	TVPADDB
	TVPADDD
	TVPADDQ
	TVPADDSB
	TVPADDSW
	TVPADDUSB
	TVPADDUSW
	TVPADDW
	TVPALIGNR
	TVPAND
	TVPANDD
	TVPANDN
	TVPANDND
	TVPANDNQ
	TVPANDQ
	TVPAVGB
	TVPAVGW
	TVPBLENDD
	TVPBLENDMB
	TVPBLENDMD
	TVPBLENDMQ
	TVPBLENDMW
	TVPBLENDVB
	TVPBLENDW
	TVPBROADCASTB
	TVPBROADCASTD
	TVPBROADCASTMB2Q
	TVPBROADCASTMW2D
	TVPBROADCASTQ
	TVPBROADCASTW
	TVPCLMULQDQ
	TVPCMOV
	TVPCMPB
	TVPCMPD
	TVPCMPEQB
	TVPCMPEQD
	TVPCMPEQQ
	TVPCMPEQW
	TVPCMPESTRI
	TVPCMPESTRM
	TVPCMPGTB
	TVPCMPGTD
	TVPCMPGTQ
	TVPCMPGTW
	TVPCMPISTRI
	TVPCMPISTRM
	TVPCMPQ
	TVPCMPUB
	TVPCMPUD
	TVPCMPUQ
	TVPCMPUW
	TVPCMPW
	TVPCOMB
	TVPCOMD
	TVPCOMPRESSD
	TVPCOMPRESSQ
	TVPCOMQ
	TVPCOMUB
	TVPCOMUD
	TVPCOMUQ
	TVPCOMUW
	TVPCOMW
	TVPCONFLICTD
	TVPCONFLICTQ
	TVPERM2F128
	TVPERM2I128
	TVPERMB
	TVPERMD
	TVPERMI2B
	TVPERMI2D
	TVPERMI2PD
	TVPERMI2PS
	TVPERMI2Q
	TVPERMI2W
	TVPERMIL2PD
	TVPERMIL2PS
	TVPERMILPD
	TVPERMILPS
	TVPERMPD
	TVPERMPS
	TVPERMQ
	TVPERMT2B
	TVPERMT2D
	TVPERMT2PD
	TVPERMT2PS
	TVPERMT2Q
	TVPERMT2W
	TVPERMW
	TVPEXPANDD
	TVPEXPANDQ
	TVPEXTRB
	TVPEXTRD
	TVPEXTRQ
	TVPEXTRW
	TVPGATHERDD
	TVPGATHERDQ
	TVPGATHERQD
	TVPGATHERQQ
	TVPHADDBD
	TVPHADDBQ
	TVPHADDBW
	TVPHADDD
	TVPHADDDQ
	TVPHADDSW
	TVPHADDUBD
	TVPHADDUBQ
	TVPHADDUBW
	TVPHADDUDQ
	TVPHADDUWD
	TVPHADDUWQ
	TVPHADDW
	TVPHADDWD
	TVPHADDWQ
	TVPHMINPOSUW
	TVPHSUBBW
	TVPHSUBD
	TVPHSUBDQ
	TVPHSUBSW
	TVPHSUBW
	TVPHSUBWD
	TVPINSRB
	TVPINSRD
	TVPINSRQ
	TVPINSRW
	TVPLZCNTD
	TVPLZCNTQ
	TVPMACSDD
	TVPMACSDQH
	TVPMACSDQL
	TVPMACSSDD
	TVPMACSSDQH
	TVPMACSSDQL
	TVPMACSSWD
	TVPMACSSWW
	TVPMACSWD
	TVPMACSWW
	TVPMADCSSWD
	TVPMADCSWD
	TVPMADD52HUQ
	TVPMADD52LUQ
	TVPMADDUBSW
	TVPMADDWD
	TVPMASKMOVD
	TVPMASKMOVQ
	TVPMAXSB
	TVPMAXSD
	TVPMAXSQ
	TVPMAXSW
	TVPMAXUB
	TVPMAXUD
	TVPMAXUQ
	TVPMAXUW
	TVPMINSB
	TVPMINSD
	TVPMINSQ
	TVPMINSW
	TVPMINUB
	TVPMINUD
	TVPMINUQ
	TVPMINUW
	TVPMOVB2M
	TVPMOVD2M
	TVPMOVDB
	TVPMOVDW
	TVPMOVM2B
	TVPMOVM2D
	TVPMOVM2Q
	TVPMOVM2W
	TVPMOVMSKB
	TVPMOVQ2M
	TVPMOVQB
	TVPMOVQD
	TVPMOVQW
	TVPMOVSDB
	TVPMOVSDW
	TVPMOVSQB
	TVPMOVSQD
	TVPMOVSQW
	TVPMOVSWB
	TVPMOVSXBD
	TVPMOVSXBQ
	TVPMOVSXBW
	TVPMOVSXDQ
	TVPMOVSXWD
	TVPMOVSXWQ
	TVPMOVUSDB
	TVPMOVUSDW
	TVPMOVUSQB
	TVPMOVUSQD
	TVPMOVUSQW
	TVPMOVUSWB
	TVPMOVW2M
	TVPMOVWB
	TVPMOVZXBD
	TVPMOVZXBQ
	TVPMOVZXBW
	TVPMOVZXDQ
	TVPMOVZXWD
	TVPMOVZXWQ
	TVPMULDQ
	TVPMULHRSW
	TVPMULHUW
	TVPMULHW
	TVPMULLD
	TVPMULLQ
	TVPMULLW
	TVPMULTISHIFTQB
	TVPMULUDQ
	TVPOR
	TVPORD
	TVPORQ
	TVPPERM
	TVPROLD
	TVPROLQ
	TVPROLVD
	TVPROLVQ
	TVPRORD
	TVPRORQ
	TVPRORVD
	TVPRORVQ
	TVPROTB
	TVPROTD
	TVPROTQ
	TVPROTW
	TVPSADBW
	TVPSCATTERDD
	TVPSCATTERDQ
	TVPSCATTERQD
	TVPSCATTERQQ
	TVPSHAB
	TVPSHAD
	TVPSHAQ
	TVPSHAW
	TVPSHLB
	TVPSHLD
	TVPSHLQ
	TVPSHLW
	TVPSHUFB
	TVPSHUFD
	TVPSHUFHW
	TVPSHUFLW
	TVPSIGNB
	TVPSIGND
	TVPSIGNW
	TVPSLLD
	TVPSLLDQ
	TVPSLLQ
	TVPSLLVD
	TVPSLLVQ
	TVPSLLVW
	TVPSLLW
	TVPSRAD
	TVPSRAQ
	TVPSRAVD
	TVPSRAVQ
	TVPSRAVW
	TVPSRAW
	TVPSRLD
	TVPSRLDQ
	TVPSRLQ
	TVPSRLVD
	TVPSRLVQ
	TVPSRLVW
	TVPSRLW
	TVPSUBB
	TVPSUBD
	TVPSUBQ
	TVPSUBSB
	TVPSUBSW
	TVPSUBUSB
	TVPSUBUSW
	TVPSUBW
	TVPTERNLOGD
	TVPTERNLOGQ
	TVPTEST
	TVPTESTMB
	TVPTESTMD
	TVPTESTMQ
	TVPTESTMW
	TVPTESTNMB
	TVPTESTNMD
	TVPTESTNMQ
	TVPTESTNMW
	TVPUNPCKHBW
	TVPUNPCKHDQ
	TVPUNPCKHQDQ
	TVPUNPCKHWD
	TVPUNPCKLBW
	TVPUNPCKLDQ
	TVPUNPCKLQDQ
	TVPUNPCKLWD
	TVPXOR
	TVPXORD
	TVPXORQ
	TVRANGEPD
	TVRANGEPS
	TVRANGESD
	TVRANGESS
	TVRCP14PD
	TVRCP14PS
	TVRCP14SD
	TVRCP14SS
	TVRCP28PD
	TVRCP28PS
	TVRCP28SD
	TVRCP28SS
	TVRCPPS
	TVRCPSS
	TVREDUCEPD
	TVREDUCEPS
	TVREDUCESD
	TVREDUCESS
	TVRNDSCALEPD
	TVRNDSCALEPS
	TVRNDSCALESD
	TVRNDSCALESS
	TVROUNDPD
	TVROUNDPS
	TVROUNDSD
	TVROUNDSS
	TVRSQRT14PD
	TVRSQRT14PS
	TVRSQRT14SD
	TVRSQRT14SS
	TVRSQRT28PD
	TVRSQRT28PS
	TVRSQRT28SD
	TVRSQRT28SS
	TVRSQRTPS
	TVRSQRTSS
	TVSCALEFPD
	TVSCALEFPS
	TVSCALEFSD
	TVSCALEFSS
	TVSCATTERDPD
	TVSCATTERDPS
	TVSCATTERPF0DPD
	TVSCATTERPF0DPS
	TVSCATTERPF0QPD
	TVSCATTERPF0QPS
	TVSCATTERPF1DPD
	TVSCATTERPF1DPS
	TVSCATTERPF1QPD
	TVSCATTERPF1QPS
	TVSCATTERQPD
	TVSCATTERQPS
	TVSHUFF32X4
	TVSHUFF64X2
	TVSHUFI32X4
	TVSHUFI64X2
	TVSHUFPD
	TVSHUFPS
	TVSQRTPD
	TVSQRTPS
	TVSQRTSD
	TVSQRTSS
	TVSTMXCSR
	TVSUBPD
	TVSUBPS
	TVSUBSD
	TVSUBSS
	TVTESTPD
	TVTESTPS
	TVUCOMISD
	TVUCOMISS
	TVUNPCKHPD
	TVUNPCKHPS
	TVUNPCKLPD
	TVUNPCKLPS
	TVXORPD
	TVXORPS
	TVZEROALL
	TVZEROUPPER
	TXADD
	TXCHG
	TXGETBV
	TXLATB
	TXOR
	TXORPD
	TXORPS
)

var instrTypes = []InstrType{
	{TADC, "Add with Carry"},
	{TADCX, "Unsigned Integer Addition of Two Operands with Carry Flag"},
	{TADD, "Add"},
	{TADDPD, "Add Packed Double-Precision Floating-Point Values"},
	{TADDPS, "Add Packed Single-Precision Floating-Point Values"},
	{TADDSD, "Add Scalar Double-Precision Floating-Point Values"},
	{TADDSS, "Add Scalar Single-Precision Floating-Point Values"},
	{TADDSUBPD, "Packed Double-FP Add/Subtract"},
	{TADDSUBPS, "Packed Single-FP Add/Subtract"},
	{TADOX, "Unsigned Integer Addition of Two Operands with Overflow Flag"},
	{TAESDEC, "Perform One Round of an AES Decryption Flow"},
	{TAESDECLAST, "Perform Last Round of an AES Decryption Flow"},
	{TAESENC, "Perform One Round of an AES Encryption Flow"},
	{TAESENCLAST, "Perform Last Round of an AES Encryption Flow"},
	{TAESIMC, "Perform the AES InvMixColumn Transformation"},
	{TAESKEYGENASSIST, "AES Round Key Generation Assist"},
	{TAND, "Logical AND"},
	{TANDN, "Logical AND NOT"},
	{TANDNPD, "Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values"},
	{TANDNPS, "Bitwise Logical AND NOT of Packed Single-Precision Floating-Point Values"},
	{TANDPD, "Bitwise Logical AND of Packed Double-Precision Floating-Point Values"},
	{TANDPS, "Bitwise Logical AND of Packed Single-Precision Floating-Point Values"},
	{TBEXTR, "Bit Field Extract"},
	{TBLCFILL, "Fill From Lowest Clear Bit"},
	{TBLCI, "Isolate Lowest Clear Bit"},
	{TBLCIC, "Isolate Lowest Set Bit and Complement"},
	{TBLCMSK, "Mask From Lowest Clear Bit"},
	{TBLCS, "Set Lowest Clear Bit"},
	{TBLENDPD, "Blend Packed Double Precision Floating-Point Values"},
	{TBLENDPS, " Blend Packed Single Precision Floating-Point Values"},
	{TBLENDVPD, " Variable Blend Packed Double Precision Floating-Point Values"},
	{TBLENDVPS, " Variable Blend Packed Single Precision Floating-Point Values"},
	{TBLSFILL, "Fill From Lowest Set Bit"},
	{TBLSI, "Isolate Lowest Set Bit"},
	{TBLSIC, "Isolate Lowest Set Bit and Complement"},
	{TBLSMSK, "Mask From Lowest Set Bit"},
	{TBLSR, "Reset Lowest Set Bit"},
	{TBSF, "Bit Scan Forward"},
	{TBSR, "Bit Scan Reverse"},
	{TBSWAP, "Byte Swap"},
	{TBT, "Bit Test"},
	{TBTC, "Bit Test and Complement"},
	{TBTR, "Bit Test and Reset"},
	{TBTS, "Bit Test and Set"},
	{TBZHI, "Zero High Bits Starting with Specified Bit Position"},
	{TCALL, "Call Procedure"},
	{TCBW, "Convert Byte to Word"},
	{TCDQ, "Convert Doubleword to Quadword"},
	{TCDQE, "Convert Doubleword to Quadword"},
	{TCLC, "Clear Carry Flag"},
	{TCLD, "Clear Direction Flag"},
	{TCMC, "Complement Carry Flag"},
	{TCMOVA, "Move if above (CF == 0 and ZF == 0)"},
	{TCMOVAE, "Move if above or equal (CF == 0)"},
	{TCMOVB, "Move if below (CF == 1)"},
	{TCMOVBE, "Move if below or equal (CF == 1 or ZF == 1)"},
	{TCMOVC, "Move if carry (CF == 1)"},
	{TCMOVE, "Move if equal (ZF == 1)"},
	{TCMOVG, "Move if greater (ZF == 0 and SF == OF)"},
	{TCMOVGE, "Move if greater or equal (SF == OF)"},
	{TCMOVL, "Move if less (SF != OF)"},
	{TCMOVLE, "Move if less or equal (ZF == 1 or SF != OF)"},
	{TCMOVNA, "Move if not above (CF == 1 or ZF == 1)"},
	{TCMOVNAE, "Move if not above or equal (CF == 1)"},
	{TCMOVNB, "Move if not below (CF == 0)"},
	{TCMOVNBE, "Move if not below or equal (CF == 0 and ZF == 0)"},
	{TCMOVNC, "Move if not carry (CF == 0)"},
	{TCMOVNE, "Move if not equal (ZF == 0)"},
	{TCMOVNG, "Move if not greater (ZF == 1 or SF != OF)"},
	{TCMOVNGE, "Move if not greater or equal (SF != OF)"},
	{TCMOVNL, "Move if not less (SF == OF)"},
	{TCMOVNLE, "Move if not less or equal (ZF == 0 and SF == OF)"},
	{TCMOVNO, "Move if not overflow (OF == 0)"},
	{TCMOVNP, "Move if not parity (PF == 0)"},
	{TCMOVNS, "Move if not sign (SF == 0)"},
	{TCMOVNZ, "Move if not zero (ZF == 0)"},
	{TCMOVO, "Move if overflow (OF == 1)"},
	{TCMOVP, "Move if parity (PF == 1)"},
	{TCMOVPE, "Move if parity even (PF == 1)"},
	{TCMOVPO, "Move if parity odd (PF == 0)"},
	{TCMOVS, "Move if sign (SF == 1)"},
	{TCMOVZ, "Move if zero (ZF == 1)"},
	{TCMP, "Compare Two Operands"},
	{TCMPPD, "Compare Packed Double-Precision Floating-Point Values"},
	{TCMPPS, "Compare Packed Single-Precision Floating-Point Values"},
	{TCMPSD, "Compare Scalar Double-Precision Floating-Point Values"},
	{TCMPSS, "Compare Scalar Single-Precision Floating-Point Values"},
	{TCMPXCHG, "Compare and Exchange"},
	{TCMPXCHG16B, "Compare and Exchange 16 Bytes"},
	{TCMPXCHG8B, "Compare and Exchange 8 Bytes"},
	{TCOMISD, "Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS"},
	{TCOMISS, "Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS"},
	{TCPUID, "CPU Identification"},
	{TCQO, "Convert Quadword to Octaword"},
	{TCRC32, "Accumulate CRC32 Value"},
	{TCVTDQ2PD, "Convert Packed Dword Integers to Packed Double-Precision FP Values"},
	{TCVTDQ2PS, "Convert Packed Dword Integers to Packed Single-Precision FP Values"},
	{TCVTPD2DQ, "Convert Packed Double-Precision FP Values to Packed Dword Integers"},
	{TCVTPD2PI, "Convert Packed Double-Precision FP Values to Packed Dword Integers"},
	{TCVTPD2PS, "Convert Packed Double-Precision FP Values to Packed Single-Precision FP Values"},
	{TCVTPI2PD, "Convert Packed Dword Integers to Packed Double-Precision FP Values"},
	{TCVTPI2PS, "Convert Packed Dword Integers to Packed Single-Precision FP Values"},
	{TCVTPS2DQ, "Convert Packed Single-Precision FP Values to Packed Dword Integers"},
	{TCVTPS2PD, "Convert Packed Single-Precision FP Values to Packed Double-Precision FP Values"},
	{TCVTPS2PI, "Convert Packed Single-Precision FP Values to Packed Dword Integers"},
	{TCVTSD2SI, "Convert Scalar Double-Precision FP Value to Integer"},
	{TCVTSD2SS, "Convert Scalar Double-Precision FP Value to Scalar Single-Precision FP Value"},
	{TCVTSI2SD, "Convert Dword Integer to Scalar Double-Precision FP Value"},
	{TCVTSI2SS, "Convert Dword Integer to Scalar Single-Precision FP Value"},
	{TCVTSS2SD, "Convert Scalar Single-Precision FP Value to Scalar Double-Precision FP Value"},
	{TCVTSS2SI, "Convert Scalar Single-Precision FP Value to Dword Integer"},
	{TCVTTPD2DQ, "Convert with Truncation Packed Double-Precision FP Values to Packed Dword Integers"},
	{TCVTTPD2PI, "Convert with Truncation Packed Double-Precision FP Values to Packed Dword Integers"},
	{TCVTTPS2DQ, "Convert with Truncation Packed Single-Precision FP Values to Packed Dword Integers"},
	{TCVTTPS2PI, "Convert with Truncation Packed Single-Precision FP Values to Packed Dword Integers"},
	{TCVTTSD2SI, "Convert with Truncation Scalar Double-Precision FP Value to Signed Integer"},
	{TCVTTSS2SI, "Convert with Truncation Scalar Single-Precision FP Value to Dword Integer"},
	{TCWD, "Convert Word to Doubleword"},
	{TCWDE, "Convert Word to Doubleword"},
	{TDEC, "Decrement by 1"},
	{TDIV, "Unsigned Divide"},
	{TDIVPD, "Divide Packed Double-Precision Floating-Point Values"},
	{TDIVPS, "Divide Packed Single-Precision Floating-Point Values"},
	{TDIVSD, "Divide Scalar Double-Precision Floating-Point Values"},
	{TDIVSS, "Divide Scalar Single-Precision Floating-Point Values"},
	{TDPPD, "Dot Product of Packed Double Precision Floating-Point Values"},
	{TDPPS, "Dot Product of Packed Single Precision Floating-Point Values"},
	{TEMMS, "Exit MMX State"},
	{TEXTRACTPS, "Extract Packed Single Precision Floating-Point Value"},
	{TEXTRQ, "Extract Field"},
	{TFEMMS, "Fast Exit Multimedia State"},
	{THADDPD, "Packed Double-FP Horizontal Add"},
	{THADDPS, "Packed Single-FP Horizontal Add"},
	{THSUBPD, "Packed Double-FP Horizontal Subtract"},
	{THSUBPS, "Packed Single-FP Horizontal Subtract"},
	{TIDIV, "Signed Divide"},
	{TIMUL, "Signed Multiply"},
	{TINC, "Increment by 1"},
	{TINSERTPS, "Insert Packed Single Precision Floating-Point Value"},
	{TINSERTQ, "Insert Field"},
	{TINT, "Call to Interrupt Procedure"},
	{TJA, "Jump if above (CF == 0 and ZF == 0)"},
	{TJAE, "Jump if above or equal (CF == 0)"},
	{TJB, "Jump if below (CF == 1)"},
	{TJBE, "Jump if below or equal (CF == 1 or ZF == 1)"},
	{TJC, "Jump if carry (CF == 1)"},
	{TJE, "Jump if equal (ZF == 1)"},
	{TJECXZ, "Jump if ECX register is 0"},
	{TJG, "Jump if greater (ZF == 0 and SF == OF)"},
	{TJGE, "Jump if greater or equal (SF == OF)"},
	{TJL, "Jump if less (SF != OF)"},
	{TJLE, "Jump if less or equal (ZF == 1 or SF != OF)"},
	{TJMP, "Jump Unconditionally"},
	{TJNA, "Jump if not above (CF == 1 or ZF == 1)"},
	{TJNAE, "Jump if not above or equal (CF == 1)"},
	{TJNB, "Jump if not below (CF == 0)"},
	{TJNBE, "Jump if not below or equal (CF == 0 and ZF == 0)"},
	{TJNC, "Jump if not carry (CF == 0)"},
	{TJNE, "Jump if not equal (ZF == 0)"},
	{TJNG, "Jump if not greater (ZF == 1 or SF != OF)"},
	{TJNGE, "Jump if not greater or equal (SF != OF)"},
	{TJNL, "Jump if not less (SF == OF)"},
	{TJNLE, "Jump if not less or equal (ZF == 0 and SF == OF)"},
	{TJNO, "Jump if not overflow (OF == 0)"},
	{TJNP, "Jump if not parity (PF == 0)"},
	{TJNS, "Jump if not sign (SF == 0)"},
	{TJNZ, "Jump if not zero (ZF == 0)"},
	{TJO, "Jump if overflow (OF == 1)"},
	{TJP, "Jump if parity (PF == 1)"},
	{TJPE, "Jump if parity even (PF == 1)"},
	{TJPO, "Jump if parity odd (PF == 0)"},
	{TJRCXZ, "Jump if RCX register is 0"},
	{TJS, "Jump if sign (SF == 1)"},
	{TJZ, "Jump if zero (ZF == 1)"},
	{TKADDB, "ADD Two 8-bit Masks"},
	{TKADDD, "ADD Two 32-bit Masks"},
	{TKADDQ, "ADD Two 64-bit Masks"},
	{TKADDW, "ADD Two 16-bit Masks"},
	{TKANDB, "Bitwise Logical AND 8-bit Masks"},
	{TKANDD, "Bitwise Logical AND 32-bit Masks"},
	{TKANDNB, "Bitwise Logical AND NOT 8-bit Masks"},
	{TKANDND, "Bitwise Logical AND NOT 32-bit Masks"},
	{TKANDNQ, "Bitwise Logical AND NOT 64-bit Masks"},
	{TKANDNW, "Bitwise Logical AND NOT 16-bit Masks"},
	{TKANDQ, "Bitwise Logical AND 64-bit Masks"},
	{TKANDW, "Bitwise Logical AND 16-bit Masks"},
	{TKMOVB, "Move 8-bit Mask"},
	{TKMOVD, "Move 32-bit Mask"},
	{TKMOVQ, "Move 64-bit Mask"},
	{TKMOVW, "Move 16-bit Mask"},
	{TKNOTB, "NOT 8-bit Mask Register"},
	{TKNOTD, "NOT 32-bit Mask Register"},
	{TKNOTQ, "NOT 64-bit Mask Register"},
	{TKNOTW, "NOT 16-bit Mask Register"},
	{TKORB, "Bitwise Logical OR 8-bit Masks"},
	{TKORD, "Bitwise Logical OR 32-bit Masks"},
	{TKORQ, "Bitwise Logical OR 64-bit Masks"},
	{TKORTESTB, "OR 8-bit Masks and Set Flags"},
	{TKORTESTD, "OR 32-bit Masks and Set Flags"},
	{TKORTESTQ, "OR 64-bit Masks and Set Flags"},
	{TKORTESTW, "OR 16-bit Masks and Set Flags"},
	{TKORW, "Bitwise Logical OR 16-bit Masks"},
	{TKSHIFTLB, "Shift Left 8-bit Masks"},
	{TKSHIFTLD, "Shift Left 32-bit Masks"},
	{TKSHIFTLQ, "Shift Left 64-bit Masks"},
	{TKSHIFTLW, "Shift Left 16-bit Masks"},
	{TKSHIFTRB, "Shift Right 8-bit Masks"},
	{TKSHIFTRD, "Shift Right 32-bit Masks"},
	{TKSHIFTRQ, "Shift Right 64-bit Masks"},
	{TKSHIFTRW, "Shift Right 16-bit Masks"},
	{TKTESTB, "Bit Test 8-bit Masks and Set Flags"},
	{TKTESTD, "Bit Test 32-bit Masks and Set Flags"},
	{TKTESTQ, "Bit Test 64-bit Masks and Set Flags"},
	{TKTESTW, "Bit Test 16-bit Masks and Set Flags"},
	{TKUNPCKBW, "Unpack and Interleave 8-bit Masks"},
	{TKUNPCKDQ, "Unpack and Interleave 32-bit Masks"},
	{TKUNPCKWD, "Unpack and Interleave 16-bit Masks"},
	{TKXNORB, "Bitwise Logical XNOR 8-bit Masks"},
	{TKXNORD, "Bitwise Logical XNOR 32-bit Masks"},
	{TKXNORQ, "Bitwise Logical XNOR 64-bit Masks"},
	{TKXNORW, "Bitwise Logical XNOR 16-bit Masks"},
	{TKXORB, "Bitwise Logical XOR 8-bit Masks"},
	{TKXORD, "Bitwise Logical XOR 32-bit Masks"},
	{TKXORQ, "Bitwise Logical XOR 64-bit Masks"},
	{TKXORW, "Bitwise Logical XOR 16-bit Masks"},
	{TLDDQU, "Load Unaligned Integer 128 Bits"},
	{TLDMXCSR, "Load MXCSR Register"},
	{TLEA, "Load Effective Address"},
	{TLFENCE, "Load Fence"},
	{TLZCNT, "Count the Number of Leading Zero Bits"},
	{TMASKMOVDQU, "Store Selected Bytes of Double Quadword"},
	{TMASKMOVQ, "Store Selected Bytes of Quadword"},
	{TMAXPD, "Return Maximum Packed Double-Precision Floating-Point Values"},
	{TMAXPS, "Return Maximum Packed Single-Precision Floating-Point Values"},
	{TMAXSD, "Return Maximum Scalar Double-Precision Floating-Point Value"},
	{TMAXSS, "Return Maximum Scalar Single-Precision Floating-Point Value"},
	{TMFENCE, "Memory Fence"},
	{TMINPD, "Return Minimum Packed Double-Precision Floating-Point Values"},
	{TMINPS, "Return Minimum Packed Single-Precision Floating-Point Values"},
	{TMINSD, "Return Minimum Scalar Double-Precision Floating-Point Value"},
	{TMINSS, "Return Minimum Scalar Single-Precision Floating-Point Value"},
	{TMOV, "Move"},
	{TMOVAPD, "Move Aligned Packed Double-Precision Floating-Point Values"},
	{TMOVAPS, "Move Aligned Packed Single-Precision Floating-Point Values"},
	{TMOVBE, "Move Data After Swapping Bytes"},
	{TMOVD, "Move Doubleword"},
	{TMOVDDUP, "Move One Double-FP and Duplicate"},
	{TMOVDQ2Q, "Move Quadword from XMM to MMX Technology Register"},
	{TMOVDQA, "Move Aligned Double Quadword"},
	{TMOVDQU, "Move Unaligned Double Quadword"},
	{TMOVHLPS, "Move Packed Single-Precision Floating-Point Values High to Low"},
	{TMOVHPD, "Move High Packed Double-Precision Floating-Point Value"},
	{TMOVHPS, "Move High Packed Single-Precision Floating-Point Values"},
	{TMOVLHPS, "Move Packed Single-Precision Floating-Point Values Low to High"},
	{TMOVLPD, "Move Low Packed Double-Precision Floating-Point Value"},
	{TMOVLPS, "Move Low Packed Single-Precision Floating-Point Values"},
	{TMOVMSKPD, "Extract Packed Double-Precision Floating-Point Sign Mask"},
	{TMOVMSKPS, "Extract Packed Single-Precision Floating-Point Sign Mask"},
	{TMOVNTDQ, "Store Double Quadword Using Non-Temporal Hint"},
	{TMOVNTDQA, "Load Double Quadword Non-Temporal Aligned Hint"},
	{TMOVNTI, "Store Doubleword Using Non-Temporal Hint"},
	{TMOVNTPD, "Store Packed Double-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TMOVNTPS, "Store Packed Single-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TMOVNTQ, "Store of Quadword Using Non-Temporal Hint"},
	{TMOVNTSD, "Store Scalar Double-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TMOVNTSS, "Store Scalar Single-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TMOVQ, "Move Quadword"},
	{TMOVQ2DQ, "Move Quadword from MMX Technology to XMM Register"},
	{TMOVSD, "Move Scalar Double-Precision Floating-Point Value"},
	{TMOVSHDUP, "Move Packed Single-FP High and Duplicate"},
	{TMOVSLDUP, "Move Packed Single-FP Low and Duplicate"},
	{TMOVSS, "Move Scalar Single-Precision Floating-Point Values"},
	{TMOVSX, "Move with Sign-Extension"},
	{TMOVSXD, "Move Doubleword to Quadword with Sign-Extension"},
	{TMOVUPD, "Move Unaligned Packed Double-Precision Floating-Point Values"},
	{TMOVUPS, "Move Unaligned Packed Single-Precision Floating-Point Values"},
	{TMOVZX, "Move with Zero-Extend"},
	{TMPSADBW, "Compute Multiple Packed Sums of Absolute Difference"},
	{TMUL, "Unsigned Multiply"},
	{TMULPD, "Multiply Packed Double-Precision Floating-Point Values"},
	{TMULPS, "Multiply Packed Single-Precision Floating-Point Values"},
	{TMULSD, "Multiply Scalar Double-Precision Floating-Point Values"},
	{TMULSS, "Multiply Scalar Single-Precision Floating-Point Values"},
	{TMULX, "Unsigned Multiply Without Affecting Flags"},
	{TNEG, "Two's Complement Negation"},
	{TNOP, "No Operation"},
	{TNOT, "One's Complement Negation"},
	{TOR, "Logical Inclusive OR"},
	{TORPD, "Bitwise Logical OR of Double-Precision Floating-Point Values"},
	{TORPS, "Bitwise Logical OR of Single-Precision Floating-Point Values"},
	{TPABSB, "Packed Absolute Value of Byte Integers"},
	{TPABSD, "Packed Absolute Value of Doubleword Integers"},
	{TPABSW, "Packed Absolute Value of Word Integers"},
	{TPACKSSDW, "Pack Doublewords into Words with Signed Saturation"},
	{TPACKSSWB, "Pack Words into Bytes with Signed Saturation"},
	{TPACKUSDW, "Pack Doublewords into Words with Unsigned Saturation"},
	{TPACKUSWB, "Pack Words into Bytes with Unsigned Saturation"},
	{TPADDB, "Add Packed Byte Integers"},
	{TPADDD, "Add Packed Doubleword Integers"},
	{TPADDQ, "Add Packed Quadword Integers"},
	{TPADDSB, "Add Packed Signed Byte Integers with Signed Saturation"},
	{TPADDSW, "Add Packed Signed Word Integers with Signed Saturation"},
	{TPADDUSB, "Add Packed Unsigned Byte Integers with Unsigned Saturation"},
	{TPADDUSW, "Add Packed Unsigned Word Integers with Unsigned Saturation"},
	{TPADDW, "Add Packed Word Integers"},
	{TPALIGNR, "Packed Align Right"},
	{TPAND, "Packed Bitwise Logical AND"},
	{TPANDN, "Packed Bitwise Logical AND NOT"},
	{TPAUSE, "Spin Loop Hint"},
	{TPAVGB, "Average Packed Byte Integers"},
	{TPAVGUSB, "Average Packed Byte Integers"},
	{TPAVGW, "Average Packed Word Integers"},
	{TPBLENDVB, "Variable Blend Packed Bytes"},
	{TPBLENDW, "Blend Packed Words"},
	{TPCLMULQDQ, "Carry-Less Quadword Multiplication"},
	{TPCMPEQB, "Compare Packed Byte Data for Equality"},
	{TPCMPEQD, "Compare Packed Doubleword Data for Equality"},
	{TPCMPEQQ, "Compare Packed Quadword Data for Equality"},
	{TPCMPEQW, "Compare Packed Word Data for Equality"},
	{TPCMPESTRI, "Packed Compare Explicit Length Strings, Return Index"},
	{TPCMPESTRM, "Packed Compare Explicit Length Strings, Return Mask"},
	{TPCMPGTB, "Compare Packed Signed Byte Integers for Greater Than"},
	{TPCMPGTD, "Compare Packed Signed Doubleword Integers for Greater Than"},
	{TPCMPGTQ, "Compare Packed Data for Greater Than"},
	{TPCMPGTW, "Compare Packed Signed Word Integers for Greater Than"},
	{TPCMPISTRI, "Packed Compare Implicit Length Strings, Return Index"},
	{TPCMPISTRM, "Packed Compare Implicit Length Strings, Return Mask"},
	{TPDEP, "Parallel Bits Deposit"},
	{TPEXT, "Parallel Bits Extract"},
	{TPEXTRB, "Extract Byte"},
	{TPEXTRD, "Extract Doubleword"},
	{TPEXTRQ, "Extract Quadword"},
	{TPEXTRW, "Extract Word"},
	{TPF2ID, "Packed Floating-Point to Integer Doubleword Converson"},
	{TPF2IW, "Packed Floating-Point to Integer Word Conversion"},
	{TPFACC, "Packed Floating-Point Accumulate"},
	{TPFADD, "Packed Floating-Point Add"},
	{TPFCMPEQ, "Packed Floating-Point Compare for Equal"},
	{TPFCMPGE, "Packed Floating-Point Compare for Greater or Equal"},
	{TPFCMPGT, "Packed Floating-Point Compare for Greater Than"},
	{TPFMAX, "Packed Floating-Point Maximum"},
	{TPFMIN, "Packed Floating-Point Minimum"},
	{TPFMUL, "Packed Floating-Point Multiply"},
	{TPFNACC, "Packed Floating-Point Negative Accumulate"},
	{TPFPNACC, "Packed Floating-Point Positive-Negative Accumulate"},
	{TPFRCP, "Packed Floating-Point Reciprocal Approximation"},
	{TPFRCPIT1, "Packed Floating-Point Reciprocal Iteration 1"},
	{TPFRCPIT2, "Packed Floating-Point Reciprocal Iteration 2"},
	{TPFRSQIT1, "Packed Floating-Point Reciprocal Square Root Iteration 1"},
	{TPFRSQRT, "Packed Floating-Point Reciprocal Square Root Approximation"},
	{TPFSUB, "Packed Floating-Point Subtract"},
	{TPFSUBR, "Packed Floating-Point Subtract Reverse"},
	{TPHADDD, "Packed Horizontal Add Doubleword Integer"},
	{TPHADDSW, "Packed Horizontal Add Signed Word Integers with Signed Saturation"},
	{TPHADDW, "Packed Horizontal Add Word Integers"},
	{TPHMINPOSUW, "Packed Horizontal Minimum of Unsigned Word Integers"},
	{TPHSUBD, "Packed Horizontal Subtract Doubleword Integers"},
	{TPHSUBSW, "Packed Horizontal Subtract Signed Word Integers with Signed Saturation"},
	{TPHSUBW, "Packed Horizontal Subtract Word Integers"},
	{TPI2FD, "Packed Integer to Floating-Point Doubleword Conversion"},
	{TPI2FW, "Packed Integer to Floating-Point Word Conversion"},
	{TPINSRB, "Insert Byte"},
	{TPINSRD, "Insert Doubleword"},
	{TPINSRQ, "Insert Quadword"},
	{TPINSRW, "Insert Word"},
	{TPMADDUBSW, "Multiply and Add Packed Signed and Unsigned Byte Integers"},
	{TPMADDWD, "Multiply and Add Packed Signed Word Integers"},
	{TPMAXSB, "Maximum of Packed Signed Byte Integers"},
	{TPMAXSD, "Maximum of Packed Signed Doubleword Integers"},
	{TPMAXSW, "Maximum of Packed Signed Word Integers"},
	{TPMAXUB, "Maximum of Packed Unsigned Byte Integers"},
	{TPMAXUD, "Maximum of Packed Unsigned Doubleword Integers"},
	{TPMAXUW, "Maximum of Packed Unsigned Word Integers"},
	{TPMINSB, "Minimum of Packed Signed Byte Integers"},
	{TPMINSD, "Minimum of Packed Signed Doubleword Integers"},
	{TPMINSW, "Minimum of Packed Signed Word Integers"},
	{TPMINUB, "Minimum of Packed Unsigned Byte Integers"},
	{TPMINUD, "Minimum of Packed Unsigned Doubleword Integers"},
	{TPMINUW, "Minimum of Packed Unsigned Word Integers"},
	{TPMOVMSKB, "Move Byte Mask"},
	{TPMOVSXBD, "Move Packed Byte Integers to Doubleword Integers with Sign Extension"},
	{TPMOVSXBQ, "Move Packed Byte Integers to Quadword Integers with Sign Extension"},
	{TPMOVSXBW, "Move Packed Byte Integers to Word Integers with Sign Extension"},
	{TPMOVSXDQ, "Move Packed Doubleword Integers to Quadword Integers with Sign Extension"},
	{TPMOVSXWD, "Move Packed Word Integers to Doubleword Integers with Sign Extension"},
	{TPMOVSXWQ, "Move Packed Word Integers to Quadword Integers with Sign Extension"},
	{TPMOVZXBD, "Move Packed Byte Integers to Doubleword Integers with Zero Extension"},
	{TPMOVZXBQ, "Move Packed Byte Integers to Quadword Integers with Zero Extension"},
	{TPMOVZXBW, "Move Packed Byte Integers to Word Integers with Zero Extension"},
	{TPMOVZXDQ, "Move Packed Doubleword Integers to Quadword Integers with Zero Extension"},
	{TPMOVZXWD, "Move Packed Word Integers to Doubleword Integers with Zero Extension"},
	{TPMOVZXWQ, "Move Packed Word Integers to Quadword Integers with Zero Extension"},
	{TPMULDQ, "Multiply Packed Signed Doubleword Integers and Store Quadword Result"},
	{TPMULHRSW, "Packed Multiply Signed Word Integers and Store High Result with Round and Scale"},
	{TPMULHRW, "Packed Multiply High Rounded Word"},
	{TPMULHUW, "Multiply Packed Unsigned Word Integers and Store High Result"},
	{TPMULHW, "Multiply Packed Signed Word Integers and Store High Result"},
	{TPMULLD, "Multiply Packed Signed Doubleword Integers and Store Low Result"},
	{TPMULLW, "Multiply Packed Signed Word Integers and Store Low Result"},
	{TPMULUDQ, "Multiply Packed Unsigned Doubleword Integers"},
	{TPOP, "Pop a Value from the Stack"},
	{TPOPCNT, "Count of Number of Bits Set to 1"},
	{TPOR, "Packed Bitwise Logical OR"},
	{TPREFETCHNTA, "Prefetch Data Into Caches using NTA Hint"},
	{TPREFETCHT0, "Prefetch Data Into Caches using T0 Hint"},
	{TPREFETCHT1, "Prefetch Data Into Caches using T1 Hint"},
	{TPREFETCHT2, "Prefetch Data Into Caches using T2 Hint"},
	{TPREFETCHW, "Prefetch Data into Caches in Anticipation of a Write"},
	{TPREFETCHWT1, "Prefetch Vector Data Into Caches with Intent to Write and T1 Hint"},
	{TPSADBW, "Compute Sum of Absolute Differences"},
	{TPSHUFB, "Packed Shuffle Bytes"},
	{TPSHUFD, "Shuffle Packed Doublewords"},
	{TPSHUFHW, "Shuffle Packed High Words"},
	{TPSHUFLW, "Shuffle Packed Low Words"},
	{TPSHUFW, "Shuffle Packed Words"},
	{TPSIGNB, "Packed Sign of Byte Integers"},
	{TPSIGND, "Packed Sign of Doubleword Integers"},
	{TPSIGNW, "Packed Sign of Word Integers"},
	{TPSLLD, "Shift Packed Doubleword Data Left Logical"},
	{TPSLLDQ, "Shift Packed Double Quadword Left Logical"},
	{TPSLLQ, "Shift Packed Quadword Data Left Logical"},
	{TPSLLW, "Shift Packed Word Data Left Logical"},
	{TPSRAD, "Shift Packed Doubleword Data Right Arithmetic"},
	{TPSRAW, "Shift Packed Word Data Right Arithmetic"},
	{TPSRLD, "Shift Packed Doubleword Data Right Logical"},
	{TPSRLDQ, "Shift Packed Double Quadword Right Logical"},
	{TPSRLQ, "Shift Packed Quadword Data Right Logical"},
	{TPSRLW, "Shift Packed Word Data Right Logical"},
	{TPSUBB, "Subtract Packed Byte Integers"},
	{TPSUBD, "Subtract Packed Doubleword Integers"},
	{TPSUBQ, "Subtract Packed Quadword Integers"},
	{TPSUBSB, "Subtract Packed Signed Byte Integers with Signed Saturation"},
	{TPSUBSW, "Subtract Packed Signed Word Integers with Signed Saturation"},
	{TPSUBUSB, "Subtract Packed Unsigned Byte Integers with Unsigned Saturation"},
	{TPSUBUSW, "Subtract Packed Unsigned Word Integers with Unsigned Saturation"},
	{TPSUBW, "Subtract Packed Word Integers"},
	{TPSWAPD, "Packed Swap Doubleword"},
	{TPTEST, "Packed Logical Compare"},
	{TPUNPCKHBW, "Unpack and Interleave High-Order Bytes into Words"},
	{TPUNPCKHDQ, "Unpack and Interleave High-Order Doublewords into Quadwords"},
	{TPUNPCKHQDQ, "Unpack and Interleave High-Order Quadwords into Double Quadwords"},
	{TPUNPCKHWD, "Unpack and Interleave High-Order Words into Doublewords"},
	{TPUNPCKLBW, "Unpack and Interleave Low-Order Bytes into Words"},
	{TPUNPCKLDQ, "Unpack and Interleave Low-Order Doublewords into Quadwords"},
	{TPUNPCKLQDQ, "Unpack and Interleave Low-Order Quadwords into Double Quadwords"},
	{TPUNPCKLWD, "Unpack and Interleave Low-Order Words into Doublewords"},
	{TPUSH, "Push Value Onto the Stack"},
	{TPXOR, "Packed Bitwise Logical Exclusive OR"},
	{TRCL, "Rotate Left through Carry Flag"},
	{TRCPPS, "Compute Approximate Reciprocals of Packed Single-Precision Floating-Point Values"},
	{TRCPSS, "Compute Approximate Reciprocal of Scalar Single-Precision Floating-Point Values"},
	{TRCR, "Rotate Right through Carry Flag"},
	{TRDRAND, "Read Random Number"},
	{TRDSEED, "Read Random SEED"},
	{TRDTSC, "Read Time-Stamp Counter"},
	{TRDTSCP, "Read Time-Stamp Counter and Processor ID"},
	{TRET, "Return from Procedure"},
	{TROL, "Rotate Left"},
	{TROR, "Rotate Right"},
	{TRORX, "Rotate Right Logical Without Affecting Flags"},
	{TROUNDPD, "Round Packed Double Precision Floating-Point Values"},
	{TROUNDPS, "Round Packed Single Precision Floating-Point Values"},
	{TROUNDSD, "Round Scalar Double Precision Floating-Point Values"},
	{TROUNDSS, "Round Scalar Single Precision Floating-Point Values"},
	{TRSQRTPS, "Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values"},
	{TRSQRTSS, "Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value"},
	{TSAL, "Arithmetic Shift Left"},
	{TSAR, "Arithmetic Shift Right"},
	{TSARX, "Arithmetic Shift Right Without Affecting Flags"},
	{TSBB, "Subtract with Borrow"},
	{TSETA, "Set byte if above (CF == 0 and ZF == 0)"},
	{TSETAE, "Set byte if above or equal (CF == 0)"},
	{TSETB, "Set byte if below (CF == 1)"},
	{TSETBE, "Set byte if below or equal (CF == 1 or ZF == 1)"},
	{TSETC, "Set byte if carry (CF == 1)"},
	{TSETE, "Set byte if equal (ZF == 1)"},
	{TSETG, "Set byte if greater (ZF == 0 and SF == OF)"},
	{TSETGE, "Set byte if greater or equal (SF == OF)"},
	{TSETL, "Set byte if less (SF != OF)"},
	{TSETLE, "Set byte if less or equal (ZF == 1 or SF != OF)"},
	{TSETNA, "Set byte if not above (CF == 1 or ZF == 1)"},
	{TSETNAE, "Set byte if not above or equal (CF == 1)"},
	{TSETNB, "Set byte if not below (CF == 0)"},
	{TSETNBE, "Set byte if not below or equal (CF == 0 and ZF == 0)"},
	{TSETNC, "Set byte if not carry (CF == 0)"},
	{TSETNE, "Set byte if not equal (ZF == 0)"},
	{TSETNG, "Set byte if not greater (ZF == 1 or SF != OF)"},
	{TSETNGE, "Set byte if not greater or equal (SF != OF)"},
	{TSETNL, "Set byte if not less (SF == OF)"},
	{TSETNLE, "Set byte if not less or equal (ZF == 0 and SF == OF)"},
	{TSETNO, "Set byte if not overflow (OF == 0)"},
	{TSETNP, "Set byte if not parity (PF == 0)"},
	{TSETNS, "Set byte if not sign (SF == 0)"},
	{TSETNZ, "Set byte if not zero (ZF == 0)"},
	{TSETO, "Set byte if overflow (OF == 1)"},
	{TSETP, "Set byte if parity (PF == 1)"},
	{TSETPE, "Set byte if parity even (PF == 1)"},
	{TSETPO, "Set byte if parity odd (PF == 0)"},
	{TSETS, "Set byte if sign (SF == 1)"},
	{TSETZ, "Set byte if zero (ZF == 1)"},
	{TSFENCE, "Store Fence"},
	{TSHA1MSG1, "Perform an Intermediate Calculation for the Next Four SHA1 Message Doublewords"},
	{TSHA1MSG2, "Perform a Final Calculation for the Next Four SHA1 Message Doublewords"},
	{TSHA1NEXTE, "Calculate SHA1 State Variable E after Four Rounds"},
	{TSHA1RNDS4, "Perform Four Rounds of SHA1 Operation"},
	{TSHA256MSG1, "Perform an Intermediate Calculation for the Next Four SHA256 Message Doublewords"},
	{TSHA256MSG2, "Perform a Final Calculation for the Next Four SHA256 Message Doublewords"},
	{TSHA256RNDS2, "Perform Two Rounds of SHA256 Operation"},
	{TSHL, "Logical Shift Left"},
	{TSHLD, "Integer Double Precision Shift Left"},
	{TSHLX, "Logical Shift Left Without Affecting Flags"},
	{TSHR, "Logical Shift Right"},
	{TSHRD, "Integer Double Precision Shift Right"},
	{TSHRX, "Logical Shift Right Without Affecting Flags"},
	{TSHUFPD, "Shuffle Packed Double-Precision Floating-Point Values"},
	{TSHUFPS, "Shuffle Packed Single-Precision Floating-Point Values"},
	{TSQRTPD, "Compute Square Roots of Packed Double-Precision Floating-Point Values"},
	{TSQRTPS, "Compute Square Roots of Packed Single-Precision Floating-Point Values"},
	{TSQRTSD, "Compute Square Root of Scalar Double-Precision Floating-Point Value"},
	{TSQRTSS, "Compute Square Root of Scalar Single-Precision Floating-Point Value"},
	{TSTC, "Set Carry Flag"},
	{TSTD, "Set Direction Flag"},
	{TSTMXCSR, "Store MXCSR Register State"},
	{TSUB, "Subtract"},
	{TSUBPD, "Subtract Packed Double-Precision Floating-Point Values"},
	{TSUBPS, "Subtract Packed Single-Precision Floating-Point Values"},
	{TSUBSD, "Subtract Scalar Double-Precision Floating-Point Values"},
	{TSUBSS, "Subtract Scalar Single-Precision Floating-Point Values"},
	{TT1MSKC, "Inverse Mask From Trailing Ones"},
	{TTEST, "Logical Compare"},
	{TTZCNT, "Count the Number of Trailing Zero Bits"},
	{TTZMSK, "Mask From Trailing Zeros"},
	{TUCOMISD, "Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS"},
	{TUCOMISS, "Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS"},
	{TUD2, "Undefined Instruction"},
	{TUNPCKHPD, "Unpack and Interleave High Packed Double-Precision Floating-Point Values"},
	{TUNPCKHPS, "Unpack and Interleave High Packed Single-Precision Floating-Point Values"},
	{TUNPCKLPD, "Unpack and Interleave Low Packed Double-Precision Floating-Point Values"},
	{TUNPCKLPS, "Unpack and Interleave Low Packed Single-Precision Floating-Point Values"},
	{TVADDPD, "Add Packed Double-Precision Floating-Point Values"},
	{TVADDPS, "Add Packed Single-Precision Floating-Point Values"},
	{TVADDSD, "Add Scalar Double-Precision Floating-Point Values"},
	{TVADDSS, "Add Scalar Single-Precision Floating-Point Values"},
	{TVADDSUBPD, "Packed Double-FP Add/Subtract"},
	{TVADDSUBPS, "Packed Single-FP Add/Subtract"},
	{TVAESDEC, "Perform One Round of an AES Decryption Flow"},
	{TVAESDECLAST, "Perform Last Round of an AES Decryption Flow"},
	{TVAESENC, "Perform One Round of an AES Encryption Flow"},
	{TVAESENCLAST, "Perform Last Round of an AES Encryption Flow"},
	{TVAESIMC, "Perform the AES InvMixColumn Transformation"},
	{TVAESKEYGENASSIST, "AES Round Key Generation Assist"},
	{TVALIGND, "Align Doubleword Vectors"},
	{TVALIGNQ, "Align Quadword Vectors"},
	{TVANDNPD, "Bitwise Logical AND NOT of Packed Double-Precision Floating-Point Values"},
	{TVANDNPS, "Bitwise Logical AND NOT of Packed Single-Precision Floating-Point Values"},
	{TVANDPD, "Bitwise Logical AND of Packed Double-Precision Floating-Point Values"},
	{TVANDPS, "Bitwise Logical AND of Packed Single-Precision Floating-Point Values"},
	{TVBLENDMPD, "Blend Packed Double-Precision Floating-Point Vectors Using an OpMask Control"},
	{TVBLENDMPS, "Blend Packed Single-Precision Floating-Point Vectors Using an OpMask Control"},
	{TVBLENDPD, "Blend Packed Double Precision Floating-Point Values"},
	{TVBLENDPS, " Blend Packed Single Precision Floating-Point Values"},
	{TVBLENDVPD, " Variable Blend Packed Double Precision Floating-Point Values"},
	{TVBLENDVPS, " Variable Blend Packed Single Precision Floating-Point Values"},
	{TVBROADCASTF128, "Broadcast 128 Bit of Floating-Point Data"},
	{TVBROADCASTF32X2, "Broadcast Two Single-Precision Floating-Point Elements"},
	{TVBROADCASTF32X4, "Broadcast Four Single-Precision Floating-Point Elements"},
	{TVBROADCASTF32X8, "Broadcast Eight Single-Precision Floating-Point Elements"},
	{TVBROADCASTF64X2, "Broadcast Two Double-Precision Floating-Point Elements"},
	{TVBROADCASTF64X4, "Broadcast Four Double-Precision Floating-Point Elements"},
	{TVBROADCASTI128, "Broadcast 128 Bits of Integer Data"},
	{TVBROADCASTI32X2, "Broadcast Two Doubleword Elements"},
	{TVBROADCASTI32X4, "Broadcast Four Doubleword Elements"},
	{TVBROADCASTI32X8, "Broadcast Eight Doubleword Elements"},
	{TVBROADCASTI64X2, "Broadcast Two Quadword Elements"},
	{TVBROADCASTI64X4, "Broadcast Four Quadword Elements"},
	{TVBROADCASTSD, "Broadcast Double-Precision Floating-Point Element"},
	{TVBROADCASTSS, "Broadcast Single-Precision Floating-Point Element"},
	{TVCMPPD, "Compare Packed Double-Precision Floating-Point Values"},
	{TVCMPPS, "Compare Packed Single-Precision Floating-Point Values"},
	{TVCMPSD, "Compare Scalar Double-Precision Floating-Point Values"},
	{TVCMPSS, "Compare Scalar Single-Precision Floating-Point Values"},
	{TVCOMISD, "Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS"},
	{TVCOMISS, "Compare Scalar Ordered Single-Precision Floating-Point Values and Set EFLAGS"},
	{TVCOMPRESSPD, "Store Sparse Packed Double-Precision Floating-Point Values into Dense Memory/Register"},
	{TVCOMPRESSPS, "Store Sparse Packed Single-Precision Floating-Point Values into Dense Memory/Register"},
	{TVCVTDQ2PD, "Convert Packed Dword Integers to Packed Double-Precision FP Values"},
	{TVCVTDQ2PS, "Convert Packed Dword Integers to Packed Single-Precision FP Values"},
	{TVCVTPD2DQ, "Convert Packed Double-Precision FP Values to Packed Dword Integers"},
	{TVCVTPD2PS, "Convert Packed Double-Precision FP Values to Packed Single-Precision FP Values"},
	{TVCVTPD2QQ, "Convert Packed Double-Precision Floating-Point Values to Packed Quadword Integers"},
	{TVCVTPD2UDQ, "Convert Packed Double-Precision Floating-Point Values to Packed Unsigned Doubleword Integers"},
	{TVCVTPD2UQQ, "Convert Packed Double-Precision Floating-Point Values to Packed Unsigned Quadword Integers"},
	{TVCVTPH2PS, "Convert Half-Precision FP Values to Single-Precision FP Values"},
	{TVCVTPS2DQ, "Convert Packed Single-Precision FP Values to Packed Dword Integers"},
	{TVCVTPS2PD, "Convert Packed Single-Precision FP Values to Packed Double-Precision FP Values"},
	{TVCVTPS2PH, "Convert Single-Precision FP value to Half-Precision FP value"},
	{TVCVTPS2QQ, "Convert Packed Single Precision Floating-Point Values to Packed Singed Quadword Integer Values"},
	{TVCVTPS2UDQ, "Convert Packed Single-Precision Floating-Point Values to Packed Unsigned Doubleword Integer Values"},
	{TVCVTPS2UQQ, "Convert Packed Single Precision Floating-Point Values to Packed Unsigned Quadword Integer Values"},
	{TVCVTQQ2PD, "Convert Packed Quadword Integers to Packed Double-Precision Floating-Point Values"},
	{TVCVTQQ2PS, "Convert Packed Quadword Integers to Packed Single-Precision Floating-Point Values"},
	{TVCVTSD2SI, "Convert Scalar Double-Precision FP Value to Integer"},
	{TVCVTSD2SS, "Convert Scalar Double-Precision FP Value to Scalar Single-Precision FP Value"},
	{TVCVTSD2USI, "Convert Scalar Double-Precision Floating-Point Value to Unsigned Doubleword Integer"},
	{TVCVTSI2SD, "Convert Dword Integer to Scalar Double-Precision FP Value"},
	{TVCVTSI2SS, "Convert Dword Integer to Scalar Single-Precision FP Value"},
	{TVCVTSS2SD, "Convert Scalar Single-Precision FP Value to Scalar Double-Precision FP Value"},
	{TVCVTSS2SI, "Convert Scalar Single-Precision FP Value to Dword Integer"},
	{TVCVTSS2USI, "Convert Scalar Single-Precision Floating-Point Value to Unsigned Doubleword Integer"},
	{TVCVTTPD2DQ, "Convert with Truncation Packed Double-Precision FP Values to Packed Dword Integers"},
	{TVCVTTPD2QQ, "Convert with Truncation Packed Double-Precision Floating-Point Values to Packed Quadword Integers"},
	{TVCVTTPD2UDQ, "Convert with Truncation Packed Double-Precision Floating-Point Values to Packed Unsigned Doubleword Integers"},
	{TVCVTTPD2UQQ, "Convert with Truncation Packed Double-Precision Floating-Point Values to Packed Unsigned Quadword Integers"},
	{TVCVTTPS2DQ, "Convert with Truncation Packed Single-Precision FP Values to Packed Dword Integers"},
	{TVCVTTPS2QQ, "Convert with Truncation Packed Single Precision Floating-Point Values to Packed Singed Quadword Integer Values"},
	{TVCVTTPS2UDQ, "Convert with Truncation Packed Single-Precision Floating-Point Values to Packed Unsigned Doubleword Integer Values"},
	{TVCVTTPS2UQQ, "Convert with Truncation Packed Single Precision Floating-Point Values to Packed Unsigned Quadword Integer Values"},
	{TVCVTTSD2SI, "Convert with Truncation Scalar Double-Precision FP Value to Signed Integer"},
	{TVCVTTSD2USI, "Convert with Truncation Scalar Double-Precision Floating-Point Value to Unsigned Integer"},
	{TVCVTTSS2SI, "Convert with Truncation Scalar Single-Precision FP Value to Dword Integer"},
	{TVCVTTSS2USI, "Convert with Truncation Scalar Single-Precision Floating-Point Value to Unsigned Integer"},
	{TVCVTUDQ2PD, "Convert Packed Unsigned Doubleword Integers to Packed Double-Precision Floating-Point Values"},
	{TVCVTUDQ2PS, "Convert Packed Unsigned Doubleword Integers to Packed Single-Precision Floating-Point Values"},
	{TVCVTUQQ2PD, "Convert Packed Unsigned Quadword Integers to Packed Double-Precision Floating-Point Values"},
	{TVCVTUQQ2PS, "Convert Packed Unsigned Quadword Integers to Packed Single-Precision Floating-Point Values"},
	{TVCVTUSI2SD, "Convert Unsigned Integer to Scalar Double-Precision Floating-Point Value"},
	{TVCVTUSI2SS, "Convert Unsigned Integer to Scalar Single-Precision Floating-Point Value"},
	{TVDBPSADBW, "Double Block Packed Sum-Absolute-Differences on Unsigned Bytes"},
	{TVDIVPD, "Divide Packed Double-Precision Floating-Point Values"},
	{TVDIVPS, "Divide Packed Single-Precision Floating-Point Values"},
	{TVDIVSD, "Divide Scalar Double-Precision Floating-Point Values"},
	{TVDIVSS, "Divide Scalar Single-Precision Floating-Point Values"},
	{TVDPPD, "Dot Product of Packed Double Precision Floating-Point Values"},
	{TVDPPS, "Dot Product of Packed Single Precision Floating-Point Values"},
	{TVEXP2PD, "Approximation to the Exponential 2^x of Packed Double-Precision Floating-Point Values with Less Than 2^-23 Relative Error"},
	{TVEXP2PS, "Approximation to the Exponential 2^x of Packed Single-Precision Floating-Point Values with Less Than 2^-23 Relative Error"},
	{TVEXPANDPD, "Load Sparse Packed Double-Precision Floating-Point Values from Dense Memory"},
	{TVEXPANDPS, "Load Sparse Packed Single-Precision Floating-Point Values from Dense Memory"},
	{TVEXTRACTF128, "Extract Packed Floating-Point Values"},
	{TVEXTRACTF32X4, "Extract 128 Bits of Packed Single-Precision Floating-Point Values"},
	{TVEXTRACTF32X8, "Extract 256 Bits of Packed Single-Precision Floating-Point Values"},
	{TVEXTRACTF64X2, "Extract 128 Bits of Packed Double-Precision Floating-Point Values"},
	{TVEXTRACTF64X4, "Extract 256 Bits of Packed Double-Precision Floating-Point Values"},
	{TVEXTRACTI128, "Extract Packed Integer Values"},
	{TVEXTRACTI32X4, "Extract 128 Bits of Packed Doubleword Integer Values"},
	{TVEXTRACTI32X8, "Extract 256 Bits of Packed Doubleword Integer Values"},
	{TVEXTRACTI64X2, "Extract 128 Bits of Packed Quadword Integer Values"},
	{TVEXTRACTI64X4, "Extract 256 Bits of Packed Quadword Integer Values"},
	{TVEXTRACTPS, "Extract Packed Single Precision Floating-Point Value"},
	{TVFIXUPIMMPD, "Fix Up Special Packed Double-Precision Floating-Point Values"},
	{TVFIXUPIMMPS, "Fix Up Special Packed Single-Precision Floating-Point Values"},
	{TVFIXUPIMMSD, "Fix Up Special Scalar Double-Precision Floating-Point Value"},
	{TVFIXUPIMMSS, "Fix Up Special Scalar Single-Precision Floating-Point Value"},
	{TVFMADD132PD, "Fused Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFMADD132PS, "Fused Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFMADD132SD, "Fused Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFMADD132SS, "Fused Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFMADD213PD, "Fused Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFMADD213PS, "Fused Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFMADD213SD, "Fused Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFMADD213SS, "Fused Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFMADD231PD, "Fused Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFMADD231PS, "Fused Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFMADD231SD, "Fused Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFMADD231SS, "Fused Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFMADDPD, "Fused Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFMADDPS, "Fused Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFMADDSD, "Fused Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFMADDSS, "Fused Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFMADDSUB132PD, "Fused Multiply-Alternating Add/Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMADDSUB132PS, "Fused Multiply-Alternating Add/Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMADDSUB213PD, "Fused Multiply-Alternating Add/Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMADDSUB213PS, "Fused Multiply-Alternating Add/Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMADDSUB231PD, "Fused Multiply-Alternating Add/Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMADDSUB231PS, "Fused Multiply-Alternating Add/Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMADDSUBPD, "Fused Multiply-Alternating Add/Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMADDSUBPS, "Fused Multiply-Alternating Add/Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMSUB132PD, "Fused Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMSUB132PS, "Fused Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMSUB132SD, "Fused Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFMSUB132SS, "Fused Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFMSUB213PD, "Fused Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMSUB213PS, "Fused Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMSUB213SD, "Fused Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFMSUB213SS, "Fused Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFMSUB231PD, "Fused Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMSUB231PS, "Fused Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMSUB231SD, "Fused Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFMSUB231SS, "Fused Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFMSUBADD132PD, "Fused Multiply-Alternating Subtract/Add of Packed Double-Precision Floating-Point Values"},
	{TVFMSUBADD132PS, "Fused Multiply-Alternating Subtract/Add of Packed Single-Precision Floating-Point Values"},
	{TVFMSUBADD213PD, "Fused Multiply-Alternating Subtract/Add of Packed Double-Precision Floating-Point Values"},
	{TVFMSUBADD213PS, "Fused Multiply-Alternating Subtract/Add of Packed Single-Precision Floating-Point Values"},
	{TVFMSUBADD231PD, "Fused Multiply-Alternating Subtract/Add of Packed Double-Precision Floating-Point Values"},
	{TVFMSUBADD231PS, "Fused Multiply-Alternating Subtract/Add of Packed Single-Precision Floating-Point Values"},
	{TVFMSUBADDPD, "Fused Multiply-Alternating Subtract/Add of Packed Double-Precision Floating-Point Values"},
	{TVFMSUBADDPS, "Fused Multiply-Alternating Subtract/Add of Packed Single-Precision Floating-Point Values"},
	{TVFMSUBPD, "Fused Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFMSUBPS, "Fused Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFMSUBSD, "Fused Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFMSUBSS, "Fused Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFNMADD132PD, "Fused Negative Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFNMADD132PS, "Fused Negative Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFNMADD132SD, "Fused Negative Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFNMADD132SS, "Fused Negative Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFNMADD213PD, "Fused Negative Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFNMADD213PS, "Fused Negative Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFNMADD213SD, "Fused Negative Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFNMADD213SS, "Fused Negative Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFNMADD231PD, "Fused Negative Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFNMADD231PS, "Fused Negative Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFNMADD231SD, "Fused Negative Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFNMADD231SS, "Fused Negative Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFNMADDPD, "Fused Negative Multiply-Add of Packed Double-Precision Floating-Point Values"},
	{TVFNMADDPS, "Fused Negative Multiply-Add of Packed Single-Precision Floating-Point Values"},
	{TVFNMADDSD, "Fused Negative Multiply-Add of Scalar Double-Precision Floating-Point Values"},
	{TVFNMADDSS, "Fused Negative Multiply-Add of Scalar Single-Precision Floating-Point Values"},
	{TVFNMSUB132PD, "Fused Negative Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFNMSUB132PS, "Fused Negative Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFNMSUB132SD, "Fused Negative Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFNMSUB132SS, "Fused Negative Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFNMSUB213PD, "Fused Negative Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFNMSUB213PS, "Fused Negative Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFNMSUB213SD, "Fused Negative Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFNMSUB213SS, "Fused Negative Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFNMSUB231PD, "Fused Negative Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFNMSUB231PS, "Fused Negative Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFNMSUB231SD, "Fused Negative Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFNMSUB231SS, "Fused Negative Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFNMSUBPD, "Fused Negative Multiply-Subtract of Packed Double-Precision Floating-Point Values"},
	{TVFNMSUBPS, "Fused Negative Multiply-Subtract of Packed Single-Precision Floating-Point Values"},
	{TVFNMSUBSD, "Fused Negative Multiply-Subtract of Scalar Double-Precision Floating-Point Values"},
	{TVFNMSUBSS, "Fused Negative Multiply-Subtract of Scalar Single-Precision Floating-Point Values"},
	{TVFPCLASSPD, "Test Class of Packed Double-Precision Floating-Point Values"},
	{TVFPCLASSPS, "Test Class of Packed Single-Precision Floating-Point Values"},
	{TVFPCLASSSD, "Test Class of Scalar Double-Precision Floating-Point Value"},
	{TVFPCLASSSS, "Test Class of Scalar Single-Precision Floating-Point Value"},
	{TVFRCZPD, "Extract Fraction Packed Double-Precision Floating-Point"},
	{TVFRCZPS, "Extract Fraction Packed Single-Precision Floating-Point"},
	{TVFRCZSD, "Extract Fraction Scalar Double-Precision Floating-Point"},
	{TVFRCZSS, "Extract Fraction Scalar Single-Precision Floating Point"},
	{TVGATHERDPD, "Gather Packed Double-Precision Floating-Point Values Using Signed Doubleword Indices"},
	{TVGATHERDPS, "Gather Packed Single-Precision Floating-Point Values Using Signed Doubleword Indices"},
	{TVGATHERPF0DPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Doubleword Indices Using T0 Hint"},
	{TVGATHERPF0DPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Doubleword Indices Using T0 Hint"},
	{TVGATHERPF0QPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Quadword Indices Using T0 Hint"},
	{TVGATHERPF0QPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Quadword Indices Using T0 Hint"},
	{TVGATHERPF1DPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Doubleword Indices Using T1 Hint"},
	{TVGATHERPF1DPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Doubleword Indices Using T1 Hint"},
	{TVGATHERPF1QPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Quadword Indices Using T1 Hint"},
	{TVGATHERPF1QPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Quadword Indices Using T1 Hint"},
	{TVGATHERQPD, "Gather Packed Double-Precision Floating-Point Values Using Signed Quadword Indices"},
	{TVGATHERQPS, "Gather Packed Single-Precision Floating-Point Values Using Signed Quadword Indices"},
	{TVGETEXPPD, "Extract Exponents of Packed Double-Precision Floating-Point Values as Double-Precision Floating-Point Values"},
	{TVGETEXPPS, "Extract Exponents of Packed Single-Precision Floating-Point Values as Single-Precision Floating-Point Values"},
	{TVGETEXPSD, "Extract Exponent of Scalar Double-Precision Floating-Point Value as Double-Precision Floating-Point Value"},
	{TVGETEXPSS, "Extract Exponent of Scalar Single-Precision Floating-Point Value as Single-Precision Floating-Point Value"},
	{TVGETMANTPD, "Extract Normalized Mantissas from Packed Double-Precision Floating-Point Values"},
	{TVGETMANTPS, "Extract Normalized Mantissas from Packed Single-Precision Floating-Point Values"},
	{TVGETMANTSD, "Extract Normalized Mantissa from Scalar Double-Precision Floating-Point Value"},
	{TVGETMANTSS, "Extract Normalized Mantissa from Scalar Single-Precision Floating-Point Value"},
	{TVHADDPD, "Packed Double-FP Horizontal Add"},
	{TVHADDPS, "Packed Single-FP Horizontal Add"},
	{TVHSUBPD, "Packed Double-FP Horizontal Subtract"},
	{TVHSUBPS, "Packed Single-FP Horizontal Subtract"},
	{TVINSERTF128, "Insert Packed Floating-Point Values"},
	{TVINSERTF32X4, "Insert 128 Bits of Packed Single-Precision Floating-Point Values"},
	{TVINSERTF32X8, "Insert 256 Bits of Packed Single-Precision Floating-Point Values"},
	{TVINSERTF64X2, "Insert 128 Bits of Packed Double-Precision Floating-Point Values"},
	{TVINSERTF64X4, "Insert 256 Bits of Packed Double-Precision Floating-Point Values"},
	{TVINSERTI128, "Insert Packed Integer Values"},
	{TVINSERTI32X4, "Insert 128 Bits of Packed Doubleword Integer Values"},
	{TVINSERTI32X8, "Insert 256 Bits of Packed Doubleword Integer Values"},
	{TVINSERTI64X2, "Insert 128 Bits of Packed Quadword Integer Values"},
	{TVINSERTI64X4, "Insert 256 Bits of Packed Quadword Integer Values"},
	{TVINSERTPS, "Insert Packed Single Precision Floating-Point Value"},
	{TVLDDQU, "Load Unaligned Integer 128 Bits"},
	{TVLDMXCSR, "Load MXCSR Register"},
	{TVMASKMOVDQU, "Store Selected Bytes of Double Quadword"},
	{TVMASKMOVPD, "Conditional Move Packed Double-Precision Floating-Point Values"},
	{TVMASKMOVPS, "Conditional Move Packed Single-Precision Floating-Point Values"},
	{TVMAXPD, "Return Maximum Packed Double-Precision Floating-Point Values"},
	{TVMAXPS, "Return Maximum Packed Single-Precision Floating-Point Values"},
	{TVMAXSD, "Return Maximum Scalar Double-Precision Floating-Point Value"},
	{TVMAXSS, "Return Maximum Scalar Single-Precision Floating-Point Value"},
	{TVMINPD, "Return Minimum Packed Double-Precision Floating-Point Values"},
	{TVMINPS, "Return Minimum Packed Single-Precision Floating-Point Values"},
	{TVMINSD, "Return Minimum Scalar Double-Precision Floating-Point Value"},
	{TVMINSS, "Return Minimum Scalar Single-Precision Floating-Point Value"},
	{TVMOVAPD, "Move Aligned Packed Double-Precision Floating-Point Values"},
	{TVMOVAPS, "Move Aligned Packed Single-Precision Floating-Point Values"},
	{TVMOVD, "Move Doubleword"},
	{TVMOVDDUP, "Move One Double-FP and Duplicate"},
	{TVMOVDQA, "Move Aligned Double Quadword"},
	{TVMOVDQA32, "Move Aligned Doubleword Values"},
	{TVMOVDQA64, "Move Aligned Quadword Values"},
	{TVMOVDQU, "Move Unaligned Double Quadword"},
	{TVMOVDQU16, "Move Unaligned Word Values"},
	{TVMOVDQU32, "Move Unaligned Doubleword Values"},
	{TVMOVDQU64, "Move Unaligned Quadword Values"},
	{TVMOVDQU8, "Move Unaligned Byte Values"},
	{TVMOVHLPS, "Move Packed Single-Precision Floating-Point Values High to Low"},
	{TVMOVHPD, "Move High Packed Double-Precision Floating-Point Value"},
	{TVMOVHPS, "Move High Packed Single-Precision Floating-Point Values"},
	{TVMOVLHPS, "Move Packed Single-Precision Floating-Point Values Low to High"},
	{TVMOVLPD, "Move Low Packed Double-Precision Floating-Point Value"},
	{TVMOVLPS, "Move Low Packed Single-Precision Floating-Point Values"},
	{TVMOVMSKPD, "Extract Packed Double-Precision Floating-Point Sign Mask"},
	{TVMOVMSKPS, "Extract Packed Single-Precision Floating-Point Sign Mask"},
	{TVMOVNTDQ, "Store Double Quadword Using Non-Temporal Hint"},
	{TVMOVNTDQA, "Load Double Quadword Non-Temporal Aligned Hint"},
	{TVMOVNTPD, "Store Packed Double-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TVMOVNTPS, "Store Packed Single-Precision Floating-Point Values Using Non-Temporal Hint"},
	{TVMOVQ, "Move Quadword"},
	{TVMOVSD, "Move Scalar Double-Precision Floating-Point Value"},
	{TVMOVSHDUP, "Move Packed Single-FP High and Duplicate"},
	{TVMOVSLDUP, "Move Packed Single-FP Low and Duplicate"},
	{TVMOVSS, "Move Scalar Single-Precision Floating-Point Values"},
	{TVMOVUPD, "Move Unaligned Packed Double-Precision Floating-Point Values"},
	{TVMOVUPS, "Move Unaligned Packed Single-Precision Floating-Point Values"},
	{TVMPSADBW, "Compute Multiple Packed Sums of Absolute Difference"},
	{TVMULPD, "Multiply Packed Double-Precision Floating-Point Values"},
	{TVMULPS, "Multiply Packed Single-Precision Floating-Point Values"},
	{TVMULSD, "Multiply Scalar Double-Precision Floating-Point Values"},
	{TVMULSS, "Multiply Scalar Single-Precision Floating-Point Values"},
	{TVORPD, "Bitwise Logical OR of Double-Precision Floating-Point Values"},
	{TVORPS, "Bitwise Logical OR of Single-Precision Floating-Point Values"},
	{TVPABSB, "Packed Absolute Value of Byte Integers"},
	{TVPABSD, "Packed Absolute Value of Doubleword Integers"},
	{TVPABSQ, "Packed Absolute Value of Quadword Integers"},
	{TVPABSW, "Packed Absolute Value of Word Integers"},
	{TVPACKSSDW, "Pack Doublewords into Words with Signed Saturation"},
	{TVPACKSSWB, "Pack Words into Bytes with Signed Saturation"},
	{TVPACKUSDW, "Pack Doublewords into Words with Unsigned Saturation"},
	{TVPACKUSWB, "Pack Words into Bytes with Unsigned Saturation"},
	{TVPADDB, "Add Packed Byte Integers"},
	{TVPADDD, "Add Packed Doubleword Integers"},
	{TVPADDQ, "Add Packed Quadword Integers"},
	{TVPADDSB, "Add Packed Signed Byte Integers with Signed Saturation"},
	{TVPADDSW, "Add Packed Signed Word Integers with Signed Saturation"},
	{TVPADDUSB, "Add Packed Unsigned Byte Integers with Unsigned Saturation"},
	{TVPADDUSW, "Add Packed Unsigned Word Integers with Unsigned Saturation"},
	{TVPADDW, "Add Packed Word Integers"},
	{TVPALIGNR, "Packed Align Right"},
	{TVPAND, "Packed Bitwise Logical AND"},
	{TVPANDD, "Bitwise Logical AND of Packed Doubleword Integers"},
	{TVPANDN, "Packed Bitwise Logical AND NOT"},
	{TVPANDND, "Bitwise Logical AND NOT of Packed Doubleword Integers"},
	{TVPANDNQ, "Bitwise Logical AND NOT of Packed Quadword Integers"},
	{TVPANDQ, "Bitwise Logical AND of Packed Quadword Integers"},
	{TVPAVGB, "Average Packed Byte Integers"},
	{TVPAVGW, "Average Packed Word Integers"},
	{TVPBLENDD, "Blend Packed Doublewords"},
	{TVPBLENDMB, "Blend Byte Vectors Using an OpMask Control"},
	{TVPBLENDMD, "Blend Doubleword Vectors Using an OpMask Control"},
	{TVPBLENDMQ, "Blend Quadword Vectors Using an OpMask Control"},
	{TVPBLENDMW, "Blend Word Vectors Using an OpMask Control"},
	{TVPBLENDVB, "Variable Blend Packed Bytes"},
	{TVPBLENDW, "Blend Packed Words"},
	{TVPBROADCASTB, "Broadcast Byte Integer"},
	{TVPBROADCASTD, "Broadcast Doubleword Integer"},
	{TVPBROADCASTMB2Q, "Broadcast Low Byte of Mask Register to Packed Quadword Values"},
	{TVPBROADCASTMW2D, "Broadcast Low Word of Mask Register to Packed Doubleword Values"},
	{TVPBROADCASTQ, "Broadcast Quadword Integer"},
	{TVPBROADCASTW, "Broadcast Word Integer"},
	{TVPCLMULQDQ, "Carry-Less Quadword Multiplication"},
	{TVPCMOV, "Packed Conditional Move"},
	{TVPCMPB, "Compare Packed Signed Byte Values"},
	{TVPCMPD, "Compare Packed Signed Doubleword Values"},
	{TVPCMPEQB, "Compare Packed Byte Data for Equality"},
	{TVPCMPEQD, "Compare Packed Doubleword Data for Equality"},
	{TVPCMPEQQ, "Compare Packed Quadword Data for Equality"},
	{TVPCMPEQW, "Compare Packed Word Data for Equality"},
	{TVPCMPESTRI, "Packed Compare Explicit Length Strings, Return Index"},
	{TVPCMPESTRM, "Packed Compare Explicit Length Strings, Return Mask"},
	{TVPCMPGTB, "Compare Packed Signed Byte Integers for Greater Than"},
	{TVPCMPGTD, "Compare Packed Signed Doubleword Integers for Greater Than"},
	{TVPCMPGTQ, "Compare Packed Data for Greater Than"},
	{TVPCMPGTW, "Compare Packed Signed Word Integers for Greater Than"},
	{TVPCMPISTRI, "Packed Compare Implicit Length Strings, Return Index"},
	{TVPCMPISTRM, "Packed Compare Implicit Length Strings, Return Mask"},
	{TVPCMPQ, "Compare Packed Signed Quadword Values"},
	{TVPCMPUB, "Compare Packed Unsigned Byte Values"},
	{TVPCMPUD, "Compare Packed Unsigned Doubleword Values"},
	{TVPCMPUQ, "Compare Packed Unsigned Quadword Values"},
	{TVPCMPUW, "Compare Packed Unsigned Word Values"},
	{TVPCMPW, "Compare Packed Signed Word Values"},
	{TVPCOMB, "Compare Packed Signed Byte Integers"},
	{TVPCOMD, "Compare Packed Signed Doubleword Integers"},
	{TVPCOMPRESSD, "Store Sparse Packed Doubleword Integer Values into Dense Memory/Register"},
	{TVPCOMPRESSQ, "Store Sparse Packed Quadword Integer Values into Dense Memory/Register"},
	{TVPCOMQ, "Compare Packed Signed Quadword Integers"},
	{TVPCOMUB, "Compare Packed Unsigned Byte Integers"},
	{TVPCOMUD, "Compare Packed Unsigned Doubleword Integers"},
	{TVPCOMUQ, "Compare Packed Unsigned Quadword Integers"},
	{TVPCOMUW, "Compare Packed Unsigned Word Integers"},
	{TVPCOMW, "Compare Packed Signed Word Integers"},
	{TVPCONFLICTD, "Detect Conflicts Within a Vector of Packed Doubleword Values into Dense Memory/Register"},
	{TVPCONFLICTQ, "Detect Conflicts Within a Vector of Packed Quadword Values into Dense Memory/Register"},
	{TVPERM2F128, "Permute Floating-Point Values"},
	{TVPERM2I128, "Permute 128-Bit Integer Values"},
	{TVPERMB, "Permute Byte Integers"},
	{TVPERMD, "Permute Doubleword Integers"},
	{TVPERMI2B, "Full Permute of Bytes From Two Tables Overwriting the Index"},
	{TVPERMI2D, "Full Permute of Doublewords From Two Tables Overwriting the Index"},
	{TVPERMI2PD, "Full Permute of Double-Precision Floating-Point Values From Two Tables Overwriting the Index"},
	{TVPERMI2PS, "Full Permute of Single-Precision Floating-Point Values From Two Tables Overwriting the Index"},
	{TVPERMI2Q, "Full Permute of Quadwords From Two Tables Overwriting the Index"},
	{TVPERMI2W, "Full Permute of Words From Two Tables Overwriting the Index"},
	{TVPERMIL2PD, "Permute Two-Source Double-Precision Floating-Point Vectors"},
	{TVPERMIL2PS, "Permute Two-Source Single-Precision Floating-Point Vectors"},
	{TVPERMILPD, "Permute Double-Precision Floating-Point Values"},
	{TVPERMILPS, "Permute Single-Precision Floating-Point Values"},
	{TVPERMPD, "Permute Double-Precision Floating-Point Elements"},
	{TVPERMPS, "Permute Single-Precision Floating-Point Elements"},
	{TVPERMQ, "Permute Quadword Integers"},
	{TVPERMT2B, "Full Permute of Bytes From Two Tables Overwriting a Table"},
	{TVPERMT2D, "Full Permute of Doublewords From Two Tables Overwriting a Table"},
	{TVPERMT2PD, "Full Permute of Double-Precision Floating-Point Values From Two Tables Overwriting a Table"},
	{TVPERMT2PS, "Full Permute of Single-Precision Floating-Point Values From Two Tables Overwriting a Table"},
	{TVPERMT2Q, "Full Permute of Quadwords From Two Tables Overwriting a Table"},
	{TVPERMT2W, "Full Permute of Words From Two Tables Overwriting a Table"},
	{TVPERMW, "Permute Word Integers"},
	{TVPEXPANDD, "Load Sparse Packed Doubleword Integer Values from Dense Memory/Register"},
	{TVPEXPANDQ, "Load Sparse Packed Quadword Integer Values from Dense Memory/Register"},
	{TVPEXTRB, "Extract Byte"},
	{TVPEXTRD, "Extract Doubleword"},
	{TVPEXTRQ, "Extract Quadword"},
	{TVPEXTRW, "Extract Word"},
	{TVPGATHERDD, "Gather Packed Doubleword Values Using Signed Doubleword Indices"},
	{TVPGATHERDQ, "Gather Packed Quadword Values Using Signed Doubleword Indices"},
	{TVPGATHERQD, "Gather Packed Doubleword Values Using Signed Quadword Indices"},
	{TVPGATHERQQ, "Gather Packed Quadword Values Using Signed Quadword Indices"},
	{TVPHADDBD, "Packed Horizontal Add Signed Byte to Signed Doubleword"},
	{TVPHADDBQ, "Packed Horizontal Add Signed Byte to Signed Quadword"},
	{TVPHADDBW, "Packed Horizontal Add Signed Byte to Signed Word"},
	{TVPHADDD, "Packed Horizontal Add Doubleword Integer"},
	{TVPHADDDQ, "Packed Horizontal Add Signed Doubleword to Signed Quadword"},
	{TVPHADDSW, "Packed Horizontal Add Signed Word Integers with Signed Saturation"},
	{TVPHADDUBD, "Packed Horizontal Add Unsigned Byte to Doubleword"},
	{TVPHADDUBQ, "Packed Horizontal Add Unsigned Byte to Quadword"},
	{TVPHADDUBW, "Packed Horizontal Add Unsigned Byte to Word"},
	{TVPHADDUDQ, "Packed Horizontal Add Unsigned Doubleword to Quadword"},
	{TVPHADDUWD, "Packed Horizontal Add Unsigned Word to Doubleword"},
	{TVPHADDUWQ, "Packed Horizontal Add Unsigned Word to Quadword"},
	{TVPHADDW, "Packed Horizontal Add Word Integers"},
	{TVPHADDWD, "Packed Horizontal Add Signed Word to Signed Doubleword"},
	{TVPHADDWQ, "Packed Horizontal Add Signed Word to Signed Quadword"},
	{TVPHMINPOSUW, "Packed Horizontal Minimum of Unsigned Word Integers"},
	{TVPHSUBBW, "Packed Horizontal Subtract Signed Byte to Signed Word"},
	{TVPHSUBD, "Packed Horizontal Subtract Doubleword Integers"},
	{TVPHSUBDQ, "Packed Horizontal Subtract Signed Doubleword to Signed Quadword"},
	{TVPHSUBSW, "Packed Horizontal Subtract Signed Word Integers with Signed Saturation"},
	{TVPHSUBW, "Packed Horizontal Subtract Word Integers"},
	{TVPHSUBWD, "Packed Horizontal Subtract Signed Word to Signed Doubleword"},
	{TVPINSRB, "Insert Byte"},
	{TVPINSRD, "Insert Doubleword"},
	{TVPINSRQ, "Insert Quadword"},
	{TVPINSRW, "Insert Word"},
	{TVPLZCNTD, "Count the Number of Leading Zero Bits for Packed Doubleword Values"},
	{TVPLZCNTQ, "Count the Number of Leading Zero Bits for Packed Quadword Values"},
	{TVPMACSDD, "Packed Multiply Accumulate Signed Doubleword to Signed Doubleword"},
	{TVPMACSDQH, "Packed Multiply Accumulate Signed High Doubleword to Signed Quadword"},
	{TVPMACSDQL, "Packed Multiply Accumulate Signed Low Doubleword to Signed Quadword"},
	{TVPMACSSDD, "Packed Multiply Accumulate with Saturation Signed Doubleword to Signed Doubleword"},
	{TVPMACSSDQH, "Packed Multiply Accumulate with Saturation Signed High Doubleword to Signed Quadword"},
	{TVPMACSSDQL, "Packed Multiply Accumulate with Saturation Signed Low Doubleword to Signed Quadword"},
	{TVPMACSSWD, "Packed Multiply Accumulate with Saturation Signed Word to Signed Doubleword"},
	{TVPMACSSWW, "Packed Multiply Accumulate with Saturation Signed Word to Signed Word"},
	{TVPMACSWD, "Packed Multiply Accumulate Signed Word to Signed Doubleword"},
	{TVPMACSWW, "Packed Multiply Accumulate Signed Word to Signed Word"},
	{TVPMADCSSWD, "Packed Multiply Add Accumulate with Saturation Signed Word to Signed Doubleword"},
	{TVPMADCSWD, "Packed Multiply Add Accumulate Signed Word to Signed Doubleword"},
	{TVPMADD52HUQ, "Packed Multiply of Unsigned 52-bit Unsigned Integers and Add High 52-bit Products to Quadword Accumulators"},
	{TVPMADD52LUQ, "Packed Multiply of Unsigned 52-bit Integers and Add the Low 52-bit Products to Quadword Accumulators"},
	{TVPMADDUBSW, "Multiply and Add Packed Signed and Unsigned Byte Integers"},
	{TVPMADDWD, "Multiply and Add Packed Signed Word Integers"},
	{TVPMASKMOVD, "Conditional Move Packed Doubleword Integers"},
	{TVPMASKMOVQ, "Conditional Move Packed Quadword Integers"},
	{TVPMAXSB, "Maximum of Packed Signed Byte Integers"},
	{TVPMAXSD, "Maximum of Packed Signed Doubleword Integers"},
	{TVPMAXSQ, "Maximum of Packed Signed Quadword Integers"},
	{TVPMAXSW, "Maximum of Packed Signed Word Integers"},
	{TVPMAXUB, "Maximum of Packed Unsigned Byte Integers"},
	{TVPMAXUD, "Maximum of Packed Unsigned Doubleword Integers"},
	{TVPMAXUQ, "Maximum of Packed Unsigned Quadword Integers"},
	{TVPMAXUW, "Maximum of Packed Unsigned Word Integers"},
	{TVPMINSB, "Minimum of Packed Signed Byte Integers"},
	{TVPMINSD, "Minimum of Packed Signed Doubleword Integers"},
	{TVPMINSQ, "Minimum of Packed Signed Quadword Integers"},
	{TVPMINSW, "Minimum of Packed Signed Word Integers"},
	{TVPMINUB, "Minimum of Packed Unsigned Byte Integers"},
	{TVPMINUD, "Minimum of Packed Unsigned Doubleword Integers"},
	{TVPMINUQ, "Minimum of Packed Unsigned Quadword Integers"},
	{TVPMINUW, "Minimum of Packed Unsigned Word Integers"},
	{TVPMOVB2M, "Move Signs of Packed Byte Integers to Mask Register"},
	{TVPMOVD2M, "Move Signs of Packed Doubleword Integers to Mask Register"},
	{TVPMOVDB, "Down Convert Packed Doubleword Values to Byte Values with Truncation"},
	{TVPMOVDW, "Down Convert Packed Doubleword Values to Word Values with Truncation"},
	{TVPMOVM2B, "Expand Bits of Mask Register to Packed Byte Integers"},
	{TVPMOVM2D, "Expand Bits of Mask Register to Packed Doubleword Integers"},
	{TVPMOVM2Q, "Expand Bits of Mask Register to Packed Quadword Integers"},
	{TVPMOVM2W, "Expand Bits of Mask Register to Packed Word Integers"},
	{TVPMOVMSKB, "Move Byte Mask"},
	{TVPMOVQ2M, "Move Signs of Packed Quadword Integers to Mask Register"},
	{TVPMOVQB, "Down Convert Packed Quadword Values to Byte Values with Truncation"},
	{TVPMOVQD, "Down Convert Packed Quadword Values to Doubleword Values with Truncation"},
	{TVPMOVQW, "Down Convert Packed Quadword Values to Word Values with Truncation"},
	{TVPMOVSDB, "Down Convert Packed Doubleword Values to Byte Values with Signed Saturation"},
	{TVPMOVSDW, "Down Convert Packed Doubleword Values to Word Values with Signed Saturation"},
	{TVPMOVSQB, "Down Convert Packed Quadword Values to Byte Values with Signed Saturation"},
	{TVPMOVSQD, "Down Convert Packed Quadword Values to Doubleword Values with Signed Saturation"},
	{TVPMOVSQW, "Down Convert Packed Quadword Values to Word Values with Signed Saturation"},
	{TVPMOVSWB, "Down Convert Packed Word Values to Byte Values with Signed Saturation"},
	{TVPMOVSXBD, "Move Packed Byte Integers to Doubleword Integers with Sign Extension"},
	{TVPMOVSXBQ, "Move Packed Byte Integers to Quadword Integers with Sign Extension"},
	{TVPMOVSXBW, "Move Packed Byte Integers to Word Integers with Sign Extension"},
	{TVPMOVSXDQ, "Move Packed Doubleword Integers to Quadword Integers with Sign Extension"},
	{TVPMOVSXWD, "Move Packed Word Integers to Doubleword Integers with Sign Extension"},
	{TVPMOVSXWQ, "Move Packed Word Integers to Quadword Integers with Sign Extension"},
	{TVPMOVUSDB, "Down Convert Packed Doubleword Values to Byte Values with Unsigned Saturation"},
	{TVPMOVUSDW, "Down Convert Packed Doubleword Values to Word Values with Unsigned Saturation"},
	{TVPMOVUSQB, "Down Convert Packed Quadword Values to Byte Values with Unsigned Saturation"},
	{TVPMOVUSQD, "Down Convert Packed Quadword Values to Doubleword Values with Unsigned Saturation"},
	{TVPMOVUSQW, "Down Convert Packed Quadword Values to Word Values with Unsigned Saturation"},
	{TVPMOVUSWB, "Down Convert Packed Word Values to Byte Values with Unsigned Saturation"},
	{TVPMOVW2M, "Move Signs of Packed Word Integers to Mask Register"},
	{TVPMOVWB, "Down Convert Packed Word Values to Byte Values with Truncation"},
	{TVPMOVZXBD, "Move Packed Byte Integers to Doubleword Integers with Zero Extension"},
	{TVPMOVZXBQ, "Move Packed Byte Integers to Quadword Integers with Zero Extension"},
	{TVPMOVZXBW, "Move Packed Byte Integers to Word Integers with Zero Extension"},
	{TVPMOVZXDQ, "Move Packed Doubleword Integers to Quadword Integers with Zero Extension"},
	{TVPMOVZXWD, "Move Packed Word Integers to Doubleword Integers with Zero Extension"},
	{TVPMOVZXWQ, "Move Packed Word Integers to Quadword Integers with Zero Extension"},
	{TVPMULDQ, "Multiply Packed Signed Doubleword Integers and Store Quadword Result"},
	{TVPMULHRSW, "Packed Multiply Signed Word Integers and Store High Result with Round and Scale"},
	{TVPMULHUW, "Multiply Packed Unsigned Word Integers and Store High Result"},
	{TVPMULHW, "Multiply Packed Signed Word Integers and Store High Result"},
	{TVPMULLD, "Multiply Packed Signed Doubleword Integers and Store Low Result"},
	{TVPMULLQ, "Multiply Packed Signed Quadword Integers and Store Low Result"},
	{TVPMULLW, "Multiply Packed Signed Word Integers and Store Low Result"},
	{TVPMULTISHIFTQB, "Select Packed Unaligned Bytes from Quadword Sources"},
	{TVPMULUDQ, "Multiply Packed Unsigned Doubleword Integers"},
	{TVPOR, "Packed Bitwise Logical OR"},
	{TVPORD, "Bitwise Logical OR of Packed Doubleword Integers"},
	{TVPORQ, "Bitwise Logical OR of Packed Quadword Integers"},
	{TVPPERM, "Packed Permute Bytes"},
	{TVPROLD, "Rotate Packed Doubleword Left"},
	{TVPROLQ, "Rotate Packed Quadword Left"},
	{TVPROLVD, "Variable Rotate Packed Doubleword Left"},
	{TVPROLVQ, "Variable Rotate Packed Quadword Left"},
	{TVPRORD, "Rotate Packed Doubleword Right"},
	{TVPRORQ, "Rotate Packed Quadword Right"},
	{TVPRORVD, "Variable Rotate Packed Doubleword Right"},
	{TVPRORVQ, "Variable Rotate Packed Quadword Right"},
	{TVPROTB, "Packed Rotate Bytes"},
	{TVPROTD, "Packed Rotate Doublewords"},
	{TVPROTQ, "Packed Rotate Quadwords"},
	{TVPROTW, "Packed Rotate Words"},
	{TVPSADBW, "Compute Sum of Absolute Differences"},
	{TVPSCATTERDD, "Scatter Packed Doubleword Values with Signed Doubleword Indices"},
	{TVPSCATTERDQ, "Scatter Packed Quadword Values with Signed Doubleword Indices"},
	{TVPSCATTERQD, "Scatter Packed Doubleword Values with Signed Quadword Indices"},
	{TVPSCATTERQQ, "Scatter Packed Quadword Values with Signed Quadword Indices"},
	{TVPSHAB, "Packed Shift Arithmetic Bytes"},
	{TVPSHAD, "Packed Shift Arithmetic Doublewords"},
	{TVPSHAQ, "Packed Shift Arithmetic Quadwords"},
	{TVPSHAW, "Packed Shift Arithmetic Words"},
	{TVPSHLB, "Packed Shift Logical Bytes"},
	{TVPSHLD, "Packed Shift Logical Doublewords"},
	{TVPSHLQ, "Packed Shift Logical Quadwords"},
	{TVPSHLW, "Packed Shift Logical Words"},
	{TVPSHUFB, "Packed Shuffle Bytes"},
	{TVPSHUFD, "Shuffle Packed Doublewords"},
	{TVPSHUFHW, "Shuffle Packed High Words"},
	{TVPSHUFLW, "Shuffle Packed Low Words"},
	{TVPSIGNB, "Packed Sign of Byte Integers"},
	{TVPSIGND, "Packed Sign of Doubleword Integers"},
	{TVPSIGNW, "Packed Sign of Word Integers"},
	{TVPSLLD, "Shift Packed Doubleword Data Left Logical"},
	{TVPSLLDQ, "Shift Packed Double Quadword Left Logical"},
	{TVPSLLQ, "Shift Packed Quadword Data Left Logical"},
	{TVPSLLVD, "Variable Shift Packed Doubleword Data Left Logical"},
	{TVPSLLVQ, "Variable Shift Packed Quadword Data Left Logical"},
	{TVPSLLVW, "Variable Shift Packed Word Data Left Logical"},
	{TVPSLLW, "Shift Packed Word Data Left Logical"},
	{TVPSRAD, "Shift Packed Doubleword Data Right Arithmetic"},
	{TVPSRAQ, "Shift Packed Quadword Data Right Arithmetic"},
	{TVPSRAVD, "Variable Shift Packed Doubleword Data Right Arithmetic"},
	{TVPSRAVQ, "Variable Shift Packed Quadword Data Right Arithmetic"},
	{TVPSRAVW, "Variable Shift Packed Word Data Right Arithmetic"},
	{TVPSRAW, "Shift Packed Word Data Right Arithmetic"},
	{TVPSRLD, "Shift Packed Doubleword Data Right Logical"},
	{TVPSRLDQ, "Shift Packed Double Quadword Right Logical"},
	{TVPSRLQ, "Shift Packed Quadword Data Right Logical"},
	{TVPSRLVD, "Variable Shift Packed Doubleword Data Right Logical"},
	{TVPSRLVQ, "Variable Shift Packed Quadword Data Right Logical"},
	{TVPSRLVW, "Variable Shift Packed Word Data Right Logical"},
	{TVPSRLW, "Shift Packed Word Data Right Logical"},
	{TVPSUBB, "Subtract Packed Byte Integers"},
	{TVPSUBD, "Subtract Packed Doubleword Integers"},
	{TVPSUBQ, "Subtract Packed Quadword Integers"},
	{TVPSUBSB, "Subtract Packed Signed Byte Integers with Signed Saturation"},
	{TVPSUBSW, "Subtract Packed Signed Word Integers with Signed Saturation"},
	{TVPSUBUSB, "Subtract Packed Unsigned Byte Integers with Unsigned Saturation"},
	{TVPSUBUSW, "Subtract Packed Unsigned Word Integers with Unsigned Saturation"},
	{TVPSUBW, "Subtract Packed Word Integers"},
	{TVPTERNLOGD, "Bitwise Ternary Logical Operation on Doubleword Values"},
	{TVPTERNLOGQ, "Bitwise Ternary Logical Operation on Quadword Values"},
	{TVPTEST, "Packed Logical Compare"},
	{TVPTESTMB, "Logical AND of Packed Byte Integer Values and Set Mask"},
	{TVPTESTMD, "Logical AND of Packed Doubleword Integer Values and Set Mask"},
	{TVPTESTMQ, "Logical AND of Packed Quadword Integer Values and Set Mask"},
	{TVPTESTMW, "Logical AND of Packed Word Integer Values and Set Mask"},
	{TVPTESTNMB, "Logical NAND of Packed Byte Integer Values and Set Mask"},
	{TVPTESTNMD, "Logical NAND of Packed Doubleword Integer Values and Set Mask"},
	{TVPTESTNMQ, "Logical NAND of Packed Quadword Integer Values and Set Mask"},
	{TVPTESTNMW, "Logical NAND of Packed Word Integer Values and Set Mask"},
	{TVPUNPCKHBW, "Unpack and Interleave High-Order Bytes into Words"},
	{TVPUNPCKHDQ, "Unpack and Interleave High-Order Doublewords into Quadwords"},
	{TVPUNPCKHQDQ, "Unpack and Interleave High-Order Quadwords into Double Quadwords"},
	{TVPUNPCKHWD, "Unpack and Interleave High-Order Words into Doublewords"},
	{TVPUNPCKLBW, "Unpack and Interleave Low-Order Bytes into Words"},
	{TVPUNPCKLDQ, "Unpack and Interleave Low-Order Doublewords into Quadwords"},
	{TVPUNPCKLQDQ, "Unpack and Interleave Low-Order Quadwords into Double Quadwords"},
	{TVPUNPCKLWD, "Unpack and Interleave Low-Order Words into Doublewords"},
	{TVPXOR, "Packed Bitwise Logical Exclusive OR"},
	{TVPXORD, "Bitwise Logical Exclusive OR of Packed Doubleword Integers"},
	{TVPXORQ, "Bitwise Logical Exclusive OR of Packed Quadword Integers"},
	{TVRANGEPD, "Range Restriction Calculation For Packed Pairs of Double-Precision Floating-Point Values"},
	{TVRANGEPS, "Range Restriction Calculation For Packed Pairs of Single-Precision Floating-Point Values"},
	{TVRANGESD, "Range Restriction Calculation For a pair of Scalar Double-Precision Floating-Point Values"},
	{TVRANGESS, "Range Restriction Calculation For a pair of Scalar Single-Precision Floating-Point Values"},
	{TVRCP14PD, "Compute Approximate Reciprocals of Packed Double-Precision Floating-Point Values"},
	{TVRCP14PS, "Compute Approximate Reciprocals of Packed Single-Precision Floating-Point Values"},
	{TVRCP14SD, "Compute Approximate Reciprocal of a Scalar Double-Precision Floating-Point Value"},
	{TVRCP14SS, "Compute Approximate Reciprocal of a Scalar Single-Precision Floating-Point Value"},
	{TVRCP28PD, "Approximation to the Reciprocal of Packed Double-Precision Floating-Point Values with Less Than 2^-28 Relative Error"},
	{TVRCP28PS, "Approximation to the Reciprocal of Packed Single-Precision Floating-Point Values with Less Than 2^-28 Relative Error"},
	{TVRCP28SD, "Approximation to the Reciprocal of a Scalar Double-Precision Floating-Point Value with Less Than 2^-28 Relative Error"},
	{TVRCP28SS, "Approximation to the Reciprocal of a Scalar Single-Precision Floating-Point Value with Less Than 2^-28 Relative Error"},
	{TVRCPPS, "Compute Approximate Reciprocals of Packed Single-Precision Floating-Point Values"},
	{TVRCPSS, "Compute Approximate Reciprocal of Scalar Single-Precision Floating-Point Values"},
	{TVREDUCEPD, "Perform Reduction Transformation on Packed Double-Precision Floating-Point Values"},
	{TVREDUCEPS, "Perform Reduction Transformation on Packed Single-Precision Floating-Point Values"},
	{TVREDUCESD, "Perform Reduction Transformation on a Scalar Double-Precision Floating-Point Value"},
	{TVREDUCESS, "Perform Reduction Transformation on a Scalar Single-Precision Floating-Point Value"},
	{TVRNDSCALEPD, "Round Packed Double-Precision Floating-Point Values To Include A Given Number Of Fraction Bits"},
	{TVRNDSCALEPS, "Round Packed Single-Precision Floating-Point Values To Include A Given Number Of Fraction Bits"},
	{TVRNDSCALESD, "Round Scalar Double-Precision Floating-Point Value To Include A Given Number Of Fraction Bits"},
	{TVRNDSCALESS, "Round Scalar Single-Precision Floating-Point Value To Include A Given Number Of Fraction Bits"},
	{TVROUNDPD, "Round Packed Double Precision Floating-Point Values"},
	{TVROUNDPS, "Round Packed Single Precision Floating-Point Values"},
	{TVROUNDSD, "Round Scalar Double Precision Floating-Point Values"},
	{TVROUNDSS, "Round Scalar Single Precision Floating-Point Values"},
	{TVRSQRT14PD, "Compute Approximate Reciprocals of Square Roots of Packed Double-Precision Floating-Point Values"},
	{TVRSQRT14PS, "Compute Approximate Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values"},
	{TVRSQRT14SD, "Compute Approximate Reciprocal of a Square Root of a Scalar Double-Precision Floating-Point Value"},
	{TVRSQRT14SS, "Compute Approximate Reciprocal of a Square Root of a Scalar Single-Precision Floating-Point Value"},
	{TVRSQRT28PD, "Approximation to the Reciprocal Square Root of Packed Double-Precision Floating-Point Values with Less Than 2^-28 Relative Error"},
	{TVRSQRT28PS, "Approximation to the Reciprocal Square Root of Packed Single-Precision Floating-Point Values with Less Than 2^-28 Relative Error"},
	{TVRSQRT28SD, "Approximation to the Reciprocal Square Root of a Scalar Double-Precision Floating-Point Value with Less Than 2^-28 Relative Error"},
	{TVRSQRT28SS, "Approximation to the Reciprocal Square Root of a Scalar Single-Precision Floating-Point Value with Less Than 2^-28 Relative Error"},
	{TVRSQRTPS, "Compute Reciprocals of Square Roots of Packed Single-Precision Floating-Point Values"},
	{TVRSQRTSS, "Compute Reciprocal of Square Root of Scalar Single-Precision Floating-Point Value"},
	{TVSCALEFPD, "Scale Packed Double-Precision Floating-Point Values With Double-Precision Floating-Point Values"},
	{TVSCALEFPS, "Scale Packed Single-Precision Floating-Point Values With Single-Precision Floating-Point Values"},
	{TVSCALEFSD, "Scale Scalar Double-Precision Floating-Point Value With a Double-Precision Floating-Point Value"},
	{TVSCALEFSS, "Scale Scalar Single-Precision Floating-Point Value With a Single-Precision Floating-Point Value"},
	{TVSCATTERDPD, "Scatter Packed Double-Precision Floating-Point Values with Signed Doubleword Indices"},
	{TVSCATTERDPS, "Scatter Packed Single-Precision Floating-Point Values with Signed Doubleword Indices"},
	{TVSCATTERPF0DPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Doubleword Indices Using T0 Hint with Intent to Write"},
	{TVSCATTERPF0DPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Doubleword Indices Using T0 Hint with Intent to Write"},
	{TVSCATTERPF0QPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Quadword Indices Using T0 Hint with Intent to Write"},
	{TVSCATTERPF0QPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Quadword Indices Using T0 Hint with Intent to Write"},
	{TVSCATTERPF1DPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Doubleword Indices Using T1 Hint with Intent to Write"},
	{TVSCATTERPF1DPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Doubleword Indices Using T1 Hint with Intent to Write"},
	{TVSCATTERPF1QPD, "Sparse Prefetch Packed Double-Precision Floating-Point Data Values with Signed Quadword Indices Using T1 Hint with Intent to Write"},
	{TVSCATTERPF1QPS, "Sparse Prefetch Packed Single-Precision Floating-Point Data Values with Signed Quadword Indices Using T1 Hint with Intent to Write"},
	{TVSCATTERQPD, "Scatter Packed Double-Precision Floating-Point Values with Signed Quadword Indices"},
	{TVSCATTERQPS, "Scatter Packed Single-Precision Floating-Point Values with Signed Quadword Indices"},
	{TVSHUFF32X4, "Shuffle 128-Bit Packed Single-Precision Floating-Point Values"},
	{TVSHUFF64X2, "Shuffle 128-Bit Packed Double-Precision Floating-Point Values"},
	{TVSHUFI32X4, "Shuffle 128-Bit Packed Doubleword Integer Values"},
	{TVSHUFI64X2, "Shuffle 128-Bit Packed Quadword Integer Values"},
	{TVSHUFPD, "Shuffle Packed Double-Precision Floating-Point Values"},
	{TVSHUFPS, "Shuffle Packed Single-Precision Floating-Point Values"},
	{TVSQRTPD, "Compute Square Roots of Packed Double-Precision Floating-Point Values"},
	{TVSQRTPS, "Compute Square Roots of Packed Single-Precision Floating-Point Values"},
	{TVSQRTSD, "Compute Square Root of Scalar Double-Precision Floating-Point Value"},
	{TVSQRTSS, "Compute Square Root of Scalar Single-Precision Floating-Point Value"},
	{TVSTMXCSR, "Store MXCSR Register State"},
	{TVSUBPD, "Subtract Packed Double-Precision Floating-Point Values"},
	{TVSUBPS, "Subtract Packed Single-Precision Floating-Point Values"},
	{TVSUBSD, "Subtract Scalar Double-Precision Floating-Point Values"},
	{TVSUBSS, "Subtract Scalar Single-Precision Floating-Point Values"},
	{TVTESTPD, "Packed Double-Precision Floating-Point Bit Test"},
	{TVTESTPS, "Packed Single-Precision Floating-Point Bit Test"},
	{TVUCOMISD, "Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS"},
	{TVUCOMISS, "Unordered Compare Scalar Single-Precision Floating-Point Values and Set EFLAGS"},
	{TVUNPCKHPD, "Unpack and Interleave High Packed Double-Precision Floating-Point Values"},
	{TVUNPCKHPS, "Unpack and Interleave High Packed Single-Precision Floating-Point Values"},
	{TVUNPCKLPD, "Unpack and Interleave Low Packed Double-Precision Floating-Point Values"},
	{TVUNPCKLPS, "Unpack and Interleave Low Packed Single-Precision Floating-Point Values"},
	{TVXORPD, "Bitwise Logical XOR for Double-Precision Floating-Point Values"},
	{TVXORPS, "Bitwise Logical XOR for Single-Precision Floating-Point Values"},
	{TVZEROALL, "Zero All YMM Registers"},
	{TVZEROUPPER, "Zero Upper Bits of YMM Registers"},
	{TXADD, "Exchange and Add"},
	{TXCHG, "Exchange Register/Memory with Register"},
	{TXGETBV, "Get Value of Extended Control Register"},
	{TXLATB, "Table Look-up Translation"},
	{TXOR, "Logical Exclusive OR"},
	{TXORPD, "Bitwise Logical XOR for Double-Precision Floating-Point Values"},
	{TXORPS, "Bitwise Logical XOR for Single-Precision Floating-Point Values"}}

func GetInstrType(typ int) InstrType {
	for _, instr := range instrTypes {
		if instr.typ == typ {
			return instr
		}
	}
	panic(fmt.Sprintf("Couldn't find matching InstrType in GetInstrType, typ:%v", typ))
}

func (instr InstrType) String() string {
	switch instr.typ {
	default:
		return "Unknown Instr!!"
	case TADC:
		return "ADC"
	case TADCX:
		return "ADCX"
	case TADD:
		return "ADD"
	case TADDPD:
		return "ADDPD"
	case TADDPS:
		return "ADDPS"
	case TADDSD:
		return "ADDSD"
	case TADDSS:
		return "ADDSS"
	case TADDSUBPD:
		return "ADDSUBPD"
	case TADDSUBPS:
		return "ADDSUBPS"
	case TADOX:
		return "ADOX"
	case TAESDEC:
		return "AESDEC"
	case TAESDECLAST:
		return "AESDECLAST"
	case TAESENC:
		return "AESENC"
	case TAESENCLAST:
		return "AESENCLAST"
	case TAESIMC:
		return "AESIMC"
	case TAESKEYGENASSIST:
		return "AESKEYGENASSIST"
	case TAND:
		return "AND"
	case TANDN:
		return "ANDN"
	case TANDNPD:
		return "ANDNPD"
	case TANDNPS:
		return "ANDNPS"
	case TANDPD:
		return "ANDPD"
	case TANDPS:
		return "ANDPS"
	case TBEXTR:
		return "BEXTR"
	case TBLCFILL:
		return "BLCFILL"
	case TBLCI:
		return "BLCI"
	case TBLCIC:
		return "BLCIC"
	case TBLCMSK:
		return "BLCMSK"
	case TBLCS:
		return "BLCS"
	case TBLENDPD:
		return "BLENDPD"
	case TBLENDPS:
		return "BLENDPS"
	case TBLENDVPD:
		return "BLENDVPD"
	case TBLENDVPS:
		return "BLENDVPS"
	case TBLSFILL:
		return "BLSFILL"
	case TBLSI:
		return "BLSI"
	case TBLSIC:
		return "BLSIC"
	case TBLSMSK:
		return "BLSMSK"
	case TBLSR:
		return "BLSR"
	case TBSF:
		return "BSF"
	case TBSR:
		return "BSR"
	case TBSWAP:
		return "BSWAP"
	case TBT:
		return "BT"
	case TBTC:
		return "BTC"
	case TBTR:
		return "BTR"
	case TBTS:
		return "BTS"
	case TBZHI:
		return "BZHI"
	case TCALL:
		return "CALL"
	case TCBW:
		return "CBW"
	case TCDQ:
		return "CDQ"
	case TCDQE:
		return "CDQE"
	case TCLC:
		return "CLC"
	case TCLD:
		return "CLD"
	case TCMC:
		return "CMC"
	case TCMOVA:
		return "CMOVA"
	case TCMOVAE:
		return "CMOVAE"
	case TCMOVB:
		return "CMOVB"
	case TCMOVBE:
		return "CMOVBE"
	case TCMOVC:
		return "CMOVC"
	case TCMOVE:
		return "CMOVE"
	case TCMOVG:
		return "CMOVG"
	case TCMOVGE:
		return "CMOVGE"
	case TCMOVL:
		return "CMOVL"
	case TCMOVLE:
		return "CMOVLE"
	case TCMOVNA:
		return "CMOVNA"
	case TCMOVNAE:
		return "CMOVNAE"
	case TCMOVNB:
		return "CMOVNB"
	case TCMOVNBE:
		return "CMOVNBE"
	case TCMOVNC:
		return "CMOVNC"
	case TCMOVNE:
		return "CMOVNE"
	case TCMOVNG:
		return "CMOVNG"
	case TCMOVNGE:
		return "CMOVNGE"
	case TCMOVNL:
		return "CMOVNL"
	case TCMOVNLE:
		return "CMOVNLE"
	case TCMOVNO:
		return "CMOVNO"
	case TCMOVNP:
		return "CMOVNP"
	case TCMOVNS:
		return "CMOVNS"
	case TCMOVNZ:
		return "CMOVNZ"
	case TCMOVO:
		return "CMOVO"
	case TCMOVP:
		return "CMOVP"
	case TCMOVPE:
		return "CMOVPE"
	case TCMOVPO:
		return "CMOVPO"
	case TCMOVS:
		return "CMOVS"
	case TCMOVZ:
		return "CMOVZ"
	case TCMP:
		return "CMP"
	case TCMPPD:
		return "CMPPD"
	case TCMPPS:
		return "CMPPS"
	case TCMPSD:
		return "CMPSD"
	case TCMPSS:
		return "CMPSS"
	case TCMPXCHG:
		return "CMPXCHG"
	case TCMPXCHG16B:
		return "CMPXCHG16B"
	case TCMPXCHG8B:
		return "CMPXCHG8B"
	case TCOMISD:
		return "COMISD"
	case TCOMISS:
		return "COMISS"
	case TCPUID:
		return "CPUID"
	case TCQO:
		return "CQO"
	case TCRC32:
		return "CRC32"
	case TCVTDQ2PD:
		return "CVTDQ2PD"
	case TCVTDQ2PS:
		return "CVTDQ2PS"
	case TCVTPD2DQ:
		return "CVTPD2DQ"
	case TCVTPD2PI:
		return "CVTPD2PI"
	case TCVTPD2PS:
		return "CVTPD2PS"
	case TCVTPI2PD:
		return "CVTPI2PD"
	case TCVTPI2PS:
		return "CVTPI2PS"
	case TCVTPS2DQ:
		return "CVTPS2DQ"
	case TCVTPS2PD:
		return "CVTPS2PD"
	case TCVTPS2PI:
		return "CVTPS2PI"
	case TCVTSD2SI:
		return "CVTSD2SI"
	case TCVTSD2SS:
		return "CVTSD2SS"
	case TCVTSI2SD:
		return "CVTSI2SD"
	case TCVTSI2SS:
		return "CVTSI2SS"
	case TCVTSS2SD:
		return "CVTSS2SD"
	case TCVTSS2SI:
		return "CVTSS2SI"
	case TCVTTPD2DQ:
		return "CVTTPD2DQ"
	case TCVTTPD2PI:
		return "CVTTPD2PI"
	case TCVTTPS2DQ:
		return "CVTTPS2DQ"
	case TCVTTPS2PI:
		return "CVTTPS2PI"
	case TCVTTSD2SI:
		return "CVTTSD2SI"
	case TCVTTSS2SI:
		return "CVTTSS2SI"
	case TCWD:
		return "CWD"
	case TCWDE:
		return "CWDE"
	case TDEC:
		return "DEC"
	case TDIV:
		return "DIV"
	case TDIVPD:
		return "DIVPD"
	case TDIVPS:
		return "DIVPS"
	case TDIVSD:
		return "DIVSD"
	case TDIVSS:
		return "DIVSS"
	case TDPPD:
		return "DPPD"
	case TDPPS:
		return "DPPS"
	case TEMMS:
		return "EMMS"
	case TEXTRACTPS:
		return "EXTRACTPS"
	case TEXTRQ:
		return "EXTRQ"
	case TFEMMS:
		return "FEMMS"
	case THADDPD:
		return "HADDPD"
	case THADDPS:
		return "HADDPS"
	case THSUBPD:
		return "HSUBPD"
	case THSUBPS:
		return "HSUBPS"
	case TIDIV:
		return "IDIV"
	case TIMUL:
		return "IMUL"
	case TINC:
		return "INC"
	case TINSERTPS:
		return "INSERTPS"
	case TINSERTQ:
		return "INSERTQ"
	case TINT:
		return "INT"
	case TJA:
		return "JA"
	case TJAE:
		return "JAE"
	case TJB:
		return "JB"
	case TJBE:
		return "JBE"
	case TJC:
		return "JC"
	case TJE:
		return "JE"
	case TJECXZ:
		return "JECXZ"
	case TJG:
		return "JG"
	case TJGE:
		return "JGE"
	case TJL:
		return "JL"
	case TJLE:
		return "JLE"
	case TJMP:
		return "JMP"
	case TJNA:
		return "JNA"
	case TJNAE:
		return "JNAE"
	case TJNB:
		return "JNB"
	case TJNBE:
		return "JNBE"
	case TJNC:
		return "JNC"
	case TJNE:
		return "JNE"
	case TJNG:
		return "JNG"
	case TJNGE:
		return "JNGE"
	case TJNL:
		return "JNL"
	case TJNLE:
		return "JNLE"
	case TJNO:
		return "JNO"
	case TJNP:
		return "JNP"
	case TJNS:
		return "JNS"
	case TJNZ:
		return "JNZ"
	case TJO:
		return "JO"
	case TJP:
		return "JP"
	case TJPE:
		return "JPE"
	case TJPO:
		return "JPO"
	case TJRCXZ:
		return "JRCXZ"
	case TJS:
		return "JS"
	case TJZ:
		return "JZ"
	case TKADDB:
		return "KADDB"
	case TKADDD:
		return "KADDD"
	case TKADDQ:
		return "KADDQ"
	case TKADDW:
		return "KADDW"
	case TKANDB:
		return "KANDB"
	case TKANDD:
		return "KANDD"
	case TKANDNB:
		return "KANDNB"
	case TKANDND:
		return "KANDND"
	case TKANDNQ:
		return "KANDNQ"
	case TKANDNW:
		return "KANDNW"
	case TKANDQ:
		return "KANDQ"
	case TKANDW:
		return "KANDW"
	case TKMOVB:
		return "KMOVB"
	case TKMOVD:
		return "KMOVD"
	case TKMOVQ:
		return "KMOVQ"
	case TKMOVW:
		return "KMOVW"
	case TKNOTB:
		return "KNOTB"
	case TKNOTD:
		return "KNOTD"
	case TKNOTQ:
		return "KNOTQ"
	case TKNOTW:
		return "KNOTW"
	case TKORB:
		return "KORB"
	case TKORD:
		return "KORD"
	case TKORQ:
		return "KORQ"
	case TKORTESTB:
		return "KORTESTB"
	case TKORTESTD:
		return "KORTESTD"
	case TKORTESTQ:
		return "KORTESTQ"
	case TKORTESTW:
		return "KORTESTW"
	case TKORW:
		return "KORW"
	case TKSHIFTLB:
		return "KSHIFTLB"
	case TKSHIFTLD:
		return "KSHIFTLD"
	case TKSHIFTLQ:
		return "KSHIFTLQ"
	case TKSHIFTLW:
		return "KSHIFTLW"
	case TKSHIFTRB:
		return "KSHIFTRB"
	case TKSHIFTRD:
		return "KSHIFTRD"
	case TKSHIFTRQ:
		return "KSHIFTRQ"
	case TKSHIFTRW:
		return "KSHIFTRW"
	case TKTESTB:
		return "KTESTB"
	case TKTESTD:
		return "KTESTD"
	case TKTESTQ:
		return "KTESTQ"
	case TKTESTW:
		return "KTESTW"
	case TKUNPCKBW:
		return "KUNPCKBW"
	case TKUNPCKDQ:
		return "KUNPCKDQ"
	case TKUNPCKWD:
		return "KUNPCKWD"
	case TKXNORB:
		return "KXNORB"
	case TKXNORD:
		return "KXNORD"
	case TKXNORQ:
		return "KXNORQ"
	case TKXNORW:
		return "KXNORW"
	case TKXORB:
		return "KXORB"
	case TKXORD:
		return "KXORD"
	case TKXORQ:
		return "KXORQ"
	case TKXORW:
		return "KXORW"
	case TLDDQU:
		return "LDDQU"
	case TLDMXCSR:
		return "LDMXCSR"
	case TLEA:
		return "LEA"
	case TLFENCE:
		return "LFENCE"
	case TLZCNT:
		return "LZCNT"
	case TMASKMOVDQU:
		return "MASKMOVDQU"
	case TMASKMOVQ:
		return "MASKMOVQ"
	case TMAXPD:
		return "MAXPD"
	case TMAXPS:
		return "MAXPS"
	case TMAXSD:
		return "MAXSD"
	case TMAXSS:
		return "MAXSS"
	case TMFENCE:
		return "MFENCE"
	case TMINPD:
		return "MINPD"
	case TMINPS:
		return "MINPS"
	case TMINSD:
		return "MINSD"
	case TMINSS:
		return "MINSS"
	case TMOV:
		return "MOV"
	case TMOVAPD:
		return "MOVAPD"
	case TMOVAPS:
		return "MOVAPS"
	case TMOVBE:
		return "MOVBE"
	case TMOVD:
		return "MOVD"
	case TMOVDDUP:
		return "MOVDDUP"
	case TMOVDQ2Q:
		return "MOVDQ2Q"
	case TMOVDQA:
		return "MOVDQA"
	case TMOVDQU:
		return "MOVDQU"
	case TMOVHLPS:
		return "MOVHLPS"
	case TMOVHPD:
		return "MOVHPD"
	case TMOVHPS:
		return "MOVHPS"
	case TMOVLHPS:
		return "MOVLHPS"
	case TMOVLPD:
		return "MOVLPD"
	case TMOVLPS:
		return "MOVLPS"
	case TMOVMSKPD:
		return "MOVMSKPD"
	case TMOVMSKPS:
		return "MOVMSKPS"
	case TMOVNTDQ:
		return "MOVNTDQ"
	case TMOVNTDQA:
		return "MOVNTDQA"
	case TMOVNTI:
		return "MOVNTI"
	case TMOVNTPD:
		return "MOVNTPD"
	case TMOVNTPS:
		return "MOVNTPS"
	case TMOVNTQ:
		return "MOVNTQ"
	case TMOVNTSD:
		return "MOVNTSD"
	case TMOVNTSS:
		return "MOVNTSS"
	case TMOVQ:
		return "MOVQ"
	case TMOVQ2DQ:
		return "MOVQ2DQ"
	case TMOVSD:
		return "MOVSD"
	case TMOVSHDUP:
		return "MOVSHDUP"
	case TMOVSLDUP:
		return "MOVSLDUP"
	case TMOVSS:
		return "MOVSS"
	case TMOVSX:
		return "MOVSX"
	case TMOVSXD:
		return "MOVSXD"
	case TMOVUPD:
		return "MOVUPD"
	case TMOVUPS:
		return "MOVUPS"
	case TMOVZX:
		return "MOVZX"
	case TMPSADBW:
		return "MPSADBW"
	case TMUL:
		return "MUL"
	case TMULPD:
		return "MULPD"
	case TMULPS:
		return "MULPS"
	case TMULSD:
		return "MULSD"
	case TMULSS:
		return "MULSS"
	case TMULX:
		return "MULX"
	case TNEG:
		return "NEG"
	case TNOP:
		return "NOP"
	case TNOT:
		return "NOT"
	case TOR:
		return "OR"
	case TORPD:
		return "ORPD"
	case TORPS:
		return "ORPS"
	case TPABSB:
		return "PABSB"
	case TPABSD:
		return "PABSD"
	case TPABSW:
		return "PABSW"
	case TPACKSSDW:
		return "PACKSSDW"
	case TPACKSSWB:
		return "PACKSSWB"
	case TPACKUSDW:
		return "PACKUSDW"
	case TPACKUSWB:
		return "PACKUSWB"
	case TPADDB:
		return "PADDB"
	case TPADDD:
		return "PADDD"
	case TPADDQ:
		return "PADDQ"
	case TPADDSB:
		return "PADDSB"
	case TPADDSW:
		return "PADDSW"
	case TPADDUSB:
		return "PADDUSB"
	case TPADDUSW:
		return "PADDUSW"
	case TPADDW:
		return "PADDW"
	case TPALIGNR:
		return "PALIGNR"
	case TPAND:
		return "PAND"
	case TPANDN:
		return "PANDN"
	case TPAUSE:
		return "PAUSE"
	case TPAVGB:
		return "PAVGB"
	case TPAVGUSB:
		return "PAVGUSB"
	case TPAVGW:
		return "PAVGW"
	case TPBLENDVB:
		return "PBLENDVB"
	case TPBLENDW:
		return "PBLENDW"
	case TPCLMULQDQ:
		return "PCLMULQDQ"
	case TPCMPEQB:
		return "PCMPEQB"
	case TPCMPEQD:
		return "PCMPEQD"
	case TPCMPEQQ:
		return "PCMPEQQ"
	case TPCMPEQW:
		return "PCMPEQW"
	case TPCMPESTRI:
		return "PCMPESTRI"
	case TPCMPESTRM:
		return "PCMPESTRM"
	case TPCMPGTB:
		return "PCMPGTB"
	case TPCMPGTD:
		return "PCMPGTD"
	case TPCMPGTQ:
		return "PCMPGTQ"
	case TPCMPGTW:
		return "PCMPGTW"
	case TPCMPISTRI:
		return "PCMPISTRI"
	case TPCMPISTRM:
		return "PCMPISTRM"
	case TPDEP:
		return "PDEP"
	case TPEXT:
		return "PEXT"
	case TPEXTRB:
		return "PEXTRB"
	case TPEXTRD:
		return "PEXTRD"
	case TPEXTRQ:
		return "PEXTRQ"
	case TPEXTRW:
		return "PEXTRW"
	case TPF2ID:
		return "PF2ID"
	case TPF2IW:
		return "PF2IW"
	case TPFACC:
		return "PFACC"
	case TPFADD:
		return "PFADD"
	case TPFCMPEQ:
		return "PFCMPEQ"
	case TPFCMPGE:
		return "PFCMPGE"
	case TPFCMPGT:
		return "PFCMPGT"
	case TPFMAX:
		return "PFMAX"
	case TPFMIN:
		return "PFMIN"
	case TPFMUL:
		return "PFMUL"
	case TPFNACC:
		return "PFNACC"
	case TPFPNACC:
		return "PFPNACC"
	case TPFRCP:
		return "PFRCP"
	case TPFRCPIT1:
		return "PFRCPIT1"
	case TPFRCPIT2:
		return "PFRCPIT2"
	case TPFRSQIT1:
		return "PFRSQIT1"
	case TPFRSQRT:
		return "PFRSQRT"
	case TPFSUB:
		return "PFSUB"
	case TPFSUBR:
		return "PFSUBR"
	case TPHADDD:
		return "PHADDD"
	case TPHADDSW:
		return "PHADDSW"
	case TPHADDW:
		return "PHADDW"
	case TPHMINPOSUW:
		return "PHMINPOSUW"
	case TPHSUBD:
		return "PHSUBD"
	case TPHSUBSW:
		return "PHSUBSW"
	case TPHSUBW:
		return "PHSUBW"
	case TPI2FD:
		return "PI2FD"
	case TPI2FW:
		return "PI2FW"
	case TPINSRB:
		return "PINSRB"
	case TPINSRD:
		return "PINSRD"
	case TPINSRQ:
		return "PINSRQ"
	case TPINSRW:
		return "PINSRW"
	case TPMADDUBSW:
		return "PMADDUBSW"
	case TPMADDWD:
		return "PMADDWD"
	case TPMAXSB:
		return "PMAXSB"
	case TPMAXSD:
		return "PMAXSD"
	case TPMAXSW:
		return "PMAXSW"
	case TPMAXUB:
		return "PMAXUB"
	case TPMAXUD:
		return "PMAXUD"
	case TPMAXUW:
		return "PMAXUW"
	case TPMINSB:
		return "PMINSB"
	case TPMINSD:
		return "PMINSD"
	case TPMINSW:
		return "PMINSW"
	case TPMINUB:
		return "PMINUB"
	case TPMINUD:
		return "PMINUD"
	case TPMINUW:
		return "PMINUW"
	case TPMOVMSKB:
		return "PMOVMSKB"
	case TPMOVSXBD:
		return "PMOVSXBD"
	case TPMOVSXBQ:
		return "PMOVSXBQ"
	case TPMOVSXBW:
		return "PMOVSXBW"
	case TPMOVSXDQ:
		return "PMOVSXDQ"
	case TPMOVSXWD:
		return "PMOVSXWD"
	case TPMOVSXWQ:
		return "PMOVSXWQ"
	case TPMOVZXBD:
		return "PMOVZXBD"
	case TPMOVZXBQ:
		return "PMOVZXBQ"
	case TPMOVZXBW:
		return "PMOVZXBW"
	case TPMOVZXDQ:
		return "PMOVZXDQ"
	case TPMOVZXWD:
		return "PMOVZXWD"
	case TPMOVZXWQ:
		return "PMOVZXWQ"
	case TPMULDQ:
		return "PMULDQ"
	case TPMULHRSW:
		return "PMULHRSW"
	case TPMULHRW:
		return "PMULHRW"
	case TPMULHUW:
		return "PMULHUW"
	case TPMULHW:
		return "PMULHW"
	case TPMULLD:
		return "PMULLD"
	case TPMULLW:
		return "PMULLW"
	case TPMULUDQ:
		return "PMULUDQ"
	case TPOP:
		return "POP"
	case TPOPCNT:
		return "POPCNT"
	case TPOR:
		return "POR"
	case TPREFETCHNTA:
		return "PREFETCHNTA"
	case TPREFETCHT0:
		return "PREFETCHT0"
	case TPREFETCHT1:
		return "PREFETCHT1"
	case TPREFETCHT2:
		return "PREFETCHT2"
	case TPREFETCHW:
		return "PREFETCHW"
	case TPREFETCHWT1:
		return "PREFETCHWT1"
	case TPSADBW:
		return "PSADBW"
	case TPSHUFB:
		return "PSHUFB"
	case TPSHUFD:
		return "PSHUFD"
	case TPSHUFHW:
		return "PSHUFHW"
	case TPSHUFLW:
		return "PSHUFLW"
	case TPSHUFW:
		return "PSHUFW"
	case TPSIGNB:
		return "PSIGNB"
	case TPSIGND:
		return "PSIGND"
	case TPSIGNW:
		return "PSIGNW"
	case TPSLLD:
		return "PSLLD"
	case TPSLLDQ:
		return "PSLLDQ"
	case TPSLLQ:
		return "PSLLQ"
	case TPSLLW:
		return "PSLLW"
	case TPSRAD:
		return "PSRAD"
	case TPSRAW:
		return "PSRAW"
	case TPSRLD:
		return "PSRLD"
	case TPSRLDQ:
		return "PSRLDQ"
	case TPSRLQ:
		return "PSRLQ"
	case TPSRLW:
		return "PSRLW"
	case TPSUBB:
		return "PSUBB"
	case TPSUBD:
		return "PSUBD"
	case TPSUBQ:
		return "PSUBQ"
	case TPSUBSB:
		return "PSUBSB"
	case TPSUBSW:
		return "PSUBSW"
	case TPSUBUSB:
		return "PSUBUSB"
	case TPSUBUSW:
		return "PSUBUSW"
	case TPSUBW:
		return "PSUBW"
	case TPSWAPD:
		return "PSWAPD"
	case TPTEST:
		return "PTEST"
	case TPUNPCKHBW:
		return "PUNPCKHBW"
	case TPUNPCKHDQ:
		return "PUNPCKHDQ"
	case TPUNPCKHQDQ:
		return "PUNPCKHQDQ"
	case TPUNPCKHWD:
		return "PUNPCKHWD"
	case TPUNPCKLBW:
		return "PUNPCKLBW"
	case TPUNPCKLDQ:
		return "PUNPCKLDQ"
	case TPUNPCKLQDQ:
		return "PUNPCKLQDQ"
	case TPUNPCKLWD:
		return "PUNPCKLWD"
	case TPUSH:
		return "PUSH"
	case TPXOR:
		return "PXOR"
	case TRCL:
		return "RCL"
	case TRCPPS:
		return "RCPPS"
	case TRCPSS:
		return "RCPSS"
	case TRCR:
		return "RCR"
	case TRDRAND:
		return "RDRAND"
	case TRDSEED:
		return "RDSEED"
	case TRDTSC:
		return "RDTSC"
	case TRDTSCP:
		return "RDTSCP"
	case TRET:
		return "RET"
	case TROL:
		return "ROL"
	case TROR:
		return "ROR"
	case TRORX:
		return "RORX"
	case TROUNDPD:
		return "ROUNDPD"
	case TROUNDPS:
		return "ROUNDPS"
	case TROUNDSD:
		return "ROUNDSD"
	case TROUNDSS:
		return "ROUNDSS"
	case TRSQRTPS:
		return "RSQRTPS"
	case TRSQRTSS:
		return "RSQRTSS"
	case TSAL:
		return "SAL"
	case TSAR:
		return "SAR"
	case TSARX:
		return "SARX"
	case TSBB:
		return "SBB"
	case TSETA:
		return "SETA"
	case TSETAE:
		return "SETAE"
	case TSETB:
		return "SETB"
	case TSETBE:
		return "SETBE"
	case TSETC:
		return "SETC"
	case TSETE:
		return "SETE"
	case TSETG:
		return "SETG"
	case TSETGE:
		return "SETGE"
	case TSETL:
		return "SETL"
	case TSETLE:
		return "SETLE"
	case TSETNA:
		return "SETNA"
	case TSETNAE:
		return "SETNAE"
	case TSETNB:
		return "SETNB"
	case TSETNBE:
		return "SETNBE"
	case TSETNC:
		return "SETNC"
	case TSETNE:
		return "SETNE"
	case TSETNG:
		return "SETNG"
	case TSETNGE:
		return "SETNGE"
	case TSETNL:
		return "SETNL"
	case TSETNLE:
		return "SETNLE"
	case TSETNO:
		return "SETNO"
	case TSETNP:
		return "SETNP"
	case TSETNS:
		return "SETNS"
	case TSETNZ:
		return "SETNZ"
	case TSETO:
		return "SETO"
	case TSETP:
		return "SETP"
	case TSETPE:
		return "SETPE"
	case TSETPO:
		return "SETPO"
	case TSETS:
		return "SETS"
	case TSETZ:
		return "SETZ"
	case TSFENCE:
		return "SFENCE"
	case TSHA1MSG1:
		return "SHA1MSG1"
	case TSHA1MSG2:
		return "SHA1MSG2"
	case TSHA1NEXTE:
		return "SHA1NEXTE"
	case TSHA1RNDS4:
		return "SHA1RNDS4"
	case TSHA256MSG1:
		return "SHA256MSG1"
	case TSHA256MSG2:
		return "SHA256MSG2"
	case TSHA256RNDS2:
		return "SHA256RNDS2"
	case TSHL:
		return "SHL"
	case TSHLD:
		return "SHLD"
	case TSHLX:
		return "SHLX"
	case TSHR:
		return "SHR"
	case TSHRD:
		return "SHRD"
	case TSHRX:
		return "SHRX"
	case TSHUFPD:
		return "SHUFPD"
	case TSHUFPS:
		return "SHUFPS"
	case TSQRTPD:
		return "SQRTPD"
	case TSQRTPS:
		return "SQRTPS"
	case TSQRTSD:
		return "SQRTSD"
	case TSQRTSS:
		return "SQRTSS"
	case TSTC:
		return "STC"
	case TSTD:
		return "STD"
	case TSTMXCSR:
		return "STMXCSR"
	case TSUB:
		return "SUB"
	case TSUBPD:
		return "SUBPD"
	case TSUBPS:
		return "SUBPS"
	case TSUBSD:
		return "SUBSD"
	case TSUBSS:
		return "SUBSS"
	case TT1MSKC:
		return "T1MSKC"
	case TTEST:
		return "TEST"
	case TTZCNT:
		return "TZCNT"
	case TTZMSK:
		return "TZMSK"
	case TUCOMISD:
		return "UCOMISD"
	case TUCOMISS:
		return "UCOMISS"
	case TUD2:
		return "UD2"
	case TUNPCKHPD:
		return "UNPCKHPD"
	case TUNPCKHPS:
		return "UNPCKHPS"
	case TUNPCKLPD:
		return "UNPCKLPD"
	case TUNPCKLPS:
		return "UNPCKLPS"
	case TVADDPD:
		return "VADDPD"
	case TVADDPS:
		return "VADDPS"
	case TVADDSD:
		return "VADDSD"
	case TVADDSS:
		return "VADDSS"
	case TVADDSUBPD:
		return "VADDSUBPD"
	case TVADDSUBPS:
		return "VADDSUBPS"
	case TVAESDEC:
		return "VAESDEC"
	case TVAESDECLAST:
		return "VAESDECLAST"
	case TVAESENC:
		return "VAESENC"
	case TVAESENCLAST:
		return "VAESENCLAST"
	case TVAESIMC:
		return "VAESIMC"
	case TVAESKEYGENASSIST:
		return "VAESKEYGENASSIST"
	case TVALIGND:
		return "VALIGND"
	case TVALIGNQ:
		return "VALIGNQ"
	case TVANDNPD:
		return "VANDNPD"
	case TVANDNPS:
		return "VANDNPS"
	case TVANDPD:
		return "VANDPD"
	case TVANDPS:
		return "VANDPS"
	case TVBLENDMPD:
		return "VBLENDMPD"
	case TVBLENDMPS:
		return "VBLENDMPS"
	case TVBLENDPD:
		return "VBLENDPD"
	case TVBLENDPS:
		return "VBLENDPS"
	case TVBLENDVPD:
		return "VBLENDVPD"
	case TVBLENDVPS:
		return "VBLENDVPS"
	case TVBROADCASTF128:
		return "VBROADCASTF128"
	case TVBROADCASTF32X2:
		return "VBROADCASTF32X2"
	case TVBROADCASTF32X4:
		return "VBROADCASTF32X4"
	case TVBROADCASTF32X8:
		return "VBROADCASTF32X8"
	case TVBROADCASTF64X2:
		return "VBROADCASTF64X2"
	case TVBROADCASTF64X4:
		return "VBROADCASTF64X4"
	case TVBROADCASTI128:
		return "VBROADCASTI128"
	case TVBROADCASTI32X2:
		return "VBROADCASTI32X2"
	case TVBROADCASTI32X4:
		return "VBROADCASTI32X4"
	case TVBROADCASTI32X8:
		return "VBROADCASTI32X8"
	case TVBROADCASTI64X2:
		return "VBROADCASTI64X2"
	case TVBROADCASTI64X4:
		return "VBROADCASTI64X4"
	case TVBROADCASTSD:
		return "VBROADCASTSD"
	case TVBROADCASTSS:
		return "VBROADCASTSS"
	case TVCMPPD:
		return "VCMPPD"
	case TVCMPPS:
		return "VCMPPS"
	case TVCMPSD:
		return "VCMPSD"
	case TVCMPSS:
		return "VCMPSS"
	case TVCOMISD:
		return "VCOMISD"
	case TVCOMISS:
		return "VCOMISS"
	case TVCOMPRESSPD:
		return "VCOMPRESSPD"
	case TVCOMPRESSPS:
		return "VCOMPRESSPS"
	case TVCVTDQ2PD:
		return "VCVTDQ2PD"
	case TVCVTDQ2PS:
		return "VCVTDQ2PS"
	case TVCVTPD2DQ:
		return "VCVTPD2DQ"
	case TVCVTPD2PS:
		return "VCVTPD2PS"
	case TVCVTPD2QQ:
		return "VCVTPD2QQ"
	case TVCVTPD2UDQ:
		return "VCVTPD2UDQ"
	case TVCVTPD2UQQ:
		return "VCVTPD2UQQ"
	case TVCVTPH2PS:
		return "VCVTPH2PS"
	case TVCVTPS2DQ:
		return "VCVTPS2DQ"
	case TVCVTPS2PD:
		return "VCVTPS2PD"
	case TVCVTPS2PH:
		return "VCVTPS2PH"
	case TVCVTPS2QQ:
		return "VCVTPS2QQ"
	case TVCVTPS2UDQ:
		return "VCVTPS2UDQ"
	case TVCVTPS2UQQ:
		return "VCVTPS2UQQ"
	case TVCVTQQ2PD:
		return "VCVTQQ2PD"
	case TVCVTQQ2PS:
		return "VCVTQQ2PS"
	case TVCVTSD2SI:
		return "VCVTSD2SI"
	case TVCVTSD2SS:
		return "VCVTSD2SS"
	case TVCVTSD2USI:
		return "VCVTSD2USI"
	case TVCVTSI2SD:
		return "VCVTSI2SD"
	case TVCVTSI2SS:
		return "VCVTSI2SS"
	case TVCVTSS2SD:
		return "VCVTSS2SD"
	case TVCVTSS2SI:
		return "VCVTSS2SI"
	case TVCVTSS2USI:
		return "VCVTSS2USI"
	case TVCVTTPD2DQ:
		return "VCVTTPD2DQ"
	case TVCVTTPD2QQ:
		return "VCVTTPD2QQ"
	case TVCVTTPD2UDQ:
		return "VCVTTPD2UDQ"
	case TVCVTTPD2UQQ:
		return "VCVTTPD2UQQ"
	case TVCVTTPS2DQ:
		return "VCVTTPS2DQ"
	case TVCVTTPS2QQ:
		return "VCVTTPS2QQ"
	case TVCVTTPS2UDQ:
		return "VCVTTPS2UDQ"
	case TVCVTTPS2UQQ:
		return "VCVTTPS2UQQ"
	case TVCVTTSD2SI:
		return "VCVTTSD2SI"
	case TVCVTTSD2USI:
		return "VCVTTSD2USI"
	case TVCVTTSS2SI:
		return "VCVTTSS2SI"
	case TVCVTTSS2USI:
		return "VCVTTSS2USI"
	case TVCVTUDQ2PD:
		return "VCVTUDQ2PD"
	case TVCVTUDQ2PS:
		return "VCVTUDQ2PS"
	case TVCVTUQQ2PD:
		return "VCVTUQQ2PD"
	case TVCVTUQQ2PS:
		return "VCVTUQQ2PS"
	case TVCVTUSI2SD:
		return "VCVTUSI2SD"
	case TVCVTUSI2SS:
		return "VCVTUSI2SS"
	case TVDBPSADBW:
		return "VDBPSADBW"
	case TVDIVPD:
		return "VDIVPD"
	case TVDIVPS:
		return "VDIVPS"
	case TVDIVSD:
		return "VDIVSD"
	case TVDIVSS:
		return "VDIVSS"
	case TVDPPD:
		return "VDPPD"
	case TVDPPS:
		return "VDPPS"
	case TVEXP2PD:
		return "VEXP2PD"
	case TVEXP2PS:
		return "VEXP2PS"
	case TVEXPANDPD:
		return "VEXPANDPD"
	case TVEXPANDPS:
		return "VEXPANDPS"
	case TVEXTRACTF128:
		return "VEXTRACTF128"
	case TVEXTRACTF32X4:
		return "VEXTRACTF32X4"
	case TVEXTRACTF32X8:
		return "VEXTRACTF32X8"
	case TVEXTRACTF64X2:
		return "VEXTRACTF64X2"
	case TVEXTRACTF64X4:
		return "VEXTRACTF64X4"
	case TVEXTRACTI128:
		return "VEXTRACTI128"
	case TVEXTRACTI32X4:
		return "VEXTRACTI32X4"
	case TVEXTRACTI32X8:
		return "VEXTRACTI32X8"
	case TVEXTRACTI64X2:
		return "VEXTRACTI64X2"
	case TVEXTRACTI64X4:
		return "VEXTRACTI64X4"
	case TVEXTRACTPS:
		return "VEXTRACTPS"
	case TVFIXUPIMMPD:
		return "VFIXUPIMMPD"
	case TVFIXUPIMMPS:
		return "VFIXUPIMMPS"
	case TVFIXUPIMMSD:
		return "VFIXUPIMMSD"
	case TVFIXUPIMMSS:
		return "VFIXUPIMMSS"
	case TVFMADD132PD:
		return "VFMADD132PD"
	case TVFMADD132PS:
		return "VFMADD132PS"
	case TVFMADD132SD:
		return "VFMADD132SD"
	case TVFMADD132SS:
		return "VFMADD132SS"
	case TVFMADD213PD:
		return "VFMADD213PD"
	case TVFMADD213PS:
		return "VFMADD213PS"
	case TVFMADD213SD:
		return "VFMADD213SD"
	case TVFMADD213SS:
		return "VFMADD213SS"
	case TVFMADD231PD:
		return "VFMADD231PD"
	case TVFMADD231PS:
		return "VFMADD231PS"
	case TVFMADD231SD:
		return "VFMADD231SD"
	case TVFMADD231SS:
		return "VFMADD231SS"
	case TVFMADDPD:
		return "VFMADDPD"
	case TVFMADDPS:
		return "VFMADDPS"
	case TVFMADDSD:
		return "VFMADDSD"
	case TVFMADDSS:
		return "VFMADDSS"
	case TVFMADDSUB132PD:
		return "VFMADDSUB132PD"
	case TVFMADDSUB132PS:
		return "VFMADDSUB132PS"
	case TVFMADDSUB213PD:
		return "VFMADDSUB213PD"
	case TVFMADDSUB213PS:
		return "VFMADDSUB213PS"
	case TVFMADDSUB231PD:
		return "VFMADDSUB231PD"
	case TVFMADDSUB231PS:
		return "VFMADDSUB231PS"
	case TVFMADDSUBPD:
		return "VFMADDSUBPD"
	case TVFMADDSUBPS:
		return "VFMADDSUBPS"
	case TVFMSUB132PD:
		return "VFMSUB132PD"
	case TVFMSUB132PS:
		return "VFMSUB132PS"
	case TVFMSUB132SD:
		return "VFMSUB132SD"
	case TVFMSUB132SS:
		return "VFMSUB132SS"
	case TVFMSUB213PD:
		return "VFMSUB213PD"
	case TVFMSUB213PS:
		return "VFMSUB213PS"
	case TVFMSUB213SD:
		return "VFMSUB213SD"
	case TVFMSUB213SS:
		return "VFMSUB213SS"
	case TVFMSUB231PD:
		return "VFMSUB231PD"
	case TVFMSUB231PS:
		return "VFMSUB231PS"
	case TVFMSUB231SD:
		return "VFMSUB231SD"
	case TVFMSUB231SS:
		return "VFMSUB231SS"
	case TVFMSUBADD132PD:
		return "VFMSUBADD132PD"
	case TVFMSUBADD132PS:
		return "VFMSUBADD132PS"
	case TVFMSUBADD213PD:
		return "VFMSUBADD213PD"
	case TVFMSUBADD213PS:
		return "VFMSUBADD213PS"
	case TVFMSUBADD231PD:
		return "VFMSUBADD231PD"
	case TVFMSUBADD231PS:
		return "VFMSUBADD231PS"
	case TVFMSUBADDPD:
		return "VFMSUBADDPD"
	case TVFMSUBADDPS:
		return "VFMSUBADDPS"
	case TVFMSUBPD:
		return "VFMSUBPD"
	case TVFMSUBPS:
		return "VFMSUBPS"
	case TVFMSUBSD:
		return "VFMSUBSD"
	case TVFMSUBSS:
		return "VFMSUBSS"
	case TVFNMADD132PD:
		return "VFNMADD132PD"
	case TVFNMADD132PS:
		return "VFNMADD132PS"
	case TVFNMADD132SD:
		return "VFNMADD132SD"
	case TVFNMADD132SS:
		return "VFNMADD132SS"
	case TVFNMADD213PD:
		return "VFNMADD213PD"
	case TVFNMADD213PS:
		return "VFNMADD213PS"
	case TVFNMADD213SD:
		return "VFNMADD213SD"
	case TVFNMADD213SS:
		return "VFNMADD213SS"
	case TVFNMADD231PD:
		return "VFNMADD231PD"
	case TVFNMADD231PS:
		return "VFNMADD231PS"
	case TVFNMADD231SD:
		return "VFNMADD231SD"
	case TVFNMADD231SS:
		return "VFNMADD231SS"
	case TVFNMADDPD:
		return "VFNMADDPD"
	case TVFNMADDPS:
		return "VFNMADDPS"
	case TVFNMADDSD:
		return "VFNMADDSD"
	case TVFNMADDSS:
		return "VFNMADDSS"
	case TVFNMSUB132PD:
		return "VFNMSUB132PD"
	case TVFNMSUB132PS:
		return "VFNMSUB132PS"
	case TVFNMSUB132SD:
		return "VFNMSUB132SD"
	case TVFNMSUB132SS:
		return "VFNMSUB132SS"
	case TVFNMSUB213PD:
		return "VFNMSUB213PD"
	case TVFNMSUB213PS:
		return "VFNMSUB213PS"
	case TVFNMSUB213SD:
		return "VFNMSUB213SD"
	case TVFNMSUB213SS:
		return "VFNMSUB213SS"
	case TVFNMSUB231PD:
		return "VFNMSUB231PD"
	case TVFNMSUB231PS:
		return "VFNMSUB231PS"
	case TVFNMSUB231SD:
		return "VFNMSUB231SD"
	case TVFNMSUB231SS:
		return "VFNMSUB231SS"
	case TVFNMSUBPD:
		return "VFNMSUBPD"
	case TVFNMSUBPS:
		return "VFNMSUBPS"
	case TVFNMSUBSD:
		return "VFNMSUBSD"
	case TVFNMSUBSS:
		return "VFNMSUBSS"
	case TVFPCLASSPD:
		return "VFPCLASSPD"
	case TVFPCLASSPS:
		return "VFPCLASSPS"
	case TVFPCLASSSD:
		return "VFPCLASSSD"
	case TVFPCLASSSS:
		return "VFPCLASSSS"
	case TVFRCZPD:
		return "VFRCZPD"
	case TVFRCZPS:
		return "VFRCZPS"
	case TVFRCZSD:
		return "VFRCZSD"
	case TVFRCZSS:
		return "VFRCZSS"
	case TVGATHERDPD:
		return "VGATHERDPD"
	case TVGATHERDPS:
		return "VGATHERDPS"
	case TVGATHERPF0DPD:
		return "VGATHERPF0DPD"
	case TVGATHERPF0DPS:
		return "VGATHERPF0DPS"
	case TVGATHERPF0QPD:
		return "VGATHERPF0QPD"
	case TVGATHERPF0QPS:
		return "VGATHERPF0QPS"
	case TVGATHERPF1DPD:
		return "VGATHERPF1DPD"
	case TVGATHERPF1DPS:
		return "VGATHERPF1DPS"
	case TVGATHERPF1QPD:
		return "VGATHERPF1QPD"
	case TVGATHERPF1QPS:
		return "VGATHERPF1QPS"
	case TVGATHERQPD:
		return "VGATHERQPD"
	case TVGATHERQPS:
		return "VGATHERQPS"
	case TVGETEXPPD:
		return "VGETEXPPD"
	case TVGETEXPPS:
		return "VGETEXPPS"
	case TVGETEXPSD:
		return "VGETEXPSD"
	case TVGETEXPSS:
		return "VGETEXPSS"
	case TVGETMANTPD:
		return "VGETMANTPD"
	case TVGETMANTPS:
		return "VGETMANTPS"
	case TVGETMANTSD:
		return "VGETMANTSD"
	case TVGETMANTSS:
		return "VGETMANTSS"
	case TVHADDPD:
		return "VHADDPD"
	case TVHADDPS:
		return "VHADDPS"
	case TVHSUBPD:
		return "VHSUBPD"
	case TVHSUBPS:
		return "VHSUBPS"
	case TVINSERTF128:
		return "VINSERTF128"
	case TVINSERTF32X4:
		return "VINSERTF32X4"
	case TVINSERTF32X8:
		return "VINSERTF32X8"
	case TVINSERTF64X2:
		return "VINSERTF64X2"
	case TVINSERTF64X4:
		return "VINSERTF64X4"
	case TVINSERTI128:
		return "VINSERTI128"
	case TVINSERTI32X4:
		return "VINSERTI32X4"
	case TVINSERTI32X8:
		return "VINSERTI32X8"
	case TVINSERTI64X2:
		return "VINSERTI64X2"
	case TVINSERTI64X4:
		return "VINSERTI64X4"
	case TVINSERTPS:
		return "VINSERTPS"
	case TVLDDQU:
		return "VLDDQU"
	case TVLDMXCSR:
		return "VLDMXCSR"
	case TVMASKMOVDQU:
		return "VMASKMOVDQU"
	case TVMASKMOVPD:
		return "VMASKMOVPD"
	case TVMASKMOVPS:
		return "VMASKMOVPS"
	case TVMAXPD:
		return "VMAXPD"
	case TVMAXPS:
		return "VMAXPS"
	case TVMAXSD:
		return "VMAXSD"
	case TVMAXSS:
		return "VMAXSS"
	case TVMINPD:
		return "VMINPD"
	case TVMINPS:
		return "VMINPS"
	case TVMINSD:
		return "VMINSD"
	case TVMINSS:
		return "VMINSS"
	case TVMOVAPD:
		return "VMOVAPD"
	case TVMOVAPS:
		return "VMOVAPS"
	case TVMOVD:
		return "VMOVD"
	case TVMOVDDUP:
		return "VMOVDDUP"
	case TVMOVDQA:
		return "VMOVDQA"
	case TVMOVDQA32:
		return "VMOVDQA32"
	case TVMOVDQA64:
		return "VMOVDQA64"
	case TVMOVDQU:
		return "VMOVDQU"
	case TVMOVDQU16:
		return "VMOVDQU16"
	case TVMOVDQU32:
		return "VMOVDQU32"
	case TVMOVDQU64:
		return "VMOVDQU64"
	case TVMOVDQU8:
		return "VMOVDQU8"
	case TVMOVHLPS:
		return "VMOVHLPS"
	case TVMOVHPD:
		return "VMOVHPD"
	case TVMOVHPS:
		return "VMOVHPS"
	case TVMOVLHPS:
		return "VMOVLHPS"
	case TVMOVLPD:
		return "VMOVLPD"
	case TVMOVLPS:
		return "VMOVLPS"
	case TVMOVMSKPD:
		return "VMOVMSKPD"
	case TVMOVMSKPS:
		return "VMOVMSKPS"
	case TVMOVNTDQ:
		return "VMOVNTDQ"
	case TVMOVNTDQA:
		return "VMOVNTDQA"
	case TVMOVNTPD:
		return "VMOVNTPD"
	case TVMOVNTPS:
		return "VMOVNTPS"
	case TVMOVQ:
		return "VMOVQ"
	case TVMOVSD:
		return "VMOVSD"
	case TVMOVSHDUP:
		return "VMOVSHDUP"
	case TVMOVSLDUP:
		return "VMOVSLDUP"
	case TVMOVSS:
		return "VMOVSS"
	case TVMOVUPD:
		return "VMOVUPD"
	case TVMOVUPS:
		return "VMOVUPS"
	case TVMPSADBW:
		return "VMPSADBW"
	case TVMULPD:
		return "VMULPD"
	case TVMULPS:
		return "VMULPS"
	case TVMULSD:
		return "VMULSD"
	case TVMULSS:
		return "VMULSS"
	case TVORPD:
		return "VORPD"
	case TVORPS:
		return "VORPS"
	case TVPABSB:
		return "VPABSB"
	case TVPABSD:
		return "VPABSD"
	case TVPABSQ:
		return "VPABSQ"
	case TVPABSW:
		return "VPABSW"
	case TVPACKSSDW:
		return "VPACKSSDW"
	case TVPACKSSWB:
		return "VPACKSSWB"
	case TVPACKUSDW:
		return "VPACKUSDW"
	case TVPACKUSWB:
		return "VPACKUSWB"
	case TVPADDB:
		return "VPADDB"
	case TVPADDD:
		return "VPADDD"
	case TVPADDQ:
		return "VPADDQ"
	case TVPADDSB:
		return "VPADDSB"
	case TVPADDSW:
		return "VPADDSW"
	case TVPADDUSB:
		return "VPADDUSB"
	case TVPADDUSW:
		return "VPADDUSW"
	case TVPADDW:
		return "VPADDW"
	case TVPALIGNR:
		return "VPALIGNR"
	case TVPAND:
		return "VPAND"
	case TVPANDD:
		return "VPANDD"
	case TVPANDN:
		return "VPANDN"
	case TVPANDND:
		return "VPANDND"
	case TVPANDNQ:
		return "VPANDNQ"
	case TVPANDQ:
		return "VPANDQ"
	case TVPAVGB:
		return "VPAVGB"
	case TVPAVGW:
		return "VPAVGW"
	case TVPBLENDD:
		return "VPBLENDD"
	case TVPBLENDMB:
		return "VPBLENDMB"
	case TVPBLENDMD:
		return "VPBLENDMD"
	case TVPBLENDMQ:
		return "VPBLENDMQ"
	case TVPBLENDMW:
		return "VPBLENDMW"
	case TVPBLENDVB:
		return "VPBLENDVB"
	case TVPBLENDW:
		return "VPBLENDW"
	case TVPBROADCASTB:
		return "VPBROADCASTB"
	case TVPBROADCASTD:
		return "VPBROADCASTD"
	case TVPBROADCASTMB2Q:
		return "VPBROADCASTMB2Q"
	case TVPBROADCASTMW2D:
		return "VPBROADCASTMW2D"
	case TVPBROADCASTQ:
		return "VPBROADCASTQ"
	case TVPBROADCASTW:
		return "VPBROADCASTW"
	case TVPCLMULQDQ:
		return "VPCLMULQDQ"
	case TVPCMOV:
		return "VPCMOV"
	case TVPCMPB:
		return "VPCMPB"
	case TVPCMPD:
		return "VPCMPD"
	case TVPCMPEQB:
		return "VPCMPEQB"
	case TVPCMPEQD:
		return "VPCMPEQD"
	case TVPCMPEQQ:
		return "VPCMPEQQ"
	case TVPCMPEQW:
		return "VPCMPEQW"
	case TVPCMPESTRI:
		return "VPCMPESTRI"
	case TVPCMPESTRM:
		return "VPCMPESTRM"
	case TVPCMPGTB:
		return "VPCMPGTB"
	case TVPCMPGTD:
		return "VPCMPGTD"
	case TVPCMPGTQ:
		return "VPCMPGTQ"
	case TVPCMPGTW:
		return "VPCMPGTW"
	case TVPCMPISTRI:
		return "VPCMPISTRI"
	case TVPCMPISTRM:
		return "VPCMPISTRM"
	case TVPCMPQ:
		return "VPCMPQ"
	case TVPCMPUB:
		return "VPCMPUB"
	case TVPCMPUD:
		return "VPCMPUD"
	case TVPCMPUQ:
		return "VPCMPUQ"
	case TVPCMPUW:
		return "VPCMPUW"
	case TVPCMPW:
		return "VPCMPW"
	case TVPCOMB:
		return "VPCOMB"
	case TVPCOMD:
		return "VPCOMD"
	case TVPCOMPRESSD:
		return "VPCOMPRESSD"
	case TVPCOMPRESSQ:
		return "VPCOMPRESSQ"
	case TVPCOMQ:
		return "VPCOMQ"
	case TVPCOMUB:
		return "VPCOMUB"
	case TVPCOMUD:
		return "VPCOMUD"
	case TVPCOMUQ:
		return "VPCOMUQ"
	case TVPCOMUW:
		return "VPCOMUW"
	case TVPCOMW:
		return "VPCOMW"
	case TVPCONFLICTD:
		return "VPCONFLICTD"
	case TVPCONFLICTQ:
		return "VPCONFLICTQ"
	case TVPERM2F128:
		return "VPERM2F128"
	case TVPERM2I128:
		return "VPERM2I128"
	case TVPERMB:
		return "VPERMB"
	case TVPERMD:
		return "VPERMD"
	case TVPERMI2B:
		return "VPERMI2B"
	case TVPERMI2D:
		return "VPERMI2D"
	case TVPERMI2PD:
		return "VPERMI2PD"
	case TVPERMI2PS:
		return "VPERMI2PS"
	case TVPERMI2Q:
		return "VPERMI2Q"
	case TVPERMI2W:
		return "VPERMI2W"
	case TVPERMIL2PD:
		return "VPERMIL2PD"
	case TVPERMIL2PS:
		return "VPERMIL2PS"
	case TVPERMILPD:
		return "VPERMILPD"
	case TVPERMILPS:
		return "VPERMILPS"
	case TVPERMPD:
		return "VPERMPD"
	case TVPERMPS:
		return "VPERMPS"
	case TVPERMQ:
		return "VPERMQ"
	case TVPERMT2B:
		return "VPERMT2B"
	case TVPERMT2D:
		return "VPERMT2D"
	case TVPERMT2PD:
		return "VPERMT2PD"
	case TVPERMT2PS:
		return "VPERMT2PS"
	case TVPERMT2Q:
		return "VPERMT2Q"
	case TVPERMT2W:
		return "VPERMT2W"
	case TVPERMW:
		return "VPERMW"
	case TVPEXPANDD:
		return "VPEXPANDD"
	case TVPEXPANDQ:
		return "VPEXPANDQ"
	case TVPEXTRB:
		return "VPEXTRB"
	case TVPEXTRD:
		return "VPEXTRD"
	case TVPEXTRQ:
		return "VPEXTRQ"
	case TVPEXTRW:
		return "VPEXTRW"
	case TVPGATHERDD:
		return "VPGATHERDD"
	case TVPGATHERDQ:
		return "VPGATHERDQ"
	case TVPGATHERQD:
		return "VPGATHERQD"
	case TVPGATHERQQ:
		return "VPGATHERQQ"
	case TVPHADDBD:
		return "VPHADDBD"
	case TVPHADDBQ:
		return "VPHADDBQ"
	case TVPHADDBW:
		return "VPHADDBW"
	case TVPHADDD:
		return "VPHADDD"
	case TVPHADDDQ:
		return "VPHADDDQ"
	case TVPHADDSW:
		return "VPHADDSW"
	case TVPHADDUBD:
		return "VPHADDUBD"
	case TVPHADDUBQ:
		return "VPHADDUBQ"
	case TVPHADDUBW:
		return "VPHADDUBW"
	case TVPHADDUDQ:
		return "VPHADDUDQ"
	case TVPHADDUWD:
		return "VPHADDUWD"
	case TVPHADDUWQ:
		return "VPHADDUWQ"
	case TVPHADDW:
		return "VPHADDW"
	case TVPHADDWD:
		return "VPHADDWD"
	case TVPHADDWQ:
		return "VPHADDWQ"
	case TVPHMINPOSUW:
		return "VPHMINPOSUW"
	case TVPHSUBBW:
		return "VPHSUBBW"
	case TVPHSUBD:
		return "VPHSUBD"
	case TVPHSUBDQ:
		return "VPHSUBDQ"
	case TVPHSUBSW:
		return "VPHSUBSW"
	case TVPHSUBW:
		return "VPHSUBW"
	case TVPHSUBWD:
		return "VPHSUBWD"
	case TVPINSRB:
		return "VPINSRB"
	case TVPINSRD:
		return "VPINSRD"
	case TVPINSRQ:
		return "VPINSRQ"
	case TVPINSRW:
		return "VPINSRW"
	case TVPLZCNTD:
		return "VPLZCNTD"
	case TVPLZCNTQ:
		return "VPLZCNTQ"
	case TVPMACSDD:
		return "VPMACSDD"
	case TVPMACSDQH:
		return "VPMACSDQH"
	case TVPMACSDQL:
		return "VPMACSDQL"
	case TVPMACSSDD:
		return "VPMACSSDD"
	case TVPMACSSDQH:
		return "VPMACSSDQH"
	case TVPMACSSDQL:
		return "VPMACSSDQL"
	case TVPMACSSWD:
		return "VPMACSSWD"
	case TVPMACSSWW:
		return "VPMACSSWW"
	case TVPMACSWD:
		return "VPMACSWD"
	case TVPMACSWW:
		return "VPMACSWW"
	case TVPMADCSSWD:
		return "VPMADCSSWD"
	case TVPMADCSWD:
		return "VPMADCSWD"
	case TVPMADD52HUQ:
		return "VPMADD52HUQ"
	case TVPMADD52LUQ:
		return "VPMADD52LUQ"
	case TVPMADDUBSW:
		return "VPMADDUBSW"
	case TVPMADDWD:
		return "VPMADDWD"
	case TVPMASKMOVD:
		return "VPMASKMOVD"
	case TVPMASKMOVQ:
		return "VPMASKMOVQ"
	case TVPMAXSB:
		return "VPMAXSB"
	case TVPMAXSD:
		return "VPMAXSD"
	case TVPMAXSQ:
		return "VPMAXSQ"
	case TVPMAXSW:
		return "VPMAXSW"
	case TVPMAXUB:
		return "VPMAXUB"
	case TVPMAXUD:
		return "VPMAXUD"
	case TVPMAXUQ:
		return "VPMAXUQ"
	case TVPMAXUW:
		return "VPMAXUW"
	case TVPMINSB:
		return "VPMINSB"
	case TVPMINSD:
		return "VPMINSD"
	case TVPMINSQ:
		return "VPMINSQ"
	case TVPMINSW:
		return "VPMINSW"
	case TVPMINUB:
		return "VPMINUB"
	case TVPMINUD:
		return "VPMINUD"
	case TVPMINUQ:
		return "VPMINUQ"
	case TVPMINUW:
		return "VPMINUW"
	case TVPMOVB2M:
		return "VPMOVB2M"
	case TVPMOVD2M:
		return "VPMOVD2M"
	case TVPMOVDB:
		return "VPMOVDB"
	case TVPMOVDW:
		return "VPMOVDW"
	case TVPMOVM2B:
		return "VPMOVM2B"
	case TVPMOVM2D:
		return "VPMOVM2D"
	case TVPMOVM2Q:
		return "VPMOVM2Q"
	case TVPMOVM2W:
		return "VPMOVM2W"
	case TVPMOVMSKB:
		return "VPMOVMSKB"
	case TVPMOVQ2M:
		return "VPMOVQ2M"
	case TVPMOVQB:
		return "VPMOVQB"
	case TVPMOVQD:
		return "VPMOVQD"
	case TVPMOVQW:
		return "VPMOVQW"
	case TVPMOVSDB:
		return "VPMOVSDB"
	case TVPMOVSDW:
		return "VPMOVSDW"
	case TVPMOVSQB:
		return "VPMOVSQB"
	case TVPMOVSQD:
		return "VPMOVSQD"
	case TVPMOVSQW:
		return "VPMOVSQW"
	case TVPMOVSWB:
		return "VPMOVSWB"
	case TVPMOVSXBD:
		return "VPMOVSXBD"
	case TVPMOVSXBQ:
		return "VPMOVSXBQ"
	case TVPMOVSXBW:
		return "VPMOVSXBW"
	case TVPMOVSXDQ:
		return "VPMOVSXDQ"
	case TVPMOVSXWD:
		return "VPMOVSXWD"
	case TVPMOVSXWQ:
		return "VPMOVSXWQ"
	case TVPMOVUSDB:
		return "VPMOVUSDB"
	case TVPMOVUSDW:
		return "VPMOVUSDW"
	case TVPMOVUSQB:
		return "VPMOVUSQB"
	case TVPMOVUSQD:
		return "VPMOVUSQD"
	case TVPMOVUSQW:
		return "VPMOVUSQW"
	case TVPMOVUSWB:
		return "VPMOVUSWB"
	case TVPMOVW2M:
		return "VPMOVW2M"
	case TVPMOVWB:
		return "VPMOVWB"
	case TVPMOVZXBD:
		return "VPMOVZXBD"
	case TVPMOVZXBQ:
		return "VPMOVZXBQ"
	case TVPMOVZXBW:
		return "VPMOVZXBW"
	case TVPMOVZXDQ:
		return "VPMOVZXDQ"
	case TVPMOVZXWD:
		return "VPMOVZXWD"
	case TVPMOVZXWQ:
		return "VPMOVZXWQ"
	case TVPMULDQ:
		return "VPMULDQ"
	case TVPMULHRSW:
		return "VPMULHRSW"
	case TVPMULHUW:
		return "VPMULHUW"
	case TVPMULHW:
		return "VPMULHW"
	case TVPMULLD:
		return "VPMULLD"
	case TVPMULLQ:
		return "VPMULLQ"
	case TVPMULLW:
		return "VPMULLW"
	case TVPMULTISHIFTQB:
		return "VPMULTISHIFTQB"
	case TVPMULUDQ:
		return "VPMULUDQ"
	case TVPOR:
		return "VPOR"
	case TVPORD:
		return "VPORD"
	case TVPORQ:
		return "VPORQ"
	case TVPPERM:
		return "VPPERM"
	case TVPROLD:
		return "VPROLD"
	case TVPROLQ:
		return "VPROLQ"
	case TVPROLVD:
		return "VPROLVD"
	case TVPROLVQ:
		return "VPROLVQ"
	case TVPRORD:
		return "VPRORD"
	case TVPRORQ:
		return "VPRORQ"
	case TVPRORVD:
		return "VPRORVD"
	case TVPRORVQ:
		return "VPRORVQ"
	case TVPROTB:
		return "VPROTB"
	case TVPROTD:
		return "VPROTD"
	case TVPROTQ:
		return "VPROTQ"
	case TVPROTW:
		return "VPROTW"
	case TVPSADBW:
		return "VPSADBW"
	case TVPSCATTERDD:
		return "VPSCATTERDD"
	case TVPSCATTERDQ:
		return "VPSCATTERDQ"
	case TVPSCATTERQD:
		return "VPSCATTERQD"
	case TVPSCATTERQQ:
		return "VPSCATTERQQ"
	case TVPSHAB:
		return "VPSHAB"
	case TVPSHAD:
		return "VPSHAD"
	case TVPSHAQ:
		return "VPSHAQ"
	case TVPSHAW:
		return "VPSHAW"
	case TVPSHLB:
		return "VPSHLB"
	case TVPSHLD:
		return "VPSHLD"
	case TVPSHLQ:
		return "VPSHLQ"
	case TVPSHLW:
		return "VPSHLW"
	case TVPSHUFB:
		return "VPSHUFB"
	case TVPSHUFD:
		return "VPSHUFD"
	case TVPSHUFHW:
		return "VPSHUFHW"
	case TVPSHUFLW:
		return "VPSHUFLW"
	case TVPSIGNB:
		return "VPSIGNB"
	case TVPSIGND:
		return "VPSIGND"
	case TVPSIGNW:
		return "VPSIGNW"
	case TVPSLLD:
		return "VPSLLD"
	case TVPSLLDQ:
		return "VPSLLDQ"
	case TVPSLLQ:
		return "VPSLLQ"
	case TVPSLLVD:
		return "VPSLLVD"
	case TVPSLLVQ:
		return "VPSLLVQ"
	case TVPSLLVW:
		return "VPSLLVW"
	case TVPSLLW:
		return "VPSLLW"
	case TVPSRAD:
		return "VPSRAD"
	case TVPSRAQ:
		return "VPSRAQ"
	case TVPSRAVD:
		return "VPSRAVD"
	case TVPSRAVQ:
		return "VPSRAVQ"
	case TVPSRAVW:
		return "VPSRAVW"
	case TVPSRAW:
		return "VPSRAW"
	case TVPSRLD:
		return "VPSRLD"
	case TVPSRLDQ:
		return "VPSRLDQ"
	case TVPSRLQ:
		return "VPSRLQ"
	case TVPSRLVD:
		return "VPSRLVD"
	case TVPSRLVQ:
		return "VPSRLVQ"
	case TVPSRLVW:
		return "VPSRLVW"
	case TVPSRLW:
		return "VPSRLW"
	case TVPSUBB:
		return "VPSUBB"
	case TVPSUBD:
		return "VPSUBD"
	case TVPSUBQ:
		return "VPSUBQ"
	case TVPSUBSB:
		return "VPSUBSB"
	case TVPSUBSW:
		return "VPSUBSW"
	case TVPSUBUSB:
		return "VPSUBUSB"
	case TVPSUBUSW:
		return "VPSUBUSW"
	case TVPSUBW:
		return "VPSUBW"
	case TVPTERNLOGD:
		return "VPTERNLOGD"
	case TVPTERNLOGQ:
		return "VPTERNLOGQ"
	case TVPTEST:
		return "VPTEST"
	case TVPTESTMB:
		return "VPTESTMB"
	case TVPTESTMD:
		return "VPTESTMD"
	case TVPTESTMQ:
		return "VPTESTMQ"
	case TVPTESTMW:
		return "VPTESTMW"
	case TVPTESTNMB:
		return "VPTESTNMB"
	case TVPTESTNMD:
		return "VPTESTNMD"
	case TVPTESTNMQ:
		return "VPTESTNMQ"
	case TVPTESTNMW:
		return "VPTESTNMW"
	case TVPUNPCKHBW:
		return "VPUNPCKHBW"
	case TVPUNPCKHDQ:
		return "VPUNPCKHDQ"
	case TVPUNPCKHQDQ:
		return "VPUNPCKHQDQ"
	case TVPUNPCKHWD:
		return "VPUNPCKHWD"
	case TVPUNPCKLBW:
		return "VPUNPCKLBW"
	case TVPUNPCKLDQ:
		return "VPUNPCKLDQ"
	case TVPUNPCKLQDQ:
		return "VPUNPCKLQDQ"
	case TVPUNPCKLWD:
		return "VPUNPCKLWD"
	case TVPXOR:
		return "VPXOR"
	case TVPXORD:
		return "VPXORD"
	case TVPXORQ:
		return "VPXORQ"
	case TVRANGEPD:
		return "VRANGEPD"
	case TVRANGEPS:
		return "VRANGEPS"
	case TVRANGESD:
		return "VRANGESD"
	case TVRANGESS:
		return "VRANGESS"
	case TVRCP14PD:
		return "VRCP14PD"
	case TVRCP14PS:
		return "VRCP14PS"
	case TVRCP14SD:
		return "VRCP14SD"
	case TVRCP14SS:
		return "VRCP14SS"
	case TVRCP28PD:
		return "VRCP28PD"
	case TVRCP28PS:
		return "VRCP28PS"
	case TVRCP28SD:
		return "VRCP28SD"
	case TVRCP28SS:
		return "VRCP28SS"
	case TVRCPPS:
		return "VRCPPS"
	case TVRCPSS:
		return "VRCPSS"
	case TVREDUCEPD:
		return "VREDUCEPD"
	case TVREDUCEPS:
		return "VREDUCEPS"
	case TVREDUCESD:
		return "VREDUCESD"
	case TVREDUCESS:
		return "VREDUCESS"
	case TVRNDSCALEPD:
		return "VRNDSCALEPD"
	case TVRNDSCALEPS:
		return "VRNDSCALEPS"
	case TVRNDSCALESD:
		return "VRNDSCALESD"
	case TVRNDSCALESS:
		return "VRNDSCALESS"
	case TVROUNDPD:
		return "VROUNDPD"
	case TVROUNDPS:
		return "VROUNDPS"
	case TVROUNDSD:
		return "VROUNDSD"
	case TVROUNDSS:
		return "VROUNDSS"
	case TVRSQRT14PD:
		return "VRSQRT14PD"
	case TVRSQRT14PS:
		return "VRSQRT14PS"
	case TVRSQRT14SD:
		return "VRSQRT14SD"
	case TVRSQRT14SS:
		return "VRSQRT14SS"
	case TVRSQRT28PD:
		return "VRSQRT28PD"
	case TVRSQRT28PS:
		return "VRSQRT28PS"
	case TVRSQRT28SD:
		return "VRSQRT28SD"
	case TVRSQRT28SS:
		return "VRSQRT28SS"
	case TVRSQRTPS:
		return "VRSQRTPS"
	case TVRSQRTSS:
		return "VRSQRTSS"
	case TVSCALEFPD:
		return "VSCALEFPD"
	case TVSCALEFPS:
		return "VSCALEFPS"
	case TVSCALEFSD:
		return "VSCALEFSD"
	case TVSCALEFSS:
		return "VSCALEFSS"
	case TVSCATTERDPD:
		return "VSCATTERDPD"
	case TVSCATTERDPS:
		return "VSCATTERDPS"
	case TVSCATTERPF0DPD:
		return "VSCATTERPF0DPD"
	case TVSCATTERPF0DPS:
		return "VSCATTERPF0DPS"
	case TVSCATTERPF0QPD:
		return "VSCATTERPF0QPD"
	case TVSCATTERPF0QPS:
		return "VSCATTERPF0QPS"
	case TVSCATTERPF1DPD:
		return "VSCATTERPF1DPD"
	case TVSCATTERPF1DPS:
		return "VSCATTERPF1DPS"
	case TVSCATTERPF1QPD:
		return "VSCATTERPF1QPD"
	case TVSCATTERPF1QPS:
		return "VSCATTERPF1QPS"
	case TVSCATTERQPD:
		return "VSCATTERQPD"
	case TVSCATTERQPS:
		return "VSCATTERQPS"
	case TVSHUFF32X4:
		return "VSHUFF32X4"
	case TVSHUFF64X2:
		return "VSHUFF64X2"
	case TVSHUFI32X4:
		return "VSHUFI32X4"
	case TVSHUFI64X2:
		return "VSHUFI64X2"
	case TVSHUFPD:
		return "VSHUFPD"
	case TVSHUFPS:
		return "VSHUFPS"
	case TVSQRTPD:
		return "VSQRTPD"
	case TVSQRTPS:
		return "VSQRTPS"
	case TVSQRTSD:
		return "VSQRTSD"
	case TVSQRTSS:
		return "VSQRTSS"
	case TVSTMXCSR:
		return "VSTMXCSR"
	case TVSUBPD:
		return "VSUBPD"
	case TVSUBPS:
		return "VSUBPS"
	case TVSUBSD:
		return "VSUBSD"
	case TVSUBSS:
		return "VSUBSS"
	case TVTESTPD:
		return "VTESTPD"
	case TVTESTPS:
		return "VTESTPS"
	case TVUCOMISD:
		return "VUCOMISD"
	case TVUCOMISS:
		return "VUCOMISS"
	case TVUNPCKHPD:
		return "VUNPCKHPD"
	case TVUNPCKHPS:
		return "VUNPCKHPS"
	case TVUNPCKLPD:
		return "VUNPCKLPD"
	case TVUNPCKLPS:
		return "VUNPCKLPS"
	case TVXORPD:
		return "VXORPD"
	case TVXORPS:
		return "VXORPS"
	case TVZEROALL:
		return "VZEROALL"
	case TVZEROUPPER:
		return "VZEROUPPER"
	case TXADD:
		return "XADD"
	case TXCHG:
		return "XCHG"
	case TXGETBV:
		return "XGETBV"
	case TXLATB:
		return "XLATB"
	case TXOR:
		return "XOR"
	case TXORPD:
		return "XORPD"
	case TXORPS:
		return "XORPS"
	}

}
