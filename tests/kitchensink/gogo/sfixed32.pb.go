// Code generated by protoc-gen-gogo.
// source: sfixed32.proto
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

type SFixed32 struct {
	None      int32 `protobuf:"fixed32,1,opt,name=none,proto3" json:"none,omitempty"`
	Lt        int32 `protobuf:"fixed32,2,opt,name=lt,proto3" json:"lt,omitempty"`
	Lte       int32 `protobuf:"fixed32,3,opt,name=lte,proto3" json:"lte,omitempty"`
	Gt        int32 `protobuf:"fixed32,4,opt,name=gt,proto3" json:"gt,omitempty"`
	Gte       int32 `protobuf:"fixed32,5,opt,name=gte,proto3" json:"gte,omitempty"`
	LtGt      int32 `protobuf:"fixed32,6,opt,name=lt_gt,json=ltGt,proto3" json:"lt_gt,omitempty"`
	LtGte     int32 `protobuf:"fixed32,7,opt,name=lt_gte,json=ltGte,proto3" json:"lt_gte,omitempty"`
	LteGt     int32 `protobuf:"fixed32,8,opt,name=lte_gt,json=lteGt,proto3" json:"lte_gt,omitempty"`
	LteGte    int32 `protobuf:"fixed32,9,opt,name=lte_gte,json=lteGte,proto3" json:"lte_gte,omitempty"`
	LtGtInv   int32 `protobuf:"fixed32,10,opt,name=lt_gt_inv,json=ltGtInv,proto3" json:"lt_gt_inv,omitempty"`
	LtGteInv  int32 `protobuf:"fixed32,11,opt,name=lt_gte_inv,json=ltGteInv,proto3" json:"lt_gte_inv,omitempty"`
	LteGtInv  int32 `protobuf:"fixed32,12,opt,name=lte_gt_inv,json=lteGtInv,proto3" json:"lte_gt_inv,omitempty"`
	LteGteInv int32 `protobuf:"fixed32,13,opt,name=lte_gte_inv,json=lteGteInv,proto3" json:"lte_gte_inv,omitempty"`
	In        int32 `protobuf:"fixed32,14,opt,name=in,proto3" json:"in,omitempty"`
	NotIn     int32 `protobuf:"fixed32,15,opt,name=not_in,json=notIn,proto3" json:"not_in,omitempty"`
	Const     int32 `protobuf:"fixed32,16,opt,name=const,proto3" json:"const,omitempty"`
}

func (m *SFixed32) Reset()                    { *m = SFixed32{} }
func (m *SFixed32) String() string            { return proto.CompactTextString(m) }
func (*SFixed32) ProtoMessage()               {}
func (*SFixed32) Descriptor() ([]byte, []int) { return fileDescriptorSfixed32, []int{0} }

func (m *SFixed32) GetNone() int32 {
	if m != nil {
		return m.None
	}
	return 0
}

func (m *SFixed32) GetLt() int32 {
	if m != nil {
		return m.Lt
	}
	return 0
}

func (m *SFixed32) GetLte() int32 {
	if m != nil {
		return m.Lte
	}
	return 0
}

func (m *SFixed32) GetGt() int32 {
	if m != nil {
		return m.Gt
	}
	return 0
}

func (m *SFixed32) GetGte() int32 {
	if m != nil {
		return m.Gte
	}
	return 0
}

func (m *SFixed32) GetLtGt() int32 {
	if m != nil {
		return m.LtGt
	}
	return 0
}

func (m *SFixed32) GetLtGte() int32 {
	if m != nil {
		return m.LtGte
	}
	return 0
}

func (m *SFixed32) GetLteGt() int32 {
	if m != nil {
		return m.LteGt
	}
	return 0
}

func (m *SFixed32) GetLteGte() int32 {
	if m != nil {
		return m.LteGte
	}
	return 0
}

func (m *SFixed32) GetLtGtInv() int32 {
	if m != nil {
		return m.LtGtInv
	}
	return 0
}

func (m *SFixed32) GetLtGteInv() int32 {
	if m != nil {
		return m.LtGteInv
	}
	return 0
}

func (m *SFixed32) GetLteGtInv() int32 {
	if m != nil {
		return m.LteGtInv
	}
	return 0
}

func (m *SFixed32) GetLteGteInv() int32 {
	if m != nil {
		return m.LteGteInv
	}
	return 0
}

func (m *SFixed32) GetIn() int32 {
	if m != nil {
		return m.In
	}
	return 0
}

func (m *SFixed32) GetNotIn() int32 {
	if m != nil {
		return m.NotIn
	}
	return 0
}

func (m *SFixed32) GetConst() int32 {
	if m != nil {
		return m.Const
	}
	return 0
}

func init() {
	proto.RegisterType((*SFixed32)(nil), "tests.kitchensink.SFixed32")
}

func init() { proto.RegisterFile("sfixed32.proto", fileDescriptorSfixed32) }

