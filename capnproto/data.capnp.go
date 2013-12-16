package capnproto

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/jmckaskill/go-capnproto"
	"unsafe"
)

type TestST C.Struct

func NewTestST(s *C.Segment) TestST      { return TestST(s.NewStruct(8, 2)) }
func NewRootTestST(s *C.Segment) TestST  { return TestST(s.NewRootStruct(8, 2)) }
func ReadRootTestST(s *C.Segment) TestST { return TestST(s.Root(0).ToStruct()) }
func (s TestST) N() int32                { return int32(C.Struct(s).Get32(0)) }
func (s TestST) SetN(v int32)            { C.Struct(s).Set32(0, uint32(v)) }
func (s TestST) S() string               { return C.Struct(s).GetObject(0).ToText() }
func (s TestST) SetS(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s TestST) A() []byte               { return C.Struct(s).GetObject(1).ToData() }
func (s TestST) SetA(v []byte)           { C.Struct(s).SetObject(1, s.Segment.NewData(v)) }

type TestST_List C.PointerList

func NewTestSTList(s *C.Segment, sz int) TestST_List { return TestST_List(s.NewCompositeList(8, 2, sz)) }
func (s TestST_List) Len() int                       { return C.PointerList(s).Len() }
func (s TestST_List) At(i int) TestST                { return TestST(C.PointerList(s).At(i).ToStruct()) }
func (s TestST_List) ToArray() []TestST {
	return *(*[]TestST)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
