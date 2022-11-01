package config

import (
	"bytes"
	"context"
	"encoding/json"
	"filippo.io/age"
	"filippo.io/age/agessh"
	"fmt"
	"github.com/gshopify/service-wrapper/auth"
	"github.com/spf13/pflag"
	v3 "go.etcd.io/etcd/client/v3"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var (
	fUsername  = pflag.StringP("username", "u", "", "configurator username")
	fPassword  = pflag.String("password", "", "configurator password")
	fEndpoints = pflag.StringArray("endpoints", []string{}, "configurator endpoint list")
	fIdentity  = pflag.StringP("identity", "i", "", "configurator private key")
	fRecipient = pflag.StringP("recipient", "r", "", "configurator public key")

	once     sync.Once
	instance *loader
)

type loader struct {
	cli        *v3.Client
	identities []age.Identity
}

//goland:noinspection ALL
func Instance() *loader {
	once.Do(func() {
		if *fUsername == "" {
			panic("configurator username must be provided")
		}

		if *fPassword == "" {
			panic("configurator password must be provided")
		}

		if len(*fEndpoints) == 0 {
			panic("configurator endpoints must be provided")
		}

		if *fIdentity == "" {
			panic("configurator private key must be provided")
		}

		var err error
		instance = new(loader)

		if instance.cli, err = v3.New(v3.Config{
			Endpoints:   *fEndpoints,
			DialTimeout: etcdTimeout,
			Username:    *fUsername,
			Password:    *fPassword,
		}); err != nil {
			panic(err)
		}

		if err = auth.Init(instance.cli, etcdTimeout); err != nil {
			panic(err)
		}

		id, err := instance.parseIdentity(*fIdentity)
		if err != nil {
			panic(err)
		}

		instance.identities = []age.Identity{id}
	})

	return instance
}

func (l *loader) parseIdentity(path string) (age.Identity, error) {
	crt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return agessh.ParseIdentity(crt)
}

func (l *loader) Load(parent context.Context, c IConfig) error {
	if c == nil {
		return fmt.Errorf("`IConfig` must be provided")
	}

	if parent == nil {
		return fmt.Errorf("parent `Context` must be provided")
	}

	ctx, cancel := context.WithTimeout(parent, etcdTimeout)
	defer cancel()

	r, err := l.cli.Get(ctx, l.component(c), v3.WithLastRev()...)
	if err != nil {
		return fmt.Errorf("could not retrieve config data: %v", err)
	}

	if r.Count < 1 {
		return fmt.Errorf("could not retrieve config data: no data")
	}

	rdata, err := age.Decrypt(bytes.NewBuffer(r.Kvs[0].Value), l.identities...)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(rdata)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return c.Validate()
}

func (l *loader) Set(parent context.Context, c IConfig, data io.Reader) error {
	if *fRecipient == "" {
		return fmt.Errorf("configurator public key must be provided")
	}

	crtData, err := ioutil.ReadFile(*fRecipient)
	if err != nil {
		return err
	}

	crt, err := agessh.ParseRecipient(string(crtData))
	if err != nil {
		return err
	}

	out := &bytes.Buffer{}
	w, err := age.Encrypt(out, crt)
	if err != nil {
		return err
	}

	if _, err = io.Copy(w, data); err != nil {
		return err
	}

	if err = w.Close(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(parent, etcdTimeout)
	defer cancel()

	if _, err = l.cli.Put(ctx, l.component(c), out.String()); err != nil {
		return err
	}

	return nil
}

func (l *loader) component(c IConfig) string {
	return fmt.Sprintf("component/%s", c.Key())
}

// PrintBanner
// @ref http://patorjk.com/software/taag/#p=display&f=Speed&t=gshopper
func PrintBanner() {
	log.Println("               ______                                    ")
	log.Println("_______ __________  /____________________________________")
	log.Println("__  __ `/_  ___/_  __ \\  __ \\__  __ \\__  __ \\  _ \\_  ___/")
	log.Println("_  /_/ /_(__  )_  / / / /_/ /_  /_/ /_  /_/ /  __/  /    ")
	log.Println("_\\__, / /____/ /_/ /_/\\____/_  .___/_  .___/\\___//_/     ")
	log.Println("/____/                      /_/     /_/                  ")
	log.Println()
}
