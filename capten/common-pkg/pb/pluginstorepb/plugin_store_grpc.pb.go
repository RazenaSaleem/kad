// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: plugin_store.proto

package pluginstorepb

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

const (
	PluginStore_ConfigurePluginStore_FullMethodName = "/pluginstorepb.PluginStore/ConfigurePluginStore"
	PluginStore_GetPluginStoreConfig_FullMethodName = "/pluginstorepb.PluginStore/GetPluginStoreConfig"
	PluginStore_SyncPluginStore_FullMethodName      = "/pluginstorepb.PluginStore/SyncPluginStore"
	PluginStore_GetPlugins_FullMethodName           = "/pluginstorepb.PluginStore/GetPlugins"
	PluginStore_GetPluginValues_FullMethodName      = "/pluginstorepb.PluginStore/GetPluginValues"
	PluginStore_GetPluginData_FullMethodName        = "/pluginstorepb.PluginStore/GetPluginData"
	PluginStore_DeployPlugin_FullMethodName         = "/pluginstorepb.PluginStore/DeployPlugin"
	PluginStore_UnDeployPlugin_FullMethodName       = "/pluginstorepb.PluginStore/UnDeployPlugin"
)

// PluginStoreClient is the client API for PluginStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginStoreClient interface {
	ConfigurePluginStore(ctx context.Context, in *ConfigurePluginStoreRequest, opts ...grpc.CallOption) (*ConfigurePluginStoreResponse, error)
	GetPluginStoreConfig(ctx context.Context, in *GetPluginStoreConfigRequest, opts ...grpc.CallOption) (*GetPluginStoreConfigResponse, error)
	SyncPluginStore(ctx context.Context, in *SyncPluginStoreRequest, opts ...grpc.CallOption) (*SyncPluginStoreResponse, error)
	GetPlugins(ctx context.Context, in *GetPluginsRequest, opts ...grpc.CallOption) (*GetPluginsResponse, error)
	GetPluginValues(ctx context.Context, in *GetPluginValuesRequest, opts ...grpc.CallOption) (*GetPluginValuesResponse, error)
	GetPluginData(ctx context.Context, in *GetPluginDataRequest, opts ...grpc.CallOption) (*GetPluginDataResponse, error)
	DeployPlugin(ctx context.Context, in *DeployPluginRequest, opts ...grpc.CallOption) (*DeployPluginResponse, error)
	UnDeployPlugin(ctx context.Context, in *UnDeployPluginRequest, opts ...grpc.CallOption) (*UnDeployPluginResponse, error)
}

type pluginStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginStoreClient(cc grpc.ClientConnInterface) PluginStoreClient {
	return &pluginStoreClient{cc}
}

