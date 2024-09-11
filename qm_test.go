package qm_test

import (
	"github.com/judah-caruso/qm"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	cases := []float32{
		0, 1,
		-0, -1,
		+0, +1,
		qm.Pi,
		qm.Tau,
		qm.Pi / qm.Tau,
	}

	for _, c := range cases {
		ours := qm.Abs(c)
		theirs := float32(math.Abs(float64(c)))
		diff := float32(math.Abs(float64(ours - theirs)))
		Expectf(t, diff <= qm.Epsilon, "qm.Abs(%.2f) deviated from math.Abs too much! %.4f vs. %.4f (+- %.4f)", c, ours, theirs, diff)

		ours = qm.Abs(-c)
		theirs = float32(math.Abs(float64(-c)))
		diff = float32(math.Abs(float64(ours - theirs)))
		Expectf(t, diff <= qm.Epsilon, "qm.Abs(-%.2f) deviated from math.Abs too much! %.4f vs. %.4f (+- %.4f)", c, ours, theirs, diff)
	}
}
