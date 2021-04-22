package controller

import (
	"net/http"

	"github.com/javiertlopez/numeral"
)

// Controller handles the HTTP requests
type Controller interface {
	CreateLog(w http.ResponseWriter, r *http.Request)
}

// controller struct holds the usecase
type controller struct {
	repository numeral.Repository
	storage    numeral.Storage
}

// New returns a controller
func New(
	r numeral.Repository,
	s numeral.Storage,
) Controller {
	return &controller{
		repository: r,
		storage:    s,
	}
}
