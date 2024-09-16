package qm

import (
	"github.com/judah-caruso/qm/fx"
)

type Color Vec4

func Rgbf(r, g, b float32) Color {
	return Rgbaf(r, g, b, 1.0)
}

func Rgbaf(r, g, b, a float32) Color {
	return Color{
		fx.F(r),
		fx.F(g),
		fx.F(b),
		fx.F(a),
	}
}

func (c Color) Channels() (r, g, b, a fx.T) {
	return Vec4.Elements(Vec4(c))
}

func (c Color) String() string {
	return Vec4.String(Vec4(c))
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c[R].Int()), uint32(c[G].Int()), uint32(c[B].Int()), uint32(c[A].Int())
}
