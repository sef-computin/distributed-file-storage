package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)


func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)

	expected_path := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	expected_filename := "6804429f74181a63c50c3d81d733a12f14a353ff"
	if pathname.Pathname != expected_path {
		t.Error(t, "have %s, want %s", pathname.Pathname, expected_path)
	}
	if pathname.Filename != expected_filename {
		t.Error(t, "have %s, want %s", pathname.Filename, expected_filename)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
  id := generateID()

	key := "momsspecials"

	data := []byte("some jpeg bytes")

	if _, err := s.writeStream(id, key, bytes.NewBuffer(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(id, key); err != nil {
		t.Error(err)
	}

}

func TestStore(t *testing.T) {
  count := 3
  s := newStore()
  id := generateID()
  defer teardown(t, s)

	// key := "foobar"
  for i:= 0; i<count; i++{
    key := fmt.Sprintf("foo_%d", i)
	  data := []byte("some jpeg bytes")
  	if _, err := s.writeStream(id, key, bytes.NewBuffer(data)); err != nil {
	  	t.Error(err)
	  }

	  if ok := s.Has(id, key); !ok{
		  t.Errorf("Expected to have key %s", key)
  	}

	  _, r, err := s.Read(id, key)
	  if err != nil {
		  t.Error(err)
	  }

	  b, err := ioutil.ReadAll(r)

	  fmt.Println(string(b))

	  if string(b) != string(data) {
		  t.Errorf("want %s, have %s", data, b)
  	}
	// fmt.Println(string(b))

    if err := s.Delete(id, key); err != nil{
      t.Error(err)
    }
    if ok := s.Has(id, key); ok{
      t.Errorf("expected not to have key %s", key)
    }
  }
}

func newStore() *Store{
  opts := StoreOpts{
    PathTransformFunc: CASPathTransformFunc,
  }
  return NewStore(opts)
}

func teardown(t *testing.T, s *Store){
  if err := s.Clear(); err != nil{
    t.Error(err)
  }
}
