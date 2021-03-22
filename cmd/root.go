package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "example-service",
	Short: "A RESTful API boilerplate",
	Long:  `A RESTful API boilerplate.......`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(ServeCmd)
	RootCmd.AddCommand(CmdMigrateMysql)
	RootCmd.AddCommand(CmdRollbackMysql)
	RootCmd.AddCommand(CmdMigratePostgres)
	RootCmd.AddCommand(CmdRollbackPostgres)
}
