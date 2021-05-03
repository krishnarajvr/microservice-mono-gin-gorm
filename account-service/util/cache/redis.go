package cache

import (
	"context"
	"encoding/json"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Valid   bool
	Client  *redis.Client
	Context context.Context
}

// Setup Initialize the Redis instance
func New(client *redis.Client) *RedisClient {
	redisClient := RedisClient{}
	ctx := context.Background()
	redisClient.Client = client
	redisClient.Context = ctx
	redisClient.Valid = true

	return &redisClient
}

// Set a key/value
func (c *RedisClient) Set(key string, data interface{}, duration time.Duration) error {
	value, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = c.Client.Set(c.Context, key, value, duration).Err()

	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func (c *RedisClient) Exists(key string) bool {
	return false
}

// Get get a key
func (c *RedisClient) Get(key string) (string, error) {
	val, err := c.Client.Get(c.Context, key).Result()

	switch {
	case err == redis.Nil:
		return "", nil
	case err != nil:
		return "", err
	}

	return val, nil
}

// Delete delete a kye
func (c *RedisClient) Delete(key string) (int64, error) {
	deleted, err := c.Client.Del(c.Context, key).Result()

	if err != nil {
		return 0, err
	}

	return deleted, nil
}

// LikeDeletes batch delete
func (c *RedisClient) LikeDeletes(key string) error {
	return nil
}

func (c *RedisClient) Ping() bool {
	_, err := c.Client.Ping(c.Context).Result()

	if err == nil {
		return true
	}

	return false
}
