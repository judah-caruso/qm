package qm_test

import (
	"github.com/judah-caruso/qm"
	"testing"
)

func TestVec3_Ops(t *testing.T) {
	a := qm.V3i(1, 2, 3)

	{ // addition
		b := a.Add(qm.V3i(1, 1, 1))
		Expectf(t, b[qm.X].Int() == 2 && b[qm.Y].Int() == 3 && b[qm.Z].Int() == 4, "incorrect result: %s", b)
	}

	{ // subtraction
		b := a.Sub(qm.V3i(1, 1, 1))
		Expectf(t, b[qm.X].Int() == 0 && b[qm.Y].Int() == 1 && b[qm.Z].Int() == 2, "incorrect result: %s", b)
	}

	{ // multiplication
		b := a.Mul(qm.V3f(2, 0.5, 10))
		Expectf(t, b[qm.X].Int() == 2 && b[qm.Y].Int() == 1 && b[qm.Z].Int() == 30, "incorrect result: %s", b)
	}

	{ // division
		b := a.Div(qm.V3i(2, 2, 2))
		Expectf(t, b[qm.X].Float() == 0.5 && b[qm.Y].Int() == 1 && b[qm.Z].Float() == 1.5, "incorrect result: %s", b)
	}
}

func TestVec3_Swizzle(t *testing.T) {
	a := qm.V3i(1, -2, 3)

	b := a.Swizzle(qm.Y, qm.X, qm.Z)
	Expect(t, b[qm.X].Int() == -2 && b[qm.Y].Int() == 1 && b[qm.Z].Int() == 3)

	c := a.Swizzle(qm.X, qm.X, qm.X)
	Expect(t, c[qm.X].Int() == 1 && c[qm.Y].Int() == 1 && c[qm.Z].Int() == 1)

	d := a.Swizzle(qm.Y, qm.Y, qm.Y)
	Expect(t, d[qm.X].Int() == -2 && d[qm.Y].Int() == -2 && d[qm.Z].Int() == -2)

	e := a.Swizzle(qm.Z, qm.Z, qm.Z)
	Expect(t, e[qm.X].Int() == 3 && e[qm.Y].Int() == 3 && e[qm.Z].Int() == 3)
}

func BenchmarkVec3_QmAdd(b *testing.B) {
	a := qm.V3i(1, 2, 3)
	for i := 0; i < b.N; i++ {
		a = a.Add(qm.V3i(i, i-1, i+1))
	}
}

func BenchmarkVec3_GoAdd(b *testing.B) {
	a := [3]float32{1, 2, 3}
	for i := 0; i < b.N; i++ {
		a = puregoVec3_Add(a, [3]float32{float32(i), float32(i - 1), float32(i + 1)})
	}
}

func puregoVec3_Add(l, r [3]float32) [3]float32 {
	return [3]float32{l[0] + r[0], l[1] + r[1], l[2] + r[2]}
}
