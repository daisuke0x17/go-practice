package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const lockPrefix = "Lock:"

type Redis struct {
	client redis.UniversalClient
}

func NewClient() *Redis {
	options := &redis.UniversalOptions{
		Addrs:    []string{"localhost:6379"},
		Password: "",
		DB:       0,
		PoolSize: 1000,
	}
	return &Redis{redis.NewClient(options.Simple())}
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func (r *Redis) Lock(ctx context.Context, key string, timeout time.Duration) (func(ctx context.Context), error) {
	lockKey := lockPrefix + key
	b, err := r.client.SetNX(ctx, lockKey, key, timeout).Result()
	if err != nil {
		return nil, err
	}
	if !b {
		return nil, fmt.Errorf("%s is already locked", key)
	}

	fmt.Printf("%s is locked\n", key)
	unlock := func(ctx context.Context) {
		fmt.Printf("%s is unlocked\n", key)
		r.client.Del(ctx, lockKey)
	}
	return unlock, nil

}

func main() {
	ctx := context.Background()
	redisClient := NewClient()
	defer redisClient.Close()

	key := "myKey"
	unlock, err := redisClient.Lock(ctx, key, time.Second*10)
	if err != nil {
		fmt.Print(err)
		return
	}

	unlock(ctx)
}
