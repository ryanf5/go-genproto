// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [AdGroupFeedService.GetAdGroupFeed][google.ads.googleads.v1.services.AdGroupFeedService.GetAdGroupFeed].
type GetAdGroupFeedRequest struct {
	// Required. The resource name of the ad group feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupFeedRequest) Reset()         { *m = GetAdGroupFeedRequest{} }
func (m *GetAdGroupFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupFeedRequest) ProtoMessage()    {}
func (*GetAdGroupFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f37f6e6fd89895, []int{0}
}

func (m *GetAdGroupFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupFeedRequest.Unmarshal(m, b)
}
func (m *GetAdGroupFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupFeedRequest.Merge(m, src)
}
func (m *GetAdGroupFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupFeedRequest.Size(m)
}
func (m *GetAdGroupFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupFeedRequest proto.InternalMessageInfo

func (m *GetAdGroupFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupFeedService.MutateAdGroupFeeds][google.ads.googleads.v1.services.AdGroupFeedService.MutateAdGroupFeeds].
type MutateAdGroupFeedsRequest struct {
	// Required. The ID of the customer whose ad group feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad group feeds.
	Operations []*AdGroupFeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupFeedsRequest) Reset()         { *m = MutateAdGroupFeedsRequest{} }
func (m *MutateAdGroupFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupFeedsRequest) ProtoMessage()    {}
func (*MutateAdGroupFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f37f6e6fd89895, []int{1}
}

func (m *MutateAdGroupFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupFeedsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupFeedsRequest.Merge(m, src)
}
func (m *MutateAdGroupFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupFeedsRequest.Size(m)
}
func (m *MutateAdGroupFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupFeedsRequest proto.InternalMessageInfo

func (m *MutateAdGroupFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupFeedsRequest) GetOperations() []*AdGroupFeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group feed.
type AdGroupFeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupFeedOperation_Create
	//	*AdGroupFeedOperation_Update
	//	*AdGroupFeedOperation_Remove
	Operation            isAdGroupFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AdGroupFeedOperation) Reset()         { *m = AdGroupFeedOperation{} }
func (m *AdGroupFeedOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeedOperation) ProtoMessage()    {}
func (*AdGroupFeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f37f6e6fd89895, []int{2}
}

func (m *AdGroupFeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeedOperation.Unmarshal(m, b)
}
func (m *AdGroupFeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeedOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupFeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeedOperation.Merge(m, src)
}
func (m *AdGroupFeedOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeedOperation.Size(m)
}
func (m *AdGroupFeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeedOperation proto.InternalMessageInfo

func (m *AdGroupFeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupFeedOperation_Operation interface {
	isAdGroupFeedOperation_Operation()
}

type AdGroupFeedOperation_Create struct {
	Create *resources.AdGroupFeed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupFeedOperation_Update struct {
	Update *resources.AdGroupFeed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupFeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupFeedOperation_Create) isAdGroupFeedOperation_Operation() {}

func (*AdGroupFeedOperation_Update) isAdGroupFeedOperation_Operation() {}

func (*AdGroupFeedOperation_Remove) isAdGroupFeedOperation_Operation() {}

func (m *AdGroupFeedOperation) GetOperation() isAdGroupFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupFeedOperation) GetCreate() *resources.AdGroupFeed {
	if x, ok := m.GetOperation().(*AdGroupFeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupFeedOperation) GetUpdate() *resources.AdGroupFeed {
	if x, ok := m.GetOperation().(*AdGroupFeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupFeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupFeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupFeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupFeedOperation_Create)(nil),
		(*AdGroupFeedOperation_Update)(nil),
		(*AdGroupFeedOperation_Remove)(nil),
	}
}

// Response message for an ad group feed mutate.
type MutateAdGroupFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateAdGroupFeedsResponse) Reset()         { *m = MutateAdGroupFeedsResponse{} }
func (m *MutateAdGroupFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupFeedsResponse) ProtoMessage()    {}
func (*MutateAdGroupFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f37f6e6fd89895, []int{3}
}

func (m *MutateAdGroupFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupFeedsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupFeedsResponse.Merge(m, src)
}
func (m *MutateAdGroupFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupFeedsResponse.Size(m)
}
func (m *MutateAdGroupFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupFeedsResponse proto.InternalMessageInfo

func (m *MutateAdGroupFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupFeedsResponse) GetResults() []*MutateAdGroupFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad group feed mutate.
type MutateAdGroupFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupFeedResult) Reset()         { *m = MutateAdGroupFeedResult{} }
func (m *MutateAdGroupFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupFeedResult) ProtoMessage()    {}
func (*MutateAdGroupFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f37f6e6fd89895, []int{4}
}

func (m *MutateAdGroupFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupFeedResult.Unmarshal(m, b)
}
func (m *MutateAdGroupFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupFeedResult.Merge(m, src)
}
func (m *MutateAdGroupFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupFeedResult.Size(m)
}
func (m *MutateAdGroupFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupFeedResult proto.InternalMessageInfo

func (m *MutateAdGroupFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupFeedRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupFeedRequest")
	proto.RegisterType((*MutateAdGroupFeedsRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupFeedsRequest")
	proto.RegisterType((*AdGroupFeedOperation)(nil), "google.ads.googleads.v1.services.AdGroupFeedOperation")
	proto.RegisterType((*MutateAdGroupFeedsResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupFeedsResponse")
	proto.RegisterType((*MutateAdGroupFeedResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_feed_service.proto", fileDescriptor_83f37f6e6fd89895)
}

var fileDescriptor_83f37f6e6fd89895 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x41, 0x6f, 0xd3, 0x48,
	0x14, 0x5e, 0x3b, 0xab, 0xee, 0x76, 0xd2, 0x76, 0xa5, 0xd9, 0xed, 0x36, 0xcd, 0xae, 0xb4, 0x91,
	0x37, 0xda, 0xad, 0xa2, 0xca, 0x26, 0x41, 0x14, 0xe1, 0xb6, 0x48, 0x8e, 0x44, 0x52, 0x0e, 0xa5,
	0xc5, 0x15, 0x3d, 0x40, 0x24, 0x6b, 0x6a, 0x4f, 0x82, 0xa9, 0xed, 0x31, 0x33, 0x76, 0xa4, 0xaa,
	0xea, 0x85, 0x1f, 0xc0, 0x85, 0x7f, 0xc0, 0x91, 0x7f, 0xc0, 0x91, 0x6b, 0xaf, 0xdc, 0x7a, 0x40,
	0x3d, 0x20, 0x0e, 0x88, 0x5f, 0xc0, 0x01, 0x21, 0x7b, 0x3c, 0x89, 0xd3, 0x24, 0x8a, 0xe8, 0xed,
	0x79, 0xde, 0xf7, 0x7d, 0xef, 0xbd, 0x79, 0xef, 0x8d, 0xc1, 0x56, 0x8f, 0x90, 0x9e, 0x87, 0x35,
	0xe4, 0x30, 0x8d, 0x9b, 0x89, 0xd5, 0xaf, 0x6b, 0x0c, 0xd3, 0xbe, 0x6b, 0x63, 0xa6, 0x21, 0xc7,
	0xea, 0x51, 0x12, 0x87, 0x56, 0x17, 0x63, 0xc7, 0xca, 0x8e, 0xd5, 0x90, 0x92, 0x88, 0xc0, 0x0a,
	0xa7, 0xa8, 0xc8, 0x61, 0xea, 0x80, 0xad, 0xf6, 0xeb, 0xaa, 0x60, 0x97, 0x6f, 0x4d, 0xd3, 0xa7,
	0x98, 0x91, 0x98, 0x8e, 0x05, 0xe0, 0xc2, 0xe5, 0xbf, 0x05, 0x2d, 0x74, 0x35, 0x14, 0x04, 0x24,
	0x42, 0x91, 0x4b, 0x02, 0x96, 0x79, 0x57, 0x72, 0x5e, 0xdb, 0x73, 0x71, 0x10, 0x65, 0x8e, 0x7f,
	0x72, 0x8e, 0xae, 0x8b, 0x3d, 0xc7, 0x3a, 0xc2, 0x4f, 0x51, 0xdf, 0x25, 0x34, 0x03, 0xac, 0xe6,
	0x00, 0x22, 0x83, 0xcc, 0x95, 0xd5, 0xa2, 0xa5, 0x5f, 0x47, 0x71, 0x37, 0x13, 0xf0, 0x11, 0x3b,
	0xbe, 0x12, 0x96, 0x86, 0xb6, 0xc6, 0x22, 0x14, 0xc5, 0x59, 0x3e, 0xca, 0x33, 0xb0, 0xdc, 0xc6,
	0x91, 0xe1, 0xb4, 0x93, 0x32, 0x5a, 0x18, 0x3b, 0x26, 0x7e, 0x1e, 0x63, 0x16, 0xc1, 0x87, 0x60,
	0x51, 0x44, 0xb1, 0x02, 0xe4, 0xe3, 0x92, 0x54, 0x91, 0xd6, 0xe6, 0x9b, 0xeb, 0x97, 0x86, 0xfc,
	0xd5, 0xf8, 0x0f, 0x54, 0x87, 0x77, 0x96, 0x59, 0xa1, 0xcb, 0x54, 0x9b, 0xf8, 0x5a, 0x5e, 0x6b,
	0x41, 0x48, 0x3c, 0x40, 0x3e, 0x56, 0xbe, 0x48, 0x60, 0x75, 0x37, 0x8e, 0x50, 0x84, 0x73, 0x18,
	0x26, 0x02, 0x56, 0x41, 0xd1, 0x8e, 0x59, 0x44, 0x7c, 0x4c, 0x2d, 0xd7, 0xc9, 0xc2, 0x15, 0x2e,
	0x0d, 0xd9, 0x04, 0xe2, 0xfc, 0xbe, 0x03, 0x9f, 0x00, 0x40, 0x42, 0x4c, 0xf9, 0x9d, 0x96, 0xe4,
	0x4a, 0x61, 0xad, 0xd8, 0xd8, 0x50, 0x67, 0xf5, 0x52, 0xcd, 0x05, 0xdc, 0x13, 0xf4, 0x4c, 0x7c,
	0x28, 0x07, 0xff, 0x07, 0xbf, 0x85, 0x88, 0x46, 0x2e, 0xf2, 0xac, 0x2e, 0x72, 0xbd, 0x98, 0xe2,
	0x52, 0xa1, 0x22, 0xad, 0xfd, 0x6a, 0x2e, 0x65, 0xc7, 0x2d, 0x7e, 0x0a, 0xff, 0x05, 0x8b, 0x7d,
	0xe4, 0xb9, 0x0e, 0x8a, 0xb0, 0x45, 0x02, 0xef, 0xa4, 0xf4, 0x73, 0x0a, 0x5b, 0x10, 0x87, 0x7b,
	0x81, 0x77, 0xa2, 0xbc, 0x94, 0xc1, 0x1f, 0x93, 0xe2, 0xc2, 0x4d, 0x50, 0x8c, 0xc3, 0x94, 0x9b,
	0x74, 0x28, 0xe5, 0x16, 0x1b, 0x65, 0x51, 0x84, 0x68, 0xa2, 0xda, 0x4a, 0x9a, 0xb8, 0x8b, 0xd8,
	0xb1, 0x09, 0x38, 0x3c, 0xb1, 0xe1, 0x0e, 0x98, 0xb3, 0x29, 0x46, 0x11, 0x6f, 0x48, 0xb1, 0xa1,
	0x4e, 0x2d, 0x7e, 0x30, 0xa6, 0xf9, 0xea, 0x77, 0x7e, 0x32, 0x33, 0x7e, 0xa2, 0xc4, 0x75, 0x4b,
	0xf2, 0x75, 0x95, 0x38, 0x1f, 0x96, 0xc0, 0x1c, 0xc5, 0x3e, 0xe9, 0xf3, 0xeb, 0x9a, 0x4f, 0x3c,
	0xfc, 0xbb, 0x59, 0x04, 0xf3, 0x83, 0xfb, 0x55, 0xde, 0x4a, 0xa0, 0x3c, 0xa9, 0xff, 0x2c, 0x24,
	0x01, 0xc3, 0xb0, 0x05, 0x96, 0xaf, 0xdc, 0xbe, 0x85, 0x29, 0x25, 0x34, 0x15, 0x2d, 0x36, 0xa0,
	0x48, 0x8f, 0x86, 0xb6, 0x7a, 0x90, 0xce, 0xb0, 0xf9, 0xfb, 0x68, 0x5f, 0xee, 0x25, 0x70, 0x78,
	0x00, 0x7e, 0xa1, 0x98, 0xc5, 0x5e, 0x24, 0xe6, 0xe3, 0xce, 0xec, 0xf9, 0x18, 0x4b, 0xcb, 0x4c,
	0x15, 0x4c, 0xa1, 0xa4, 0xdc, 0x05, 0x2b, 0x53, 0x30, 0xc9, 0x30, 0x4c, 0xd8, 0x94, 0xd1, 0xd9,
	0x6f, 0x7c, 0x28, 0x00, 0x98, 0xa3, 0x1e, 0xf0, 0xc0, 0xf0, 0x9d, 0x04, 0x96, 0x46, 0xf7, 0x0f,
	0xde, 0x9e, 0x9d, 0xed, 0xc4, 0x8d, 0x2d, 0xff, 0x60, 0xff, 0x94, 0xd6, 0x85, 0x31, 0x9a, 0xf8,
	0x8b, 0xf7, 0x1f, 0x5f, 0xc9, 0x37, 0xa0, 0x9a, 0xbc, 0x71, 0xa7, 0x23, 0x9e, 0x6d, 0xb1, 0x81,
	0x4c, 0xab, 0x69, 0x28, 0xd7, 0x3c, 0xad, 0x76, 0x06, 0x3f, 0x49, 0x00, 0x8e, 0xb7, 0x15, 0x6e,
	0x5e, 0xe3, 0xd6, 0xc5, 0x63, 0x50, 0xde, 0xba, 0x1e, 0x99, 0x4f, 0x92, 0xf2, 0xe8, 0xc2, 0xf8,
	0x33, 0xf7, 0x96, 0xac, 0x0f, 0x57, 0x3c, 0x2d, 0x71, 0x43, 0xa9, 0x27, 0x25, 0x0e, 0x6b, 0x3a,
	0xcd, 0x81, 0xb7, 0x6b, 0x67, 0x23, 0x15, 0xea, 0x7e, 0x1a, 0x47, 0x97, 0x6a, 0xe5, 0xbf, 0xce,
	0x8d, 0xd2, 0xb4, 0x67, 0xaf, 0xf9, 0x4d, 0x02, 0x55, 0x9b, 0xf8, 0x33, 0xf3, 0x6e, 0xae, 0x8c,
	0x8f, 0xc1, 0x7e, 0xb2, 0xf2, 0xfb, 0xd2, 0xe3, 0x9d, 0x8c, 0xdc, 0x23, 0x1e, 0x0a, 0x7a, 0x2a,
	0xa1, 0x3d, 0xad, 0x87, 0x83, 0xf4, 0x41, 0xd0, 0x86, 0xe1, 0xa6, 0xff, 0xf0, 0x36, 0x85, 0xf1,
	0x5a, 0x2e, 0xb4, 0x0d, 0xe3, 0x8d, 0x5c, 0x69, 0x73, 0x41, 0xc3, 0x61, 0x2a, 0x37, 0x13, 0xeb,
	0xb0, 0xae, 0x66, 0x81, 0xd9, 0xb9, 0x80, 0x74, 0x0c, 0x87, 0x75, 0x06, 0x90, 0xce, 0x61, 0xbd,
	0x23, 0x20, 0x9f, 0xe5, 0x2a, 0x3f, 0xd7, 0x75, 0xc3, 0x61, 0xba, 0x3e, 0x00, 0xe9, 0xfa, 0x61,
	0x5d, 0xd7, 0x05, 0xec, 0x68, 0x2e, 0xcd, 0xf3, 0xe6, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0xe0, 0xcf, 0xfd, 0x97, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupFeedServiceClient is the client API for AdGroupFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupFeedServiceClient interface {
	// Returns the requested ad group feed in full detail.
	GetAdGroupFeed(ctx context.Context, in *GetAdGroupFeedRequest, opts ...grpc.CallOption) (*resources.AdGroupFeed, error)
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error)
}

type adGroupFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupFeedServiceClient(cc grpc.ClientConnInterface) AdGroupFeedServiceClient {
	return &adGroupFeedServiceClient{cc}
}

func (c *adGroupFeedServiceClient) GetAdGroupFeed(ctx context.Context, in *GetAdGroupFeedRequest, opts ...grpc.CallOption) (*resources.AdGroupFeed, error) {
	out := new(resources.AdGroupFeed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupFeedService/GetAdGroupFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupFeedServiceClient) MutateAdGroupFeeds(ctx context.Context, in *MutateAdGroupFeedsRequest, opts ...grpc.CallOption) (*MutateAdGroupFeedsResponse, error) {
	out := new(MutateAdGroupFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupFeedService/MutateAdGroupFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupFeedServiceServer is the server API for AdGroupFeedService service.
type AdGroupFeedServiceServer interface {
	// Returns the requested ad group feed in full detail.
	GetAdGroupFeed(context.Context, *GetAdGroupFeedRequest) (*resources.AdGroupFeed, error)
	// Creates, updates, or removes ad group feeds. Operation statuses are
	// returned.
	MutateAdGroupFeeds(context.Context, *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error)
}

// UnimplementedAdGroupFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupFeedServiceServer struct {
}

func (*UnimplementedAdGroupFeedServiceServer) GetAdGroupFeed(ctx context.Context, req *GetAdGroupFeedRequest) (*resources.AdGroupFeed, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupFeed not implemented")
}
func (*UnimplementedAdGroupFeedServiceServer) MutateAdGroupFeeds(ctx context.Context, req *MutateAdGroupFeedsRequest) (*MutateAdGroupFeedsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupFeeds not implemented")
}

func RegisterAdGroupFeedServiceServer(s *grpc.Server, srv AdGroupFeedServiceServer) {
	s.RegisterService(&_AdGroupFeedService_serviceDesc, srv)
}

func _AdGroupFeedService_GetAdGroupFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupFeedServiceServer).GetAdGroupFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupFeedService/GetAdGroupFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupFeedServiceServer).GetAdGroupFeed(ctx, req.(*GetAdGroupFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupFeedService_MutateAdGroupFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupFeedService/MutateAdGroupFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupFeedServiceServer).MutateAdGroupFeeds(ctx, req.(*MutateAdGroupFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupFeedService",
	HandlerType: (*AdGroupFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupFeed",
			Handler:    _AdGroupFeedService_GetAdGroupFeed_Handler,
		},
		{
			MethodName: "MutateAdGroupFeeds",
			Handler:    _AdGroupFeedService_MutateAdGroupFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_feed_service.proto",
}
