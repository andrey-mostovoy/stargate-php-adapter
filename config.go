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
}

func LoadConfig() Config {

	numConns, _ := strconv.Atoi(getEnv("SCYLLA_NUM_CONNS", "4"))
	timeout, _ := strconv.Atoi(getEnv("SCYLLA_TIMEOUT_MS", "100"))

	return Config{
		Hosts: strings.Split(getEnv("SCYLLA_HOSTS", "127.0.0.1"), ","),
		Keyspace: getEnv("SCYLLA_KEYSPACE", "app"),
		Port: getEnv("PORT", "8080"),
		NumConns: numConns,
		TimeoutMs: timeout,
	}
}

func getEnv(k, d string) string {
	v := os.Getenv(k)
	if v == "" {
		return d
	}
	return v
}