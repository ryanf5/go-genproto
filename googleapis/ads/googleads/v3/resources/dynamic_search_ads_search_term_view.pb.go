// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/dynamic_search_ads_search_term_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A dynamic search ads search term view.
type DynamicSearchAdsSearchTermView struct {
	// Output only. The resource name of the dynamic search ads search term view.
	// Dynamic search ads search term view resource names have the form:
	//
	// `customers/{customer_id}/dynamicSearchAdsSearchTermViews/{ad_group_id}~{search_term_fp}~{headline_fp}~{landing_page_fp}~{page_url_fp}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Search term
	//
	// This field is read-only.
	SearchTerm *wrappers.StringValue `protobuf:"bytes,2,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	// Output only. The dynamically generated headline of the Dynamic Search Ad.
	//
	// This field is read-only.
	Headline *wrappers.StringValue `protobuf:"bytes,3,opt,name=headline,proto3" json:"headline,omitempty"`
	// Output only. The dynamically selected landing page URL of the impression.
	//
	// This field is read-only.
	LandingPage *wrappers.StringValue `protobuf:"bytes,4,opt,name=landing_page,json=landingPage,proto3" json:"landing_page,omitempty"`
	// Output only. The URL of page feed item served for the impression.
	//
	// This field is read-only.
	PageUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	// Output only. True if query matches a negative keyword.
	//
	// This field is read-only.
	HasNegativeKeyword *wrappers.BoolValue `protobuf:"bytes,6,opt,name=has_negative_keyword,json=hasNegativeKeyword,proto3" json:"has_negative_keyword,omitempty"`
	// Output only. True if query is added to targeted keywords.
	//
	// This field is read-only.
	HasMatchingKeyword *wrappers.BoolValue `protobuf:"bytes,7,opt,name=has_matching_keyword,json=hasMatchingKeyword,proto3" json:"has_matching_keyword,omitempty"`
	// Output only. True if query matches a negative url.
	//
	// This field is read-only.
	HasNegativeUrl       *wrappers.BoolValue `protobuf:"bytes,8,opt,name=has_negative_url,json=hasNegativeUrl,proto3" json:"has_negative_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DynamicSearchAdsSearchTermView) Reset()         { *m = DynamicSearchAdsSearchTermView{} }
func (m *DynamicSearchAdsSearchTermView) String() string { return proto.CompactTextString(m) }
func (*DynamicSearchAdsSearchTermView) ProtoMessage()    {}
func (*DynamicSearchAdsSearchTermView) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0384f99ab33a76a, []int{0}
}

func (m *DynamicSearchAdsSearchTermView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Unmarshal(m, b)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Marshal(b, m, deterministic)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.Merge(m, src)
}
func (m *DynamicSearchAdsSearchTermView) XXX_Size() int {
	return xxx_messageInfo_DynamicSearchAdsSearchTermView.Size(m)
}
func (m *DynamicSearchAdsSearchTermView) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicSearchAdsSearchTermView.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicSearchAdsSearchTermView proto.InternalMessageInfo

