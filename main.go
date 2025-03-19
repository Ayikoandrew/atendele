package main

import (
	"time"

	"github.com/Ayikoandrew/atendele/network"
)

func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trRemote.Connect(trLocal)
	trLocal.Connect(trRemote)

	//actorSystem, _ := actor.NewEngine(actor.NewEngineConfig())

	//actorPID := actorSystem.Spawn(actors.NewActor(trLocal), "LOCAL TRANSPORT")

	go func() {
		for {
			data := []byte("Hello, World!")
			trRemote.SendMessage(trLocal.Addr(), data)
			time.Sleep(1 * time.Second)

		}
	}()
	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)

	s.Start()
}
