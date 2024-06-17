package kvsore

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type KVstore struct {
	client *redis.Client
	db     int
}

func GetKVStore(addr, password string, database int) KVStoreI {
	return &KVstore{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       database,
		}),
		db: database,
	}
}

// Del implements KVStoreI.
func (k *KVstore) Del(ctx context.Context, key string) error {
	return k.client.Del(ctx, "key").Err()
}

// Get implements KVStoreI.
func (k *KVstore) Get(ctx context.Context, key string) (string, error) {
	return k.client.Get(ctx, key).Result()
}

// Set implements KVStoreI.
func (k *KVstore) Set(ctx context.Context, key string, val *string) error {
	return k.client.Set(ctx, key, &val, 0).Err()
}
