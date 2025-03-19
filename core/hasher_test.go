package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ayikoandrew/atendele/types"
)

func TestHashBlock(t *testing.T) {
	b := &Block{
		BlockNumber:  1,
		Timestamp:    uint64(time.Now().UnixNano()),
		Transactions: []Transaction{},
		StateDiff:    StateDiff{},
		ParentHash:   types.Hash{},
	}

	data, _ := HashBlock(b)

	fmt.Println(data)
	fmt.Print()
	fmt.Printf("Size of bytes using grpc is %d", len(data))
}
