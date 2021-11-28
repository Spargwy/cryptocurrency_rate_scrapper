package service

import (
	"cryptorate/api"
	"cryptorate/scrapper"
	"log"
	"net/http"
	"os"

	"github.com/sevlyar/go-daemon"
)

//Run - run cryptocompare scrapper service
func Run(port string) {
	if len(os.Args) == 2 {
		if os.Args[1] != "-d" {
			log.Fatalf("Unexpected argument: %s", os.Args[1])
		}
		cntxt := &daemon.Context{
			PidFileName: "sample.pid",
			PidFilePerm: 0644,
			LogFileName: "sample.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
			return
		}

		log.Print("Running as daemon")

		if d != nil {
			return
		}
		defer cntxt.Release()
	}
	if len(os.Args) > 2 {
		log.Fatal("Only '-d' argument available")
	}
	err := scrapper.PeriodicScrapping()
	if err != nil {
		log.Fatal("Cant add periodic task: ", err)
	}
	api.SetupRoutes()
	log.Fatal(http.ListenAndServe(port, nil))
}
