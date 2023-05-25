package cmd

import (
	"github.com/ansurfen/yock/util"
	"github.com/ansurfen/yock/ypm"
	"github.com/spf13/cobra"
)

type newCmdParameter struct {
	opt ypm.YpmNewOpt
}

var (
	newParameter newCmdParameter
	newCmd       = &cobra.Command{
		Use:   "new [module-name]",
		Short: `Initialize new module in current directory`,
		Long: `Init initializes and writes a new index.lua file in the current directory, in
		effect creating a new module rooted at the current directory. The index.lua file
		must not already exist.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || args[0][0] == '-' {
				util.Ycho.Fatal(util.ErrInvalidModuleName.Error())
			}
			newParameter.opt.Module = args[0]
			if err := ypm.New(newParameter.opt); err != nil {
				util.Ycho.Fatal(err.Error())
			}
		},
	}
)

func init() {
	yockCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().StringVarP(&newParameter.opt.Lang, "lang", "l", "", "")
	newCmd.PersistentFlags().StringVarP(&newParameter.opt.Version, "ver", "v", "v1", "")
	newCmd.PersistentFlags().BoolVarP(&newParameter.opt.CreateDir, "create", "c", false, "")
}
