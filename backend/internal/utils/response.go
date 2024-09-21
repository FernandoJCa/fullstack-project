package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	log.Printf("Sent response with status code %d: %s", code, string(response))
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, JsonResponse{
		Status:  "error",
		Message: message,
	})
	log.Printf("Sent error response with status code %d: %s", code, message)
}

func RespondNoContent(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNoContent)
	log.Printf("Sent no content response: %s", message)
}
