package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lokeshkarthik5/url-shortner/internal/database"
	"github.com/lokeshkarthik5/url-shortner/utils"
)

type Controllers struct {
	DB *database.Queries
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status Ok"))
}

func (c *Controllers) CreateUrl(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		url string `json:"url"`
	}

	type response struct {
		id        uuid.UUID `json:"id"`
		longUrl   string    `json:"long_url"`
		short     string    `json:"short_url"`
		createdAt time.Time `json:"created_at"`
		updatedAt time.Time `json:"updated_at"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		return
	}

	short := utils.GenerateShortUrls()

	link, err := c.DB.CreateUrl(r.Context(), database.CreateUrlParams{
		Longurl: params.url,
		Short:   short,
	})
	if err != nil {
		log.Println(err)
		return
	}

	res := response{
		id:        link.ID,
		longUrl:   link.Longurl,
		short:     link.Short,
		createdAt: link.Createdat,
		updatedAt: link.Updatedat,
	}

	respondWithJSON(w, http.StatusCreated, res)

}
