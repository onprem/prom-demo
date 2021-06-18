package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	addr = "0.0.0.0:8080"
	name = "demo"
)

var (
	requestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name:      "http_requests_total",
		Namespace: name,
		Help:      "Total number of incoming HTTP requests.",
	}, []string{"handler", "code"})
)

func sayPong(w http.ResponseWriter, r *http.Request) {
	requestsTotal.WithLabelValues("ping", "200").Inc()

	delay := rand.Float64() * 5 * 1000
	time.Sleep(time.Millisecond * time.Duration(delay))
	fmt.Fprintf(w, "pong after %ds!", int(delay/1000))
}

func russianRoulette(w http.ResponseWriter, r *http.Request) {
	chamber := rand.Intn(6)
	if chamber == 0 {
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("you die!"))

		requestsTotal.WithLabelValues("roulette", fmt.Sprint(http.StatusTeapot)).Inc()

		return
	}

	requestsTotal.WithLabelValues("roulette", fmt.Sprint(http.StatusOK)).Inc()
	w.Write([]byte("you live!"))
}

func main() {
	// Seeding to get different pseudo random results everytime.
	rand.Seed(time.Now().Unix())

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", sayPong)
	mux.HandleFunc("/roulette", russianRoulette)

	// Expose Prometheus metrics.
	mux.Handle("/metrics", promhttp.Handler())

	log.Printf("starting web server on: %s\n", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Printf("error starting server: %v\n", err)
	}
}