func (c *pluginStoreClient) ConfigurePluginStore(ctx context.Context, in *ConfigurePluginStoreRequest, opts ...grpc.CallOption) (*ConfigurePluginStoreResponse, error) {
	out := new(ConfigurePluginStoreResponse)
	err := c.cc.Invoke(ctx, PluginStore_ConfigurePluginStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) GetPluginStoreConfig(ctx context.Context, in *GetPluginStoreConfigRequest, opts ...grpc.CallOption) (*GetPluginStoreConfigResponse, error) {
	out := new(GetPluginStoreConfigResponse)
	err := c.cc.Invoke(ctx, PluginStore_GetPluginStoreConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) SyncPluginStore(ctx context.Context, in *SyncPluginStoreRequest, opts ...grpc.CallOption) (*SyncPluginStoreResponse, error) {
	out := new(SyncPluginStoreResponse)
	err := c.cc.Invoke(ctx, PluginStore_SyncPluginStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) GetPlugins(ctx context.Context, in *GetPluginsRequest, opts ...grpc.CallOption) (*GetPluginsResponse, error) {
	out := new(GetPluginsResponse)
	err := c.cc.Invoke(ctx, PluginStore_GetPlugins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) GetPluginValues(ctx context.Context, in *GetPluginValuesRequest, opts ...grpc.CallOption) (*GetPluginValuesResponse, error) {
	out := new(GetPluginValuesResponse)
	err := c.cc.Invoke(ctx, PluginStore_GetPluginValues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) GetPluginData(ctx context.Context, in *GetPluginDataRequest, opts ...grpc.CallOption) (*GetPluginDataResponse, error) {
	out := new(GetPluginDataResponse)
	err := c.cc.Invoke(ctx, PluginStore_GetPluginData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) DeployPlugin(ctx context.Context, in *DeployPluginRequest, opts ...grpc.CallOption) (*DeployPluginResponse, error) {
	out := new(DeployPluginResponse)
	err := c.cc.Invoke(ctx, PluginStore_DeployPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginStoreClient) UnDeployPlugin(ctx context.Context, in *UnDeployPluginRequest, opts ...grpc.CallOption) (*UnDeployPluginResponse, error) {
	out := new(UnDeployPluginResponse)
	err := c.cc.Invoke(ctx, PluginStore_UnDeployPlugin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginStoreServer is the server API for PluginStore service.
// All implementations must embed UnimplementedPluginStoreServer
// for forward compatibility
type PluginStoreServer interface {
	ConfigurePluginStore(context.Context, *ConfigurePluginStoreRequest) (*ConfigurePluginStoreResponse, error)
	GetPluginStoreConfig(context.Context, *GetPluginStoreConfigRequest) (*GetPluginStoreConfigResponse, error)
	SyncPluginStore(context.Context, *SyncPluginStoreRequest) (*SyncPluginStoreResponse, error)
	GetPlugins(context.Context, *GetPluginsRequest) (*GetPluginsResponse, error)
	GetPluginValues(context.Context, *GetPluginValuesRequest) (*GetPluginValuesResponse, error)
	GetPluginData(context.Context, *GetPluginDataRequest) (*GetPluginDataResponse, error)
	DeployPlugin(context.Context, *DeployPluginRequest) (*DeployPluginResponse, error)
	UnDeployPlugin(context.Context, *UnDeployPluginRequest) (*UnDeployPluginResponse, error)
	mustEmbedUnimplementedPluginStoreServer()
}

// UnimplementedPluginStoreServer must be embedded to have forward compatible implementations.
type UnimplementedPluginStoreServer struct {
}

func (UnimplementedPluginStoreServer) ConfigurePluginStore(context.Context, *ConfigurePluginStoreRequest) (*ConfigurePluginStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigurePluginStore not implemented")
}
func (UnimplementedPluginStoreServer) GetPluginStoreConfig(context.Context, *GetPluginStoreConfigRequest) (*GetPluginStoreConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginStoreConfig not implemented")
}
func (UnimplementedPluginStoreServer) SyncPluginStore(context.Context, *SyncPluginStoreRequest) (*SyncPluginStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncPluginStore not implemented")
}
func (UnimplementedPluginStoreServer) GetPlugins(context.Context, *GetPluginsRequest) (*GetPluginsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugins not implemented")
}
func (UnimplementedPluginStoreServer) GetPluginValues(context.Context, *GetPluginValuesRequest) (*GetPluginValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginValues not implemented")
}
func (UnimplementedPluginStoreServer) GetPluginData(context.Context, *GetPluginDataRequest) (*GetPluginDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginData not implemented")
}
func (UnimplementedPluginStoreServer) DeployPlugin(context.Context, *DeployPluginRequest) (*DeployPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployPlugin not implemented")
}
func (UnimplementedPluginStoreServer) UnDeployPlugin(context.Context, *UnDeployPluginRequest) (*UnDeployPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnDeployPlugin not implemented")
}
func (UnimplementedPluginStoreServer) mustEmbedUnimplementedPluginStoreServer() {}

// UnsafePluginStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginStoreServer will
// result in compilation errors.
type UnsafePluginStoreServer interface {
	mustEmbedUnimplementedPluginStoreServer()
}

func RegisterPluginStoreServer(s grpc.ServiceRegistrar, srv PluginStoreServer) {
	s.RegisterService(&PluginStore_ServiceDesc, srv)
}

func _PluginStore_ConfigurePluginStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurePluginStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).ConfigurePluginStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_ConfigurePluginStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).ConfigurePluginStore(ctx, req.(*ConfigurePluginStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_GetPluginStoreConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginStoreConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).GetPluginStoreConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_GetPluginStoreConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).GetPluginStoreConfig(ctx, req.(*GetPluginStoreConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_SyncPluginStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncPluginStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).SyncPluginStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_SyncPluginStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).SyncPluginStore(ctx, req.(*SyncPluginStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_GetPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).GetPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_GetPlugins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).GetPlugins(ctx, req.(*GetPluginsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_GetPluginValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).GetPluginValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_GetPluginValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).GetPluginValues(ctx, req.(*GetPluginValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_GetPluginData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).GetPluginData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_GetPluginData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).GetPluginData(ctx, req.(*GetPluginDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_DeployPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).DeployPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_DeployPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).DeployPlugin(ctx, req.(*DeployPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginStore_UnDeployPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnDeployPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginStoreServer).UnDeployPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PluginStore_UnDeployPlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginStoreServer).UnDeployPlugin(ctx, req.(*UnDeployPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginStore_ServiceDesc is the grpc.ServiceDesc for PluginStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginstorepb.PluginStore",
	HandlerType: (*PluginStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigurePluginStore",
			Handler:    _PluginStore_ConfigurePluginStore_Handler,
		},
		{
			MethodName: "GetPluginStoreConfig",
			Handler:    _PluginStore_GetPluginStoreConfig_Handler,
		},
		{
			MethodName: "SyncPluginStore",
			Handler:    _PluginStore_SyncPluginStore_Handler,
		},
		{
			MethodName: "GetPlugins",
			Handler:    _PluginStore_GetPlugins_Handler,
		},
		{
			MethodName: "GetPluginValues",
			Handler:    _PluginStore_GetPluginValues_Handler,
		},
		{
			MethodName: "GetPluginData",
			Handler:    _PluginStore_GetPluginData_Handler,
		},
		{
			MethodName: "DeployPlugin",
			Handler:    _PluginStore_DeployPlugin_Handler,
		},
		{
			MethodName: "UnDeployPlugin",
			Handler:    _PluginStore_UnDeployPlugin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin_store.proto",
}