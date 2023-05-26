package util

import (
	"os"
	"path"
	"path/filepath"

	"github.com/ansurfen/cushion/utils"
	"go.uber.org/zap"
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
	err = initLogger(LoggerOpt{
		Compress:    false,
		FileName:    "yock.log",
		Level:       "debug",
		FileMaxSize: 1024,
		Path:        Pathf("@/log"),
		Stdout:      true,
	})
	if err != nil {
		panic(err)
	}
	Ycho = zap.L()

	yockCpu = newCPU()
	yockMem = newMem()
	yockDisk = newDisk()
	yockHost = newHost()
	yockNet = newNet()
}
