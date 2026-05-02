package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// This command does two things: show today's push up repetitions if called without args and register repetitions if called with args
func RootCommand(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	if len(args) == 0 {
		reps := infra.FindTodayPushups(db)

		if reps == 0 {
			color.Magenta("Nenhuma flexão de braço executada hoje. Bora lá!\n")
			return
		}

		currentDate := time.Now().Local().Format("2006-01-02")

		fmt.Printf("Flexões de braço hoje (%v): %d\n", currentDate, reps)
		return
	}

	newRepetitions, err := strconv.Atoi(args[0])
	if err != nil {
		color.Red("Apenas números inteiros são aceitos! Exemplo: flexoeshoje 10\n")
		os.Exit(1)
	}

	err = infra.CreateOnePushup(db, newRepetitions)

	if err != nil {
		log.Fatalln("Ocorreu um erro ao registrar suas flexões")
	}

	color.Green("Suas flexões foram registradas com sucesso, EXCELENTE!\n")
}
