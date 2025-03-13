package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalTransport_Addr(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.Addr(), NetAddr("A"))
	assert.Equal(t, trb.Addr(), NetAddr("B"))

	assert.NotEqual(t, tra.Addr(), trb.Addr())
	assert.NotEqual(t, trb.Addr(), tra.Addr())

}

func TestLocalTransport_Connect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	assert.Nil(t, tra.Connect(trb))
	assert.Nil(t, trb.Connect(tra))

	trc := NewLocalTransport("C")
	trb.Connect(trc)
	assert.Nil(t, tra.Connect(trc))
}

func TestLocalTransport_SendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	assert.Nil(t, tra.Connect(trb))
	assert.Nil(t, trb.Connect(tra))

	data := []byte("Hello, World!")
	assert.Nil(t, tra.SendMessage(trb.Addr(), data))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.From, tra.Addr())
	assert.Equal(t, rpc.Payload, data)
}
