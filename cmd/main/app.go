package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api-template/internal/config"
	"go-rest-api-template/internal/user"
	"go-rest-api-template/pkg/logging"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info(">==============================================================>")
	logger.Info("create router...")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("register user handler...")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application...")

	var listener net.Listener
	var listenerErr error

	if cfg.Listen.Type == "socket" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("creat socket")
		socketPath := path.Join(appDir, "app.socket")
		logger.Debugf("socket path: %s", socketPath)

		logger.Info("listen unix socket")
		listener, listenerErr = net.Listen("unix", socketPath)

	} else {
		logger.Info("listen tcp")
		listener, listenerErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenerErr != nil {
		logger.Fatal(listenerErr)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
