package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// Set assigns a key, value pair in the cache
func Set(key, value string) (map[string]string, error) {
	data := make(map[string]string)
	res, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		return nil, err
	}
	data["status"] = res.Result

	return data, nil
}
