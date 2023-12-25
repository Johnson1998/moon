// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: system/api.proto

package system

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	// 创建API数据
	CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiReply, error)
	// 更新API数据
	UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiReply, error)
	// 删除API数据
	DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiReply, error)
	// 获取API数据
	GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiReply, error)
	// 获取API列表
	ListApi(ctx context.Context, in *ListApiRequest, opts ...grpc.CallOption) (*ListApiReply, error)
	// 获取API下拉列表
	SelectApi(ctx context.Context, in *SelectApiRequest, opts ...grpc.CallOption) (*SelectApiReply, error)
	// 修改API状态
	EditApiStatus(ctx context.Context, in *EditApiStatusRequest, opts ...grpc.CallOption) (*EditApiStatusReply, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) CreateApi(ctx context.Context, in *CreateApiRequest, opts ...grpc.CallOption) (*CreateApiReply, error) {
	out := new(CreateApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/CreateApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateApi(ctx context.Context, in *UpdateApiRequest, opts ...grpc.CallOption) (*UpdateApiReply, error) {
	out := new(UpdateApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/UpdateApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteApi(ctx context.Context, in *DeleteApiRequest, opts ...grpc.CallOption) (*DeleteApiReply, error) {
	out := new(DeleteApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/DeleteApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetApi(ctx context.Context, in *GetApiRequest, opts ...grpc.CallOption) (*GetApiReply, error) {
	out := new(GetApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/GetApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListApi(ctx context.Context, in *ListApiRequest, opts ...grpc.CallOption) (*ListApiReply, error) {
	out := new(ListApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/ListApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SelectApi(ctx context.Context, in *SelectApiRequest, opts ...grpc.CallOption) (*SelectApiReply, error) {
	out := new(SelectApiReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/SelectApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) EditApiStatus(ctx context.Context, in *EditApiStatusRequest, opts ...grpc.CallOption) (*EditApiStatusReply, error) {
	out := new(EditApiStatusReply)
	err := c.cc.Invoke(ctx, "/api.system.Api/EditApiStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	// 创建API数据
	CreateApi(context.Context, *CreateApiRequest) (*CreateApiReply, error)
	// 更新API数据
	UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiReply, error)
	// 删除API数据
	DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiReply, error)
	// 获取API数据
	GetApi(context.Context, *GetApiRequest) (*GetApiReply, error)
	// 获取API列表
	ListApi(context.Context, *ListApiRequest) (*ListApiReply, error)
	// 获取API下拉列表
	SelectApi(context.Context, *SelectApiRequest) (*SelectApiReply, error)
	// 修改API状态
	EditApiStatus(context.Context, *EditApiStatusRequest) (*EditApiStatusReply, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) CreateApi(context.Context, *CreateApiRequest) (*CreateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedApiServer) UpdateApi(context.Context, *UpdateApiRequest) (*UpdateApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedApiServer) DeleteApi(context.Context, *DeleteApiRequest) (*DeleteApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedApiServer) GetApi(context.Context, *GetApiRequest) (*GetApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApi not implemented")
}
func (UnimplementedApiServer) ListApi(context.Context, *ListApiRequest) (*ListApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApi not implemented")
}
func (UnimplementedApiServer) SelectApi(context.Context, *SelectApiRequest) (*SelectApiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectApi not implemented")
}
func (UnimplementedApiServer) EditApiStatus(context.Context, *EditApiStatusRequest) (*EditApiStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditApiStatus not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/CreateApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateApi(ctx, req.(*CreateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/UpdateApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateApi(ctx, req.(*UpdateApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/DeleteApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteApi(ctx, req.(*DeleteApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/GetApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetApi(ctx, req.(*GetApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/ListApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListApi(ctx, req.(*ListApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SelectApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SelectApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/SelectApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SelectApi(ctx, req.(*SelectApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_EditApiStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditApiStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).EditApiStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.system.Api/EditApiStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).EditApiStatus(ctx, req.(*EditApiStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApi",
			Handler:    _Api_CreateApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _Api_UpdateApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _Api_DeleteApi_Handler,
		},
		{
			MethodName: "GetApi",
			Handler:    _Api_GetApi_Handler,
		},
		{
			MethodName: "ListApi",
			Handler:    _Api_ListApi_Handler,
		},
		{
			MethodName: "SelectApi",
			Handler:    _Api_SelectApi_Handler,
		},
		{
			MethodName: "EditApiStatus",
			Handler:    _Api_EditApiStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/api.proto",
}