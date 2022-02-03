package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/picunada/http-rest-balance/internal/users"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	logging.Init()
	router := httprouter.New()
	usersHandler := users.NewHandler()
	//balanceHandler := users.NewHandler()

	log.Println("Register handlers")
	//balanceHandler.Register(router)
	usersHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("Start server")
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("server in listening on port 8001")
	log.Fatalln(server.Serve(listener))
}
