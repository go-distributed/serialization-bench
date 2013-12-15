package gobson

import (
	"fmt"
	"testing"
)

var _ = fmt.Printf

type Data struct {
	N int
	S string
	A []byte
}

func TestGobson(t *testing.T) {

	length := 10
	a := make([]byte, length, length)

	putData := &Data{N: 1, S: "abc", A: a}

	data, err := Marshal(putData)
	if err != nil {
		t.Errorf("%v", err)
	}
	//fmt.Printf("%q\n", data)

	getData := &Data{}

	err = Unmarshal(data, getData)
	if err != nil {
		t.Errorf("%v", err)
	}

	if getData.N != putData.N || getData.S != putData.S {
		t.Errorf("Unmarshalled values doesn't equal to original values")
	}

	//fmt.Printf("%v\n", getData)
	for i := 0; i < length; i++ {
		if getData.A[i] != putData.A[i] {
			t.Errorf("Unmarshalled values doesn't equal to original values")
		}
	}

}
