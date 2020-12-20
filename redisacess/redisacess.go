package redisacess

import (
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

//GetRedisClient here
func GetRedisClient() redis.Client {

	var redisClient *redis.Client
	//var ctx = context.Background()
	godotenv.Load()

	host := os.Getenv("REDISHOST")
	port := os.Getenv("REDISPORT")
	password := os.Getenv("REDISPASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:        host + ":" + port,
		Password:    password,
		DB:          0,
		DialTimeout: 10 * time.Minute,
	})

	return *redisClient
}
