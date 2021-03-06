// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/build/v1/build_status.proto

package build

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// The end result of the Build.
type BuildStatus_Result int32

const (
	// Unspecified or unknown.
	BuildStatus_UNKNOWN_STATUS BuildStatus_Result = 0
	// Build was successful and tests (if requested) all pass.
	BuildStatus_COMMAND_SUCCEEDED BuildStatus_Result = 1
	// Build error and/or test failure.
	BuildStatus_COMMAND_FAILED BuildStatus_Result = 2
	// Unable to obtain a result due to input provided by the user.
	BuildStatus_USER_ERROR BuildStatus_Result = 3
	// Unable to obtain a result due to a failure within the build system.
	BuildStatus_SYSTEM_ERROR BuildStatus_Result = 4
	// Build required too many resources, such as build tool RAM.
	BuildStatus_RESOURCE_EXHAUSTED BuildStatus_Result = 5
	// An invocation attempt time exceeded its deadline.
	BuildStatus_INVOCATION_DEADLINE_EXCEEDED BuildStatus_Result = 6
	// Build request time exceeded the request_deadline
	BuildStatus_REQUEST_DEADLINE_EXCEEDED BuildStatus_Result = 8
	// The build was cancelled by a call to CancelBuild.
	BuildStatus_CANCELLED BuildStatus_Result = 7
)

var BuildStatus_Result_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "COMMAND_SUCCEEDED",
	2: "COMMAND_FAILED",
	3: "USER_ERROR",
	4: "SYSTEM_ERROR",
	5: "RESOURCE_EXHAUSTED",
	6: "INVOCATION_DEADLINE_EXCEEDED",
	8: "REQUEST_DEADLINE_EXCEEDED",
	7: "CANCELLED",
}

var BuildStatus_Result_value = map[string]int32{
	"UNKNOWN_STATUS":               0,
	"COMMAND_SUCCEEDED":            1,
	"COMMAND_FAILED":               2,
	"USER_ERROR":                   3,
	"SYSTEM_ERROR":                 4,
	"RESOURCE_EXHAUSTED":           5,
	"INVOCATION_DEADLINE_EXCEEDED": 6,
	"REQUEST_DEADLINE_EXCEEDED":    8,
	"CANCELLED":                    7,
}

func (x BuildStatus_Result) String() string {
	return proto.EnumName(BuildStatus_Result_name, int32(x))
}

func (BuildStatus_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5ec8917bb205301, []int{0, 0}
}

