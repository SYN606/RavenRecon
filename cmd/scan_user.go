package cmd

import (
	"RavenRecon/functions"
	"github.com/spf13/cobra"
	"log"
)

var username string

// scanUserCmd represents the scan-user command
var scanUserCmd = &cobra.Command{
	Use:   "scanuser",
	Short: "Scan for username across websites",
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" {
			log.Fatal("You must provide a username")
		}
		err := functions.SearchUserAcrossWebsites(username)
		if err != nil {
			log.Fatal("Error scanning websites:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanUserCmd)

	scanUserCmd.Flags().StringVarP(&username, "username", "u", "", "The username to scan for")
}
