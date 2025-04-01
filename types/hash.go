package types

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log/slog"
)

type Hash [32]uint8

func HashFromByte(d []byte) (Hash, error) {
	if len(d) != 32 {
		slog.Info("byte is not of length 32")
	}

	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = d[i]
	}
	return Hash(value), nil
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	hash, err := HashFromByte(RandomBytes(32))
	if err != nil {
		fmt.Printf("error generating 32 bytes")
	}
	return hash
}

func BlockNumber() uint64 {
	buf := make([]byte, 8)
	rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
