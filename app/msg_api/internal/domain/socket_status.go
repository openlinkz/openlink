package domain

import "github.com/go-redis/redis/v8"

type socketStatus struct {
	rdb *redis.Client
}
