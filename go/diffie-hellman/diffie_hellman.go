package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey randomize private key between (1, p)
func PrivateKey(p *big.Int) *big.Int {
	result, _ := rand.Int(rand.Reader, big.NewInt(p.Int64()).Sub(p, big.NewInt(3)))
	return result.Add(result, big.NewInt(2))
}

// PublicKey calculate public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

// NewPair generate public and private keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey calculate secret
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}
