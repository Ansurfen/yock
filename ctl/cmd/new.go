package cmd

import "github.com/spf13/cobra"

var newCmd = &cobra.Command{
	Use:   "new",
	Short: ``,
	Long:  ``,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	yockCmd.AddCommand(newCmd)
}
