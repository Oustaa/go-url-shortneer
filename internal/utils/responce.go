package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponce(w http.ResponseWriter, body interface{}) {
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, "Enable to encode the urls.", http.StatusInternalServerError)
	}
}
