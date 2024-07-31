package p2p

import "net"

// represents the remote node
type Peer interface{
  Send([]byte) error
  RemoteAddr() net.Addr
  Close() error
}

// anything that handles the communication between the
// nodes in the network 
type Transport interface{
  Dial(string) error
  ListenAndAccept() error
  Consume() <-chan RPC
  Close() error
}
