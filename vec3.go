package qm

import "github.com/judah-caruso/qm/fx"

type Vec3 [3]fx.T

func V3f(x, y, z float32) Vec3 {
	return Vec3{fx.F(x), fx.F(y), fx.F(z)}
}

func V3i(x, y, z int) Vec3 {
	return Vec3{fx.I(x), fx.I(y), fx.I(z)}
}

func (v Vec3) Elements() (fx.T, fx.T, fx.T) {
	return v[X], v[Y], v[Z]
}

func (v Vec3) String() string {
	return "(" + v[X].String() + ", " + v[Y].String() + ", " + v[Z].String() + ")"
}

func (l Vec3) Eq(r Vec3) bool {
	return l[X] == r[X] && l[Y] == r[Y] && l[Z] == r[Z]
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

// Lerp linearly interpolates from a to b via t. Returns a new Vec3.
func (a Vec3) Lerp(b Vec3, t fx.T) Vec3 {
	return Vec3.Add(
		a.Mulf(fx.Sub(fx.One(), t)),
		b.Mulf(t),
	)
}

// Negate negates each component. Returns a new Vec3.
func (a Vec3) Negate() Vec3 {
	return Vec3{
		fx.Negate(a[X]),
		fx.Negate(a[Y]),
		fx.Negate(a[Z]),
	}
}

// Invert calculates the inverse of a Vec3. Returns a new Vec3.
func (a Vec3) Invert() Vec3 {
	return Vec3{
		fx.Div(fx.One(), a[X]),
		fx.Div(fx.One(), a[Y]),
		fx.Div(fx.One(), a[Z]),
	}
}

// Abs calculates the absolute value of a Vec2. Returns a new Vec2.
func (a Vec3) Abs() Vec3 {
	return Vec3{
		fx.Abs(a[X]),
		fx.Abs(a[Y]),
		fx.Abs(a[Z]),
	}
}
