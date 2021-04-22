package router

import (
	"github.com/javiertlopez/numeral/controller"

	"github.com/gorilla/mux"
)

// setupLogController setup the router with the log controller
func setupLogController(router *mux.Router, cont controller.Controller) {
	router.HandleFunc("/logs", cont.CreateLog).Methods("POST")
}
