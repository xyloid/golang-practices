/*
Package romannumerals implements a function that convert a arabic number into roman numerals
*/
package romannumerals

import "fmt"

// ToRomanNumeral converts an arabic number into a roman number
func ToRomanNumeral(num int) (string, error) {
	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("Invalid number")
	}

	var table = map[int]string{1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ret := ""
	for i := 0; i < len(nums); i++ {
		if nums[i] <= num {
			num -= nums[i]
			ret += table[nums[i]]
			i--
		}
	}

	return ret, nil
}
