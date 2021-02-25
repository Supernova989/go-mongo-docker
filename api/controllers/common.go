package controllers

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(w http.ResponseWriter, payload interface{}, status int, errorCode *int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := make(map[string]interface{})
	if errorCode != nil {
		resp["errorCode"] = &errorCode
	}
	if errorMessage != "" {
		resp["errorMessage"] = &errorMessage
	}
	resp["payload"] = payload
	json.NewEncoder(w).Encode(resp)
}
