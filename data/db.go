package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "root"
	password1 = "password"
	dbname    = "postgres"
)

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password1, dbname)
	db, err := sql.Open("postgres", dbinfo)

	CheckError(err)

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
