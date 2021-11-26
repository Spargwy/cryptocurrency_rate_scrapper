package storage

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Init - open database
func InitDB(connectionString string) error {
	err := migrate(connectionString)
	if err != nil {
		log.Print("Cant execute migrations: ", err)
		return err
	}
	return nil
}

func migrate(connectionString string) error {
	db, err := DBConnect(connectionString)
	if err != nil {
		log.Print("InitDB error: ", err)
		return err
	}
	query, err := ioutil.ReadFile("storage/schema.sql")
	if err != nil {
		log.Print("ReadFile error: ", err)
		return err
	}
	_, err = db.Exec(string(query))
	if err != nil {
		log.Print("Query exec error: ", err)
		return err
	}

	return nil
}
func DBConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("CANT CONNECT TO DB: ", err)
		return &sql.DB{}, err
	}

	return db, nil
}
