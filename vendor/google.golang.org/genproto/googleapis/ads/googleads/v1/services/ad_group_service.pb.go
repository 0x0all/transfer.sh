// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/ad_group_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for [AdGroupService.GetAdGroup][google.ads.googleads.v1.services.AdGroupService.GetAdGroup].
type GetAdGroupRequest struct {
	// The resource name of the ad group to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupRequest) Reset()         { *m = GetAdGroupRequest{} }
func (m *GetAdGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupRequest) ProtoMessage()    {}
func (*GetAdGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_service_d5cdb2e7bdfeb784, []int{0}
}
func (m *GetAdGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupRequest.Unmarshal(m, b)
}
func (m *GetAdGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupRequest.Marshal(b, m, deterministic)
}
func (dst *GetAdGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupRequest.Merge(dst, src)
}
func (m *GetAdGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupRequest.Size(m)
}
func (m *GetAdGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupRequest proto.InternalMessageInfo

func (m *GetAdGroupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupService.MutateAdGroups][google.ads.googleads.v1.services.AdGroupService.MutateAdGroups].
type MutateAdGroupsRequest struct {
	// The ID of the customer whose ad groups are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ad groups.
	Operations []*AdGroupOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupsRequest) Reset()         { *m = MutateAdGroupsRequest{} }
func (m *MutateAdGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupsRequest) ProtoMessage()    {}
func (*MutateAdGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_service_d5cdb2e7bdfeb784, []int{1}
}
func (m *MutateAdGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupsRequest.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupsRequest.Merge(dst, src)
}
func (m *MutateAdGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupsRequest.Size(m)
}
func (m *MutateAdGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupsRequest proto.InternalMessageInfo

func (m *MutateAdGroupsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupsRequest) GetOperations() []*AdGroupOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group.
type AdGroupOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupOperation_Create
	//	*AdGroupOperation_Update
	//	*AdGroupOperation_Remove
	Operation            isAdGroupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AdGroupOperation) Reset()         { *m = AdGroupOperation{} }
func (m *AdGroupOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupOperation) ProtoMessage()    {}
func (*AdGroupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_service_d5cdb2e7bdfeb784, []int{2}
}
func (m *AdGroupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupOperation.Unmarshal(m, b)
}
func (m *AdGroupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupOperation.Marshal(b, m, deterministic)
}
func (dst *AdGroupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupOperation.Merge(dst, src)
}
func (m *AdGroupOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupOperation.Size(m)
}
func (m *AdGroupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupOperation proto.InternalMessageInfo

func (m *AdGroupOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupOperation_Operation interface {
	isAdGroupOperation_Operation()
}

type AdGroupOperation_Create struct {
	Create *resources.AdGroup `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupOperation_Update struct {
	Update *resources.AdGroup `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupOperation_Create) isAdGroupOperation_Operation() {}

func (*AdGroupOperation_Update) isAdGroupOperation_Operation() {}

func (*AdGroupOperation_Remove) isAdGroupOperation_Operation() {}

func (m *AdGroupOperation) GetOperation() isAdGroupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupOperation) GetCreate() *resources.AdGroup {
	if x, ok := m.GetOperation().(*AdGroupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupOperation) GetUpdate() *resources.AdGroup {
	if x, ok := m.GetOperation().(*AdGroupOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AdGroupOperation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AdGroupOperation_OneofMarshaler, _AdGroupOperation_OneofUnmarshaler, _AdGroupOperation_OneofSizer, []interface{}{
		(*AdGroupOperation_Create)(nil),
		(*AdGroupOperation_Update)(nil),
		(*AdGroupOperation_Remove)(nil),
	}
}

func _AdGroupOperation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AdGroupOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupOperation_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *AdGroupOperation_Update:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *AdGroupOperation_Remove:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Remove)
	case nil:
	default:
		return fmt.Errorf("AdGroupOperation.Operation has unexpected type %T", x)
	}
	return nil
}

func _AdGroupOperation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AdGroupOperation)
	switch tag {
	case 1: // operation.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroup)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupOperation_Create{msg}
		return true, err
	case 2: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(resources.AdGroup)
		err := b.DecodeMessage(msg)
		m.Operation = &AdGroupOperation_Update{msg}
		return true, err
	case 3: // operation.remove
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &AdGroupOperation_Remove{x}
		return true, err
	default:
		return false, nil
	}
}

