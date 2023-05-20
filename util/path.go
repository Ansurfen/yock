package util

import (
	"os"
	"path"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
)

var (
	WorkSpace  string
	PluginPath string
	DriverPath string
	// executable file path
	YockPath string
)

func init() {
	WorkSpace = filepath.ToSlash(path.Join(utils.GetEnv().Workdir(), ".yock"))
	PluginPath = path.Join(WorkSpace, "plugin")
	DriverPath = path.Join(WorkSpace, "driver")
	exfPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	if YockBuild == "dev" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		exfPath = wd
	}
	YockPath = filepath.Join(exfPath, "..")
	utils.InitLogger(utils.LoggerOpt{
		FileName: "yock.log",
		Path:     Pathf("@/log"),
		Stdout:   true,
	})
}

// Pathf to format path
//
// @/abc => {WorkSpace}/abc (WorkSpace = UserHome + .yock)
//
// ~/abc => {YockPath}/abc (YockPath = executable file path)
func Pathf(path string) string {
	if len(path) > 0 {
		if path[0] == '@' {
			path = WorkSpace + path[1:]
		} else if path[0] == '~' {
			path = YockPath + path[1:]
		}
	}
	return path
}
