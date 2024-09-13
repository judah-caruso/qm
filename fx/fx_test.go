package fx_test

import (
	"github.com/judah-caruso/qm/fx"
	"testing"
)

const shift = 1 << 16

func TestStringification(t *testing.T) {
	cases := []struct {
		fixed    fx.T
		expected string
	}{
		{fx.F(3.14159), "3.1415863"},
		{fx.F(-3.14159), "-3.1415863"},
		{fx.I(1), "1"},
		{fx.I(-1), "-1"},
		{fx.MinFP, "-32768"},
		{fx.MaxFP, "32768"},
	}

	for _, c := range cases {
		str := c.fixed.String()
		if str != c.expected {
			t.Errorf("expected %q, got %q", c.expected, str)
		}
	}
}

func TestAdd(t *testing.T) {
	f := fx.Add(fx.I(2), fx.I(2))
	if f != 4*shift {
		t.Errorf("expected %d, got %d", 4*shift, f)
	}

	f = fx.Addi(fx.I(-1), 1)
	if f != 0 {
		t.Errorf("expected %d, got %d", 0, f)
	}

	f = fx.Addf(fx.F(1.0), 3.5)
	if fx.Round(f) != fx.F(5.0) {
		t.Errorf("expected %.2f, got %.2f", 5.0, f.Float())
	}
}

func TestSub(t *testing.T) {
	f := fx.Sub(fx.I(43), fx.I(42))
	if f != 1*shift {
		t.Errorf("expected %d, got %d", 1*shift, f)
	}
}

func TestMul(t *testing.T) {
	f := fx.Mul(fx.I(3), fx.I(2))
	if f != 6*shift {
		t.Errorf("expected %d, got %d", 6*shift, f)
	}
}

func TestDiv(t *testing.T) {
	f := fx.Div(fx.I(99), fx.I(3))
	if f != 33*shift {
		t.Errorf("expected %d, got %d", 33*shift, f)
	}

	f = fx.Div(fx.I(30_000), fx.I(12))
	if f != 2500*shift {
		t.Errorf("expected %d, got %d", 2500*shift, f)
	}
}

func TestF(t *testing.T) {
	f := fx.F(250 * 2)
	if f.Float() != 500 {
		t.Errorf("expected %d, got %d", 500, f)
	}
}

func TestI(t *testing.T) {
	f := fx.I(250 * 2)
	if f.Int() != 500 {
		t.Errorf("expected %d, got %d", 500, f)
	}
}

func TestRound(t *testing.T) {
	f := fx.Round(fx.Add(fx.I(5), fx.T(shift/2))) // 5.5
	if f != fx.F(6) {
		t.Errorf("expected %d, got %.2f", 6, f.Float())
	}

	f = fx.Round(fx.Sub(fx.I(-5), fx.T(shift/2))) // -5.5
	if f != fx.F(-5) {
		t.Errorf("expected %.2f, got %.2f", -5.0, f.Float())
	}
}
