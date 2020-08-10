/*
Package lsproduct implements a function that finds the largest product of N contiguous numbers
*/
package lsproduct

import (
	"fmt"
	"unicode"
)

// LargestSeriesProduct finds the largest series product
func LargestSeriesProduct(digits string, span int) (int, error) {

	if span < 0 {
		return -1, fmt.Errorf("Expect non-negative span")
	}

	numbers := make([]int, 0)

	for _, digit := range digits {
		if unicode.IsDigit(digit) {
			numbers = append(numbers, int(digit-'0'))
		} else {
			return -1, fmt.Errorf("Invalid digit")
		}

	}

	if len(numbers) < span {
		return -1, fmt.Errorf("Span can not be larger than the digits number")
	}

	if len(numbers) == 0 && span == 0 {
		return 1, nil
	}

	count := 0
	product := 1
	maxProduct := 0
	index := 0
	for index < len(numbers) {
		if numbers[index] == 0 {
			product = 1
			count = 0
		} else {
			if count == span {
				product /= numbers[index-count]
				product *= numbers[index]
				if product > maxProduct {
					maxProduct = product
				}

			} else {
				product *= numbers[index]
				count++
				if count == span {
					if product > maxProduct {
						maxProduct = product
					}
				}
			}
		}
		index++
	}
	return maxProduct, nil
}
