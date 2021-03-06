// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_bid_modifier.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// Represents a bid-modifiable only criterion at the campaign level.
type CampaignBidModifier struct {
	// Immutable. The resource name of the campaign bid modifier.
	// Campaign bid modifier resource names have the form:
	//
	// `customers/{customer_id}/campaignBidModifiers/{campaign_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign to which this criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The criterion of this campaign bid modifier.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignBidModifier_InteractionType
	Criterion            isCampaignBidModifier_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CampaignBidModifier) Reset()         { *m = CampaignBidModifier{} }
func (m *CampaignBidModifier) String() string { return proto.CompactTextString(m) }
func (*CampaignBidModifier) ProtoMessage()    {}
func (*CampaignBidModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_c88d20810b244847, []int{0}
}

func (m *CampaignBidModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBidModifier.Unmarshal(m, b)
}
func (m *CampaignBidModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBidModifier.Marshal(b, m, deterministic)
}
func (m *CampaignBidModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBidModifier.Merge(m, src)
}
func (m *CampaignBidModifier) XXX_Size() int {
	return xxx_messageInfo_CampaignBidModifier.Size(m)
}
func (m *CampaignBidModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBidModifier.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBidModifier proto.InternalMessageInfo

func (m *CampaignBidModifier) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignBidModifier) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignBidModifier) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

type isCampaignBidModifier_Criterion interface {
	isCampaignBidModifier_Criterion()
}

type CampaignBidModifier_InteractionType struct {
	InteractionType *common.InteractionTypeInfo `protobuf:"bytes,5,opt,name=interaction_type,json=interactionType,proto3,oneof"`
}

func (*CampaignBidModifier_InteractionType) isCampaignBidModifier_Criterion() {}

func (m *CampaignBidModifier) GetCriterion() isCampaignBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignBidModifier) GetInteractionType() *common.InteractionTypeInfo {
	if x, ok := m.GetCriterion().(*CampaignBidModifier_InteractionType); ok {
		return x.InteractionType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignBidModifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignBidModifier_InteractionType)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignBidModifier)(nil), "google.ads.googleads.v1.resources.CampaignBidModifier")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_bid_modifier.proto", fileDescriptor_c88d20810b244847)
}

