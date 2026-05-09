package main

import (
	"log"

	"github.com/kauefraga/flexoeshoje-cli/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "flexoeshoje",
	Version: "1.1.0",
	Aliases: []string{"fh"},
	Short:   "Registre suas flexões diárias",
	Long:    "Registre quantas flexões de braço você executou hoje sem sair do terminal.",
	Example: `  flexoeshoje
  flexoeshoje 25
  flexoeshoje 5 --subtrair`,
	Args: cobra.MaximumNArgs(1),
	Run:  commands.RootCommand,
}

func init() {
	rootCmd.Flags().BoolP("subtrair", "s", false, "subtrai flexões")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Ocorreu um erro durante a execução")
	}
}
