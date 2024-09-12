//go:build arm64

#include "textflag.h"

// func add_arrayof_2float32(l [2]float32, r [2]float32) [2]float32
TEXT ·add_arrayof_2float32(SB), NOSPLIT|NOFRAME, $0-24
    FMOVS r_0+8(FP), F0
    FMOVS l+0(FP), F1
    FADDS F0, F1, F0
    FMOVS F0, r_0+16(FP)
    FMOVS r_0+12(FP), F0
    FMOVS l+4(FP), F1
    FADDS F0, F1, F0
    FMOVS F0, r_0+20(FP)
	RET (R30)

// func sub_arrayof_2float32(l [2]float32, r [2]float32) [2]float32
TEXT ·sub_arrayof_2float32(SB), NOSPLIT|NOFRAME, $0-24
    FMOVS r_0+8(FP), F0
    FMOVS l+0(FP), F1
    FSUBS F0, F1, F0
    FMOVS F0, r_0+16(FP)
    FMOVS r_0+12(FP), F0
    FMOVS l+4(FP), F1
    FSUBS F0, F1, F0
    FMOVS F0, r_0+20(FP)
	RET (R30)

// func mul_arrayof_2float32(l [2]float32, r [2]float32) [2]float32
TEXT ·mul_arrayof_2float32(SB), NOSPLIT|NOFRAME, $0-24
    FMOVS r_0+8(FP), F0
    FMOVS l+0(FP), F1
    FMULS F0, F1, F0
    FMOVS F0, r_0+16(FP)
    FMOVS r_0+12(FP), F0
    FMOVS l+4(FP), F1
    FMULS F0, F1, F0
    FMOVS F0, r_0+20(FP)
	RET (R30)

// func div_arrayof_2float32(l [2]float32, r [2]float32) [2]float32
TEXT ·div_arrayof_2float32(SB), NOSPLIT|NOFRAME, $0-24
    FMOVS r_0+8(FP), F0
    FMOVS l+0(FP), F1
    FDIVS F0, F1, F0
    FMOVS F0, r_0+16(FP)
    FMOVS r_0+12(FP), F0
    FMOVS l+4(FP), F1
    FDIVS F0, F1, F0
    FMOVS F0, r_0+20(FP)
	RET (R30)
