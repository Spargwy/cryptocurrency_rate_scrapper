package storage

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Init - open database
func DBConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("CANT CONNECT TO DB: ", err)
		return &sql.DB{}, err
	}
	return db, nil
}
