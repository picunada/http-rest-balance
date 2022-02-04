package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/picunada/http-rest-balance/internal/config"
	"github.com/picunada/http-rest-balance/internal/users"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create router")
	router := httprouter.New()

	cfg := config.GetConfig()
	logger.Info("Register handlers")
	usersHandler := users.NewHandler(logger)
	usersHandler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.ServerConfig) {
	logger := logging.GetLogger()
	logger.Info("Start server")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		// /path/to/binary
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("Create socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("Socket path: %s", socketPath)

		logger.Info("Listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		if listenErr != nil {
			logger.Fatal(listenErr)
		}
	} else {
		logger.Info("Listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("server in listening on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	logger.Fatal(server.Serve(listener))
}
