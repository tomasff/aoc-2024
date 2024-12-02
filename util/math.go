package util

import "math"

func Abs(x int) int {
	if x == math.MinInt64 {
		return x
	}

	return max(x, -x)
}
