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
func StoreSession(userid string) bool {
	redisClient := redisacess.GetRedisClient()
	//ctx := context.Background()
	user := model.UserSession{
		UserID: userid,
	}
	data, _ := json.Marshal(user)

	fmt.Println(data)
	//var ctx = context.Background()
	//defer ctx.Done()
	redisClient.Set(userid, string(data), 10*time.Minute)
	defer redisClient.Close()
	return true
}

//StartedSockectGetPrice ok
func StartedSockectGetPrice(userid string) {
	redisClient := redisacess.GetRedisClient()
	user := model.UserSession{
		UserID:           userid,
		ScoketGetPriceOn: true,
	}
	data, _ := json.Marshal(user)
	defer redisClient.Close()
	redisClient.Set(userid, data, 10*time.Minute)
}

func StopSockectGetPrice(userID string) {
	redisClient := redisacess.GetRedisClient()
	user := model.UserSession{
		UserID:           userID,
		ScoketGetPriceOn: false,
	}
	data, _ := json.Marshal(user)
	defer redisClient.Close()
	redisClient.Set(userID, data, 10*time.Minute)
}

//GetSession is used to retrive the user session
func GetSession(userRequest string) (*model.UserSession, error) {

	user := &model.UserSession{}
	json.Unmarshal([]byte(userRequest), user)
	redisClient := redisacess.GetRedisClient()
	defer redisClient.Close()

	fmt.Println("trying get userID" + user.UserID)
	res := redisClient.Get(user.UserID)

	userString := res.Val()

	if userString == "" {
		err := errors.New("User not logged in")
		fmt.Println("User not logged in")
		return user, err
	} else {
		fmt.Println("user loged")
		fmt.Println("json returned", userString)
		json.Unmarshal([]byte(userString), user)
	}

	return user, nil
}

//CloseSession to clean user session
func CloseSession(userID string) {
	redisClient := redisacess.GetRedisClient()
	redisClient.Del(userID)
	defer redisClient.Close()
}
