package infra

import (
	"database/sql"
	"log"
	"time"
)

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
