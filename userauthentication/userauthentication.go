package userauthentication

import (
	"GoSocket/model"
	"GoSocket/redisacess"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	//	"github.com/go-redis/redis/v8"

	//	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//var redisClient *redis.Client

//UserValid is true
func UserValid(userID string) bool {
	return true
}

//StoreSession is used to store user session
func StoreSession(userID string) bool {
	redisClient := redisacess.GetRedisClient()
	//ctx := context.Background()
	user := model.UserSession{
		UserId: userID,
	}
	data, _ := json.Marshal(user)

	fmt.Println(data)
	//var ctx = context.Background()
	//defer ctx.Done()
	redisClient.Set(userID, string(data), 10*time.Minute)
	defer redisClient.Close()
	return true
}

//StartedSockectGetPrice ok
func StartedSockectGetPrice(userID string) {
	redisClient := redisacess.GetRedisClient()
	user := model.UserSession{
		UserId:           userID,
		ScoketGetPriceOn: true,
	}
	data, _ := json.Marshal(user)
	defer redisClient.Close()
	redisClient.Set(userID, data, 10*time.Minute)
}

func StopSockectGetPrice(userID string) {
	redisClient := redisacess.GetRedisClient()
	user := model.UserSession{
		UserId:           userID,
		ScoketGetPriceOn: false,
	}
	data, _ := json.Marshal(user)
	defer redisClient.Close()
	redisClient.Set(userID, data, 10*time.Minute)
}

//GetSession is used to retrive the user session
func GetSession(userID string) (*model.UserSession, error) {
	redisClient := redisacess.GetRedisClient()
	res := redisClient.Get(userID)

	user := &model.UserSession{}
	var err error
	userString := res.Val()
	if userString == "" {
		err := errors.New("User not logged in")
		return user, err
	}

	err = json.Unmarshal([]byte(userString), user)
	defer redisClient.Close()
	return user, err
}

//CloseSession to clean user session
func CloseSession(userID string) {
	redisClient := redisacess.GetRedisClient()
	redisClient.Del(userID)
	defer redisClient.Close()
}
