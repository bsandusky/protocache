package cache

import (
	"context"
	"errors"

	"github.com/bsandusky/protocache/pb"
)

// Get returns map with key and value for given key
func Get(key string) (*pb.GetResponse, error) {
	res, err := client.Get(context.Background(), &pb.GetRequest{Key: key})
	if err != nil {
		return nil, err
	}

	if res.Value == "" {
		return nil, errors.New("Key not recognized")
	}

	return res, nil
}

// GetAll returns map with keys and values for all items in the cache
func GetAll() (*pb.GetAllResponse, error) {
	res, err := client.GetAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	if len(res.GetResponse) == 0 {
		return nil, errors.New("Empty cache")
	}

	return res, nil
}
