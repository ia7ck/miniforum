package main

import (
	"net/http"
	"time"

	"github.com/ia7ck/miniforum/controller"
)

func main() {
	mux := http.NewServeMux()
	controller.Routes(mux)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func init() {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = location
}
