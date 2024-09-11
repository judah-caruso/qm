package qm_test

import (
	"fmt"
	"github.com/judah-caruso/qm"
	"math"
	"testing"
)

// @note(judah): Sin/Cos don't use Epsilon because their general approximations, so some error is expected.

func TestSin(t *testing.T) {
	var i = float32(-10_000)
	for i < 10_000 {
		ours := qm.Sin(i)
		theirs := float32(math.Sin(float64(i)))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.05, fmt.Sprintf("qm.Sin(%.4f) deviated from math.Sin too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff))
		i += 0.01
	}
}

func TestCos(t *testing.T) {
	var i = float32(-10_000)
	for i < 10_000 {
		ours := qm.Cos(i)
		theirs := float32(math.Cos(float64(i)))
		diff := qm.Abs(ours - theirs)
		Expectf(t, diff <= 0.05, fmt.Sprintf("qm.Cos(%.4f) deviated from math.Cos too much! %.4f vs %.4f (+- %.4f)", i, ours, theirs, diff))
		i += 0.01
	}
}

func BenchmarkQmSinCos(b *testing.B) {
	for i := -(b.N / 2); i < b.N/2; i += 1 {
		s := qm.Sin(float32(i))
		c := qm.Cos(float32(i))
		s, c = c, s
	}
}

func BenchmarkStdSinCos(b *testing.B) {
	for i := -(b.N / 2); i < b.N/2; i += 1 {
		s := math.Sin(float64(i))
		c := math.Cos(float64(i))
		s, c = c, s
	}
}
