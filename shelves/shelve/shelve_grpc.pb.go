// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: shelve.proto

package shelve

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
	ShelveService_GetShelves_FullMethodName  = "/ShelveService/GetShelves"
	ShelveService_CreateShelf_FullMethodName = "/ShelveService/CreateShelf"
	ShelveService_UpdateShelf_FullMethodName = "/ShelveService/UpdateShelf"
	ShelveService_RemoveShelf_FullMethodName = "/ShelveService/RemoveShelf"
	ShelveService_GetShelf_FullMethodName    = "/ShelveService/GetShelf"
)

// ShelveServiceClient is the client API for ShelveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShelveServiceClient interface {
	GetShelves(ctx context.Context, in *GetShelvesRequest, opts ...grpc.CallOption) (*GetShelvesResponse, error)
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error)
	UpdateShelf(ctx context.Context, in *UpdateShelfRequest, opts ...grpc.CallOption) (*UpdateShelfResponse, error)
	RemoveShelf(ctx context.Context, in *RemoveShelfRequest, opts ...grpc.CallOption) (*RemoveShelfResponse, error)
	GetShelf(ctx context.Context, in *RemoveShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error)
}

type shelveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShelveServiceClient(cc grpc.ClientConnInterface) ShelveServiceClient {
	return &shelveServiceClient{cc}
}

func (c *shelveServiceClient) GetShelves(ctx context.Context, in *GetShelvesRequest, opts ...grpc.CallOption) (*GetShelvesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetShelvesResponse)
	err := c.cc.Invoke(ctx, ShelveService_GetShelves_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShelfResponse)
	err := c.cc.Invoke(ctx, ShelveService_CreateShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) UpdateShelf(ctx context.Context, in *UpdateShelfRequest, opts ...grpc.CallOption) (*UpdateShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShelfResponse)
	err := c.cc.Invoke(ctx, ShelveService_UpdateShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) RemoveShelf(ctx context.Context, in *RemoveShelfRequest, opts ...grpc.CallOption) (*RemoveShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveShelfResponse)
	err := c.cc.Invoke(ctx, ShelveService_RemoveShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelveServiceClient) GetShelf(ctx context.Context, in *RemoveShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateShelfResponse)
	err := c.cc.Invoke(ctx, ShelveService_GetShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShelveServiceServer is the server API for ShelveService service.
// All implementations must embed UnimplementedShelveServiceServer
// for forward compatibility
type ShelveServiceServer interface {
	GetShelves(context.Context, *GetShelvesRequest) (*GetShelvesResponse, error)
	CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error)
	UpdateShelf(context.Context, *UpdateShelfRequest) (*UpdateShelfResponse, error)
	RemoveShelf(context.Context, *RemoveShelfRequest) (*RemoveShelfResponse, error)
	GetShelf(context.Context, *RemoveShelfRequest) (*CreateShelfResponse, error)
	mustEmbedUnimplementedShelveServiceServer()
}

// UnimplementedShelveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShelveServiceServer struct {
}

