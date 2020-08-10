// Package grains implements functions about counting the grain numbers on a 64 squares chessboard.
package grains

import "fmt"

// Square returns the number of grains in n-th sqaure
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("n must be in the range [1,64]")
	}

	return 1 << (n - 1), nil
}

// Total returns the total number of grains on the chessboard
func Total() uint64 {
	return 1<<64 - 1
}
