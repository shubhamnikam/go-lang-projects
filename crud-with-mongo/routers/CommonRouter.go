package routers

import (
	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/controllers"
	
	"github.com/gorilla/mux"
)

func Router() (*mux.Router){
	router := mux.NewRouter()

	//setup route
	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.UpdateMovieAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteMovieById).Methods("DELETE")

	return router
}