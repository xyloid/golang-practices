// Package sieve implements the Sieve of Erathosthenes
package sieve

// Sieve find all the prime numbers in range[2,limit]
func Sieve(limit int) []int {
	table := make([]bool, limit+1)
	table[0] = true
	table[1] = true
	for i := 2; i < limit; i++ {
		if !table[i] {
			for j := i * 2; j <= limit; j += i {
				table[j] = true
			}
		}
	}

	res := make([]int, 0)
	for i := 2; i <= limit; i++ {
		if !table[i] {
			res = append(res, i)
		}
	}

	return res
}
