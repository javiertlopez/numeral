package controller

import (
	"net/http"

	"github.com/javiertlopez/numeral"
)

// Controller handles the HTTP requests
type Controller interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	Statusz(w http.ResponseWriter, r *http.Request)

	CreateLog(w http.ResponseWriter, r *http.Request)
	UpdateLog(w http.ResponseWriter, r *http.Request)
}

// controller struct holds the usecase
type controller struct {
	commit  string
	version string

	repository numeral.Repository
	storage    numeral.Storage
}

// New returns a controller
func New(
	c string,
	v string,
	r numeral.Repository,
	s numeral.Storage,
) Controller {
	return &controller{
		commit:     c,
		version:    v,
		repository: r,
		storage:    s,
	}
}
