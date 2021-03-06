// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import route "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/route"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/filter/fault/v2"
import _type "github.com/cilium/cilium/pkg/envoy/envoy/type"
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

type FaultAbort struct {
	// An integer between 0-100 indicating the percentage of requests/operations/connections
	// that will be aborted with the error code provided.
	//
	// .. attention::
	//
	//   Use of integer `percent` value is deprecated. Use fractional `percentage` field instead.
	Percent uint32 `protobuf:"varint,1,opt,name=percent,proto3" json:"percent,omitempty"` // Deprecated: Do not use.
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}
func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (dst *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(dst, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *FaultAbort) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultAbort_OneofMarshaler, _FaultAbort_OneofUnmarshaler, _FaultAbort_OneofSizer, []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

func _FaultAbort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("FaultAbort.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _FaultAbort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultAbort)
	switch tag {
	case 2: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &FaultAbort_HttpStatus{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _FaultAbort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object. At least *abort* or *delay* must be specified.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percent
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percent>` field. The filter will
	// check the request's headers against all the specified headers in the filter
	// config. A match will happen if all the headers in the config are present in
	// the request with the same values (or based on presence if the *value* field
	// is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes      []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}
func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (dst *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(dst, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0xcd, 0xbf, 0x7b, 0xbd, 0x13, 0x2e, 0x96, 0x71, 0xe1, 0x50, 0x50, 0xd2, 0x82, 0x10,
	0x0b, 0x9d, 0x48, 0x5c, 0x8a, 0x82, 0xad, 0x94, 0x6e, 0x94, 0x12, 0xbb, 0x72, 0x53, 0xa6, 0xc9,
	0x69, 0x1b, 0x88, 0x99, 0x30, 0x99, 0x44, 0xfa, 0x3c, 0xbe, 0x84, 0xba, 0xea, 0xd2, 0x47, 0xf0,
	0x15, 0xfa, 0x16, 0x32, 0x33, 0x09, 0x75, 0x51, 0xb8, 0x9b, 0x30, 0x9c, 0xf9, 0x7d, 0xe7, 0x3b,
	0xe7, 0x9b, 0xa0, 0x29, 0x94, 0x2d, 0x3f, 0x46, 0x29, 0x2f, 0x77, 0xf9, 0x3e, 0xda, 0xe5, 0x85,
	0x04, 0x11, 0x1d, 0xa4, 0xac, 0xa2, 0x1d, 0x6b, 0x0a, 0x19, 0xb5, 0xb1, 0x39, 0xd0, 0x4a, 0x70,
	0xc9, 0xf1, 0x48, 0xe3, 0xd4, 0xe0, 0xd4, 0xe0, 0x54, 0xe1, 0xd4, 0x50, 0x6d, 0x3c, 0x7c, 0x61,
	0x3a, 0xb2, 0x2a, 0x57, 0x62, 0xc1, 0x1b, 0x09, 0xe6, 0x6b, 0x5a, 0x0c, 0xc3, 0x6b, 0x8e, 0xd7,
	0xcc, 0x86, 0xc4, 0x90, 0xf2, 0x58, 0x41, 0x54, 0x81, 0x48, 0xa1, 0xec, 0x6f, 0x9e, 0xb5, 0xac,
	0xc8, 0x33, 0x26, 0x21, 0xea, 0x0f, 0xe6, 0x62, 0xfc, 0xcb, 0x42, 0x68, 0xa1, 0x5a, 0x7c, 0xd8,
	0x72, 0x21, 0xf1, 0x4b, 0x74, 0xdb, 0x09, 0x89, 0x15, 0x58, 0xe1, 0xfd, 0xcc, 0xff, 0x7d, 0x3e,
	0x39, 0xee, 0xc4, 0x26, 0x19, 0xb1, 0x92, 0xfe, 0x0e, 0xbf, 0x46, 0xbe, 0xda, 0x61, 0x53, 0x4b,
	0x26, 0x9b, 0x9a, 0xd8, 0x1a, 0xbd, 0x57, 0xe8, 0xe3, 0xc9, 0xcd, 0xe0, 0xaf, 0x1b, 0xfe, 0xb1,
	0x96, 0x8f, 0x12, 0xa4, 0x98, 0x2f, 0x1a, 0xc1, 0xef, 0x10, 0xea, 0xc4, 0x6c, 0x0f, 0xc4, 0x09,
	0xac, 0xd0, 0x8f, 0x9f, 0x53, 0x13, 0x8e, 0x9a, 0x97, 0x2e, 0x04, 0x4b, 0x65, 0xce, 0x4b, 0x56,
	0xac, 0x0c, 0x97, 0xfc, 0x27, 0x98, 0x3d, 0x45, 0x08, 0x84, 0xe0, 0x62, 0xa3, 0x58, 0xec, 0xfd,
	0x3c, 0x9f, 0x1c, 0x6b, 0xfc, 0xc3, 0x46, 0x77, 0xcb, 0xf5, 0x7a, 0xa5, 0xe7, 0xc7, 0xef, 0x91,
	0x97, 0x41, 0xc1, 0x8e, 0x7a, 0x70, 0x3f, 0x0e, 0xe9, 0xb5, 0xe4, 0xfb, 0xd0, 0xa9, 0xd6, 0x7c,
	0x54, 0x7c, 0x62, 0x64, 0x78, 0x8e, 0x3c, 0xa6, 0x32, 0xd0, 0xdb, 0xf8, 0xf1, 0x94, 0x3e, 0xf8,
	0x72, 0xf4, 0x12, 0x5c, 0x62, 0xb4, 0xf8, 0x15, 0x1a, 0x34, 0x55, 0x2d, 0x05, 0xb0, 0x6f, 0x9b,
	0xb4, 0x68, 0x6a, 0x09, 0x42, 0x2f, 0x7b, 0x97, 0x3c, 0xe9, 0xeb, 0x73, 0x53, 0xc6, 0x6f, 0xd1,
	0xed, 0x01, 0x58, 0x06, 0xa2, 0x26, 0x6e, 0xe0, 0x84, 0x7e, 0x3c, 0xea, 0x1c, 0x59, 0x95, 0xab,
	0xe6, 0xe6, 0x17, 0x58, 0x6a, 0xe4, 0x13, 0x93, 0xe9, 0x01, 0x44, 0xd2, 0x2b, 0x94, 0x4f, 0xc6,
	0xbf, 0x97, 0x9d, 0x53, 0xc9, 0x33, 0xa8, 0x89, 0x17, 0x38, 0xca, 0xe7, 0x52, 0xff, 0xac, 0xca,
	0x33, 0xf7, 0xab, 0xdd, 0xc6, 0xdb, 0x1b, 0xfd, 0xdc, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x68, 0x1e, 0x13, 0x6b, 0xbf, 0x02, 0x00, 0x00,
}
