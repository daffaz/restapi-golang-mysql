package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	w.Header().Set("Content-type", "application/json")
	dataInByte, err := json.Marshal(p)

	if err != nil {
		http.Error(w, "Error occured...", http.StatusBadRequest)
	}

	w.WriteHeader(status)
	w.Write(dataInByte)
	// w.Write([]byte(dataInByte))
}
