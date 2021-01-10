package main

import (
	"GoSocket/mycustomhandler"
	"flag"
	"log"
	"net/http"
	"os"

	"fmt"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/getstock", mycustomhandler.GetStock)
	http.HandleFunc("/auth", mycustomhandler.Authentication)
	http.HandleFunc("/", index)
	fmt.Println("iniciando")

	port := os.Getenv("PORT")
	fmt.Println("Port?", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
