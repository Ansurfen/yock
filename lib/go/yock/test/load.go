package test

import yocki "github.com/ansurfen/yock/interface"

type TestInterface interface {
	Hello()
}

func LoadTest(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("test")
	lib.SetField(map[string]any{
		"TestInterface": func(t TestInterface) {
			t.Hello()
		},
	})
}
