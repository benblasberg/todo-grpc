package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "bb.com/todo/grpc/todo"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
)

const (
	address            = "localhost:5000"
	defaultNumRequests = 100
)

func createTask(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	title, desc := uuid.New(), uuid.New()

	_, err := client.AddTask(ctx, &pb.AddTaskRequest{Title: title, Description: desc})
	if err != nil {
		log.Println("Error creating task:", err)
		return
	}
}

func getTasks(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.GetTasks(ctx, &pb.GetTasksRequest{})
	if err != nil {
		log.Println("Error creating task:", err)
		return
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	numRequests := defaultNumRequests
	if len(os.Args) > 1 {
		numRequests, err = strconv.Atoi(os.Args[1])

		if err != nil {
			log.Print("Invalid number of requests", os.Args[1], err)
			numRequests = defaultNumRequests
		}
	}

	for i := 0; i < numRequests; i++ {
		go createTask(c)
		go getTasks(c)
		time.Sleep(100 * time.Millisecond)
		go getTasks(c)
	}

	time.Sleep(5 * time.Second)
}
