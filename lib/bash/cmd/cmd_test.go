package cmd

import (
	"fmt"
	"os"
	"testing"

	"github.com/ansurfen/yock/util"
)

func init() {
	util.SafeBatchMkdirs([]string{"a", "b"})
}

func TestMove(t *testing.T) {
	mv := &MoveCmd{}
	out, err := mv.Exec("a b")
	fmt.Println(util.ConvertByte2String(out, util.GB18030), err)
	os.RemoveAll("b")
}

func TestCurl(t *testing.T) {
	curl := &CurlCmd{}
	out, err := curl.Exec("https://www.github.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func TestEcho(t *testing.T) {
	echo := &EchoCmd{}
	out, err := echo.Exec("$GOPATH")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}

func TestRm(t *testing.T) {
	rm := &RmCmd{}
	rm.Exec("-r a")
}
