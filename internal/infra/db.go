package infra

import (
	"database/sql"
	"log"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("sqlite", "./flexoeshoje.db")

	if err != nil {
		log.Fatalln("Ocorreu um erro ao conectar com o banco de dados")
	}

	return db
}
