// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
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

type ThriftProxy_TransportType int32

const (
	// For every new connection, the Thrift proxy will determine which transport to use.
	ThriftProxy_AUTO_TRANSPORT ThriftProxy_TransportType = 0
	// The Thrift proxy will assume the client is using the Thrift framed transport.
	ThriftProxy_FRAMED ThriftProxy_TransportType = 1
	// The Thrift proxy will assume the client is using the Thrift unframed transport.
	ThriftProxy_UNFRAMED ThriftProxy_TransportType = 2
	// The Thrift proxy will assume the client is using the Thrift header transport.
	ThriftProxy_HEADER ThriftProxy_TransportType = 3
)

var ThriftProxy_TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var ThriftProxy_TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x ThriftProxy_TransportType) String() string {
	return proto.EnumName(ThriftProxy_TransportType_name, int32(x))
}

func (ThriftProxy_TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0, 0}
}

type ThriftProxy_ProtocolType int32

const (
	// For every new connection, the Thrift proxy will determine which protocol to use.
	// N.B. The older, non-strict binary protocol is not included in automatic protocol
	// detection.
	ThriftProxy_AUTO_PROTOCOL ThriftProxy_ProtocolType = 0
	// The Thrift proxy will assume the client is using the Thrift binary protocol.
	ThriftProxy_BINARY ThriftProxy_ProtocolType = 1
	// The Thrift proxy will assume the client is using the Thrift non-strict binary protocol.
	ThriftProxy_LAX_BINARY ThriftProxy_ProtocolType = 2
	// The Thrift proxy will assume the client is using the Thrift compact protocol.
	ThriftProxy_COMPACT ThriftProxy_ProtocolType = 3
)

var ThriftProxy_ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
}

var ThriftProxy_ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
}

func (x ThriftProxy_ProtocolType) String() string {
	return proto.EnumName(ThriftProxy_ProtocolType_name, int32(x))
}

func (ThriftProxy_ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0, 1}
}

