package common

import "time"

type HealthCheck struct {
	Version    string    `json:"version"`
	ServerTime time.Time `json:"server_time"`
}

type DummyLog struct {
	Action     string    `json:"action"`
	ServerTime time.Time `json:"server_time"`
}
