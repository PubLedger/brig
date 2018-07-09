// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp2 "github.com/sahib/brig/catfs/nodes/capnp"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Change describes a single change
type Change struct{ capnp.Struct }

// Change_TypeID is the unique identifier for the type Change.
const Change_TypeID = 0x9592300df48789af

func NewChange(s *capnp.Segment) (Change, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Change{st}, err
}

func NewRootChange(s *capnp.Segment) (Change, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Change{st}, err
}

func ReadRootChange(msg *capnp.Message) (Change, error) {
	root, err := msg.RootPtr()
	return Change{root.Struct()}, err
}

func (s Change) String() string {
	str, _ := text.Marshal(0x9592300df48789af, s.Struct)
	return str
}

func (s Change) Mask() uint64 {
	return s.Struct.Uint64(0)
}

func (s Change) SetMask(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Change) Head() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(0)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasHead() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Change) SetHead(v capnp2.Node) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewHead sets the head field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewHead() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) Next() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(1)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasNext() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Change) SetNext(v capnp2.Node) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewNext sets the next field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewNext() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) Curr() (capnp2.Node, error) {
	p, err := s.Struct.Ptr(2)
	return capnp2.Node{Struct: p.Struct()}, err
}

func (s Change) HasCurr() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Change) SetCurr(v capnp2.Node) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewCurr sets the curr field to a newly
// allocated capnp2.Node struct, preferring placement in s's segment.
func (s Change) NewCurr() (capnp2.Node, error) {
	ss, err := capnp2.NewNode(s.Struct.Segment())
	if err != nil {
		return capnp2.Node{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Change) ReferToPath() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s Change) HasReferToPath() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Change) ReferToPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s Change) SetReferToPath(v string) error {
	return s.Struct.SetText(3, v)
}

// Change_List is a list of Change.
type Change_List struct{ capnp.List }

// NewChange creates a new list of Change.
func NewChange_List(s *capnp.Segment, sz int32) (Change_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return Change_List{l}, err
}

func (s Change_List) At(i int) Change { return Change{s.List.Struct(i)} }

func (s Change_List) Set(i int, v Change) error { return s.List.SetStruct(i, v.Struct) }

func (s Change_List) String() string {
	str, _ := text.MarshalList(0x9592300df48789af, s.List)
	return str
}

// Change_Promise is a wrapper for a Change promised by a client call.
type Change_Promise struct{ *capnp.Pipeline }

func (p Change_Promise) Struct() (Change, error) {
	s, err := p.Pipeline.Struct()
	return Change{s}, err
}

func (p Change_Promise) Head() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Change_Promise) Next() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Change_Promise) Curr() capnp2.Node_Promise {
	return capnp2.Node_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

// Patch contains a single change
type Patch struct{ capnp.Struct }

// Patch_TypeID is the unique identifier for the type Patch.
const Patch_TypeID = 0x927c7336e3054805

func NewPatch(s *capnp.Segment) (Patch, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Patch{st}, err
}

func NewRootPatch(s *capnp.Segment) (Patch, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Patch{st}, err
}

func ReadRootPatch(msg *capnp.Message) (Patch, error) {
	root, err := msg.RootPtr()
	return Patch{root.Struct()}, err
}

func (s Patch) String() string {
	str, _ := text.Marshal(0x927c7336e3054805, s.Struct)
	return str
}

func (s Patch) FromIndex() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Patch) SetFromIndex(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Patch) CurrIndex() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s Patch) SetCurrIndex(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s Patch) Changes() (Change_List, error) {
	p, err := s.Struct.Ptr(0)
	return Change_List{List: p.List()}, err
}

