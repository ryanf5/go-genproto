// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/ad_group_feed.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// An ad group feed.
type AdGroupFeed struct {
	// Immutable. The resource name of the ad group feed.
	// Ad group feed resource names have the form:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}~{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The feed being linked to the ad group.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// Immutable. The ad group being linked to the feed.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// ad group. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v3.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the AdGroupFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Output only. Status of the ad group feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v3.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupFeed) Reset()         { *m = AdGroupFeed{} }
func (m *AdGroupFeed) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeed) ProtoMessage()    {}
func (*AdGroupFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_9221b5cf8e75124f, []int{0}
}

func (m *AdGroupFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeed.Unmarshal(m, b)
}
func (m *AdGroupFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeed.Marshal(b, m, deterministic)
}
func (m *AdGroupFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeed.Merge(m, src)
}
func (m *AdGroupFeed) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeed.Size(m)
}
func (m *AdGroupFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeed.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeed proto.InternalMessageInfo

func (m *AdGroupFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *AdGroupFeed) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *AdGroupFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *AdGroupFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupFeed)(nil), "google.ads.googleads.v3.resources.AdGroupFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/ad_group_feed.proto", fileDescriptor_9221b5cf8e75124f)
}

var fileDescriptor_9221b5cf8e75124f = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x55, 0xda, 0xad, 0x80, 0x07, 0x53, 0x97, 0x53, 0x98, 0x06, 0xeb, 0x26, 0x8a, 0x7a, 0x40,
	0x36, 0x34, 0xc0, 0x21, 0x5c, 0x48, 0x24, 0x56, 0x09, 0x01, 0x2a, 0x19, 0x2a, 0x02, 0x75, 0x8a,
	0xdc, 0xc4, 0x4d, 0xa3, 0x25, 0x76, 0x14, 0x27, 0x45, 0xd3, 0xb4, 0x2f, 0xc3, 0x0d, 0x3e, 0x0a,
	0x9f, 0x62, 0xe7, 0x7d, 0x04, 0x24, 0x24, 0x14, 0xc7, 0x49, 0xff, 0xa0, 0xd2, 0xde, 0x5e, 0xed,
	0xf7, 0xde, 0xef, 0x4f, 0x9e, 0x0b, 0x5e, 0xf8, 0x8c, 0xf9, 0x21, 0x41, 0xd8, 0xe3, 0xa8, 0x80,
	0x39, 0x9a, 0xea, 0x28, 0x21, 0x9c, 0x65, 0x89, 0x4b, 0x38, 0xc2, 0x9e, 0xe3, 0x27, 0x2c, 0x8b,
	0x9d, 0x31, 0x21, 0x1e, 0x8c, 0x13, 0x96, 0x32, 0xf5, 0xa8, 0xe0, 0x42, 0xec, 0x71, 0x58, 0xc9,
	0xe0, 0x54, 0x87, 0x95, 0x6c, 0xff, 0xe5, 0x2a, 0x67, 0x97, 0x45, 0x11, 0xa3, 0x28, 0xc2, 0xa9,
	0x3b, 0x09, 0xa8, 0xef, 0x8c, 0x33, 0xea, 0xa6, 0x01, 0xa3, 0x85, 0xf5, 0xfe, 0xf3, 0x55, 0x3a,
	0x42, 0xb3, 0x88, 0xa3, 0xbc, 0x09, 0x27, 0x0c, 0xe8, 0xb9, 0xc3, 0x53, 0x9c, 0x66, 0x7c, 0x33,
	0x55, 0x1c, 0x62, 0x97, 0x4c, 0x58, 0xe8, 0x91, 0xc4, 0x49, 0x2f, 0x62, 0x22, 0x55, 0x87, 0xa5,
	0x2a, 0x0e, 0xd0, 0x38, 0x20, 0xa1, 0xe7, 0x8c, 0xc8, 0x04, 0x4f, 0x03, 0x96, 0x48, 0xc2, 0xfd,
	0x39, 0x42, 0x39, 0x9a, 0xbc, 0x7a, 0x28, 0xaf, 0xc4, 0xaf, 0x51, 0x36, 0x46, 0xdf, 0x12, 0x1c,
	0xc7, 0x24, 0x29, 0x3b, 0x3a, 0x98, 0x93, 0x62, 0x4a, 0x59, 0x8a, 0xf3, 0x21, 0xe5, 0xed, 0xf1,
	0x8f, 0x6d, 0xb0, 0x63, 0x7a, 0xbd, 0x7c, 0xaf, 0x27, 0x84, 0x78, 0xea, 0x47, 0x70, 0xaf, 0xf4,
	0x77, 0x28, 0x8e, 0x88, 0xa6, 0xb4, 0x94, 0xce, 0x1d, 0xeb, 0xc9, 0xb5, 0xb9, 0xfd, 0xdb, 0x7c,
	0x0c, 0x1e, 0xcd, 0x96, 0x2c, 0x51, 0x1c, 0x70, 0xe8, 0xb2, 0x08, 0xcd, 0x99, 0xd8, 0x77, 0x4b,
	0x8b, 0x0f, 0x38, 0x22, 0xea, 0x67, 0xb0, 0x95, 0x2f, 0x4b, 0xab, 0xb5, 0x94, 0xce, 0x4e, 0xf7,
	0x40, 0x0a, 0x61, 0xd9, 0x2f, 0x3c, 0x4d, 0x93, 0x80, 0xfa, 0x03, 0x1c, 0x66, 0xc4, 0x6a, 0x8b,
	0x3a, 0x87, 0xe0, 0xc1, 0xca, 0x3a, 0xa2, 0x80, 0x30, 0x54, 0x31, 0xb8, 0x5d, 0x66, 0x42, 0xab,
	0x6f, 0x60, 0xde, 0x11, 0xe6, 0xc7, 0xa0, 0xb5, 0x6e, 0x08, 0xfb, 0x16, 0x2e, 0x80, 0xca, 0xc0,
	0xde, 0xf2, 0x27, 0xe3, 0xda, 0x56, 0xab, 0xde, 0xd9, 0xed, 0x5a, 0x70, 0x55, 0xf6, 0xc4, 0xa7,
	0x86, 0xfd, 0x99, 0xee, 0xd3, 0x45, 0x4c, 0xde, 0xd0, 0x2c, 0x5a, 0x3e, 0xb3, 0x9b, 0xf1, 0xe2,
	0x01, 0x57, 0xcf, 0xc0, 0xde, 0x3f, 0x81, 0xd4, 0xb6, 0xc5, 0x70, 0x4f, 0x57, 0x16, 0x2c, 0x92,
	0x0c, 0xdf, 0x4b, 0xe1, 0x89, 0xd4, 0xd9, 0xcd, 0x68, 0xe9, 0x44, 0x3d, 0x03, 0x8d, 0x22, 0xae,
	0x5a, 0xa3, 0xa5, 0x74, 0x76, 0xbb, 0xaf, 0xd7, 0x0c, 0x91, 0x6f, 0xfb, 0x5d, 0x40, 0xcf, 0x4f,
	0x85, 0x48, 0xcc, 0xb0, 0x78, 0x64, 0xd5, 0xaf, 0xcd, 0xba, 0x2d, 0x4d, 0x8d, 0xe1, 0x8d, 0xf9,
	0x65, 0xb3, 0x8c, 0xa8, 0xcf, 0xdc, 0x8c, 0xa7, 0x2c, 0x22, 0x09, 0x47, 0x97, 0x25, 0xbc, 0x42,
	0x78, 0xc6, 0xe0, 0xe8, 0x72, 0xe1, 0xc5, 0x5f, 0x59, 0x7f, 0x14, 0xd0, 0x76, 0x59, 0x04, 0xd7,
	0xbe, 0x79, 0xab, 0x39, 0x57, 0xa9, 0x9f, 0x47, 0xa1, 0xaf, 0x7c, 0x7d, 0x2b, 0x65, 0x3e, 0x0b,
	0x31, 0xf5, 0x21, 0x4b, 0x7c, 0xe4, 0x13, 0x2a, 0x82, 0x82, 0x66, 0x7d, 0xfe, 0xe7, 0x0f, 0xe8,
	0x55, 0x85, 0xbe, 0xd7, 0xea, 0x3d, 0xd3, 0xfc, 0x59, 0x3b, 0xea, 0x15, 0x96, 0xa6, 0xc7, 0x61,
	0x01, 0x73, 0x34, 0xd0, 0xa1, 0x5d, 0x32, 0x7f, 0x95, 0x9c, 0xa1, 0xe9, 0xf1, 0x61, 0xc5, 0x19,
	0x0e, 0xf4, 0x61, 0xc5, 0xb9, 0xa9, 0xb5, 0x8b, 0x0b, 0xc3, 0x30, 0x3d, 0x6e, 0x18, 0x15, 0xcb,
	0x30, 0x06, 0xba, 0x61, 0x54, 0xbc, 0x51, 0x43, 0x34, 0xab, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x7e, 0x19, 0x09, 0x2c, 0x05, 0x00, 0x00,
}
