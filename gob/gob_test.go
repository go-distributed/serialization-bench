package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

var _ = fmt.Printf

type Data struct {
	N int
	S []byte
	A []byte
}

func TestGobEncoding(t *testing.T) {
	length := 100
	a := make([]byte, length, length)
	for i := 0; i < length; i++ {
		a[i] = byte(i)
	}

	putData := &Data{N: 1, S: []byte{'a', 'b', 'c'}, A: a}

	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	enc.Encode(putData)

	getData := &Data{}
	dec := gob.NewDecoder(m)
	dec.Decode(&getData)

	if getData.N != putData.N ||
		bytes.Compare(getData.S, putData.S) != 0 ||
		bytes.Compare(getData.A, putData.A) != 0 {
		t.Fatal("Gob encoding failed!")
	}
}

func BenchmarkMars(b *testing.B) {
	length := 100
	a := make([]byte, length, length)
	for i := 0; i < length; i++ {
		a[i] = byte(i)
	}

	putData := &Data{N: 1, S: []byte{'a', 'b', 'c'}, A: a}

	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	for i := 0; i < b.N; i++ {
		enc.Encode(putData)
	}
}

func BenchmarkUnma(b *testing.B) {

	length := 100
	a := make([]byte, length, length)
	for i := 0; i < length; i++ {
		a[i] = byte(i)
	}

	putData := &Data{N: 1, S: []byte{'a', 'b', 'c'}, A: a}

	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	for i := 0; i < b.N; i++ {
		enc.Encode(putData)
	}

	b.ResetTimer()

	getData := &Data{}
	dec := gob.NewDecoder(m)
	for i := 0; i < b.N; i++ {
		dec.Decode(&getData)
	}
}
