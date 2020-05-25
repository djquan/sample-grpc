package main

import (
	"github.com/djquan/skeleton/grpcservice/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	s := newServer()
	if err = s.Serve(lis); err != nil {
		log.Fatalf("oh nos %v", err)
	}
}

func newServer() *grpc.Server {
	s := grpc.NewServer()
	health.RegisterHealthServiceServer(s, health.NewHealthService())
	reflection.Register(s)
	return s
}
