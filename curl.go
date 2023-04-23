package yock

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/ansurfen/cushion/utils"
)

type CurlCmd struct {
	urls   []string
	body   string
	method string
	o      string
	O      string
}

func NewCurlCmd() Cmd {
	return &CurlCmd{}
}

func (curl *CurlCmd) Exec(arg string) ([]byte, error) {
	initCmd(curl, arg, func(cli *flag.FlagSet, cc *CurlCmd) {
		cli.StringVar(&cc.body, "d", "", "")
		cli.StringVar(&cc.method, "x", "GET", "")
		cli.StringVar(&cc.O, "O", "", "")
		cli.StringVar(&cc.o, "o", ".", "")
	}, map[string]uint8{
		"-d": FlagString,
		"-x": FlagString,
		"-O": FlagString,
		"-o": FlagString,
	}, func(cc *CurlCmd, s string) error {
		if utils.IsURL(s) {
			cc.urls = append(cc.urls, s)
		}
		return nil
	})
	if len(curl.O) > 0 {
		if curl.o == "." {
			curl.o = path.Base(curl.O)
		}
		_, err := utils.FetchFile(curl.O, curl.o)
		return NilByte, err
	} else {
		for _, url := range curl.urls {
			req, err := http.NewRequest(curl.method, url, nil)
			if err != nil {
				return NilByte, fmt.Errorf("error creating request: %v", err)
			}
			switch strings.ToUpper(curl.method) {
			case "GET":
			case "POST":
				req.Method = "POST"
				req.Body = ioutil.NopCloser(strings.NewReader(curl.body))
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return NilByte, fmt.Errorf("error sending request: %v", err)
			}
			defer resp.Body.Close()

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return NilByte, fmt.Errorf("error reading response body: %v", err)
			}
			if len(bodyBytes) > 0 {
				return bodyBytes, nil
			}
		}
	}
	return NilByte, nil
}
