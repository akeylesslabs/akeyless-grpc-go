//
//Akeyless API
//
//The purpose of this application is to provide access to Akeyless API.
//
//The version of the OpenAPI document: 2.0
//
//Contact: support@akeyless.io
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: services/v2_service.proto

package models

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
	AkeylessV2Service_Auth_FullMethodName                  = "/akeyless_grpc.AkeylessV2Service/Auth"
	AkeylessV2Service_CreateSecret_FullMethodName          = "/akeyless_grpc.AkeylessV2Service/CreateSecret"
	AkeylessV2Service_DeleteItem_FullMethodName            = "/akeyless_grpc.AkeylessV2Service/DeleteItem"
	AkeylessV2Service_DescribeItem_FullMethodName          = "/akeyless_grpc.AkeylessV2Service/DescribeItem"
	AkeylessV2Service_GetRotatedSecretValue_FullMethodName = "/akeyless_grpc.AkeylessV2Service/GetRotatedSecretValue"
	AkeylessV2Service_GetSecretValue_FullMethodName        = "/akeyless_grpc.AkeylessV2Service/GetSecretValue"
	AkeylessV2Service_ListItems_FullMethodName             = "/akeyless_grpc.AkeylessV2Service/ListItems"
	AkeylessV2Service_UpdateSecretVal_FullMethodName       = "/akeyless_grpc.AkeylessV2Service/UpdateSecretVal"
)

// AkeylessV2ServiceClient is the client API for AkeylessV2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AkeylessV2ServiceClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthOutput, error)
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretOutput, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemOutput, error)
	DescribeItem(ctx context.Context, in *DescribeItemRequest, opts ...grpc.CallOption) (*Item, error)
	GetRotatedSecretValue(ctx context.Context, in *GetRotatedSecretValueRequest, opts ...grpc.CallOption) (*GetRotatedSecretValueResponse, error)
	GetSecretValue(ctx context.Context, in *GetSecretValueRequest, opts ...grpc.CallOption) (*GetSecretValueResponse, error)
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsInPathOutput, error)
	UpdateSecretVal(ctx context.Context, in *UpdateSecretValRequest, opts ...grpc.CallOption) (*UpdateSecretValOutput, error)
}

type akeylessV2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAkeylessV2ServiceClient(cc grpc.ClientConnInterface) AkeylessV2ServiceClient {
	return &akeylessV2ServiceClient{cc}
}

func (c *akeylessV2ServiceClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthOutput)
	err := c.cc.Invoke(ctx, AkeylessV2Service_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSecretOutput)
	err := c.cc.Invoke(ctx, AkeylessV2Service_CreateSecret_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteItemOutput)
	err := c.cc.Invoke(ctx, AkeylessV2Service_DeleteItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) DescribeItem(ctx context.Context, in *DescribeItemRequest, opts ...grpc.CallOption) (*Item, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Item)
	err := c.cc.Invoke(ctx, AkeylessV2Service_DescribeItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) GetRotatedSecretValue(ctx context.Context, in *GetRotatedSecretValueRequest, opts ...grpc.CallOption) (*GetRotatedSecretValueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRotatedSecretValueResponse)
	err := c.cc.Invoke(ctx, AkeylessV2Service_GetRotatedSecretValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) GetSecretValue(ctx context.Context, in *GetSecretValueRequest, opts ...grpc.CallOption) (*GetSecretValueResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSecretValueResponse)
	err := c.cc.Invoke(ctx, AkeylessV2Service_GetSecretValue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsInPathOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListItemsInPathOutput)
	err := c.cc.Invoke(ctx, AkeylessV2Service_ListItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *akeylessV2ServiceClient) UpdateSecretVal(ctx context.Context, in *UpdateSecretValRequest, opts ...grpc.CallOption) (*UpdateSecretValOutput, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateSecretValOutput)
	err := c.cc.Invoke(ctx, AkeylessV2Service_UpdateSecretVal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AkeylessV2ServiceServer is the server API for AkeylessV2Service service.
// All implementations must embed UnimplementedAkeylessV2ServiceServer
// for forward compatibility
type AkeylessV2ServiceServer interface {
	Auth(context.Context, *AuthRequest) (*AuthOutput, error)
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretOutput, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemOutput, error)
	DescribeItem(context.Context, *DescribeItemRequest) (*Item, error)
	GetRotatedSecretValue(context.Context, *GetRotatedSecretValueRequest) (*GetRotatedSecretValueResponse, error)
	GetSecretValue(context.Context, *GetSecretValueRequest) (*GetSecretValueResponse, error)
	ListItems(context.Context, *ListItemsRequest) (*ListItemsInPathOutput, error)
	UpdateSecretVal(context.Context, *UpdateSecretValRequest) (*UpdateSecretValOutput, error)
	mustEmbedUnimplementedAkeylessV2ServiceServer()
}

