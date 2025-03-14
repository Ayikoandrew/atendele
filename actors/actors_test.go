package actors

import (
	"testing"

	"github.com/Ayikoandrew/atendele/network"
	"github.com/anthdm/hollywood/actor"
	"github.com/stretchr/testify/assert"
)

func TestActors_Receive(t *testing.T) {
	tra := network.NewLocalTransport("A")
	trb := network.NewLocalTransport("B")

	actorSystem, _ := actor.NewEngine(actor.NewEngineConfig())

	actorPID := actorSystem.Spawn(NewActor(tra), "Actor")

	data := []byte("Hello, World!")
	actorSystem.Send(actorPID, SendMessage{From: trb.Addr(), Payload: data})
	assert.Nil(t, tra.Connect(trb))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, data)
	assert.Equal(t, rpc.From, tra.Addr())
}
