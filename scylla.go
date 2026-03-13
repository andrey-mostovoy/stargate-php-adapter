package main

import (
	"context"
	"time"

	"github.com/gocql/gocql"
)

type Scylla struct {
	session *gocql.Session
	cache   *PreparedCache
	cfg     Config
}

func NewScylla(cfg Config) *Scylla {

	cluster := gocql.NewCluster(cfg.Hosts...)

	cluster.Keyspace = cfg.Keyspace
	cluster.NumConns = cfg.NumConns
	cluster.Timeout = time.Duration(cfg.TimeoutMs) * time.Millisecond
	cluster.PageSize = cfg.PageSize
	cluster.Consistency = gocql.Quorum
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	return &Scylla{
		session: session,
		cache:   NewPreparedCache(cfg.PreparedSize),
		cfg:     cfg,
	}
}

func (s *Scylla) Query(ctx context.Context, cql string, params []interface{}) (*gocql.Iter, error) {

	q := s.session.Query(cql, params...)
	q = q.WithContext(ctx)

	return q.Iter(), nil
}

func (s *Scylla) Exec(ctx context.Context, cql string, params []interface{}) error {

	q := s.session.Query(cql, params...)
	q = q.WithContext(ctx)

	return q.Exec()
}

func (s *Scylla) Batch(ctx context.Context, queries []QueryRequest) error {

	b := s.session.NewBatch(gocql.LoggedBatch)

	for _, q := range queries {
		b.Query(q.CQL, q.Params...)
	}

	return s.session.ExecuteBatch(b.WithContext(ctx))
}