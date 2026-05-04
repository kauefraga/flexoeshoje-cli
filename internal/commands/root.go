package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/internal/entities"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// This command does two things: show today's push up repetitions if called without args and register repetitions if called with args
func RootCommand(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	if len(args) == 0 {
		pushups, err := infra.FindTodayPushups(db)
		if err != nil {
			log.Fatalln(err)
		}

		if len(pushups) == 0 {
			highlight := color.New(color.FgHiYellow, color.Bold)

			fmt.Printf("Nenhuma flexão de braço executada hoje. %s\n", highlight.Sprint("Bora lá! 10 ZERO, GUERREIRO!!"))
			return
		}

		var totalReps int

		for _, p := range pushups {
			if p.Type == "subtract" {
				color.Magenta("[%v] menos %d repetições\n", p.CreatedAt.Format("15:04:05"), p.Repetitions)
				totalReps -= p.Repetitions
				continue
			}

			color.Magenta("[%v] mais %d repetições\n", p.CreatedAt.Format("15:04:05"), p.Repetitions)
			totalReps += p.Repetitions
		}

		today := time.Now().Local().Format("02-01-2006")

		fmt.Printf(
			"Hoje você fez %s flexões de braço (%s)\n",
			color.HiCyanString("%d", totalReps),
			color.HiCyanString("%v", today),
		)
		return
	}

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

	err = infra.CreateOnePushup(db, newPushup)
	if err != nil {
		log.Fatalln("Ocorreu um erro ao registrar suas flexões")
	}

	fmt.Printf("Suas flexões foram registradas com sucesso, %s\n", color.GreenString("EXCELENTE!"))
}
