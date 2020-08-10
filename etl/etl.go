/*
Package etl implements Extract-Transform-Load (ETL).
*/
package etl

import "strings"

// Transform implements ETL.
func Transform(given map[int][]string) map[string]int {
	var ret = map[string]int{}

	for score, letters := range given {
		for _, letter := range letters {
			ret[strings.ToLower(letter)] = score
		}
	}
	return ret
}
