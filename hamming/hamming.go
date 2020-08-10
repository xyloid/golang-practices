/*
Package hamming computes the hamming distance between two strings.
*/
package hamming

import "fmt"

// Distance computes the hamming distance between two strings with equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("DNA strands should have same length")
	}
	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil

}
