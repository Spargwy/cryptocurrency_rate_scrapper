package scrapper

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
)

func cronInit() *cron.Cron {
	cron := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	return cron
}

//PeriodicScrapping - sets up periodic scrapping data from cryptocompare
func PeriodicScrapping() error {
	cron := cronInit()
	params := ParseAvailableParams()
	period := os.Getenv("CRONTAB_FOR_SCRAPPING")
	if period == "" {
		period = "* * * * *"
	}
	_, err := cron.AddFunc(period, func() {
		_, err := GetCurrentPrice(params, true, true)
		if err != nil {
			log.Print("Cant get current price: ", err)
		}
	})
	if err != nil {
		log.Print("AddFunc error: ", err)
		return err
	}
	cron.Start()
	return nil
}
