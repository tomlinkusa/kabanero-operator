// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/advertising_channel_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing the various advertising channel types.
type AdvertisingChannelTypeEnum_AdvertisingChannelType int32

const (
	// Not specified.
	AdvertisingChannelTypeEnum_UNSPECIFIED AdvertisingChannelTypeEnum_AdvertisingChannelType = 0
	// Used for return value only. Represents value unknown in this version.
	AdvertisingChannelTypeEnum_UNKNOWN AdvertisingChannelTypeEnum_AdvertisingChannelType = 1
	// Search Network. Includes display bundled, and Search+ campaigns.
	AdvertisingChannelTypeEnum_SEARCH AdvertisingChannelTypeEnum_AdvertisingChannelType = 2
	// Google Display Network only.
	AdvertisingChannelTypeEnum_DISPLAY AdvertisingChannelTypeEnum_AdvertisingChannelType = 3
	// Shopping campaigns serve on the shopping property
	// and on google.com search results.
	AdvertisingChannelTypeEnum_SHOPPING AdvertisingChannelTypeEnum_AdvertisingChannelType = 4
	// Hotel Ads campaigns.
	AdvertisingChannelTypeEnum_HOTEL AdvertisingChannelTypeEnum_AdvertisingChannelType = 5
	// Video campaigns.
	AdvertisingChannelTypeEnum_VIDEO AdvertisingChannelTypeEnum_AdvertisingChannelType = 6
)

var AdvertisingChannelTypeEnum_AdvertisingChannelType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SEARCH",
	3: "DISPLAY",
	4: "SHOPPING",
	5: "HOTEL",
	6: "VIDEO",
}
var AdvertisingChannelTypeEnum_AdvertisingChannelType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SEARCH":      2,
	"DISPLAY":     3,
	"SHOPPING":    4,
	"HOTEL":       5,
	"VIDEO":       6,
}

func (x AdvertisingChannelTypeEnum_AdvertisingChannelType) String() string {
	return proto.EnumName(AdvertisingChannelTypeEnum_AdvertisingChannelType_name, int32(x))
}
func (AdvertisingChannelTypeEnum_AdvertisingChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_type_ba3b48a2ec07c28a, []int{0, 0}
}

// The channel type a campaign may target to serve on.
type AdvertisingChannelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdvertisingChannelTypeEnum) Reset()         { *m = AdvertisingChannelTypeEnum{} }
func (m *AdvertisingChannelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdvertisingChannelTypeEnum) ProtoMessage()    {}
func (*AdvertisingChannelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertising_channel_type_ba3b48a2ec07c28a, []int{0}
}
func (m *AdvertisingChannelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Unmarshal(m, b)
}
func (m *AdvertisingChannelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AdvertisingChannelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertisingChannelTypeEnum.Merge(dst, src)
}
func (m *AdvertisingChannelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdvertisingChannelTypeEnum.Size(m)
}
func (m *AdvertisingChannelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertisingChannelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertisingChannelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdvertisingChannelTypeEnum)(nil), "google.ads.googleads.v0.enums.AdvertisingChannelTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdvertisingChannelTypeEnum_AdvertisingChannelType", AdvertisingChannelTypeEnum_AdvertisingChannelType_name, AdvertisingChannelTypeEnum_AdvertisingChannelType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/advertising_channel_type.proto", fileDescriptor_advertising_channel_type_ba3b48a2ec07c28a)
}

