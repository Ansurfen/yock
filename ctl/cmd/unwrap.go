package cmd

import (
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	yockc "github.com/ansurfen/yock/cmd"
	"github.com/ansurfen/yock/util"
	"github.com/spf13/cobra"
)

var unwarpCmd = &cobra.Command{
	Use:   "unwrap [name]",
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.YchoFatal("", util.ErrArgsTooLittle.Error())
		}
		if err := utils.SafeMkdirs(util.Pathf("@/unmount")); err != nil {
			util.YchoFatal("", err.Error())
		}
		file := args[0]
		exf := ""
		switch utils.CurPlatform.OS {
		case "windows":
			exf = ".bat"
		default:
			exf = ".sh"
		}
		if err := yockc.Mv(yockc.MvOpt{}, filepath.Join(util.Pathf("@/mount"), file+exf),
			filepath.Join(util.Pathf("@/unmount"), file+exf)); err != nil {
			util.YchoFatal("", err.Error())
		}
	},
}

func init() {
	yockCmd.AddCommand(unwarpCmd)
}
