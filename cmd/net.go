package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ansurfen/cushion/utils"
)

type HttpOpt struct {
	Data   string
	Method string
	Save   bool
	Dir    string
	Debug  bool
	Header map[string]string
	Fn     func(string) string
	Cookie *http.Cookie
}

func HTTP(opt HttpOpt, urls []string) error {
	for _, url := range urls {
		if opt.Debug {
			fmt.Print("curl: ", url)
		}
		if !utils.IsURL(url) {
			if opt.Debug {
				fmt.Println("\nerr url: ", url)
			}
			continue
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error creating request: %v", err)
		}
		for k, v := range opt.Header {
			req.Header.Add(k, v)
		}
		if opt.Cookie != nil {
			req.AddCookie(opt.Cookie)
		}
		switch strings.ToUpper(opt.Method) {
		case "GET":
		case "POST":
			req.Method = "POST"
			req.Body = ioutil.NopCloser(strings.NewReader(opt.Data))
		default:
			return fmt.Errorf("error method")
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error sending request: %v", err)
		}
		defer res.Body.Close()
		if opt.Save {
			dst := path.Join(opt.Dir, opt.Fn(url)+".lua")
			dir := filepath.Dir(dst)
			utils.SafeMkdirs(dir)
			file, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return fmt.Errorf("fail to create file")
			}
			defer file.Close()
			io.Copy(file, res.Body)
		}
		if opt.Debug {
			fmt.Println("✔")
		}
	}
	return nil
}
