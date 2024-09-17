package qm

import (
	"github.com/judah-caruso/qm/fx"
)

type Color Vec4

func Rgbf(r, g, b float32) Color {
	return Rgbaf(r, g, b, 1.0)
}

func Rgbaf(r, g, b, a float32) Color {
	return Color{fx.F(r), fx.F(g), fx.F(b), fx.F(a)}
}

func (c Color) Channels() (r, g, b, a fx.T) {
	return Vec4.Elements(Vec4(c))
}

func (c Color) String() string {
	return Vec4.String(Vec4(c))
}

func (l Color) Eq(r Color) bool {
	return Vec4.Eq(Vec4(r), Vec4(l))
}

func (c Color) Swizzle(el ...VecElementIndex) Color {
	return Color(Vec4.Swizzle(Vec4(c), el...))
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c[R].Float() * 255)
	r |= r << 8

	g = uint32(c[G].Float() * 255)
	g |= g << 8

	b = uint32(c[B].Float() * 255)
	b |= b << 8

	a = uint32(c[A].Float() * 255)
	a |= a << 8

	return
}
