//go:build amd64

package qm

// Sqrt calculates the approximate square root of x.
//
//go:nosplit
//go:noescape
func Sqrt(x float32) float32

// InvSqrt calculates the approximate inverse square root of x.
//
//go:nosplit
//go:noescape
func InvSqrt(x float32) float32
