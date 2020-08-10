/*
Package twofer implements the function ShareWith
*/
package twofer

// ShareWith returns a string message.
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
