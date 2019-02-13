package lib

import (
	"math"
	"math/rand"
	"time"
	"wasmgo/types"
)

//var Math_abs = Math.abs;
//var Math_cos = Math.cos;
//var Math_sin = Math.sin;
//var Math_tan = Math.tan;
//var Math_acos = Math.acos;
//var Math_asin = Math.asin;
//var Math_atan = Math.atan;
//var Math_atan2 = Math.atan2;
//var Math_exp = Math.exp;
//var Math_log = Math.log;
//var Math_sqrt = Math.sqrt;
//var Math_ceil = Math.ceil;
//var Math_floor = Math.floor;
//var Math_pow = Math.pow;
//var Math_imul = Math.imul;
//var Math_fround = Math.fround;
//var Math_round = Math.round;
//var Math_min = Math.min;
//var Math_max = Math.max;
//var Math_clz32 = Math.clz32;
//var Math_trunc = Math.trunc;

type Math struct {
	types.VMInterface
}

func (*Math) Random() int64 {
	rand.Seed(time.Now().Unix())
	return int64(rand.Int())
}

//double log(double x)
func (*Math) Log(x float64) float64 {
	return math.Log(x)
}

//double pow(double x, double y)
func (*Math) Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

//double sqrt(double x)
func (*Math) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

//double log10(double x)
func (*Math) Log10(x float64) float64 {
	return math.Log10(x)
}

//double sin(double x)
func (*Math) Sin(x float64) float64 {
	return math.Sin(x)
}

//double sinh(double x)
func (*Math) Sinh(x float64) float64 {
	return math.Sinh(x)
}

//double exp(double x)
func (*Math) Exp(x float64) float64 {
	return math.Exp(x)
}

//double floor(double x)
func (*Math) Floor(x float64) float64 {
	return math.Floor(x)
}

//double cos(double x)
func (*Math) Cos(x float64) float64 {
	return math.Cos(x)
}

//double cosh(double x)
func (*Math) Cosh(x float64) float64 {
	return math.Cosh(x)
}

//double tan(double x)
func (*Math) Tan(x float64) float64 {
	return math.Tan(x)
}

//double tan(double x)
func (*Math) Atan(x float64) float64 {
	return math.Atan(x)
}

//double tan(double x)
func (*Math) Atan2(x float64, y float64) float64 {
	return math.Atan2(x, y)
}

//double tan(double x)
//func (*Math) Imul(x float64, y float64) float64 {
//	return math(x, y)
//}
