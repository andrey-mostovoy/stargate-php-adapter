package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocql/gocql"
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

		ctx, cancel := context.WithTimeout(
			r.Context(),
			time.Duration(s.cfg.TimeoutMs)*time.Millisecond,
		)

		defer cancel()

		var req QueryRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		iter, err := s.Query(ctx, req.CQL, req.Params)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		streamRows(w, iter)
	}
}

func ExecHandler(s *Scylla) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			time.Duration(s.cfg.TimeoutMs)*time.Millisecond,
		)

		defer cancel()

		var req QueryRequest

		json.NewDecoder(r.Body).Decode(&req)

		err := s.Exec(ctx, req.CQL, req.Params)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(204)
	}
}

func BatchHandler(s *Scylla) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(
			r.Context(),
			time.Duration(s.cfg.TimeoutMs)*time.Millisecond,
		)

		defer cancel()

		var req BatchRequest

		json.NewDecoder(r.Body).Decode(&req)

		err := s.Batch(ctx, req.Queries)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(204)
	}
}

func streamRows(w http.ResponseWriter, iter *gocql.Iter) {

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)

	w.Write([]byte("["))

	first := true

	m := map[string]interface{}{}

	for iter.MapScan(m) {

		if !first {
			w.Write([]byte(","))
		}

		first = false

		enc.Encode(m)

		m = map[string]interface{}{}
	}

	w.Write([]byte("]"))

	iter.Close()
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("ok"))
}