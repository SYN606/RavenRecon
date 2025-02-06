package cmd

import (
	"RavenRecon/utils"
	"github.com/spf13/cobra"
)

var username string

var scanUserCmd = &cobra.Command{
	Use:   "scan-user",
	Short: "Scan websites for user information",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username = args[0]
		utils.ScanWebsites(username)
	},
}

func init() {
	rootCmd.AddCommand(scanUserCmd)
}
