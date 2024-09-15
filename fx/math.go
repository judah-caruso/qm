package fx

// Pi returns the approximate value of pi.
func Pi() T {
	return T{shift * 3141 / 1000}
}

// Pi2 returns the approximate value of pi/2.
func Pi2() T {
	return T{shift * 1570 / 1000}
}

// Tau returns the approximate value of 2*pi.
func Tau() T {
	return T{shift * 6283 / 1000}
}

// Sin calculates the approximate sine of an angle specified in radians.
func Sin(a T) T {
	return lookupRadian(sinLookup, a)
}

// Asin calculates the approximate arc-sine an angle specified in radians.
func Asin(a T) T {
	return lookupInterval(asinLookup, a)
}

// Cos calculates the approximate cosine of an angle specified in radians.
func Cos(a T) T {
	return lookupRadian(cosLookup, a)
}

// Acos calculates the approximate arc-cosine an angle specified in radians.
func Acos(a T) T {
	return lookupInterval(acosLookup, a)
}

// Sqrt calculates the approximate square root of x.
func Sqrt(x T) T {
	if x.int32 <= 0 {
		return Zero()
	}

	var (
		iterations = 10
		two        = T{shift * 2}
	)

	guess := Div(x, two)
	for range iterations {
		next := Divi(Add(guess, Div(x, guess)), 2)
		if Abs(Sub(next, guess)).int32 <= 1 {
			return next
		}

		guess = next
	}

	return guess
}

// InvSqrt calculates the approximate inverse square root of x.
func InvSqrt(x T) T {
	return Div(One(), Sqrt(x))
}

// Tan calculates the approximate tangent of an angle specified in radians.
func Tan(a T) T {
	return Div(Sin(a), Cos(a))
}

func lookupInterval(data [256]T, i T) T {
	i = Clamp(i, NegOne(), One())
	fi := Muli(Add(i, One()), 256/2)
	idx := Min(Floor(fi), I(255)).Int()
	return data[idx]
}

func lookupRadian(data [256]T, rad T) T {
	return lookupNormalized(data, Div(Mod(rad, Tau()), Tau()))
}

func lookupNormalized(data [256]T, frac T) T {
	fi := Muli(frac, 256)
	idx := Min(Floor(fi), I(255)).Int()
	return data[idx]
}
