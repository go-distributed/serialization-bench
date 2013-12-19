package gobin

import (
	"bytes"
	"fmt"
	math_rand "math/rand"
	"testing"
	"time"
)

var _ = fmt.Printf

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

func TestGobin(t *testing.T) {
	p := &PreAccept{LeaderId: 1}
	data := new(bytes.Buffer)
	p.Marshal(data)
	msg := &PreAccept{}
	msg.Unmarshal(data)

	if msg.LeaderId != p.LeaderId {
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

		b.StartTimer()
		p.Marshal(data)
		b.StopTimer()

		total += len(data.Bytes())
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkUnma(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		p := NewPopulatedPreAccept(popr)

		data := new(bytes.Buffer)

		p.Marshal(data)

		total += len(data.Bytes())
		msg := &PreAccept{}

		b.StartTimer()
		msg.Unmarshal(data)

	}
	b.SetBytes(int64(total / b.N))
}
