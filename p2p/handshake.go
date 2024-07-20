package p2p

import "errors"

var ErrInvalidHandshake = errors.New("Invalid handshake")

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
