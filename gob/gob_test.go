package gob

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"runtime"
	math_rand "math/rand"
	"testing"
	"time"
)

var _ = fmt.Printf

type PreAccept struct {
	LeaderId int32
	Replica  int32
	Instance int32
	Ballot   int32
	Command  []byte
	Seq      int32
	Deps     [5]int32
}

func NewPopulatedPreAccept(r *math_rand.Rand) *PreAccept {
	this := &PreAccept{}
	this.LeaderId = r.Int31()
	this.Replica = r.Int31()
	this.Instance = r.Int31()
	this.Ballot = r.Int31()
	v1 := r.Intn(100)
	this.Command = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Command[i] = byte(r.Intn(256))
	}
	this.Seq = r.Int31()
	for i := 0; i < 5; i++ {
		this.Deps[i] = r.Int31()
	}
	return this
}

func TestGob(t *testing.T) {

	p := &PreAccept{LeaderId: 1}

	data := new(bytes.Buffer)
	enc := gob.NewEncoder(data)
	dec := gob.NewDecoder(data)

	msg := &PreAccept{}
	msg2 := &PreAccept{}

	enc.Encode(p)
	dec.Decode(&msg)

	enc.Encode(p)
	dec.Decode(&msg2)

	if msg.LeaderId != p.LeaderId || msg2.LeaderId != p.LeaderId {
		t.Fatal("not equal")
	}
}

func BenchmarkMars(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		p := NewPopulatedPreAccept(popr)
		data := new(bytes.Buffer)
		enc := gob.NewEncoder(data)

		b.StartTimer()
		enc.Encode(p)
		b.StopTimer()

		total += len(data.Bytes())

	}

	b.SetBytes(int64(total / b.N))
}

func BenchmarkUnma(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	data := new(bytes.Buffer)
	enc := gob.NewEncoder(data)
	dec := gob.NewDecoder(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		p := NewPopulatedPreAccept(popr)

		enc.Encode(p)

		total += len(data.Bytes())

		msg := &PreAccept{}

		b.StartTimer()
		dec.Decode(&msg)

	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkStreamMars(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	data := new(bytes.Buffer)
	enc := gob.NewEncoder(data)

	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		p := NewPopulatedPreAccept(popr)

		b.StartTimer()
		enc.Encode(p)
		b.StopTimer()

	}

	total += len(data.Bytes())
	b.SetBytes(int64(total / b.N))
}

func BenchmarkStreamUnma(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	data := new(bytes.Buffer)
	enc := gob.NewEncoder(data)
	dec := gob.NewDecoder(data)

	for i := 0; i < b.N; i++ {
		p := NewPopulatedPreAccept(popr)
		enc.Encode(p)

	}
	total += len(data.Bytes())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()

                runtime.GC() // it speeds up by 3000 ns/op.
		msg := &PreAccept{}

		b.StartTimer()
		dec.Decode(&msg)

	}
	b.SetBytes(int64(total / b.N))
}
