package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"time"
)

type Opts struct {
	Port         int
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
	HandlerPath  string

	middlewares []mux.MiddlewareFunc
}

func NewDefaultOpts(port int, timeout time.Duration) (*Opts, error) {
	if port < 1 {
		return nil, fmt.Errorf("illegal port to expose: [%d]", port)
	}

	if timeout.Seconds() < 1 {
		return nil, fmt.Errorf("illegal timeout settings: [%fs]", timeout.Seconds())
	}

	return &Opts{
		Port:         port,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
		IdleTimeout:  timeout * 2,
		HandlerPath:  "/query",

		middlewares: make([]mux.MiddlewareFunc, 0),
	}, nil
}

func (o *Opts) addr() string {
	return fmt.Sprintf(":%d", o.Port)
}
