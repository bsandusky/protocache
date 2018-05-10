package cache

import (
	"context"
	"errors"

	"github.com/bsandusky/protocache/pb"
)

// Server is gRPC cache server object
type Server struct{}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, errors.New("Not implemented")
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	return &pb.SetResponse{}, errors.New("Not implemented")
}

func (s *Server) FlushAll(ctx context.Context, req *pb.Empty) (*pb.FlushResponse, error) {
	return &pb.FlushResponse{}, errors.New("Not implemented")
}

func (s *Server) FlushKey(ctx context.Context, req *pb.FlushKeyRequest) (*pb.FlushResponse, error) {
	return &pb.FlushResponse{}, errors.New("Not implemented")
}
