// Package luhn implements the Luhn Algorithm in the function Valid().
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a number sequence is valid.
func Valid(numbers string) bool {
	numbers = strings.ReplaceAll(numbers, " ", "")

	if len(numbers) < 2 {
		return false
	}

	sum := 0
	isEvenDigit := len(numbers)%2 == 0

	for _, letter := range numbers {
		if !unicode.IsDigit(letter) {
			return false
		}

		val := int(letter - '0')
		if isEvenDigit {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		sum += val
		isEvenDigit = !isEvenDigit
	}

	return sum%10 == 0
}
