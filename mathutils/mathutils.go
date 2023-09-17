package mathutils

import (
	"math"

	"github.com/joselws/go-utils/mytypes"
)

func Max[T mytypes.Number](numbers ...T) T {
	var maximum T
	var hasNaN bool
	for index, number := range numbers {
		numberFloat := float64(number)
		switch {
		case index == 0:
			maximum = number
		case math.IsInf(numberFloat, 1):
			return T(math.Inf(1))
		case math.IsNaN(numberFloat):
			hasNaN = true
		case number > maximum:
			maximum = number
		}
	}
	if hasNaN {
		return T(math.NaN())
	}
	return maximum
}

func Min[T mytypes.Number](numbers ...T) T {
	var minimum T
	var hasNaN bool
	for index, number := range numbers {
		numberFloat := float64(number)
		switch {
		case index == 0:
			minimum = number
		case math.IsInf(numberFloat, 1):
			return T(math.Inf(-1))
		case math.IsNaN(numberFloat):
			hasNaN = true
		case number < minimum:
			minimum = number
		}
	}
	if hasNaN {
		return T(math.NaN())
	}
	return minimum
}
