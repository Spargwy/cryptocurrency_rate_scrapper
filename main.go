package main

import (
	"cryptorate/service"
	"cryptorate/storage"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cant load envs: ", err)
	}
	log.Print("Envs loaded")
	err = storage.InitDB()
	if err != nil {
		log.Fatal("InitDB error: ", err)
	}
	log.Print("DB inited")

}

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	service.Run(port)
}
