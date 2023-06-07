package math

import "math"

func SumFloat(x []float64) (sum float64) {
	for _, v := range x {
		sum += v
	}
	return
}

func Mean(x []float64) (mean float64) {
	return Sum(x) / float64(len(x))
}

func MeanVar(x []float64) (mean, variance float64) {
	mean = Mean(x)
	for _, v := range x {
		variance += (v - mean) * (v - mean)
	}
	return mean, variance / float64(len(x)-1)
}

// MeanStdDev
func MeanStdDev(x []float64) (mean, stdDev float64) {
	mean, stdDev = MeanVar(x)
	return mean, math.Sqrt(stdDev)
}

// DivisionInt return float64(x)/float64(y)
func DivisionInt(x, y int) float64 {
	return float64(x) / float64(y)
}

func Sum[G int | float64](x []G) (sum G) {
	for _, v := range x {
		sum += v
	}
	return
}

// SumInt sum int
func SumInt(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}
