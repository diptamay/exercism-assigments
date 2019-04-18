// Package diffsquares finds the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSum finds the square of the sum of first N natural numbers
func SquareOfSum(N int) int {
	sum := N * (N + 1) / 2
	return sum * sum
}

// SumOfSquares finds the sum of squares of the first N natural numbers
func SumOfSquares(N int) int {
	return N * (N + 1) * (2*N + 1) / 6
}

// Difference finds the difference between the square of the sum and the sum of the squares of the first N natural numbers
func Difference(N int) int {
	return SquareOfSum(N) - SumOfSquares(N)
}
