package main

import (
	"fmt"
	"log"
	"time"

	"github.com/sef-comp/distrfs/p2p"
)

func main(){
  Test2()
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


func Test2(){
  tcptransportOpts := p2p.TCPTransportOpts{
    ListenAddr: ":3000",
    ShakeHands: p2p.NOPHandshakeFunc,
    Decoder: p2p.DefaultDecoder{},
    // OnPeer: func(p2p.Peer) error,
  }

  tcptransport := p2p.NewTCPTransport(tcptransportOpts)

  fileServerOpts := FileServerOpts{
    StorageRoot: "3000_network",
    PathTransformFunc: CASPathTransformFunc,
    Transport: tcptransport,
    BootstrapNodes: []string{":4000"},
  }

  s := NewFileServer(fileServerOpts)
 
  go func(){
    time.Sleep(time.Second * 3)
    s.Stop()
  }()
 
  if err := s.Start(); err != nil{
    log.Fatal(err)
  }


}
