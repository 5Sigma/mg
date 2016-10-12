package cmd

import (
	"github.com/spf13/cobra"
	"mg/console"
	"mg/project"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate the database",
	Long:  `Run all migrations that have not yet been run on the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := project.Migrate()
		if err != nil {
			console.Error(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
