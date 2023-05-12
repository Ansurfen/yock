package cmd

import "github.com/spf13/cobra"

var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	yockCmd.AddCommand(tidyCmd)
}
