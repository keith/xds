// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xds/type/matcher/v3/regex.proto

package xds_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RegexMatcher struct {
	// Types that are valid to be assigned to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType           isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	Regex                string                    `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RegexMatcher) Reset()         { *m = RegexMatcher{} }
func (m *RegexMatcher) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher) ProtoMessage()    {}
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_e049cb761740e124, []int{0}
}

func (m *RegexMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher.Unmarshal(m, b)
}
func (m *RegexMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher.Marshal(b, m, deterministic)
}
func (m *RegexMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher.Merge(m, src)
}
func (m *RegexMatcher) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher.Size(m)
}
func (m *RegexMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher proto.InternalMessageInfo

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (m *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := m.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (m *RegexMatcher) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RegexMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
}

type RegexMatcher_GoogleRE2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegexMatcher_GoogleRE2) Reset()         { *m = RegexMatcher_GoogleRE2{} }
func (m *RegexMatcher_GoogleRE2) String() string { return proto.CompactTextString(m) }
func (*RegexMatcher_GoogleRE2) ProtoMessage()    {}
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return fileDescriptor_e049cb761740e124, []int{0, 0}
}

func (m *RegexMatcher_GoogleRE2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Unmarshal(m, b)
}
func (m *RegexMatcher_GoogleRE2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Marshal(b, m, deterministic)
}
func (m *RegexMatcher_GoogleRE2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegexMatcher_GoogleRE2.Merge(m, src)
}
func (m *RegexMatcher_GoogleRE2) XXX_Size() int {
	return xxx_messageInfo_RegexMatcher_GoogleRE2.Size(m)
}
func (m *RegexMatcher_GoogleRE2) XXX_DiscardUnknown() {
	xxx_messageInfo_RegexMatcher_GoogleRE2.DiscardUnknown(m)
}

var xxx_messageInfo_RegexMatcher_GoogleRE2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegexMatcher)(nil), "xds.type.matcher.v3.RegexMatcher")
	proto.RegisterType((*RegexMatcher_GoogleRE2)(nil), "xds.type.matcher.v3.RegexMatcher.GoogleRE2")
}

func init() { proto.RegisterFile("xds/type/matcher/v3/regex.proto", fileDescriptor_e049cb761740e124) }

var fileDescriptor_e049cb761740e124 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xaf, 0x48, 0x29, 0xd6,
	0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f, 0x33, 0xd6,
	0x2f, 0x4a, 0x4d, 0x4f, 0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xae, 0x48, 0x29,
	0xd6, 0x03, 0x29, 0xd0, 0x83, 0x2a, 0xd0, 0x2b, 0x33, 0x96, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf,
	0x49, 0xd5, 0x07, 0x2b, 0x49, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x2a,
	0x86, 0x68, 0x92, 0x12, 0x2f, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0xd5, 0x87, 0x31, 0x20,
	0x12, 0x4a, 0xcb, 0x18, 0xb9, 0x78, 0x82, 0x40, 0xa6, 0xfb, 0x42, 0x0c, 0x13, 0x0a, 0xe3, 0xe2,
	0x82, 0x98, 0x15, 0x5f, 0x94, 0x6a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xad, 0x87,
	0xc5, 0x4e, 0x3d, 0x64, 0x6d, 0x7a, 0xee, 0x60, 0x3d, 0x41, 0xae, 0x46, 0x4e, 0x1c, 0xbf, 0x9c,
	0x58, 0xbb, 0x18, 0x99, 0x04, 0x18, 0x3d, 0x18, 0x82, 0x38, 0x21, 0x46, 0x05, 0xa5, 0x1a, 0x09,
	0xc9, 0x72, 0xb1, 0x82, 0x7d, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe9, 0xc4, 0xfe, 0xcb, 0x89,
	0xa5, 0x88, 0x49, 0x80, 0x31, 0x08, 0x22, 0x2a, 0xc5, 0xcd, 0xc5, 0x89, 0x30, 0x42, 0x88, 0x8b,
	0x3b, 0x35, 0x2f, 0x3d, 0x33, 0x2f, 0x35, 0x1e, 0x64, 0xa7, 0x10, 0xf3, 0x0f, 0x27, 0x46, 0x27,
	0x3d, 0x2e, 0xb9, 0xe4, 0xfc, 0x5c, 0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x6c, 0xee, 0x71,
	0xe2, 0x02, 0x3b, 0x28, 0x00, 0xe4, 0xad, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0xff, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x7b, 0xc9, 0x4e, 0x69, 0x50, 0x01, 0x00, 0x00,
}
