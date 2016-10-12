package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mg",
	Short: "A general purpose database migration tool.",
	Long: `
Mg aims to help manage database migrations in a technology agnostic way.

To setup a new project run "mg init" in the project's base folder.
For information on what this does run "mg help init".`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is mg.yaml)")

	RootCmd.PersistentFlags().StringP("database", "d", "", "Database name")
	viper.BindPFlag("dbname", RootCmd.PersistentFlags().Lookup("dbname"))

	RootCmd.PersistentFlags().StringP("host", "r", "localhost", "Database host")
	viper.BindPFlag("host", RootCmd.PersistentFlags().Lookup("host"))
	viper.SetDefault("host", "localhost")

	RootCmd.PersistentFlags().IntP("port", "p", 0, "Database port")
	viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", 0)

	RootCmd.PersistentFlags().StringP("driver", "x", "", "Database driver")
	viper.BindPFlag("driver", RootCmd.PersistentFlags().Lookup("driver"))
	viper.SetDefault("driver", "localhost")

	RootCmd.PersistentFlags().StringP("user", "u", "", "Database username")
	viper.BindPFlag("user", RootCmd.PersistentFlags().Lookup("user"))
	viper.SetDefault("user", "")

	RootCmd.PersistentFlags().StringP("password", "P", "", "Database password")
	viper.BindPFlag("password", RootCmd.PersistentFlags().Lookup("password"))
	viper.SetDefault("password", "")

	RootCmd.PersistentFlags().StringP("migration-path", "m", "db", "Migration file path")
	viper.BindPFlag("migration-path", RootCmd.PersistentFlags().Lookup("migration-path"))
	viper.SetDefault("migration-path", "db")

	viper.SetDefault("pgssl", "disable")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("mg")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("mg")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		if viper.GetString("driver") == "postgres" {
			viper.SetDefault("port", 5432)
		}
	}
}
