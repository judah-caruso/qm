package fx

import "strconv"

const (
	MinFP  = T(-1 << 31)  // 32768.00
	MaxFP  = T(1<<31 - 1) // -32768.00
	One    = T(shift)
	NegOne = T(-shift)
)

// F converts a float32 into a fixed-point number.
func F(f float32) T {
	return T(f * shiftF32)
}

// I converts an integer into a fixed-point number.
func I(i int) T {
	return T(i) * shift
}

// T represents a 16.16 fixed-point number.
type T int32

func (f T) String() string {
	return strconv.FormatFloat(f.Float64(), 'f', -1, 32)
}

// Float converts a fixed-point number into a float32.
func (f T) Float() float32 {
	return float32(f) / shiftF32
}

// Float64 converts a fixed-point number into a float64.
func (f T) Float64() float64 {
	return float64(f.Float())
}

// Int converts a fixed-point number into an integer.
func (f T) Int() int {
	return int(f / shift)
}

// Add returns the result of adding two fixed-point numbers.
func Add(lhs, rhs T) T {
	return lhs + rhs
}

// Addi returns the result of adding a fixed-point number with an integer.
func Addi(lhs T, rhs int) T {
	return lhs + I(rhs)
}

// Addf returns the result of adding a fixed-point number with a float32.
func Addf(lhs T, rhs float32) T {
	return lhs + F(rhs)
}

// Sub returns the result of subtracting two fixed-point numbers.
func Sub(lhs, rhs T) T {
	return lhs - rhs
}

// Subi returns the result of subtracting a fixed-point number and an integer.
func Subi(lhs T, rhs int) T {
	return lhs - I(rhs)
}

// Subf returns the result of subtracting a fixed-point number and a float32.
func Subf(lhs T, rhs float32) T {
	return lhs - F(rhs)
}

// Mul returns the result of multiplying two fixed-point numbers.
func Mul(lhs, rhs T) T {
	return T((int64(lhs) * int64(rhs)) >> scale)
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
	return T((int64(lhs) << scale) / int64(rhs))
}

// Divi returns the result of dividing a fixed-point number and an integer.
func Divi(lhs T, rhs int) T {
	return Div(lhs, I(rhs))
}

// Divf returns the result of dividing a fixed-point number and a float32.
func Divf(lhs T, rhs float32) T {
	return Div(lhs, F(rhs))
}

// Abs returns the absolute value of a fixed-point number.
func Abs(f T) T {
	if f < 0 {
		return -f
	}
	return f
}

func Round(f T) T {
	return (f + shiftHalf) & -1
}

// Floor returns the larger integer value <= x.
func Floor(f T) T {
	rem := f & (shift - 1)
	if rem == 0 {
		return f
	}

	return (f - rem) & integerPart
}

// Ceil returns the smaller integer value >= x
func Ceil(f T) T {
	rem := f & (shift - 1)
	if rem == 0 {
		return f
	}

	return f + shift - rem
}

// Min returns the smaller of two fixed-point numbers.
func Min(a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

// Max returns the larger of two fixed-point numbers.
func Max(a, b T) T {
	if a > b {
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
	scale        = 16
	shift        = 1 << scale
	shiftHalf    = shift / 2
	shiftI64     = int64(shift)
	shiftF32     = float32(shift)
	fractionPart = T(0xFFFFFFFF >> int32(32-scale))
	integerPart  = -1 | fractionPart
)
