// Code generated by protoc-gen-gogo.
// source: wrapper.proto
// DO NOT EDIT!

package tests_kitchensink

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"
import google_protobuf3 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Wrappers struct {
	None        *google_protobuf3.DoubleValue `protobuf:"bytes,1,opt,name=none" json:"none,omitempty"`
	DoubleValue *google_protobuf3.DoubleValue `protobuf:"bytes,2,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	FloatValue  *google_protobuf3.FloatValue  `protobuf:"bytes,3,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	Int64Value  *google_protobuf3.Int64Value  `protobuf:"bytes,4,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Uint64Value *google_protobuf3.UInt64Value `protobuf:"bytes,5,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Int32Value  *google_protobuf3.Int32Value  `protobuf:"bytes,6,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Uint32Value *google_protobuf3.UInt32Value `protobuf:"bytes,7,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	BoolValue   *google_protobuf3.BoolValue   `protobuf:"bytes,8,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	StringValue *google_protobuf3.StringValue `protobuf:"bytes,9,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	BytesValue  *google_protobuf3.BytesValue  `protobuf:"bytes,10,opt,name=bytes_value,json=bytesValue" json:"bytes_value,omitempty"`
}

func (m *Wrappers) Reset()                    { *m = Wrappers{} }
func (m *Wrappers) String() string            { return proto.CompactTextString(m) }
func (*Wrappers) ProtoMessage()               {}
func (*Wrappers) Descriptor() ([]byte, []int) { return fileDescriptorWrapper, []int{0} }

func (m *Wrappers) GetNone() *google_protobuf3.DoubleValue {
	if m != nil {
		return m.None
	}
	return nil
}

func (m *Wrappers) GetDoubleValue() *google_protobuf3.DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *Wrappers) GetFloatValue() *google_protobuf3.FloatValue {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *Wrappers) GetInt64Value() *google_protobuf3.Int64Value {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

func (m *Wrappers) GetUint64Value() *google_protobuf3.UInt64Value {
	if m != nil {
		return m.Uint64Value
	}
	return nil
}

func (m *Wrappers) GetInt32Value() *google_protobuf3.Int32Value {
	if m != nil {
		return m.Int32Value
	}
	return nil
}

func (m *Wrappers) GetUint32Value() *google_protobuf3.UInt32Value {
	if m != nil {
		return m.Uint32Value
	}
	return nil
}

func (m *Wrappers) GetBoolValue() *google_protobuf3.BoolValue {
	if m != nil {
		return m.BoolValue
	}
	return nil
}

func (m *Wrappers) GetStringValue() *google_protobuf3.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *Wrappers) GetBytesValue() *google_protobuf3.BytesValue {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Wrappers)(nil), "tests.kitchensink.Wrappers")
}

func init() { proto.RegisterFile("wrapper.proto", fileDescriptorWrapper) }