// Status used for both invocation attempt and overall build completion.
type BuildStatus struct {
	// The end result.
	Result BuildStatus_Result `protobuf:"varint,1,opt,name=result,proto3,enum=google.devtools.build.v1.BuildStatus_Result" json:"result,omitempty"`
	// Final invocation ID of the build, if there was one.
	// This field is only set on a status in BuildFinished event.
	FinalInvocationId string `protobuf:"bytes,3,opt,name=final_invocation_id,json=finalInvocationId,proto3" json:"final_invocation_id,omitempty"`
	// Build tool exit code. Integer value returned by the executed build tool.
	// Might not be available in some cases, e.g., a build timeout.
	BuildToolExitCode *wrappers.Int32Value `protobuf:"bytes,4,opt,name=build_tool_exit_code,json=buildToolExitCode,proto3" json:"build_tool_exit_code,omitempty"`
	// Fine-grained diagnostic information to complement the status.
	Details              *any.Any `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildStatus) Reset()         { *m = BuildStatus{} }
func (m *BuildStatus) String() string { return proto.CompactTextString(m) }
func (*BuildStatus) ProtoMessage()    {}
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ec8917bb205301, []int{0}
}

func (m *BuildStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildStatus.Unmarshal(m, b)
}
func (m *BuildStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildStatus.Marshal(b, m, deterministic)
}
func (m *BuildStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildStatus.Merge(m, src)
}
func (m *BuildStatus) XXX_Size() int {
	return xxx_messageInfo_BuildStatus.Size(m)
}
func (m *BuildStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildStatus.DiscardUnknown(m)
}

var xxx_messageInfo_BuildStatus proto.InternalMessageInfo

func (m *BuildStatus) GetResult() BuildStatus_Result {
	if m != nil {
		return m.Result
	}
	return BuildStatus_UNKNOWN_STATUS
}

func (m *BuildStatus) GetFinalInvocationId() string {
	if m != nil {
		return m.FinalInvocationId
	}
	return ""
}

func (m *BuildStatus) GetBuildToolExitCode() *wrappers.Int32Value {
	if m != nil {
		return m.BuildToolExitCode
	}
	return nil
}

func (m *BuildStatus) GetDetails() *any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.devtools.build.v1.BuildStatus_Result", BuildStatus_Result_name, BuildStatus_Result_value)
	proto.RegisterType((*BuildStatus)(nil), "google.devtools.build.v1.BuildStatus")
}

func init() {
	proto.RegisterFile("google/devtools/build/v1/build_status.proto", fileDescriptor_f5ec8917bb205301)
}

var fileDescriptor_f5ec8917bb205301 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0xc9, 0x3a, 0x3a, 0xe6, 0x41, 0x95, 0x9a, 0x81, 0xb2, 0x31, 0x50, 0xb5, 0x53, 0x25,
	0x90, 0xa3, 0x75, 0x47, 0xc4, 0x21, 0x4d, 0x8c, 0x88, 0x68, 0x93, 0x61, 0x27, 0xe3, 0xc7, 0xc5,
	0x4a, 0x1b, 0x2f, 0x8a, 0x64, 0xe2, 0x2a, 0x71, 0xca, 0xf6, 0x67, 0x72, 0xe0, 0x7f, 0xe1, 0x88,
	0xe2, 0x24, 0x68, 0x62, 0xec, 0x66, 0xf9, 0x7d, 0xde, 0x7b, 0x5f, 0xff, 0x00, 0xaf, 0x33, 0x29,
	0x33, 0xc1, 0xed, 0x94, 0x6f, 0x95, 0x94, 0xa2, 0xb2, 0x57, 0x75, 0x2e, 0x52, 0x7b, 0x7b, 0xd6,
	0x2e, 0x58, 0xa5, 0x12, 0x55, 0x57, 0x68, 0x53, 0x4a, 0x25, 0xa1, 0xd5, 0xc2, 0xa8, 0x87, 0x91,
	0x66, 0xd0, 0xf6, 0xec, 0xf8, 0x55, 0x17, 0xa3, 0xb9, 0x55, 0x7d, 0x65, 0xff, 0x28, 0x93, 0xcd,
	0x86, 0x97, 0x9d, 0xf3, 0xf8, 0xe8, 0x5f, 0x3d, 0x29, 0x6e, 0x5a, 0xe9, 0xf4, 0xd7, 0x00, 0x1c,
	0xcc, 0x9b, 0x1c, 0xaa, 0xab, 0xa0, 0x07, 0x86, 0x25, 0xaf, 0x6a, 0xa1, 0x2c, 0x63, 0x62, 0x4c,
	0x47, 0xb3, 0x37, 0xe8, 0xbe, 0x56, 0x74, 0xcb, 0x86, 0x88, 0xf6, 0x90, 0xce, 0x0b, 0x11, 0x78,
	0x7a, 0x95, 0x17, 0x89, 0x60, 0x79, 0xb1, 0x95, 0xeb, 0x44, 0xe5, 0xb2, 0x60, 0x79, 0x6a, 0x0d,
	0x26, 0xc6, 0x74, 0x9f, 0x8c, 0xb5, 0xe4, 0xff, 0x55, 0xfc, 0x14, 0x2e, 0xc0, 0x61, 0x7b, 0xe0,
	0xa6, 0x82, 0xf1, 0xeb, 0x5c, 0xb1, 0xb5, 0x4c, 0xb9, 0xb5, 0x3b, 0x31, 0xa6, 0x07, 0xb3, 0x17,
	0xfd, 0x0c, 0xfd, 0xfc, 0xc8, 0x2f, 0xd4, 0xf9, 0xec, 0x32, 0x11, 0x35, 0x27, 0x63, 0x6d, 0x8c,
	0xa4, 0x14, 0xf8, 0x3a, 0x57, 0xae, 0x4c, 0x39, 0x44, 0x60, 0x2f, 0xe5, 0x2a, 0xc9, 0x45, 0x65,
	0xed, 0xe8, 0x80, 0xc3, 0x3b, 0x01, 0x4e, 0x71, 0x43, 0x7a, 0xe8, 0xf4, 0xa7, 0x01, 0x86, 0xed,
	0x01, 0x20, 0x04, 0xa3, 0x38, 0xf8, 0x18, 0x84, 0x9f, 0x03, 0x46, 0x23, 0x27, 0x8a, 0xa9, 0xf9,
	0x00, 0x3e, 0x03, 0x63, 0x37, 0x5c, 0x2e, 0x9d, 0xc0, 0x63, 0x34, 0x76, 0x5d, 0x8c, 0x3d, 0xec,
	0x99, 0x46, 0x83, 0xf6, 0xdb, 0xef, 0x1d, 0x7f, 0x81, 0x3d, 0x73, 0x07, 0x8e, 0x00, 0x88, 0x29,
	0x26, 0x0c, 0x13, 0x12, 0x12, 0x73, 0x00, 0x4d, 0xf0, 0x98, 0x7e, 0xa5, 0x11, 0x5e, 0x76, 0x3b,
	0xbb, 0xf0, 0x39, 0x80, 0x04, 0xd3, 0x30, 0x26, 0x2e, 0x66, 0xf8, 0xcb, 0x07, 0x27, 0xa6, 0x11,
	0xf6, 0xcc, 0x87, 0x70, 0x02, 0x4e, 0xfc, 0xe0, 0x32, 0x74, 0x9d, 0xc8, 0x0f, 0x03, 0xe6, 0x61,
	0xc7, 0x5b, 0xf8, 0x41, 0x83, 0x74, 0x7d, 0x43, 0xf8, 0x12, 0x1c, 0x11, 0xfc, 0x29, 0xc6, 0x34,
	0xfa, 0x8f, 0xfc, 0x08, 0x3e, 0x01, 0xfb, 0xae, 0x13, 0xb8, 0x78, 0xd1, 0x4c, 0xb2, 0x37, 0x57,
	0xe0, 0x64, 0x2d, 0xbf, 0xdf, 0xfb, 0x78, 0x73, 0xf3, 0xd6, 0xeb, 0x5d, 0x34, 0xb7, 0x72, 0x61,
	0x7c, 0x7b, 0xd7, 0xd1, 0x99, 0x14, 0x49, 0x91, 0x21, 0x59, 0x66, 0x76, 0xc6, 0x0b, 0x7d, 0x67,
	0x76, 0x2b, 0x25, 0x9b, 0xbc, 0xba, 0xfb, 0x59, 0xdf, 0xea, 0xc5, 0x6f, 0xc3, 0x58, 0x0d, 0x35,
	0x7c, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x13, 0x41, 0xbc, 0xd8, 0x02, 0x00, 0x00,
}
