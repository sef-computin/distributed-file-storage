package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	// if we dial, outbound = true
	// if we accept, outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr string
	ShakeHands HandshakeFunc
	Decoder    Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error %s\n:", err)
		}

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.ShakeHands(peer); err != nil {
		conn.Close()
		return
	}

	msg := &Temp{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %v\n", err)
			continue
		}
	}

	// fmt.Printf("new incoming connection %+v\n", conn)
}
