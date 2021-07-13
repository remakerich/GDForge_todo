package main

import (
	"log"
	"net"

	"example.com/GDForge/server/methods"
	"example.com/GDForge/todo"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	log.Println("Server successfully started. Listening on port 9000")

	grpcServer := grpc.NewServer()
	todo.RegisterTodoServiceServer(grpcServer, &methods.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000: %v", err)
	}
}
