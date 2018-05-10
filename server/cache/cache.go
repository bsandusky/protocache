package cache

import (
	"context"
	"errors"

	"github.com/bsandusky/protocache/pb"
)

// Server is gRPC cache server object
type Server struct {
	Store map[string][]byte
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, errors.New("Not implemented")
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.Result, error) {
	return &pb.Result{}, errors.New("Not implemented")
}

func (s *Server) FlushAll(ctx context.Context, req *pb.Empty) (*pb.Result, error) {
	return &pb.Result{}, errors.New("Not implemented")
}

func (s *Server) FlushKey(ctx context.Context, req *pb.FlushKeyRequest) (*pb.Result, error) {
	return &pb.Result{}, errors.New("Not implemented")
}

func (s *Server) Listener(req *pb.Empty, stream pb.Cache_ListenerServer) error {
	return errors.New("Not implemented")
}
