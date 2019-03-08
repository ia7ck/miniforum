package main

import (
	"miniforum/controller"
	"net/http"
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

}
