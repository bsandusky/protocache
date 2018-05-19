package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

// Set assigns a key, value pair in the cache
func Set(key, value string) (map[string]string, error) {
	res, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	data["status"] = res.Result
	return data, nil
}
