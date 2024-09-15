package fx

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// MaxT returns the highest value representable by T (approximately 32768.00)
func MaxT() T {
	return T{1<<31 - 1}
}

// MinT returns the lowest value representable by T (approximately -32768.00)
func MinT() T {
	return T{-1 << 31}
}

// Zero returns the fixed-point value of 0.
func Zero() T {
	return T{0}
}

// One returns the fixed-point value of 1.
func One() T {
	return T{shift}
}

// NegOne returns the fixed-point value of -1.
func NegOne() T {
	return T{-shift}
}

// F converts a float into a fixed-point number.
func F[FT ~float32 | ~float64](f FT) T {
	return T{int32(f * FT(shiftF32))}
}

// I converts an integer into a fixed-point number.
func I(i int) T {
	return T{int32(i) * shift}
}

// Expr parses a math expression and returns its fixed-point value.
// The expression follows the same operator precedence rules as Go.
// Note: only constant numeric values are valid in the given expression.
// To use named values, call ExprVars instead.
func Expr(expr string) T {
	e, err := parser.ParseExpr(expr)
	if err != nil {
		panic(fmt.Errorf("invalid fixed-point expression: %w", err))
	}

	return eval(e, nil)
}

func eval(e ast.Expr, lookup ExprVarMap) T {
	switch v := e.(type) {
	case *ast.BinaryExpr:
		lhs := eval(v.X, lookup)
		rhs := eval(v.Y, lookup)

		switch v.Op {
		case token.ADD:
			return Add(lhs, rhs)
		case token.SUB:
			return Sub(lhs, rhs)
		case token.MUL:
			return Mul(lhs, rhs)
		case token.QUO:
			return Div(lhs, rhs)
		case token.REM:
			return Mod(lhs, rhs)
		default:
			panic(fmt.Errorf("unexpected operator in fixed-point expression: %s", v.Op))
		}

	case *ast.BasicLit:
		if v.Kind != token.INT && v.Kind != token.FLOAT {
			panic(fmt.Errorf("unexpected value in fixed-point expression: %s", v.Value))
		}

		f, err := strconv.ParseFloat(v.Value, 32)
		if err != nil {
			panic(fmt.Errorf("invalid number in fixed-point expression: %s", v.Value))
		}

		return F(float32(f))

	case *ast.Ident:
		if lookup != nil {
			f, ok := lookup[v.Name]
			if !ok {
				goto error
			}

			return f
		}

	error:
		panic(fmt.Errorf("unexpected identifier in fixed-point expression: %s", v.Name))

	case *ast.ParenExpr:
		return eval(v.X, lookup)

	default:
		panic(fmt.Errorf("unexpected node in parsed fixed-point expression: %T", v))
	}

	return Zero()
}

type ExprVarMap map[string]T

// ExprVars parses a math expression and returns its fixed-point value.
func ExprVars(expr string, values ExprVarMap) T {
	return T{0}
}

// T represents a 16.16 fixed-point number.
type T struct{ int32 }

func (f T) String() string {
	return strconv.FormatFloat(f.Float64(), 'f', -1, 32)
}

// Float converts a fixed-point number into a float32.
func (f T) Float() float32 {
	return float32(f.int32) / shiftF32
}

// Float64 converts a fixed-point number into a float64.
func (f T) Float64() float64 {
	return float64(f.Float())
}

// Int converts a fixed-point number into an integer.
func (f T) Int() int {
	return int(f.int32 / shift)
}

// Raw returns the internal representation of f.
func (f T) Raw() int32 {
	return f.int32
}

func Negate(f T) T {
	return T{-f.int32}
}

// Add returns the result of adding two fixed-point numbers.
func Add(lhs, rhs T) T {
	return T{lhs.int32 + rhs.int32}
}

// Addi returns the result of adding a fixed-point number with an integer.
func Addi(lhs T, rhs int) T {
	return Add(lhs, I(rhs))
}

// Addf returns the result of adding a fixed-point number with a float32.
func Addf(lhs T, rhs float32) T {
	return Add(lhs, F(rhs))
}

// Sub returns the result of subtracting two fixed-point numbers.
func Sub(lhs, rhs T) T {
	return T{lhs.int32 - rhs.int32}
}

// Subi returns the result of subtracting a fixed-point number and an integer.
func Subi(lhs T, rhs int) T {
	return Sub(lhs, I(rhs))
}

// Subf returns the result of subtracting a fixed-point number and a float32.
func Subf(lhs T, rhs float32) T {
	return Sub(lhs, F(rhs))
}

// Mul returns the result of multiplying two fixed-point numbers.
func Mul(lhs, rhs T) T {
	return T{int32((int64(lhs.int32) * int64(rhs.int32)) >> scale)}
}

// Muli returns the result of multiplying a fixed-point number and an integer.
func Muli(lhs T, rhs int) T {
	return Mul(lhs, I(rhs))
}

// Mulf returns the result of multiplying a fixed-point number and a float32.
func Mulf(lhs T, rhs float32) T {
	return Mul(lhs, F(rhs))
}

// Div returns the result of dividing two fixed-point numbers.
func Div(lhs, rhs T) T {
	return T{int32((int64(lhs.int32) << scale) / int64(rhs.int32))}
}

// Divi returns the result of dividing a fixed-point number and an integer.
func Divi(lhs T, rhs int) T {
	return Div(lhs, I(rhs))
}

// Divf returns the result of dividing a fixed-point number and a float32.
func Divf(lhs T, rhs float32) T {
	return Div(lhs, F(rhs))
}

func Square(f T) T {
	return Mul(f, f)
}

func Mod(lhs, rhs T) T {
	return T{lhs.int32 % rhs.int32}
}

// Abs returns the absolute value of a fixed-point number.
func Abs(f T) T {
	if f.int32 < 0 {
		return T{-f.int32}
	}
	return f
}

func Round(f T) T {
	return T{(f.int32 + shiftHalf) & -1}
}

// Floor returns the larger integer value <= x.
func Floor(f T) T {
	rem := f.int32 & (shift - 1)
	if rem == 0 {
		return f
	}

	return T{(f.int32 - rem) & integerPart}
}

// Ceil returns the smaller integer value >= x
func Ceil(f T) T {
	rem := f.int32 & (shift - 1)
	if rem == 0 {
		return f
	}

	return T{f.int32 + shift - rem}
}

// Min returns the smaller of two fixed-point numbers.
func Min(a, b T) T {
	if a.int32 < b.int32 {
		return a
	} else {
		return b
	}
}

// Max returns the larger of two fixed-point numbers.
func Max(a, b T) T {
	if a.int32 > b.int32 {
		return a
	} else {
		return b
	}
}

// Clamp returns a fixed-point number within the range of min and max.
func Clamp(f, min, max T) T {
	return Max(min, Min(f, max))
}

const (
	scale        = int32(16)
	shift        = int32(1 << scale)
	shiftHalf    = shift / 2
	shiftF32     = float32(shift)
	fractionPart = int32(0xFFFFFFFF >> int32(32-scale))
	integerPart  = -1 | fractionPart
)
