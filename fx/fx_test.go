package fx_test

import (
	"github.com/judah-caruso/qm/fx"
	"math"
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
		{fx.MinimumValue(), "-32768"},
		{fx.MaximumValue(), "32768"},
	}

	for _, c := range cases {
		str := c.fixed.String()
		if str != c.expected {
			t.Errorf("expected %q, got %q", c.expected, str)
		}
	}
}

func TestComparison(t *testing.T) {
	cases := []struct {
		lhs, rhs fx.T
		op       func(lhs, rhs fx.T) bool
	}{
		{lhs: fx.F(3.14), rhs: fx.F(1.0), op: fx.T.Gt},
		{lhs: fx.F(3.145), rhs: fx.F(3.14), op: fx.T.GtEq},
		{lhs: fx.F(0.0), rhs: fx.F(-1.0), op: fx.T.Gt},
		{lhs: fx.F(1.0), rhs: fx.F(3.14), op: fx.T.Lt},
		{lhs: fx.F(-1.0), rhs: fx.F(0.0), op: fx.T.LtEq},
		{lhs: fx.F(-1.1), rhs: fx.F(-1.0), op: fx.T.LtEq},
		{lhs: fx.F(1.1), rhs: fx.F(1.1), op: fx.T.Eq},
		{lhs: fx.F(0.0), rhs: fx.F(1.0), op: func(lhs, rhs fx.T) bool {
			return !fx.T.Eq(lhs, rhs)
		}},
	}

	for i, c := range cases {
		if !c.op(c.lhs, c.rhs) {
			t.Errorf("%d: comparsion failed!", i)
		}
	}
}

func TestAdd(t *testing.T) {
	f := fx.Add(fx.I(2), fx.I(2))
	if f.Raw() != 4*shift {
		t.Errorf("expected %d, got %d", 4*shift, f)
	}

	f = fx.Addi(fx.I(-1), 1)
	if f.Raw() != 0 {
		t.Errorf("expected %d, got %d", 0, f)
	}

	f = fx.Addf(fx.F(1.0), 3.5)
	if fx.Round(f) != fx.F(5.0) {
		t.Errorf("expected %.2f, got %.2f", 5.0, f.Float())
	}
}

func TestSub(t *testing.T) {
	f := fx.Sub(fx.I(43), fx.I(42))
	if f.Raw() != 1*shift {
		t.Errorf("expected %d, got %d", 1*shift, f)
	}
}

func TestMul(t *testing.T) {
	f := fx.Mul(fx.I(3), fx.I(2))
	if f.Raw() != 6*shift {
		t.Errorf("expected %d, got %d", 6*shift, f)
	}
}

func TestDiv(t *testing.T) {
	f := fx.Div(fx.I(99), fx.I(3))
	if f.Raw() != 33*shift {
		t.Errorf("expected %d, got %d", 33*shift, f)
	}

	f = fx.Div(fx.I(30_000), fx.I(12))
	if f.Raw() != 2500*shift {
		t.Errorf("expected %d, got %d", 2500*shift, f)
	}
}

func TestF(t *testing.T) {
	f := fx.F(250.0 * 2.0)
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
	f := fx.Round(fx.Add(fx.I(5), fx.F(0.5))) // 5.5
	if f != fx.F(6.0) {
		t.Errorf("expected %d, got %.2f", 6, f.Float())
	}

	f = fx.Round(fx.Sub(fx.I(-5), fx.F(0.5))) // -5.5
	if f != fx.F(-5.0) {
		t.Errorf("expected %.2f, got %.2f", -5.0, f.Float())
	}
}

func TestClamp(t *testing.T) {
	cases := []struct {
		in, min, max float32
		out          fx.T
	}{
		{in: 3.14, min: -1, max: 1, out: fx.One()},
		{in: 1.25, min: -0.5, max: 0.5, out: fx.F(0.5)},
		{in: -1, min: 0, max: 1, out: fx.Zero()},
		{in: 0, min: 0, max: 1, out: fx.Zero()},
		{in: 1, min: 0, max: 1, out: fx.One()},
		{in: 0.5, min: 0, max: 1, out: fx.F(0.5)},
	}

	for _, c := range cases {
		infp := fx.F(c.in)
		minfp := fx.F(c.min)
		maxfp := fx.F(c.max)

		res := fx.Clamp(infp, minfp, maxfp)
		if res != c.out {
			t.Errorf("expected %v to be clamped between %v and %v, was %v", infp, minfp, maxfp, res)
		}
	}
}

func TestExpr(t *testing.T) {
	cases := []struct {
		input string
		res   fx.T
	}{
		{"  1 + 1 ", fx.I(2)},
		{"2 + 5 - 3.0", fx.F(2 + 5 - 3.0)},
		{" 10.3 * (2   + 1.25)", fx.F(10.3 * (2.0 + 1.25))},
		{"-1/2*4", fx.F(-1.0 / 2 * 4)},
	}

	for _, c := range cases {
		res := fx.Expr(c.input)
		if math.Abs(res.Float64()-c.res.Float64()) >= 0.0001 {
			t.Errorf("expected %v to be %v, was %v", c.input, c.res, res)
		}
	}
}

func TestExprVars(t *testing.T) {
	cases := []struct {
		input string
		vars  fx.ExprVarMap
		res   fx.T
	}{
		{"Pi / 2 ", fx.ExprVarMap{"Pi": fx.Pi()}, fx.Div(fx.Pi(), fx.F(2.0))},
		{"1 - t", fx.ExprVarMap{"t": fx.F(3.14)}, fx.Sub(fx.One(), fx.F(3.14))},
		{"-1 / n", fx.ExprVarMap{"n": fx.I(5)}, fx.Div(fx.NegOne(), fx.I(5))},
	}

	for _, c := range cases {
		res := fx.ExprVars(c.input, c.vars)
		if math.Abs(res.Float64()-c.res.Float64()) >= 0.0001 {
			t.Errorf("expected %v to be %v, was %v", c.input, c.res, res)
		}
	}
}
