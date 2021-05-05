package web

import (
	"encoding/json"
	"errors"
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

func Ok(w http.ResponseWriter, body interface{}) {
	writeResponseWithStatusCode(w, body, 200)
}

func WriteSecurityError(w http.ResponseWriter, errorBody interface{}) {
	writeResponseWithStatusCode(w, errorBody, 401)
}

func Confict(w http.ResponseWriter, errorBody interface{}) {
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
