package client

import (
	"fmt"
	"testing"
)

func TestYocki(t *testing.T) {
	c := New("localhost", 9090)
	defer c.Close()
	fmt.Println(c.Call("SayHello", ""))
}
