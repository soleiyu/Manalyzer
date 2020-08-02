package lib

import (
	"math"
)

func Dft_rn(darr []float64, k float64) float64 {
	cv := 2.0 * k * math.Pi / float64(len(darr))
	res := 0.0

	for i := 0; i < len(darr); i++ {
		res += darr[i] * math.Cos((float64(i) * cv))
	}

	return res / float64(len(darr))
}

