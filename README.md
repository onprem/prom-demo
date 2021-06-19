# `prom-demo`

A small demo for monitoring an application with Prometheus.

## Instructions

- Run the application using `go run main.go`

- To run the Prometheus stack:
  - Go to test directory `cd test`
  - Run the docker-compose setup `docker-compose up`

- Prometheus UI will be now available at `http://localhost:9090`

- Send requests to `http://localhost:8080/ping` and `http://localhost:8080/roulette` to generate metrics from our demo application.
