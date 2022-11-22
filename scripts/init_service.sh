#!/bin/bash

while [ $# -gt 0 ]; do
  case "$1" in
    --path=*)
      path="${1#*=}"
      ;;
    --name=*)
      name="${1#*=}"
      ;;
    *)
      printf "***************************\n"
      printf "* Error: Invalid argument.*\n"
      printf "***************************\n"
      exit 1
  esac
  shift
done

if [[ -z "$name" ]]; then
  name="service"
fi

if [[ -z "$path" ]]; then
  path="$(pwd)"
fi

[ -d "$path" ] && echo "Directory $path ALREADY exists." && exit 1

cd "$(dirname "$path")" || exit 1
path="$(cd "$(dirname "$path")" || exit 1; pwd)/$(basename "$path")"

printf "Creating root dir: %s\n" "$path"

mkdir -p "$path"
mkdir -p "$path/cmd"
mkdir -p "$path/graphql"

cd "$path" || exit 1

go mod init "gshopper.com/gshopify/$name"
go get github.com/gshopify/service-wrapper
go get github.com/spf13/pflag

printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > "$path/tools.go"
echo 'package main

import(
  "context"
  "github.com/gshopify/service-wrapper/config"
  "github.com/spf13/pflag"
  "log"
  "os"
  "os/signal"
  "syscall"
  "time"
)

var (
  signals  = make(chan os.Signal, 1)
)

func init() {
	pflag.Parse()
	config.Instance()
}

func main() {
  config.PrintBanner()

  go func() {
    // server must be run here
  }()

  signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  <-signals

  ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
  defer cancel()

  _ = srv.Shutdown(ctx)

  log.Println("shutting down")
  os.Exit(0)
}' | gofmt > "$path/cmd/main.go"

go mod tidy

echo "" > "$path/graphql/schema.graphql"
echo "federation:
  filename: generated/federation.go
  package: generated
  version: 2
schema:
  - schema.graphql
exec:
  filename: generated/generated.go
model:
  filename: generated/models_gen.go
resolver:
  filename: generated/resolver.go
  type: Resolver
autobind: []
omit_getters: true
#omit_getters: true
#omit_slice_element_pointers: true
models:
  DateTime:
    model: github.com/gshopify/service-wrapper/scalar.DateTime
  LanguageCode:
    model: github.com/gshopify/service-wrapper/model.LanguageCode
  CountryCode:
    model: github.com/gshopify/service-wrapper/model.CountryCode
  UnsignedInt64:
    model: github.com/gshopify/service-wrapper/model.UInt
  PageInfo:
    model: github.com/gshopify/service-wrapper/model.PageInfo
  HTML:
    model: github.com/gshopify/service-wrapper/scalar.Html
  JSON:
    model: github.com/gshopify/service-wrapper/scalar.Json" > "$path/graphql/gqlgen.yml"

echo "== DONE"