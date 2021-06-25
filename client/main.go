package main

import (
	"context"
	"log"
	"time"

	pb "bb.com/todo/grpc/todo"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
)

const address = "localhost:5000"

func createTask(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	title, desc := uuid.New(), uuid.New()

	resp, err := client.AddTask(ctx, &pb.AddTaskRequest{Title: title, Description: desc})
	if err != nil {
		log.Println("Error creating task:", err)
		return
	}

	log.Println("Created task:", resp.Id)
}

func getTasks(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetTasks(ctx, &pb.GetTasksRequest{})
	if err != nil {
		log.Println("Error creating task:", err)
		return
	}

	log.Println(resp.Tasks)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	for i := 0; i < 100; i++ {
		go createTask(c)
		go getTasks(c)
	}

	time.Sleep(5 * time.Second)
}
