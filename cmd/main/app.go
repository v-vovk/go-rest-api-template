package main

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api-template/internal/user"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("create router...")
	router := httprouter.New()

	log.Println("register user handler...")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application...")
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
