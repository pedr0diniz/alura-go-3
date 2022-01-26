package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	connection := "user=postgres dbname=alura_store password=12345678 host=localhost sslmode=disable port=5433"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
