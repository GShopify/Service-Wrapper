package auth

import (
	"fmt"
	"net/url"
	"strings"
)

type Config struct {
	Endpoint string     `json:"endpoint"`
	Admin    Credential `json:"admin"`
	Cli      Credential `json:"cli"`
}

type Credential struct {
	Realm        string `json:"realm"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Login        string `json:"login"`
	Password     string `json:"password"`
}

func (cred *Credential) Validate() error {
	if strings.TrimSpace(cred.Realm) == "" {
		return fmt.Errorf("credential realm must be provided")
	}

	if strings.TrimSpace(cred.ClientId) == "" {
		return fmt.Errorf("credential ClientId must be provided")
	}

	if strings.TrimSpace(cred.ClientSecret) == "" && strings.TrimSpace(cred.Login) == "" {
		return fmt.Errorf("credential secrets must be provided")
	}

	if strings.TrimSpace(cred.ClientSecret) == "" && strings.TrimSpace(cred.Password) == "" {
		return fmt.Errorf("credential secrets must be provided")
	}

	return nil
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

	if err = c.Admin.Validate(); err != nil {
		return err
	}

	if err = c.Cli.Validate(); err != nil {
		return err
	}

	return nil
}
