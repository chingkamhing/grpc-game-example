package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "T Shooter game",
	Run: func(cmd *cobra.Command, args []string) {
		// default command
		cmd.Usage()
	},
}

const version string = "1.0.0"

var address string
var port int
var password string
var numBots int

func init() {
	// set --version flag
	rootCmd.Version = version
	rootCmd.InitDefaultVersionFlag()
}

// Execute init cli commands, flags and read configuration
func Execute() {
	// run root command
	if err := rootCmd.Execute(); err != nil {
		log.Print(err)
	}
}
