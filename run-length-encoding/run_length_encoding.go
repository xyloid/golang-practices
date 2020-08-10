/*
Package encode implements Run Length Encoding
*/
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode implements Run Length Encoding
func RunLengthEncode(sentence string) string {
	if sentence == "" {
		return ""
	}
	var sb strings.Builder

	runes := []rune(sentence)
	count := make([]int, len(runes))

	for i, r := range runes {
		if i == 0 {
			count[0] = 1
		} else if r == runes[i-1] {
			count[i] = count[i-1] + 1
		} else {
			count[i] = 1
		}
	}

	for i, r := range runes {
		if i < len(runes)-1 && r != runes[i+1] || i == len(runes)-1 {
			if count[i] > 1 {
				sb.WriteString(strconv.Itoa(count[i]))
			}
			sb.WriteRune(runes[i])
		}
	}

	return sb.String()
}

// RunLengthDecode decodes
func RunLengthDecode(sentence string) string {
	runes := []rune(sentence)
	var sb strings.Builder
	count := 0
	for i := 0; i < len(runes); {
		if unicode.IsDigit(runes[i]) {
			count = count*10 + int(runes[i]-'0')
		} else {
			if count > 0 {
				for count > 0 {
					sb.WriteRune(runes[i])
					count--
				}
			} else {
				sb.WriteRune(runes[i])
			}
		}
		i++
	}

	return sb.String()
}
