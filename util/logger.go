package util

import (
	"fmt"

	"go.uber.org/zap"
)

func ychof(call, msg string) string {
	return fmt.Sprintf("%s\t%s", call, msg)
}

func YchoInfo(call, msg string) {
	zap.S().Info(ychof(call, msg))
}

func YchoWarn(call, msg string) {
	zap.S().Warn(ychof(call, msg))
}

func YchoPanic(call, msg string) {
	zap.S().Panic(ychof(call, msg))
}

func YchoFatal(call, msg string) {
	zap.S().Fatal(ychof(call, msg))
}

func YchoDebug(call, msg string) {
	zap.S().Debug(ychof(call, msg))
}
