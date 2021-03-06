// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: generate.proto

// The package name determines the name of the directories that truss creates
// for `package proto;` truss will create the directory "generate-service".

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a150c98491dbd900, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Int64KeyResponse struct {
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Id   int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Int64KeyResponse) Reset()         { *m = Int64KeyResponse{} }
func (m *Int64KeyResponse) String() string { return proto.CompactTextString(m) }
func (*Int64KeyResponse) ProtoMessage()    {}
func (*Int64KeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a150c98491dbd900, []int{1}
}
func (m *Int64KeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Int64KeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Int64KeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Int64KeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64KeyResponse.Merge(m, src)
}
func (m *Int64KeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *Int64KeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64KeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_Int64KeyResponse proto.InternalMessageInfo

func (m *Int64KeyResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Int64KeyResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StringKeyResponse struct {
	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *StringKeyResponse) Reset()         { *m = StringKeyResponse{} }
func (m *StringKeyResponse) String() string { return proto.CompactTextString(m) }
func (*StringKeyResponse) ProtoMessage()    {}
func (*StringKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a150c98491dbd900, []int{2}
}
func (m *StringKeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StringKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StringKeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StringKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringKeyResponse.Merge(m, src)
}
func (m *StringKeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *StringKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StringKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StringKeyResponse proto.InternalMessageInfo

func (m *StringKeyResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StringKeyResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*Int64KeyResponse)(nil), "proto.Int64KeyResponse")
	proto.RegisterType((*StringKeyResponse)(nil), "proto.StringKeyResponse")
}

func init() { proto.RegisterFile("generate.proto", fileDescriptor_a150c98491dbd900) }

var fileDescriptor_a150c98491dbd900 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcd, 0x4b,
	0x2d, 0x4a, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xec,
	0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x4a, 0x66, 0x5c, 0x02, 0x9e, 0x79, 0x25, 0x66, 0x26,
	0xde, 0xa9, 0x95, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0xc9,
	0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x10, 0x1f, 0x17, 0x53,
	0x66, 0x8a, 0x04, 0x13, 0x58, 0x84, 0x29, 0x33, 0x45, 0xc9, 0x9c, 0x4b, 0x30, 0xb8, 0xa4, 0x28,
	0x33, 0x2f, 0x9d, 0x78, 0x8d, 0x9c, 0x20, 0x8d, 0x46, 0x6d, 0x8c, 0x5c, 0x1c, 0xee, 0x50, 0x37,
	0x09, 0x59, 0x73, 0x09, 0xc0, 0xd8, 0x30, 0x57, 0x08, 0xf1, 0x40, 0x5c, 0xaa, 0x07, 0x76, 0x9f,
	0x94, 0x38, 0x94, 0x87, 0xee, 0x48, 0x25, 0x06, 0x21, 0x5b, 0x2e, 0x41, 0x98, 0x66, 0xb8, 0x53,
	0xd0, 0x74, 0x4b, 0x40, 0x79, 0x18, 0x4e, 0x55, 0x62, 0x70, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e,
	0xe1, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x5e, 0xb5, 0x27, 0x3c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GenerateClient is the client API for Generate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenerateClient interface {
	GenerateInt64Key(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Int64KeyResponse, error)
	GenerateStringKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringKeyResponse, error)
}

type generateClient struct {
	cc *grpc.ClientConn
}

func NewGenerateClient(cc *grpc.ClientConn) GenerateClient {
	return &generateClient{cc}
}

func (c *generateClient) GenerateInt64Key(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Int64KeyResponse, error) {
	out := new(Int64KeyResponse)
	err := c.cc.Invoke(ctx, "/proto.Generate/GenerateInt64Key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateClient) GenerateStringKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringKeyResponse, error) {
	out := new(StringKeyResponse)
	err := c.cc.Invoke(ctx, "/proto.Generate/GenerateStringKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateServer is the server API for Generate service.
type GenerateServer interface {
	GenerateInt64Key(context.Context, *Empty) (*Int64KeyResponse, error)
	GenerateStringKey(context.Context, *Empty) (*StringKeyResponse, error)
}

// UnimplementedGenerateServer can be embedded to have forward compatible implementations.
type UnimplementedGenerateServer struct {
}

func (*UnimplementedGenerateServer) GenerateInt64Key(ctx context.Context, req *Empty) (*Int64KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInt64Key not implemented")
}
func (*UnimplementedGenerateServer) GenerateStringKey(ctx context.Context, req *Empty) (*StringKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateStringKey not implemented")
}

func RegisterGenerateServer(s *grpc.Server, srv GenerateServer) {
	s.RegisterService(&_Generate_serviceDesc, srv)
}

func _Generate_GenerateInt64Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateServer).GenerateInt64Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Generate/GenerateInt64Key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateServer).GenerateInt64Key(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generate_GenerateStringKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateServer).GenerateStringKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Generate/GenerateStringKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateServer).GenerateStringKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Generate",
	HandlerType: (*GenerateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInt64Key",
			Handler:    _Generate_GenerateInt64Key_Handler,
		},
		{
			MethodName: "GenerateStringKey",
			Handler:    _Generate_GenerateStringKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate.proto",
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *Int64KeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Int64KeyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerate(dAtA, i, uint64(m.Code))
	}
	if m.Id != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGenerate(dAtA, i, uint64(m.Id))
	}
	return i, nil
}

func (m *StringKeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StringKeyResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGenerate(dAtA, i, uint64(m.Code))
	}
	if len(m.Id) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGenerate(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	return i, nil
}

func encodeVarintGenerate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Int64KeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGenerate(uint64(m.Code))
	}
	if m.Id != 0 {
		n += 1 + sovGenerate(uint64(m.Id))
	}
	return n
}

func (m *StringKeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGenerate(uint64(m.Code))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGenerate(uint64(l))
	}
	return n
}

func sovGenerate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerate(x uint64) (n int) {
	return sovGenerate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerate
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGenerate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerate
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
func (m *Int64KeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerate
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
			return fmt.Errorf("proto: Int64KeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Int64KeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerate
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
func (m *StringKeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerate
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
			return fmt.Errorf("proto: StringKeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StringKeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerate
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerate
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
func skipGenerate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerate
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
					return 0, ErrIntOverflowGenerate
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerate
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
				return 0, ErrInvalidLengthGenerate
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerate
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerate(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerate
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerate   = fmt.Errorf("proto: integer overflow")
)
