package go_crud

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnection() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:@(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}

	return db
}
