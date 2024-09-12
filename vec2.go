package qm

import "fmt"

// Vec2 represents a two-component vector.
type Vec2 [2]float32

// Elements returns each element of the Vec2 in order.
func (v Vec2) Elements() (float32, float32) {
	return v[X], v[Y]
}

func (v Vec2) String() string {
	return "(" + floatString(v[X]) + ", " + floatString(v[Y]) + ")"
}

// Swizzle is a slower, but possibly clearer alternative to swizzling the Vec2 via array access.
// Swizzle panics if given invalid or too many elements.
func (a Vec2) Swizzle(el ...VecElementIndex) Vec2 {
	if len(el) != 2 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (given %d elements)", len(el)))
	}

	if el[0] > 1 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (element %d does not exist)", el[0]))
	}

	if el[1] > 1 {
		panic(fmt.Sprintf("invalid swizzle of Vec2 (element %d does not exist)", el[1]))
	}

	return Vec2{a[el[0]], a[el[1]]}
}

// Eq compares two Vec2 with strict equality.
func (l Vec2) Eq(r Vec2) bool {
	return l[X] == r[X] && l[Y] == r[Y]
}

// KindaEq compares two Vec2 with approximate equality via Epsilon.
func (l Vec2) KindaEq(r Vec2) bool {
	return Abs(l[X]-r[X]) <= Epsilon && Abs(l[Y]-r[Y]) <= Epsilon
}

// Add component-wise adds two Vec2. Returns a new Vec2.
func (l Vec2) Add(r Vec2) Vec2 {
	return add_array_2_float32(l, r)
}

// Addf component-wise adds Vec2 and float32. Returns a new Vec2.
func (l Vec2) Addf(r float32) Vec2 {
	return add_array_2_float32(l, Vec2{r, r})
}

// Sub component-wise subtracts two Vec2. Returns a new Vec2.
func (l Vec2) Sub(r Vec2) Vec2 {
	return sub_array_2_float32(l, r)
}

// Subf component-wise subtracts Vec2 and float32. Returns a new Vec2.
func (l Vec2) Subf(r float32) Vec2 {
	return sub_array_2_float32(l, Vec2{r, r})
}

// Mul component-wise multiplies two Vec2. Returns a new Vec2.
func (l Vec2) Mul(r Vec2) Vec2 {
	return mul_array_2_float32(l, r)
}

// Mulf component-wise multiplies Vec2 and float32. Returns a new Vec2.
func (l Vec2) Mulf(r float32) Vec2 {
	return mul_array_2_float32(l, Vec2{r, r})
}

// Div component-wise divides two Vec2. Returns a new Vec2.
func (l Vec2) Div(r Vec2) Vec2 {
	return div_array_2_float32(l, r)
}

// Divf component-wise divides Vec2 and float32. Returns a new Vec2.
func (l Vec2) Divf(r float32) Vec2 {
	return div_array_2_float32(l, Vec2{r, r})
}

// Dot calculates the dot-product of two Vec2.
func (l Vec2) Dot(r Vec2) float32 {
	return l[X]*r[X] + l[Y]*r[Y]
}

// Mag calculates the magnitude (length) of a Vec2.
func (l Vec2) Mag() float32 {
	return Sqrt(l.MagSqr())
}

// MagSqr calculates the square magnitude (length) of a Vec2.
func (l Vec2) MagSqr() float32 {
	return l.Dot(l)
}

// Lerp linearly interpolates from a to b via t. Returns a new Vec2.
func (a Vec2) Lerp(b Vec2, t float32) Vec2 {
	return Vec2.Add(a.Mulf(1.0-t), b.Mulf(t))
}

// Rotate rotates a Vec2 via an angle specified in radians. Returns a new Vec2.
func (a Vec2) Rotate(angle float32) Vec2 {
	sin := Sin(angle)
	cos := Cos(angle)
	return Vec2{a[X]*cos - a[Y]*sin, a[X]*sin + a[Y]*cos}
}

// Normalize returns a new Vec2 with the same direction, but a magnitude of 1.
func (a Vec2) Normalize() Vec2 {
	return a.Mulf(InvSqrt(a.Dot(a)))
}

// Negate negates each component. Returns a new Vec2.
func (a Vec2) Negate() Vec2 {
	return Vec2{-a[X], -a[Y]}
}

// Invert calculates the inverse of a Vec2. Returns a new Vec2.
func (a Vec2) Invert() Vec2 {
	return Vec2{1.0 / a[X], 1.0 / a[Y]}
}

// Abs calculates the absolute value of a Vec2. Returns a new Vec2.
func (a Vec2) Abs() Vec2 {
	return Vec2{Abs(a[X]), Abs(a[Y])}
}

// Distance calculates the distance between two Vec2.
func (a Vec2) Distance(b Vec2) float32 {
	return Sqrt(Square(a[X]-b[X]) + Square(a[Y]-b[Y]))
}

// Reflect calculates a reflected Vec2 via a normal. Returns a new Vec2.
func (a Vec2) Reflect(normal Vec2) Vec2 {
	return a.Sub(normal.Mulf(2.0).Mulf(a.Dot(normal)))
}
