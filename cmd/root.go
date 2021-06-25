package cmd

import (
	"fmt"
	"my-test-app/app"

	"github.com/spf13/cobra"
)

var config *string

var Root = &cobra.Command{
	Use:   "my-test-app",
	Short: "Тестовое приложение",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app.InitApp(*config)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Type 'help' for usage details")
	},
}

func init() {
	config = Root.PersistentFlags().StringP("config", "c", app.Name+".yml", "config file name *.yml, *.json")
}