var fileDescriptor_c88d20810b244847 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x77, 0x77, 0x5b, 0xb1, 0xb3, 0x2b, 0x4a, 0xbc, 0xc4, 0xb5, 0x68, 0x57, 0x29, 0x54, 0xd0,
	0x19, 0xd6, 0x16, 0x0f, 0x11, 0x91, 0x49, 0x85, 0xba, 0x82, 0x52, 0xa2, 0x2c, 0x28, 0x0b, 0x61,
	0x92, 0x99, 0x8d, 0x03, 0x9b, 0x99, 0x38, 0x93, 0x5d, 0x29, 0xd2, 0xa3, 0x5f, 0xc4, 0xa3, 0xdf,
	0xc3, 0x8b, 0x9f, 0x62, 0xcf, 0xfd, 0x08, 0x3d, 0x49, 0x92, 0xc9, 0x34, 0xe2, 0xae, 0xc5, 0xdb,
	0x0b, 0xef, 0xf7, 0xe7, 0xbd, 0x5f, 0xe6, 0x81, 0xe7, 0x89, 0x94, 0xc9, 0x8c, 0x21, 0x42, 0x35,
	0xaa, 0xca, 0xa2, 0x5a, 0x0c, 0x91, 0x62, 0x5a, 0xce, 0x55, 0xcc, 0x34, 0x8a, 0x49, 0x9a, 0x11,
	0x9e, 0x88, 0x30, 0xe2, 0x34, 0x4c, 0x25, 0xe5, 0x53, 0xce, 0x14, 0xcc, 0x94, 0xcc, 0xa5, 0x33,
	0xa8, 0x38, 0x90, 0x50, 0x0d, 0x2d, 0x1d, 0x2e, 0x86, 0xd0, 0xd2, 0xfb, 0x8f, 0xd7, 0x39, 0xc4,
	0x32, 0x4d, 0xa5, 0x40, 0xb1, 0xe2, 0x39, 0x53, 0x9c, 0x54, 0x8a, 0xfd, 0x7b, 0x35, 0x3c, 0xe3,
	0x68, 0xca, 0xd9, 0x8c, 0x86, 0x11, 0xfb, 0x44, 0x16, 0x5c, 0x1a, 0xcb, 0xfe, 0xed, 0x06, 0xa0,
	0x76, 0x31, 0xad, 0xbb, 0xa6, 0x55, 0x7e, 0x45, 0xf3, 0x29, 0xfa, 0xa2, 0x48, 0x96, 0x31, 0xa5,
	0x4d, 0x7f, 0xbb, 0x41, 0x25, 0x42, 0xc8, 0x9c, 0xe4, 0x5c, 0x0a, 0xd3, 0xbd, 0xff, 0x73, 0x03,
	0xdc, 0x3a, 0x34, 0xbb, 0xfa, 0x9c, 0xbe, 0x31, 0x9b, 0x3a, 0x1f, 0xc0, 0xf5, 0xda, 0x27, 0x14,
	0x24, 0x65, 0x6e, 0x6b, 0xa7, 0xb5, 0xb7, 0xe5, 0x1f, 0x2c, 0xf1, 0xe6, 0x39, 0x86, 0xe0, 0xd1,
	0xc5, 0xde, 0xa6, 0xca, 0xb8, 0x86, 0xb1, 0x4c, 0xd1, 0x0a, 0xb1, 0xa0, 0x57, 0x4b, 0xbd, 0x25,
	0x29, 0x73, 0x62, 0x70, 0xad, 0x4e, 0xd7, 0x6d, 0xef, 0xb4, 0xf6, 0xba, 0x4f, 0xb6, 0x8d, 0x08,
	0xac, 0x77, 0x80, 0xef, 0x72, 0xc5, 0x45, 0x32, 0x26, 0xb3, 0x39, 0xf3, 0x1f, 0x2e, 0x71, 0xe7,
	0x1c, 0x3f, 0x00, 0x83, 0x4b, 0x3d, 0x03, 0x2b, 0xec, 0x1c, 0x82, 0x9e, 0xc9, 0x58, 0x8a, 0x90,
	0x53, 0xb7, 0x53, 0x1a, 0xdd, 0xf9, 0xcb, 0x68, 0x24, 0xf2, 0xa7, 0x07, 0x95, 0x4f, 0x67, 0x89,
	0x3b, 0x41, 0xd7, 0xb2, 0x46, 0xd4, 0x79, 0x01, 0x7a, 0xcd, 0xdf, 0xef, 0x6e, 0xac, 0x99, 0xf6,
	0xa5, 0x9c, 0x47, 0x33, 0x56, 0xaa, 0x04, 0xdd, 0xa8, 0x91, 0x22, 0x03, 0x37, 0xb9, 0xc8, 0x99,
	0x22, 0x71, 0x91, 0x79, 0x98, 0x9f, 0x64, 0xcc, 0xdd, 0x2c, 0x45, 0xf6, 0xe1, 0xba, 0x47, 0x54,
	0xbd, 0x90, 0x62, 0xb0, 0x9a, 0xf7, 0xfe, 0x24, 0x63, 0x23, 0x31, 0x95, 0xc5, 0x84, 0x9b, 0xaf,
	0xae, 0x04, 0x37, 0xf8, 0x9f, 0x3d, 0x2f, 0x3f, 0xc3, 0x9f, 0xff, 0xef, 0x97, 0x38, 0x38, 0x9e,
	0xeb, 0x5c, 0xa6, 0x4c, 0x69, 0xf4, 0xb5, 0x2e, 0x4f, 0xed, 0xab, 0x6f, 0x20, 0x8b, 0xfe, 0xaa,
	0x5b, 0x38, 0xf5, 0xbb, 0x60, 0xcb, 0x86, 0xe5, 0x7f, 0x6b, 0x83, 0xdd, 0x58, 0xa6, 0xf0, 0xd2,
	0xd3, 0xf0, 0xdd, 0x15, 0xe3, 0x1c, 0x17, 0x49, 0x1e, 0xb7, 0x3e, 0xbe, 0x36, 0xf4, 0x44, 0xce,
	0x88, 0x48, 0xa0, 0x54, 0x09, 0x4a, 0x98, 0x28, 0x73, 0x46, 0x17, 0x4b, 0xfd, 0xe3, 0x6e, 0x9f,
	0xd9, 0xea, 0x7b, 0xbb, 0x73, 0x84, 0xf1, 0x8f, 0xf6, 0xe0, 0xa8, 0x92, 0xc4, 0x54, 0xc3, 0xaa,
	0x2c, 0xaa, 0xf1, 0x10, 0x06, 0x35, 0xf2, 0x57, 0x8d, 0x99, 0x60, 0xaa, 0x27, 0x16, 0x33, 0x19,
	0x0f, 0x27, 0x16, 0x73, 0xd6, 0xde, 0xad, 0x1a, 0x9e, 0x87, 0xa9, 0xf6, 0x3c, 0x8b, 0xf2, 0xbc,
	0xf1, 0xd0, 0xf3, 0x2c, 0x2e, 0xba, 0x5a, 0x0e, 0xbb, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x9e,
	0x30, 0xa2, 0xa3, 0x63, 0x04, 0x00, 0x00,
}
