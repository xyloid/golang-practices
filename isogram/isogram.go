/*
Package isogram implements function IsIsogram.
*/
package isogram

import (
	"unicode"
)

// IsIsogram checks if a word is an isogram
func IsIsogram(word string) bool {
	var record = map[rune]bool{}
	for _, letter := range word {
		if letter == ' ' || letter == '-' {
			continue
		}
		letter = unicode.ToLower(letter)
		if record[letter] {
			return false
		}
		record[letter] = true
	}
	return true
}
