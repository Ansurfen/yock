package yockp

import "testing"

func TestCC(t *testing.T) {
	yp := New()
	yp.CC("./test/c", "./test/c/main/main.lua")
}

func TestCCA(t *testing.T) {
	yp := New()
	yp.CC("./test/a", "./test/a/main.lua")
}
