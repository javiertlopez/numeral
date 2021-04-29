package controller

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/javiertlopez/numeral"
	"github.com/javiertlopez/numeral/errorcodes"

	"github.com/gorilla/mux"
)

// Create controller
func (c *controller) CreateLog(w http.ResponseWriter, r *http.Request) {
	device := r.Header.Get("x-device-id")
	if len(device) == 0 {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: "x-device-id missing",
				Status:  http.StatusBadRequest,
			},
		)

		return
	}

	file, handler, err := r.FormFile("filename")
	if err != nil {
		JSONResponse(
			w, http.StatusUnprocessableEntity,
			Response{
				Message: err.Error(),
				Status:  http.StatusUnprocessableEntity,
			},
		)

		return
	}
	defer file.Close()

	response, err := c.repository.CreateLog(
		r.Context(),
		numeral.Log{
			DeviceID: device,
			ImageKey: handler.Filename,
		},
	)
	if err != nil {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		)

		return
	}

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
		KeyID:  response.ID,
		Binary: buf.Bytes(),
	}

	_, err = c.storage.PutImage(r.Context(), image)
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
		response,
	)
}

// UpdateLog controller
func (c *controller) UpdateLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var log numeral.Log
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&log); err != nil {
		JSONResponse(
			w, http.StatusBadRequest,
			Response{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		)
		return
	}
	defer r.Body.Close()

	response, err := c.repository.UpdateLog(r.Context(), id, log)
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
		response,
	)
}