func (m *DynamicSearchAdsSearchTermView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DynamicSearchAdsSearchTermView) GetSearchTerm() *wrappers.StringValue {
	if m != nil {
		return m.SearchTerm
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHeadline() *wrappers.StringValue {
	if m != nil {
		return m.Headline
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetLandingPage() *wrappers.StringValue {
	if m != nil {
		return m.LandingPage
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetPageUrl() *wrappers.StringValue {
	if m != nil {
		return m.PageUrl
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHasNegativeKeyword() *wrappers.BoolValue {
	if m != nil {
		return m.HasNegativeKeyword
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHasMatchingKeyword() *wrappers.BoolValue {
	if m != nil {
		return m.HasMatchingKeyword
	}
	return nil
}

func (m *DynamicSearchAdsSearchTermView) GetHasNegativeUrl() *wrappers.BoolValue {
	if m != nil {
		return m.HasNegativeUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*DynamicSearchAdsSearchTermView)(nil), "google.ads.googleads.v3.resources.DynamicSearchAdsSearchTermView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/dynamic_search_ads_search_term_view.proto", fileDescriptor_c0384f99ab33a76a)
}

var fileDescriptor_c0384f99ab33a76a = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6a, 0xd4, 0x40,
	0x1c, 0xc6, 0xc9, 0xae, 0xb6, 0x75, 0x5a, 0x45, 0x82, 0x87, 0xb8, 0x94, 0xda, 0x2a, 0x85, 0x9e,
	0x26, 0x60, 0x0e, 0x62, 0x54, 0x24, 0xa1, 0x50, 0xb4, 0x58, 0xd6, 0xad, 0xdd, 0x43, 0x59, 0x08,
	0xb3, 0x99, 0x7f, 0x27, 0x83, 0xc9, 0x4c, 0x98, 0xc9, 0xee, 0x52, 0x4a, 0xdf, 0xc1, 0x8b, 0x07,
	0x3d, 0x7a, 0xf4, 0x51, 0x7c, 0x8a, 0x9e, 0xfb, 0x08, 0x9e, 0x24, 0x9b, 0x4c, 0xba, 0x22, 0xae,
	0xc1, 0xdb, 0xb7, 0xfc, 0xbf, 0xef, 0x37, 0xff, 0x6f, 0xd8, 0x09, 0x3a, 0x64, 0x52, 0xb2, 0x14,
	0x5c, 0x42, 0xb5, 0x5b, 0xc9, 0x52, 0x4d, 0x3d, 0x57, 0x81, 0x96, 0x13, 0x15, 0x83, 0x76, 0xe9,
	0xb9, 0x20, 0x19, 0x8f, 0x23, 0x0d, 0x44, 0xc5, 0x49, 0x44, 0xa8, 0x36, 0xb2, 0x00, 0x95, 0x45,
	0x53, 0x0e, 0x33, 0x9c, 0x2b, 0x59, 0x48, 0x7b, 0xa7, 0x22, 0x60, 0x42, 0x35, 0x6e, 0x60, 0x78,
	0xea, 0xe1, 0x06, 0xd6, 0x7b, 0x64, 0xce, 0xcb, 0xb9, 0x7b, 0xc6, 0x21, 0xa5, 0xd1, 0x18, 0x12,
	0x32, 0xe5, 0x52, 0x55, 0x8c, 0xde, 0xc3, 0x05, 0x83, 0x89, 0xd5, 0xa3, 0xad, 0x7a, 0x34, 0xff,
	0x35, 0x9e, 0x9c, 0xb9, 0x33, 0x45, 0xf2, 0x1c, 0x94, 0xae, 0xe7, 0x9b, 0x0b, 0x51, 0x22, 0x84,
	0x2c, 0x48, 0xc1, 0xa5, 0xa8, 0xa7, 0x8f, 0xbf, 0xae, 0xa0, 0xad, 0xfd, 0xaa, 0xca, 0xf1, 0x7c,
	0xfd, 0x80, 0xea, 0x4a, 0x7c, 0x00, 0x95, 0x0d, 0x39, 0xcc, 0x6c, 0x8a, 0xee, 0x9a, 0x23, 0x23,
	0x41, 0x32, 0x70, 0xac, 0x6d, 0x6b, 0xef, 0x4e, 0xf8, 0xfa, 0x2a, 0xe8, 0xfe, 0x0c, 0x9e, 0xa3,
	0x67, 0x37, 0x9d, 0x6a, 0x95, 0x73, 0x8d, 0x63, 0x99, 0xb9, 0xcb, 0xb9, 0x83, 0x0d, 0x43, 0x3d,
	0x22, 0x19, 0xd8, 0x21, 0x5a, 0x5f, 0xb8, 0x3f, 0xa7, 0xb3, 0x6d, 0xed, 0xad, 0x3f, 0xdd, 0xac,
	0x91, 0xd8, 0x94, 0xc3, 0xc7, 0x85, 0xe2, 0x82, 0x0d, 0x49, 0x3a, 0x81, 0xb0, 0x7b, 0x15, 0x74,
	0x07, 0x48, 0x37, 0x54, 0xfb, 0x15, 0x5a, 0x4b, 0x80, 0xd0, 0x94, 0x0b, 0x70, 0xba, 0x6d, 0x01,
	0x4d, 0xc4, 0xde, 0x47, 0x1b, 0x29, 0x11, 0x94, 0x0b, 0x16, 0xe5, 0x84, 0x81, 0x73, 0xab, 0x2d,
	0x62, 0xbd, 0x8e, 0xf5, 0x09, 0x03, 0xfb, 0x25, 0x5a, 0x2b, 0xd3, 0xd1, 0x44, 0xa5, 0xce, 0xed,
	0xb6, 0x84, 0xd5, 0x32, 0x72, 0xa2, 0x52, 0xfb, 0x3d, 0x7a, 0x90, 0x10, 0x1d, 0x09, 0x60, 0xa4,
	0xe0, 0x53, 0x88, 0x3e, 0xc2, 0xf9, 0x4c, 0x2a, 0xea, 0xac, 0xcc, 0x49, 0xbd, 0x3f, 0x48, 0xa1,
	0x94, 0xe9, 0x02, 0xc7, 0x4e, 0x88, 0x3e, 0xaa, 0xb3, 0x87, 0x55, 0xd4, 0x20, 0x33, 0x52, 0xc4,
	0x49, 0xd9, 0xcd, 0x20, 0x57, 0xdb, 0x23, 0xdf, 0xd5, 0x59, 0x83, 0x7c, 0x83, 0xee, 0xff, 0xb6,
	0x65, 0xd9, 0x75, 0xad, 0x1d, 0xee, 0xde, 0xc2, 0x86, 0x27, 0x2a, 0xf5, 0xbf, 0x58, 0xd7, 0xc1,
	0x67, 0xeb, 0xbf, 0xff, 0x45, 0xf6, 0x69, 0x3c, 0xd1, 0x85, 0xcc, 0x40, 0x69, 0xf7, 0xc2, 0xc8,
	0x4b, 0xf3, 0x3a, 0xff, 0x12, 0xd2, 0xee, 0x45, 0x8b, 0xe7, 0x7b, 0x19, 0x7e, 0xea, 0xa0, 0xdd,
	0x58, 0x66, 0xf8, 0x9f, 0x0f, 0x38, 0x7c, 0xb2, 0x7c, 0xcb, 0x7e, 0x79, 0x19, 0x7d, 0xeb, 0xf4,
	0x6d, 0x4d, 0x62, 0x32, 0x25, 0x82, 0x61, 0xa9, 0x98, 0xcb, 0x40, 0xcc, 0xaf, 0xca, 0xbd, 0xa9,
	0xbd, 0xe4, 0xb3, 0xf3, 0xa2, 0x51, 0xdf, 0x3a, 0xdd, 0x83, 0x20, 0xf8, 0xde, 0xd9, 0x39, 0xa8,
	0x90, 0x01, 0xd5, 0xb8, 0x92, 0xa5, 0x1a, 0x7a, 0x78, 0x60, 0x9c, 0x3f, 0x8c, 0x67, 0x14, 0x50,
	0x3d, 0x6a, 0x3c, 0xa3, 0xa1, 0x37, 0x6a, 0x3c, 0xd7, 0x9d, 0xdd, 0x6a, 0xe0, 0xfb, 0x01, 0xd5,
	0xbe, 0xdf, 0xb8, 0x7c, 0x7f, 0xe8, 0xf9, 0x7e, 0xe3, 0x1b, 0xaf, 0xcc, 0x97, 0xf5, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x61, 0x54, 0x7a, 0x26, 0x22, 0x05, 0x00, 0x00,
}
