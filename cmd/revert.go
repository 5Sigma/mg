package cmd

import (
	"github.com/5sigma/mg/console"
	"github.com/5sigma/mg/project"
	"github.com/spf13/cobra"
)

// revertCmd represents the revert command
var revertCmd = &cobra.Command{
	Use:   "revert [version]",
	Short: "Revert back migrations",
	Long: `Undo the last migration or the last several migrations.

When run alone it will rollback the last migration. Optionally a version number
can be specified and migrations will be rolled back to that version.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var version = ""
		if len(args) > 0 {
			version = args[0]
		}
		err := project.Revert(version)
		if err != nil {
			console.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(revertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// revertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// revertCmd.Flags().BoolP("all", "a", false, "Revert all migrations")

}
