package main

import (
	"log"
	"net"
	"os"

	server "bb.com/todo/grpc/server/todoserver"
	"bb.com/todo/grpc/server/todoserver/data"
	pb "bb.com/todo/grpc/todo"
	"google.golang.org/grpc"
	"gopkg.in/alexcesaro/statsd.v2"
)

const defaultPort = ":5000"

func main() {
	lis, err := net.Listen("tcp", defaultPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	statsd, err := statsd.New(statsd.Address(os.Getenv("STATSD")))
	if err != nil {
		log.Fatalf("Could not connect to statsd: %v", err)
	}

	taskServer := server.New(data.NewInMemoryDataStore())
	taskServer = server.NewMetricsServer(statsd, taskServer)

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, taskServer)
	log.Printf("Listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
