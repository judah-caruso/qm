package qm_test

import (
	"github.com/judah-caruso/qm"
	"github.com/judah-caruso/qm/fx"
	"testing"
)

func TestVec2_Ops(t *testing.T) {
	a := qm.Vec2{1, 2}

	{ // addition
		b := a.Add(qm.Vec2{1, 1})
		Expectf(t, b[qm.X] == 2 && b[qm.Y] == 3, "incorrect result: %s", b)
	}
}

func TestVec2_Swizzle(t *testing.T) {
	a := qm.Vec2{1, -2}

	b := a.Swizzle(qm.Y, qm.X)
	Expect(t, b[qm.X] == -2 && b[qm.Y] == 1)

	c := a.Swizzle(qm.X, qm.X)
	Expect(t, c[qm.X] == 1 && c[qm.Y] == 1)

	d := a.Swizzle(qm.Y, qm.Y)
	Expect(t, d[qm.X] == -2 && d[qm.Y] == -2)
}

func TestVec2_SwizzleNoMethod(t *testing.T) {
	a := qm.Vec2{1, -2}

	b := qm.Vec2{a[qm.Y], a[qm.X]}
	Expect(t, b[qm.X] == -2 && b[qm.Y] == 1)

	c := qm.Vec2{a[qm.X], a[qm.X]}
	Expect(t, c[qm.X] == 1 && c[qm.Y] == 1)

	d := qm.Vec2{a[qm.Y], a[qm.Y]}
	Expect(t, d[qm.X] == -2 && d[qm.Y] == -2)
}

func BenchmarkVec2_QmAdd(b *testing.B) {
	a := qm.Vec2{1, 2}
	for i := 0; i < b.N; i++ {
		a = a.Add(qm.Vec2{fx.I(i), fx.I(i - 1)})
	}
}

func BenchmarkVec2_GoAdd(b *testing.B) {
	a := [2]float32{1, 2}
	for i := 0; i < b.N; i++ {
		a = puregoVec2_Add(a, [2]float32{float32(i), float32(i - 1)})
	}
}

func puregoVec2_Add(l, r [2]float32) [2]float32 {
	return [2]float32{l[0] + r[0], l[1] + r[1]}
}
