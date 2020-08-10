// Package prime imlements function factor
package prime

// Factors return all the prime factors of a give number
func Factors(n int64) []int64 {
	ret := make([]int64, 0)
	for factor := int64(2); factor <= n; {

		if isPrime(factor) && n%factor == 0 {
			ret = append(ret, factor)
			n = n / factor
		} else {
			factor++
		}
	}
	return ret
}

func isPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 || n < 2 {
		return false
	}

	for i := int64(3); i < n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
