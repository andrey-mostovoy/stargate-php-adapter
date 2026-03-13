package main

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

type Scylla struct {
	session *gocql.Session
	cache   *PreparedCache
}

func NewScylla(cfg Config) *Scylla {

	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Timeout = time.Duration(cfg.TimeoutMs) * time.Millisecond
	cluster.NumConns = cfg.NumConns
	cluster.Consistency = gocql.Quorum
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	return &Scylla{
		session: session,
		cache:   NewCache(),
	}
}

func (s *Scylla) Query(cql string, params []interface{}) ([]map[string]interface{}, error) {

	iter := s.session.Query(cql, params...).Iter()

	rows := []map[string]interface{}{}

	m := map[string]interface{}{}

	for iter.MapScan(m) {

		row := map[string]interface{}{}

		for k, v := range m {
			row[k] = v
		}

		rows = append(rows, row)

		m = map[string]interface{}{}
	}

	return rows, iter.Close()
}

func (s *Scylla) Exec(cql string, params []interface{}) error {

	return s.session.Query(cql, params...).Exec()
}

func (s *Scylla) Batch(queries []QueryRequest) error {

	b := s.session.NewBatch(gocql.LoggedBatch)

	for _, q := range queries {
		b.Query(q.CQL, q.Params...)
	}

	return s.session.ExecuteBatch(b)
}