/*
Package prime is about prime numbers
*/
package prime

// Nth returns nth prime number
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	p := 2
	for i := 2; n > 0; i++ {
		if isPrime(i) {
			n--
			p = i
		}
	}
	return p, true
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num > 2 && num%2 == 0 {
		return false
	}

	for f := 3; f <= num/2; f += 2 {
		if num%f == 0 {
			return false
		}
	}
	return true
}
