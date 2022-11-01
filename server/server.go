package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"net/http"
)

func NewServer(service *handler.Server, opts *Opts) *http.Server {
	return &http.Server{
		Addr:         opts.addr(),
		WriteTimeout: opts.WriteTimeout,
		ReadTimeout:  opts.ReadTimeout,
		IdleTimeout:  opts.IdleTimeout,
		Handler:      newRouter(service, opts.HandlerPath, opts.middlewares...),
	}
}
