package gmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

func Acos[T Number](x T) T {
	return T(math.Acos(float64(x)))
}

func Acosh[T Number](x T) T {
	return T(math.Acosh(float64(x)))
}

func Asin[T Number](x T) T {
	return T(math.Asin(float64(x)))
}

func Asinh[T Number](x T) T {
	return T(math.Asinh(float64(x)))
}

func Atan[T Number](x T) T {
	return T(math.Atan(float64(x)))
}

func Atan2[T Number](y, x T) T {
	return T(math.Atan2(float64(y), float64(x)))
}

func Atanh[T Number](x T) T {
	return T(math.Atanh(float64(x)))
}

func Cbrt[T Number](x T) T {
	return T(math.Cbrt(float64(x)))
}

func Ceil[T Number](x T) T {
	return T(math.Ceil(float64(x)))
}

func Copysign[T Number](x, y T) T {
	return T(math.Copysign(float64(x), float64(y)))
}

func Cos[T Number](x T) T {
	return T(math.Cos(float64(x)))
}

func Cosh[T Number](x T) T {
	return T(math.Cosh(float64(x)))
}

func Dim[T Number](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

func Erf[T Number](x T) T {
	return T(math.Erf(float64(x)))
}

func Erfc[T Number](x T) T {
	return T(math.Erfc(float64(x)))
}

func Erfcinv[T Number](x T) T {
	return T(math.Erfcinv(float64(x)))
}

func Erfinv[T Number](x T) T {
	return T(math.Erfinv(float64(x)))
}

func Exp[T Number](x T) T {
	return T(math.Exp(float64(x)))
}

func Exp2[T Number](x T) T {
	return T(math.Exp2(float64(x)))
}

func Expm1[T Number](x T) T {
	return T(math.Expm1(float64(x)))
}

func FMA[T Number](x, y, z T) T {
	return T(math.FMA(float64(x), float64(y), float64(z)))
}

func Floor[T Number](x T) T {
	return T(math.Floor(float64(x)))
}

func Frexp[T constraints.Float](f T) (frac T, exp int) {
	fra, exp := math.Frexp(float64(f))
	return T(fra), exp
}

func Gamma[T Number](x T) T {
	return T(math.Gamma(float64(x)))
}

func Hypot[T Number](p, q T) T {
	return T(math.Hypot(float64(p), float64(q)))
}

func Ilogb[T constraints.Float](x T) int {
	return math.Ilogb(float64(x))
}

func Inf[T constraints.Float](sign int) T {
	return T(math.Inf(sign))
}

func IsInf[T constraints.Float](f T, sign int) bool {
	return math.IsInf(float64(f), sign)
}

func IsNaN[T constraints.Float](f T) (is bool) {
	return math.IsNaN(float64(f))
}

func J0[T Number](x T) T {
	return T(math.J0(float64(x)))
}

func J1[T Number](x T) T {
	return T(math.J1(float64(x)))
}

func Jn[T Number](n int, x T) T {
	return T(math.Jn(n, float64(x)))
}

func Ldexp[T Number](frac T, exp int) T {
	return T(math.Ldexp(float64(frac), exp))
}

func Lgamma[T Number](x T) (lgamma T, sign int) {
	g, sign := math.Lgamma(float64(x))
	return T(g), sign
}

func Log[T Number](x T) T {
	return T(math.Log(float64(x)))
}

func Log10[T Number](x T) T {
	return T(math.Log10(float64(x)))
}

func Log1p[T Number](x T) T {
	return T(math.Log1p(float64(x)))
}

func Log2[T Number](x T) T {
	return T(math.Log2(float64(x)))
}

func Logb[T Number](x T) T {
	return T(math.Logb(float64(x)))
}

func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Mod[T Number](x, y T) T {
	return T(math.Mod(float64(x), float64(y)))
}

func Modf[I constraints.Integer, T constraints.Float](f T) (int I, frac T) {
	i, fr := math.Modf(float64(f))
	return I(i), T(fr)
}

func Pow[T Number](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

func Pow10[T Number](n int) T {
	return T(math.Pow10(n))
}

func Remainder[T Number](x, y T) T {
	return T(math.Remainder(float64(x), float64(y)))
}

func Round[T constraints.Float, I constraints.Integer](x T) I {
	return I(math.Round(float64(x)))
}

func RoundToEven[T constraints.Float, I constraints.Integer](x T) I {
	return I(math.RoundToEven(float64(x)))
}

func Signbit[T Number](x T) bool {
	return math.Signbit(float64(x))
}

func Sin[T Number](x T) T {
	return T(math.Sin(float64(x)))
}

func Sincos[T Number](x T) (sin, cos T) {
	s, c := math.Sincos(float64(x))
	return T(s), T(c)
}

func Sinh[T Number](x T) T {
	return T(math.Sinh(float64(x)))
}

func Sqrt[T Number](x T) T {
	return T(math.Sqrt(float64(x)))
}

func Tan[T Number](x T) T {
	return T(math.Tan(float64(x)))
}

func Tanh[T Number](x T) T {
	return T(math.Tanh(float64(x)))
}

func Trunc[T constraints.Float, I constraints.Integer](x T) I {
	return I(math.Trunc(float64(x)))
}

func Y0[T Number](x T) T {
	return T(math.Y0(float64(x)))
}

func Y1[T Number](x T) T {
	return T(math.Y1(float64(x)))
}

func Yn[T Number](n int, x T) T {
	return T(math.Yn(n, float64(x)))
}
