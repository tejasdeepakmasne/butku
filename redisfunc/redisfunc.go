package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var Ctx = context.Background()

func InitializeRDB(dbNo int) *redis.Client {
	loadenv()
	redis_addr := os.Getenv("REDIS_ADDR")
	redis_pass := os.Getenv("REDIS_PASS")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
		DB:       dbNo,
	})

	return rdb

}

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error Loading env file %v", err)
	}
}
