// generated by stringer -type=Instr,InstrOpType,InstructionType,XmmData codegen; DO NOT EDIT

package codegen

import "fmt"

const _Instr_name = "NONEAADAAMAASADCBADCLADCWADDBADDLADDWADJSPANDBANDLANDWARPLBOUNDLBOUNDWBSFLBSFWBSRLBSRWBTLBTWBTCLBTCWBTRLBTRWBTSLBTSWBYTECLCCLDCLICLTSCMCCMPBCMPLCMPWCMPSBCMPSLCMPSWDAADASDECBDECLDECQDECWDIVBDIVLDIVWENTERHLTIDIVBIDIVLIDIVWIMULBIMULLIMULWINBINLINWINCBINCLINCQINCWINSBINSLINSWINTINTOIRETLIRETWJCCJCSJCXZLJEQJGEJGTJHIJLEJLSJLTJMIJNEJOCJOSJPCJPLJPSLAHFLARLLARWLEALLEAWLEAVELLEAVEWLOCKLODSBLODSLLODSWLONGLOOPLOOPEQLOOPNELSLLLSLWMOVBMOVLMOVWMOVBLSXMOVBLZXMOVBQSXMOVBQZXMOVBWSXMOVBWZXMOVWLSXMOVWLZXMOVWQSXMOVWQZXMOVSBMOVSLMOVSWMULBMULLMULWNEGBNEGLNEGWNOTBNOTLNOTWORBORLORWOUTBOUTLOUTWOUTSBOUTSLOUTSWPAUSEPOPALPOPAWPOPFLPOPFWPOPLPOPWPUSHALPUSHAWPUSHFLPUSHFWPUSHLPUSHWRCLBRCLLRCLWRCRBRCRLRCRWREPREPNROLBROLLROLWRORBRORLRORWSAHFSALBSALLSALWSARBSARLSARWSBBBSBBLSBBWSCASBSCASLSCASWSETCCSETCSSETEQSETGESETGTSETHISETLESETLSSETLTSETMISETNESETOCSETOSSETPCSETPLSETPSCDQCWDSHLBSHLLSHLWSHRBSHRLSHRWSTCSTDSTISTOSBSTOSLSTOSWSUBBSUBLSUBWSYSCALLTESTBTESTLTESTWVERRVERWWAITWORDXCHGBXCHGLXCHGWXLATXORBXORLXORWFMOVBFMOVBPFMOVDFMOVDPFMOVFFMOVFPFMOVLFMOVLPFMOVVFMOVVPFMOVWFMOVWPFMOVXFMOVXPFCOMBFCOMBPFCOMDFCOMDPFCOMDPPFCOMFFCOMFPFCOMLFCOMLPFCOMWFCOMWPFUCOMFUCOMPFUCOMPPFADDDPFADDWFADDLFADDFFADDDFMULDPFMULWFMULLFMULFFMULDFSUBDPFSUBWFSUBLFSUBFFSUBDFSUBRDPFSUBRWFSUBRLFSUBRFFSUBRDFDIVDPFDIVWFDIVLFDIVFFDIVDFDIVRDPFDIVRWFDIVRLFDIVRFFDIVRDFXCHDFFREEFLDCWFLDENVFRSTORFSAVEFSTCWFSTENVFSTSWF2XM1FABSFCHSFCLEXFCOSFDECSTPFINCSTPFINITFLD1FLDL2EFLDL2TFLDLG2FLDLN2FLDPIFLDZFNOPFPATANFPREMFPREM1FPTANFRNDINTFSCALEFSINFSINCOSFSQRTFTSTFXAMFXTRACTFYL2XFYL2XP1CMPXCHGBCMPXCHGLCMPXCHGWCMPXCHG8BCPUIDINVDINVLPGLFENCEMFENCEMOVNTILRDMSRRDPMCRDTSCRSMSFENCESYSRETWBINVDWRMSRXADDBXADDLXADDWCMOVLCCCMOVLCSCMOVLEQCMOVLGECMOVLGTCMOVLHICMOVLLECMOVLLSCMOVLLTCMOVLMICMOVLNECMOVLOCCMOVLOSCMOVLPCCMOVLPLCMOVLPSCMOVQCCCMOVQCSCMOVQEQCMOVQGECMOVQGTCMOVQHICMOVQLECMOVQLSCMOVQLTCMOVQMICMOVQNECMOVQOCCMOVQOSCMOVQPCCMOVQPLCMOVQPSCMOVWCCCMOVWCSCMOVWEQCMOVWGECMOVWGTCMOVWHICMOVWLECMOVWLSCMOVWLTCMOVWMICMOVWNECMOVWOCCMOVWOSCMOVWPCCMOVWPLCMOVWPSADCQADDQANDQBSFQBSRQBTCQBTQBTRQBTSQCMPQCMPSQCMPXCHGQCQODIVQIDIVQIMULQIRETQJCXZQLEAQLEAVEQLODSQMOVQMOVLQSXMOVLQZXMOVNTIQMOVSQMULQNEGQNOTQORQPOPFQPOPQPUSHFQPUSHQRCLQRCRQROLQRORQQUADSALQSARQSBBQSCASQSHLQSHRQSTOSQSUBQTESTQXADDQXCHGQXORQADDPDADDPSADDSDADDSSANDNPDANDNPSANDPDANDPSCMPPDCMPPSCMPSDCMPSSCOMISDCOMISSCVTPD2PLCVTPD2PSCVTPL2PDCVTPL2PSCVTPS2PDCVTPS2PLCVTSD2SLCVTSD2SQCVTSD2SSCVTSL2SDCVTSL2SSCVTSQ2SDCVTSQ2SSCVTSS2SDCVTSS2SLCVTSS2SQCVTTPD2PLCVTTPS2PLCVTTSD2SLCVTTSD2SQCVTTSS2SLCVTTSS2SQDIVPDDIVPSDIVSDDIVSSEMMSFXRSTORFXRSTOR64FXSAVEFXSAVE64LDMXCSRMASKMOVOUMASKMOVQMAXPDMAXPSMAXSDMAXSSMINPDMINPSMINSDMINSSMOVAPDMOVAPSMOVOUMOVHLPSMOVHPDMOVHPSMOVLHPSMOVLPDMOVLPSMOVMSKPDMOVMSKPSMOVNTOMOVNTPDMOVNTPSMOVNTQMOVOMOVQOZXMOVSDMOVSSMOVUPDMOVUPSMULPDMULPSMULSDMULSSORPDORPSPACKSSLWPACKSSWBPACKUSWBPADDBPADDLPADDQPADDSBPADDSWPADDUSBPADDUSWPADDWPANDBPANDLPANDSBPANDSWPANDUSBPANDUSWPANDWPANDPANDNPAVGBPAVGWPCMPEQBPCMPEQLPCMPEQWPCMPGTBPCMPGTLPCMPGTWPEXTRWPFACCPFADDPFCMPEQPFCMPGEPFCMPGTPFMAXPFMINPFMULPFNACCPFPNACCPFRCPPFRCPIT1PFRCPI2TPFRSQIT1PFRSQRTPFSUBPFSUBRPINSRWPINSRDPINSRQPMADDWLPMAXSWPMAXUBPMINSWPMINUBPMOVMSKBPMULHRWPMULHUWPMULHWPMULLWPMULULQPORPSADBWPSHUFHWPSHUFLPSHUFLWPSHUFWPSHUFBPSLLOPSLLLPSLLQPSLLWPSRALPSRAWPSRLOPSRLLPSRLQPSRLWPSUBBPSUBLPSUBQPSUBSBPSUBSWPSUBUSBPSUBUSWPSUBWPSWAPLPUNPCKHBWPUNPCKHLQPUNPCKHQDQPUNPCKHWLPUNPCKLBWPUNPCKLLQPUNPCKLQDQPUNPCKLWLPXORRCPPSRCPSSRSQRTPSRSQRTSSSHUFPDSHUFPSSQRTPDSQRTPSSQRTSDSQRTSSSTMXCSRSUBPDSUBPSSUBSDSUBSSUCOMISDUCOMISSUNPCKHPDUNPCKHPSUNPCKLPDUNPCKLPSXORPDXORPSPF2IWPF2ILPI2FWPI2FLRETFWRETFLRETFQSWAPGSMODECRC32BCRC32QIMUL3QPREFETCHT0PREFETCHT1PREFETCHT2PREFETCHNTAMOVQLBSWAPLBSWAPQAESENCAESENCLASTAESDECAESDECLASTAESIMCAESKEYGENASSISTROUNDPSROUNDSSROUNDPDROUNDSDPSHUFDPCLMULQDQJCXZWFCMOVCCFCMOVCSFCMOVEQFCMOVHIFCMOVLSFCMOVNEFCMOVNUFCMOVUNFCOMIFCOMIPFUCOMIFUCOMIPLAST"

