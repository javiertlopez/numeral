package controller

import (
	"net/http"
)

// Healthz controller
func (c *controller) Healthz(w http.ResponseWriter, r *http.Request) {
	JSONResponse(
		w,
		http.StatusOK,
		Response{
			Message: "Hello World!",
			Status:  http.StatusOK,
		},
	)
}

// Statusz controller
func (c *controller) Statusz(w http.ResponseWriter, r *http.Request) {
	JSONResponse(
		w,
		http.StatusOK,
		map[string]interface{}{
			"commit":  c.commit,
			"version": c.version,
		},
	)
}