var fileDescriptorSfixed32 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x4f, 0x2a, 0x31,
	0x14, 0x86, 0xcb, 0x30, 0x5f, 0x1c, 0xbe, 0x2e, 0xcd, 0x25, 0xb7, 0xd7, 0x08, 0x51, 0x0c, 0x86,
	0x98, 0x0c, 0x44, 0x08, 0x4b, 0x36, 0x2e, 0x24, 0x6c, 0x71, 0xc7, 0x86, 0x20, 0xd4, 0x71, 0xc2,
	0xa4, 0x63, 0x9c, 0x13, 0xe2, 0x6f, 0x73, 0xe5, 0xdf, 0xf1, 0x17, 0xb8, 0x35, 0x3d, 0x15, 0x49,
	0x33, 0xbb, 0x86, 0xf7, 0x79, 0xce, 0x79, 0x69, 0x07, 0x1a, 0xf9, 0x53, 0xf2, 0x26, 0x77, 0x93,
	0xf1, 0xf0, 0xe5, 0x35, 0xc3, 0x8c, 0xb7, 0x50, 0xe6, 0x98, 0x0f, 0xf7, 0x09, 0x6e, 0x9f, 0xa5,
	0xca, 0x13, 0xb5, 0x3f, 0xfb, 0x77, 0xd8, 0xa4, 0xc9, 0x6e, 0x83, 0x72, 0x74, 0x3c, 0x18, 0xb6,
	0xf7, 0xe5, 0x42, 0xf8, 0x70, 0x6f, 0x74, 0xce, 0xc1, 0x55, 0x99, 0x92, 0xa2, 0x74, 0x51, 0x1a,
	0x34, 0x97, 0x74, 0xe6, 0xe7, 0xe0, 0xa4, 0x28, 0x1c, 0xfd, 0xcb, 0x5d, 0xed, 0xfd, 0xf3, 0xa3,
	0x1c, 0xac, 0xbc, 0xb6, 0xc7, 0x18, 0x5b, 0x3a, 0x29, 0xf2, 0x2e, 0x94, 0x53, 0x94, 0xa2, 0x6c,
	0xc5, 0x1d, 0x8a, 0x75, 0xa0, 0xed, 0x18, 0x85, 0x6b, 0xc5, 0x7d, 0x63, 0xc7, 0x64, 0xc7, 0x28,
	0x85, 0x67, 0xc5, 0x91, 0xb1, 0x63, 0x94, 0xfc, 0x1a, 0xbc, 0x14, 0xd7, 0x31, 0x0a, 0x9f, 0x88,
	0x96, 0x26, 0x6a, 0x2b, 0x68, 0x37, 0x19, 0x63, 0x7d, 0xd0, 0x98, 0x9b, 0xe2, 0x1c, 0xf9, 0x00,
	0x7c, 0xe2, 0xa4, 0x08, 0x8a, 0x60, 0x44, 0xa0, 0xa7, 0x41, 0x69, 0x48, 0xa9, 0x47, 0x86, 0x16,
	0xd9, 0x39, 0x8d, 0xf4, 0x52, 0x94, 0x73, 0xe4, 0x37, 0x10, 0x18, 0x52, 0x8a, 0x4a, 0x11, 0x35,
	0x43, 0x7d, 0x42, 0x25, 0x8f, 0xa0, 0x42, 0xfb, 0xd7, 0x89, 0x3a, 0x08, 0xb0, 0x2b, 0xfc, 0xd5,
	0x83, 0xff, 0x6b, 0x3a, 0xd0, 0x15, 0x16, 0xea, 0xc0, 0x47, 0x00, 0xa6, 0x2e, 0xf1, 0xd5, 0x22,
	0x1f, 0x11, 0x1f, 0x52, 0xe5, 0x5f, 0x41, 0x1e, 0x17, 0xd4, 0xec, 0x3a, 0xa7, 0x05, 0x21, 0xd5,
	0xd1, 0xc2, 0x2d, 0x54, 0x7f, 0xca, 0x93, 0x51, 0x2f, 0x1a, 0x66, 0x45, 0xc5, 0xfc, 0x01, 0xad,
	0x5c, 0x82, 0x93, 0x28, 0xd1, 0xb0, 0xc8, 0x69, 0x97, 0x31, 0x36, 0xbd, 0xa2, 0xe7, 0x4a, 0x94,
	0xbe, 0x3c, 0x95, 0xe9, 0x0e, 0xa2, 0x69, 0x61, 0xb3, 0x01, 0x63, 0x6c, 0x16, 0xd1, 0xe5, 0xa9,
	0x0c, 0x17, 0x8a, 0xf7, 0xc0, 0xdb, 0x66, 0x2a, 0x47, 0xf1, 0xc7, 0x7a, 0xda, 0xfa, 0x98, 0x18,
	0x8a, 0x1e, 0x7d, 0xfa, 0x00, 0x27, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x43, 0x16, 0xf9,
	0xbe, 0x02, 0x00, 0x00,
}
