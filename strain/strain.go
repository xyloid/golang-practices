/*
Package strain implements keep and discard functions
*/
package strain

// Ints is a list of integers
type Ints []int

// Keep keeps certain items
func (l Ints) Keep(fn func(int) bool) Ints {
	var ret []int
	for _, num := range l {
		if fn(num) {
			ret = append(ret, num)
		}
	}
	return Ints(ret)
}

// Discard discards certain items
func (l Ints) Discard(fn func(int) bool) Ints {
	var ret []int
	for _, num := range l {
		if !fn(num) {
			ret = append(ret, num)
		}
	}
	return Ints(ret)
}

// Strings is a list of strings
type Strings []string

// Keep keeps certain items
func (l Strings) Keep(fn func(string) bool) Strings {
	var ret []string
	for _, num := range l {
		if fn(num) {
			ret = append(ret, num)
		}
	}
	return Strings(ret)
}

// Discard discards certain items
func (l Strings) Discard(fn func(string) bool) Strings {
	var ret []string
	for _, num := range l {
		if !fn(num) {
			ret = append(ret, num)
		}
	}
	return Strings(ret)
}

// Lists is a list of lists
type Lists [][]int

// Keep keeps certain items
func (l Lists) Keep(fn func([]int) bool) Lists {
	var ret [][]int
	for _, num := range l {
		if fn(num) {
			ret = append(ret, num)
		}
	}
	return Lists(ret)
}

// Discard discards certain items
func (l Lists) Discard(fn func([]int) bool) Lists {
	var ret [][]int
	for _, num := range l {
		if !fn(num) {
			ret = append(ret, num)
		}
	}
	return Lists(ret)
}
