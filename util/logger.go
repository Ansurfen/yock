package util

import (
	"fmt"

	"go.uber.org/zap"
)

func Ycho(call, msg string) string {
	return fmt.Sprintf("%s\t%s", call, msg)
}

func YchoInfo(call, msg string) {
	zap.S().Info(Ycho(call, msg))
}

func YchoWarn(call, msg string) {
	zap.S().Warn(Ycho(call, msg))
}

func YchoPanic(call, msg string) {
	zap.S().Panic(Ycho(call, msg))
}

func YchoFatal(call, msg string) {
	zap.S().Fatal(Ycho(call, msg))
}

func YchoDebug(call, msg string) {
	zap.S().Debug(Ycho(call, msg))
}
