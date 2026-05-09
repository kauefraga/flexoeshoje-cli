package infra

import (
	"database/sql"
	"log"
	"os"

	"github.com/fatih/color"
)

func checkIfDBExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func createPushupsTable(db *sql.DB) {
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

func ConnectToDB() *sql.DB {
	// Check before it's automatically created by sql.Open
	hasDB := checkIfDBExists("./flexoeshoje.db")

	db, err := sql.Open("sqlite", "./flexoeshoje.db")

	if err != nil {
		log.Fatalln("Ocorreu um erro ao conectar com o banco de dados")
	}

	if !hasDB {
		color.Black("Banco de dados não encontrado. Um novo foi criado em `./flexoeshoje.db`.")
		createPushupsTable(db)
		color.Black("Preparando banco de dados... OK")
	}

	return db
}
