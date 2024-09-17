package qm

import (
	"fmt"
	"github.com/judah-caruso/qm/fx"
)

type Vec4 [4]fx.T

func V4f(x, y, z, w float32) Vec4 {
	return Vec4{fx.F(x), fx.F(y), fx.F(z), fx.F(w)}
}

func V4i(x, y, z, w int) (out Vec4) {
	return Vec4{fx.I(x), fx.I(y), fx.I(z), fx.I(w)}
}

func (v Vec4) Elements() (x, y, z, w fx.T) {
	return v[X], v[Y], v[Z], v[W]
}

func (v Vec4) String() string {
	return "(" + v[X].String() + ", " + v[Y].String() + ", " + v[Z].String() + ", " + v[W].String() + ")"
}

func (l Vec4) Eq(r Vec4) bool {
	return l[X] == r[X] && l[Y] == r[Y] && l[Z] == r[Z] && l[W] == r[W]
}

func (v Vec4) Swizzle(el ...VecElementIndex) Vec4 {
	if len(el) != 4 {
		panic(fmt.Sprintf("invalid swizzle of Vec4 (given %d elements, required 4)", len(el)))
	}

	px := el[0]
	if px < 0 || px > 3 {
		panic(fmt.Sprintf("invalid swizzle of Vec4 (element %d does not exist)", px))
	}

	py := el[1]
	if py < 0 || py > 3 {
		panic(fmt.Sprintf("invalid swizzle of Vec4 (element %d does not exist)", py))
	}

	pz := el[2]
	if pz < 0 || pz > 3 {
		panic(fmt.Sprintf("invalid swizzle of Vec4 (element %d does not exist)", pz))
	}

	pw := el[3]
	if pw < 0 || pw > 3 {
		panic(fmt.Sprintf("invalid swizzle of Vec4 (element %d does not exist)", pw))
	}

	return Vec4{v[px], v[py], v[pz], v[pw]}
}
