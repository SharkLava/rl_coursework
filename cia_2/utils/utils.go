package utils

import "math"

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Abs(a float64) float64 {
	return math.Abs(a)
}