var fileDescriptor_advertising_channel_type_ba3b48a2ec07c28a = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xc2, 0x30,
	0x1c, 0xc6, 0xdd, 0x10, 0xd4, 0x62, 0x62, 0xb3, 0x83, 0x07, 0x0d, 0x07, 0x78, 0x80, 0x6e, 0x89,
	0xb7, 0xea, 0xa5, 0xc0, 0x84, 0x45, 0xb2, 0x2d, 0x0e, 0x66, 0x34, 0x4b, 0xc8, 0x64, 0x4d, 0x25,
	0x81, 0x76, 0xa1, 0x40, 0xc2, 0x63, 0xf8, 0x0a, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0x93, 0x8f, 0x60,
	0xda, 0x0a, 0x5e, 0xd0, 0xcb, 0xf2, 0x65, 0xbf, 0xff, 0xf7, 0xe5, 0xeb, 0x07, 0x6e, 0x98, 0x10,
	0x6c, 0x46, 0xdd, 0xbc, 0x90, 0xae, 0x91, 0x4a, 0xad, 0x3d, 0x97, 0xf2, 0xd5, 0x5c, 0xba, 0x79,
	0xb1, 0xa6, 0x8b, 0xe5, 0x54, 0x4e, 0x39, 0x1b, 0x4f, 0x5e, 0x72, 0xce, 0xe9, 0x6c, 0xbc, 0xdc,
	0x94, 0x14, 0x95, 0x0b, 0xb1, 0x14, 0x4e, 0xc3, 0x58, 0x50, 0x5e, 0x48, 0xb4, 0x73, 0xa3, 0xb5,
	0x87, 0xb4, 0xbb, 0xf5, 0x6a, 0x81, 0x0b, 0xf2, 0x9b, 0xd0, 0x31, 0x01, 0xc3, 0x4d, 0x49, 0x7d,
	0xbe, 0x9a, 0xb7, 0x24, 0x38, 0xdf, 0x4f, 0x9d, 0x33, 0x50, 0x1f, 0x85, 0x49, 0xec, 0x77, 0x82,
	0xdb, 0xc0, 0xef, 0xc2, 0x03, 0xa7, 0x0e, 0x8e, 0x46, 0xe1, 0x5d, 0x18, 0x3d, 0x84, 0xd0, 0x72,
	0x00, 0xa8, 0x25, 0x3e, 0xb9, 0xef, 0xf4, 0xa1, 0xad, 0x40, 0x37, 0x48, 0xe2, 0x01, 0x79, 0x84,
	0x15, 0xe7, 0x14, 0x1c, 0x27, 0xfd, 0x28, 0x8e, 0x83, 0xb0, 0x07, 0x0f, 0x9d, 0x13, 0x50, 0xed,
	0x47, 0x43, 0x7f, 0x00, 0xab, 0x4a, 0xa6, 0x41, 0xd7, 0x8f, 0x60, 0xad, 0xfd, 0x65, 0x81, 0xe6,
	0x44, 0xcc, 0xd1, 0xbf, 0xcd, 0xdb, 0x97, 0xfb, 0x8b, 0xc5, 0xea, 0xd5, 0xb1, 0xf5, 0xd4, 0xfe,
	0x71, 0x33, 0x31, 0xcb, 0x39, 0x43, 0x62, 0xc1, 0x5c, 0x46, 0xb9, 0xde, 0x64, 0xbb, 0x62, 0x39,
	0x95, 0x7f, 0x8c, 0x7a, 0xad, 0xbf, 0x6f, 0x76, 0xa5, 0x47, 0xc8, 0xbb, 0xdd, 0xe8, 0x99, 0x28,
	0x52, 0x48, 0x64, 0xa4, 0x52, 0xa9, 0x87, 0xd4, 0x44, 0xf2, 0x63, 0xcb, 0x33, 0x52, 0xc8, 0x6c,
	0xc7, 0xb3, 0xd4, 0xcb, 0x34, 0xff, 0xb4, 0x9b, 0xe6, 0x27, 0xc6, 0xa4, 0x90, 0x18, 0xef, 0x2e,
	0x30, 0x4e, 0x3d, 0x8c, 0xf5, 0xcd, 0x73, 0x4d, 0x17, 0xbb, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x56, 0x7e, 0x57, 0xec, 0x01, 0x00, 0x00,
}