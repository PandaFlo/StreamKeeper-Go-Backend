package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Controller interfaces to simulate different controllers for Movies, TV Shows, and People
type Controller interface {
	RegisterRoutes(router *mux.Router)
}

// TMDBController handles generic TMDb routes
type TMDBController struct{}

func (c *TMDBController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/validate", c.ValidateAPIKey).Methods("GET")
}

func (c *TMDBController) ValidateAPIKey(w http.ResponseWriter, r *http.Request) {
	tmdbBaseURL := os.Getenv("TMDB_BASE_URL")
	apiKey := os.Getenv("TMDB_API_KEY")

	resp, err := http.Get(fmt.Sprintf("%s/configuration?api_key=%s", tmdbBaseURL, apiKey))
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "API key is invalid or TMDB service is unavailable.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("API key is valid."))
}

// MovieController handles movie-specific routes
type MovieController struct{}

func (c *MovieController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", c.HandleMovies).Methods("GET")
}

func (c *MovieController) HandleMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movies endpoint"))
}

// TVShowController handles TV-specific routes
type TVShowController struct{}

func (c *TVShowController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", c.HandleTVShows).Methods("GET")
}

func (c *TVShowController) HandleTVShows(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TV Shows endpoint"))
}

// PersonController handles person-specific routes
type PersonController struct{}

func (c *PersonController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", c.HandlePersons).Methods("GET")
}

func (c *PersonController) HandlePersons(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Persons endpoint"))
}

// Load environment variables from .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system environment variables instead.")
	}
}

// StartServer starts a new HTTP server with the given port, controller, and route path
func startServer(port string, controller Controller, routePath string) {
	router := mux.NewRouter()

	// If serving TMDb, add static file server for view directory
	if port == "3001" {
		viewPath, err := filepath.Abs("view")
		if err != nil {
			log.Fatalf("Error determining view path: %v", err)
		}
		fileServer := http.FileServer(http.Dir(viewPath))
		router.PathPrefix("/").Handler(fileServer)
	}

	subRouter := router.PathPrefix(routePath).Subrouter()
	controller.RegisterRoutes(subRouter)

	log.Printf("Server for %s running at http://localhost:%s", routePath, port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	loadEnv()

	go startServer("3001", &TMDBController{}, "/api")
	go startServer("3002", &MovieController{}, "/api/movies")
	go startServer("3003", &TVShowController{}, "/api/tv")
	go startServer("3004", &PersonController{}, "/api/person")

	// Keep the main goroutine alive
	select {}
}
