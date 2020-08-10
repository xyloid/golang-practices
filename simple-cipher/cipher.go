package cipher

// Cipher interface for any algorithm
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
