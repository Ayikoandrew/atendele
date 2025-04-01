package core

import (
	"github.com/Ayikoandrew/atendele/crypto"
	"github.com/Ayikoandrew/atendele/types"
)

type Transaction struct {
	From  types.Address
	To    types.Address
	Value []byte
	Data  []byte

	publicKey crypto.PublicKey
	signature crypto.Signature
}

func (tx Transaction) Sign(priv crypto.PrivateKey) error {
	priv.Sign(tx.Data)
	return nil
}

func (tx Transaction) Verify() error {
	tx.publicKey.Verify(tx.Data, tx.signature)

	return nil
}
