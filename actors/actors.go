package actors

import (
	"fmt"
	"log/slog"

	"github.com/Ayikoandrew/atendele/network"
	"github.com/anthdm/hollywood/actor"
)

type SendMessage struct {
	To      network.NetAddr
	Payload []byte
}

type Actor struct {
	Transports  []network.Transport
	TransportID *actor.PID
}

func NewActor(tr network.Transport) actor.Producer {
	return func() actor.Receiver {
		return &Actor{
			Transports: []network.Transport{tr},
		}
	}
}

// Receive implements actor.Receiver.
func (a *Actor) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		slog.Info("Transport actor has started")
		_ = msg
	case SendMessage:
		fmt.Println("Sending RPC...")
		for _, tr := range a.Transports {
			if err := tr.SendMessage(msg.To, msg.Payload); err != nil {
				slog.Error("Failed to send message", "error", err)
			}
		}
	}
}
