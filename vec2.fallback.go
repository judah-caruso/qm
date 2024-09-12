//go:build !amd64

package qm

func add_array_2_float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] + r[0], l[1] + r[1]}
}

func sub_array_2_float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] - r[0], l[1] - r[1]}
}

func mul_array_2_float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] * r[0], l[1] * r[1]}
}

func div_array_2_float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] / r[0], l[1] / r[1]}
}
