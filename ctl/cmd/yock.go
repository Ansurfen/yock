package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var yockCmd = &cobra.Command{
	Use:   "yock",
	Short: "",
	Long:  ``,
}

func Execute() {
	err := yockCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
