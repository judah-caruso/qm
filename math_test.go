package qm_test

import (
	"github.com/judah-caruso/qm"
	"math"
	"math/rand/v2"
	"testing"
)

func TestSin(t *testing.T) {
	for i := -100_000.0; i < 100_000.0; i += 0.01 {
		ours := qm.Sin(float32(i))
		theirs := float32(math.Sin(i))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.05, "qm.Sin(%.4f) deviated from math.Sin too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
		i += 0.01
	}
}

func TestCos(t *testing.T) {
	for i := -100_000.0; i < 100_000.0; i += 0.01 {
		ours := qm.Cos(float32(i))
		theirs := float32(math.Cos(i))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.05, "qm.Cos(%.4f) deviated from math.Cos too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestAcos(t *testing.T) {
	for i := -qm.Pi / 2; i < qm.Pi/2; i += 0.1 {
		ours := qm.Acos(i)
		theirs := float32(math.Acos(float64(i)))
		if math.IsNaN(float64(theirs)) || math.IsNaN(float64(ours)) {
			continue
		}

		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.0001, "qm.Acos(%.4f) deviated from math.Acos too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestAtan(t *testing.T) {
	for i := -qm.Pi / 2; i < qm.Pi/2; i += 0.1 {
		ours := qm.Atan(i)
		theirs := float32(math.Atan(float64(i)))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.03, "qm.Atan(%.2f) deviated from math.Atan too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestAtan2(t *testing.T) {
	for i := -qm.Pi / 2; i < qm.Pi/2; i += 0.1 {
		ours := qm.Atan2(-i, i)
		theirs := float32(math.Atan2(float64(-i), float64(i)))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.03, "qm.Atan2(%.2f) deviated from math.Atan2 too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestLog(t *testing.T) {
	for i := 1.0; i < 10_000.0; i += 0.1 {
		ours := qm.Log(float32(i))
		theirs := float32(math.Log(i))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.03, "qm.Log(%.2f) deviated from math.Log too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestSqrt(t *testing.T) {
	for i := 0.0; i < 100_000.0; i += 0.1 {
		ours := qm.Sqrt(float32(i))
		theirs := float32(math.Sqrt(i))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.1, "qm.Sqrt(%.2f) deviated from math.Sqrt too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestInvSqrt(t *testing.T) {
	for i := 0.1; i < 100_000.0; i += 0.1 {
		ours := qm.InvSqrt(float32(i))
		theirs := float32(1.0 / math.Sqrt(i))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.1, "qm.InvSqrt(%.2f) deviated from 1.0 / math.Sqrt too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff)
	}
}

func TestExp(t *testing.T) {
	for range 100_0000 {
		x := float32(rand.NormFloat64())

		ours := qm.Exp(x)
		theirs := float32(math.Exp(float64(x)))
		diff := qm.Abs(ours-theirs) / theirs
		Expectf(t, diff <= 0.03, "qm.Exp(%.2f) deviated from math.Exp too much! %.4f vs %.4f (+- %.4f)", x, ours, theirs, diff)
	}

	Expectf(t, qm.Exp(-90.0) == 0, "qm.Exp(-90.0) was expected to be 0, instead was %.4f", qm.Exp(-90.0))
	Expectf(t, qm.Exp(90.0) == float32(math.Inf(1)), "qm.Exp(90.0) was expected to be +Inf, instead was %.4f", qm.Exp(90.0))
}