// [#protodoc-title: Extensions Thrift Proxy]
// Thrift Proxy filter configuration.
// [#comment:next free field: 5]
type ThriftProxy struct {
	// Supplies the type of transport that the Thrift proxy should use. Defaults to `AUTO_TRANSPORT`.
	Transport ThriftProxy_TransportType `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy_TransportType" json:"transport,omitempty"`
	// Supplies the type of protocol that the Thrift proxy should use. Defaults to `AUTO_PROTOCOL`.
	Protocol ThriftProxy_ProtocolType `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy_ProtocolType" json:"protocol,omitempty"`
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig          *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8fab7646d88fc90, []int{0}
}
func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (dst *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(dst, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() ThriftProxy_TransportType {
	if m != nil {
		return m.Transport
	}
	return ThriftProxy_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ThriftProxy_ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ThriftProxy_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy")
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy_TransportType", ThriftProxy_TransportType_name, ThriftProxy_TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v2alpha1.ThriftProxy_ProtocolType", ThriftProxy_ProtocolType_name, ThriftProxy_ProtocolType_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto", fileDescriptor_e8fab7646d88fc90)
}

var fileDescriptor_e8fab7646d88fc90 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xbb, 0x8e, 0x29, 0xcd, 0x38, 0x8d, 0xcc, 0x0a, 0x09, 0x2b, 0x07, 0x64, 0xe5, 0x14,
	0x71, 0x58, 0xab, 0xe6, 0xcc, 0xc1, 0x76, 0x8c, 0x40, 0xa4, 0xb1, 0xb5, 0x6c, 0x25, 0xe0, 0x62,
	0x99, 0x62, 0x27, 0x16, 0x96, 0xd7, 0xda, 0x6c, 0x4d, 0x73, 0xe5, 0xc4, 0x7b, 0xf0, 0x08, 0xdc,
	0x38, 0xf1, 0x3a, 0xbc, 0x05, 0xf2, 0xda, 0xa6, 0x8d, 0x38, 0x95, 0xde, 0x66, 0xe6, 0x9f, 0xfd,
	0xfe, 0xd9, 0x1f, 0x96, 0x59, 0xd5, 0xf0, 0xbd, 0x73, 0xc9, 0xab, 0xbc, 0xd8, 0x38, 0x79, 0x51,
	0xca, 0x4c, 0x38, 0x55, 0x26, 0xbf, 0x70, 0xf1, 0xd9, 0x91, 0x5b, 0x51, 0xe4, 0x32, 0xa9, 0x05,
	0xbf, 0xde, 0x3b, 0x8d, 0x9b, 0x96, 0xf5, 0x36, 0x3d, 0x3b, 0x98, 0x92, 0x5a, 0x70, 0xc9, 0xf1,
	0x99, 0xa2, 0x90, 0x8e, 0x42, 0x3a, 0x0a, 0xe9, 0x29, 0xe4, 0x60, 0x7f, 0xa0, 0xcc, 0x5e, 0xdc,
	0xdd, 0x58, 0xf0, 0x2b, 0x99, 0x75, 0x8e, 0xb3, 0x27, 0x4d, 0x5a, 0x16, 0x9f, 0x52, 0x99, 0x39,
	0x43, 0xd1, 0x0b, 0x8f, 0x37, 0x7c, 0xc3, 0x55, 0xe9, 0xb4, 0x55, 0x37, 0x9d, 0xff, 0xd0, 0xc1,
	0x60, 0x0a, 0x1a, 0xb7, 0x4c, 0xdc, 0xc0, 0x58, 0x8a, 0xb4, 0xda, 0xd5, 0x5c, 0x48, 0x4b, 0xb3,
	0xd1, 0x62, 0xea, 0xae, 0xc8, 0x9d, 0x3f, 0x41, 0x6e, 0x21, 0x09, 0x1b, 0x78, 0x6c, 0x5f, 0x67,
	0x3e, 0xfc, 0xfc, 0xfd, 0x6b, 0xf4, 0xe0, 0x2b, 0xd2, 0x4c, 0x44, 0x6f, 0xac, 0xf0, 0x0e, 0x4e,
	0xd4, 0x41, 0x97, 0xbc, 0xb4, 0x46, 0xca, 0xf6, 0xcd, 0x3d, 0x6d, 0xe3, 0x1e, 0xf7, 0x8f, 0xeb,
	0x5f, 0x23, 0xfc, 0x0c, 0x8c, 0x9d, 0x4c, 0x5b, 0x4a, 0x96, 0x17, 0xd7, 0x16, 0xb2, 0xd1, 0x62,
	0xec, 0x8f, 0xdb, 0x55, 0x5d, 0x68, 0x36, 0xa2, 0xd0, 0xaa, 0xb1, 0x12, 0xf1, 0x16, 0x26, 0x2a,
	0xe6, 0xa4, 0xbb, 0xc7, 0xd2, 0x6d, 0xb4, 0x30, 0xdc, 0xf0, 0x3f, 0x8e, 0xa4, 0x2d, 0x26, 0x50,
	0x0f, 0xae, 0x44, 0x2a, 0x0b, 0x5e, 0x51, 0x43, 0xdc, 0xcc, 0xe6, 0x11, 0x9c, 0x1e, 0x44, 0x86,
	0x31, 0x4c, 0xbd, 0x0b, 0x16, 0x25, 0x8c, 0x7a, 0xeb, 0xb7, 0x71, 0x44, 0x99, 0x79, 0x84, 0x01,
	0x8e, 0x5f, 0x52, 0xef, 0x3c, 0x5c, 0x9a, 0x08, 0x4f, 0xe0, 0xe4, 0x62, 0xdd, 0x77, 0x5a, 0xab,
	0xbc, 0x0a, 0xbd, 0x65, 0x48, 0xcd, 0xd1, 0x4c, 0xff, 0xf6, 0xfd, 0xe9, 0xd1, 0x3c, 0x86, 0xc9,
	0xed, 0x30, 0xf0, 0x23, 0x38, 0x55, 0xbc, 0x98, 0x46, 0x2c, 0x0a, 0xa2, 0x55, 0x87, 0xf3, 0x5f,
	0xaf, 0x3d, 0xfa, 0xde, 0x44, 0x78, 0x0a, 0xb0, 0xf2, 0xde, 0x25, 0x7d, 0xaf, 0x61, 0x03, 0x1e,
	0x06, 0xd1, 0x79, 0xec, 0x05, 0x6c, 0x20, 0xfa, 0xfa, 0x07, 0xad, 0x71, 0x3f, 0x1e, 0xab, 0x20,
	0x9f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xe2, 0xc4, 0x10, 0x2b, 0x03, 0x00, 0x00,
}
