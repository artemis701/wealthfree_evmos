// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/erc20/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// MsgConvertCoin defines a Msg to convert a native Cosmos coin to a ERC20 token
type MsgConvertCoin struct {
	// coin is a Cosmos coin whose denomination is registered in a token pair. The coin
	// amount defines the amount of coins to convert.
	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin"`
	// receiver is the hex address to receive ERC20 token
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender is the cosmos bech32 address from the owner of the given Cosmos coins
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertCoin) Reset()         { *m = MsgConvertCoin{} }
func (m *MsgConvertCoin) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoin) ProtoMessage()    {}
func (*MsgConvertCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8926fc6cb676914, []int{0}
}
func (m *MsgConvertCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoin.Merge(m, src)
}
func (m *MsgConvertCoin) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoin.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoin proto.InternalMessageInfo

func (m *MsgConvertCoin) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *MsgConvertCoin) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertCoin) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertCoinResponse returns no fields
type MsgConvertCoinResponse struct {
}

func (m *MsgConvertCoinResponse) Reset()         { *m = MsgConvertCoinResponse{} }
func (m *MsgConvertCoinResponse) String() string { return proto.CompactTextString(m) }
func (*MsgConvertCoinResponse) ProtoMessage()    {}
func (*MsgConvertCoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8926fc6cb676914, []int{1}
}
func (m *MsgConvertCoinResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertCoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertCoinResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertCoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertCoinResponse.Merge(m, src)
}
func (m *MsgConvertCoinResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertCoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertCoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertCoinResponse proto.InternalMessageInfo

// MsgConvertERC20 defines a Msg to convert a ERC20 token to a native Cosmos
// coin.
type MsgConvertERC20 struct {
	// contract_address of an ERC20 token contract, that is registered in a token pair
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// amount of ERC20 tokens to convert
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// receiver is the bech32 address to receive native Cosmos coins
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// sender is the hex address from the owner of the given ERC20 tokens
	Sender string `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgConvertERC20) Reset()         { *m = MsgConvertERC20{} }
func (m *MsgConvertERC20) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20) ProtoMessage()    {}
func (*MsgConvertERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8926fc6cb676914, []int{2}
}
func (m *MsgConvertERC20) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20.Merge(m, src)
}
func (m *MsgConvertERC20) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20 proto.InternalMessageInfo

func (m *MsgConvertERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *MsgConvertERC20) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgConvertERC20) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgConvertERC20Response returns no fields
type MsgConvertERC20Response struct {
}

func (m *MsgConvertERC20Response) Reset()         { *m = MsgConvertERC20Response{} }
func (m *MsgConvertERC20Response) String() string { return proto.CompactTextString(m) }
func (*MsgConvertERC20Response) ProtoMessage()    {}
func (*MsgConvertERC20Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8926fc6cb676914, []int{3}
}
func (m *MsgConvertERC20Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgConvertERC20Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgConvertERC20Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgConvertERC20Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgConvertERC20Response.Merge(m, src)
}
func (m *MsgConvertERC20Response) XXX_Size() int {
	return m.Size()
}
func (m *MsgConvertERC20Response) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgConvertERC20Response.DiscardUnknown(m)
}

var xxx_messageInfo_MsgConvertERC20Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgConvertCoin)(nil), "evmos.erc20.v1.MsgConvertCoin")
	proto.RegisterType((*MsgConvertCoinResponse)(nil), "evmos.erc20.v1.MsgConvertCoinResponse")
	proto.RegisterType((*MsgConvertERC20)(nil), "evmos.erc20.v1.MsgConvertERC20")
	proto.RegisterType((*MsgConvertERC20Response)(nil), "evmos.erc20.v1.MsgConvertERC20Response")
}

func init() { proto.RegisterFile("evmos/erc20/v1/tx.proto", fileDescriptor_f8926fc6cb676914) }

var fileDescriptor_f8926fc6cb676914 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0x8e, 0x93, 0x28, 0xa2, 0x2e, 0x6a, 0x91, 0x85, 0xda, 0xf4, 0x84, 0x9c, 0x90, 0xa1, 0x0d,
	0x03, 0x76, 0x2e, 0xfd, 0x05, 0x24, 0x02, 0x89, 0xa1, 0xcb, 0x8d, 0x2c, 0x95, 0xcf, 0x67, 0x1d,
	0x27, 0x88, 0xdf, 0xc9, 0x76, 0x4f, 0xed, 0xc2, 0xd0, 0x91, 0x09, 0x89, 0x3f, 0xc3, 0x4f, 0xe8,
	0x58, 0x89, 0x05, 0x31, 0x54, 0x28, 0xe1, 0x87, 0xa0, 0xf3, 0x5d, 0x43, 0x0f, 0xd4, 0x74, 0x49,
	0xec, 0xf7, 0x7d, 0xf7, 0xbe, 0xef, 0x7d, 0xcf, 0x78, 0x5f, 0x15, 0x0b, 0xb0, 0x5c, 0x19, 0x39,
	0x9d, 0xf0, 0x22, 0xe4, 0xee, 0x9c, 0xe5, 0x06, 0x1c, 0x90, 0x1d, 0x0f, 0x30, 0x0f, 0xb0, 0x22,
	0x0c, 0xa8, 0x04, 0x5b, 0x32, 0x63, 0x61, 0x15, 0x2f, 0xc2, 0x58, 0x39, 0x11, 0x72, 0x09, 0x99,
	0xae, 0xf8, 0xc1, 0xd3, 0x14, 0x52, 0xf0, 0x47, 0x5e, 0x9e, 0xea, 0xea, 0xb3, 0x14, 0x20, 0xfd,
	0xa8, 0xb8, 0xc8, 0x33, 0x2e, 0xb4, 0x06, 0x27, 0x5c, 0x06, 0xda, 0x56, 0xe8, 0xe8, 0x02, 0xef,
	0x9c, 0xd8, 0x74, 0x0e, 0xba, 0x50, 0xc6, 0xcd, 0x21, 0xd3, 0xe4, 0x18, 0x77, 0xcb, 0x9e, 0x7d,
	0x34, 0x44, 0xe3, 0xed, 0xe9, 0x01, 0xab, 0x44, 0x59, 0x29, 0xca, 0x6a, 0x51, 0x56, 0x12, 0x67,
	0xdd, 0xab, 0x9b, 0x41, 0x2b, 0xf2, 0x64, 0x12, 0xe0, 0x47, 0x46, 0x49, 0x95, 0x15, 0xca, 0xf4,
	0xdb, 0x43, 0x34, 0xde, 0x8a, 0xd6, 0x77, 0xb2, 0x87, 0x7b, 0x56, 0xe9, 0x44, 0x99, 0x7e, 0xc7,
	0x23, 0xf5, 0x6d, 0xd4, 0xc7, 0x7b, 0x4d, 0xe9, 0x48, 0xd9, 0x1c, 0xb4, 0x55, 0xa3, 0x6f, 0x08,
	0xef, 0xfe, 0x85, 0x5e, 0x47, 0xf3, 0xe9, 0x84, 0xbc, 0xc0, 0x4f, 0x24, 0x68, 0x67, 0x84, 0x74,
	0xa7, 0x22, 0x49, 0x8c, 0xb2, 0xd6, 0x5b, 0xdc, 0x8a, 0x76, 0x6f, 0xeb, 0xaf, 0xaa, 0x32, 0x79,
	0x83, 0x7b, 0x62, 0x01, 0x67, 0xda, 0x55, 0x56, 0x66, 0xac, 0x34, 0xfa, 0xf3, 0x66, 0x70, 0x98,
	0x66, 0xee, 0xfd, 0x59, 0xcc, 0x24, 0x2c, 0x78, 0x1d, 0x65, 0xf5, 0xf7, 0xd2, 0x26, 0x1f, 0xb8,
	0xbb, 0xc8, 0x95, 0x65, 0x6f, 0xb5, 0x8b, 0xea, 0xaf, 0x1b, 0x43, 0x75, 0xee, 0x1d, 0xaa, 0xdb,
	0x18, 0xea, 0x00, 0xef, 0xff, 0xe3, 0xfc, 0x76, 0xaa, 0xe9, 0xe7, 0x36, 0xee, 0x9c, 0xd8, 0x94,
	0x7c, 0xc2, 0xdb, 0x77, 0xf3, 0xa6, 0xac, 0xb9, 0x66, 0xd6, 0x0c, 0x25, 0x38, 0xdc, 0x8c, 0xaf,
	0x43, 0x3b, 0xba, 0xfc, 0xfe, 0xfb, 0x6b, 0xfb, 0x39, 0x19, 0xf0, 0xff, 0xde, 0x13, 0x97, 0x15,
	0xff, 0xd4, 0xef, 0xea, 0x12, 0xe1, 0xc7, 0x8d, 0x68, 0x07, 0xf7, 0x2b, 0x78, 0x42, 0x70, 0xf4,
	0x00, 0x61, 0xed, 0x61, 0xec, 0x3d, 0x8c, 0xc8, 0x70, 0x83, 0x07, 0x5f, 0x9b, 0xcd, 0xae, 0x96,
	0x14, 0x5d, 0x2f, 0x29, 0xfa, 0xb5, 0xa4, 0xe8, 0xcb, 0x8a, 0xb6, 0xae, 0x57, 0xb4, 0xf5, 0x63,
	0x45, 0x5b, 0xef, 0xc6, 0x77, 0xb6, 0x54, 0x77, 0xf1, 0xbf, 0x45, 0x38, 0xe1, 0xe7, 0x75, 0x47,
	0xbf, 0xab, 0xb8, 0xe7, 0x9f, 0xf0, 0xf1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x52, 0x11,
	0x3e, 0x41, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ConvertCoin(ctx context.Context, in *MsgConvertCoin, opts ...grpc.CallOption) (*MsgConvertCoinResponse, error) {
	out := new(MsgConvertCoinResponse)
	err := c.cc.Invoke(ctx, "/evmos.erc20.v1.Msg/ConvertCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ConvertERC20(ctx context.Context, in *MsgConvertERC20, opts ...grpc.CallOption) (*MsgConvertERC20Response, error) {
	out := new(MsgConvertERC20Response)
	err := c.cc.Invoke(ctx, "/evmos.erc20.v1.Msg/ConvertERC20", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConvertCoin mints a ERC20 representation of the native Cosmos coin denom
	// that is registered on the token mapping.
	ConvertCoin(context.Context, *MsgConvertCoin) (*MsgConvertCoinResponse, error)
	// ConvertERC20 mints a native Cosmos coin representation of the ERC20 token
	// contract that is registered on the token mapping.
	ConvertERC20(context.Context, *MsgConvertERC20) (*MsgConvertERC20Response, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ConvertCoin(ctx context.Context, req *MsgConvertCoin) (*MsgConvertCoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertCoin not implemented")
}
func (*UnimplementedMsgServer) ConvertERC20(ctx context.Context, req *MsgConvertERC20) (*MsgConvertERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertERC20 not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ConvertCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertCoin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.erc20.v1.Msg/ConvertCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertCoin(ctx, req.(*MsgConvertCoin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ConvertERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgConvertERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ConvertERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evmos.erc20.v1.Msg/ConvertERC20",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ConvertERC20(ctx, req.(*MsgConvertERC20))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evmos.erc20.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertCoin",
			Handler:    _Msg_ConvertCoin_Handler,
		},
		{
			MethodName: "ConvertERC20",
			Handler:    _Msg_ConvertERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evmos/erc20/v1/tx.proto",
}

func (m *MsgConvertCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgConvertCoinResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertCoinResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertCoinResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgConvertERC20Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgConvertERC20Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgConvertERC20Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgConvertCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Coin.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertCoinResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgConvertERC20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgConvertERC20Response) Size() (n int) {
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
func (m *MsgConvertCoin) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgConvertCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
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
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgConvertCoinResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgConvertCoinResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertCoinResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgConvertERC20) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgConvertERC20: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
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
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
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
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
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
func (m *MsgConvertERC20Response) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgConvertERC20Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgConvertERC20Response: illegal tag %d (wire type %d)", fieldNum, wire)
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
