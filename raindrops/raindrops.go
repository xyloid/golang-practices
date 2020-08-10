/*
Package raindrops implemented Convert(), which converts a number into raindrop sounds.
*/
package raindrops

import "strconv"

// Convert a number into a string represeting rain drop sounds
func Convert(num int) string {
	ret := ""
	if num%3 == 0 {
		ret += "Pling"
	}
	if num%5 == 0 {
		ret += "Plang"
	}
	if num%7 == 0 {
		ret += "Plong"
	}
	if len(ret) == 0 {
		ret = strconv.Itoa(num)
	}
	return ret
}
