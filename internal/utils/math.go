package utils

import "math"

// Approximately 判断两个float64是否相等
func Approximately(left, right float64) bool {
	// 默认精度EPSILON = 0.00001
	EPSILON := 0.00001
	return math.Abs(left-right) < EPSILON
}
