// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: authorization/v1/authorization.proto

package authorization

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthorizationService_AddActor_FullMethodName          = "/authorization.v1.AuthorizationService/AddActor"
	AuthorizationService_GetActor_FullMethodName          = "/authorization.v1.AuthorizationService/GetActor"
	AuthorizationService_GetActors_FullMethodName         = "/authorization.v1.AuthorizationService/GetActors"
	AuthorizationService_GetTeams_FullMethodName          = "/authorization.v1.AuthorizationService/GetTeams"
	AuthorizationService_AddResource_FullMethodName       = "/authorization.v1.AuthorizationService/AddResource"
	AuthorizationService_GetResources_FullMethodName      = "/authorization.v1.AuthorizationService/GetResources"
	AuthorizationService_GetActorResources_FullMethodName = "/authorization.v1.AuthorizationService/GetActorResources"
	AuthorizationService_ArchiveResource_FullMethodName   = "/authorization.v1.AuthorizationService/ArchiveResource"
)

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	AddActor(ctx context.Context, in *AddActorRequest, opts ...grpc.CallOption) (*AddActorResponse, error)
	GetActor(ctx context.Context, in *GetActorRequest, opts ...grpc.CallOption) (*GetActorResponse, error)
	GetActors(ctx context.Context, in *GetActorsRequest, opts ...grpc.CallOption) (*GetActorsResponse, error)
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error)
	AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*AddResourceResponse, error)
	GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error)
	GetActorResources(ctx context.Context, in *GetActorResourcesRequest, opts ...grpc.CallOption) (*GetActorResourcesResponse, error)
	ArchiveResource(ctx context.Context, in *ArchiveResourceRequest, opts ...grpc.CallOption) (*ArchiveResourceResponse, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) AddActor(ctx context.Context, in *AddActorRequest, opts ...grpc.CallOption) (*AddActorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddActorResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_AddActor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetActor(ctx context.Context, in *GetActorRequest, opts ...grpc.CallOption) (*GetActorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActorResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetActor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetActors(ctx context.Context, in *GetActorsRequest, opts ...grpc.CallOption) (*GetActorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActorsResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetActors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTeamsResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) AddResource(ctx context.Context, in *AddResourceRequest, opts ...grpc.CallOption) (*AddResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddResourceResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_AddResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetActorResources(ctx context.Context, in *GetActorResourcesRequest, opts ...grpc.CallOption) (*GetActorResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetActorResourcesResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_GetActorResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) ArchiveResource(ctx context.Context, in *ArchiveResourceRequest, opts ...grpc.CallOption) (*ArchiveResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ArchiveResourceResponse)
	err := c.cc.Invoke(ctx, AuthorizationService_ArchiveResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations must embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	AddActor(context.Context, *AddActorRequest) (*AddActorResponse, error)
	GetActor(context.Context, *GetActorRequest) (*GetActorResponse, error)
	GetActors(context.Context, *GetActorsRequest) (*GetActorsResponse, error)
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsResponse, error)
	AddResource(context.Context, *AddResourceRequest) (*AddResourceResponse, error)
	GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error)
	GetActorResources(context.Context, *GetActorResourcesRequest) (*GetActorResourcesResponse, error)
	ArchiveResource(context.Context, *ArchiveResourceRequest) (*ArchiveResourceResponse, error)
	mustEmbedUnimplementedAuthorizationServiceServer()
}

// UnimplementedAuthorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) AddActor(context.Context, *AddActorRequest) (*AddActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActor not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetActor(context.Context, *GetActorRequest) (*GetActorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActor not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetActors(context.Context, *GetActorsRequest) (*GetActorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActors not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}
func (UnimplementedAuthorizationServiceServer) AddResource(context.Context, *AddResourceRequest) (*AddResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddResource not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetActorResources(context.Context, *GetActorResourcesRequest) (*GetActorResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActorResources not implemented")
}
func (UnimplementedAuthorizationServiceServer) ArchiveResource(context.Context, *ArchiveResourceRequest) (*ArchiveResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveResource not implemented")
}
func (UnimplementedAuthorizationServiceServer) mustEmbedUnimplementedAuthorizationServiceServer() {}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_AddActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).AddActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_AddActor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).AddActor(ctx, req.(*AddActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetActor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetActor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetActor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetActor(ctx, req.(*GetActorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetActors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetActors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetActors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetActors(ctx, req.(*GetActorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetTeams(ctx, req.(*GetTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_AddResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).AddResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_AddResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).AddResource(ctx, req.(*AddResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetResources(ctx, req.(*GetResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetActorResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActorResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetActorResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetActorResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetActorResources(ctx, req.(*GetActorResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_ArchiveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).ArchiveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_ArchiveResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).ArchiveResource(ctx, req.(*ArchiveResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.v1.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddActor",
			Handler:    _AuthorizationService_AddActor_Handler,
		},
		{
			MethodName: "GetActor",
			Handler:    _AuthorizationService_GetActor_Handler,
		},
		{
			MethodName: "GetActors",
			Handler:    _AuthorizationService_GetActors_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _AuthorizationService_GetTeams_Handler,
		},
		{
			MethodName: "AddResource",
			Handler:    _AuthorizationService_AddResource_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _AuthorizationService_GetResources_Handler,
		},
		{
			MethodName: "GetActorResources",
			Handler:    _AuthorizationService_GetActorResources_Handler,
		},
		{
			MethodName: "ArchiveResource",
			Handler:    _AuthorizationService_ArchiveResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization/v1/authorization.proto",
}
