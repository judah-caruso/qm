package fx_test

import (
	"github.com/judah-caruso/qm/fx"
	"math"
	"testing"
)

func TestSin(t *testing.T) {
	const pi32 = float32(math.Pi)
	cases := []float32{
		0,
		pi32 / 2,
		5 * pi32 / 2,
		pi32, 0,
		3 * pi32 / 2,
		2 * pi32,
		pi32 / 6.0,
		.01598529,
	}

	for i, c := range cases {
		ours := fx.F(c)
		theirs := float64(c)

		ourRes := fx.Sin(ours).Float64()
		theirRes := math.Sin(theirs)

		diff := math.Abs(ourRes - theirRes)
		if diff >= 0.03 {
			t.Fatalf("%d: fx.Sin(%v)=%v deviated too much from math.Sin(%.2f)=%.2f (+- %.4f)", i, ours, ourRes, theirs, theirRes, diff)
		}
	}
}
