package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "elencdeo:12EA56@tcp(localhost:3306)/tcc_outlier")
	if err != nil {
		panic(err.Error())
	}
	return db
}