// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if isInvalidLength(a) || isInvalidLength(b) || isInvalidLength(c) {
		return NaT
	}

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	if a+b < c {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}

	if a == b || b == c {
		return Iso
	}

	return Sca
}

func isInvalidLength(l float64) bool {
	return l <= 0 || math.IsNaN(l) || math.IsInf(l, 0)
}
