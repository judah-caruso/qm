package qm

import (
	"fmt"
	"github.com/judah-caruso/qm/fx"
)

// Vec2 represents a two-component vector.
type Vec2 [2]fx.T

// V2f creates a Vec2 from two float32.
func V2f(x, y float32) Vec2 {
	return Vec2{fx.F(x), fx.F(y)}
}

// V2i creates a Vec2 from two int.
func V2i(x, y int) Vec2 {
	return Vec2{fx.I(x), fx.I(y)}
}

// Elements returns each element of the Vec2 in order.
func (v Vec2) Elements() (x, y fx.T) {
	return v[X], v[Y]
}

func (v Vec2) String() string {
	return "(" + v[X].String() + ", " + v[Y].String() + ")"
}

// Swizzle is a slower, but possibly clearer alternative to swizzling the Vec2 via array access.
// Swizzle panics if given invalid or too many elements.
func (v Vec2) Swizzle(el ...VecElementIndex) Vec2 {
	if len(el) != 2 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (given %d elements, required 2)", len(el)))
	}

	px := el[0]
	if px < 0 || px > 1 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (element %d does not exist)", px))
	}

	py := el[1]
	if py < 0 || py > 1 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (element %d does not exist)", py))
	}

	return Vec2{v[px], v[py]}
}

// Eq compares two Vec2 with strict equality.
func (l Vec2) Eq(r Vec2) bool {
	return l[X] == r[X] && l[Y] == r[Y]
}

// Add component-wise adds two Vec2. Returns a new Vec2.
func (l Vec2) Add(r Vec2) Vec2 {
	return Vec2{
		fx.Add(l[X], r[X]),
		fx.Add(l[Y], r[Y]),
	}
}

// Addf component-wise adds Vec2 and fixed-point number. Returns a new Vec2.
func (l Vec2) Addf(r fx.T) Vec2 {
	return Vec2{
		fx.Add(l[X], r),
		fx.Add(l[Y], r),
	}
}

// Sub component-wise subtracts two Vec2. Returns a new Vec2.
func (l Vec2) Sub(r Vec2) Vec2 {
	return Vec2{
		fx.Sub(l[X], r[X]),
		fx.Sub(l[Y], r[Y]),
	}
}

// Subf component-wise subtracts Vec2 and fixed-point number. Returns a new Vec2.
func (l Vec2) Subf(r fx.T) Vec2 {
	return Vec2{
		fx.Sub(l[X], r),
		fx.Sub(l[Y], r),
	}
}

// Mul component-wise multiplies two Vec2. Returns a new Vec2.
func (l Vec2) Mul(r Vec2) Vec2 {
	return Vec2{
		fx.Mul(l[X], r[X]),
		fx.Mul(l[Y], r[Y]),
	}
}

// Mulf component-wise multiplies Vec2 and fixed-point number. Returns a new Vec2.
func (l Vec2) Mulf(r fx.T) Vec2 {
	return Vec2{
		fx.Mul(l[X], r),
		fx.Mul(l[Y], r),
	}
}

// Div component-wise divides two Vec2. Returns a new Vec2.
func (l Vec2) Div(r Vec2) Vec2 {
	return Vec2{
		fx.Div(l[X], r[X]),
		fx.Div(l[Y], r[Y]),
	}
}

// Divf component-wise divides Vec2 and fixed-point number. Returns a new Vec2.
func (l Vec2) Divf(r fx.T) Vec2 {
	return Vec2{
		fx.Div(l[X], r),
		fx.Div(l[Y], r),
	}
}

// Dot calculates the dot-product of two Vec2.
func (l Vec2) Dot(r Vec2) fx.T {
	return fx.Add(fx.Mul(l[X], r[X]), fx.Mul(l[Y], r[Y]))
}

// Mag calculates the magnitude (length) of a Vec2.
func (v Vec2) Mag() fx.T {
	return fx.Sqrt(v.MagSqr())
}

// MagSqr calculates the square magnitude (length) of a Vec2.
func (v Vec2) MagSqr() fx.T {
	return v.Dot(v)
}

// Lerp linearly interpolates from a to b via t. Returns a new Vec2.
func (a Vec2) Lerp(b Vec2, t fx.T) Vec2 {
	return Vec2.Add(
		a.Mulf(fx.Sub(fx.One(), t)),
		b.Mulf(t),
	)
}

// Clamp returns a new Vec2 that is >= min and <= max.
func (v Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2{
		fx.Clamp(v[X], min[X], max[X]),
		fx.Clamp(v[Y], min[Y], max[Y]),
	}
}

// Rotate rotates a Vec2 via an angle specified in radians. Returns a new Vec2.
func (v Vec2) Rotate(angle fx.T) Vec2 {
	sin := fx.Sin(angle)
	cos := fx.Cos(angle)
	return Vec2{
		fx.Sub(fx.Mul(v[X], cos), fx.Mul(v[Y], sin)),
		fx.Add(fx.Mul(v[X], sin), fx.Mul(v[Y], cos)),
	}
}

// Normalize returns a new Vec2 with the same direction, but a magnitude of 1.
func (v Vec2) Normalize() Vec2 {
	return v.Mulf(fx.InvSqrt(v.Dot(v)))
}

// Negate negates each component. Returns a new Vec2.
func (v Vec2) Negate() Vec2 {
	return Vec2{fx.Negate(v[X]), fx.Negate(v[Y])}
}

// Invert calculates the inverse of a Vec2. Returns a new Vec2.
func (v Vec2) Invert() Vec2 {
	return Vec2{
		fx.Div(fx.One(), v[X]),
		fx.Div(fx.One(), v[Y]),
	}
}

// Abs calculates the absolute value of a Vec2. Returns a new Vec2.
func (v Vec2) Abs() Vec2 {
	return Vec2{fx.Abs(v[X]), fx.Abs(v[Y])}
}

// Distance calculates the distance between two Vec2.
func (a Vec2) Distance(b Vec2) fx.T {
	x := fx.Square(fx.Sub(a[X], b[X]))
	y := fx.Square(fx.Sub(a[Y], b[Y]))
	return fx.Sqrt(fx.Add(x, y))
}

// Reflect calculates a reflected Vec2 via a normal. Returns a new Vec2.
func (v Vec2) Reflect(normal Vec2) Vec2 {
	return v.Sub(normal.Mulf(fx.I(2)).Mulf(v.Dot(normal)))
}
