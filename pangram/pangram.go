/*
Package pangram implements IsPangram
*/
package pangram

import "unicode"

// IsPangram checks if a sentence is a pangram
func IsPangram(sentence string) bool {
	table := map[rune]bool{}

	for _, r := range sentence {

		if letter, ok := isLetter(r); ok {
			table[letter] = true
		}
	}

	return len(table) == 26
}

func isLetter(r rune) (rune, bool) {
	r = unicode.ToLower(r)
	return r, 'a' <= r && r <= 'z'
}
