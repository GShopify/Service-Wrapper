package db

import (
	"fmt"
	"strings"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Params   params `json:"params"`
}

func New() *Config {
	return &Config{}
}

type params struct {
	Timeout      uint `json:"timeout"`
	ReadTimeout  uint `json:"read_timeout"`
	WriteTimeout uint `json:"write_timeout"`
	Debug        bool `json:"debug"`
}

func (c *Config) Key() string {
	return "database"
}

func (c *Config) Validate() error {
	if strings.TrimSpace(c.Username) == "" {
		return fmt.Errorf("`username` must not be an empty string")
	}

	if strings.TrimSpace(c.Password) == "" {
		return fmt.Errorf("`password` must not be an empty string")
	}

	if strings.TrimSpace(c.Host) == "" {
		c.Host = "localhost"
	}

	if c.Port < 1 {
		c.Port = 8123
	}

	if strings.TrimSpace(c.Database) == "" {
		c.Database = "default"
	}

	if c.Params.Timeout < 1 {
		return fmt.Errorf("`params.timeout` must be longer than 1s")
	}

	if c.Params.ReadTimeout < 1 {
		return fmt.Errorf("`params.read_timeout` must be longer than 1s")
	}

	if c.Params.WriteTimeout < 1 {
		return fmt.Errorf("`params.write_timeout` must be longer than 1s")
	}

	return nil
}
