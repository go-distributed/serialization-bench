package bson

import (
	"fmt"
	math_rand "math/rand"
	"testing"
	"time"

	"labix.org/v2/mgo/bson"
)

var _ = fmt.Printf
var _ = time.Now

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

func TestBson(t *testing.T) {
	p := &PreAccept{LeaderId: 1}

	data, err := bson.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	msg := &PreAccept{}
	if err := bson.Unmarshal(data, msg); err != nil {
		t.Fatal(err)
	}

	if msg.LeaderId != p.LeaderId {
		t.Fatal("not equal")
	}
}

func BenchmarkBsonMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		p := NewPopulatedPreAccept(popr)

		b.StartTimer()
		data, err := bson.Marshal(p)
		if err != nil {
			panic(err)
		}
		b.StopTimer()

		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkBsonUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	total := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		p := NewPopulatedPreAccept(popr)

		data, err := bson.Marshal(p)
		if err != nil {
			panic(err)
		}

		total += len(data)

		msg := &PreAccept{}

		b.StartTimer()
		if err := bson.Unmarshal(data, msg); err != nil {
			panic(err)
		}

	}
	b.SetBytes(int64(total / b.N))
}
