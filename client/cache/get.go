package cache

import (
	"context"
	"errors"

	"github.com/bsandusky/protocache/pb"
)

// Get returns map with key and value for given key
func Get(key string) (map[string]string, error) {
	res, err := client.Get(context.Background(), &pb.GetRequest{Key: key})
	if err != nil {
		return nil, err
	}

	if res.Value == "" {
		return nil, errors.New("Key not recognized")
	}

	data := make(map[string]string)
	data[res.Key] = res.Value
	return data, nil
}

// GetAll returns map with keys and values for all items in the cache
func GetAll() (map[string]string, error) {
	res, err := client.GetAll(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, err
	}

	if len(res.GetResponse) == 0 {
		return nil, errors.New("Empty cache")
	}

	data := make(map[string]string)
	for _, r := range res.GetResponse {
		data[r.Key] = r.Value
	}
	return data, nil
}