var _Instr_index = [...]uint16{0, 4, 7, 10, 13, 17, 21, 25, 29, 33, 37, 42, 46, 50, 54, 58, 64, 70, 74, 78, 82, 86, 89, 92, 96, 100, 104, 108, 112, 116, 120, 123, 126, 129, 133, 136, 140, 144, 148, 153, 158, 163, 166, 169, 173, 177, 181, 185, 189, 193, 197, 202, 205, 210, 215, 220, 225, 230, 235, 238, 241, 244, 248, 252, 256, 260, 264, 268, 272, 275, 279, 284, 289, 292, 295, 300, 303, 306, 309, 312, 315, 318, 321, 324, 327, 330, 333, 336, 339, 342, 346, 350, 354, 358, 362, 368, 374, 378, 383, 388, 393, 397, 401, 407, 413, 417, 421, 425, 429, 433, 440, 447, 454, 461, 468, 475, 482, 489, 496, 503, 508, 513, 518, 522, 526, 530, 534, 538, 542, 546, 550, 554, 557, 560, 563, 567, 571, 575, 580, 585, 590, 595, 600, 605, 610, 615, 619, 623, 629, 635, 641, 647, 652, 657, 661, 665, 669, 673, 677, 681, 684, 688, 692, 696, 700, 704, 708, 712, 716, 720, 724, 728, 732, 736, 740, 744, 748, 752, 757, 762, 767, 772, 777, 782, 787, 792, 797, 802, 807, 812, 817, 822, 827, 832, 837, 842, 847, 850, 853, 857, 861, 865, 869, 873, 877, 880, 883, 886, 891, 896, 901, 905, 909, 913, 920, 925, 930, 935, 939, 943, 947, 951, 956, 961, 966, 970, 974, 978, 982, 987, 993, 998, 1004, 1009, 1015, 1020, 1026, 1031, 1037, 1042, 1048, 1053, 1059, 1064, 1070, 1075, 1081, 1088, 1093, 1099, 1104, 1110, 1115, 1121, 1126, 1132, 1139, 1145, 1150, 1155, 1160, 1165, 1171, 1176, 1181, 1186, 1191, 1197, 1202, 1207, 1212, 1217, 1224, 1230, 1236, 1242, 1248, 1254, 1259, 1264, 1269, 1274, 1281, 1287, 1293, 1299, 1305, 1310, 1315, 1320, 1326, 1332, 1337, 1342, 1348, 1353, 1358, 1362, 1366, 1371, 1375, 1382, 1389, 1394, 1398, 1404, 1410, 1416, 1422, 1427, 1431, 1435, 1441, 1446, 1452, 1457, 1464, 1470, 1474, 1481, 1486, 1490, 1494, 1501, 1506, 1513, 1521, 1529, 1537, 1546, 1551, 1555, 1561, 1567, 1573, 1580, 1585, 1590, 1595, 1598, 1604, 1610, 1616, 1621, 1626, 1631, 1636, 1643, 1650, 1657, 1664, 1671, 1678, 1685, 1692, 1699, 1706, 1713, 1720, 1727, 1734, 1741, 1748, 1755, 1762, 1769, 1776, 1783, 1790, 1797, 1804, 1811, 1818, 1825, 1832, 1839, 1846, 1853, 1860, 1867, 1874, 1881, 1888, 1895, 1902, 1909, 1916, 1923, 1930, 1937, 1944, 1951, 1958, 1965, 1972, 1976, 1980, 1984, 1988, 1992, 1996, 1999, 2003, 2007, 2011, 2016, 2024, 2027, 2031, 2036, 2041, 2046, 2051, 2055, 2061, 2066, 2070, 2077, 2084, 2091, 2096, 2100, 2104, 2108, 2111, 2116, 2120, 2126, 2131, 2135, 2139, 2143, 2147, 2151, 2155, 2159, 2163, 2168, 2172, 2176, 2181, 2185, 2190, 2195, 2200, 2204, 2209, 2214, 2219, 2224, 2230, 2236, 2241, 2246, 2251, 2256, 2261, 2266, 2272, 2278, 2286, 2294, 2302, 2310, 2318, 2326, 2334, 2342, 2350, 2358, 2366, 2374, 2382, 2390, 2398, 2406, 2415, 2424, 2433, 2442, 2451, 2460, 2465, 2470, 2475, 2480, 2484, 2491, 2500, 2506, 2514, 2521, 2530, 2538, 2543, 2548, 2553, 2558, 2563, 2568, 2573, 2578, 2584, 2590, 2595, 2602, 2608, 2614, 2621, 2627, 2633, 2641, 2649, 2655, 2662, 2669, 2675, 2679, 2686, 2691, 2696, 2702, 2708, 2713, 2718, 2723, 2728, 2732, 2736, 2744, 2752, 2760, 2765, 2770, 2775, 2781, 2787, 2794, 2801, 2806, 2811, 2816, 2822, 2828, 2835, 2842, 2847, 2851, 2856, 2861, 2866, 2873, 2880, 2887, 2894, 2901, 2908, 2914, 2919, 2924, 2931, 2938, 2945, 2950, 2955, 2960, 2966, 2973, 2978, 2986, 2994, 3002, 3009, 3014, 3020, 3026, 3032, 3038, 3045, 3051, 3057, 3063, 3069, 3077, 3084, 3091, 3097, 3103, 3110, 3113, 3119, 3126, 3132, 3139, 3145, 3151, 3156, 3161, 3166, 3171, 3176, 3181, 3186, 3191, 3196, 3201, 3206, 3211, 3216, 3222, 3228, 3235, 3242, 3247, 3253, 3262, 3271, 3281, 3290, 3299, 3308, 3318, 3327, 3331, 3336, 3341, 3348, 3355, 3361, 3367, 3373, 3379, 3385, 3391, 3398, 3403, 3408, 3413, 3418, 3425, 3432, 3440, 3448, 3456, 3464, 3469, 3474, 3479, 3484, 3489, 3494, 3499, 3504, 3509, 3515, 3519, 3525, 3531, 3537, 3547, 3557, 3567, 3578, 3583, 3589, 3595, 3601, 3611, 3617, 3627, 3633, 3648, 3655, 3662, 3669, 3676, 3682, 3691, 3696, 3703, 3710, 3717, 3724, 3731, 3738, 3745, 3752, 3757, 3763, 3769, 3776, 3780}