func _AdGroupOperation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AdGroupOperation)
	// operation
	switch x := m.Operation.(type) {
	case *AdGroupOperation_Create:
		s := proto.Size(x.Create)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupOperation_Update:
		s := proto.Size(x.Update)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AdGroupOperation_Remove:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Remove)))
		n += len(x.Remove)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for an ad group mutate.
type MutateAdGroupsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MutateAdGroupsResponse) Reset()         { *m = MutateAdGroupsResponse{} }
func (m *MutateAdGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupsResponse) ProtoMessage()    {}
func (*MutateAdGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_service_d5cdb2e7bdfeb784, []int{3}
}
func (m *MutateAdGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupsResponse.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupsResponse.Merge(dst, src)
}
func (m *MutateAdGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupsResponse.Size(m)
}
func (m *MutateAdGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupsResponse proto.InternalMessageInfo

func (m *MutateAdGroupsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupsResponse) GetResults() []*MutateAdGroupResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad group mutate.
type MutateAdGroupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupResult) Reset()         { *m = MutateAdGroupResult{} }
func (m *MutateAdGroupResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupResult) ProtoMessage()    {}
func (*MutateAdGroupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_service_d5cdb2e7bdfeb784, []int{4}
}
func (m *MutateAdGroupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupResult.Unmarshal(m, b)
}
func (m *MutateAdGroupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupResult.Marshal(b, m, deterministic)
}
func (dst *MutateAdGroupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupResult.Merge(dst, src)
}
func (m *MutateAdGroupResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupResult.Size(m)
}
func (m *MutateAdGroupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupResult proto.InternalMessageInfo

func (m *MutateAdGroupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupRequest)(nil), "google.ads.googleads.v1.services.GetAdGroupRequest")
	proto.RegisterType((*MutateAdGroupsRequest)(nil), "google.ads.googleads.v1.services.MutateAdGroupsRequest")
	proto.RegisterType((*AdGroupOperation)(nil), "google.ads.googleads.v1.services.AdGroupOperation")
	proto.RegisterType((*MutateAdGroupsResponse)(nil), "google.ads.googleads.v1.services.MutateAdGroupsResponse")
	proto.RegisterType((*MutateAdGroupResult)(nil), "google.ads.googleads.v1.services.MutateAdGroupResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupServiceClient is the client API for AdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupServiceClient interface {
	// Returns the requested ad group in full detail.
	GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error)
}

type adGroupServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupServiceClient(cc *grpc.ClientConn) AdGroupServiceClient {
	return &adGroupServiceClient{cc}
}

