package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// Get returns map with key and value for given key
func Get(key string) (map[string]string, error) {
	data := make(map[string]string)
	res, err := client.Get(context.Background(), &pb.GetRequest{Key: []byte(key)})
	if err != nil {
		return nil, err
	}
	data[string(res.Key)] = string(res.Value)

	return data, nil
}

// GetAll returns map with keys and values for all items in the cache
func GetAll() (map[string]string, error) {
	data := make(map[string]string)
	res, err := client.GetAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	for _, r := range res.GetResponses {
		data[string(r.Key)] = string(r.Value)
	}

	return data, nil
}
