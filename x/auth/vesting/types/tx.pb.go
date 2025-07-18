// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vesting/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// MsgCreatePeriodicVestingAccount defines a message that enables creating a periodic vesting
// account.
type MsgCreatePeriodicVestingAccount struct {
	FromAddress    string         `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress      string         `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	StartTime      int64          `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods []types.Period `protobuf:"bytes,4,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods"`
}

func (m *MsgCreatePeriodicVestingAccount) Reset()         { *m = MsgCreatePeriodicVestingAccount{} }
func (m *MsgCreatePeriodicVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeriodicVestingAccount) ProtoMessage()    {}
func (*MsgCreatePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66d8e2aee35e411, []int{0}
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeriodicVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeriodicVestingAccount.Merge(m, src)
}
func (m *MsgCreatePeriodicVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeriodicVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeriodicVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeriodicVestingAccount proto.InternalMessageInfo

func (m *MsgCreatePeriodicVestingAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgCreatePeriodicVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreatePeriodicVestingAccount) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *MsgCreatePeriodicVestingAccount) GetVestingPeriods() []types.Period {
	if m != nil {
		return m.VestingPeriods
	}
	return nil
}

// MsgCreateVestingAccountResponse defines the Msg/CreatePeriodicVestingAccount
// response type.
type MsgCreatePeriodicVestingAccountResponse struct {
}

func (m *MsgCreatePeriodicVestingAccountResponse) Reset() {
	*m = MsgCreatePeriodicVestingAccountResponse{}
}
func (m *MsgCreatePeriodicVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeriodicVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreatePeriodicVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66d8e2aee35e411, []int{1}
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeriodicVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeriodicVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePeriodicVestingAccount)(nil), "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount")
	proto.RegisterType((*MsgCreatePeriodicVestingAccountResponse)(nil), "cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse")
}

func init() { proto.RegisterFile("vesting/v1beta1/tx.proto", fileDescriptor_c66d8e2aee35e411) }

var fileDescriptor_c66d8e2aee35e411 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0x73, 0xff, 0x96, 0x3f, 0xe4, 0x2a, 0x0a, 0x41, 0x24, 0x16, 0xbd, 0xd4, 0x22, 0x58,
	0x41, 0x13, 0x5a, 0x07, 0xc1, 0x45, 0x5a, 0xd1, 0xad, 0x20, 0x41, 0x1c, 0x5c, 0xc2, 0x25, 0x39,
	0xd3, 0x0c, 0xc9, 0x85, 0xdc, 0x9b, 0x5a, 0xbf, 0x85, 0xa3, 0xa3, 0x93, 0x9f, 0xa5, 0x63, 0x37,
	0x9d, 0x44, 0xda, 0xc5, 0x8f, 0x21, 0xcd, 0xa5, 0x45, 0xc4, 0x5a, 0x70, 0x4a, 0x78, 0x9e, 0xdf,
	0xfb, 0xbe, 0xcf, 0xbd, 0x77, 0x58, 0xef, 0x33, 0x01, 0x61, 0x1c, 0x58, 0xfd, 0xa6, 0xcb, 0x80,
	0x36, 0x2d, 0x18, 0x98, 0x49, 0xca, 0x81, 0x6b, 0x1b, 0x1e, 0x17, 0x11, 0x17, 0x66, 0x01, 0x98,
	0x05, 0x50, 0x5d, 0x0f, 0x78, 0xc0, 0x73, 0xc4, 0x9a, 0xfe, 0x49, 0xba, 0x4a, 0x24, 0x6d, 0xb9,
	0x54, 0xb0, 0x79, 0x2f, 0x8f, 0x87, 0x71, 0xe1, 0x6f, 0x4a, 0xdf, 0x91, 0x85, 0x45, 0x6b, 0x69,
	0xed, 0x16, 0xa5, 0xdf, 0x93, 0xcc, 0x06, 0x4b, 0xca, 0x58, 0x40, 0xcd, 0xf2, 0xd6, 0x5f, 0x10,
	0x36, 0xba, 0x22, 0x38, 0x4b, 0x19, 0x05, 0x76, 0xc9, 0xd2, 0x90, 0xfb, 0xa1, 0x77, 0x2d, 0xe9,
	0xb6, 0xe7, 0xf1, 0x2c, 0x06, 0x6d, 0x07, 0xaf, 0xdc, 0xa6, 0x3c, 0x72, 0xa8, 0xef, 0xa7, 0x4c,
	0x08, 0x1d, 0xd5, 0x50, 0x43, 0xb5, 0x2b, 0x53, 0xad, 0x2d, 0x25, 0x6d, 0x1b, 0x63, 0xe0, 0x73,
	0xe0, 0x5f, 0x0e, 0xa8, 0xc0, 0xbf, 0xd8, 0x02, 0x68, 0x0a, 0x0e, 0x84, 0x11, 0xd3, 0x4b, 0x35,
	0xd4, 0x28, 0xd9, 0x6a, 0xae, 0x5c, 0x85, 0x11, 0xd3, 0xba, 0x78, 0xad, 0x08, 0xe8, 0x24, 0x79,
	0x04, 0xa1, 0x97, 0x6b, 0xa5, 0x46, 0xa5, 0x45, 0xcc, 0x9f, 0xd7, 0x69, 0xca, 0xa4, 0x9d, 0xf2,
	0xf0, 0xcd, 0x50, 0xec, 0xd5, 0xc2, 0x95, 0xa2, 0x38, 0x29, 0x7f, 0x3c, 0x19, 0x4a, 0x7d, 0x1f,
	0xef, 0x2d, 0x39, 0x98, 0xcd, 0x44, 0xc2, 0x63, 0xc1, 0x5a, 0xcf, 0x08, 0xab, 0x5d, 0x11, 0x9c,
	0x0f, 0x80, 0xc5, 0xbe, 0xf6, 0x88, 0xf0, 0xd6, 0xaf, 0xfb, 0x38, 0x5e, 0x94, 0x6a, 0xc9, 0xbc,
	0xea, 0xe9, 0x1f, 0x0b, 0x67, 0x41, 0x3b, 0x17, 0xc3, 0x31, 0x41, 0xa3, 0x31, 0x41, 0xef, 0x63,
	0x82, 0x1e, 0x26, 0x44, 0x19, 0x4d, 0x88, 0xf2, 0x3a, 0x21, 0xca, 0xcd, 0x41, 0x10, 0x42, 0x2f,
	0x73, 0x4d, 0x8f, 0x47, 0x16, 0xcd, 0x52, 0x7a, 0x18, 0xdf, 0xe5, 0x5f, 0x6b, 0x60, 0xd1, 0x0c,
	0x7a, 0xf3, 0x07, 0x00, 0xf7, 0x09, 0x13, 0xee, 0xff, 0xfc, 0xf2, 0x8f, 0x3e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x9c, 0xda, 0xe4, 0x15, 0xc8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgExtendClient is the client API for MsgExtend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgExtendClient interface {
	// anam move from cosmos-sdk 0.46
	CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error)
}

type msgExtendClient struct {
	cc grpc1.ClientConn
}

func NewMsgExtendClient(cc grpc1.ClientConn) MsgExtendClient {
	return &msgExtendClient{cc}
}

func (c *msgExtendClient) CreatePeriodicVestingAccount(ctx context.Context, in *MsgCreatePeriodicVestingAccount, opts ...grpc.CallOption) (*MsgCreatePeriodicVestingAccountResponse, error) {
	out := new(MsgCreatePeriodicVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.MsgExtend/CreatePeriodicVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgExtendServer is the server API for MsgExtend service.
type MsgExtendServer interface {
	// anam move from cosmos-sdk 0.46
	CreatePeriodicVestingAccount(context.Context, *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error)
}

// UnimplementedMsgExtendServer can be embedded to have forward compatible implementations.
type UnimplementedMsgExtendServer struct {
}

func (*UnimplementedMsgExtendServer) CreatePeriodicVestingAccount(ctx context.Context, req *MsgCreatePeriodicVestingAccount) (*MsgCreatePeriodicVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeriodicVestingAccount not implemented")
}

func RegisterMsgExtendServer(s grpc1.Server, srv MsgExtendServer) {
	s.RegisterService(&_MsgExtend_serviceDesc, srv)
}

func _MsgExtend_CreatePeriodicVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeriodicVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgExtendServer).CreatePeriodicVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.MsgExtend/CreatePeriodicVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgExtendServer).CreatePeriodicVestingAccount(ctx, req.(*MsgCreatePeriodicVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgExtend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vesting.v1beta1.MsgExtend",
	HandlerType: (*MsgExtendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePeriodicVestingAccount",
			Handler:    _MsgExtend_CreatePeriodicVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vesting/v1beta1/tx.proto",
}

func (m *MsgCreatePeriodicVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StartTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePeriodicVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeriodicVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeriodicVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovTx(uint64(m.StartTime))
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreatePeriodicVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePeriodicVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, types.Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePeriodicVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeriodicVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
