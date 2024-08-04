package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/sef-comp/distrfs/p2p"
)

func main() {
	Test2()
}

func makeServer(listenAddr string, nodes ...string) *FileServer {

	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr: listenAddr,
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    p2p.DefaultDecoder{},
		// OnPeer: func(p2p.Peer) error,
	}

	tcptransport := p2p.NewTCPTransport(tcptransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcptransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)
	tcptransport.OnPeer = s.OnPeer

	return s
}

// func OnPeer(peer p2p.Peer) error {
// 	fmt.Println("Doing onpeer logic")
// 	peer.Close()
// 	return nil
// }

func Test1() {

	opts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		ShakeHands: p2p.NOPHandshakeFunc,
		Decoder:    p2p.DefaultDecoder{},
		// OnPeer:     OnPeer,
	}
	tr := p2p.NewTCPTransport(opts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n\tmessage:%s", msg, string(msg.Payload))
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}

	// fmt.Println("AGAZEGA!")

}

func Test2() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()


  time.Sleep(1 * time.Second)
	
  go s2.Start()

  time.Sleep(1 * time.Second)

	data := bytes.NewReader([]byte("my big data file here!"))

	s2.StoreData("myprivatekey", data)

  select{}


}
