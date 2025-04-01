package core

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/Ayikoandrew/atendele/types"
	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	res := types.BlockNumber()
	par := types.RandomHash()
	tx := []Transaction{}
	b := NewBlock(tx, &StateDiff{}, par, res)

	assert.Equal(t, b.BlockNumber, res)
	assert.Equal(t, b.ParentHash, par)
}

func TestBlockWithNilStateDiff(t *testing.T) {
	res := types.BlockNumber()
	par := types.RandomHash()
	tx := []Transaction{}

	b := NewBlock(tx, nil, par, res)

	assert.NotNil(t, b.StateDiff.Balance)
	assert.NotNil(t, b.StateDiff.Nonce)
	assert.NotNil(t, b.StateDiff.Storage)
}

func TestBlockTimestamp(t *testing.T) {
	before := uint64(time.Now().UnixNano())
	b := NewBlock(nil, nil, types.Hash{}, 1)

	after := uint64(time.Now().UnixNano())

	assert.GreaterOrEqual(t, b.Timestamp, before)
	assert.LessOrEqual(t, b.Timestamp, after)

}

type MockBlockEncoder struct{}

func (m *MockBlockEncoder) Encode(w io.Writer, b *Block) error {

	return nil
}

type MockBlockDecoder struct{}

func (m *MockBlockDecoder) Decode(r io.Reader, b *Block) (*Block, error) {
	return b, nil
}

func TestBlockEncoding(t *testing.T) {
	b := NewBlock(nil, nil, types.RandomHash(), 1)

	mockEncoder := &MockBlockEncoder{}
	mockDecoder := &MockBlockDecoder{}

	var buf bytes.Buffer

	err := b.Encode(&buf, mockEncoder)
	assert.NoError(t, err)

	newBlock := &Block{}
	err = newBlock.Decode(&buf, mockDecoder)
	assert.NoError(t, err)

}

type MockBlockHasher struct{}

func (m *MockBlockHasher) HashBlock(b *Block) types.Hash {
	return [32]byte{1, 2, 3}
}

func TestBlockHash(t *testing.T) {
	b := NewBlock(nil, nil, types.RandomHash(), 1)
	mockHasher := &MockBlockHasher{}

	hash := b.Hash(mockHasher)
	assert.NotEqual(t, types.Hash{}, hash)
}
