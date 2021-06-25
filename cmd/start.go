package cmd

import (
	"my-test-app/app"
	"my-test-app/migrations"
	"os"

	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "Starts my-test-app server",
	Run: func(cmd *cobra.Command, args []string) {
		app.Start()
	},
}

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Start migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := app.DB()
		if err != nil {
			panic(err)
		}
		err = migrations.Migrate(db, args...)
		if err != nil {
			os.Exit(-1)
		}
	},
}

func init() {
	Root.AddCommand(start, migrate)
}
