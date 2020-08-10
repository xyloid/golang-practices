// Package diffiehellman implements DH cryption
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

// PrivateKey calculates private key
func PrivateKey(p *big.Int) *big.Int {
	var res, key big.Int
	// range is (1,p), so we need [0,p-2) which has 0...p-3 range, that is 2...p-1 range in the end
	key.Sub(p, big.NewInt(2))
	res.Rand(rand.New(rand.NewSource(time.Now().UnixNano())), &key)
	return res.Add(&res, big.NewInt(2))
}

// PublicKey generates the public key
func PublicKey(a, p *big.Int, g int64) *big.Int {
	var res big.Int
	var m big.Int
	m.SetInt64(g)
	res.Exp(&m, a, p)
	return &res
}

// SecretKey generates the secret key
func SecretKey(a, B, p *big.Int) *big.Int {
	var res big.Int
	res.Exp(B, a, p)
	return &res
}

// NewPair creates a new pair of keys
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	key := PublicKey(a, p, g)
	return a, key
}
