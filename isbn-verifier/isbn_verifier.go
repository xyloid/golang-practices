/*
Package isbn implements an isbn verifier
*/
package isbn

import "unicode"

// IsValidISBN check if an isbn number is valid
func IsValidISBN(isbn string) bool {
	digits := make([]int, 0)
	count := 0
	for _, r := range isbn {
		if unicode.IsDigit(r) || r == 'X' {
			count++

			if r == 'X' && count != 10 {
				return false
			}

			if r == 'X' {
				digits = append(digits, 10)
			} else {
				digits = append(digits, int(r-'0'))
			}

		} else if r != '-' {
			return false
		}
	}
	if len(digits) != 10 {
		return false
	}

	// calculate sum

	sum := 0
	for i, val := range digits {
		sum += (10 - i) * val
	}

	return sum%11 == 0
}
