package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("Data received: %+v\n", rpc)

		case <-s.quitCh:
			fmt.Println("Received quit signal, shutting down")
			break free
		case <-ticker.C:
			fmt.Println("Every five seconds do something")
		}
	}
	fmt.Println("Server shutting down")

}

func (s *Server) Stop() {
	close(s.quitCh)
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
