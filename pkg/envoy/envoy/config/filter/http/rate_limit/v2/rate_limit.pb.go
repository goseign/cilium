// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rate_limit/v2/rate_limit.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny      bool     `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_rate_limit_ae6b3709d7e9d60d, []int{0}
}
func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (dst *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(dst, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.http.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rate_limit/v2/rate_limit.proto", fileDescriptor_rate_limit_ae6b3709d7e9d60d)
}

var fileDescriptor_rate_limit_ae6b3709d7e9d60d = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0x34, 0xed, 0xf7, 0x75, 0xaa, 0x88, 0x41, 0x30, 0x76, 0x61, 0x53, 0x17, 0x12,
	0xba, 0x98, 0x81, 0xb8, 0x10, 0xb7, 0xa5, 0x4b, 0xdd, 0x04, 0x57, 0x6e, 0xc2, 0xd4, 0xdc, 0xc4,
	0x81, 0x24, 0x37, 0x4e, 0x6f, 0x02, 0x79, 0x13, 0x9f, 0xc5, 0x95, 0xaf, 0xe1, 0xca, 0xb5, 0x6f,
	0x21, 0xf9, 0x07, 0xdd, 0xdd, 0x7b, 0xce, 0xf9, 0x1d, 0x38, 0xfc, 0x1e, 0x8a, 0x1a, 0x1b, 0xf9,
	0x8a, 0x45, 0xa2, 0x53, 0x99, 0xe8, 0x8c, 0xc0, 0xc8, 0x37, 0xa2, 0x52, 0x1a, 0x45, 0x10, 0x65,
	0x3a, 0xd7, 0x24, 0xeb, 0xe0, 0xe8, 0x13, 0xa5, 0x41, 0x42, 0xe7, 0xb6, 0x03, 0x45, 0x0f, 0x8a,
	0x1e, 0x14, 0x2d, 0x28, 0x8e, 0xa2, 0x75, 0xb0, 0xbc, 0x4e, 0x11, 0xd3, 0x0c, 0x64, 0x47, 0xed,
	0xab, 0x44, 0xc6, 0x95, 0x51, 0xa4, 0xb1, 0xe8, 0x7b, 0x96, 0x97, 0xb5, 0xca, 0x74, 0xac, 0x08,
	0xe4, 0x78, 0x0c, 0xc6, 0x45, 0x8a, 0x29, 0x76, 0xa7, 0x6c, 0xaf, 0x5e, 0xbd, 0xf9, 0x66, 0x7c,
	0x1e, 0x2a, 0x82, 0xc7, 0xb6, 0xdf, 0x59, 0xf3, 0x59, 0x8c, 0xb9, 0xd2, 0x85, 0xcb, 0x3c, 0xe6,
	0xcf, 0xb7, 0xf3, 0xcf, 0xdf, 0xaf, 0x89, 0x6d, 0x2c, 0x8f, 0x85, 0x83, 0xe1, 0xac, 0xf8, 0xf4,
	0x40, 0x2a, 0x05, 0xd7, 0xf2, 0x98, 0x7f, 0x3a, 0x24, 0x36, 0x96, 0xcb, 0xc3, 0x5e, 0x77, 0xd6,
	0xfc, 0xc4, 0xc0, 0x7b, 0x05, 0x07, 0x8a, 0xa8, 0x29, 0xc1, 0x9d, 0xb4, 0x4d, 0xe1, 0x62, 0xd0,
	0x9e, 0x9b, 0x12, 0x9c, 0x07, 0xfe, 0x8f, 0x74, 0x0e, 0x58, 0x91, 0x6b, 0x7b, 0xcc, 0x5f, 0x04,
	0x57, 0xa2, 0x5f, 0x25, 0xc6, 0x55, 0x62, 0x37, 0xac, 0xda, 0xda, 0x1f, 0x3f, 0x2b, 0x16, 0x8e,
	0x79, 0x67, 0xc3, 0xcf, 0x13, 0xa5, 0xb3, 0xca, 0x40, 0x94, 0x63, 0x0c, 0x51, 0x0c, 0x45, 0xe3,
	0x4e, 0x3d, 0xe6, 0xff, 0x0f, 0xcf, 0x06, 0xe3, 0x09, 0x63, 0xd8, 0x41, 0xd1, 0x6c, 0xed, 0x17,
	0xab, 0x0e, 0xf6, 0xb3, 0xae, 0xf3, 0xee, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x32, 0xe3, 0x0e,
	0x9a, 0x01, 0x00, 0x00,
}
