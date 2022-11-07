package auth

import (
	"context"
	"fmt"
	"github.com/Nerzal/gocloak/v12"
	"net/url"
	"strings"
	"sync"
	"time"
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

	_mu        sync.Locker  `json:"-"`
	_expiresAt time.Time    `json:"-"`
	_token     *gocloak.JWT `json:"-"`
}

func (cred *Credential) Token(client *gocloak.GoCloak, ctx context.Context) (*gocloak.JWT, error) {
	if time.Now().After(cred._expiresAt) {
		cred._mu.Lock()
		defer cred._mu.Unlock()

		var err error
		if cred._token != nil {
			cred._token, err = client.RefreshToken(ctx, cred._token.RefreshToken, cred.ClientId, cred.ClientSecret, cred.Realm)
			if err != nil {
				cred._token = nil
			}
		}

		if cred._token == nil {
			cred._token, err = client.LoginAdmin(ctx, cred.Login, cred.Password, cred.Realm)
			if err != nil {
				cred._token = nil
			}
		}

		if cred._token != nil {
			cred._expiresAt = time.Now().
				Add(time.Duration(cred._token.ExpiresIn) * time.Second).
				Add(-time.Second * 5)
		} else {
			cred._expiresAt = time.Time{}
			return nil, fmt.Errorf("could obtain access token: %v", err)
		}
	}

	return cred._token, nil
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
