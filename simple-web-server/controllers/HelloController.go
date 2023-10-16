package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamnikam/go-lang-projects/simple-web-server/models"
)

func GetDemo(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusOK))

	json.NewEncoder(w).Encode("Hi from demo")
}

func PostCreateStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("StatusCode", http.StatusText(http.StatusCreated))

	var student models.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	
	json.NewEncoder(w).Encode(student.Name + " - created successfully")
}