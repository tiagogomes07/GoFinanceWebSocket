package mycustomhandler

import (
	"GoSocket/tickerdetails"
	"GoSocket/userauthentication"

	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func GetStock(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Println("getStock")

	for {

		if err != nil {
			log.Println("read:", err)
			break
		}
		mt, userMessage, err := c.ReadMessage()

		log.Println("reciving ticker " + string(userMessage))

		userObj := string(userMessage)
		userSession, err := userauthentication.GetSession(userObj)

		log.Println("user logged", userSession.UserID, "try getting", userSession.Ticker)

		if err != nil {
			//log.Println(err.Error())
			c.WriteMessage(1, []byte("User not connected"))
			defer c.Close()
			return
		}

		go tickerdetails.SendStocksPrices(c, mt, string(userMessage))

		if err != nil {
			log.Println("write:", err)
			break
		}
		log.Println("finalized for")
	}
	log.Println("finalized")
}
