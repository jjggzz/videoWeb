// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sms.proto

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

// 短信的策略
// 1验证码短信
// 2推广短信
// ...
type SmsStrategy int32

const (
	SmsStrategy_verify SmsStrategy = 0
)

var SmsStrategy_name = map[int32]string{
	0: "verify",
}

var SmsStrategy_value = map[string]int32{
	"verify": 0,
}

func (x SmsStrategy) String() string {
	return proto.EnumName(SmsStrategy_name, int32(x))
}

func (SmsStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{0}
}

type SendSmsRequest struct {
	// 策略
	Strategy SmsStrategy `protobuf:"varint,1,opt,name=strategy,proto3,enum=proto.SmsStrategy" json:"strategy,omitempty"`
	// 内容
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// 手机号码列表
	Phone []string `protobuf:"bytes,3,rep,name=phone,proto3" json:"phone,omitempty"`
}

func (m *SendSmsRequest) Reset()         { *m = SendSmsRequest{} }
func (m *SendSmsRequest) String() string { return proto.CompactTextString(m) }
func (*SendSmsRequest) ProtoMessage()    {}
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{0}
}
func (m *SendSmsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendSmsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendSmsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendSmsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSmsRequest.Merge(m, src)
}
func (m *SendSmsRequest) XXX_Size() int {
	return m.Size()
}
func (m *SendSmsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSmsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSmsRequest proto.InternalMessageInfo

func (m *SendSmsRequest) GetStrategy() SmsStrategy {
	if m != nil {
		return m.Strategy
	}
	return SmsStrategy_verify
}

func (m *SendSmsRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SendSmsRequest) GetPhone() []string {
	if m != nil {
		return m.Phone
	}
	return nil
}

type SendSmsResponse struct {
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *SendSmsResponse) Reset()         { *m = SendSmsResponse{} }
func (m *SendSmsResponse) String() string { return proto.CompactTextString(m) }
func (*SendSmsResponse) ProtoMessage()    {}
func (*SendSmsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8d8bdc537111860, []int{1}
}
func (m *SendSmsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendSmsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendSmsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendSmsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSmsResponse.Merge(m, src)
}
func (m *SendSmsResponse) XXX_Size() int {
	return m.Size()
}
func (m *SendSmsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSmsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendSmsResponse proto.InternalMessageInfo

func (m *SendSmsResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.SmsStrategy", SmsStrategy_name, SmsStrategy_value)
	proto.RegisterType((*SendSmsRequest)(nil), "proto.SendSmsRequest")
	proto.RegisterType((*SendSmsResponse)(nil), "proto.SendSmsResponse")
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor_c8d8bdc537111860) }

var fileDescriptor_c8d8bdc537111860 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xce, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x59, 0x5c, 0x7c, 0xc1, 0xa9, 0x79,
	0x29, 0xc1, 0xb9, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7a, 0x5c, 0x1c, 0xc5,
	0x25, 0x45, 0x89, 0x25, 0xa9, 0xe9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x42, 0x10,
	0x2d, 0x7a, 0xc1, 0xb9, 0xc5, 0xc1, 0x50, 0x99, 0x20, 0xb8, 0x1a, 0x21, 0x21, 0x2e, 0x96, 0xa4,
	0xfc, 0x94, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0xb5,
	0x20, 0x23, 0x3f, 0x2f, 0x55, 0x82, 0x59, 0x81, 0x59, 0x83, 0x33, 0x08, 0xc2, 0x51, 0x52, 0xe5,
	0xe2, 0x87, 0xdb, 0x55, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0xd2, 0x9c, 0x9c, 0x9f, 0x92, 0x0a,
	0xb6, 0x88, 0x39, 0x08, 0xcc, 0xd6, 0x92, 0xe4, 0xe2, 0x46, 0xb2, 0x49, 0x88, 0x8b, 0x8b, 0xad,
	0x2c, 0xb5, 0x28, 0x33, 0xad, 0x52, 0x80, 0xc1, 0xc8, 0x91, 0x8b, 0x39, 0x38, 0xb7, 0x58, 0xc8,
	0x8a, 0x8b, 0x1d, 0x6a, 0x90, 0x90, 0x28, 0xcc, 0x6d, 0x28, 0x9e, 0x90, 0x12, 0x43, 0x17, 0x86,
	0xd8, 0xa7, 0xc4, 0xe0, 0x24, 0x71, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49,
	0x6c, 0x60, 0x2d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xae, 0x6f, 0x26, 0x25, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SmsClient is the client API for Sms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsClient interface {
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
}

type smsClient struct {
	cc *grpc.ClientConn
}

func NewSmsClient(cc *grpc.ClientConn) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	out := new(SendSmsResponse)
	err := c.cc.Invoke(ctx, "/proto.Sms/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServer is the server API for Sms service.
type SmsServer interface {
	SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error)
}

// UnimplementedSmsServer can be embedded to have forward compatible implementations.
type UnimplementedSmsServer struct {
}

func (*UnimplementedSmsServer) SendSms(ctx context.Context, req *SendSmsRequest) (*SendSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}

func RegisterSmsServer(s *grpc.Server, srv SmsServer) {
	s.RegisterService(&_Sms_serviceDesc, srv)
}

func _Sms_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sms/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSms",
			Handler:    _Sms_SendSms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}

func (m *SendSmsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendSmsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Strategy != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSms(dAtA, i, uint64(m.Strategy))
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSms(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if len(m.Phone) > 0 {
		for _, s := range m.Phone {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SendSmsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendSmsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSms(dAtA, i, uint64(m.Code))
	}
	return i, nil
}

func encodeVarintSms(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SendSmsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Strategy != 0 {
		n += 1 + sovSms(uint64(m.Strategy))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovSms(uint64(l))
	}
	if len(m.Phone) > 0 {
		for _, s := range m.Phone {
			l = len(s)
			n += 1 + l + sovSms(uint64(l))
		}
	}
	return n
}

func (m *SendSmsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovSms(uint64(m.Code))
	}
	return n
}

func sovSms(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSms(x uint64) (n int) {
	return sovSms(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendSmsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSms
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
			return fmt.Errorf("proto: SendSmsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendSmsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Strategy", wireType)
			}
			m.Strategy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSms
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Strategy |= SmsStrategy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSms
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
				return ErrInvalidLengthSms
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSms
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSms
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
				return ErrInvalidLengthSms
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSms
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = append(m.Phone, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSms(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSms
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSms
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
func (m *SendSmsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSms
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
			return fmt.Errorf("proto: SendSmsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendSmsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSms
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
		default:
			iNdEx = preIndex
			skippy, err := skipSms(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSms
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSms
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
func skipSms(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSms
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
					return 0, ErrIntOverflowSms
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
					return 0, ErrIntOverflowSms
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
				return 0, ErrInvalidLengthSms
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSms
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSms
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
				next, err := skipSms(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSms
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
	ErrInvalidLengthSms = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSms   = fmt.Errorf("proto: integer overflow")
)
