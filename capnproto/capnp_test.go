package capnproto

import (
	"bytes"
	"testing"

	"github.com/jmckaskill/go-capnproto"
)

func TestCapnproto(t *testing.T) {
	// Write
	s := capn.NewBuffer(nil)
	st := NewRootTestST(s)
	st.SetN(1)
	st.SetS("hello world")
	st.SetA([]byte("hello world"))
	buf := bytes.Buffer{}
	s.WriteTo(&buf)

	// Read
	s, err := capn.ReadFromStream(&buf, nil)
	if err != nil {
		t.Fatal(err)
	}
	st = ReadRootTestST(s)

	if st.N() != 1 || st.S() != "hello world" || string(st.A()) != "hello world" {
		t.Fatal("error on decoding capnproto")
	}
}
