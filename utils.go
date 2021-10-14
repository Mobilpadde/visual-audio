package visual

import "math"

func minMax(vals []int16) (float64, float64) {
	min, max := int16(math.MaxInt16), int16(math.MinInt16)
	for _, x := range vals {
		if min > x {
			min = x
		}

		if max < x {
			max = x
		}
	}

	return float64(min), float64(max)
}

func mapNumber(x, inMin, inMax, outMin, outMax float64) float64 {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
