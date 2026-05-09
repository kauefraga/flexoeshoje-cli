package infra

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kauefraga/flexoeshoje-cli/internal/entities"
)

func FindTodayPushups(db *sql.DB) ([]entities.Pushup, error) {
	query := `SELECT * FROM pushups WHERE created_at >= ?`
	today := time.Now().Format("2006-01-02")

	rows, err := db.Query(query, today)
	if err != nil {
		return nil, fmt.Errorf("Ocorreu um erro ao buscar todos os registros de flexões: %v", err)
	}
	defer rows.Close()

	var pushups []entities.Pushup

	for rows.Next() {
		var p entities.Pushup
		var createdAt string
		if err := rows.Scan(&p.Id, &p.Repetitions, &p.Type, &createdAt); err != nil {
			return nil, fmt.Errorf("Ocorreu um erro durante o scan do registro de flexão: %v", err)
		}

		p.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, fmt.Errorf("Ocorreu um erro ao converter a data do registro: %v", err)
		}

		pushups = append(pushups, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Ocorreu um erro ao manipular os registros de flexões: %v", err)
	}

	return pushups, nil
}

func CreateOnePushup(db *sql.DB, pushup entities.NewPushup) error {
	query := `INSERT INTO pushups (repetitions, type, created_at)
		VALUES ($1, $2, $3)`

	_, err := db.Exec(
		query,
		pushup.Repetitions,
		pushup.Type,
		pushup.CreatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return fmt.Errorf("Ocorreu um erro ao registrar suas flexões: %v", err)
	}

	return nil
}
