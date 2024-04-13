package db

import (
	"database/sql"
	"fmt"
  _"github.com/lib/pq"
)


func Db () (*sql.DB, error) {


	/* Do not expose connection string like this in production */
	connStr := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres",connStr)

	if err != nil {
		fmt.Println(err)
		db.Close()
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT, password TEXT)")

	if err != nil {
		fmt.Println(err)
	}

	return db, err
}