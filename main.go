package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"

	"github.com/gorilla/mux"

	"github.com/rs/cors"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			polling.Default,
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})

	if err != nil {
		fmt.Println("error ini")
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	mux := mux.NewRouter()

	mux.Handle("/socket.io/", server)
	mux.HandleFunc("/", index)

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))

	// 	//index(w, r)
	// })

	//

	// log.Println("Serving at localhost:PORT " + PORT)
	// //cors.AllowAll().Handler(mux))
	// log.Fatal(http.ListenAndServe(PORT, mux))

	handler := cors.Default().Handler(mux)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler = c.Handler(handler)

	godotenv.Load(".env")
	PORT := ":" + os.Getenv("PORT")

	log.Println("Serving at localhost" + PORT + "...")
	log.Fatal(http.ListenAndServe(PORT, handler))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
