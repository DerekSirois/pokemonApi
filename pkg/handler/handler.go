package handler

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	sendJson(w, http.StatusOK, "Hello World")
}

func sendJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = w.Write(payload)
	if err != nil {
		return err
	}
	return nil
}
