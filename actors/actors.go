package actors

import (
	"log/slog"

	"github.com/Ayikoandrew/atendele/network"
	"github.com/anthdm/hollywood/actor"
)

type ConnectPeer struct {
	Peer network.Transport
}

type SendMessage struct {
	From    network.NetAddr
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
	case ConnectPeer:
		for _, tr := range a.Transports {
			if err := tr.Connect(msg.Peer); err != nil {
				slog.Error("Failed to connect to peer")
			}
		}
	case SendMessage:
		for _, tr := range a.Transports {
			if err := tr.SendMessage(msg.From, msg.Payload); err != nil {
				slog.Error("Failed to send message", "error", err)
			}
		}
	}
}
