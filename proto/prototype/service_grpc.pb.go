// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package prototype

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RaidoChainServiceClient is the client API for RaidoChainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaidoChainServiceClient interface {
	// GetUTxO get all unspent transaction outputs of given address
	GetUTxO(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*UTxOResponse, error)
	// GetStatus returns node status
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type raidoChainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaidoChainServiceClient(cc grpc.ClientConnInterface) RaidoChainServiceClient {
	return &raidoChainServiceClient{cc}
}

func (c *raidoChainServiceClient) GetUTxO(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*UTxOResponse, error) {
	out := new(UTxOResponse)
	err := c.cc.Invoke(ctx, "/rdo.service.RaidoChainService/GetUTxO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raidoChainServiceClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/rdo.service.RaidoChainService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaidoChainServiceServer is the server API for RaidoChainService service.
// All implementations must embed UnimplementedRaidoChainServiceServer
// for forward compatibility
type RaidoChainServiceServer interface {
	// GetUTxO get all unspent transaction outputs of given address
	GetUTxO(context.Context, *AddressRequest) (*UTxOResponse, error)
	// GetStatus returns node status
	GetStatus(context.Context, *emptypb.Empty) (*StatusResponse, error)
	mustEmbedUnimplementedRaidoChainServiceServer()
}

// UnimplementedRaidoChainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRaidoChainServiceServer struct {
}

func (UnimplementedRaidoChainServiceServer) GetUTxO(context.Context, *AddressRequest) (*UTxOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUTxO not implemented")
}
func (UnimplementedRaidoChainServiceServer) GetStatus(context.Context, *emptypb.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedRaidoChainServiceServer) mustEmbedUnimplementedRaidoChainServiceServer() {}

// UnsafeRaidoChainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaidoChainServiceServer will
// result in compilation errors.
type UnsafeRaidoChainServiceServer interface {
	mustEmbedUnimplementedRaidoChainServiceServer()
}

func RegisterRaidoChainServiceServer(s grpc.ServiceRegistrar, srv RaidoChainServiceServer) {
	s.RegisterService(&RaidoChainService_ServiceDesc, srv)
}

func _RaidoChainService_GetUTxO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaidoChainServiceServer).GetUTxO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdo.service.RaidoChainService/GetUTxO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaidoChainServiceServer).GetUTxO(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaidoChainService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaidoChainServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdo.service.RaidoChainService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaidoChainServiceServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RaidoChainService_ServiceDesc is the grpc.ServiceDesc for RaidoChainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaidoChainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rdo.service.RaidoChainService",
	HandlerType: (*RaidoChainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUTxO",
			Handler:    _RaidoChainService_GetUTxO_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _RaidoChainService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prototype/service.proto",
}

// AttestationServiceClient is the client API for AttestationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttestationServiceClient interface {
	// SendTx send signed transaction to the node.
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type attestationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttestationServiceClient(cc grpc.ClientConnInterface) AttestationServiceClient {
	return &attestationServiceClient{cc}
}

func (c *attestationServiceClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/rdo.service.AttestationService/SendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttestationServiceServer is the server API for AttestationService service.
// All implementations must embed UnimplementedAttestationServiceServer
// for forward compatibility
type AttestationServiceServer interface {
	// SendTx send signed transaction to the node.
	SendTx(context.Context, *SendTxRequest) (*ErrorResponse, error)
	mustEmbedUnimplementedAttestationServiceServer()
}

// UnimplementedAttestationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttestationServiceServer struct {
}

func (UnimplementedAttestationServiceServer) SendTx(context.Context, *SendTxRequest) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedAttestationServiceServer) mustEmbedUnimplementedAttestationServiceServer() {}

// UnsafeAttestationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttestationServiceServer will
// result in compilation errors.
type UnsafeAttestationServiceServer interface {
	mustEmbedUnimplementedAttestationServiceServer()
}

func RegisterAttestationServiceServer(s grpc.ServiceRegistrar, srv AttestationServiceServer) {
	s.RegisterService(&AttestationService_ServiceDesc, srv)
}

func _AttestationService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttestationServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdo.service.AttestationService/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttestationServiceServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttestationService_ServiceDesc is the grpc.ServiceDesc for AttestationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttestationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rdo.service.AttestationService",
	HandlerType: (*AttestationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTx",
			Handler:    _AttestationService_SendTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prototype/service.proto",
}