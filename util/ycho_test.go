package util

import (
	"testing"
)

func TestLog(t *testing.T) {
	Ycho.Info("Hello World")
	Ycho.Warn("This is warn")
	Ycho.Error("This is error")
	Ycho.Fatal("This is fatal")
}
