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

	log.Println("adapter running on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}