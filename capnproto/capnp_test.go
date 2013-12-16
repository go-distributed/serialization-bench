package capnproto

import (
	"bytes"
	"testing"

	"github.com/jmckaskill/go-capnproto"
)

type Data struct {
	N int32
	S string
	A []byte
}

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

func BenchmarkCapnprotoMarshal(b *testing.B) {
	length := 100
	a := make([]byte, length, length)
	buf := bytes.Buffer{}

	putData := Data{N: 1, S: "abc", A: a}

	for i := 0; i < b.N; i++ {
		s := capn.NewBuffer(nil)
		st := NewRootTestST(s)
		st.SetN(putData.N)
		st.SetS(putData.S)
		st.SetA(putData.A)
		s.WriteTo(&buf)
		buf.Reset()
	}
}

func BenchmarkCapnprotoUnmarshal(b *testing.B) {
	length := 100
	a := make([]byte, length, length)
	buf := bytes.Buffer{}

	putData := Data{N: 1, S: "abc", A: a}
	s := capn.NewBuffer(nil)
	st := NewRootTestST(s)
	st.SetN(putData.N)
	st.SetS(putData.S)
	st.SetA(putData.A)
	s.WriteTo(&buf)

	for i := 0; i < b.N; i++ {
		s, _, err := capn.ReadFromMemoryZeroCopy(buf.Bytes())
		if err != nil {
			panic(i)
		}
		st = ReadRootTestST(s)
		if st.N() != 1 || st.S() != "abc" {
			panic(i)
		}
	}

}
