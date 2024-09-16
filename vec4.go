package qm

import "github.com/judah-caruso/qm/fx"

type Vec4 [4]fx.T

func V4f(x, y, z, w float32) Vec4 {
	return Vec4{fx.F(x), fx.F(y), fx.F(z), fx.F(w)}
}

func V4i(x, y, z, w int) Vec4 {
	return Vec4{fx.I(x), fx.I(y), fx.I(z), fx.I(w)}
}

func (v Vec4) Elements() (x, y, z, w fx.T) {
	return v[X], v[Y], v[Z], v[W]
}

func (v Vec4) String() string {
	return "(" + v[X].String() + ", " + v[Y].String() + ", " + v[Z].String() + ", " + v[W].String() + ")"
}
