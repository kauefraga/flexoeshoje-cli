package commands

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/internal/entities"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
	"github.com/kauefraga/flexoeshoje-cli/internal/usecases"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// This command does two things: show today's push up repetitions if called without args and register repetitions if called with args
func RootCommand(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	if len(args) > 0 {
		newRepetitions, err := strconv.Atoi(args[0])
		if err != nil {
			color.Red("Apenas números inteiros são aceitos! Exemplo: flexoeshoje 10\n")
			os.Exit(1)
		}

		isToSubtract, err := cmd.Flags().GetBool("subtrair")
		if err != nil {
			log.Fatalln("Ocorreu um erro ao ler a flag subtrair")
		}

		newPushup := entities.NewPushup{
			Repetitions: newRepetitions,
			Type:        "add",
			CreatedAt:   time.Now(),
		}

		if isToSubtract {
			newPushup.Type = "subtract"
		}

		if err := usecases.CreatePushup(db, newPushup); err != nil {
			log.Fatalln(err)
		}

		return
	}

	if err := usecases.ListPushups(db); err != nil {
		log.Fatalln(err)
	}
}
