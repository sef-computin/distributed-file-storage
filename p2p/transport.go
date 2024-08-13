package p2p

import "net"

// represents the remote node
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// anything that handles the communication between the
// nodes in the network
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
