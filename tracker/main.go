package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
}
