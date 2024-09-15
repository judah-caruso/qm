package qm

import "github.com/judah-caruso/qm/fx"

type Vec3 [3]fx.T

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

func (l Vec3) Div(r Vec3) Vec3 {
	return Vec3{
		fx.Div(l[X], r[X]),
		fx.Div(l[Y], r[Y]),
		fx.Div(l[Z], r[Z]),
	}
}
