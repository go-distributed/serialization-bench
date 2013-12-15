package bson

import (
	"fmt"
	"testing"
	"labix.org/v2/mgo/bson"
)

var _ = fmt.Printf

type Data struct {
	N int
	S string
	A []byte
}

func BenchmarkBson(b *testing.B) {
	length := 100
	a := make([]byte, length, length)

	putData := &Data{A: a}
	getData := &Data{}

	for i := 0; i < b.N; i++ {
	        putData.N = i;
	        putData.S = string(i)
		data, _ := bson.Marshal(putData)
		bson.Unmarshal(data, getData)
	}
}
