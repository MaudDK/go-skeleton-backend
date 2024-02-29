package server

import (
	"net/http"
	"app/internal/routes"
	"github.com/julienschmidt/httprouter"
	"app/internal/bootstrap"
)

func NewServer(app *bootstrap.Application) http.Handler{
	router := httprouter.New()
	api.Routes(router, app)
	var handler http.Handler = router
	// handler = someMiddleware(handler)
	// handler = someMiddleware2(handler)
	// handler = someMiddleware3(handler)
	return handler
}