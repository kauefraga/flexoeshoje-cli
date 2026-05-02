package main

import (
	"log"

	"github.com/kauefraga/flexoeshoje-cli/internal/commands"
	"github.com/kauefraga/flexoeshoje-cli/internal/infra"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "flexoeshoje",
	Aliases: []string{"fh"},
	Short:   "Registre suas flexões diárias",
	Long:    "Registre quantas flexões de braço você executou hoje sem sair do terminal.",
	Run:     commands.RootCommand,
}

func init() {
	db := infra.ConnectToDB()
	infra.CreatePushupsTable(db)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Algo deu errado durante a execução")
	}
}
