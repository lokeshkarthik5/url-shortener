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

func (c *Controllers) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("urlId")
	if urlId == "" {
		return
	}
	err := c.DB.DeleteUrl(r.Context(), urlId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot delete", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Url Deleted"))

}

func (c *Controllers) GetStats(w http.ResponseWriter, r *http.Request) {
	type response struct {
		id          uuid.UUID `json:"id"`
		longUrl     string    `json:"long_url"`
		short       string    `json:"short_url"`
		createdAt   time.Time `json:"created_at"`
		updatedAt   time.Time `json:"updated_at"`
		accesscount int       `json:"access_count"`
	}

	urlId := r.PathValue("urlId")
	if urlId == "" {
		return
	}
	url, err := c.DB.GetCounts(r.Context(), urlId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot get db", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		id:          url.ID,
		longUrl:     url.Longurl,
		short:       url.Short,
		createdAt:   url.Createdat,
		updatedAt:   url.Updatedat,
		accesscount: url.Accesscount,
	})

}

func (c *Controllers) GetUrl(w http.ResponseWriter, r *http.Request) {

	type response struct {
		id        uuid.UUID `json:"id"`
		longUrl   string    `json:"long_url"`
		short     string    `json:"short_url"`
		createdAt time.Time `json:"created_at"`
		updatedAt time.Time `json:"updated_at"`
	}

	urlId := r.PathValue("urlId")
	if urlId == "" {
		return
	}

	url, err := c.DB.GetUrl(r.Context(), urlId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB error", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		id:        url.ID,
		longUrl:   url.Longurl,
		short:     url.Short,
		createdAt: url.Createdat,
		updatedAt: url.Updatedat,
	})
}

func (c *Controllers) UpdateUrl(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		longUrl string `json:"long_url"`
	}

	type response struct {
		id        uuid.UUID `json:"id"`
		longUrl   string    `json:"long_url"`
		short     string    `json:"short_url"`
		createdAt time.Time `json:"created_at"`
		updatedAt time.Time `json:"updated_at"`
	}

	urlId := r.PathValue("urlId")
	if urlId == "" {
		return
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding", err)
		return
	}

	url, err := c.DB.UpdateUrl(r.Context(), database.UpdateUrlParams{
		Longurl: params.longUrl,
		Short:   urlId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB Error", err)
		return
	}

	respondWithJSON(w, http.StatusAccepted, response{
		id:        url.ID,
		longUrl:   url.Longurl,
		short:     url.Short,
		createdAt: url.Createdat,
		updatedAt: url.Updatedat,
	})

}
