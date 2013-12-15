package gobson

import (
	"labix.org/v2/mgo/bson"
)

func Marshal(in interface{}) ([]byte, error) {
	data, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Unmarshal(in []byte, out interface{}) error {
	return bson.Unmarshal(in, out)
}
