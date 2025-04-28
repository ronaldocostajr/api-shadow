// Middleware to database REDIS
package middleware

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func RedisInsert(redisKey string, redisValue string) bool {
	errD := godotenv.Load()
	if errD != nil {
		log.Fatalf("Error loading .env file: %s", errD)
	}

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_uri"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})

	err := client.Set(ctx, redisKey, redisValue, 0).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func RedisRead(redisKey string) string {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_uri"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})

	val, err := client.Get(ctx, redisKey).Result()
	if err != nil {
		return ""
	}
	return val
}

func RedisDelete(redisKey string) bool {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_uri"),
		Password: os.Getenv("redis_password"),
		DB:       0,
	})

	_, err := client.Del(ctx, redisKey).Result()
	if err != nil {
		panic(err)
		return false
	}
	return true
}
