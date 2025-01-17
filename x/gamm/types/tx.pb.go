// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/osmosis-labs/osmosis/v15/x/poolmanager/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("osmosis/gamm/v1beta1/tx.proto", fileDescriptor_cfc8fd3ac7df3247) }

var fileDescriptor_cfc8fd3ac7df3247 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0xea, 0xd3, 0x40,
	0x14, 0xc5, 0x3b, 0x08, 0x52, 0x67, 0x67, 0xf0, 0x03, 0x02, 0x66, 0xa1, 0x0b, 0x37, 0x9a, 0x69,
	0x15, 0x41, 0x05, 0x95, 0x16, 0xaa, 0xb4, 0x50, 0x14, 0xbb, 0x73, 0x53, 0x26, 0xe9, 0x90, 0x0e,
	0x24, 0x73, 0x63, 0x66, 0x92, 0x26, 0x3e, 0x85, 0x82, 0x5b, 0x41, 0x70, 0xeb, 0x83, 0xb8, 0xec,
	0xd2, 0xa5, 0xb4, 0x2f, 0x22, 0xd3, 0x24, 0xc3, 0xbf, 0x5f, 0xd0, 0xb4, 0xbb, 0xcc, 0xcc, 0xf9,
	0x9d, 0x7b, 0x72, 0x2f, 0x5c, 0x7c, 0x0f, 0x64, 0x04, 0x92, 0x4b, 0x12, 0xd0, 0x28, 0x22, 0x59,
	0xd7, 0x63, 0x8a, 0x76, 0x89, 0xca, 0xdd, 0x38, 0x01, 0x05, 0xd6, 0x83, 0x59, 0x11, 0x31, 0x21,
	0x39, 0x88, 0xbc, 0xf8, 0xe2, 0x9a, 0x83, 0xab, 0xd5, 0x6e, 0xa5, 0xb6, 0x6f, 0x05, 0x10, 0xc0,
	0x46, 0x4f, 0xf4, 0x57, 0x89, 0xda, 0x8e, 0xbf, 0xb1, 0x26, 0x1e, 0x95, 0xcc, 0x18, 0xfb, 0xc0,
	0x45, 0xf5, 0xfe, 0xa8, 0xae, 0x1c, 0x03, 0x84, 0x11, 0x15, 0x34, 0x60, 0x89, 0xd1, 0xc9, 0x05,
	0x8d, 0xa7, 0x09, 0xa4, 0x8a, 0x55, 0xea, 0xfb, 0x47, 0x72, 0x4e, 0xb5, 0xb2, 0xd2, 0x3c, 0x3c,
	0xa6, 0x09, 0xf9, 0xe7, 0x94, 0xcf, 0xb8, 0x2a, 0x4a, 0xe1, 0x93, 0xdf, 0x37, 0xf0, 0xb5, 0xb1,
	0x0c, 0xac, 0x0c, 0xb7, 0x47, 0xc0, 0xc5, 0x07, 0x80, 0xd0, 0xea, 0xb8, 0x27, 0xfc, 0xaa, 0x3b,
	0x96, 0x41, 0x4d, 0xd8, 0xcf, 0x9b, 0x12, 0x1f, 0x99, 0x8c, 0x41, 0x48, 0x66, 0xfd, 0x44, 0xf8,
	0x8e, 0xbe, 0x9c, 0x2c, 0x68, 0x3c, 0xc8, 0x15, 0x4b, 0x44, 0x2f, 0x82, 0x54, 0xa8, 0xa1, 0xb0,
	0x5e, 0x37, 0x31, 0xdd, 0xe7, 0xed, 0xb7, 0x97, 0xf1, 0x07, 0x23, 0x4e, 0xe6, 0x34, 0x61, 0xa5,
	0xe2, 0x7d, 0xaa, 0x9a, 0x47, 0xdc, 0xe6, 0x9b, 0x47, 0xdc, 0xe6, 0x4d, 0xc4, 0x0c, 0xb7, 0x07,
	0x39, 0x57, 0xcd, 0xa6, 0x57, 0x13, 0xa7, 0x4f, 0xaf, 0x26, 0x4c, 0xdd, 0x5f, 0x08, 0xdf, 0xd5,
	0x97, 0xbb, 0xdd, 0xd3, 0xbd, 0x79, 0xd3, 0xc4, 0xf5, 0x80, 0x81, 0xfd, 0xee, 0x42, 0x03, 0x93,
	0xf2, 0x07, 0xc2, 0xb7, 0x6b, 0xcd, 0x95, 0x06, 0x0e, 0x85, 0xf5, 0xaa, 0x69, 0x89, 0x2d, 0xdc,
	0x1e, 0x5c, 0x84, 0x9b, 0x7c, 0xdf, 0x10, 0xbe, 0x59, 0xe6, 0xa7, 0xbe, 0x32, 0xd9, 0x5e, 0x9c,
	0x6a, 0xbe, 0x87, 0xda, 0xbd, 0xb3, 0x51, 0x93, 0xe9, 0x3b, 0xc2, 0xd6, 0xce, 0xab, 0x1e, 0xea,
	0xcb, 0x33, 0x9d, 0xf5, 0x3c, 0xfb, 0xe7, 0xb3, 0x75, 0xac, 0xfe, 0xe8, 0xcf, 0xca, 0x41, 0xcb,
	0x95, 0x83, 0xfe, 0xad, 0x1c, 0xf4, 0x75, 0xed, 0xb4, 0x96, 0x6b, 0xa7, 0xf5, 0x77, 0xed, 0xb4,
	0x3e, 0x75, 0x02, 0xae, 0xe6, 0xa9, 0xe7, 0xfa, 0x10, 0x91, 0x6a, 0xf9, 0x3d, 0x0e, 0xa9, 0x27,
	0xeb, 0x03, 0xc9, 0xba, 0xcf, 0x48, 0x5e, 0xee, 0x43, 0x55, 0xc4, 0x4c, 0x7a, 0xd7, 0x37, 0x1b,
	0xf0, 0xe9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x9d, 0x4b, 0x7f, 0xf8, 0x05, 0x00, 0x00,
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
	JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error)
	JoinSwapExternAmountIn(ctx context.Context, in *MsgJoinSwapExternAmountIn, opts ...grpc.CallOption) (*MsgJoinSwapExternAmountInResponse, error)
	JoinSwapShareAmountOut(ctx context.Context, in *MsgJoinSwapShareAmountOut, opts ...grpc.CallOption) (*MsgJoinSwapShareAmountOutResponse, error)
	ExitPool(ctx context.Context, in *MsgExitPool, opts ...grpc.CallOption) (*MsgExitPoolResponse, error)
	ExitSwapExternAmountOut(ctx context.Context, in *MsgExitSwapExternAmountOut, opts ...grpc.CallOption) (*MsgExitSwapExternAmountOutResponse, error)
	ExitSwapShareAmountIn(ctx context.Context, in *MsgExitSwapShareAmountIn, opts ...grpc.CallOption) (*MsgExitSwapShareAmountInResponse, error)
	SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error) {
	out := new(MsgJoinPoolResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinSwapExternAmountIn(ctx context.Context, in *MsgJoinSwapExternAmountIn, opts ...grpc.CallOption) (*MsgJoinSwapExternAmountInResponse, error) {
	out := new(MsgJoinSwapExternAmountInResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinSwapExternAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinSwapShareAmountOut(ctx context.Context, in *MsgJoinSwapShareAmountOut, opts ...grpc.CallOption) (*MsgJoinSwapShareAmountOutResponse, error) {
	out := new(MsgJoinSwapShareAmountOutResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinSwapShareAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitPool(ctx context.Context, in *MsgExitPool, opts ...grpc.CallOption) (*MsgExitPoolResponse, error) {
	out := new(MsgExitPoolResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitSwapExternAmountOut(ctx context.Context, in *MsgExitSwapExternAmountOut, opts ...grpc.CallOption) (*MsgExitSwapExternAmountOutResponse, error) {
	out := new(MsgExitSwapExternAmountOutResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitSwapExternAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitSwapShareAmountIn(ctx context.Context, in *MsgExitSwapShareAmountIn, opts ...grpc.CallOption) (*MsgExitSwapShareAmountInResponse, error) {
	out := new(MsgExitSwapShareAmountInResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitSwapShareAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error) {
	out := new(MsgSwapExactAmountInResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/SwapExactAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error) {
	out := new(MsgSwapExactAmountOutResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.gamm.v1beta1.Msg/SwapExactAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	JoinPool(context.Context, *MsgJoinPool) (*MsgJoinPoolResponse, error)
	JoinSwapExternAmountIn(context.Context, *MsgJoinSwapExternAmountIn) (*MsgJoinSwapExternAmountInResponse, error)
	JoinSwapShareAmountOut(context.Context, *MsgJoinSwapShareAmountOut) (*MsgJoinSwapShareAmountOutResponse, error)
	ExitPool(context.Context, *MsgExitPool) (*MsgExitPoolResponse, error)
	ExitSwapExternAmountOut(context.Context, *MsgExitSwapExternAmountOut) (*MsgExitSwapExternAmountOutResponse, error)
	ExitSwapShareAmountIn(context.Context, *MsgExitSwapShareAmountIn) (*MsgExitSwapShareAmountInResponse, error)
	SwapExactAmountIn(context.Context, *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(context.Context, *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) JoinPool(ctx context.Context, req *MsgJoinPool) (*MsgJoinPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPool not implemented")
}
func (*UnimplementedMsgServer) JoinSwapExternAmountIn(ctx context.Context, req *MsgJoinSwapExternAmountIn) (*MsgJoinSwapExternAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSwapExternAmountIn not implemented")
}
func (*UnimplementedMsgServer) JoinSwapShareAmountOut(ctx context.Context, req *MsgJoinSwapShareAmountOut) (*MsgJoinSwapShareAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSwapShareAmountOut not implemented")
}
func (*UnimplementedMsgServer) ExitPool(ctx context.Context, req *MsgExitPool) (*MsgExitPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitPool not implemented")
}
func (*UnimplementedMsgServer) ExitSwapExternAmountOut(ctx context.Context, req *MsgExitSwapExternAmountOut) (*MsgExitSwapExternAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitSwapExternAmountOut not implemented")
}
func (*UnimplementedMsgServer) ExitSwapShareAmountIn(ctx context.Context, req *MsgExitSwapShareAmountIn) (*MsgExitSwapShareAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitSwapShareAmountIn not implemented")
}
func (*UnimplementedMsgServer) SwapExactAmountIn(ctx context.Context, req *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountIn not implemented")
}
func (*UnimplementedMsgServer) SwapExactAmountOut(ctx context.Context, req *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountOut not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_JoinPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinPool(ctx, req.(*MsgJoinPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinSwapExternAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinSwapExternAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinSwapExternAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinSwapExternAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinSwapExternAmountIn(ctx, req.(*MsgJoinSwapExternAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinSwapShareAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinSwapShareAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinSwapShareAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/JoinSwapShareAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinSwapShareAmountOut(ctx, req.(*MsgJoinSwapShareAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitPool(ctx, req.(*MsgExitPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitSwapExternAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitSwapExternAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitSwapExternAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitSwapExternAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitSwapExternAmountOut(ctx, req.(*MsgExitSwapExternAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitSwapShareAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitSwapShareAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitSwapShareAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/ExitSwapShareAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitSwapShareAmountIn(ctx, req.(*MsgExitSwapShareAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapExactAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/SwapExactAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountIn(ctx, req.(*MsgSwapExactAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapExactAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.gamm.v1beta1.Msg/SwapExactAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountOut(ctx, req.(*MsgSwapExactAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.gamm.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPool",
			Handler:    _Msg_JoinPool_Handler,
		},
		{
			MethodName: "JoinSwapExternAmountIn",
			Handler:    _Msg_JoinSwapExternAmountIn_Handler,
		},
		{
			MethodName: "JoinSwapShareAmountOut",
			Handler:    _Msg_JoinSwapShareAmountOut_Handler,
		},
		{
			MethodName: "ExitPool",
			Handler:    _Msg_ExitPool_Handler,
		},
		{
			MethodName: "ExitSwapExternAmountOut",
			Handler:    _Msg_ExitSwapExternAmountOut_Handler,
		},
		{
			MethodName: "ExitSwapShareAmountIn",
			Handler:    _Msg_ExitSwapShareAmountIn_Handler,
		},
		{
			MethodName: "SwapExactAmountIn",
			Handler:    _Msg_SwapExactAmountIn_Handler,
		},
		{
			MethodName: "SwapExactAmountOut",
			Handler:    _Msg_SwapExactAmountOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/v1beta1/tx.proto",
}
