package main

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api-template/internal/user"
	"go-rest-api-template/pkg/logging"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router...")
	router := httprouter.New()

	logger.Info("register user handler...")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application...")
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server listening port :8888")
	logger.Fatal(server.Serve(listener))
}
