package web

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

func ReadBody(body io.Reader, object interface{}) error {
	err := json.NewDecoder(body).Decode(object)
	if err != nil {
		return errors.New("json.bad.format")
	}
	return nil
}

func GetPathVariable(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func Ok(w http.ResponseWriter, body interface{}) {
	writeResponseWithStatusCode(w, body, 200)
}

func WriteSecurityError(w http.ResponseWriter, errorBody interface{}) {
	writeResponseWithStatusCode(w, errorBody, 401)
}

func Conflict(w http.ResponseWriter, errorBody interface{}) {
	writeResponseWithStatusCode(w, errorBody, 409)
}

func InternalServerError(w http.ResponseWriter, errorBody interface{}) {
	writeResponseWithStatusCode(w, errorBody, 500)
}

func writeResponseWithStatusCode(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)

}
