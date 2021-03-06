// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/feed_placeholder_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [FeedPlaceholderViewService.GetFeedPlaceholderView][google.ads.googleads.v1.services.FeedPlaceholderViewService.GetFeedPlaceholderView].
type GetFeedPlaceholderViewRequest struct {
	// Required. The resource name of the feed placeholder view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedPlaceholderViewRequest) Reset()         { *m = GetFeedPlaceholderViewRequest{} }
func (m *GetFeedPlaceholderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedPlaceholderViewRequest) ProtoMessage()    {}
func (*GetFeedPlaceholderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c806439cff2804f, []int{0}
}

func (m *GetFeedPlaceholderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Unmarshal(m, b)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.Merge(m, src)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Size(m)
}
func (m *GetFeedPlaceholderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedPlaceholderViewRequest proto.InternalMessageInfo

func (m *GetFeedPlaceholderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedPlaceholderViewRequest)(nil), "google.ads.googleads.v1.services.GetFeedPlaceholderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/feed_placeholder_view_service.proto", fileDescriptor_8c806439cff2804f)
}

var fileDescriptor_8c806439cff2804f = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x25, 0x11, 0x04, 0x83, 0x36, 0x29, 0xf4, 0x8c, 0xca, 0x2d, 0xc7, 0x15, 0x72, 0xc8, 0x0c,
	0x51, 0x39, 0x64, 0xe4, 0x90, 0x59, 0xc4, 0xb5, 0x3a, 0x16, 0x85, 0x80, 0x12, 0x08, 0x73, 0x99,
	0xef, 0x72, 0x03, 0x49, 0x26, 0xce, 0x64, 0x73, 0xa0, 0xd8, 0x58, 0xf8, 0x07, 0xac, 0x6d, 0x2c,
	0xfd, 0x29, 0xd7, 0xda, 0x59, 0x29, 0x58, 0xf9, 0x13, 0xac, 0x24, 0x3b, 0x99, 0xec, 0xae, 0x24,
	0x6e, 0xf7, 0xd8, 0xf7, 0xbe, 0xf7, 0x7d, 0xf3, 0xde, 0xc6, 0x7b, 0x9a, 0x49, 0x99, 0xe5, 0x80,
	0x19, 0xd7, 0xd8, 0xc0, 0x16, 0x35, 0x21, 0xd6, 0xa0, 0x1a, 0x91, 0x82, 0xc6, 0xa7, 0x00, 0x3c,
	0xa9, 0x72, 0x96, 0xc2, 0x99, 0xcc, 0x39, 0xa8, 0xa4, 0x11, 0x70, 0x9e, 0x74, 0x34, 0xaa, 0x94,
	0xac, 0xa5, 0x3f, 0x31, 0xa3, 0x88, 0x71, 0x8d, 0x7a, 0x17, 0xd4, 0x84, 0xc8, 0xba, 0x04, 0x47,
	0x63, 0x7b, 0x14, 0x68, 0xb9, 0x50, 0xa3, 0x8b, 0xcc, 0x82, 0xe0, 0xb6, 0x1d, 0xaf, 0x04, 0x66,
	0x65, 0x29, 0x6b, 0x56, 0x0b, 0x59, 0xea, 0x8e, 0xbd, 0xb1, 0xc6, 0xa6, 0xb9, 0x80, 0xb2, 0xee,
	0x88, 0xdd, 0x35, 0xe2, 0x54, 0x40, 0xce, 0x93, 0x13, 0x38, 0x63, 0x8d, 0x90, 0xaa, 0x13, 0xdc,
	0x5c, 0x13, 0xd8, 0x4b, 0x0c, 0xb5, 0xf7, 0xd6, 0xbb, 0x33, 0x83, 0xfa, 0x19, 0x00, 0x9f, 0xaf,
	0x6e, 0x8a, 0x04, 0x9c, 0xbf, 0x80, 0x37, 0x0b, 0xd0, 0xb5, 0xff, 0xca, 0xbb, 0x66, 0x47, 0x92,
	0x92, 0x15, 0xb0, 0xe3, 0x4c, 0x9c, 0xbb, 0x57, 0xa6, 0x0f, 0x7f, 0x50, 0xf7, 0x0f, 0x45, 0xde,
	0xbd, 0x55, 0x10, 0x1d, 0xaa, 0x84, 0x46, 0xa9, 0x2c, 0xf0, 0x90, 0xe7, 0x55, 0x6b, 0x75, 0xcc,
	0x0a, 0xb8, 0xff, 0xd9, 0xf5, 0x82, 0x01, 0xd5, 0x4b, 0x93, 0xa6, 0xff, 0xd3, 0xf1, 0xae, 0x0f,
	0xdf, 0xe6, 0x3f, 0x41, 0xdb, 0xaa, 0x40, 0xff, 0x7d, 0x55, 0x70, 0x38, 0x6a, 0xd0, 0x37, 0x85,
	0x06, 0xc6, 0xf7, 0x8e, 0xbf, 0xd3, 0xcd, 0x38, 0x3e, 0x7c, 0xfb, 0xf5, 0xc9, 0x7d, 0xe4, 0x1f,
	0xb6, 0x25, 0xbf, 0xdb, 0x60, 0x8e, 0xd2, 0x85, 0xae, 0x65, 0x01, 0x4a, 0xe3, 0x83, 0x65, 0xeb,
	0xff, 0x78, 0x69, 0x7c, 0xf0, 0x3e, 0xb8, 0x75, 0x41, 0x77, 0xc6, 0x12, 0x9c, 0x7e, 0x74, 0xbd,
	0xfd, 0x54, 0x16, 0x5b, 0xdf, 0x3a, 0xdd, 0x1d, 0x4f, 0x71, 0xde, 0xb6, 0x3c, 0x77, 0x5e, 0x3f,
	0xef, 0x4c, 0x32, 0x99, 0xb3, 0x32, 0x43, 0x52, 0x65, 0x38, 0x83, 0x72, 0xf9, 0x1f, 0xc0, 0xab,
	0xb5, 0xe3, 0x1f, 0xc8, 0x63, 0x0b, 0xbe, 0xb8, 0x97, 0x66, 0x94, 0x7e, 0x75, 0x27, 0x33, 0x63,
	0x48, 0xb9, 0x46, 0x06, 0xb6, 0x28, 0x0a, 0x51, 0xb7, 0x58, 0x5f, 0x58, 0x49, 0x4c, 0xb9, 0x8e,
	0x7b, 0x49, 0x1c, 0x85, 0xb1, 0x95, 0xfc, 0x76, 0xf7, 0xcd, 0xef, 0x84, 0x50, 0xae, 0x09, 0xe9,
	0x45, 0x84, 0x44, 0x21, 0x21, 0x56, 0x76, 0x72, 0x79, 0x79, 0xe7, 0x83, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8d, 0xae, 0x36, 0xef, 0xc7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedPlaceholderViewServiceClient is the client API for FeedPlaceholderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedPlaceholderViewServiceClient interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error)
}

type feedPlaceholderViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedPlaceholderViewServiceClient(cc grpc.ClientConnInterface) FeedPlaceholderViewServiceClient {
	return &feedPlaceholderViewServiceClient{cc}
}

func (c *feedPlaceholderViewServiceClient) GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error) {
	out := new(resources.FeedPlaceholderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedPlaceholderViewService/GetFeedPlaceholderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedPlaceholderViewServiceServer is the server API for FeedPlaceholderViewService service.
type FeedPlaceholderViewServiceServer interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(context.Context, *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error)
}

// UnimplementedFeedPlaceholderViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedPlaceholderViewServiceServer struct {
}

func (*UnimplementedFeedPlaceholderViewServiceServer) GetFeedPlaceholderView(ctx context.Context, req *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedPlaceholderView not implemented")
}

func RegisterFeedPlaceholderViewServiceServer(s *grpc.Server, srv FeedPlaceholderViewServiceServer) {
	s.RegisterService(&_FeedPlaceholderViewService_serviceDesc, srv)
}

func _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedPlaceholderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedPlaceholderViewService/GetFeedPlaceholderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, req.(*GetFeedPlaceholderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedPlaceholderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.FeedPlaceholderViewService",
	HandlerType: (*FeedPlaceholderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedPlaceholderView",
			Handler:    _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/feed_placeholder_view_service.proto",
}
