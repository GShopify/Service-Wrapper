package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/gshopify/service-wrapper/server/middleware"
	"net/http"
)

func newRouter(service *handler.Server, path string, middlewares ...mux.MiddlewareFunc) http.Handler {
	router := mux.NewRouter()
	router.Handle(path, service)

	router.Use(middleware.GShopify())
	router.Use(middlewares...)

	return router
}
