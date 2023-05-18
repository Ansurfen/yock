package ypm

import (
	"fmt"
	"testing"
)

func TestModule(t *testing.T) {
	mod, err := CreateModule("module.json", "bash")
	if err != nil {
		panic(err)
	}
	fmt.Println(mod)
}
