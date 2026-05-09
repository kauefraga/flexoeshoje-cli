package commands

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/v2/internal/entities"
	"github.com/kauefraga/flexoeshoje-cli/v2/internal/infra"
	"github.com/kauefraga/flexoeshoje-cli/v2/internal/usecases"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:     "adicionar",
	Aliases: []string{"add", "a"},
	Short:   "Registre suas flexões diárias",
	Example: `  flexoeshoje adicionar 10
  flexoeshoje add 5
  flexoeshoje a 5`,
	Args: cobra.ExactArgs(1),
	Run:  runAddCmd,
}

func runAddCmd(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	reps, err := strconv.Atoi(args[0])
	if err != nil || reps <= 0 {
		color.Red("Apenas números positivos são aceitos! Exemplo: flexoeshoje adicionar 10\n")
		os.Exit(1)
	}

	newPushup := entities.NewPushup{
		Repetitions: reps,
		Type:        entities.OpAdd,
		CreatedAt:   time.Now(),
	}

	if err := usecases.CreatePushup(db, newPushup); err != nil {
		log.Fatalln(err)
	}
}
