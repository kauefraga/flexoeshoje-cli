package commands

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
	"github.com/spf13/cobra"
	_ "modernc.org/sqlite"
)

// This command does two things: show today's push up repetitions if called without args and register repetitions if called with args
func RootCommand(cmd *cobra.Command, args []string) {
	currentDate := time.Now().Local()

	db := infra.ConnectToDB()
	defer db.Close()

	infra.CreatePushupsTable(db)

	if len(args) == 0 {
		reps := infra.FindTodayPushups(db)

		if reps == 0 {
			fmt.Println("Nenhuma flexão de braço executada hoje. Bora lá!")
			return
		}

		fmt.Printf("Flexões de braço hoje (%v): %d\n", currentDate.Format("2006-01-02"), reps)
		return
	}

	reps, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln("Apenas números inteiros são aceitos! Exemplo: flexoeshoje 10")
	}

	err = infra.CreateOnePushup(db, reps)

	if err != nil {
		log.Fatalln("Ocorreu um erro ao registrar suas flexões")
	}

	fmt.Println("Suas flexões foram registradas com sucesso")
}
