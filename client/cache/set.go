package cache

import (
	"context"

	"github.com/bsandusky/protocache/pb"
)

func Set(key, value string) (string, error) {
	res, err := client.Set(context.Background(), &pb.SetRequest{
		Key:   key,
		Value: []byte(value),
	})
	if err != nil {
		return "", err
	}
	return res.Status, nil
}
