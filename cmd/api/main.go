package main

import (
	"github.com/caleflat/boilerplate/cmd/api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	router := mux.NewRouter()
	router.Use(middleware.Auth)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatalf("Server failed to start: %v", http.ListenAndServe(":8080", router))
}
