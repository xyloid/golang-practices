// Package rotationalcipher implements a rotaional cipher
package rotationalcipher

import (
	"unicode"
)

// RotationalCipher performs rotational cipher aloorithm
func RotationalCipher(plain string, key int) (res string) {
	for _, letter := range plain {
		if unicode.IsLetter(letter) {
			if unicode.IsLower(letter) {
				index := (int(letter-'a') + key) % 26
				letter = rune('a' + index)
			} else {
				index := (int(letter-'A') + key) % 26
				letter = rune('A' + index)
			}
		}
		res += string(letter)
	}
	return
}
