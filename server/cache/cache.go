package cache

import (
	"context"
	"errors"

	"github.com/bsandusky/protocache/pb"
)

// Server is gRPC cache server object
type Server struct {
	Store map[string]string
}

// Get returns value for given key
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{
		Value: s.Store[req.Key],
	}, nil
}

// GetAll returns all cache keys and values
func (s *Server) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllResponse, error) {
	res := &pb.GetAllResponse{}
	for k, v := range s.Store {
		res.GetResponses = append(res.GetResponses, &pb.GetResponse{
			Key:   k,
			Value: v,
		})
	}
	return res, nil
}

// Set creates entry in server store with key and value from request
func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.Result, error) {
	s.Store[req.Key] = req.Value
	return &pb.Result{Result: "OK"}, nil
}

// FlushAll clears cache
func (s *Server) FlushAll(ctx context.Context, req *pb.Empty) (*pb.Result, error) {
	for k := range s.Store {
		delete(s.Store, k)
	}
	return &pb.Result{Result: "OK"}, nil
}

// FlushKey clears given key in server store
func (s *Server) FlushKey(ctx context.Context, req *pb.FlushKeyRequest) (*pb.Result, error) {
	delete(s.Store, req.Key)
	return &pb.Result{Result: "OK"}, nil
}

// Listener provides stream of data activity for monitoring and logging
func (s *Server) Listener(req *pb.Empty, stream pb.Cache_ListenerServer) error {
	return errors.New("Not implemented")
}
