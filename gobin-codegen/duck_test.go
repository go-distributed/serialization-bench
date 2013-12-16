package duck

import (
	"bytes"
	"fmt"
	"testing"
)

var _ = fmt.Printf

func BenchmarkMars(b *testing.B) {
	buf := new(bytes.Buffer)
	q := &Quack{X: 1, Y: [3]byte{'a', 'b', 'c'}}

	for i := 0; i < 100; i++ {
		q.Z[i] = byte(i)
	}
	for i := 0; i < b.N; i++ {
		q.Marshal(buf)
	}
}
func BenchmarkUnma(b *testing.B) {
	buf := new(bytes.Buffer)
	q := &Quack{X: 1, Y: [3]byte{'a', 'b', 'c'}}
	for i := 0; i < 100; i++ {
		q.Z[i] = byte(i)
	}

	q.Marshal(buf)

	q2 := &Quack{}

	for i := 0; i < b.N; i++ {
		q2.Marshal(buf)
	}
}
