package main

import (
	"cryptorate/api"
	"cryptorate/scrapper"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load envs: ", err)
	}
	api.SetupRoutes()
	err = scrapper.PeriodicScrapping()
	if err != nil {
		log.Fatal("Cant add periodic task: ", err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
