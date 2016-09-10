package main

import (
	"log"
	"net/http"

	"github.com/cyberroadie/webservices-spike/server"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// websocket server
	server := server.NewServer("/entry")
	go server.Listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot/public")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
