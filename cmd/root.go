package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd is the base command for your CLI tool
var rootCmd = &cobra.Command{
	Use:   "RavenRecon",
	Short: "RavenRecon is an OSINT tool for scanning websites",
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, print the usage
		fmt.Println("Welcome to RavenRecon. Use 'scanuser' command to start scanning.")
	},
}

// Execute initializes the commands and starts the application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
