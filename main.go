package main

import (
	"log"
	"net/http"
)

func main() {

	cfg := LoadConfig()

	scylla := NewScylla(cfg)

	mux := http.NewServeMux()

	mux.Handle("/query", QueryHandler(scylla))
	mux.Handle("/exec", ExecHandler(scylla))
	mux.Handle("/batch", BatchHandler(scylla))

	mux.HandleFunc("/health", HealthHandler)
	mux.Handle("/metrics", MetricsHandler())

	log.Println("scylla adapter listening :" + cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, mux)

	if err != nil {
		panic(err)
	}
}