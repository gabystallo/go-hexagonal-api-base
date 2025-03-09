package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		panic("Error writing JSON to client")
	}
}
