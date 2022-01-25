package math

import "math"

func Sum(x []float64) (sum float64) {
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

func MeanStdDev(x []float64) (mean, stdDev float64) {
	mean, stdDev = MeanVar(x)
	return mean, math.Sqrt(stdDev)
}
