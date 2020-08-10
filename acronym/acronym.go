package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	re := regexp.MustCompile("[^a-zA-Z']+")
	sp := re.Split(s, -1)
	var ret string
	for _, str := range sp {
		ret += string(str[0])
	}
	return strings.ToUpper(ret)
}
