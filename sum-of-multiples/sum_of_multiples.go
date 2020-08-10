/*
Package summultiples implements the Solution of Sum Of Multiples
*/
package summultiples

import (
	"sort"
)

// SumMultiples calculates the sum of mupltiples of the divisors that smaller than the limit.
func SumMultiples(limit int, divisors ...int) int {

	ret := 0
	if len(divisors) == 0 {
		return ret
	}
	sort.Ints(divisors)
	candidate := make([]int, len(divisors))
	copy(candidate, divisors)

	start := 0
	for start < len(divisors) && divisors[start] < 1 {
		start++
	}

	length := len(candidate)

	for start < length && candidate[start] < limit {
		for i := start; i < length; i++ {
			if candidate[i] < limit {
				curNum := candidate[i]
				for j := i + 1; j < len(divisors); j++ {
					if curNum%divisors[j] == 0 {
						curNum = 0
						break
					}
				}
				ret += curNum
			} else {
				length = i
			}
			candidate[i] += divisors[i]
		}
	}
	return ret
}
