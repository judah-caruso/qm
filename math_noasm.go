//go:build !amd64

package qm

// Sqrt calculates the approximate square root of x.
func Sqrt(x float32) float32 {
	var (
		z = float32(1)
		l = 11
	)

	for i := 0; i < l && Abs(z*z-x) > 0.0001; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

// InvSqrt calculates the approximate inverse square root of x.
func InvSqrt(x float32) float32 {
	return 1 / Sqrt(x)
}
