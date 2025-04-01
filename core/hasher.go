package core

import (
	"encoding/binary"
	"fmt"

	"github.com/Ayikoandrew/atendele/types"
	"golang.org/x/crypto/sha3"
	"google.golang.org/protobuf/proto"
)

type Hasher[T any] interface {
	HashBlock(T) types.Hash
}

type BlockHasher struct{}

func HashBlock(d *Block) (types.Hash, error) {

	pbBlock := &types.Block{
		BlockNumber:        d.BlockNumber,
		Timestamp:          d.Timestamp,
		Transactions:       make([]*types.Transaction, len(d.Transactions)),
		StateDiff:          make([]*types.StateDiff, 0),
		SettlementMetadata: d.SettlementMetadata,
	}

	for i, tx := range d.Transactions {
		pbTx := types.Transaction{
			From: tx.From[:],
			To:   tx.To[:],
			Data: tx.Data,
		}

		if tx.Value != nil {
			pbTx.Value = tx.Value
		} else {
			pbTx.Value = []byte{}
		}
		pbBlock.Transactions[i] = &pbTx
	}

	stateDiff := &types.StateDiff{
		BalanceUpdate: make([]*types.BalanceUpdate, 0),
		NonceUpdate:   make([]*types.NonceUpdate, 0),
		StorageUpdate: make([]*types.StorageUpdate, 0),
	}

	for acc, tokenBalance := range d.StateDiff.Balance {
		for token, balance := range tokenBalance {
			stateDiff.BalanceUpdate = append(stateDiff.BalanceUpdate, &types.BalanceUpdate{
				Account: acc[:],
				Token:   token[:],
				Balance: balance.Bytes(),
			})
		}
	}

	for addr, nonce := range d.StateDiff.Nonce {
		nonceBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(nonceBytes, nonce)
		stateDiff.NonceUpdate = append(stateDiff.NonceUpdate, &types.NonceUpdate{

			Account: addr[:],
			Nonce:   nonceBytes,
		})
	}

	for addr, storage := range d.StateDiff.Storage {
		for key, value := range storage {
			stateDiff.StorageUpdate = append(stateDiff.StorageUpdate, &types.StorageUpdate{
				Account: addr[:],
				SlotKey: []byte(key),
				Value:   value,
			})
		}
	}

	b, err := proto.Marshal(pbBlock)
	if err != nil {
		return types.Hash{}, err
	}

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(b)
	fmt.Println("Protobuf encoded size:", len(b))

	h := hasher.Sum(nil)

	return types.Hash(h), nil

}

// func HashBlock(d *Block) (types.Hash, error) {
// 	buf := bytes.Buffer{}

// 	enc := gob.NewEncoder(&buf)
// 	if err := enc.Encode(d); err != nil {
// 		return types.Hash{}, err
// 	}

// 	fmt.Println("Gob encoded size:", len(buf.Bytes()))

// 	hasher := sha3.NewLegacyKeccak256()
// 	_, err := hasher.Write(buf.Bytes())
// 	if err != nil {
// 		return types.Hash{}, err
// 	}
// 	h := hasher.Sum(nil)

// 	return types.Hash(h), nil
// }
