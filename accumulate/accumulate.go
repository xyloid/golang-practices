/*
Package accumulate implements Accumulate()
*/
package accumulate

// Accumulate function will apply fn onto collection arr.
func Accumulate(arr []string, fn func(string) string) []string {
	ret := make([]string, len(arr))
	copy(ret, arr)
	for i, s := range arr {
		ret[i] = fn(s)
	}
	return ret
}
