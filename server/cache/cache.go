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

// Get resturns value for given key
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Value: s.Store[req.Key],
		Result: &pb.Result{
			Status: "OK",
			Error:  "",
		},
	}, nil
}

// Set creates entry in server store with key and value from request
func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.Result, error) {
	s.Store[req.Key] = req.Value
	return &pb.Result{Status: "OK", Error: ""}, nil
}

// FlushAll clears cache
func (s *Server) FlushAll(ctx context.Context, req *pb.Empty) (*pb.Result, error) {
	for k := range s.Store {
		delete(s.Store, k)
	}
	return &pb.Result{Status: "OK", Error: ""}, nil
}

// FlushKey clears given key in server store
func (s *Server) FlushKey(ctx context.Context, req *pb.FlushKeyRequest) (*pb.Result, error) {
	delete(s.Store, req.Key)
	return &pb.Result{Status: "OK", Error: ""}, nil
}

// Listener provides stream of data activity for monitoring and logging
func (s *Server) Listener(req *pb.Empty, stream pb.Cache_ListenerServer) error {
	return errors.New("Not implemented")
}
