package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/raeandres/golang-rest-product.git/model"
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

	log.Println("Creating tables..")

	/*Create Table*/
	createProductTable(db)

	/*Drop table */
	// tableName := "product"
	// dropTableErr := dropTable(db, tableName)

	// if dropTableErr != nil {
	// 	log.Fatalf("Failed to drop table %s: %v", tableName, dropTableErr)
	// }

	// fmt.Printf("Table %s dropped successfully!\n", tableName)

	DB = DBInstance{
		Db: db,
	}

}

// func dropTable(db *sql.DB, tableName string) error {

// 	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
// 	_, err := db.Exec(query)

// 	return err
// }

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

func GetAllProducts(db *sql.DB) string {

	data := []model.Product{}
	rows, err := db.Query(`SELECT id, name, product_type, picture, price, description FROM product`)

	if err != nil {
		log.Fatal("GET || Database Error: ", err)
		return ""
	}

	// defer rows.Close()
	// to scan DB values
	var Id, Name, ProductType, Picture, Description string
	var Price float64

	for rows.Next() {
		rows.Scan(&Id, &Name, &ProductType, &Picture, &Price, &Description)
		if err != nil {
			log.Fatal("GET || Parsing to model Error: ", err)
		}
		data = append(data, model.Product{ID: Id, Name: Name, ProductType: ProductType, Picture: Picture, Price: Price, Description: Description})
	}

	//convert struct into json string

	jsonString, jsonError := json.Marshal(data)

	if jsonError != nil {
		log.Fatal("GET || JSON Parsing error: ", jsonError)
		return ""
	}

	return string(jsonString)

}

func InsertProduct(db *sql.DB, product *model.Product) int {

	query := `INSERT INTO PRODUCT (name, product_type, picture, price, description) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.ProductType, product.Picture, product.Price, product.Description).Scan(&pk)

	if err != nil {
		log.Fatal("POST || Database Error: ", err)
		return 0
	}

	return pk
}

func UpdateProduct(db *sql.DB, product *model.Product) string {

	query := `UPDATE PRODUCT SET name=$1, product_type=$2, picture=$3, price=$4, description=$5 WHERE id=$6`

	_, err := db.Exec(query, product.Name, product.ProductType, product.Picture, product.Price, product.Description, product.ID)

	if err != nil {
		log.Fatal("POST || Database Error: ", err)
		return ""
	}

	jsonString, jsonError := json.Marshal(product)

	if jsonError != nil {
		log.Fatal("GET || JSON Parsing error: ", jsonError)
		return ""
	}

	return string(jsonString)
}
