package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	}
	if isUpper(remark) && isQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}
	if isUpper(remark) {
		return "Whoa, chill out!"
	}
	if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}

func isUpper(remark string) bool {
	hasLetter := false
	for _, c := range remark {
		if unicode.IsLetter(c) && !unicode.IsUpper(c) {
			return false
		}
		if unicode.IsLetter(c) {
			hasLetter = true
		}
	}
	return hasLetter
}

func isQuestion(remark string) bool {
	remark = strings.TrimSpace(remark)
	return remark[len(remark)-1] == '?'
}
