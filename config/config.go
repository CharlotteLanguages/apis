package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	db, error := sql.Open("mysql", "root:btPy7MOat7IPPA6aNBdp@tcp(containers-us-west-153.railway.app:5905)/railway")

	if error != nil {
		log.Fatal(error)
	}
	return
}
