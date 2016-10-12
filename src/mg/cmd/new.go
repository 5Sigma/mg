package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"mg/console"
	"mg/project"
	"os"
	"strings"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "add a new migration",
	Long: `Create a new migration file. This will generate a new file named
with the version and description you provide. It will also be opened in the
default editor.

Descriptions should be a short name using a couple of words. Something like:

	create users
	add amount to items
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var template = `
----UP
-- Your migration command here

----DOWN
-- Your rollback command here
		`
		var name string
		if len(args) == 0 {
			console.Printf("Enter a short name for the migration: For example %s",
				chalk.Yellow.Color("create customer table"))
			name = console.Ask("Migration name:")
		} else {
			name = strings.Join(args, " ")
		}
		mFile := project.NewMigration(name)
		f, err := os.Create(mFile.FilePath)
		if err != nil {
			console.Fatalf("Could not write file: %s\n%s", mFile.FilePath, err.Error())
		}
		_, err = f.Write([]byte(template))
		if err != nil {
			console.Fatalf("Could not write file: %s\n%s", mFile.FilePath, err.Error())
		}
		f.Close()
		mFile.Edit()
	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
