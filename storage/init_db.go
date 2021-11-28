package storage

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" //need to use mysql driver
)

//InitDB - open database and execute sql file
func InitDB() error {
	err := migrate()
	if err != nil {
		log.Print("Cant execute migrations: ", err)
		return err
	}
	return nil
}

func migrate() error {
	db, err := DBConnect()
	if err != nil {
		log.Print("InitDB error: ", err)
		return err
	}

	//Separating several queries
	queries, err := parseQueries()
	if err != nil {
		log.Print("PARSE QUERIES ERROR: ", err)
		return err
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

//parseQueries for parse several queries from schema.sql file
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

//DBConnect - connects to db
func DBConnect() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("CANT CONNECT TO DB: ", err)
		return &sql.DB{}, err
	}

	return db, nil
}