var fileDescriptorWrapper = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0xcf, 0x6a, 0xdb, 0x30,
	0x1c, 0x07, 0x70, 0xfc, 0x37, 0xb6, 0xec, 0xc0, 0x26, 0x06, 0x0b, 0xde, 0x1f, 0xb2, 0x9c, 0xc6,
	0x0e, 0x8e, 0x71, 0xc6, 0x08, 0xbb, 0x6c, 0x84, 0xb0, 0xb0, 0x8c, 0x0d, 0xa6, 0xb1, 0xed, 0x38,
	0xec, 0x45, 0x4e, 0xdd, 0x18, 0x2b, 0xd8, 0x72, 0x4a, 0x73, 0xef, 0x03, 0xf4, 0x45, 0xf2, 0x00,
	0x3d, 0xf5, 0x55, 0x7a, 0x2c, 0x39, 0xf6, 0x05, 0x8a, 0x64, 0xd9, 0x31, 0x04, 0xb7, 0xb7, 0xe4,
	0xf7, 0xfb, 0xea, 0xc3, 0x57, 0x16, 0xe8, 0x9e, 0x65, 0xc1, 0x7a, 0x8d, 0x33, 0x77, 0x9d, 0x11,
	0x4a, 0xe0, 0x53, 0x8a, 0x73, 0x9a, 0xbb, 0xab, 0x98, 0xfe, 0x3f, 0xc1, 0x69, 0x1e, 0xa7, 0x2b,
	0xe7, 0xf9, 0x26, 0x48, 0xe2, 0x45, 0x40, 0xf1, 0xb0, 0xfa, 0x51, 0x66, 0x9d, 0xd7, 0x4b, 0x42,
	0x96, 0x09, 0x1e, 0xf2, 0x7f, 0x61, 0x11, 0x0d, 0x05, 0x95, 0x97, 0xfb, 0xc1, 0x5e, 0x03, 0xc6,
	0x5f, 0x31, 0x82, 0x1e, 0x50, 0x53, 0x92, 0xe2, 0x9e, 0xd4, 0x97, 0xde, 0x5a, 0xfe, 0x4b, 0xb7,
	0x3c, 0xeb, 0x56, 0x67, 0xdd, 0x29, 0x29, 0xc2, 0x04, 0xff, 0x09, 0x92, 0x02, 0x23, 0x9e, 0x84,
	0x3f, 0x81, 0xbd, 0xe0, 0xc3, 0x7f, 0x1b, 0x36, 0xed, 0xc9, 0x8f, 0x9f, 0x9c, 0x3c, 0xb9, 0xba,
	0xbd, 0x56, 0x2c, 0x68, 0xbe, 0xd9, 0xcd, 0x6e, 0xb6, 0xcf, 0x76, 0x77, 0x9f, 0x90, 0xb5, 0x38,
	0xac, 0xe1, 0x1c, 0x58, 0x51, 0x42, 0x02, 0x2a, 0x44, 0x85, 0x8b, 0x2f, 0x8e, 0xc4, 0x2f, 0x2c,
	0x53, 0x82, 0x36, 0x03, 0x3b, 0x40, 0x7b, 0x75, 0xb1, 0xbf, 0xfc, 0x8c, 0x40, 0x54, 0x6f, 0xe0,
	0x0c, 0x58, 0x71, 0x4a, 0x3f, 0xbc, 0x17, 0x96, 0xda, 0x62, 0x7d, 0x65, 0x99, 0xd2, 0x32, 0x99,
	0xa5, 0x0e, 0x64, 0x43, 0x43, 0x20, 0xae, 0xc7, 0xf0, 0x07, 0xb0, 0x8b, 0xa6, 0xa4, 0xb5, 0xdc,
	0xf3, 0x77, 0x83, 0xea, 0x32, 0xca, 0xf0, 0x75, 0x4f, 0xf2, 0x64, 0x4f, 0x41, 0x56, 0xd1, 0xf0,
	0xbe, 0xf1, 0x62, 0x23, 0x5f, 0x70, 0x7a, 0x7b, 0xb1, 0x91, 0xdf, 0xd4, 0x1c, 0x7d, 0xac, 0x8e,
	0xb5, 0xb1, 0xce, 0xcb, 0x89, 0x15, 0x9c, 0x97, 0xe5, 0x6a, 0xad, 0xf3, 0x40, 0xb9, 0x8a, 0x2b,
	0xef, 0xf9, 0x4e, 0xee, 0x8b, 0x62, 0x95, 0x35, 0x05, 0x20, 0x24, 0x24, 0x11, 0x92, 0xc1, 0x25,
	0xe7, 0x48, 0x9a, 0x10, 0x92, 0x34, 0x9d, 0x53, 0xd9, 0x90, 0x90, 0x19, 0x56, 0x53, 0xf8, 0x1d,
	0xd8, 0x39, 0xcd, 0xe2, 0x74, 0x29, 0x1c, 0xb3, 0xa5, 0xd1, 0x2f, 0x1e, 0x6a, 0xbe, 0x62, 0xa6,
	0x7d, 0x54, 0x22, 0x42, 0x90, 0x95, 0x1f, 0x56, 0xec, 0x19, 0xc3, 0x73, 0x8a, 0x73, 0xa1, 0x81,
	0x96, 0xaf, 0x35, 0x61, 0x99, 0x66, 0xad, 0xad, 0xdc, 0x33, 0x10, 0x08, 0xeb, 0x71, 0xa8, 0xf3,
	0xec, 0xe8, 0x3e, 0x00, 0x00, 0xff, 0xff, 0x23, 0x03, 0xd6, 0x07, 0x51, 0x03, 0x00, 0x00,
}
