package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
    EncKey:            newEncryptionKey(),
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
	s2 := makeServer(":7000", "")
	s3 := makeServer(":6000", ":3000", ":7000")

	go func() { log.Fatal(s1.Start()) }()
	time.Sleep(500 * time.Millisecond)
	go func() { log.Fatal(s2.Start()) }()

	time.Sleep(2 * time.Second)

	go s3.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("picture_%d.png", i)
		data := bytes.NewReader([]byte("my big data file here!"))
		s3.Store(key, data)

		if err := s3.store.Delete(s3.ID, key); err != nil {
			log.Fatal(err)
		}

		r, err := s3.Get(key)
		if err != nil {
			log.Fatal(err)
		}

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}
