package router

import (
	"github.com/javiertlopez/numeral/controller"

	"github.com/gorilla/mux"
)

// setupAppController setup the router with the log controller
func setupAppController(router *mux.Router, cont controller.Controller) {
	router.HandleFunc("/healthz", cont.Healthz).Methods("GET")
	router.HandleFunc("/statusz", cont.Statusz).Methods("GET")
}
