package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type RedisCacheCounter struct {
	client *redis.Client
}

func NewRedisCacheCounter(client *redis.Client) *RedisCacheCounter {
	return &RedisCacheCounter{
		client: client,
	}
}

func (c *RedisCacheCounter) Incr(ctx context.Context, key string) error {
	_, err := c.client.Incr(ctx, key).Result()

	if err == redis.Nil {
		log.WithFields(log.Fields{
			"key": key,
		}).Warn("Key does not exist, initializing to 1")
	}

	if err != nil {
		log.WithFields(log.Fields{
			"key": key,
			"err": err,
		}).Error("Failed to increment key in Redis")
	}

	return err
}

func (c *RedisCacheCounter) Get(ctx context.Context, key string) (int64, error) {
	val, err := c.client.Get(ctx, key).Int64()
	if err == redis.Nil {
		return 0, nil // Key does not exist
	}
	if err != nil {
		return 0, err // Some other error occurred
	}
	return val, nil
}

func (c *RedisCacheCounter) Set(ctx context.Context, key string, value int64) error {
	_, err := c.client.Set(ctx, key, value, 0).Result()
	return err
}

func (c *RedisCacheCounter) Delete(ctx context.Context, key string) error {
	_, err := c.client.Del(ctx, key).Result()
	return err
}