func (UnimplementedShelveServiceServer) GetShelves(context.Context, *GetShelvesRequest) (*GetShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelves not implemented")
}
func (UnimplementedShelveServiceServer) CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (UnimplementedShelveServiceServer) UpdateShelf(context.Context, *UpdateShelfRequest) (*UpdateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShelf not implemented")
}
func (UnimplementedShelveServiceServer) RemoveShelf(context.Context, *RemoveShelfRequest) (*RemoveShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveShelf not implemented")
}
func (UnimplementedShelveServiceServer) GetShelf(context.Context, *RemoveShelfRequest) (*CreateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (UnimplementedShelveServiceServer) mustEmbedUnimplementedShelveServiceServer() {}

// UnsafeShelveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShelveServiceServer will
// result in compilation errors.
type UnsafeShelveServiceServer interface {
	mustEmbedUnimplementedShelveServiceServer()
}

func RegisterShelveServiceServer(s grpc.ServiceRegistrar, srv ShelveServiceServer) {
	s.RegisterService(&ShelveService_ServiceDesc, srv)
}

func _ShelveService_GetShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).GetShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_GetShelves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).GetShelves(ctx, req.(*GetShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_CreateShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_UpdateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).UpdateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_UpdateShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).UpdateShelf(ctx, req.(*UpdateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_RemoveShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).RemoveShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_RemoveShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).RemoveShelf(ctx, req.(*RemoveShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelveService_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelveServiceServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShelveService_GetShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelveServiceServer).GetShelf(ctx, req.(*RemoveShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShelveService_ServiceDesc is the grpc.ServiceDesc for ShelveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShelveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ShelveService",
	HandlerType: (*ShelveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShelves",
			Handler:    _ShelveService_GetShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _ShelveService_CreateShelf_Handler,
		},
		{
			MethodName: "UpdateShelf",
			Handler:    _ShelveService_UpdateShelf_Handler,
		},
		{
			MethodName: "RemoveShelf",
			Handler:    _ShelveService_RemoveShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _ShelveService_GetShelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shelve.proto",
}

const (
	BookShelfService_AddBook_FullMethodName         = "/BookShelfService/AddBook"
	BookShelfService_RemoveBook_FullMethodName      = "/BookShelfService/RemoveBook"
	BookShelfService_SearchBookShelf_FullMethodName = "/BookShelfService/SearchBookShelf"
)

// BookShelfServiceClient is the client API for BookShelfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookShelfServiceClient interface {
	AddBook(ctx context.Context, in *CreateBookShelf, opts ...grpc.CallOption) (*BookShelfResponse, error)
	RemoveBook(ctx context.Context, in *CreateBookShelf, opts ...grpc.CallOption) (*RemoveBookShelfResponse, error)
	SearchBookShelf(ctx context.Context, in *BookShelf, opts ...grpc.CallOption) (*BookShelvesResponse, error)
}

type bookShelfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookShelfServiceClient(cc grpc.ClientConnInterface) BookShelfServiceClient {
	return &bookShelfServiceClient{cc}
}

func (c *bookShelfServiceClient) AddBook(ctx context.Context, in *CreateBookShelf, opts ...grpc.CallOption) (*BookShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookShelfResponse)
	err := c.cc.Invoke(ctx, BookShelfService_AddBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfServiceClient) RemoveBook(ctx context.Context, in *CreateBookShelf, opts ...grpc.CallOption) (*RemoveBookShelfResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveBookShelfResponse)
	err := c.cc.Invoke(ctx, BookShelfService_RemoveBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfServiceClient) SearchBookShelf(ctx context.Context, in *BookShelf, opts ...grpc.CallOption) (*BookShelvesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookShelvesResponse)
	err := c.cc.Invoke(ctx, BookShelfService_SearchBookShelf_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookShelfServiceServer is the server API for BookShelfService service.
// All implementations must embed UnimplementedBookShelfServiceServer
// for forward compatibility
type BookShelfServiceServer interface {
	AddBook(context.Context, *CreateBookShelf) (*BookShelfResponse, error)
	RemoveBook(context.Context, *CreateBookShelf) (*RemoveBookShelfResponse, error)
	SearchBookShelf(context.Context, *BookShelf) (*BookShelvesResponse, error)
	mustEmbedUnimplementedBookShelfServiceServer()
}

// UnimplementedBookShelfServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookShelfServiceServer struct {
}

func (UnimplementedBookShelfServiceServer) AddBook(context.Context, *CreateBookShelf) (*BookShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedBookShelfServiceServer) RemoveBook(context.Context, *CreateBookShelf) (*RemoveBookShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBook not implemented")
}
func (UnimplementedBookShelfServiceServer) SearchBookShelf(context.Context, *BookShelf) (*BookShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBookShelf not implemented")
}
func (UnimplementedBookShelfServiceServer) mustEmbedUnimplementedBookShelfServiceServer() {}

// UnsafeBookShelfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookShelfServiceServer will
// result in compilation errors.
type UnsafeBookShelfServiceServer interface {
	mustEmbedUnimplementedBookShelfServiceServer()
}

func RegisterBookShelfServiceServer(s grpc.ServiceRegistrar, srv BookShelfServiceServer) {
	s.RegisterService(&BookShelfService_ServiceDesc, srv)
}

func _BookShelfService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookShelf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookShelfService_AddBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServiceServer).AddBook(ctx, req.(*CreateBookShelf))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookShelfService_RemoveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookShelf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServiceServer).RemoveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookShelfService_RemoveBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServiceServer).RemoveBook(ctx, req.(*CreateBookShelf))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookShelfService_SearchBookShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookShelf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServiceServer).SearchBookShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookShelfService_SearchBookShelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServiceServer).SearchBookShelf(ctx, req.(*BookShelf))
	}
	return interceptor(ctx, in, info, handler)
}

// BookShelfService_ServiceDesc is the grpc.ServiceDesc for BookShelfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookShelfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookShelfService",
	HandlerType: (*BookShelfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _BookShelfService_AddBook_Handler,
		},
		{
			MethodName: "RemoveBook",
			Handler:    _BookShelfService_RemoveBook_Handler,
		},
		{
			MethodName: "SearchBookShelf",
			Handler:    _BookShelfService_SearchBookShelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shelve.proto",
}
