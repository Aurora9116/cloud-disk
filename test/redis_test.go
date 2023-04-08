package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var rdb = NewRedisClient()

func TestSetValue(t *testing.T) {
	err := rdb.Set(context.Background(), "key", "value", time.Second*10).Err()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetValue(t *testing.T) {
	result, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
