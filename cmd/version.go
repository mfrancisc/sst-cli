package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of SST CLI",
	Long:  `All software has versions. This is SST's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Security Scan Tools CLI v0.9 -- HEAD")
	},
}
