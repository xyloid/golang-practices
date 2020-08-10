// Package secret implements a secret handshake protocol
package secret

// Handshake converts a number into handshake actions
func Handshake(n uint) []string {
	res := make([]string, 0)

	isReversed := n > 16

	for bit := uint(1); bit < 16; bit = bit << 1 {
		if n&bit > 0 {
			if isReversed {
				res = append([]string{actions[bit]}, res...)
			} else {
				res = append(res, actions[bit])
			}
		}
	}
	return res
}

var actions = map[uint]string{
	1:      "wink",
	1 << 1: "double blink",
	1 << 2: "close your eyes",
	1 << 3: "jump"}
