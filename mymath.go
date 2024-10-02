package mymath

import "math"

func Abs(x float64) float64 {
	return math.Abs(float64(x))
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
