package storage

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

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
	queries, err := parseQueries()
	if err != nil {
		log.Print("PARSE QUERIES ERROR: ", err)
	}
	for _, query := range queries {
		_, err = db.Exec(string(query))
		if err != nil {
			log.Print("Query exec error: ", err)
			return err
		}
	}

	return nil
}

func parseQueries() (splitedqueries []string, err error) {
	queries, err := ioutil.ReadFile("storage/schema.sql")
	if err != nil {
		log.Print("ReadFile error: ", err)
		return
	}
	splitedqueries = strings.Split(string(queries), ";")
	for i := range splitedqueries {
		if strings.TrimSpace(splitedqueries[i]) == "" {
			splitedqueries = append(splitedqueries[:i], splitedqueries[i+1:]...)
		}
	}
	for i := range splitedqueries {
		splitedqueries[i] += ";"
	}

	return
}
func DBConnect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("CANT CONNECT TO DB: ", err)
		return &sql.DB{}, err
	}

	return db, nil
}
