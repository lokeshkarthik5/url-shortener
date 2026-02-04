package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func createUrl(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		url string `json:"url"`
	}

	type response struct{
		
	}

	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		return
	}

	

}
