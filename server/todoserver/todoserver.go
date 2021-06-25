package server

import (
	"context"
	"errors"

	"bb.com/todo/grpc/server/todoserver/data"
	pb "bb.com/todo/grpc/todo"
)

type server struct {
	store data.DataStore
	pb.UnimplementedTodoServiceServer
}

func New(store data.DataStore) pb.TodoServiceServer {
	return server{store: store}
}

var ErrInvalidArgument = errors.New("invalid argument received")

func (s server) AddTask(ctx context.Context, in *pb.AddTaskRequest) (*pb.AddTaskResponse, error) {
	if in.Title == "" || in.Description == "" {
		return nil, ErrInvalidArgument
	}

	t := data.Task{Title: in.Title, Description: in.Description}
	err := s.store.AddTodo(&t)

	if err != nil {
		return nil, err
	}

	return &pb.AddTaskResponse{Id: t.Id}, nil
}

func (s server) GetTasks(ctx context.Context, in *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	tasks, err := s.store.GetTasks()
	if err != nil {
		return nil, err
	}

	var respTasks []*pb.Task
	for _, t := range *tasks {
		respTasks = append(respTasks, &pb.Task{Id: t.Id, Title: t.Title, Description: t.Description})
	}

	return &pb.GetTasksResponse{Tasks: respTasks}, nil
}
