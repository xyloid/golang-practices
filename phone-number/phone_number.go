// Package phonenumber implements functions related to telephone numbers
package phonenumber

import (
	"fmt"
	"regexp"
)

// Number parses an input string
func Number(s string) (string, error) {
	number, err := validate(s)
	if err != nil {
		return "", err
	}
	return number, nil
}

// AreaCode returns the area code of the given number
func AreaCode(s string) (string, error) {
	number, err := validate(s)
	if err != nil {
		return "", err
	}
	return number[0:3], nil
}

// Format returns a number in formatted way.
func Format(s string) (string, error) {

	number, err := validate(s)
	if err != nil {
		return "", err
	}
	return "(" + string(number[0:3]) + ") " + string(number[3:6]) + "-" + string(number[6:]), nil
}

func countryCode(s string) (code, num string) {
	num = getNumber(s)

	if len(num) == 10 {
		return "", num
	} else {
		return string(num[0]), num[1:]
	}
}

func getNumber(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}

func validateLetters(s string) bool {
	reg, err := regexp.Compile(`[\d\s\+\.\-\(\)]+`)
	if err != nil {
		return false
	}
	return reg.ReplaceAllString(s, "") == ""
}

func validate(s string) (string, error) {
	if validateLetters(s) {
		country, number := countryCode(s)
		if country == "" || country == "1" {
			if len(number) != 10 {
				return "", fmt.Errorf("Invalid digit number")
			}
			if number[0] == '0' || number[0] == '1' || number[3] == '0' || number[3] == '1' {
				return "", fmt.Errorf("Invalid format of the telephone number")
			}

			return number, nil
		} else {
			return "", fmt.Errorf("Invalid country code")
		}

	} else {
		return "", fmt.Errorf("Invalid letters in the number")
	}
}
