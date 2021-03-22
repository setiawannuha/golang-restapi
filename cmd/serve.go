package cmd

import (
	"example.com/app"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "start-server",
	Short: "start http server",
	Long:  "Starts a http server",
	Run: func(cmd *cobra.Command, args []string) {
		app.NewServer()
	},
}
