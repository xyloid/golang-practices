/*
Package wordcount implements a word count function
*/
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is the data structure for WordCount
type Frequency map[string]int

// WordCount counts the words in a sentence.
func WordCount(sentence string) Frequency {
	ret := make(map[string]int)

	splited := split(sentence, " \t\n^$%")
	for _, sp := range splited {
		ret[sp]++
	}
	return ret
}

func split(sentence string, sep string) []string {
	seps := make(map[rune]bool)
	for _, letter := range sep {
		seps[letter] = true
	}

	ret := make([]string, 0)

	var sb strings.Builder
	for _, letter := range sentence {
		letter = unicode.ToLower(letter)
		if letter != '\'' && (unicode.IsPunct(letter) || seps[letter]) {
			if sb.Len() > 0 {
				word := sb.String()
				if word[len(word)-1] == '\'' {
					word = string([]rune(word)[0 : len(word)-1])
				}
				ret = append(ret, word)
				sb.Reset()
			}
		} else {
			if sb.Len() == 0 && letter == '\'' {
				continue
			}
			sb.WriteRune(letter)
		}
	}

	if sb.Len() > 0 {
		word := sb.String()
		if word[len(word)-1] == '\'' {
			word = string([]rune(word)[0 : len(word)-1])
		}
		ret = append(ret, word)
	}

	return ret
}
