// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: customer.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/jjggzz/truss/deftree/googlethirdparty"
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

type RegisterResult int32

const (
	RegisterResult_SUCCESS RegisterResult = 0
	RegisterResult_FAIL    RegisterResult = 1
)

var RegisterResult_name = map[int32]string{
	0: "SUCCESS",
	1: "FAIL",
}

var RegisterResult_value = map[string]int32{
	"SUCCESS": 0,
	"FAIL":    1,
}

func (x RegisterResult) String() string {
	return proto.EnumName(RegisterResult_name, int32(x))
}

func (RegisterResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

type RegisterByPhoneRequest struct {
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (m *RegisterByPhoneRequest) Reset()         { *m = RegisterByPhoneRequest{} }
func (m *RegisterByPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterByPhoneRequest) ProtoMessage()    {}
func (*RegisterByPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}
func (m *RegisterByPhoneRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterByPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterByPhoneRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterByPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterByPhoneRequest.Merge(m, src)
}
func (m *RegisterByPhoneRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterByPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterByPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterByPhoneRequest proto.InternalMessageInfo

func (m *RegisterByPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type RegisterByPhoneResponse struct {
	Result  RegisterResult `protobuf:"varint,1,opt,name=result,proto3,enum=proto.RegisterResult" json:"result,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *RegisterByPhoneResponse) Reset()         { *m = RegisterByPhoneResponse{} }
func (m *RegisterByPhoneResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterByPhoneResponse) ProtoMessage()    {}
func (*RegisterByPhoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}
func (m *RegisterByPhoneResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterByPhoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterByPhoneResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterByPhoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterByPhoneResponse.Merge(m, src)
}
func (m *RegisterByPhoneResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterByPhoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterByPhoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterByPhoneResponse proto.InternalMessageInfo

func (m *RegisterByPhoneResponse) GetResult() RegisterResult {
	if m != nil {
		return m.Result
	}
	return RegisterResult_SUCCESS
}

func (m *RegisterByPhoneResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.RegisterResult", RegisterResult_name, RegisterResult_value)
	proto.RegisterType((*RegisterByPhoneRequest)(nil), "proto.RegisterByPhoneRequest")
	proto.RegisterType((*RegisterByPhoneResponse)(nil), "proto.RegisterByPhoneResponse")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x33, 0x7f, 0xfe, 0xfd, 0x70, 0x94, 0x5a, 0x06, 0x3f, 0x42, 0xc1, 0x41, 0xba, 0x51,
	0x04, 0x13, 0xa8, 0x4f, 0x60, 0x8b, 0x82, 0xe0, 0x42, 0xa6, 0xb8, 0x15, 0xd2, 0xf6, 0x3a, 0x4d,
	0x69, 0xe6, 0xc6, 0x99, 0x9b, 0x45, 0xfb, 0x14, 0x3e, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0xf2, 0x22,
	0x62, 0xd2, 0x2e, 0x5a, 0x71, 0x35, 0xfc, 0xb8, 0xe7, 0xcc, 0x3d, 0xe7, 0xf2, 0xd6, 0x38, 0x73,
	0x84, 0x09, 0xd8, 0x20, 0xb5, 0x48, 0x28, 0x6a, 0xe5, 0xd3, 0xe9, 0xeb, 0x98, 0xa6, 0xd9, 0x28,
	0x18, 0x63, 0x12, 0xce, 0x66, 0x5a, 0x2f, 0x97, 0x21, 0xd9, 0xcc, 0xb9, 0x70, 0x02, 0xaf, 0x64,
	0x01, 0x42, 0x8d, 0xa8, 0xe7, 0x40, 0xd3, 0xd8, 0x4e, 0xd2, 0xc8, 0xd2, 0x22, 0x8c, 0x8c, 0x41,
	0x8a, 0x28, 0x46, 0xe3, 0xaa, 0xaf, 0x3a, 0x07, 0x63, 0x4c, 0x12, 0x34, 0x15, 0x75, 0x03, 0x7e,
	0xa2, 0x40, 0xc7, 0x8e, 0xc0, 0xf6, 0x17, 0x4f, 0x53, 0x34, 0xa0, 0xe0, 0x2d, 0x03, 0x47, 0xe2,
	0x88, 0xd7, 0xd2, 0x1f, 0xf6, 0xd9, 0x39, 0xbb, 0xdc, 0x53, 0x15, 0x74, 0x47, 0xfc, 0xf4, 0x97,
	0xde, 0xa5, 0x68, 0x1c, 0x88, 0x6b, 0x5e, 0xb7, 0xe0, 0xb2, 0x39, 0x95, 0x8e, 0x56, 0xef, 0xb8,
	0x5a, 0x11, 0x6c, 0xf4, 0xaa, 0x1c, 0xaa, 0xb5, 0x48, 0xf8, 0xbc, 0x91, 0x80, 0x73, 0x91, 0x06,
	0xff, 0x5f, 0xb9, 0x61, 0x83, 0x57, 0x17, 0xbc, 0xb5, 0xed, 0x11, 0xfb, 0xbc, 0x31, 0x7c, 0x1e,
	0x0c, 0xee, 0x86, 0xc3, 0xb6, 0x27, 0x9a, 0xfc, 0xff, 0xfd, 0xed, 0xc3, 0x63, 0x9b, 0xf5, 0x5e,
	0x78, 0x73, 0xb0, 0xbe, 0x93, 0x50, 0xfc, 0x70, 0x27, 0x98, 0x38, 0xdb, 0x09, 0xb0, 0x5d, 0xb0,
	0x23, 0xff, 0x1a, 0x57, 0x7d, 0xba, 0x5e, 0xdf, 0xff, 0xc8, 0x25, 0x5b, 0xe5, 0x92, 0x7d, 0xe5,
	0x92, 0xbd, 0x17, 0xd2, 0x5b, 0x15, 0xd2, 0xfb, 0x2c, 0xa4, 0x37, 0xaa, 0x97, 0xd6, 0x9b, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x02, 0xd1, 0x93, 0xa8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClient interface {
	RegisterByPhone(ctx context.Context, in *RegisterByPhoneRequest, opts ...grpc.CallOption) (*RegisterByPhoneResponse, error)
}

type customerClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClient(cc *grpc.ClientConn) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) RegisterByPhone(ctx context.Context, in *RegisterByPhoneRequest, opts ...grpc.CallOption) (*RegisterByPhoneResponse, error) {
	out := new(RegisterByPhoneResponse)
	err := c.cc.Invoke(ctx, "/proto.Customer/RegisterByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
type CustomerServer interface {
	RegisterByPhone(context.Context, *RegisterByPhoneRequest) (*RegisterByPhoneResponse, error)
}

// UnimplementedCustomerServer can be embedded to have forward compatible implementations.
type UnimplementedCustomerServer struct {
}

func (*UnimplementedCustomerServer) RegisterByPhone(ctx context.Context, req *RegisterByPhoneRequest) (*RegisterByPhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterByPhone not implemented")
}

func RegisterCustomerServer(s *grpc.Server, srv CustomerServer) {
	s.RegisterService(&_Customer_serviceDesc, srv)
}

func _Customer_RegisterByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).RegisterByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Customer/RegisterByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).RegisterByPhone(ctx, req.(*RegisterByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterByPhone",
			Handler:    _Customer_RegisterByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}

func (m *RegisterByPhoneRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterByPhoneRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Phone) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Phone)))
		i += copy(dAtA[i:], m.Phone)
	}
	return i, nil
}

func (m *RegisterByPhoneResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterByPhoneResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(m.Result))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCustomer(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintCustomer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterByPhoneRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Phone)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func (m *RegisterByPhoneResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovCustomer(uint64(m.Result))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCustomer(uint64(l))
	}
	return n
}

func sovCustomer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCustomer(x uint64) (n int) {
	return sovCustomer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterByPhoneRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: RegisterByPhoneRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterByPhoneRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func (m *RegisterByPhoneResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCustomer
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
			return fmt.Errorf("proto: RegisterByPhoneResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterByPhoneResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= RegisterResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCustomer
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
				return ErrInvalidLengthCustomer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCustomer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCustomer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCustomer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCustomer
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
func skipCustomer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCustomer
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
					return 0, ErrIntOverflowCustomer
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
					return 0, ErrIntOverflowCustomer
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
				return 0, ErrInvalidLengthCustomer
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCustomer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCustomer
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
				next, err := skipCustomer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCustomer
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
	ErrInvalidLengthCustomer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCustomer   = fmt.Errorf("proto: integer overflow")
)
