package infra

import (
	"database/sql"
	"time"
)

func FindTodayPushups(db *sql.DB) (int, error) {
	total := 0

	query := `SELECT COALESCE(SUM(repetitions), 0) FROM pushups WHERE created_at >= ? LIMIT 1`
	today := time.Now().Format("2006-01-02")

	err := db.QueryRow(query, today).Scan(&total)

	return total, err
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
