//go:build !(amd64 || arm64)

package qm

func add_arrayof_2float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] + r[0], l[1] + r[1]}
}

func sub_arrayof_2float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] - r[0], l[1] - r[1]}
}

func mul_arrayof_2float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] * r[0], l[1] * r[1]}
}

func div_arrayof_2float32(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] / r[0], l[1] / r[1]}
}
