package p2p

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func OnPeer(Peer) error{
  fmt.Println("failed the onpeer func")
  return nil
}

func TestTCPTransport(t *testing.T){
  opts := TCPTransportOpts{
    ListenAddr: ":3000",
    ShakeHands: NOPHandshakeFunc,
    Decoder: DefaultDecoder{},
    OnPeer: OnPeer,
  }


  tr := NewTCPTransport(opts)
  
  assert.Equal(t, tr.ListenAddr, ":3000")


  assert.Nil(t, tr.ListenAndAccept())
}
