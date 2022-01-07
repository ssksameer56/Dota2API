// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Dota2ServiceClient is the client API for Dota2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Dota2ServiceClient interface {
	GetAllHeroes(ctx context.Context, in *GetAllHeroesRequest, opts ...grpc.CallOption) (*GetAllHeroesResponse, error)
	GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Hero, error)
	GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*Item, error)
	GetLiveMatches(ctx context.Context, in *MatchIDsRequest, opts ...grpc.CallOption) (Dota2Service_GetLiveMatchesClient, error)
	GetMatchDetails(ctx context.Context, in *MatchDetailsRequest, opts ...grpc.CallOption) (Dota2Service_GetMatchDetailsClient, error)
}

type dota2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDota2ServiceClient(cc grpc.ClientConnInterface) Dota2ServiceClient {
	return &dota2ServiceClient{cc}
}

func (c *dota2ServiceClient) GetAllHeroes(ctx context.Context, in *GetAllHeroesRequest, opts ...grpc.CallOption) (*GetAllHeroesResponse, error) {
	out := new(GetAllHeroesResponse)
	err := c.cc.Invoke(ctx, "/grpc.Dota2Service/GetAllHeroes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dota2ServiceClient) GetHero(ctx context.Context, in *GetHeroRequest, opts ...grpc.CallOption) (*Hero, error) {
	out := new(Hero)
	err := c.cc.Invoke(ctx, "/grpc.Dota2Service/GetHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dota2ServiceClient) GetAllItems(ctx context.Context, in *GetAllItemsRequest, opts ...grpc.CallOption) (*GetAllItemsResponse, error) {
	out := new(GetAllItemsResponse)
	err := c.cc.Invoke(ctx, "/grpc.Dota2Service/GetAllItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dota2ServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/grpc.Dota2Service/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dota2ServiceClient) GetLiveMatches(ctx context.Context, in *MatchIDsRequest, opts ...grpc.CallOption) (Dota2Service_GetLiveMatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dota2Service_ServiceDesc.Streams[0], "/grpc.Dota2Service/GetLiveMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &dota2ServiceGetLiveMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dota2Service_GetLiveMatchesClient interface {
	Recv() (*MatchIDsResponse, error)
	grpc.ClientStream
}

type dota2ServiceGetLiveMatchesClient struct {
	grpc.ClientStream
}

func (x *dota2ServiceGetLiveMatchesClient) Recv() (*MatchIDsResponse, error) {
	m := new(MatchIDsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dota2ServiceClient) GetMatchDetails(ctx context.Context, in *MatchDetailsRequest, opts ...grpc.CallOption) (Dota2Service_GetMatchDetailsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dota2Service_ServiceDesc.Streams[1], "/grpc.Dota2Service/GetMatchDetails", opts...)
	if err != nil {
		return nil, err
	}
	x := &dota2ServiceGetMatchDetailsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dota2Service_GetMatchDetailsClient interface {
	Recv() (*MatchDetailsResponse, error)
	grpc.ClientStream
}

type dota2ServiceGetMatchDetailsClient struct {
	grpc.ClientStream
}

func (x *dota2ServiceGetMatchDetailsClient) Recv() (*MatchDetailsResponse, error) {
	m := new(MatchDetailsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Dota2ServiceServer is the server API for Dota2Service service.
// All implementations must embed UnimplementedDota2ServiceServer
// for forward compatibility
type Dota2ServiceServer interface {
	GetAllHeroes(context.Context, *GetAllHeroesRequest) (*GetAllHeroesResponse, error)
	GetHero(context.Context, *GetHeroRequest) (*Hero, error)
	GetAllItems(context.Context, *GetAllItemsRequest) (*GetAllItemsResponse, error)
	GetItem(context.Context, *GetItemRequest) (*Item, error)
	GetLiveMatches(*MatchIDsRequest, Dota2Service_GetLiveMatchesServer) error
	GetMatchDetails(*MatchDetailsRequest, Dota2Service_GetMatchDetailsServer) error
	mustEmbedUnimplementedDota2ServiceServer()
}

// UnimplementedDota2ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDota2ServiceServer struct {
}

func (UnimplementedDota2ServiceServer) GetAllHeroes(context.Context, *GetAllHeroesRequest) (*GetAllHeroesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllHeroes not implemented")
}
func (UnimplementedDota2ServiceServer) GetHero(context.Context, *GetHeroRequest) (*Hero, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHero not implemented")
}
func (UnimplementedDota2ServiceServer) GetAllItems(context.Context, *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllItems not implemented")
}
func (UnimplementedDota2ServiceServer) GetItem(context.Context, *GetItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedDota2ServiceServer) GetLiveMatches(*MatchIDsRequest, Dota2Service_GetLiveMatchesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLiveMatches not implemented")
}
func (UnimplementedDota2ServiceServer) GetMatchDetails(*MatchDetailsRequest, Dota2Service_GetMatchDetailsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMatchDetails not implemented")
}
func (UnimplementedDota2ServiceServer) mustEmbedUnimplementedDota2ServiceServer() {}

// UnsafeDota2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Dota2ServiceServer will
// result in compilation errors.
type UnsafeDota2ServiceServer interface {
	mustEmbedUnimplementedDota2ServiceServer()
}

func RegisterDota2ServiceServer(s grpc.ServiceRegistrar, srv Dota2ServiceServer) {
	s.RegisterService(&Dota2Service_ServiceDesc, srv)
}

func _Dota2Service_GetAllHeroes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllHeroesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dota2ServiceServer).GetAllHeroes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dota2Service/GetAllHeroes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dota2ServiceServer).GetAllHeroes(ctx, req.(*GetAllHeroesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dota2Service_GetHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dota2ServiceServer).GetHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dota2Service/GetHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dota2ServiceServer).GetHero(ctx, req.(*GetHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dota2Service_GetAllItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dota2ServiceServer).GetAllItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dota2Service/GetAllItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dota2ServiceServer).GetAllItems(ctx, req.(*GetAllItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dota2Service_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dota2ServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Dota2Service/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dota2ServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dota2Service_GetLiveMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchIDsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Dota2ServiceServer).GetLiveMatches(m, &dota2ServiceGetLiveMatchesServer{stream})
}

type Dota2Service_GetLiveMatchesServer interface {
	Send(*MatchIDsResponse) error
	grpc.ServerStream
}

type dota2ServiceGetLiveMatchesServer struct {
	grpc.ServerStream
}

func (x *dota2ServiceGetLiveMatchesServer) Send(m *MatchIDsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Dota2Service_GetMatchDetails_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchDetailsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(Dota2ServiceServer).GetMatchDetails(m, &dota2ServiceGetMatchDetailsServer{stream})
}

type Dota2Service_GetMatchDetailsServer interface {
	Send(*MatchDetailsResponse) error
	grpc.ServerStream
}

type dota2ServiceGetMatchDetailsServer struct {
	grpc.ServerStream
}

func (x *dota2ServiceGetMatchDetailsServer) Send(m *MatchDetailsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Dota2Service_ServiceDesc is the grpc.ServiceDesc for Dota2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dota2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Dota2Service",
	HandlerType: (*Dota2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllHeroes",
			Handler:    _Dota2Service_GetAllHeroes_Handler,
		},
		{
			MethodName: "GetHero",
			Handler:    _Dota2Service_GetHero_Handler,
		},
		{
			MethodName: "GetAllItems",
			Handler:    _Dota2Service_GetAllItems_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _Dota2Service_GetItem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLiveMatches",
			Handler:       _Dota2Service_GetLiveMatches_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMatchDetails",
			Handler:       _Dota2Service_GetMatchDetails_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "models/grpc/dota2.proto",
}