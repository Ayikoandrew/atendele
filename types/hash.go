package types

import (
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
