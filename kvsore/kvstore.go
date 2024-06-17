package kvsore

import "context"

type KVStoreI interface {
	Set(ctx context.Context, key string, val *string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}
