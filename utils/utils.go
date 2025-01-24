package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParsearJSON(r *http.Request, payload any) error {
	//parsear el json
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(&payload)
}

func CrearJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func CrearError(w http.ResponseWriter, status int, err error) {
	CrearJSON(w, status, map[string]string{"error": err.Error()})
}
