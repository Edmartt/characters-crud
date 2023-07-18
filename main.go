package main

import (
	"log"
	"os"

	"github.com/edmartt/characters-crud/internal/api/server"
	"github.com/edmartt/characters-crud/internal/database"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("env vars error", err.Error())
	}

	database.Init() //migrations
	server.Init(os.Getenv("PORT"))

}
