package bson

import (
        "fmt"
        "testing"

        "labix.org/v2/mgo/bson"
)

var _ = fmt.Printf

type Data struct {
        N       int
        S       string
        A       []byte
}

func BenchmarkBsonMarshal(b *testing.B) {
        length := 100
        a := make([]byte, length, length)

        putData := &Data{A: a}
        putData.N = 1
        putData.S = "abc"

        for i := 0; i < b.N; i++ {
                bson.Marshal(putData)
        }
}

func BenchmarkBsonUnmarshal(b *testing.B) {
        length := 100
        a := make([]byte, length, length)

        putData := &Data{A: a}
        putData.N = 1
        putData.S = "abc"
        getData := &Data{}

        data, _ := bson.Marshal(putData)

        for i := 0; i < b.N; i++ {
                bson.Unmarshal(data, getData)
        }
}
