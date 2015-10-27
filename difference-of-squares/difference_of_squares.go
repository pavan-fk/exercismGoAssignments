package diffsquares

import (
	"math"
)

//SquareOfSums is the square of the sum of natural numbers till n.
func SquareOfSums(n int) int {
	return ((n*n + n) * (n*n + n)) / 4
}

//SumOfSquares is the sum of squares of natural numbers till n.
func SumOfSquares(n int) int {
	return int(math.Pow(float64(n), 3)/3 + math.Pow(float64(n), 2)/2 + float64(n)/6)
}

//Difference calculates the difference between SquareOfSums and SumOfSquares.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
