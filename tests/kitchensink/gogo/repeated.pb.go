// Code generated by protoc-gen-gogo.
// source: repeated.proto
// DO NOT EDIT!

package tests_kitchensink

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Repeated struct {
	None        []int64           `protobuf:"varint,1,rep,packed,name=none" json:"none,omitempty"`
	MinItems    []int64           `protobuf:"varint,2,rep,packed,name=min_items,json=minItems" json:"min_items,omitempty"`
	MaxItems    []string          `protobuf:"bytes,3,rep,name=max_items,json=maxItems" json:"max_items,omitempty"`
	MinMaxItems [][]byte          `protobuf:"bytes,4,rep,name=min_max_items,json=minMaxItems" json:"min_max_items,omitempty"`
	EqItems     []*Repeated_Embed `protobuf:"bytes,5,rep,name=eq_items,json=eqItems" json:"eq_items,omitempty"`
	Unique      []int64           `protobuf:"varint,6,rep,packed,name=unique" json:"unique,omitempty"`
	UniqueBytes [][]byte          `protobuf:"bytes,7,rep,name=unique_bytes,json=uniqueBytes" json:"unique_bytes,omitempty"`
	ItemRules   []float64         `protobuf:"fixed64,8,rep,packed,name=item_rules,json=itemRules" json:"item_rules,omitempty"`
}

func (m *Repeated) Reset()                    { *m = Repeated{} }
func (m *Repeated) String() string            { return proto.CompactTextString(m) }
func (*Repeated) ProtoMessage()               {}
func (*Repeated) Descriptor() ([]byte, []int) { return fileDescriptorRepeated, []int{0} }

func (m *Repeated) GetNone() []int64 {
	if m != nil {
		return m.None
	}
	return nil
}

func (m *Repeated) GetMinItems() []int64 {
	if m != nil {
		return m.MinItems
	}
	return nil
}

func (m *Repeated) GetMaxItems() []string {
	if m != nil {
		return m.MaxItems
	}
	return nil
}

func (m *Repeated) GetMinMaxItems() [][]byte {
	if m != nil {
		return m.MinMaxItems
	}
	return nil
}

func (m *Repeated) GetEqItems() []*Repeated_Embed {
	if m != nil {
		return m.EqItems
	}
	return nil
}

func (m *Repeated) GetUnique() []int64 {
	if m != nil {
		return m.Unique
	}
	return nil
}

func (m *Repeated) GetUniqueBytes() [][]byte {
	if m != nil {
		return m.UniqueBytes
	}
	return nil
}

func (m *Repeated) GetItemRules() []float64 {
	if m != nil {
		return m.ItemRules
	}
	return nil
}

type Repeated_Embed struct {
}

func (m *Repeated_Embed) Reset()                    { *m = Repeated_Embed{} }
func (m *Repeated_Embed) String() string            { return proto.CompactTextString(m) }
func (*Repeated_Embed) ProtoMessage()               {}
func (*Repeated_Embed) Descriptor() ([]byte, []int) { return fileDescriptorRepeated, []int{0, 0} }

func init() {
	proto.RegisterType((*Repeated)(nil), "tests.kitchensink.Repeated")
	proto.RegisterType((*Repeated_Embed)(nil), "tests.kitchensink.Repeated.Embed")
}

func init() { proto.RegisterFile("repeated.proto", fileDescriptorRepeated) }

var fileDescriptorRepeated = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd1, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x07, 0x70, 0xb7, 0x9b, 0x34, 0x9b, 0x49, 0x2b, 0xe9, 0x82, 0x18, 0x7a, 0x8a, 0xbd, 0x98,
	0x8b, 0x51, 0xd4, 0x07, 0x90, 0x80, 0x07, 0x05, 0x2f, 0xfb, 0x02, 0x25, 0xb5, 0x03, 0x2e, 0xed,
	0x6e, 0x3f, 0x76, 0x2b, 0xf5, 0xe6, 0x33, 0xf4, 0x71, 0x3c, 0xf9, 0x3a, 0xbe, 0x85, 0xec, 0x26,
	0xa5, 0x15, 0xf7, 0x34, 0xf0, 0xff, 0xcd, 0x30, 0xc3, 0xc2, 0xe9, 0x1a, 0x97, 0x58, 0x5b, 0x9c,
	0x96, 0xcb, 0xf5, 0xc2, 0x2e, 0xf8, 0xc0, 0xa2, 0xb1, 0xa6, 0x9c, 0x49, 0xfb, 0xfa, 0x86, 0xda,
	0x48, 0x3d, 0x1b, 0x9e, 0xbf, 0xd7, 0x73, 0x39, 0xad, 0x2d, 0x5e, 0xef, 0x8b, 0xc6, 0x8e, 0x3e,
	0x29, 0x30, 0xd1, 0xb6, 0x73, 0x0e, 0x81, 0x5e, 0x68, 0xcc, 0x48, 0x4e, 0x0b, 0x2a, 0x7c, 0xcd,
	0x2f, 0x21, 0x56, 0x52, 0x8f, 0xa5, 0x45, 0x65, 0xb2, 0x8e, 0x0b, 0x2a, 0xf8, 0xfa, 0xf9, 0xa6,
	0xe1, 0x8e, 0x74, 0x18, 0x11, 0x4c, 0x49, 0xfd, 0xe4, 0x32, 0x0f, 0xeb, 0x6d, 0x0b, 0x69, 0x4e,
	0x8b, 0xf8, 0x00, 0xd3, 0x8e, 0x60, 0xaa, 0xde, 0x36, 0xf0, 0x06, 0xfa, 0x6e, 0xe2, 0x01, 0x07,
	0x39, 0x2d, 0x7a, 0x55, 0xcf, 0xe1, 0x68, 0x47, 0x02, 0x46, 0xd3, 0x40, 0x24, 0x4a, 0xea, 0x97,
	0x7d, 0xc7, 0x33, 0x30, 0x5c, 0xb5, 0x38, 0xcc, 0x69, 0x91, 0xdc, 0x5e, 0x94, 0xff, 0x6e, 0x2c,
	0xf7, 0x67, 0x94, 0x8f, 0x6a, 0x82, 0xd3, 0xa3, 0x79, 0x61, 0x1a, 0x8a, 0x08, 0x57, 0xcd, 0xac,
	0x11, 0x74, 0x37, 0x5a, 0xae, 0x36, 0x98, 0x75, 0xff, 0x1e, 0x93, 0x11, 0xd1, 0x26, 0xfc, 0x0a,
	0x7a, 0x4d, 0x35, 0x9e, 0x7c, 0x58, 0x34, 0x59, 0xe4, 0x17, 0x3c, 0x96, 0x49, 0x93, 0x57, 0x2e,
	0xe6, 0xf7, 0x00, 0x6e, 0xb7, 0xf1, 0x7a, 0x33, 0x47, 0x93, 0xb1, 0x9c, 0x16, 0xa4, 0x3a, 0x73,
	0x38, 0xdd, 0x91, 0xfe, 0x28, 0xe1, 0xf1, 0xe0, 0xc4, 0x3f, 0xf6, 0x20, 0x62, 0x07, 0x85, 0x73,
	0xc3, 0x08, 0x42, 0xbf, 0xe8, 0xa4, 0xeb, 0x7f, 0xe2, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x66,
	0xa4, 0x79, 0xdc, 0xc7, 0x01, 0x00, 0x00,
}
