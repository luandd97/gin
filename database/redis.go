package database

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		DB:       0,
		Password: os.Getenv("REDIS_PASSWORD"),
		Addr:     getHost(),
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	return client
}

func getHost() string {
	return os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
}
