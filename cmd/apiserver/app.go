package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/picunada/http-rest-balance/internal/users"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()
	usersHandler := users.NewHandler(logger)
	//balanceHandler := users.NewHandler()

	logger.Info("Register handlers")
	//balanceHandler.Register(router)
	usersHandler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Start server")
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server in listening on port 8001")
	logger.Fatal(server.Serve(listener))
}
