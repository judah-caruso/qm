package qm

import (
	"strconv"
	"unsafe"
)

type xx = float64

func floatString(f float32) string {
	return strconv.FormatFloat(float64(f), 'd', 2, 32)
}

func ptrcastFloatToUint(f float32) uint32 {
	return *(*uint32)(unsafe.Pointer(&f))
}

func ptrcastUintToFloat(u uint32) float32 {
	return *(*float32)(unsafe.Pointer(&u))
}
