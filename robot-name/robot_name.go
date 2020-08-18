// Package robotname implements an random robot name generator.
package robotname

import (
	"math/rand"
	"time"
)

// Robot types store infomation about a robot.
type Robot struct {
	name string
}

var digits = []rune("0123456789")

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Name returns a formatted robot name.
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = r.randName()
	}

	return r.name, nil
}

// Reset generates a new name for the robot.
func (r *Robot) Reset() {
	r.name = r.randName()
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
