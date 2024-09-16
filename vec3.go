package qm

import (
	"fmt"
	"github.com/judah-caruso/qm/fx"
)

type Vec3 [3]fx.T

func V3f(x, y, z float32) Vec3 {
	return Vec3{fx.F(x), fx.F(y), fx.F(z)}
}

func V3i(x, y, z int) Vec3 {
	return Vec3{fx.I(x), fx.I(y), fx.I(z)}
}

func (v Vec3) Elements() (x, y, z fx.T) {
	return v[X], v[Y], v[Z]
}

func (v Vec3) String() string {
	return "(" + v[X].String() + ", " + v[Y].String() + ", " + v[Z].String() + ")"
}

func (l Vec3) Eq(r Vec3) bool {
	return l[X] == r[X] && l[Y] == r[Y] && l[Z] == r[Z]
}

func (v Vec3) Swizzle(el ...VecElementIndex) Vec3 {
	if len(el) != 3 {
		panic(fmt.Sprintf("invalid swizzle of Vec3 (given %d elements, required 3)", len(el)))
	}

	px := el[0]
	if px < 0 || px > 2 {
		panic(fmt.Sprintf("invalid swizzle of Vec3 (element %d does not exist)", px))
	}

	py := el[1]
	if py < 0 || py > 2 {
		panic(fmt.Sprintf("invalid swizzle of Vec3 (element %d does not exist)", py))
	}

	pz := el[2]
	if pz < 0 || pz > 2 {
		panic(fmt.Sprintf("invalid swizzle of Vec3 (element %d does not exist)", pz))
	}

	return Vec3{v[px], v[py], v[pz]}
}

func (l Vec3) Add(r Vec3) Vec3 {
	return Vec3{
		fx.Add(l[X], r[X]),
		fx.Add(l[Y], r[Y]),
		fx.Add(l[Z], r[Z]),
	}
}

func (l Vec3) Sub(r Vec3) Vec3 {
	return Vec3{
		fx.Sub(l[X], r[X]),
		fx.Sub(l[Y], r[Y]),
		fx.Sub(l[Z], r[Z]),
	}
}

func (l Vec3) Mul(r Vec3) Vec3 {
	return Vec3{
		fx.Mul(l[X], r[X]),
		fx.Mul(l[Y], r[Y]),
		fx.Mul(l[Z], r[Z]),
	}
}

func (l Vec3) Mulf(r fx.T) Vec3 {
	return Vec3{
		fx.Mul(l[X], r),
		fx.Mul(l[Y], r),
		fx.Mul(l[Z], r),
	}
}

func (l Vec3) Div(r Vec3) Vec3 {
	return Vec3{
		fx.Div(l[X], r[X]),
		fx.Div(l[Y], r[Y]),
		fx.Div(l[Z], r[Z]),
	}
}

func (l Vec3) Dot(r Vec3) fx.T {
	return fx.Add(
		fx.Add(fx.Mul(l[X], r[X]), fx.Mul(l[Y], r[Y])),
		fx.Mul(l[Z], r[Z]),
	)
}

func (v Vec3) Mag() fx.T {
	return fx.Sqrt(v.MagSqr())
}

func (v Vec3) MagSqr() fx.T {
	return v.Dot(v)
}

// Lerp linearly interpolates from a to b via t. Returns a new Vec3.
func (a Vec3) Lerp(b Vec3, t fx.T) Vec3 {
	return Vec3.Add(
		a.Mulf(fx.Sub(fx.One(), t)),
		b.Mulf(t),
	)
}

func (v Vec3) Normalize() Vec3 {
	return v.Mulf(fx.InvSqrt(v.Dot(v)))
}

// Negate negates each component. Returns a new Vec3.
func (v Vec3) Negate() Vec3 {
	return Vec3{
		fx.Negate(v[X]),
		fx.Negate(v[Y]),
		fx.Negate(v[Z]),
	}
}

// Invert calculates the inverse of a Vec3. Returns a new Vec3.
func (v Vec3) Invert() Vec3 {
	return Vec3{
		fx.Div(fx.One(), v[X]),
		fx.Div(fx.One(), v[Y]),
		fx.Div(fx.One(), v[Z]),
	}
}

// Abs calculates the absolute value of a Vec2. Returns a new Vec2.
func (v Vec3) Abs() Vec3 {
	return Vec3{
		fx.Abs(v[X]),
		fx.Abs(v[Y]),
		fx.Abs(v[Z]),
	}
}
