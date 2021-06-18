package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	addr = "0.0.0.0:8080"
)

func sayPong(w http.ResponseWriter, r *http.Request) {
	delay := rand.Float64() * 5 * 1000
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Fprintf(w, "pong after %ds!", int(delay/1000))
}

func main() {
	// Seeding to get different pseudo random results everytime.
	rand.Seed(time.Now().Unix())

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", sayPong)

	log.Printf("starting web server on: %s\n", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Printf("error starting server: %v\n", err)
	}
}
