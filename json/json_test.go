package json

import (
        "encoding/json"
        "fmt"
        "testing"
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

        putData := Data{N: 1, S: "abc", A: a}

        for i := 0; i < b.N; i++ {
                json.Marshal(putData)
        }
}

func BenchmarkBsonUnmarshal(b *testing.B) {
        length := 100
        a := make([]byte, length, length)

        putData := &Data{N: 1, S: "abc", A: a}
        getData := &Data{}

        data, _ := json.Marshal(putData)

        for i := 0; i < b.N; i++ {
                json.Unmarshal(data, getData)
        }
}
