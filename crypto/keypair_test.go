package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Keypair(t *testing.T) {
	priv := GeneratePrivateKey()

	pub := priv.GeneratePublicKey()

	data := []byte("Hello Jude")
	sig, err := priv.Sign(data)
	assert.Nil(t, err)

	res := pub.Verify(data, sig)

	assert.IsType(t, true, res)

	a := pub.public.X.String()
	fmt.Printf("address %+v", a)
}
