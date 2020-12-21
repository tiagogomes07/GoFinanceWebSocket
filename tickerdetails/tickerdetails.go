package tickerdetails

import (
	"GoSocket/model"
	"GoSocket/redisacess"
	"encoding/json"

	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

//GetTickerCurrentPrice get current price
func GetTickerCurrentPrice(userRequest string) (model.RequestTicker, error) {

	requestTicker := model.RequestTicker{}

	json.Unmarshal([]byte(userRequest), &requestTicker)

	redisClient := redisacess.GetRedisClient()

	requestedTicker := redisClient.Get(requestTicker.Ticker).Val()

	if requestedTicker == "" {
		err := errors.New("ticker unknow")
		return requestTicker, err
	}

	json.Unmarshal([]byte(string(requestedTicker)), &requestTicker)

	defer redisClient.Close()

	return requestTicker, nil
}

//SendStocksPrices ok
func SendStocksPrices(c *websocket.Conn, mt int, strMessage string) {
	fmt.Println("into go routine SendStocksPrices")
	for true {

		requestedTicker, errTicker := GetTickerCurrentPrice(strMessage)
		if errTicker != nil {
			fmt.Println(errTicker.Error())
			defer c.Close()
			return
		}

		obj, _ := json.Marshal(requestedTicker)

		err := c.WriteMessage(mt, obj)
		time.Sleep(3 * time.Second)
		//fmt.Println(concat)
		if err != nil {
			fmt.Println("error sending message, then stop")
			defer c.Close()
			return
		}
	}
}

func SendStocksPricesTest1(c *websocket.Conn, mt int, strMessage string) {
	fmt.Println("into go routine SendStocksPrices")
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
