package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/models"
	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/repositories"
	
	"github.com/gorilla/mux"
)

// controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := repositories.GetAllMovies()
	json.NewEncoder(w).Encode(result)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	result := repositories.GetMovieById(params["id"])
	json.NewEncoder(w).Encode(result)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusCreated))

	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	repositories.CreateMovie(movie)
	json.NewEncoder(w).Encode("Created Successfully")
}

func UpdateMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusOK))

	params := mux.Vars(r)

	repositories.UpdateMovieById(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusOK))

	params := mux.Vars(r)

	repositories.DeleteMovieById(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusOK))

	count:= repositories.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}