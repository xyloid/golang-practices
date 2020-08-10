/*
Package pythagorean implement triangle related functions
*/
package pythagorean

// Triplet is an array of 3 ints
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	res := make([]Triplet, 0)
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					res = append(res, [3]int{a, b, c})
				}
			}
		}
	}
	return res
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	res := make([]Triplet, 0)
	for a := 1; a < p/2; a++ {
		for b := a; b < p/2; b++ {
			c := p - b - a
			if a*a+b*b == c*c {
				res = append(res, [3]int{a, b, c})
			}
		}
	}

	return res
}
