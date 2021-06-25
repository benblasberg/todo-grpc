package server

import (
	"context"

	pb "bb.com/todo/grpc/todo"
	"gopkg.in/alexcesaro/statsd.v2"
)

const (
	addTaskSuccess  = "todoservergrpc.addTask.success"
	addTaskTimer    = "todoservergrpc.addTask.timing"
	getTasksSuccess = "todoservergrpc.getTasks.success"
	getTasksTimer   = "todoservergrpc.getTasks.timing"
)

type metrics struct {
	statsd *statsd.Client
	next   pb.TodoServiceServer
	pb.UnimplementedTodoServiceServer
}

func NewMetricsServer(statsd *statsd.Client, next pb.TodoServiceServer) pb.TodoServiceServer {
	return &metrics{statsd: statsd, next: next}
}

func (s metrics) AddTask(ctx context.Context, in *pb.AddTaskRequest) (resp *pb.AddTaskResponse, err error) {
	timer := s.statsd.NewTiming()

	resp, err = s.next.AddTask(ctx, in)

	if err == nil {
		s.statsd.Increment(addTaskSuccess)
	}

	timer.Send(addTaskTimer)

	return resp, err
}

func (s metrics) GetTasks(ctx context.Context, in *pb.GetTasksRequest) (resp *pb.GetTasksResponse, err error) {
	timer := s.statsd.NewTiming()

	resp, err = s.next.GetTasks(ctx, in)

	if err == nil {
		s.statsd.Increment(getTasksSuccess)
	}

	timer.Send(getTasksTimer)

	return resp, err
}
