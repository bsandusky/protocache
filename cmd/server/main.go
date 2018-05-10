package main

import (
	"log"
	"net"

	"github.com/bsandusky/protocache/pb"
	"google.golang.org/grpc"
)

func main() {
	// Server implementation
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCacheServer(grpcServer, &cache.Server{})
	grpcServer.Serve(lis)
}
