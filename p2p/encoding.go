package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface{
  Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{
}

func (dec GOBDecoder) Decode(r io.Reader, rpc *RPC) error{
  return gob.NewDecoder(r).Decode(rpc.Payload)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, rpc *RPC) error{
  buf := make([]byte, 1028)

  n, err := r.Read(buf)
  if err != nil{
    return err
  }

  // fmt.Println(string(buf[:n]))
  rpc.Payload = buf[:n]

  return nil
}

