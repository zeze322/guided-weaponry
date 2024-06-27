package api

import (
	"encoding/json"
	"net/http"
)

type ApiErr struct {
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTP(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiErr{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