func (i Instr) String() string {
	if i < 0 || i >= Instr(len(_Instr_index)-1) {
		return fmt.Sprintf("Instr(%d)", i)
	}
	return _Instr_name[_Instr_index[i]:_Instr_index[i+1]]
}

const _InstrOpType_name = "INVALID_OPINTEGER_OPXMM_OP"

var _InstrOpType_index = [...]uint8{0, 10, 20, 26}

func (i InstrOpType) String() string {
	if i < 0 || i >= InstrOpType(len(_InstrOpType_index)-1) {
		return fmt.Sprintf("InstrOpType(%d)", i)
	}
	return _InstrOpType_name[_InstrOpType_index[i]:_InstrOpType_index[i+1]]
}

const _InstructionType_name = "I_ADDI_ANDI_CMPI_CVT_FLOAT2INTI_CVT_INT2FLOATI_CVT_FLOAT2FLOATI_DIVI_IMULI_IDIVI_LEAI_MOVI_MOVBSXI_MOVWSXI_MOVLSXI_MOVBZXI_MOVWZXI_MOVLZXI_MULI_ORI_PADDI_PANDI_PANDNI_PCMPEQI_PCMPGTI_PIMULI_PMULI_PORI_PSLLI_PSRAI_PSRLI_PSUBI_PXORI_PMOVI_SALI_SARI_SHLI_SHRI_SUBI_XOR"

var _InstructionType_index = [...]uint16{0, 5, 10, 15, 30, 45, 62, 67, 73, 79, 84, 89, 97, 105, 113, 121, 129, 137, 142, 146, 152, 158, 165, 173, 181, 188, 194, 199, 205, 211, 217, 223, 229, 235, 240, 245, 250, 255, 260, 265}

func (i InstructionType) String() string {
	if i < 0 || i >= InstructionType(len(_InstructionType_index)-1) {
		return fmt.Sprintf("InstructionType(%d)", i)
	}
	return _InstructionType_name[_InstructionType_index[i]:_InstructionType_index[i+1]]
}

const _XmmData_name = "XMM_INVALIDXMM_F32XMM_F64XMM_4X_F32XMM_2X_F64"

var _XmmData_index = [...]uint8{0, 11, 18, 25, 35, 45}

func (i XmmData) String() string {
	if i < 0 || i >= XmmData(len(_XmmData_index)-1) {
		return fmt.Sprintf("XmmData(%d)", i)
	}
	return _XmmData_name[_XmmData_index[i]:_XmmData_index[i+1]]
}