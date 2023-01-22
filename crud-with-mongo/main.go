package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shubhamnikam/go-small-projects/curd-with-mongo/routers"
)

func main() {
	fmt.Println("API started...")
	
	r := routers.Router()

	log.Fatal(http.ListenAndServe(":4000", r))
}