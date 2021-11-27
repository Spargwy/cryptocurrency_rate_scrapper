package main

import (
	"cryptorate/api"
	"cryptorate/scrapper"
	"cryptorate/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load envs: ", err)
	}
	err = storage.InitDB()
	if err != nil {
		log.Fatal("InitDB error: ", err)
	}
	api.SetupRoutes()
	err = scrapper.PeriodicScrapping()
	if err != nil {
		log.Fatal("Cant add periodic task: ", err)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
