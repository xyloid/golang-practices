// Package robotname implements an random robot name generator.
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot types store infomation about a robot.
type Robot struct {
	name string
}

const nameLimit = 26 * 26 * 10 * 10 * 10

var records = make(map[string]bool)

var digits = []rune("0123456789")

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Name returns a formatted robot name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(records) == nameLimit {
			return "", errors.New("Name exhuasted")
		}
		r.name = r.randName()

		for records[r.name] == true {
			r.name = r.randName()
		}
		fmt.Println(len(records))
		records[r.name] = true
	}

	return r.name, nil
}

// Reset generates a new name for the robot.
func (r *Robot) Reset() {
	r.name = ""
}

func (r *Robot) randName() string {
	randRune := make([]rune, 5)
	for i := range randRune {
		rand.Seed(time.Now().UnixNano())
		if i < 2 {
			randRune[i] = letters[rand.Intn(len(letters))]
		} else {
			randRune[i] = digits[rand.Intn(len(digits))]
		}
	}

	return string(randRune)
}
