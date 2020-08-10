// Package cipher implements simple cipher
package cipher

import (
	"strings"
	"unicode"
)

type shiftCipher struct {
	distance int
}

func (s shiftCipher) Encode(plain string) (res string) {
	plain = strings.ToLower(plain)

	for _, letter := range plain {
		if unicode.IsLetter(letter) {
			index := (int(letter-'a') + s.distance) % 26
			letter = rune('a' + index)
			res += string(letter)
		}

	}
	return
}

func (s shiftCipher) Decode(plain string) (res string) {
	plain = strings.ToLower(plain)

	for _, letter := range plain {
		if unicode.IsLetter(letter) {
			index := (int(letter-'a') - s.distance + 26) % 26
			letter = rune('a' + index)

			res += string(letter)
		}

	}
	return
}

// NewCaesar creates an object for Caeser cipher
func NewCaesar() Cipher {
	return shiftCipher{3}
}

// NewShift creates an object for shift cipher
func NewShift(distance int) Cipher {
	if distance < 0 {
		distance += 26
	}
	if distance > 25 || distance < 1 {
		return nil
	}
	return shiftCipher{distance}
}

type vigenere struct {
	key []rune
}

func (v vigenere) Encode(plain string) (res string) {
	plain = filter(plain)
	for i, letter := range plain {
		index := i % len(v.key)
		letter = (letter-'a'+v.key[index]-'a')%26 + 'a'
		res += string(letter)
	}
	return
}

func (v vigenere) Decode(plain string) (res string) {

	for i, letter := range plain {
		index := i % len(v.key)
		letter = (letter-v.key[index]+26)%26 + 'a'
		res += string(letter)
	}
	return
}

// NewVigenere creates an object for Vigenere cipher
func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	isAllA := true

	for _, letter := range key {
		if letter != 'a' {
			isAllA = false
		}
		if unicode.IsUpper(letter) {
			return nil
		}
		if !unicode.IsLetter(letter) {
			return nil
		}
	}
	if isAllA {
		return nil
	}

	return vigenere{[]rune(key)}
}

func filter(plain string) (res string) {
	plain = strings.ToLower(plain)
	for _, letter := range plain {
		if unicode.IsLetter(letter) {
			res += string(letter)
		}
	}
	return
}
