package qm_test

import (
	"github.com/judah-caruso/qm"
	"testing"
)

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
