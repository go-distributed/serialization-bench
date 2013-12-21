package gogopb

import (
        "encoding/json"
        "unsafe"
)

type Uint128 [3]uint64

func (u Uint128) Marshal() ([]byte, error) {
        buffer := make([]byte, 24)
        u.MarshalTo(buffer)
        return buffer, nil
}

func (u Uint128) MarshalTo(data []byte) (n int, err error) {
        PutLittleEndianUint128(data, 0, u)
        return 24, nil
}

func GetLittleEndianUint64(b []byte, offset int) uint64 {
        return *(*uint64)(unsafe.Pointer(&b[offset]))
}

func PutLittleEndianUint64(b []byte, offset int, v uint64) {
        b[offset] = byte(v)
        b[offset+1] = byte(v >> 8)
        b[offset+2] = byte(v >> 16)
        b[offset+3] = byte(v >> 24)
        b[offset+4] = byte(v >> 32)
        b[offset+5] = byte(v >> 40)
        b[offset+6] = byte(v >> 48)
        b[offset+7] = byte(v >> 56)
}

func PutLittleEndianUint128(buffer []byte, offset int, v [3]uint64) {
        PutLittleEndianUint64(buffer, offset, v[0])
        PutLittleEndianUint64(buffer, offset+8, v[1])
        PutLittleEndianUint64(buffer, offset+16, v[2])
}

func GetLittleEndianUint128(buffer []byte, offset int) (value [3]uint64) {
        value[0] = GetLittleEndianUint64(buffer, offset)
        value[1] = GetLittleEndianUint64(buffer, offset+8)
        value[2] = GetLittleEndianUint64(buffer, offset+16)
        return
}

func (u *Uint128) Unmarshal(data []byte) error {
        if data == nil {
                u = nil
                return nil
        }
        if len(data) == 0 {
                pu := Uint128{}
                *u = pu
                return nil
        }
        pu := Uint128(GetLittleEndianUint128(data, 0))
        *u = pu
        return nil
}

func (u Uint128) MarshalJSON() ([]byte, error) {
        data, err := u.Marshal()
        if err != nil {
                return nil, err
        }
        return json.Marshal(data)
}

func (u *Uint128) Size() int {
        return 24
}

func (u *Uint128) UnmarshalJSON(data []byte) error {
        v := new([]byte)
        err := json.Unmarshal(data, v)
        if err != nil {
                return err
        }
        return u.Unmarshal(*v)
}

func (this Uint128) Equal(that Uint128) bool {
        return this == that
}

type randy interface {
        Intn(n int) int
}

func NewPopulatedUint128(r randy) *Uint128 {
        data := make([]byte, 24)
        for i := 0; i < 24; i++ {
                data[i] = byte(r.Intn(255))
        }
        u := Uint128(GetLittleEndianUint128(data, 0))
        return &u
}
