/*
Package diffsquares implements functions related to the calculation of square of sum, sum of squares and their difference.
*/
package diffsquares

// SquareOfSum calculates the square of the sum from 1 to n.
func SquareOfSum(n int) int {
	base := n * (n + 1) / 2
	return base * base
}

// SumOfSquares calculates the sum of squares of 1 to n.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between SquareOfSum(n) and SumOfSquares(n).
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
