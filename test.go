package main

import (
	"log"

	"github.com/sef-comp/distrfs/p2p"
)

func main(){
  Test1()
}

func Test1(){

  opts := p2p.TCPTransportOpts{
    ListenAddr: ":3000",
    ShakeHands: p2p.NOPHandshakeFunc,
    Decoder: p2p.GOBDecoder{},
  }
  tr := p2p.NewTCPTransport(opts)

  if err := tr.ListenAndAccept(); err != nil{
    log.Fatal(err)
  }
  
  select{}

  // fmt.Println("AGAZEGA!")


}
