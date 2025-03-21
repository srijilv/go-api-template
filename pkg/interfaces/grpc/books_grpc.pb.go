// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/grpc/books.proto

package grpc

import (
	context "context"
	list_books "github.com/srijilv/go-api-template.git/pkg/interfaces/grpc/list_books"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BooksServiceClient is the client API for BooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksServiceClient interface {
	Listbooks(ctx context.Context, in *list_books.ListBooksRequest, opts ...grpc.CallOption) (*list_books.ListBooksResponse, error)
}

type booksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksServiceClient(cc grpc.ClientConnInterface) BooksServiceClient {
	return &booksServiceClient{cc}
}

func (c *booksServiceClient) Listbooks(ctx context.Context, in *list_books.ListBooksRequest, opts ...grpc.CallOption) (*list_books.ListBooksResponse, error) {
	out := new(list_books.ListBooksResponse)
	err := c.cc.Invoke(ctx, "/grpc.BooksService/Listbooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BooksServiceServer is the server API for BooksService service.
// All implementations must embed UnimplementedBooksServiceServer
// for forward compatibility
type BooksServiceServer interface {
	Listbooks(context.Context, *list_books.ListBooksRequest) (*list_books.ListBooksResponse, error)
	mustEmbedUnimplementedBooksServiceServer()
}

// UnimplementedBooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServiceServer struct {
}

func (UnimplementedBooksServiceServer) Listbooks(context.Context, *list_books.ListBooksRequest) (*list_books.ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Listbooks not implemented")
}
func (UnimplementedBooksServiceServer) mustEmbedUnimplementedBooksServiceServer() {}

// UnsafeBooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServiceServer will
// result in compilation errors.
type UnsafeBooksServiceServer interface {
	mustEmbedUnimplementedBooksServiceServer()
}

func RegisterBooksServiceServer(s grpc.ServiceRegistrar, srv BooksServiceServer) {
	s.RegisterService(&BooksService_ServiceDesc, srv)
}

func _BooksService_Listbooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(list_books.ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BooksServiceServer).Listbooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.BooksService/Listbooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BooksServiceServer).Listbooks(ctx, req.(*list_books.ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BooksService_ServiceDesc is the grpc.ServiceDesc for BooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.BooksService",
	HandlerType: (*BooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Listbooks",
			Handler:    _BooksService_Listbooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/books.proto",
}
