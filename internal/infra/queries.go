package infra

import (
	"database/sql"
	"log"
	"time"
)

func CreatePushupsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS pushups (
		id INTEGER PRIMARY KEY,
		repetitions INTEGER NOT NULL,
		last_modified TEXT DEFAULT CURRENT_TIMESTAMP,
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	if _, err := db.Exec(query); err != nil {
		log.Fatal("Ocorreu um erro ao criar a tabela de flexões.")
	}
}

func FindTodayPushups(db *sql.DB) int {
	var total int

	today := time.Now().Format("2006-01-02")
	query := `SELECT COALESCE(SUM(repetitions), 0) FROM pushups WHERE created_at >= ? LIMIT 1`

	err := db.QueryRow(query, today).Scan(&total)
	if err != nil {
		log.Fatal(err)
	}

	return total
}

func FindManyPushups(db *sql.DB) {

}

func CreateOnePushup(db *sql.DB, repetitions int) error {
	query := `INSERT INTO pushups (repetitions, created_at)
		VALUES ($1, $2)`

	_, err := db.Exec(
		query,
		repetitions,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return err
}
