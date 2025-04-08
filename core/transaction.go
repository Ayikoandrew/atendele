package core

import (
	"fmt"

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

func (tx *Transaction) Sign(priv crypto.PrivateKey) error {
	data := fmt.Sprintf("%+v %+v %+v", tx.To, tx.Value, tx.Data)
	sig, err := priv.Sign([]byte(data))
	if err != nil {
		return err
	}

	tx.signature = sig
	tx.publicKey = priv.PublicKey()
	return nil
}

func (tx *Transaction) Verify() error {
	data := fmt.Sprintf("%+v %+v %+v", tx.To, tx.Value, tx.Data)
	var zeroSig = &crypto.Signature{}
	if tx.signature == *zeroSig {
		return fmt.Errorf("transaction does not have a signature")
	}
	if !tx.publicKey.Verify([]byte(data), tx.signature) {
		return fmt.Errorf("signature verification failed")
	}

	return nil
}
