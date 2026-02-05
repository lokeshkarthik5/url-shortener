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
		Url string `json:"url"`
	}

	type response struct {
		Id        uuid.UUID `json:"id"`
		LongUrl   string    `json:"long_url"`
		Short     string    `json:"short_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err)
		return
	}

	short := utils.GenerateShortUrls()

	link, err := c.DB.CreateUrl(r.Context(), database.CreateUrlParams{
		Longurl: params.Url,
		Short:   short,
	})
	if err != nil {
		log.Println(err)
		return
	}

	res := response{
		Id:        link.ID,
		LongUrl:   link.Longurl,
		Short:     link.Short,
		CreatedAt: link.Createdat,
		UpdatedAt: link.Updatedat,
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
		Id          uuid.UUID `json:"id"`
		LongUrl     string    `json:"long_url"`
		Short       string    `json:"short_url"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Accesscount int       `json:"access_count"`
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
		Id:          url.ID,
		LongUrl:     url.Longurl,
		Short:       url.Short,
		CreatedAt:   url.Createdat,
		UpdatedAt:   url.Updatedat,
		Accesscount: url.Accesscount,
	})

}

func (c *Controllers) GetUrl(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Id        uuid.UUID `json:"id"`
		LongUrl   string    `json:"long_url"`
		Short     string    `json:"short_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
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
		Id:        url.ID,
		LongUrl:   url.Longurl,
		Short:     url.Short,
		CreatedAt: url.Createdat,
		UpdatedAt: url.Updatedat,
	})
}

func (c *Controllers) UpdateUrl(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LongUrl string `json:"long_url"`
	}

	type response struct {
		Id        uuid.UUID `json:"id"`
		LongUrl   string    `json:"long_url"`
		Short     string    `json:"short_url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
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
		Longurl: params.LongUrl,
		Short:   urlId,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB Error", err)
		return
	}

	respondWithJSON(w, http.StatusAccepted, response{
		Id:        url.ID,
		LongUrl:   url.Longurl,
		Short:     url.Short,
		CreatedAt: url.Createdat,
		UpdatedAt: url.Updatedat,
	})

}
