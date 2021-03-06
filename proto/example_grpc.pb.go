// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: github.com/53jk1/todo-task/proto/example.proto

package example

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	GetAllTodos(ctx context.Context, in *GetAllTodosRequest, opts ...grpc.CallOption) (*GetAllTodosResponse, error)
	GetSpecificTodo(ctx context.Context, in *GetSpecificTodoRequest, opts ...grpc.CallOption) (*GetSpecificTodoResponse, error)
	GetIncomingTodo(ctx context.Context, in *GetIncomingTodoRequest, opts ...grpc.CallOption) (*GetIncomingTodoResponse, error)
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error)
	SetTodoPercentComplete(ctx context.Context, in *SetTodoPercentCompleteRequest, opts ...grpc.CallOption) (*SetTodoPercentCompleteResponse, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
	MarkTodoAsComplete(ctx context.Context, in *MarkTodoAsCompleteRequest, opts ...grpc.CallOption) (*MarkTodoAsCompleteResponse, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) GetAllTodos(ctx context.Context, in *GetAllTodosRequest, opts ...grpc.CallOption) (*GetAllTodosResponse, error) {
	out := new(GetAllTodosResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/GetAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetSpecificTodo(ctx context.Context, in *GetSpecificTodoRequest, opts ...grpc.CallOption) (*GetSpecificTodoResponse, error) {
	out := new(GetSpecificTodoResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/GetSpecificTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetIncomingTodo(ctx context.Context, in *GetIncomingTodoRequest, opts ...grpc.CallOption) (*GetIncomingTodoResponse, error) {
	out := new(GetIncomingTodoResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/GetIncomingTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*UpdateTodoResponse, error) {
	out := new(UpdateTodoResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) SetTodoPercentComplete(ctx context.Context, in *SetTodoPercentCompleteRequest, opts ...grpc.CallOption) (*SetTodoPercentCompleteResponse, error) {
	out := new(SetTodoPercentCompleteResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/SetTodoPercentComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) MarkTodoAsComplete(ctx context.Context, in *MarkTodoAsCompleteRequest, opts ...grpc.CallOption) (*MarkTodoAsCompleteResponse, error) {
	out := new(MarkTodoAsCompleteResponse)
	err := c.cc.Invoke(ctx, "/example.TodoService/MarkTodoAsComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	GetAllTodos(context.Context, *GetAllTodosRequest) (*GetAllTodosResponse, error)
	GetSpecificTodo(context.Context, *GetSpecificTodoRequest) (*GetSpecificTodoResponse, error)
	GetIncomingTodo(context.Context, *GetIncomingTodoRequest) (*GetIncomingTodoResponse, error)
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error)
	SetTodoPercentComplete(context.Context, *SetTodoPercentCompleteRequest) (*SetTodoPercentCompleteResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	MarkTodoAsComplete(context.Context, *MarkTodoAsCompleteRequest) (*MarkTodoAsCompleteResponse, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) GetAllTodos(context.Context, *GetAllTodosRequest) (*GetAllTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoServiceServer) GetSpecificTodo(context.Context, *GetSpecificTodoRequest) (*GetSpecificTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecificTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetIncomingTodo(context.Context, *GetIncomingTodoRequest) (*GetIncomingTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncomingTodo not implemented")
}
func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodo(context.Context, *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoServiceServer) SetTodoPercentComplete(context.Context, *SetTodoPercentCompleteRequest) (*SetTodoPercentCompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTodoPercentComplete not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoServiceServer) MarkTodoAsComplete(context.Context, *MarkTodoAsCompleteRequest) (*MarkTodoAsCompleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkTodoAsComplete not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/GetAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetAllTodos(ctx, req.(*GetAllTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetSpecificTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetSpecificTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/GetSpecificTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetSpecificTodo(ctx, req.(*GetSpecificTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetIncomingTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncomingTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetIncomingTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/GetIncomingTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetIncomingTodo(ctx, req.(*GetIncomingTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_SetTodoPercentComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTodoPercentCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).SetTodoPercentComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/SetTodoPercentComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).SetTodoPercentComplete(ctx, req.(*SetTodoPercentCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_MarkTodoAsComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkTodoAsCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).MarkTodoAsComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.TodoService/MarkTodoAsComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).MarkTodoAsComplete(ctx, req.(*MarkTodoAsCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoService_GetAllTodos_Handler,
		},
		{
			MethodName: "GetSpecificTodo",
			Handler:    _TodoService_GetSpecificTodo_Handler,
		},
		{
			MethodName: "GetIncomingTodo",
			Handler:    _TodoService_GetIncomingTodo_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "SetTodoPercentComplete",
			Handler:    _TodoService_SetTodoPercentComplete_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "MarkTodoAsComplete",
			Handler:    _TodoService_MarkTodoAsComplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/53jk1/todo-task/proto/example.proto",
}
