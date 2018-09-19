package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const baseUrl = "/"
const defaultPort = "8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	http.Handle(baseUrl, r)

	fmt.Println("Starting API server on port", defaultPort)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + defaultPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
