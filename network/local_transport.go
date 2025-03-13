package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	peer, ok := tr.(*LocalTransport)
	if !ok {
		return fmt.Errorf("transport is invalid")
	}
	t.peers[tr.Addr()] = peer
	fmt.Printf("Connected %s to %s\n", t.addr, peer.addr)
	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, data []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()
	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s is not connected to %s", to, t.addr)
	}

	peer.consumeCh <- RPC{
		From:    t.Addr(),
		Payload: data,
	}
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
