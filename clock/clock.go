// Package clock implements a clock function
package clock

import "fmt"

// Clock object struct
type Clock struct {
	minutes int
}

// New creates a new clock object with hours and minutes
func New(h, m int) Clock {
	// m is in the range [0,24*60)
	m += h * 60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}

	return Clock{m}
}

// Add will add delta minutes to the time c
func (c Clock) Add(delta int) Clock {
	return New(0, c.minutes+delta)
}

// Subtract will subtract delta minutes from the time c
func (c Clock) Subtract(delta int) Clock {
	return New(0, c.minutes-delta)
}

// String returns a string representing the time c
func (c Clock) String() string {
	h := c.minutes / 60
	m := c.minutes % 60
	return fmt.Sprintf("%02v:%02v", h, m)
}
