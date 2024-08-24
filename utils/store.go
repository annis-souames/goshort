package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func SetKey(ctx *context.Context, rdb *redis.Client, key string, value string, ttl int) {
	// We set the key value pair in Redis, we use the context defined in main by reference and a TTL of 0 (no expiration)
	rdb.Set(*ctx, key, value, 0)
	fmt.Println("The key", key, "has been set to", value, " successfully")
}

// This function retrieves the long URL from the short URL from Redis
func GetLongURL(ctx *context.Context, rdb *redis.Client, shortURL string) (string, error) {
	// We always use the context by reference from main.go to avoid creating a copy of the context
	longURL, err := rdb.Get(*ctx, shortURL).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("short URL not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve from Redis: %v", err)
	}
	return longURL, nil
}