// UnimplementedAkeylessV2ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAkeylessV2ServiceServer struct {
}

func (UnimplementedAkeylessV2ServiceServer) Auth(context.Context, *AuthRequest) (*AuthOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) DescribeItem(context.Context, *DescribeItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeItem not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) GetRotatedSecretValue(context.Context, *GetRotatedSecretValueRequest) (*GetRotatedSecretValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRotatedSecretValue not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) GetSecretValue(context.Context, *GetSecretValueRequest) (*GetSecretValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretValue not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsInPathOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) UpdateSecretVal(context.Context, *UpdateSecretValRequest) (*UpdateSecretValOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecretVal not implemented")
}
func (UnimplementedAkeylessV2ServiceServer) mustEmbedUnimplementedAkeylessV2ServiceServer() {}

// UnsafeAkeylessV2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AkeylessV2ServiceServer will
// result in compilation errors.
type UnsafeAkeylessV2ServiceServer interface {
	mustEmbedUnimplementedAkeylessV2ServiceServer()
}

func RegisterAkeylessV2ServiceServer(s grpc.ServiceRegistrar, srv AkeylessV2ServiceServer) {
	s.RegisterService(&AkeylessV2Service_ServiceDesc, srv)
}

func _AkeylessV2Service_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_CreateSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_DescribeItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).DescribeItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_DescribeItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).DescribeItem(ctx, req.(*DescribeItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_GetRotatedSecretValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRotatedSecretValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).GetRotatedSecretValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_GetRotatedSecretValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).GetRotatedSecretValue(ctx, req.(*GetRotatedSecretValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_GetSecretValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).GetSecretValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_GetSecretValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).GetSecretValue(ctx, req.(*GetSecretValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_ListItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AkeylessV2Service_UpdateSecretVal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretValRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkeylessV2ServiceServer).UpdateSecretVal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AkeylessV2Service_UpdateSecretVal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkeylessV2ServiceServer).UpdateSecretVal(ctx, req.(*UpdateSecretValRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AkeylessV2Service_ServiceDesc is the grpc.ServiceDesc for AkeylessV2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AkeylessV2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akeyless_grpc.AkeylessV2Service",
	HandlerType: (*AkeylessV2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AkeylessV2Service_Auth_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _AkeylessV2Service_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _AkeylessV2Service_DeleteItem_Handler,
		},
		{
			MethodName: "DescribeItem",
			Handler:    _AkeylessV2Service_DescribeItem_Handler,
		},
		{
			MethodName: "GetRotatedSecretValue",
			Handler:    _AkeylessV2Service_GetRotatedSecretValue_Handler,
		},
		{
			MethodName: "GetSecretValue",
			Handler:    _AkeylessV2Service_GetSecretValue_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _AkeylessV2Service_ListItems_Handler,
		},
		{
			MethodName: "UpdateSecretVal",
			Handler:    _AkeylessV2Service_UpdateSecretVal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/v2_service.proto",
}
