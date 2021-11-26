package storage

import (
	"fmt"
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

func Select() error {
	connectionString := os.Getenv("MYSQL_CONN")
	db, err := DBConnect(connectionString)
	if err != nil {
		log.Fatal("Cant connect to database: ", err)
		return err
	}
	query := "SELECT * FROM cryptocurrency_rate"
	res, err := db.Query(query)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
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
