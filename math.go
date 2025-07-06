// GENERATED! DO NOT EDIT
package gmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Abs returns the absolute value of x.
//
// Special cases are:
//
// Abs(±Inf) = +Inf
// Abs(NaN) = NaN
func Abs[N Number](x N) N {
	return N(math.Abs(float64(x)))
}

// Acos returns the arccosine, in radians, of x.
//
// Special case is:
//
// Acos(x) = NaN if x < -1 or x > 1
func Acos[N Number](x N) N {
	return N(math.Acos(float64(x)))
}

// Acosh returns the inverse hyperbolic cosine of x.
//
// Special cases are:
//
// Acosh(+Inf) = +Inf
// Acosh(x) = NaN if x < 1
// Acosh(NaN) = NaN
func Acosh[N Number](x N) N {
	return N(math.Acosh(float64(x)))
}

// Asin returns the arcsine, in radians, of x.
//
// Special cases are:
//
// Asin(±0) = ±0
// Asin(x) = NaN if x < -1 or x > 1
func Asin[N Number](x N) N {
	return N(math.Asin(float64(x)))
}

// Asinh returns the inverse hyperbolic sine of x.
//
// Special cases are:
//
// Asinh(±0) = ±0
// Asinh(±Inf) = ±Inf
// Asinh(NaN) = NaN
func Asinh[N Number](x N) N {
	return N(math.Asinh(float64(x)))
}

// Atan returns the arctangent, in radians, of x.
//
// Special cases are:
//
// Atan(±0) = ±0
// Atan(±Inf) = ±Pi/2
func Atan[N Number](x N) N {
	return N(math.Atan(float64(x)))
}

// Atan2 returns the arc tangent of y/x, using the signs of the two to
// determine the quadrant of the return value.
//
// Special cases are (in order):
//
// Atan2(y, NaN) = NaN
// Atan2(NaN, x) = NaN
// Atan2(+0, x>=0) = +0
// Atan2(-0, x>=0) = -0
// Atan2(+0, x<=-0) = +Pi
// Atan2(-0, x<=-0) = -Pi
// Atan2(y>0, 0) = +Pi/2
// Atan2(y<0, 0) = -Pi/2
// Atan2(+Inf, +Inf) = +Pi/4
// Atan2(-Inf, +Inf) = -Pi/4
// Atan2(+Inf, -Inf) = 3Pi/4
// Atan2(-Inf, -Inf) = -3Pi/4
// Atan2(y, +Inf) = 0
// Atan2(y>0, -Inf) = +Pi
// Atan2(y<0, -Inf) = -Pi
// Atan2(+Inf, x) = +Pi/2
// Atan2(-Inf, x) = -Pi/2
func Atan2[N Number](y N, x N) N {
	return N(math.Atan2(float64(y), float64(x)))
}

// Atanh returns the inverse hyperbolic tangent of x.
//
// Special cases are:
//
// Atanh(1) = +Inf
// Atanh(±0) = ±0
// Atanh(-1) = -Inf
// Atanh(x) = NaN if x < -1 or x > 1
// Atanh(NaN) = NaN
func Atanh[N Number](x N) N {
	return N(math.Atanh(float64(x)))
}

// Cbrt returns the cube root of x.
//
// Special cases are:
//
// Cbrt(±0) = ±0
// Cbrt(±Inf) = ±Inf
// Cbrt(NaN) = NaN
func Cbrt[N Number](x N) N {
	return N(math.Cbrt(float64(x)))
}

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
//
// Ceil(±0) = ±0
// Ceil(±Inf) = ±Inf
// Ceil(NaN) = NaN
func Ceil[N Number](x N) N {
	return N(math.Ceil(float64(x)))
}

// Copysign returns a value with the magnitude of f and the sign of sign.
func Copysign[N Number](f N, sign N) N {
	return N(math.Copysign(float64(f), float64(sign)))
}

// Cos returns the cosine of the radian argument x.
//
// Special cases are:
//
// Cos(±Inf) = NaN
// Cos(NaN) = NaN
func Cos[N Number](x N) N {
	return N(math.Cos(float64(x)))
}

