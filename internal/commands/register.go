package commands

import (
	"log"

	"github.com/kauefraga/flexoeshoje-cli/v2/internal/infra"
	"github.com/kauefraga/flexoeshoje-cli/v2/internal/usecases"
	"github.com/spf13/cobra"
)

var RegisterCmd = &cobra.Command{
	Use:     "registro",
	Aliases: []string{"reg", "r"},
	Short:   "Exibe todas as flexões registradas",
	Run:     runRegisterCmd,
}

func runRegisterCmd(cmd *cobra.Command, args []string) {
	db := infra.ConnectToDB()
	defer db.Close()

	if err := usecases.ListPushups(db); err != nil {
		log.Fatalln(err)
	}
}
