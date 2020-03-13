package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func RespondOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func RespondError(w http.ResponseWriter, err error, status int) {
	RespondJSON(w, status, map[string]string{"error": err.Error()})
}

func RespondInternal(w http.ResponseWriter, err error, status int) {
	log.Printf("%+v\n", err)
	RespondJSON(w, status, map[string]string{"error": http.StatusText(status)})
}
