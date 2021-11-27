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
	defer db.Close()
	query := "INSERT INTO cryptocurrency_rate(raw, display) VALUES(?, ?)"
	_, err = db.Query(query, cryptocurrencies.Raw, cryptocurrencies.Display)
	if err != nil {
		log.Print("INSERT DATA ERROR: ", err)
		return err
	}
	return nil
}

func Select() (raw, display []byte, err error) {
	connectionString := os.Getenv("MYSQL_CONN")
	db, err := DBConnect(connectionString)
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

func Update() error {
	connectionString := os.Getenv("MYSQL_CONN")
	db, err := DBConnect(connectionString)
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
		return err
	}
	query := ""
	_, err = db.Exec(query)
	if err != nil {
		log.Print("Exec error in update: ", err)
	}
	return nil
}
