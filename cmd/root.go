package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "chaoswall",
		Short: "ChaosWall",
		Long:  "ChaosWall",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(
		serveCmd,
		versionCmd,
	)

	rootCmd.PersistentFlags().StringP("config", "c", "", "Configuration path")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}