// Cosh returns the hyperbolic cosine of x.
//
// Special cases are:
//
// Cosh(±0) = 1
// Cosh(±Inf) = +Inf
// Cosh(NaN) = NaN
func Cosh[N Number](x N) N {
	return N(math.Cosh(float64(x)))
}

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//
// Dim(+Inf, +Inf) = NaN
// Dim(-Inf, -Inf) = NaN
// Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim[N Number](x N, y N) N {
	return N(math.Dim(float64(x), float64(y)))
}

// Erf returns the error function of x.
//
// Special cases are:
//
// Erf(+Inf) = 1
// Erf(-Inf) = -1
// Erf(NaN) = NaN
func Erf[N Number](x N) N {
	return N(math.Erf(float64(x)))
}

// Erfc returns the complementary error function of x.
//
// Special cases are:
//
// Erfc(+Inf) = 0
// Erfc(-Inf) = 2
// Erfc(NaN) = NaN
func Erfc[N Number](x N) N {
	return N(math.Erfc(float64(x)))
}

// Erfcinv returns the inverse of Erfc(x).
//
// Special cases are:
//
// Erfcinv(0) = +Inf
// Erfcinv(2) = -Inf
// Erfcinv(x) = NaN if x < 0 or x > 2
// Erfcinv(NaN) = NaN
func Erfcinv[N Number](x N) N {
	return N(math.Erfcinv(float64(x)))
}

// Erfinv returns the inverse error function of x.
//
// Special cases are:
//
// Erfinv(1) = +Inf
// Erfinv(-1) = -Inf
// Erfinv(x) = NaN if x < -1 or x > 1
// Erfinv(NaN) = NaN
func Erfinv[N Number](x N) N {
	return N(math.Erfinv(float64(x)))
}

// Exp returns e**x, the base-e exponential of x.
//
// Special cases are:
//
// Exp(+Inf) = +Inf
// Exp(NaN) = NaN
//
// Very large values overflow to 0 or +Inf. Very small values underflow to 1.
func Exp[N Number](x N) N {
	return N(math.Exp(float64(x)))
}

// Exp2 returns 2**x, the base-2 exponential of x.
//
// Special cases are the same as Exp.
func Exp2[N Number](x N) N {
	return N(math.Exp2(float64(x)))
}

// Expm1 returns e**x - 1, the base-e exponential of x minus 1. It is more
// accurate than Exp(x) - 1 when x is near zero.
//
// Special cases are:
//
// Expm1(+Inf) = +Inf
// Expm1(-Inf) = -1
// Expm1(NaN) = NaN
//
// Very large values overflow to -1 or +Inf.
func Expm1[N Number](x N) N {
	return N(math.Expm1(float64(x)))
}

// FMA returns x * y + z, computed with only one rounding. (That is, FMA
// returns the fused multiply-add of x, y, and z.)
func FMA[N Number](x N, y N, z N) N {
	return N(math.FMA(float64(x), float64(y), float64(z)))
}

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//
// Floor(±0) = ±0
// Floor(±Inf) = ±Inf
// Floor(NaN) = NaN
func Floor[N Number](x N) N {
	return N(math.Floor(float64(x)))
}

// Frexp breaks f into a normalized fraction and an integral power of two.
// It returns frac and exp satisfying f == frac × 2**exp, with the absolute
// value of frac in the interval [½, 1).
//
// Special cases are:
//
// Frexp(±0) = ±0, 0
// Frexp(±Inf) = ±Inf, 0
// Frexp(NaN) = NaN, 0
func Frexp[N Number, I constraints.Integer](f N) (frac N, exp I) {
	fracR, expR := math.Frexp(float64(f))
	return N(fracR), I(expR)
}

