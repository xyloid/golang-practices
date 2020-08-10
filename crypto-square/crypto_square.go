/*
Package cryptosquare implements Crypto Square
*/
package cryptosquare

import (
	"math"
	"unicode"
)

// Encode implements Crypto Square
func Encode(sentence string) string {
	letters := make([]rune, 0)
	for _, r := range sentence {
		r = unicode.ToLower(r)
		if isLetter(r) {
			letters = append(letters, r)
		}
	}
	// create squares
	cols := int(math.Ceil(math.Sqrt(float64(len(letters)))))
	rows := cols - 1

	if cols*rows < len(letters) {
		rows++
	}

	res := ""

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			index := row*cols + col
			if index < len(letters) {
				res += string(letters[index])
			} else {
				res += " "
			}
		}
		if col < cols-1 {
			res += " "
		}

	}
	return res
}

func isLetter(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	}
	if r < 'a' || r > 'z' {
		return false
	}
	return true
}
