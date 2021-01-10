package redisacess

import (
	"os"

	"sync"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

var (
	instance *redis.Client
)

//GetRedisClient
func GetRedisClient() *redis.Client {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		println("instance redis client")
		godotenv.Load()

		host := os.Getenv("REDISHOST")
		port := os.Getenv("REDISPORT")
		password := os.Getenv("REDISPASSWORD")

		instance = redis.NewClient(&redis.Options{
			Addr:     host + ":" + port,
			Password: password,
			DB:       0,
			//DialTimeout: 10 * time.Minute,
		})

	}

	return instance
}
