// Package atbash implements atbash cipher
package atbash

import (
	"strings"
	"unicode"
)

// Atbash is an atbash cipher
func Atbash(sentence string) (res string) {
	sentence = strings.ToLower(sentence)
	sentence = strings.ReplaceAll(sentence, " ", "")

	count := 0
	for _, letter := range sentence {

		if !unicode.IsDigit(letter) && !unicode.IsLetter(letter) {
			continue
		}

		if count > 0 && count%segLength == 0 {
			res += " "
		}

		newLetter := translate(letter)
		res += newLetter
		count++

	}

	return res
}

const segLength = 5

func translate(character rune) string {
	if character >= 'a' && character <= 'z' {
		character = 'a' + 25 - (character - 'a')
	}
	return string(character)
}
