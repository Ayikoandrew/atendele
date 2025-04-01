package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (a Address) ToSlice() []byte {
	v := make([]byte, 20)

	for i := 0; i < 20; i++ {
		v[i] = a[i]
	}

	return v
}

func (a *Address) String() string {
	return hex.EncodeToString(a[:])
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("given bytes with len %d should be 20", len(b))
		panic(msg)
	}

	var v [20]byte
	for i := 0; i < 10; i++ {
		v[i] = b[i]
	}

	return Address(v)
}
