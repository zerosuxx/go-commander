package main

import (
	"github.com/zerosuxx/go-commander/handler"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
)

var Version = "development"

func main() {
	port := "1234"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Println("Http c0mmander " + Version + " | Server listening on: http://localhost:" + port)

	http.HandleFunc("/healthcheck", handler.CreateHealthCheckHandler().Handle)
	log.Println("GET /healthcheck")

	http.HandleFunc("/cmd", handler.CreateCommandHandler().Handle)
	log.Println("POST /cmd")

	http.Handle("/echo", websocket.Handler(handler.ShellServer))
	log.Println("WS /echo")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
