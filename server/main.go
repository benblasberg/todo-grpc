package main

import (
	"log"
	"net"

	server "bb.com/todo/grpc/server/todoserver"
	"bb.com/todo/grpc/server/todoserver/data"
	pb "bb.com/todo/grpc/todo"
	"google.golang.org/grpc"
)

const defaultPort = ":5000"

func main() {
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, server.New(data.NewInMemoryDataStore()))
	log.Printf("Listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
