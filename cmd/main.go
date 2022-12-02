package main

import (
	"context"
	"fmt"
	"github.com/gshopify/service-wrapper/auth"
	"github.com/gshopify/service-wrapper/config"
	"github.com/gshopify/service-wrapper/db"
	"github.com/spf13/pflag"
	"os"
)

var (
	fPreset = pflag.String("preset", "", "available modes are: auth; clickhouse")
	fConfig = pflag.String("config", "", "config file path")
)

func init() {
	pflag.Parse()
	config.Instance()
}

func main() {
	if *fConfig == "" {
		panic(fmt.Errorf("config file path must be provided"))
	}

	set := preset(*fPreset)
	if !set.Valid() {
		panic(fmt.Errorf("unknown preset mode `%s`", *fPreset))
	}

	var cfg config.IConfig
	switch set {
	case modeAuth:
		cfg = auth.New()
	case modeDatabase:
		cfg = db.New()
	}

	f, err := os.Open(*fConfig)
	if err != nil {
		panic(err)
	}

	err = config.Instance().Set(context.Background(), cfg, f)
	if err != nil {
		panic(err)
	}
}
