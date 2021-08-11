// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package examplefreightv1

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

// FreightServiceClient is the client API for FreightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FreightServiceClient interface {
	// Get a shipper.
	// See: https://google.aip.dev/131 (Standard methods: Get).
	GetShipper(ctx context.Context, in *GetShipperRequest, opts ...grpc.CallOption) (*Shipper, error)
	// List shippers.
	// See: https://google.aip.dev/132 (Standard methods: List).
	ListShippers(ctx context.Context, in *ListShippersRequest, opts ...grpc.CallOption) (*ListShippersResponse, error)
	// Create a shipper.
	// See: https://google.aip.dev/133 (Standard methods: Create).
	CreateShipper(ctx context.Context, in *CreateShipperRequest, opts ...grpc.CallOption) (*Shipper, error)
	// Update a shipper.
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateShipper(ctx context.Context, in *UpdateShipperRequest, opts ...grpc.CallOption) (*Shipper, error)
	// Delete a shipper.
	// See: https://google.aip.dev/135 (Standard methods: Delete).
	// See: https://google.aip.dev/164 (Soft delete).
	DeleteShipper(ctx context.Context, in *DeleteShipperRequest, opts ...grpc.CallOption) (*Shipper, error)
	// Get a site.
	// See: https://google.aip.dev/131 (Standard methods: Get).
	GetSite(ctx context.Context, in *GetSiteRequest, opts ...grpc.CallOption) (*Site, error)
	// Batch get sites.
	// See: https://google.aip.dev/231 (Standard methods: Get).
	BatchGetSites(ctx context.Context, in *BatchGetSitesRequest, opts ...grpc.CallOption) (*BatchGetSitesResponse, error)
	// List sites for a shipper.
	// See: https://google.aip.dev/132 (Standard methods: List).
	ListSites(ctx context.Context, in *ListSitesRequest, opts ...grpc.CallOption) (*ListSitesResponse, error)
	// Create a site.
	// See: https://google.aip.dev/133 (Standard methods: Create).
	CreateSite(ctx context.Context, in *CreateSiteRequest, opts ...grpc.CallOption) (*Site, error)
	// Update a site.
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateSite(ctx context.Context, in *UpdateSiteRequest, opts ...grpc.CallOption) (*Site, error)
	// Delete a site.
	// See: https://google.aip.dev/135 (Standard methods: Delete).
	// See: https://google.aip.dev/164 (Soft delete).
	DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*Site, error)
}

type freightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFreightServiceClient(cc grpc.ClientConnInterface) FreightServiceClient {
	return &freightServiceClient{cc}
}

func (c *freightServiceClient) GetShipper(ctx context.Context, in *GetShipperRequest, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/GetShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) ListShippers(ctx context.Context, in *ListShippersRequest, opts ...grpc.CallOption) (*ListShippersResponse, error) {
	out := new(ListShippersResponse)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/ListShippers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) CreateShipper(ctx context.Context, in *CreateShipperRequest, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/CreateShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) UpdateShipper(ctx context.Context, in *UpdateShipperRequest, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/UpdateShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) DeleteShipper(ctx context.Context, in *DeleteShipperRequest, opts ...grpc.CallOption) (*Shipper, error) {
	out := new(Shipper)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/DeleteShipper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) GetSite(ctx context.Context, in *GetSiteRequest, opts ...grpc.CallOption) (*Site, error) {
	out := new(Site)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/GetSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) BatchGetSites(ctx context.Context, in *BatchGetSitesRequest, opts ...grpc.CallOption) (*BatchGetSitesResponse, error) {
	out := new(BatchGetSitesResponse)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/BatchGetSites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) ListSites(ctx context.Context, in *ListSitesRequest, opts ...grpc.CallOption) (*ListSitesResponse, error) {
	out := new(ListSitesResponse)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/ListSites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) CreateSite(ctx context.Context, in *CreateSiteRequest, opts ...grpc.CallOption) (*Site, error) {
	out := new(Site)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/CreateSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) UpdateSite(ctx context.Context, in *UpdateSiteRequest, opts ...grpc.CallOption) (*Site, error) {
	out := new(Site)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/UpdateSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freightServiceClient) DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*Site, error) {
	out := new(Site)
	err := c.cc.Invoke(ctx, "/einride.example.freight.v1.FreightService/DeleteSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FreightServiceServer is the server API for FreightService service.
// All implementations must embed UnimplementedFreightServiceServer
// for forward compatibility
type FreightServiceServer interface {
	// Get a shipper.
	// See: https://google.aip.dev/131 (Standard methods: Get).
	GetShipper(context.Context, *GetShipperRequest) (*Shipper, error)
	// List shippers.
	// See: https://google.aip.dev/132 (Standard methods: List).
	ListShippers(context.Context, *ListShippersRequest) (*ListShippersResponse, error)
	// Create a shipper.
	// See: https://google.aip.dev/133 (Standard methods: Create).
	CreateShipper(context.Context, *CreateShipperRequest) (*Shipper, error)
	// Update a shipper.
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateShipper(context.Context, *UpdateShipperRequest) (*Shipper, error)
	// Delete a shipper.
	// See: https://google.aip.dev/135 (Standard methods: Delete).
	// See: https://google.aip.dev/164 (Soft delete).
	DeleteShipper(context.Context, *DeleteShipperRequest) (*Shipper, error)
	// Get a site.
	// See: https://google.aip.dev/131 (Standard methods: Get).
	GetSite(context.Context, *GetSiteRequest) (*Site, error)
	// Batch get sites.
	// See: https://google.aip.dev/231 (Standard methods: Get).
	BatchGetSites(context.Context, *BatchGetSitesRequest) (*BatchGetSitesResponse, error)
	// List sites for a shipper.
	// See: https://google.aip.dev/132 (Standard methods: List).
	ListSites(context.Context, *ListSitesRequest) (*ListSitesResponse, error)
	// Create a site.
	// See: https://google.aip.dev/133 (Standard methods: Create).
	CreateSite(context.Context, *CreateSiteRequest) (*Site, error)
	// Update a site.
	// See: https://google.aip.dev/134 (Standard methods: Update).
	UpdateSite(context.Context, *UpdateSiteRequest) (*Site, error)
	// Delete a site.
	// See: https://google.aip.dev/135 (Standard methods: Delete).
	// See: https://google.aip.dev/164 (Soft delete).
	DeleteSite(context.Context, *DeleteSiteRequest) (*Site, error)
	mustEmbedUnimplementedFreightServiceServer()
}

// UnimplementedFreightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFreightServiceServer struct {
}

func (UnimplementedFreightServiceServer) GetShipper(context.Context, *GetShipperRequest) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShipper not implemented")
}
func (UnimplementedFreightServiceServer) ListShippers(context.Context, *ListShippersRequest) (*ListShippersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShippers not implemented")
}
func (UnimplementedFreightServiceServer) CreateShipper(context.Context, *CreateShipperRequest) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShipper not implemented")
}
func (UnimplementedFreightServiceServer) UpdateShipper(context.Context, *UpdateShipperRequest) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShipper not implemented")
}
func (UnimplementedFreightServiceServer) DeleteShipper(context.Context, *DeleteShipperRequest) (*Shipper, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShipper not implemented")
}
func (UnimplementedFreightServiceServer) GetSite(context.Context, *GetSiteRequest) (*Site, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSite not implemented")
}
func (UnimplementedFreightServiceServer) BatchGetSites(context.Context, *BatchGetSitesRequest) (*BatchGetSitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetSites not implemented")
}
func (UnimplementedFreightServiceServer) ListSites(context.Context, *ListSitesRequest) (*ListSitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSites not implemented")
}
func (UnimplementedFreightServiceServer) CreateSite(context.Context, *CreateSiteRequest) (*Site, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSite not implemented")
}
func (UnimplementedFreightServiceServer) UpdateSite(context.Context, *UpdateSiteRequest) (*Site, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSite not implemented")
}
func (UnimplementedFreightServiceServer) DeleteSite(context.Context, *DeleteSiteRequest) (*Site, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSite not implemented")
}
func (UnimplementedFreightServiceServer) mustEmbedUnimplementedFreightServiceServer() {}

// UnsafeFreightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FreightServiceServer will
// result in compilation errors.
type UnsafeFreightServiceServer interface {
	mustEmbedUnimplementedFreightServiceServer()
}

func RegisterFreightServiceServer(s grpc.ServiceRegistrar, srv FreightServiceServer) {
	s.RegisterService(&FreightService_ServiceDesc, srv)
}

func _FreightService_GetShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShipperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).GetShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/GetShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).GetShipper(ctx, req.(*GetShipperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_ListShippers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShippersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).ListShippers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/ListShippers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).ListShippers(ctx, req.(*ListShippersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_CreateShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShipperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).CreateShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/CreateShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).CreateShipper(ctx, req.(*CreateShipperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_UpdateShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShipperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).UpdateShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/UpdateShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).UpdateShipper(ctx, req.(*UpdateShipperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_DeleteShipper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShipperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).DeleteShipper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/DeleteShipper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).DeleteShipper(ctx, req.(*DeleteShipperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_GetSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).GetSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/GetSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).GetSite(ctx, req.(*GetSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_BatchGetSites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetSitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).BatchGetSites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/BatchGetSites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).BatchGetSites(ctx, req.(*BatchGetSitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_ListSites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).ListSites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/ListSites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).ListSites(ctx, req.(*ListSitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_CreateSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).CreateSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/CreateSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).CreateSite(ctx, req.(*CreateSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_UpdateSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).UpdateSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/UpdateSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).UpdateSite(ctx, req.(*UpdateSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreightService_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreightServiceServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/einride.example.freight.v1.FreightService/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreightServiceServer).DeleteSite(ctx, req.(*DeleteSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FreightService_ServiceDesc is the grpc.ServiceDesc for FreightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FreightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "einride.example.freight.v1.FreightService",
	HandlerType: (*FreightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShipper",
			Handler:    _FreightService_GetShipper_Handler,
		},
		{
			MethodName: "ListShippers",
			Handler:    _FreightService_ListShippers_Handler,
		},
		{
			MethodName: "CreateShipper",
			Handler:    _FreightService_CreateShipper_Handler,
		},
		{
			MethodName: "UpdateShipper",
			Handler:    _FreightService_UpdateShipper_Handler,
		},
		{
			MethodName: "DeleteShipper",
			Handler:    _FreightService_DeleteShipper_Handler,
		},
		{
			MethodName: "GetSite",
			Handler:    _FreightService_GetSite_Handler,
		},
		{
			MethodName: "BatchGetSites",
			Handler:    _FreightService_BatchGetSites_Handler,
		},
		{
			MethodName: "ListSites",
			Handler:    _FreightService_ListSites_Handler,
		},
		{
			MethodName: "CreateSite",
			Handler:    _FreightService_CreateSite_Handler,
		},
		{
			MethodName: "UpdateSite",
			Handler:    _FreightService_UpdateSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _FreightService_DeleteSite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "einride/example/freight/v1/freight_service.proto",
}
