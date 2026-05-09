package usecases

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
)

func ListPushups(db *sql.DB) error {
	pushups, err := infra.FindTodayPushups(db)
	if err != nil {
		return err
	}

	if len(pushups) == 0 {
		highlight := color.New(color.FgHiYellow, color.Bold)
		fmt.Printf("Nenhuma flexão de braço executada hoje. %s\n", highlight.Sprint("Bora lá! 10 ZERO, GUERREIRO!!"))

		return nil
	}

	var totalReps int

	for _, p := range pushups {
		executedAt := p.CreatedAt.Format("15:04:05")

		if p.Type == "subtract" {
			color.Magenta("[%v] menos %d repetições\n", executedAt, p.Repetitions)
			totalReps -= p.Repetitions
			continue
		}

		color.Magenta("[%v] mais %d repetições\n", executedAt, p.Repetitions)
		totalReps += p.Repetitions
	}

	today := time.Now().Local().Format("02-01-2006")

	fmt.Printf(
		"Hoje você fez %s flexões de braço (%s)\n",
		color.HiCyanString("%d", totalReps),
		color.HiCyanString("%v", today),
	)

	return nil
}
