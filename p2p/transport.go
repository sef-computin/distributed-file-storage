package p2p

// represents the remote node
type Peer interface{
  Close() error
}

// anything that handles the communication between the
// nodes in the network 
type Transport interface{
  ListenAndAccept() error
  Consume() <-chan RPC
}
