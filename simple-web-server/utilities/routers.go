package utilities

import (
	"github.com/gorilla/mux"
	"github.com/shubhamnikam/go-lang-projects/simple-web-server/controllers"
)

func SetupRouter() (*mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/demo", controllers.GetDemo).Methods("GET")
	router.HandleFunc("/create", controllers.PostCreateStudent).Methods("POST")

	return router
}