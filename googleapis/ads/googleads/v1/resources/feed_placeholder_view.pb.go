// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/feed_placeholder_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A feed placeholder view.
type FeedPlaceholderView struct {
	// Output only. The resource name of the feed placeholder view.
	// Feed placeholder view resource names have the form:
	//
	// `customers/{customer_id}/feedPlaceholderViews/{placeholder_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The placeholder type of the feed placeholder view.
	PlaceholderType      enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,2,opt,name=placeholder_type,json=placeholderType,proto3,enum=google.ads.googleads.v1.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *FeedPlaceholderView) Reset()         { *m = FeedPlaceholderView{} }
func (m *FeedPlaceholderView) String() string { return proto.CompactTextString(m) }
func (*FeedPlaceholderView) ProtoMessage()    {}
func (*FeedPlaceholderView) Descriptor() ([]byte, []int) {
	return fileDescriptor_84882bb262f263ab, []int{0}
}

func (m *FeedPlaceholderView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedPlaceholderView.Unmarshal(m, b)
}
func (m *FeedPlaceholderView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedPlaceholderView.Marshal(b, m, deterministic)
}
func (m *FeedPlaceholderView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedPlaceholderView.Merge(m, src)
}
func (m *FeedPlaceholderView) XXX_Size() int {
	return xxx_messageInfo_FeedPlaceholderView.Size(m)
}
func (m *FeedPlaceholderView) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedPlaceholderView.DiscardUnknown(m)
}

var xxx_messageInfo_FeedPlaceholderView proto.InternalMessageInfo

func (m *FeedPlaceholderView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *FeedPlaceholderView) GetPlaceholderType() enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderType
	}
	return enums.PlaceholderTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*FeedPlaceholderView)(nil), "google.ads.googleads.v1.resources.FeedPlaceholderView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/feed_placeholder_view.proto", fileDescriptor_84882bb262f263ab)
}

var fileDescriptor_84882bb262f263ab = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x49, 0x02, 0x82, 0xc1, 0x2f, 0xe2, 0xa6, 0x5e, 0x04, 0xef, 0x15, 0x2e, 0xdc, 0x85,
	0xcc, 0x10, 0xbd, 0xab, 0x11, 0x17, 0x13, 0xd0, 0x82, 0x0b, 0x29, 0x45, 0x02, 0x4a, 0x20, 0x4c,
	0x33, 0xa7, 0x69, 0x20, 0x99, 0x89, 0x33, 0x49, 0x4a, 0x29, 0x5d, 0xfa, 0x22, 0x2e, 0x7d, 0x14,
	0x1f, 0xc1, 0x55, 0xd7, 0x7d, 0x04, 0x57, 0xd2, 0x7c, 0x35, 0x96, 0x56, 0x71, 0xf7, 0x67, 0xce,
	0xef, 0xfc, 0xcf, 0x39, 0x73, 0x8e, 0xfd, 0x26, 0x96, 0x32, 0x4e, 0x01, 0x33, 0xae, 0x71, 0x23,
	0xf7, 0xaa, 0x72, 0xb1, 0x02, 0x2d, 0x4b, 0x15, 0x81, 0xc6, 0x73, 0x00, 0x1e, 0xe6, 0x29, 0x8b,
	0x60, 0x21, 0x53, 0x0e, 0x2a, 0xac, 0x12, 0x58, 0xa2, 0x5c, 0xc9, 0x42, 0x3a, 0x57, 0x4d, 0x0e,
	0x62, 0x5c, 0xa3, 0x3e, 0x1d, 0x55, 0x2e, 0xea, 0xd3, 0x2f, 0x6e, 0xcf, 0x55, 0x00, 0x51, 0x66,
	0x1a, 0x0f, 0x8d, 0x8b, 0x55, 0x0e, 0x8d, 0xf1, 0xc5, 0xb3, 0x2e, 0x2b, 0x4f, 0xf0, 0x3c, 0x81,
	0x94, 0x87, 0x33, 0x58, 0xb0, 0x2a, 0x91, 0xaa, 0x05, 0x9e, 0x0c, 0x80, 0xae, 0x58, 0x1b, 0x7a,
	0x3a, 0x08, 0x31, 0x21, 0x64, 0xc1, 0x8a, 0x44, 0x0a, 0xdd, 0x44, 0x9f, 0xff, 0x34, 0xed, 0xc7,
	0xef, 0x00, 0xf8, 0xe4, 0x50, 0xd8, 0x4f, 0x60, 0xe9, 0x7c, 0xb2, 0xef, 0x77, 0x3e, 0xa1, 0x60,
	0x19, 0x8c, 0x8c, 0x4b, 0xe3, 0xe6, 0xae, 0x77, 0xbb, 0xa5, 0xd6, 0x2f, 0x8a, 0xec, 0x17, 0x87,
	0xf1, 0x5a, 0x95, 0x27, 0x1a, 0x45, 0x32, 0xc3, 0x27, 0xcc, 0xa6, 0xf7, 0x3a, 0xab, 0x0f, 0x2c,
	0x03, 0x47, 0xd9, 0x8f, 0x8e, 0xc7, 0x1c, 0x99, 0x97, 0xc6, 0xcd, 0x83, 0x97, 0x1e, 0x3a, 0xf7,
	0x81, 0xf5, 0xef, 0xa0, 0x81, 0xef, 0xc7, 0x55, 0x0e, 0x6f, 0x45, 0x99, 0x1d, 0xbf, 0x79, 0xd6,
	0x96, 0x5a, 0xd3, 0x87, 0xf9, 0x9f, 0xaf, 0xa4, 0xd8, 0xd1, 0x2f, 0xff, 0xd7, 0xb4, 0x43, 0xa3,
	0x52, 0x17, 0x32, 0x03, 0xa5, 0xf1, 0xba, 0x93, 0x9b, 0x7a, 0xfd, 0x47, 0xa4, 0xc6, 0xeb, 0x93,
	0x47, 0xb1, 0xf1, 0xbe, 0x9a, 0xf6, 0x75, 0x24, 0x33, 0xf4, 0xcf, 0xb3, 0xf0, 0x46, 0x27, 0x3a,
	0x98, 0xec, 0x17, 0x34, 0x31, 0x3e, 0xbf, 0x6f, 0xd3, 0x63, 0x99, 0x32, 0x11, 0x23, 0xa9, 0x62,
	0x1c, 0x83, 0xa8, 0xd7, 0x87, 0x0f, 0x73, 0xfc, 0xe5, 0x66, 0x5f, 0xf7, 0xea, 0x9b, 0x69, 0x8d,
	0x29, 0xfd, 0x6e, 0x5e, 0x8d, 0x1b, 0x4b, 0xca, 0x35, 0x6a, 0xe4, 0x5e, 0xf9, 0x2e, 0x9a, 0x76,
	0xe4, 0x8f, 0x8e, 0x09, 0x28, 0xd7, 0x41, 0xcf, 0x04, 0xbe, 0x1b, 0xf4, 0xcc, 0xce, 0xbc, 0x6e,
	0x02, 0x84, 0x50, 0xae, 0x09, 0xe9, 0x29, 0x42, 0x7c, 0x97, 0x90, 0x9e, 0x9b, 0xdd, 0xa9, 0x9b,
	0x7d, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x53, 0x57, 0x4c, 0xe1, 0x5f, 0x03, 0x00, 0x00,
}
