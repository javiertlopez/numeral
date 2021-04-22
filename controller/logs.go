package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/javiertlopez/numeral"
	"github.com/javiertlopez/numeral/errorcodes"
)

// Create controller
func (c *controller) CreateLog(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("name")
	if err != nil {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		)
	}

	image := numeral.Image{
		KeyID:  handler.Filename,
		Binary: buf.Bytes(),
	}

	response, err := c.storage.PutImage(r.Context(), image)
	if err != nil {
		// Look for Custom Error
		if err == errorcodes.ErrUnprocessable {
			JSONResponse(
				w, http.StatusUnprocessableEntity,
				Response{
					Message: "Unprocessable entity",
					Status:  http.StatusUnprocessableEntity,
				},
			)
			return
		}

		JSONResponse(
			w, http.StatusInternalServerError,
			Response{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		)

		return
	}

	JSONResponse(
		w,
		http.StatusCreated,
		fmt.Sprintf("key: %s, file: %s", response.KeyID, handler.Filename),
	)
}
