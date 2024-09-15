package fx_test

import (
	"github.com/judah-caruso/qm/fx"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := [][2]fx.T{
		{fx.Zero(), fx.Zero()},
		{fx.One(), fx.One()},
		{fx.I(4), fx.I(2)},
		{fx.I(4225), fx.I(65)},
		{fx.I(-3923), fx.Zero()},
	}

	for _, c := range cases {
		res := fx.Sqrt(c[0])
		if res != c[1] {
			t.Errorf("Sqrt(%v) = %v, wanted %v", c[0], res, c[1])
		}
	}

	for i := 1.0; i < 50.0; i += 0.1 {
		ours := fx.Sqrt(fx.F(i)).Float64()
		theirs := math.Sqrt(i)

		diff := math.Abs(ours - theirs)
		if diff >= 0.003 {
			t.Fatalf("fx.Sqrt(%.2f)=%v deviated too much from math.Sqrt (+- %.4f)", i, ours, diff)
		}
	}
}

func TestInvSqrt(t *testing.T) {
	for i := 1.0; i < 50.0; i += 0.1 {
		ours := fx.InvSqrt(fx.F(i)).Float64()
		theirs := 1.0 / math.Sqrt(i)

		diff := math.Abs(ours - theirs)
		if diff >= 0.003 {
			t.Fatalf("fx.InvSqrt(%.2f)=%v deviated too much from 1.0 / math.Sqrt (+- %.4f)", i, ours, diff)
		}
	}
}

func TestSinCos(t *testing.T) {
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

	for i, c := range cases {
		ours := fx.F(c)
		theirs := float64(c)

		ourRes := fx.Cos(ours).Float64()
		theirRes := math.Cos(theirs)

		diff := math.Abs(ourRes - theirRes)
		if diff >= 0.03 {
			t.Fatalf("%d: fx.Cos(%v)=%v deviated too much from math.Cos(%.2f)=%.2f (+- %.4f)", i, ours, ourRes, theirs, theirRes, diff)
		}
	}
}

func TestAsinAcos(t *testing.T) {
	const pi32 = float32(math.Pi)
	cases := []float32{
		0,
		pi32,
		pi32 / 2,
		pi32 / 3,
		2.0944,
		pi32 / 4,
		2.3562,
		pi32 / 6,
		2.61799,
	}

	for i, c := range cases {
		ours := fx.F(c)
		theirs := float64(c)

		ourRes := fx.Asin(ours).Float64()
		theirRes := math.Asin(theirs)

		diff := math.Abs(ourRes - theirRes)
		if diff >= 0.03 {
			t.Fatalf("%d: fx.Asin(%v)=%v deviated too much from math.Asin(%.2f)=%.2f (+- %.4f)", i, ours, ourRes, theirs, theirRes, diff)
		}
	}

	for i, c := range cases {
		ours := fx.F(c)
		theirs := float64(c)

		ourRes := fx.Acos(ours).Float64()
		theirRes := math.Acos(theirs)

		diff := math.Abs(ourRes - theirRes)
		if diff >= 0.03 {
			t.Fatalf("%d: fx.Acos(%v)=%v deviated too much from math.Acos(%.2f)=%.2f (+- %.4f)", i, ours, ourRes, theirs, theirRes, diff)
		}
	}
}
