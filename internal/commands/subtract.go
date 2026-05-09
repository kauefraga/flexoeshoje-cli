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

var SubtractCmd = &cobra.Command{
	Use:     "subtrair",
	Aliases: []string{"sub", "s"},
	Short:   "Remove flexões adicionadas incorretamente",
	Example: `  flexoeshoje subtrair 10
  flexoeshoje sub 2
  flexoeshoje s 2`,
	Args: cobra.ExactArgs(1),
	Run:  runSubtractCmd,
}

func runSubtractCmd(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	reps, err := strconv.Atoi(args[0])
	if err != nil || reps <= 0 {
		color.Red("Apenas números positivos são aceitos! Exemplo: flexoeshoje subtrair 10\n")
		os.Exit(1)
	}

	newPushup := entities.NewPushup{
		Repetitions: reps,
		Type:        entities.OpSubtract,
		CreatedAt:   time.Now(),
	}

	if err := usecases.CreatePushup(db, newPushup); err != nil {
		log.Fatalln(err)
	}
}
