package infra

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	_ "modernc.org/sqlite"
)

func checkIfDBExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func findDbPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln("Ocorreu um erro ao acessar o diretório das configurações")
	}

	appDir := filepath.Join(configDir, "flexoeshoje-cli")

	err = os.MkdirAll(appDir, 0755)
	if err != nil {
		log.Fatalln("Ocorreu um erro ao criar o diretório da ferramenta")
	}

	dbPath := filepath.Join(appDir, "flexoeshoje.db")

	return dbPath
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
	dbPath := findDbPath()

	// Check before it's automatically created by sql.Open
	hasDB := checkIfDBExists(dbPath)

	db, err := sql.Open("sqlite", dbPath)

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
