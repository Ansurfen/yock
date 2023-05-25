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
	"sync"

	"github.com/ansurfen/cushion/utils"
	"github.com/ansurfen/yock/util"
)

// HttpOpt indicates configuration of HTTP request
type HttpOpt struct {
	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	Header map[string]string
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	Method string
	Data   string
	Cookie *http.Cookie
	// Save will write body into specify file
	Save bool
	// Dir set root directionary of file to be saved
	Dir string
	// Filename returns filename that will be saved according to url
	Filename func(string) string
	// Debug prints output when it's true
	Debug bool
	// Caller is used to mark parent caller of HTTP function
	//
	// It'll printed on console when debug is true
	Caller string
	// Strict will exit at once when error occur
	Strict bool

	// Async will enable goroutines when it's true.
	//
	// maxRequestCount and wg to limit count of concurrent goroutines
	Async           bool
	maxRequestCount chan struct{}
	wg              *sync.WaitGroup

	err error
}

// httpExceptionHandle controls return of HTTP function when value of return is true
func httpExceptionHandle(err error, opt *HttpOpt, exception error) bool {
	if err != nil {
		if opt.Debug {
			util.Ycho.Warn(fmt.Sprintf("%s\t%s", opt.Caller, exception.Error()))
		}
		if opt.Strict {
			return true
		} else {
			opt.err = ErrGeneral
		}
	}
	return false
}

// Http is similar with curl, which is used to send HTTP request according to opt and urls.
func HTTP(opt HttpOpt, urls []string) error {
	for _, url := range urls {
		if !utils.IsURL(url) && httpExceptionHandle(ErrGeneral, &opt, ErrInvalidURL) {
			return ErrInvalidURL
		}
		req, err := http.NewRequest("GET", url, nil)

		switch opt.Method = strings.ToUpper(opt.Method); opt.Method {
		case "": // default GET method
		case "GET", "HEAD", "PUT", "POST", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH":
			req.Method = opt.Method
		default:
			if httpExceptionHandle(ErrGeneral, &opt, ErrInvalidMethod) {
				return ErrInvalidMethod
			}
		}

		if httpExceptionHandle(err, &opt, ErrBadCreateFile) {
			return ErrBadCreateRequest
		}

		for k, v := range opt.Header {
			req.Header.Add(k, v)
		}
		if opt.Cookie != nil {
			req.AddCookie(opt.Cookie)
		}

		if len(opt.Data) != 0 {
			req.Body = ioutil.NopCloser(strings.NewReader(opt.Data))
		}

		if opt.Debug {
			util.Ycho.Info(fmt.Sprintf("%s\t%s", opt.Caller, fmt.Sprintf("%s %s", req.Method, url)))
		}

		if opt.Async {
			if opt.maxRequestCount == nil {
				opt.maxRequestCount = make(chan struct{}, 10)
				opt.wg = &sync.WaitGroup{}
			}

			opt.maxRequestCount <- struct{}{}

			u := url // to avoid loop variable url captured by func literal

			go func() {
				defer opt.wg.Done()

				res, err := http.DefaultClient.Do(req)

				if httpExceptionHandle(err, &opt, ErrBadSendRequest) {
					return
				}

				defer res.Body.Close()

				if opt.Save {
					dst := path.Join(opt.Dir, opt.Filename(u))
					dir := filepath.Dir(dst)

					if httpExceptionHandle(utils.SafeMkdirs(dir), &opt, ErrBadCreateDir) {
						return
					}

					file, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
					if err != nil {
						return
					}
					defer file.Close()
					io.Copy(file, res.Body)
				}
				<-opt.maxRequestCount
			}()
		} else {
			res, err := http.DefaultClient.Do(req)

			if httpExceptionHandle(err, &opt, ErrBadSendRequest) {
				return ErrBadSendRequest
			}

			defer res.Body.Close()

			if opt.Save {
				dst := path.Join(opt.Dir, opt.Filename(url))
				dir := filepath.Dir(dst)

				if httpExceptionHandle(utils.SafeMkdirs(dir), &opt, ErrBadCreateDir) {
					return ErrBadCreateDir
				}

				file, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
				if httpExceptionHandle(err, &opt, ErrBadCreateFile) {
					return ErrBadCreateFile
				}
				defer file.Close()
				io.Copy(file, res.Body)
			}
		}
	}

	if opt.Async {
		close(opt.maxRequestCount)
		opt.wg.Wait()
	}

	return opt.err
}
