package main

import (
	"log"

	"github.com/kauefraga/flexoeshoje-cli/v2/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "flexoeshoje",
	Version: "2.0.2",
	Short:   "Registre suas flexões diárias",
	Long:    "Registre quantas flexões de braço você executou hoje sem sair do terminal.",
	Example: `  flexoeshoje registro
  flexoeshoje r
  flexoeshoje adicionar 30
  flexoeshoje subtrair 5`,
}

func init() {
	rootCmd.AddCommand(commands.RegisterCmd)
	rootCmd.AddCommand(commands.AddCmd)
	rootCmd.AddCommand(commands.SubtractCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Ocorreu um erro durante a execução")
	}
}