func (s Patch) HasChanges() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Patch) SetChanges(v Change_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewChanges sets the changes field to a newly
// allocated Change_List, preferring placement in s's segment.
func (s Patch) NewChanges(n int32) (Change_List, error) {
	l, err := NewChange_List(s.Struct.Segment(), n)
	if err != nil {
		return Change_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Patch_List is a list of Patch.
type Patch_List struct{ capnp.List }

// NewPatch creates a new list of Patch.
func NewPatch_List(s *capnp.Segment, sz int32) (Patch_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Patch_List{l}, err
}

func (s Patch_List) At(i int) Patch { return Patch{s.List.Struct(i)} }

func (s Patch_List) Set(i int, v Patch) error { return s.List.SetStruct(i, v.Struct) }

func (s Patch_List) String() string {
	str, _ := text.MarshalList(0x927c7336e3054805, s.List)
	return str
}

// Patch_Promise is a wrapper for a Patch promised by a client call.
type Patch_Promise struct{ *capnp.Pipeline }

func (p Patch_Promise) Struct() (Patch, error) {
	s, err := p.Pipeline.Struct()
	return Patch{s}, err
}

const schema_b943b54bf1683782 = "x\xda|\xd0=k\x14]\x18\xc6\xf1\xeb\xba\xcf\xcc\xb3" +
	"\xe4a5\x19w+\x09\xec\xb4Z\xb8\x09\x01\x05\x11|" +
	"Y\x05E\xc59\xa2`%\x9c\xcc\xce\xee,&\xb3\xcb" +
	"\xcc(\x1bH\x88JD#Y\x08\x11\xc1\xceN\xb0J" +
	"e\x91\xc2\xd2\xaf\xe0\x17Hi%X\xa59r6\xbe" +
	"\x04\x0dv\x87?\x17\x1c~\xf7\xcc-^\x90Y\xbf\xe9" +
	"\x03\xfa\x9c\xff\x9f\xf5\xaf\xfa\xbb\xa7\x8b\xe5-\xe8i\x8a" +
	"}z&\xfdz\xfdCk\x07>+\xc0\xdc\x8a\x1cc" +
	"m$\x95\xdaH\x1a\xb5\x1d\xd9\x06\xed\xf6\xfa\xf3oG" +
	"f\xb6^\xbb=\x0f\xec=\xb7\xef\xa9\xe3\xac-\xa9J" +
	"mI5\xe6\xde\xab\x06\xb1icSv\x8a\xe6\xa3X" +
	"\x15\xcd\xd8\x0c\xb2As`\xca8=5~\x9f\x8dL" +
	"\x193\x8dH\xedQ\xec\xfdWo\xf5\xc7\xcf/?A" +
	"{\xc2\x8b\xd3d\x15\x08\xb8g\xdd*\x0d\xe3\xbed\xa5" +
	"\xe9eEh\xc2\xa2\x97u\x17\x92\xf0|\x9c\x9a\xac\x9b" +
	"\x00\xba\xaa<\xc0#\x10\\\xb9\x0d\xe8\xcb\x8a:\x12\x06" +
	"d\x9d.\xdet\xf1\x86\xa2\xbe'\xa4\xd4)@p\xf7" +
	"\x12\xa0#E\xbd \xb4\x9d\xbc\xbfx-k'\xe0\x90" +
	">\x84>h\xe3\x87y\xfeG[\xdd\xff\xb0\xe0Q0" +
	"R\xe4\xd4\xef{\x80.\xfe\x9b\xdbJM\xa6\xba\xc9\xe1" +
	"\xdep\xec\x9d\xe5\xff\xb4\xad\xf1/a[%E\x9c\xf7" +
	"\xe6\x93\x03\xe4\x1fb\xea\xfa/\xf1\xcaI@\x0f\x15\xf5" +
	"\x9a\xf0'\xf8\x89k\xcb\x8a\xfa\x850\x10\xee\x8b\x9f\xb9" +
	"\xf8XQo\x08\x03%u* XwqMQo" +
	"\x0a\x03O\xd5\xe9\x01\xc1h\x1e\xd0\x1b\x8a\xfa\x8dpr" +
	"\xd1\x14\x0f8\x01\xe1\x048\x99&\xa6\xcd)\xbb\xbb\xd7" +
	"\x19\xac~9\xf1\xce\xb1\xa7\xc0\xc9,\x19\x96\x87dw" +
	"\xc2\xbf\xb3\xcd\x93N\x92\xdf\xe9G\xa8\x982e\x15\xc2" +
	"*\xf8=\x00\x00\xff\xff\xf5$\xa6\x92"

func init() {
	schemas.Register(schema_b943b54bf1683782,
		0x927c7336e3054805,
		0x9592300df48789af)
}