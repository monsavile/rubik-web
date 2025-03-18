package config

import (
	"errors"
	"os"
	"strconv"
)

const portEnvName = "SERVER_PORT"

type httpServerConfig struct {
	port int
}

func NewHTTPServerConfig() (*httpServerConfig, error) {
	portStr := os.Getenv(portEnvName)
	if len(portStr) == 0 {
		return nil, errors.New("http server port not found")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("http server port is not a valid number")
	}

	return &httpServerConfig{
		port: port,
	}, nil
}

func (c *httpServerConfig) Port() int {
	return c.port
}
