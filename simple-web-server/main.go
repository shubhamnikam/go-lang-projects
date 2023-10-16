package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shubhamnikam/go-lang-projects/simple-web-server/utilities"
)

func main() {
	fmt.Println("server started...")

	log.Fatal(http.ListenAndServe(utilities.SERVER_URL, utilities.SetupRouter()))
}
