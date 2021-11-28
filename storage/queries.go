package storage

import (
	"log"
)

//Insert - inserts processed priceMultyFull response as raw
// and display to db
func Insert(cryptocurrencies PriceMultyFull) error {
	db, err := DBConnect()
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
		return err
	}
	defer db.Close()
	query := "INSERT INTO cryptocurrency_rate(raw, display) VALUES(?, ?)"
	_, err = db.Query(query, cryptocurrencies.Raw, cryptocurrencies.Display)
	if err != nil {
		log.Print("INSERT DATA ERROR: ", err)
		return err
	}
	return nil
}

//Select - selects raw and display
func Select() (raw, display []byte, err error) {
	db, err := DBConnect()
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
		return
	}
	query := "SELECT raw, display FROM cryptocurrency_rate ORDER BY id DESC LIMIT 1"
	res, err := db.Query(query)
	if err != nil {
		log.Print("SELECT Error: ", err)
		return
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(&raw, &display)
		if err != nil {
			log.Print("Scan error: ", err)
			return
		}
	}

	return
}