// Gamma returns the Gamma function of x.
//
// Special cases are:
//
// Gamma(+Inf) = +Inf
// Gamma(+0) = +Inf
// Gamma(-0) = -Inf
// Gamma(x) = NaN for integer x < 0
// Gamma(-Inf) = NaN
// Gamma(NaN) = NaN
func Gamma[N Number](x N) N {
	return N(math.Gamma(float64(x)))
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid unnecessary overflow and
// underflow.
//
// Special cases are:
//
// Hypot(±Inf, q) = +Inf
// Hypot(p, ±Inf) = +Inf
// Hypot(NaN, q) = NaN
// Hypot(p, NaN) = NaN
func Hypot[N Number](p N, q N) N {
	return N(math.Hypot(float64(p), float64(q)))
}

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
//
// Ilogb(±Inf) = MaxInt32
// Ilogb(0) = MinInt32
// Ilogb(NaN) = MaxInt32
func Ilogb[N Number, I constraints.Integer](x N) I {
	return I(math.Ilogb(float64(x)))
}

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf[N Number, I constraints.Integer](sign I) N {
	return N(math.Inf(int(sign)))
}

// IsInf reports whether f is an infinity, according to sign. If sign > 0,
// IsInf reports whether f is positive infinity. If sign < 0, IsInf reports
// whether f is negative infinity. If sign == 0, IsInf reports whether f is
// either infinity.
func IsInf[N Number, I constraints.Integer](f N, sign I) bool {
	return math.IsInf(float64(f), int(sign))
}

// IsNaN reports whether f is an IEEE 754 “not-a-number” value.
func IsNaN[N Number](f N) (is bool) {
	return math.IsNaN(float64(f))
}

// J0 returns the order-zero Bessel function of the first kind.
//
// Special cases are:
//
// J0(±Inf) = 0
// J0(0) = 1
// J0(NaN) = NaN
func J0[N Number](x N) N {
	return N(math.J0(float64(x)))
}

// J1 returns the order-one Bessel function of the first kind.
//
// Special cases are:
//
// J1(±Inf) = 0
// J1(NaN) = NaN
func J1[N Number](x N) N {
	return N(math.J1(float64(x)))
}

// Jn returns the order-n Bessel function of the first kind.
//
// Special cases are:
//
// Jn(n, ±Inf) = 0
// Jn(n, NaN) = NaN
func Jn[N Number, I constraints.Integer](n I, x N) N {
	return N(math.Jn(int(n), float64(x)))
}

// Ldexp is the inverse of Frexp. It returns frac × 2**exp.
//
// Special cases are:
//
// Ldexp(±0, exp) = ±0
// Ldexp(±Inf, exp) = ±Inf
// Ldexp(NaN, exp) = NaN
func Ldexp[N Number, I constraints.Integer](frac N, exp I) N {
	return N(math.Ldexp(float64(frac), int(exp)))
}

// Lgamma returns the natural logarithm and sign (-1 or +1) of Gamma(x).
//
// Special cases are:
//
// Lgamma(+Inf) = +Inf
// Lgamma(0) = +Inf
// Lgamma(-integer) = +Inf
// Lgamma(-Inf) = -Inf
// Lgamma(NaN) = NaN
func Lgamma[N Number, I constraints.Integer](x N) (lgamma N, sign I) {
	lgammaR, signR := math.Lgamma(float64(x))
	return N(lgammaR), I(signR)
}

// Log returns the natural logarithm of x.
//
// Special cases are:
//
// Log(+Inf) = +Inf
// Log(0) = -Inf
// Log(x < 0) = NaN
// Log(NaN) = NaN
func Log[N Number](x N) N {
	return N(math.Log(float64(x)))
}

// Log10 returns the decimal logarithm of x. The special cases are the same as
// for Log.
func Log10[N Number](x N) N {
	return N(math.Log10(float64(x)))
}

// Log1p returns the natural logarithm of 1 plus its argument x. It is more
// accurate than Log(1 + x) when x is near zero.
//
// Special cases are:
//
// Log1p(+Inf) = +Inf
// Log1p(±0) = ±0
// Log1p(-1) = -Inf
// Log1p(x < -1) = NaN
// Log1p(NaN) = NaN
func Log1p[N Number](x N) N {
	return N(math.Log1p(float64(x)))
}

// Log2 returns the binary logarithm of x. The special cases are the same as
// for Log.
func Log2[N Number](x N) N {
	return N(math.Log2(float64(x)))
}

// Logb returns the binary exponent of x.
//
// Special cases are:
//
// Logb(±Inf) = +Inf
// Logb(0) = -Inf
// Logb(NaN) = NaN
func Logb[N Number](x N) N {
	return N(math.Logb(float64(x)))
}

// Max returns the larger of x or y.
//
// Special cases are:
//
// Max(x, +Inf) = Max(+Inf, x) = +Inf
// Max(x, NaN) = Max(NaN, x) = NaN
// Max(+0, ±0) = Max(±0, +0) = +0
// Max(-0, -0) = -0
//
// Note that this differs from the built-in function max when called with NaN
// and +Inf.
func Max[N Number](x N, y N) N {
	return N(math.Max(float64(x), float64(y)))
}

// Min returns the smaller of x or y.
//
// Special cases are:
//
// Min(x, -Inf) = Min(-Inf, x) = -Inf
// Min(x, NaN) = Min(NaN, x) = NaN
// Min(-0, ±0) = Min(±0, -0) = -0
//
// Note that this differs from the built-in function min when called with NaN
// and -Inf.
func Min[N Number](x N, y N) N {
	return N(math.Min(float64(x), float64(y)))
}

// Mod returns the floating-point remainder of x/y. The magnitude of the result
// is less than y and its sign agrees with that of x.
//
// Special cases are:
//
// Mod(±Inf, y) = NaN
// Mod(NaN, y) = NaN
// Mod(x, 0) = NaN
// Mod(x, ±Inf) = x
// Mod(x, NaN) = NaN
func Mod[N Number](x N, y N) N {
	return N(math.Mod(float64(x), float64(y)))
}

// Modf returns integer and fractional floating-point numbers that sum to f.
// Both values have the same sign as f.
//
// Special cases are:
//
// Modf(±Inf) = ±Inf, NaN
// Modf(NaN) = NaN, NaN
func Modf[N Number](f N) (int N, frac N) {
	intR, fracR := math.Modf(float64(f))
	return N(intR), N(fracR)
}

// Pow returns x**y, the base-x exponential of y.
//
// Special cases are (in order):
//
// Pow(x, ±0) = 1 for any x
// Pow(1, y) = 1 for any y
// Pow(x, 1) = x for any x
// Pow(NaN, y) = NaN
// Pow(x, NaN) = NaN
// Pow(±0, y) = ±Inf for y an odd integer < 0
// Pow(±0, -Inf) = +Inf
// Pow(±0, +Inf) = +0
// Pow(±0, y) = +Inf for finite y < 0 and not an odd integer
// Pow(±0, y) = ±0 for y an odd integer > 0
// Pow(±0, y) = +0 for finite y > 0 and not an odd integer
// Pow(-1, ±Inf) = 1
// Pow(x, +Inf) = +Inf for |x| > 1
// Pow(x, -Inf) = +0 for |x| > 1
// Pow(x, +Inf) = +0 for |x| < 1
// Pow(x, -Inf) = +Inf for |x| < 1
// Pow(+Inf, y) = +Inf for y > 0
// Pow(+Inf, y) = +0 for y < 0
// Pow(-Inf, y) = Pow(-0, -y)
// Pow(x, y) = NaN for finite x < 0 and finite non-integer y
func Pow[N Number](x N, y N) N {
	return N(math.Pow(float64(x), float64(y)))
}

// Pow10 returns 10**n, the base-10 exponential of n.
//
// Special cases are:
//
// Pow10(n) =    0 for n < -323
// Pow10(n) = +Inf for n > 308
func Pow10[N Number, I constraints.Integer](n I) N {
	return N(math.Pow10(int(n)))
}

// Remainder returns the IEEE 754 floating-point remainder of x/y.
//
// Special cases are:
//
// Remainder(±Inf, y) = NaN
// Remainder(NaN, y) = NaN
// Remainder(x, 0) = NaN
// Remainder(x, ±Inf) = x
// Remainder(x, NaN) = NaN
func Remainder[N Number](x N, y N) N {
	return N(math.Remainder(float64(x), float64(y)))
}

// Round returns the nearest integer, rounding half away from zero.
//
// Special cases are:
//
// Round(±0) = ±0
// Round(±Inf) = ±Inf
// Round(NaN) = NaN
func Round[N Number](x N) N {
	return N(math.Round(float64(x)))
}

// RoundToEven returns the nearest integer, rounding ties to even.
//
// Special cases are:
//
// RoundToEven(±0) = ±0
// RoundToEven(±Inf) = ±Inf
// RoundToEven(NaN) = NaN
func RoundToEven[N Number](x N) N {
	return N(math.RoundToEven(float64(x)))
}

// Signbit reports whether x is negative or negative zero.
func Signbit[N Number](x N) bool {
	return math.Signbit(float64(x))
}

// Sin returns the sine of the radian argument x.
//
// Special cases are:
//
// Sin(±0) = ±0
// Sin(±Inf) = NaN
// Sin(NaN) = NaN
func Sin[N Number](x N) N {
	return N(math.Sin(float64(x)))
}

// Sincos returns Sin(x), Cos(x).
//
// Special cases are:
//
// Sincos(±0) = ±0, 1
// Sincos(±Inf) = NaN, NaN
// Sincos(NaN) = NaN, NaN
func Sincos[N Number](x N) (sin N, cos N) {
	sinR, cosR := math.Sincos(float64(x))
	return N(sinR), N(cosR)
}

// Sinh returns the hyperbolic sine of x.
//
// Special cases are:
//
// Sinh(±0) = ±0
// Sinh(±Inf) = ±Inf
// Sinh(NaN) = NaN
func Sinh[N Number](x N) N {
	return N(math.Sinh(float64(x)))
}

// Sqrt returns the square root of x.
//
// Special cases are:
//
// Sqrt(+Inf) = +Inf
// Sqrt(±0) = ±0
// Sqrt(x < 0) = NaN
// Sqrt(NaN) = NaN
func Sqrt[N Number](x N) N {
	return N(math.Sqrt(float64(x)))
}

// Tan returns the tangent of the radian argument x.
//
// Special cases are:
//
// Tan(±0) = ±0
// Tan(±Inf) = NaN
// Tan(NaN) = NaN
func Tan[N Number](x N) N {
	return N(math.Tan(float64(x)))
}

// Tanh returns the hyperbolic tangent of x.
//
// Special cases are:
//
// Tanh(±0) = ±0
// Tanh(±Inf) = ±1
// Tanh(NaN) = NaN
func Tanh[N Number](x N) N {
	return N(math.Tanh(float64(x)))
}

// Trunc returns the integer value of x.
//
// Special cases are:
//
// Trunc(±0) = ±0
// Trunc(±Inf) = ±Inf
// Trunc(NaN) = NaN
func Trunc[N Number](x N) N {
	return N(math.Trunc(float64(x)))
}

// Y0 returns the order-zero Bessel function of the second kind.
//
// Special cases are:
//
// Y0(+Inf) = 0
// Y0(0) = -Inf
// Y0(x < 0) = NaN
// Y0(NaN) = NaN
func Y0[N Number](x N) N {
	return N(math.Y0(float64(x)))
}

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
//
// Y1(+Inf) = 0
// Y1(0) = -Inf
// Y1(x < 0) = NaN
// Y1(NaN) = NaN
func Y1[N Number](x N) N {
	return N(math.Y1(float64(x)))
}

// Yn returns the order-n Bessel function of the second kind.
//
// Special cases are:
//
// Yn(n, +Inf) = 0
// Yn(n ≥ 0, 0) = -Inf
// Yn(n < 0, 0) = +Inf if n is odd, -Inf if n is even
// Yn(n, x < 0) = NaN
// Yn(n, NaN) = NaN
func Yn[N Number, I constraints.Integer](n I, x N) N {
	return N(math.Yn(int(n), float64(x)))
}
