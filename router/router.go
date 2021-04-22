package router

import (
	"github.com/javiertlopez/numeral/controller"

	"github.com/gorilla/mux"
)

// New returns a *mux.Router
func New(
	app controller.Controller,
) *mux.Router {
	router := mux.NewRouter()

	setupLogController(router, app)

	return router
}
