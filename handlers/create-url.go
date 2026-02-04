package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status Ok"))
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		url string `json:"url"`
	}

	type response struct {
		id uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		return
	}

}
