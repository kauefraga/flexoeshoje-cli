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

func CreatePushupsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS pushups (
		id INTEGER PRIMARY KEY,
		repetitions INTEGER NOT NULL,
		type TEXT NOT NULL CHECK(type in ('add', 'subtract')) DEFAULT 'add',
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := db.Exec(query); err != nil {
		log.Fatalln("Ocorreu um erro ao criar a tabela de flexões")
	}
}
