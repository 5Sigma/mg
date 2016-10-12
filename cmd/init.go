package cmd

import (
	"github.com/5sigma/mg/console"
	"github.com/5sigma/mg/project"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new project",
	Long: `Create a new migration project.
This will create new 'db' folder in the current directory which will be
used to store migration files. A configuration file will also be placed
in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := project.CreateProject("")
		if err != nil {
			console.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
