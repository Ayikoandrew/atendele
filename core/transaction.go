package core

import "github.com/Ayikoandrew/atendele/types"

type Transaction struct {
	From  types.Address
	To    types.Address
	Value []byte
	Data  []byte
}
