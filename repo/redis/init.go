package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var ctx = context.Background()

func InitCon() {
	redisAddr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	fmt.Printf("Redis Server %s ... \n", redisAddr)

	Client = redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     os.Getenv("REDIS_PASSWORD"), // empty string if no password
		DB:           0,                           // default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Ping Redis server to check connection
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("CONNECT REDIS FAILED: %s\n", err)
	}
}
