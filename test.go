package main

import (
	"fmt"
	"log"

	"github.com/sef-comp/distrfs/p2p"
)

func main(){
  Test1()
}

func OnPeer(peer p2p.Peer) error{
  fmt.Println("Doing onpeer logic")
  peer.Close()
  return nil
}

func Test1(){

  opts := p2p.TCPTransportOpts{
    ListenAddr: ":3000",
    ShakeHands: p2p.NOPHandshakeFunc,
    Decoder: p2p.DefaultDecoder{},
    OnPeer: OnPeer,
  }
  tr := p2p.NewTCPTransport(opts)


  go func(){
    for{
      msg := <-tr.Consume()
      fmt.Printf("%+v\n\tmessage:%s", msg, string(msg.Payload))
    }
  }()

  if err := tr.ListenAndAccept(); err != nil{
    log.Fatal(err)
  }
  
  select{}

  // fmt.Println("AGAZEGA!")


}
