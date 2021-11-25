package storage

import (
	"log"
	"os"
)

func Insert(cryptocurrencies PriceMultyFull) error {
	connectionString := os.Getenv("MYSQL_CONN")
	db, err := DBConnect(connectionString)
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
		return err
	}
	if cryptocurrencies.Display == nil || cryptocurrencies.Raw == nil {
		log.Print("cryptocurrencies is nil")
		return nil
	}
	defer db.Close()
	query := "INSERT INTO cryptocurrencies(raw, display) VALUES(?, ?)"
	_, err = db.Query(query, cryptocurrencies.Raw, cryptocurrencies.Display)
	if err != nil {
		log.Print("INSERT DATA ERROR: ", err)
		return err
	}
	return nil
}

// func Select() error {
// 	connectionString := os.Getenv("MYSQL_CONN")
// 	db, err := DBConnect(connectionString)
// 	if err != nil {
// 		log.Fatal("Cant connect to database: ", err)
// 		return err
// 	}
// 	query := "SELECT * FROM cryptocurrencies"
// 	res, err := db.Query(query)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