func (c *adGroupServiceClient) GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error) {
	out := new(resources.AdGroup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupService/GetAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupServiceClient) MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error) {
	out := new(MutateAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.AdGroupService/MutateAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupServiceServer is the server API for AdGroupService service.
type AdGroupServiceServer interface {
	// Returns the requested ad group in full detail.
	GetAdGroup(context.Context, *GetAdGroupRequest) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	MutateAdGroups(context.Context, *MutateAdGroupsRequest) (*MutateAdGroupsResponse, error)
}

func RegisterAdGroupServiceServer(s *grpc.Server, srv AdGroupServiceServer) {
	s.RegisterService(&_AdGroupService_serviceDesc, srv)
}

func _AdGroupService_GetAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupService/GetAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, req.(*GetAdGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupService_MutateAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.AdGroupService/MutateAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, req.(*MutateAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.AdGroupService",
	HandlerType: (*AdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroup",
			Handler:    _AdGroupService_GetAdGroup_Handler,
		},
		{
			MethodName: "MutateAdGroups",
			Handler:    _AdGroupService_MutateAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/ad_group_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/ad_group_service.proto", fileDescriptor_ad_group_service_d5cdb2e7bdfeb784)
}

var fileDescriptor_ad_group_service_d5cdb2e7bdfeb784 = []byte{
	// 702 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x4f, 0xd4, 0x40,
	0x14, 0xb7, 0xc5, 0xa0, 0x4c, 0x11, 0x75, 0x08, 0xba, 0xd9, 0x18, 0xdd, 0x54, 0x12, 0xc9, 0x46,
	0x5b, 0xb7, 0xc4, 0x40, 0x4a, 0x38, 0x2c, 0x51, 0x16, 0x0f, 0x08, 0x29, 0x09, 0x07, 0xb3, 0x49,
	0x33, 0x6c, 0x87, 0xa6, 0xa1, 0xed, 0xd4, 0x99, 0xe9, 0x1a, 0x42, 0xb8, 0xf0, 0x15, 0xfc, 0x04,
	0x7a, 0xf4, 0xe6, 0x07, 0xf0, 0x0b, 0x78, 0xf5, 0xe6, 0xd9, 0x93, 0x67, 0x4f, 0x9e, 0xcc, 0x74,
	0x66, 0x16, 0x76, 0x95, 0xac, 0xeb, 0xed, 0xcd, 0x9b, 0xf7, 0xfb, 0xbd, 0xdf, 0xbc, 0x3f, 0x03,
	0x56, 0x62, 0x42, 0xe2, 0x14, 0xbb, 0x28, 0x62, 0xae, 0x34, 0x85, 0xd5, 0x6f, 0xb9, 0x0c, 0xd3,
	0x7e, 0xd2, 0xc3, 0xcc, 0x45, 0x51, 0x18, 0x53, 0x52, 0x16, 0xa1, 0xf2, 0x38, 0x05, 0x25, 0x9c,
	0xc0, 0x86, 0x8c, 0x76, 0x50, 0xc4, 0x9c, 0x01, 0xd0, 0xe9, 0xb7, 0x1c, 0x0d, 0xac, 0x3f, 0xbd,
	0x8c, 0x9a, 0x62, 0x46, 0x4a, 0x7a, 0x91, 0x5b, 0x72, 0xd6, 0xef, 0x69, 0x44, 0x91, 0xb8, 0x28,
	0xcf, 0x09, 0x47, 0x3c, 0x21, 0x39, 0x53, 0xb7, 0x2a, 0xa3, 0x5b, 0x9d, 0x0e, 0xca, 0x43, 0xf7,
	0x30, 0xc1, 0x69, 0x14, 0x66, 0x88, 0x1d, 0xa9, 0x88, 0xfb, 0xa3, 0x11, 0x6f, 0x29, 0x2a, 0x0a,
	0x4c, 0x35, 0xc3, 0x5d, 0x75, 0x4f, 0x8b, 0x9e, 0xcb, 0x38, 0xe2, 0xa5, 0xba, 0xb0, 0x57, 0xc1,
	0xed, 0x0e, 0xe6, 0xed, 0xa8, 0x23, 0xc4, 0x04, 0xf8, 0x4d, 0x89, 0x19, 0x87, 0x0f, 0xc1, 0x0d,
	0xad, 0x34, 0xcc, 0x51, 0x86, 0x6b, 0x46, 0xc3, 0x58, 0x9a, 0x09, 0x66, 0xb5, 0xf3, 0x15, 0xca,
	0xb0, 0xfd, 0xcd, 0x00, 0x0b, 0xdb, 0x25, 0x47, 0x1c, 0x2b, 0x34, 0xd3, 0xf0, 0x07, 0xc0, 0xea,
	0x95, 0x8c, 0x93, 0x0c, 0xd3, 0x30, 0x89, 0x14, 0x18, 0x68, 0xd7, 0xcb, 0x08, 0x06, 0x00, 0x90,
	0x02, 0x53, 0xf9, 0xc6, 0x9a, 0xd9, 0x98, 0x5a, 0xb2, 0x3c, 0xcf, 0x19, 0x57, 0x56, 0x47, 0xe5,
	0xd9, 0xd1, 0xd0, 0xe0, 0x02, 0x0b, 0x7c, 0x04, 0x6e, 0x16, 0x88, 0xf2, 0x04, 0xa5, 0xe1, 0x21,
	0x4a, 0xd2, 0x92, 0xe2, 0xda, 0x54, 0xc3, 0x58, 0xba, 0x1e, 0xcc, 0x29, 0xf7, 0xa6, 0xf4, 0x8a,
	0xc7, 0xf5, 0x51, 0x9a, 0x44, 0x88, 0xe3, 0x90, 0xe4, 0xe9, 0x71, 0xed, 0x6a, 0x15, 0x36, 0xab,
	0x9d, 0x3b, 0x79, 0x7a, 0x6c, 0x9f, 0x99, 0xe0, 0xd6, 0x68, 0x3a, 0xb8, 0x06, 0xac, 0xb2, 0xa8,
	0x70, 0xa2, 0xf2, 0x15, 0xce, 0xf2, 0xea, 0x5a, 0xb7, 0x2e, 0xbd, 0xb3, 0x29, 0x9a, 0xb3, 0x8d,
	0xd8, 0x51, 0x00, 0x64, 0xb8, 0xb0, 0xe1, 0x73, 0x30, 0xdd, 0xa3, 0x18, 0x71, 0x59, 0x4c, 0xcb,
	0x6b, 0x5e, 0xfa, 0xde, 0xc1, 0x90, 0xe8, 0x07, 0x6f, 0x5d, 0x09, 0x14, 0x56, 0xb0, 0x48, 0xce,
	0x9a, 0xf9, 0x3f, 0x2c, 0x12, 0x0b, 0x6b, 0x60, 0x9a, 0xe2, 0x8c, 0xf4, 0x65, 0x89, 0x66, 0xc4,
	0x8d, 0x3c, 0x6f, 0x58, 0x60, 0x66, 0x50, 0x53, 0xfb, 0x93, 0x01, 0xee, 0x8c, 0x76, 0x98, 0x15,
	0x24, 0x67, 0x18, 0x6e, 0x82, 0x85, 0x91, 0x6a, 0x87, 0x98, 0x52, 0x42, 0x2b, 0x42, 0xcb, 0x83,
	0x5a, 0x16, 0x2d, 0x7a, 0xce, 0x5e, 0x35, 0x6f, 0xc1, 0xfc, 0x70, 0x1f, 0x5e, 0x88, 0x70, 0xb8,
	0x03, 0xae, 0x51, 0xcc, 0xca, 0x94, 0xeb, 0x31, 0x78, 0x36, 0x7e, 0x0c, 0x86, 0x24, 0x05, 0x15,
	0x3a, 0xd0, 0x2c, 0xb6, 0x0f, 0xe6, 0xff, 0x72, 0xff, 0x4f, 0x13, 0xed, 0xfd, 0x34, 0xc1, 0x9c,
	0x82, 0xed, 0xc9, 0x64, 0xf0, 0xbd, 0x01, 0xc0, 0xf9, 0x7e, 0xc0, 0xe5, 0xf1, 0xea, 0xfe, 0xd8,
	0xa6, 0xfa, 0x04, 0x3d, 0xb2, 0xbd, 0xb3, 0xaf, 0xdf, 0xdf, 0x99, 0x8f, 0x61, 0x53, 0xfc, 0x16,
	0x27, 0x43, 0x92, 0xd7, 0xf5, 0x02, 0x31, 0xb7, 0xe9, 0x22, 0xd5, 0x10, 0xb7, 0x79, 0x0a, 0x3f,
	0x1b, 0x60, 0x6e, 0xb8, 0x4d, 0x70, 0x65, 0xc2, 0x2a, 0xea, 0xd5, 0xad, 0xaf, 0x4e, 0x0e, 0x94,
	0x13, 0x61, 0xaf, 0x56, 0xca, 0x3d, 0xfb, 0x89, 0x50, 0x7e, 0x2e, 0xf5, 0xe4, 0xc2, 0x4f, 0xb0,
	0xde, 0x3c, 0x1d, 0x08, 0xf7, 0xb3, 0x8a, 0xc6, 0x37, 0x9a, 0x1b, 0xbf, 0x0c, 0xb0, 0xd8, 0x23,
	0xd9, 0xd8, 0xcc, 0x1b, 0xf3, 0xc3, 0xcd, 0xd9, 0x15, 0x0b, 0xb7, 0x6b, 0xbc, 0xde, 0x52, 0xc0,
	0x98, 0xa4, 0x28, 0x8f, 0x1d, 0x42, 0x63, 0x37, 0xc6, 0x79, 0xb5, 0x8e, 0xfa, 0xf7, 0x2d, 0x12,
	0x76, 0xf9, 0x3f, 0xbf, 0xa6, 0x8d, 0x0f, 0xe6, 0x54, 0xa7, 0xdd, 0xfe, 0x68, 0x36, 0x3a, 0x92,
	0xb0, 0x1d, 0x31, 0x47, 0x9a, 0xc2, 0xda, 0x6f, 0x39, 0x2a, 0x31, 0xfb, 0xa2, 0x43, 0xba, 0xed,
	0x88, 0x75, 0x07, 0x21, 0xdd, 0xfd, 0x56, 0x57, 0x87, 0xfc, 0x30, 0x17, 0xa5, 0xdf, 0xf7, 0xdb,
	0x11, 0xf3, 0xfd, 0x41, 0x90, 0xef, 0xef, 0xb7, 0x7c, 0x5f, 0x87, 0x1d, 0x4c, 0x57, 0x3a, 0x97,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x5f, 0x6a, 0x64, 0x8e, 0x06, 0x00, 0x00,
}