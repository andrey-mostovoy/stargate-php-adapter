package main

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Hosts        []string
	Keyspace     string
	Port         string
	NumConns     int
	TimeoutMs    int
	PageSize     int
	PreparedSize int
}

func LoadConfig() Config {

	numConns, _ := strconv.Atoi(getEnv("SCYLLA_NUM_CONNS", "4"))
	timeout, _ := strconv.Atoi(getEnv("SCYLLA_TIMEOUT_MS", "200"))
	pageSize, _ := strconv.Atoi(getEnv("SCYLLA_PAGE_SIZE", "500"))
	cacheSize, _ := strconv.Atoi(getEnv("SCYLLA_PREPARED_CACHE", "500"))

	return Config{
		Hosts:        strings.Split(getEnv("SCYLLA_HOSTS", "127.0.0.1"), ","),
		Keyspace:     getEnv("SCYLLA_KEYSPACE", "app"),
		Port:         getEnv("PORT", "8080"),
		NumConns:     numConns,
		TimeoutMs:    timeout,
		PageSize:     pageSize,
		PreparedSize: cacheSize,
	}
}

func getEnv(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		return d
	}
	return v
}