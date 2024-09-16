package qm

import "github.com/judah-caruso/qm/fx"

type Rect struct {
	Min Vec2
	Max Vec2
}

func Rectf(minx, miny, maxx, maxy float32) Rect {
	return Rect{
		Min: V2f(minx, miny),
		Max: V2f(maxx, maxy),
	}
}

func Recti(minx, miny, maxx, maxy int) Rect {
	return Rect{
		Min: V2i(minx, miny),
		Max: V2i(maxx, maxy),
	}
}

func (r Rect) Width() fx.T {
	return r.Max.Sub(r.Min)[X]
}

func (r Rect) Height() fx.T {
	return r.Max.Sub(r.Min)[Y]
}
