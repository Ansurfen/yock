package util

import (
	"testing"

	"github.com/ansurfen/cushion/utils"
)

func init() {
	utils.InitLogger(utils.LoggerOpt{
		FileName: "yock.log",
		Path:     "./log",
		Stdout:   true,
	})
}

func TestYEcho(t *testing.T) {
	YchoInfo("c", "...")
	YchoInfo("", "err")
}
