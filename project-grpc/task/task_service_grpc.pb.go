// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: task_service.proto

package task

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
	TaskService_TaskStages_FullMethodName        = "/task.service.v1.TaskService/TaskStages"
	TaskService_MemberProjectList_FullMethodName = "/task.service.v1.TaskService/MemberProjectList"
	TaskService_TaskList_FullMethodName          = "/task.service.v1.TaskService/TaskList"
	TaskService_SaveTask_FullMethodName          = "/task.service.v1.TaskService/SaveTask"
	TaskService_TaskSort_FullMethodName          = "/task.service.v1.TaskService/TaskSort"
	TaskService_MyTaskList_FullMethodName        = "/task.service.v1.TaskService/MyTaskList"
)

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error)
	MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error)
	TaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskListResponse, error)
	SaveTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error)
	TaskSort(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSortResponse, error)
	MyTaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MyTaskListResponse, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) TaskStages(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskStagesResponse, error) {
	out := new(TaskStagesResponse)
	err := c.cc.Invoke(ctx, TaskService_TaskStages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) MemberProjectList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MemberProjectResponse, error) {
	out := new(MemberProjectResponse)
	err := c.cc.Invoke(ctx, TaskService_MemberProjectList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, TaskService_TaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) SaveTask(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskMessage, error) {
	out := new(TaskMessage)
	err := c.cc.Invoke(ctx, TaskService_SaveTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) TaskSort(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*TaskSortResponse, error) {
	out := new(TaskSortResponse)
	err := c.cc.Invoke(ctx, TaskService_TaskSort_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) MyTaskList(ctx context.Context, in *TaskReqMessage, opts ...grpc.CallOption) (*MyTaskListResponse, error) {
	out := new(MyTaskListResponse)
	err := c.cc.Invoke(ctx, TaskService_MyTaskList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error)
	MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error)
	TaskList(context.Context, *TaskReqMessage) (*TaskListResponse, error)
	SaveTask(context.Context, *TaskReqMessage) (*TaskMessage, error)
	TaskSort(context.Context, *TaskReqMessage) (*TaskSortResponse, error)
	MyTaskList(context.Context, *TaskReqMessage) (*MyTaskListResponse, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) TaskStages(context.Context, *TaskReqMessage) (*TaskStagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskStages not implemented")
}
func (UnimplementedTaskServiceServer) MemberProjectList(context.Context, *TaskReqMessage) (*MemberProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberProjectList not implemented")
}
func (UnimplementedTaskServiceServer) TaskList(context.Context, *TaskReqMessage) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedTaskServiceServer) SaveTask(context.Context, *TaskReqMessage) (*TaskMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTask not implemented")
}
func (UnimplementedTaskServiceServer) TaskSort(context.Context, *TaskReqMessage) (*TaskSortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskSort not implemented")
}
func (UnimplementedTaskServiceServer) MyTaskList(context.Context, *TaskReqMessage) (*MyTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MyTaskList not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_TaskStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_TaskStages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskStages(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_MemberProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).MemberProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_MemberProjectList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).MemberProjectList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_TaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_SaveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).SaveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_SaveTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).SaveTask(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_TaskSort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).TaskSort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_TaskSort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).TaskSort(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_MyTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskReqMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).MyTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskService_MyTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).MyTaskList(ctx, req.(*TaskReqMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.service.v1.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TaskStages",
			Handler:    _TaskService_TaskStages_Handler,
		},
		{
			MethodName: "MemberProjectList",
			Handler:    _TaskService_MemberProjectList_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskService_TaskList_Handler,
		},
		{
			MethodName: "SaveTask",
			Handler:    _TaskService_SaveTask_Handler,
		},
		{
			MethodName: "TaskSort",
			Handler:    _TaskService_TaskSort_Handler,
		},
		{
			MethodName: "MyTaskList",
			Handler:    _TaskService_MyTaskList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task_service.proto",
}