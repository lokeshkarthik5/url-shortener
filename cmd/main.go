package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type Config struct {
	totalVisits atomic.Int32
}

func main() {

	c := Config{
		totalVisits: atomic.Int32{},
	}
	
	mux := http.NewServeMux()

	mux.HandleFunc("/")
	mux.HandleFunc("GET /api/health")
	mux.HandleFunc("POST /api/create-url")
	mux.HandleFunc("GET /api/get-url/{urlId}",c.middlewareMetrics())
	mux.HandleFunc("PUT /api/update-url/{urlId}")
	mux.HandleFunc("DELETE /api/delete-url/{urlId}")
	mux.HandleFunc("GET /api/get-stats/{urlId}",)

	srv := &http.Server{
		Addr:    ":" + "3001",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func (c *Config) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.totalVisits.Add(1)
		next.ServeHTTP(w, r)
	})
}
