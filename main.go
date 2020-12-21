package main

import (
	"GoSocket/mycustomhandler"
	"flag"
	"log"
	"net/http"

	"fmt"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/getstock", mycustomhandler.GetStock)
	http.HandleFunc("/auth", mycustomhandler.Authentication)
	http.HandleFunc("/", index)
	fmt.Println("iniciando")

	port := "5000"
	fmt.Println("Port?", port)
	//host := fmt.Sprint("localhost:", port)

	//addr := flag.String("addr", host, "http service address")
	log.Fatal(http.ListenAndServe(port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}
