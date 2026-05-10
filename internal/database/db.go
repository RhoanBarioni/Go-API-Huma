package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Db() (*sql.DB, error) {
	dbAdress := "root:123@tcp(localhost:3306)/escola"
	db, err := sql.Open("mysql", dbAdress)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
