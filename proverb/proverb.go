package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var ret []string

	for i := 0; i < len(rhyme)-1; i++ {
		current := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		ret = append(ret, current)
	}

	ret = append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return ret
}
