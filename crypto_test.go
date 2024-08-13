package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCopyEncryptDecrypt(t *testing.T){
  payload := "AGAZEGA"
  src := bytes.NewReader([]byte(payload))
  dst := new(bytes.Buffer)
  key := newEncryptionKey()

  _, err := copyEncrypt(key, src, dst)
  if err != nil{
    t.Error(err)
  }
  fmt.Println(dst.Bytes())

  out := new(bytes.Buffer)
  nw, err := copyDecrypt(key, dst, out)
  if err != nil{
    t.Error(err)
  }

  if nw != len(payload)+16{
    t.Fail()
  } 

  if payload != out.String(){
    t.Errorf("encryption changed the file")
  }

  fmt.Println(out.String())
}
