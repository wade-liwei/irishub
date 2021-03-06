// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: guardian/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryProfilersRequest is request type for the Query/Profilers RPC method
type QueryProfilersRequest struct {
}

func (m *QueryProfilersRequest) Reset()         { *m = QueryProfilersRequest{} }
func (m *QueryProfilersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProfilersRequest) ProtoMessage()    {}
func (*QueryProfilersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cf24f8e5be2110, []int{0}
}
func (m *QueryProfilersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfilersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfilersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfilersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfilersRequest.Merge(m, src)
}
func (m *QueryProfilersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfilersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfilersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfilersRequest proto.InternalMessageInfo

// QueryProfilersResponse is response type for the Query/Profilers RPC method
type QueryProfilersResponse struct {
	Profilers []Guardian `protobuf:"bytes,1,rep,name=profilers,proto3" json:"profilers"`
}

func (m *QueryProfilersResponse) Reset()         { *m = QueryProfilersResponse{} }
func (m *QueryProfilersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProfilersResponse) ProtoMessage()    {}
func (*QueryProfilersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cf24f8e5be2110, []int{1}
}
func (m *QueryProfilersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfilersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfilersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfilersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfilersResponse.Merge(m, src)
}
func (m *QueryProfilersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfilersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfilersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfilersResponse proto.InternalMessageInfo

func (m *QueryProfilersResponse) GetProfilers() []Guardian {
	if m != nil {
		return m.Profilers
	}
	return nil
}

// QueryTrusteesRequest is request type for the Query/Trustees RPC method
type QueryTrusteesRequest struct {
}

func (m *QueryTrusteesRequest) Reset()         { *m = QueryTrusteesRequest{} }
func (m *QueryTrusteesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTrusteesRequest) ProtoMessage()    {}
func (*QueryTrusteesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cf24f8e5be2110, []int{2}
}
func (m *QueryTrusteesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTrusteesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTrusteesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTrusteesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTrusteesRequest.Merge(m, src)
}
func (m *QueryTrusteesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTrusteesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTrusteesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTrusteesRequest proto.InternalMessageInfo

// QueryTrusteesResponse is response type for the Query/Trustees RPC method
type QueryTrusteesResponse struct {
	Trustees []Guardian `protobuf:"bytes,1,rep,name=trustees,proto3" json:"trustees"`
}

func (m *QueryTrusteesResponse) Reset()         { *m = QueryTrusteesResponse{} }
func (m *QueryTrusteesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTrusteesResponse) ProtoMessage()    {}
func (*QueryTrusteesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cf24f8e5be2110, []int{3}
}
func (m *QueryTrusteesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTrusteesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTrusteesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTrusteesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTrusteesResponse.Merge(m, src)
}
func (m *QueryTrusteesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTrusteesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTrusteesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTrusteesResponse proto.InternalMessageInfo

func (m *QueryTrusteesResponse) GetTrustees() []Guardian {
	if m != nil {
		return m.Trustees
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryProfilersRequest)(nil), "irishub.guardian.QueryProfilersRequest")
	proto.RegisterType((*QueryProfilersResponse)(nil), "irishub.guardian.QueryProfilersResponse")
	proto.RegisterType((*QueryTrusteesRequest)(nil), "irishub.guardian.QueryTrusteesRequest")
	proto.RegisterType((*QueryTrusteesResponse)(nil), "irishub.guardian.QueryTrusteesResponse")
}

func init() { proto.RegisterFile("guardian/query.proto", fileDescriptor_20cf24f8e5be2110) }

var fileDescriptor_20cf24f8e5be2110 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2f, 0x4d, 0x2c,
	0x4a, 0xc9, 0x4c, 0xcc, 0xd3, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0xc8, 0x2c, 0xca, 0x2c, 0xce, 0x28, 0x4d, 0xd2, 0x83, 0xc9, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x25, 0xf5, 0x41, 0x2c, 0x88, 0x3a, 0x29, 0x71, 0xb8, 0x6e, 0x18, 0x03, 0x2a,
	0x21, 0x93, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97,
	0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55, 0x12, 0xe7, 0x12, 0x0d, 0x04, 0xd9,
	0x16, 0x50, 0x94, 0x9f, 0x96, 0x99, 0x93, 0x5a, 0x54, 0x1c, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0xa2, 0x14, 0xc1, 0x25, 0x86, 0x2e, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xc7, 0xc5,
	0x59, 0x00, 0x13, 0x94, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd2, 0x43, 0x77, 0xa5, 0x9e,
	0x3b, 0x94, 0xe1, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x42, 0x8b, 0x92, 0x18, 0x97, 0x08,
	0xd8, 0xe4, 0x90, 0xa2, 0xd2, 0xe2, 0x92, 0xd4, 0x54, 0xb8, 0x8d, 0xa1, 0x50, 0xa7, 0x20, 0xc4,
	0xa1, 0x16, 0xda, 0x70, 0x71, 0x94, 0x40, 0xc5, 0x88, 0xb6, 0x0f, 0xae, 0xc3, 0xa8, 0x8f, 0x89,
	0x8b, 0x15, 0x6c, 0xae, 0x50, 0x33, 0x23, 0x17, 0x27, 0xdc, 0x3b, 0x42, 0xea, 0x98, 0x66, 0x60,
	0x0d, 0x09, 0x29, 0x0d, 0xc2, 0x0a, 0x21, 0x0e, 0x55, 0x52, 0x6e, 0xba, 0xfc, 0x64, 0x32, 0x93,
	0xac, 0x90, 0xb4, 0x3e, 0x54, 0x07, 0x3c, 0x2e, 0xf4, 0xe1, 0xde, 0x17, 0xaa, 0xe7, 0xe2, 0x80,
	0xf9, 0x50, 0x48, 0x0d, 0x87, 0xd1, 0x68, 0x41, 0x23, 0xa5, 0x4e, 0x50, 0x1d, 0xd4, 0x05, 0x4a,
	0x60, 0x17, 0xc8, 0x08, 0x49, 0x61, 0xba, 0x00, 0x16, 0x20, 0x4e, 0xde, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x98, 0x9e, 0x59, 0x02, 0xb2, 0x24, 0x39, 0x3f, 0x17,
	0xac, 0x3f, 0x2f, 0xb5, 0x04, 0x6e, 0x4e, 0x6e, 0x7e, 0x4a, 0x69, 0x4e, 0x6a, 0x31, 0x92, 0x79,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x64, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0x97, 0x78, 0x65, 0xbd, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Profilers returns all profilers
	Profilers(ctx context.Context, in *QueryProfilersRequest, opts ...grpc.CallOption) (*QueryProfilersResponse, error)
	// Trustees returns all trustees
	Trustees(ctx context.Context, in *QueryTrusteesRequest, opts ...grpc.CallOption) (*QueryTrusteesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Profilers(ctx context.Context, in *QueryProfilersRequest, opts ...grpc.CallOption) (*QueryProfilersResponse, error) {
	out := new(QueryProfilersResponse)
	err := c.cc.Invoke(ctx, "/irishub.guardian.Query/Profilers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Trustees(ctx context.Context, in *QueryTrusteesRequest, opts ...grpc.CallOption) (*QueryTrusteesResponse, error) {
	out := new(QueryTrusteesResponse)
	err := c.cc.Invoke(ctx, "/irishub.guardian.Query/Trustees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Profilers returns all profilers
	Profilers(context.Context, *QueryProfilersRequest) (*QueryProfilersResponse, error)
	// Trustees returns all trustees
	Trustees(context.Context, *QueryTrusteesRequest) (*QueryTrusteesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Profilers(ctx context.Context, req *QueryProfilersRequest) (*QueryProfilersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profilers not implemented")
}
func (*UnimplementedQueryServer) Trustees(ctx context.Context, req *QueryTrusteesRequest) (*QueryTrusteesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trustees not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Profilers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProfilersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Profilers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irishub.guardian.Query/Profilers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Profilers(ctx, req.(*QueryProfilersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Trustees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTrusteesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Trustees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irishub.guardian.Query/Trustees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Trustees(ctx, req.(*QueryTrusteesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irishub.guardian.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Profilers",
			Handler:    _Query_Profilers_Handler,
		},
		{
			MethodName: "Trustees",
			Handler:    _Query_Trustees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guardian/query.proto",
}

func (m *QueryProfilersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfilersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfilersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryProfilersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfilersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfilersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Profilers) > 0 {
		for iNdEx := len(m.Profilers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Profilers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryTrusteesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTrusteesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTrusteesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTrusteesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTrusteesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTrusteesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Trustees) > 0 {
		for iNdEx := len(m.Trustees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Trustees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProfilersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryProfilersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Profilers) > 0 {
		for _, e := range m.Profilers {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryTrusteesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTrusteesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Trustees) > 0 {
		for _, e := range m.Trustees {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProfilersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProfilersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfilersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryProfilersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryProfilersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfilersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profilers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profilers = append(m.Profilers, Guardian{})
			if err := m.Profilers[len(m.Profilers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTrusteesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTrusteesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTrusteesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryTrusteesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryTrusteesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTrusteesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trustees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trustees = append(m.Trustees, Guardian{})
			if err := m.Trustees[len(m.Trustees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
