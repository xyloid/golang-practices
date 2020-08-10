/*
Package anagram implements Detect
*/
package anagram

import (
	"strings"
	"unicode"
)

// Detect checks if a candidate is an anagram
func Detect(word string, candidates []string) []string {
	ret := make([]string, 0)
	letterCount := count(word)
	word = strings.ToLower(word)

	for _, wd := range candidates {
		if check(wd, letterCount) && strings.ToLower(wd) != word {
			ret = append(ret, wd)
		}
	}
	return ret
}

func count(word string) map[rune]int {
	ret := make(map[rune]int)
	for _, letter := range word {
		letter = unicode.ToLower(letter)
		ret[letter] = ret[letter] + 1
	}
	return ret
}

func check(word string, count map[rune]int) bool {
	table := make(map[rune]int)

	for k, v := range count {
		table[k] = v
	}

	for _, letter := range word {
		letter = unicode.ToLower(letter)
		if table[letter] == 0 {
			return false
		}
		table[letter] = table[letter] - 1
		if table[letter] == 0 {
			delete(table, letter)
		}
	}
	return len(table) == 0
}
