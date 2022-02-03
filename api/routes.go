package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/vaccine/:id", app.getOneVaccine)
	router.HandlerFunc(http.MethodGet, "/v1/vaccines", app.getAllVaccines)

	return app.enableCORS(router)
}
