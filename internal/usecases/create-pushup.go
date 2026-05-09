package usecases

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/internal/entities"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
)

func CreatePushup(db *sql.DB, pushup entities.NewPushup) error {
	err := infra.CreateOnePushup(db, pushup)
	if err != nil {
		return err
	}

	if pushup.Type == entities.OpAdd {
		fmt.Printf("Suas flexões foram registradas com sucesso, %s\n", color.GreenString("EXCELENTE!"))
	}

	if pushup.Type == entities.OpSubtract {
		fmt.Printf("As flexões foram removidas com sucesso, %s\n", color.GreenString("EXCELENTE!"))
	}

	return nil
}
