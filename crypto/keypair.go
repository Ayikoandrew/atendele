package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/Ayikoandrew/atendele/types"
)

type PrivateKey struct {
	private *ecdsa.PrivateKey
}

func (priv *PrivateKey) Sign(d []byte) (Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, priv.private, d)
	if err != nil {
		return Signature{}, err
	}
	return Signature{r: r, s: s}, nil
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		private: key,
	}
}

type PublicKey struct {
	public *ecdsa.PublicKey
}

func (p *PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		public: &p.private.PublicKey,
	}
}

func (p *PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), p.public.X, p.public.Y)
}

func (p *PublicKey) Address() types.Address {
	b := sha256.Sum256(p.ToSlice())
	return types.AddressFromBytes(b[len(b)-20:])
}

type Signature struct {
	r, s *big.Int
}

func (pub *PublicKey) Verify(d []byte, s Signature) bool {
	return ecdsa.Verify(pub.public, d, s.r, s.s)
}
