package iface

import "context"

type Counter interface {
	Incr(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (int64, error)
	Set(ctx context.Context, key string, value int64) error
	Delete(ctx context.Context, key string) error
}
