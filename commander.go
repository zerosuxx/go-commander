package main

import (
	"github.com/zerosuxx/go-http-commander/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	version := "1.0.1"
	port := "1234"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Println("Http c0mmander " + version + " | Server listening on: http://localhost:" + port)

	http.HandleFunc("/healthcheck", handler.CreateHealthCheckHandler().Handle)
	log.Println("GET /healthcheck")

	http.HandleFunc("/cmd", handler.CreateCommandHandler().Handle)
	log.Println("POST /cmd")

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
