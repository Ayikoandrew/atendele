package core

import (
	"io"
	"math/big"
	"time"

	"github.com/Ayikoandrew/atendele/types"
)

type StateDiff struct {
	Balance map[types.Address]map[types.Address]*big.Int
	Nonce   map[types.Address]uint64
	Storage map[types.Address]map[string][]byte
}

type Block struct {
	BlockNumber        uint64
	Timestamp          uint64
	Transactions       []Transaction
	StateDiff          StateDiff
	BlockHash          types.Hash
	ParentHash         types.Hash
	SettlementMetadata uint64
}

func NewBlock(tx []Transaction) *Block {
	return &Block{
		BlockNumber:        1,
		Timestamp:          uint64(time.Now().UnixNano()),
		Transactions:       tx,
		StateDiff:          StateDiff{},
		BlockHash:          types.Hash{},
		ParentHash:         types.Hash{},
		SettlementMetadata: 1,
	}
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	_, err := dec.Decode(r, b)
	return err
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	return types.Hash(hasher.HashBlock(b))
}
