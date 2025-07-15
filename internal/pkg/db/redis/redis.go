package redis

import (
	"context"
	"fmt"
	"log"

	"hackathon-dsm-server/internal/pkg/config"

	"github.com/redis/go-redis/v9"
)

func NewConnection() *redis.Client {
	cfg := config.GlobalConfig.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis")
	return client
}