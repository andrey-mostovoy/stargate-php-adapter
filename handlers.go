package main

import (
	"encoding/json"
	"net/http"
)

type QueryRequest struct {
	CQL    string        `json:"cql"`
	Params []interface{} `json:"params"`
}

type BatchRequest struct {
	Queries []QueryRequest `json:"queries"`
}

func QueryHandler(s *Scylla) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req QueryRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		rows, err := s.Query(req.CQL, req.Params)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rows)
	}
}

func ExecHandler(s *Scylla) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req QueryRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err := s.Exec(req.CQL, req.Params)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(204)
	}
}

func BatchHandler(s *Scylla) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req BatchRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		err := s.Batch(req.Queries)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(204)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("ok"))
}