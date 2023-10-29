package main

import (
	"github.com/joho/godotenv"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/services"
	"github.com/shubhamnikam/go-lang-projects/crud-with-fiber-gorm-sqlite/configs"
)

func main() {
	godotenv.Load("./configs/.env")
	
	services.DbConnect()

	configs.SetupServer()
}