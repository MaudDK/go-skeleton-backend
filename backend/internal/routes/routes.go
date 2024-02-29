package api

import (
	"app/internal/bootstrap"
	"app/internal/controllers/accounts"
	"app/internal/controllers/health"
	"app/pkg/errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes(router *httprouter.Router, app *bootstrap.Application){
	//Get Logger
	logger := app.Logger
	//TODO WRITE A WRAPPER FOR THIS
	//Not Found Responses
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		errs.NotFoundResponse(logger, w, r)
	})

	//Method Not Allowed Responses
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		errs.MethodNotAllowedResponse(logger, w, r)
	})

	//Health
	router.Handler(http.MethodGet, "/api/v1/health", health.GetHealth(app))
	
	//Users CRUD
	router.Handler(http.MethodPost, "/api/v1/accounts", accounts.CreateAccount(app))
	router.Handler(http.MethodGet, "/api/v1/accounts/:id", accounts.GetAccount(app))
	
}