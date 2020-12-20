package tickerdetails

import (
	"GoSocket/model"
	"GoSocket/redisacess"
	"encoding/json"

	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

//GetTickerCurrentPrice get current price
func GetTickerCurrentPrice(userRequest string) model.RequestTicker {

	requestTicker := model.RequestTicker{}

	json.Unmarshal([]byte(userRequest), &requestTicker)

	redisClient := redisacess.GetRedisClient()

	requestedTicker := redisClient.Get(requestTicker.Ticker).Val()

	json.Unmarshal([]byte(string(requestedTicker)), &requestTicker)

	defer redisClient.Close()

	return requestTicker
}

//SendStocksPrices ok
func SendStocksPrices(c *websocket.Conn, mt int, strMessage string) {
	for true {
		rdn := rand.Intn(50-30) + 30
		concat := fmt.Sprint(strMessage, " price: ", rdn)

		res := []byte(concat)
		err := c.WriteMessage(mt, res)
		time.Sleep(3 * time.Second)
		fmt.Println(concat)
		if err != nil {
			fmt.Println("error sending message, then stop")
			defer c.Close()
			return
		}
	}
}
