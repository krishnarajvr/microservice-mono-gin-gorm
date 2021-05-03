package db

import (
	redis "github.com/go-redis/redis/v8"

	"micro/config"
	"micro/util/cache"
)

func New(conf *config.Conf) (*cache.RedisClient, error) {
	if len(conf.Cache.Host) == 0 {
		return nil, nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	redisCache := cache.New(rdb)

	return redisCache, nil
}
