/*
Package reverse implements reverse
*/
package reverse

// Reverse reverses a string
func Reverse(str string) string {
	arr := []rune(str)
	var l, r = 0, len(arr) - 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		r--
		l++
	}
	return string(arr)
}
