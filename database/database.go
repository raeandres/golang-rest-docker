package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DBInstance struct {
	Db *sql.DB
}

var DB DBInstance

func ConnectDb() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("HOST_NAME"),
		os.Getenv("DB_NAME"))

	db, error := sql.Open("postgres", connStr)

	if error != nil {
		log.Fatal("Failed to connect to Database, \n", error)
		os.Exit(2)
	}

	log.Println("Connected To Db")

	defer db.Close()

	log.Println("Creating tables..")

	createProductTable(db)

	DB = DBInstance{
		Db: db,
	}

}

func createProductTable(db *sql.DB) {

	createTablequery := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY NOT NULL, 
		name VARCHAR(255), 
		product_type VARCHAR(100), 
		picture VARCHAR(255), 
		price NUMERIC(10,2), 
		description TEXT,
		created timestamp DEFAULT NOW()
	)`

	_, error := db.Exec(createTablequery)

	if error != nil {
		log.Fatal("Problem with: ", error)
		os.Exit(2)
		return
	}

	log.Println("No error")

}
