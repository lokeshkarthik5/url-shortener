package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	controllers "github.com/lokeshkarthik5/url-shortner/handlers"
	"github.com/lokeshkarthik5/url-shortner/internal/database"
)

type Config struct {
	totalVisits atomic.Int32
	handlers    *controllers.Controllers
}

func main() {

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening database", err)
	}
	dbQueries := database.New(dbConn)

	c := &Config{
		totalVisits: atomic.Int32{},
		handlers: &controllers.Controllers{
			DB: dbQueries,
		},
	}

	mux := http.NewServeMux()

	//mux.HandleFunc("/")
	mux.HandleFunc("GET /api/health", controllers.HealthCheck)

	mux.HandleFunc("POST /api/create-url", c.handlers.CreateUrl)
	mux.HandleFunc("GET /api/get-url/{urlId}")
	mux.HandleFunc("PUT /api/update-url/{urlId}")
	mux.HandleFunc("DELETE /api/delete-url/{urlId}")
	mux.HandleFunc("GET /api/get-stats/{urlId}")

	srv := &http.Server{
		Addr:    ":" + "3001",
		Handler: mux,
	}
	log.Println("running on 3001")
	log.Fatal(srv.ListenAndServe())
}

func (c *Config) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.totalVisits.Add(1)
		next.ServeHTTP(w, r)
	})
}
