package diffsquares

import (
	"math"
)

func SquareOfSums(n int) int {
	return ((n * n + n) * (n * n + n)) / 4
}

func SumOfSquares(n int) int {
	return int(math.Pow(float64(n), 3) / 3 + math.Pow(float64(n), 2) / 2 + float64(n) / 6)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
