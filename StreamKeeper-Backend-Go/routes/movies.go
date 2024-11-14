package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"your_project_name/helpers" // Adjust the import path based on your project structure
	"your_project_name/models"

	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie API is running"})
}

func FetchMovieImagesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/images", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie images", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func FetchMovieCreditsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/credits", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie credits", http.StatusInternalServerError)
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(data, &response); err != nil {
		http.Error(w, "Failed to parse movie credits response", http.StatusInternalServerError)
		return
	}

	cast := []models.Person{}
	crew := []models.Person{}
	if castData, ok := response["cast"].([]interface{}); ok {
		for _, person := range castData {
			cast = append(cast, *models.NewPerson(person.(map[string]interface{})))
		}
	}
	if crewData, ok := response["crew"].([]interface{}); ok {
		for _, person := range crewData {
			crew = append(crew, *models.NewPerson(person.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"cast": cast, "crew": crew})
}

func FetchMovieExternalIDsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/external_ids", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie external IDs", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func FetchMovieRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/recommendations", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie recommendations", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse movie recommendations response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchMovieReviewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/reviews", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie reviews", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse movie reviews response", http.StatusInternalServerError)
		return
	}

	reviews := []models.Review{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, reviewData := range items {
			reviews = append(reviews, *models.NewReview(reviewData.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(reviews)
}

func FetchSimilarMoviesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/similar", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch similar movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse similar movies response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchMovieVideosHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/videos", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie videos", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func FetchWatchProvidersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s/watch/providers", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie watch providers", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func SearchMoviesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/search/movie?query=%s", query))
	if err != nil {
		http.Error(w, "Failed to search movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse search results", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchPopularMoviesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := helpers.FetchFromTmdb("/movie/popular")
	if err != nil {
		http.Error(w, "Failed to fetch popular movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse popular movies response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchNowPlayingMoviesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := helpers.FetchFromTmdb("/movie/now_playing")
	if err != nil {
		http.Error(w, "Failed to fetch now playing movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse now playing movies response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchTopRatedMoviesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := helpers.FetchFromTmdb("/movie/top_rated")
	if err != nil {
		http.Error(w, "Failed to fetch top-rated movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse top-rated movies response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchUpcomingMoviesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := helpers.FetchFromTmdb("/movie/upcoming")
	if err != nil {
		http.Error(w, "Failed to fetch upcoming movies", http.StatusInternalServerError)
		return
	}

	var results map[string]interface{}
	if err := json.Unmarshal(data, &results); err != nil {
		http.Error(w, "Failed to parse upcoming movies response", http.StatusInternalServerError)
		return
	}

	movies := []models.Movie{}
	if items, ok := results["results"].([]interface{}); ok {
		for _, item := range items {
			movies = append(movies, *models.NewMovie(item.(map[string]interface{})))
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func FetchMovieDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieID := vars["movie_id"]

	data, err := helpers.FetchFromTmdb(fmt.Sprintf("/movie/%s", movieID))
	if err != nil {
		http.Error(w, "Failed to fetch movie details", http.StatusInternalServerError)
		return
	}

	var movieData map[string]interface{}
	if err := json.Unmarshal(data, &movieData); err != nil {
		http.Error(w, "Failed to parse movie details response", http.StatusInternalServerError)
		return
	}

	movie := models.NewMovie(movieData)
	json.NewEncoder(w).Encode(movie)
}

func RegisterMovieRoutes(router *mux.Router) {
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/images", FetchMovieImagesHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/credits", FetchMovieCreditsHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/external_ids", FetchMovieExternalIDsHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/recommendations", FetchMovieRecommendationsHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/reviews", FetchMovieReviewsHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/similar", FetchSimilarMoviesHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/videos", FetchMovieVideosHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}/watch/providers", FetchWatchProvidersHandler).Methods("GET")
	router.HandleFunc("/search", SearchMoviesHandler).Methods("GET")
	router.HandleFunc("/popular", FetchPopularMoviesHandler).Methods("GET")
	router.HandleFunc("/now_playing", FetchNowPlayingMoviesHandler).Methods("GET")
	router.HandleFunc("/top_rated", FetchTopRatedMoviesHandler).Methods("GET")
	router.HandleFunc("/upcoming", FetchUpcomingMoviesHandler).Methods("GET")
	router.HandleFunc("/{movie_id:[0-9]+}", FetchMovieDetailsHandler).Methods("GET")
}
