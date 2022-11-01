package auth

import (
	"fmt"
	"net/url"
	"strings"
)

type Config struct {
	Endpoint     string `json:"endpoint"`
	Realm        string `json:"realm"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Admin        string `json:"admin"`
	Password     string `json:"password"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) Key() string {
	return "auth"
}

func (c *Config) Validate() error {
	if strings.TrimSpace(c.Endpoint) == "" {
		return fmt.Errorf("endpoint must be provided")
	}

	endpoint, err := url.Parse(c.Endpoint)
	if err != nil {
		return fmt.Errorf("could not parse endpoint: %v", err)
	}

	if !strings.HasPrefix(endpoint.Scheme, "http") {
		return fmt.Errorf("illegal endpoint scheme: %s", endpoint.Scheme)
	}

	if strings.TrimSpace(endpoint.Host) == "" {
		return fmt.Errorf("endpoint host must be provided")
	}

	if strings.TrimSpace(c.Realm) == "" {
		return fmt.Errorf("keycloak realm must be provided")
	}

	if strings.TrimSpace(c.ClientId) == "" {
		return fmt.Errorf("keycloak ClientId must be provided")
	}

	if strings.TrimSpace(c.ClientSecret) == "" {
		return fmt.Errorf("keycloak ClientSecret must be provided")
	}

	if strings.TrimSpace(c.Admin) == "" {
		return fmt.Errorf("keycloak admin must be provided")
	}

	if strings.TrimSpace(c.Password) == "" {
		return fmt.Errorf("keycloak admin password must be provided")
	}

	return nil
}
